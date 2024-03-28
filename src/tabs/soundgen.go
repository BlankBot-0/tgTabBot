package tabs

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/sinshu/go-meltysynth/meltysynth"
	"github.com/viert/go-lame"
	"io"
	"math"
	"time"
)

func TextToMp3(input string, sf *meltysynth.SoundFont) (bytes.Buffer, error) {
	// Set up processors - one midi processor in this case
	midiBuf := new(bytes.Buffer)
	midiProc := NewTabProcessorMidi(midiBuf, 480)

	// Parse text
	parseErrs := Parse(input, []TabProcessor{midiProc})
	if errs := errors.Join(parseErrs...); errs != nil {
		return bytes.Buffer{}, errs
	}

	// Extract processor's errors
	midiErrs := midiProc.Errors()

	// Write mp3 from midi
	mp3Output := bytes.Buffer{}
	if err := MidiToMp3(midiBuf, sf, &mp3Output); err != nil {
		return bytes.Buffer{}, err
	}

	return mp3Output, errors.Join(midiErrs...)
}

func MidiToMp3(midiInput io.Reader, soundFont *meltysynth.SoundFont, mp3Output io.Writer) error {
	pcm, err := MidiToPcm(soundFont, midiInput)
	if err != nil {
		return err
	}
	if err = PcmToMp3(&pcm, mp3Output); err != nil {
		return err
	}
	return nil
}

func PcmToMp3(pcmInput io.Reader, mp3Output io.Writer) error {
	enc := lame.NewEncoder(mp3Output)
	if err := enc.SetMode(2); err != nil {
		return fmt.Errorf("PcmToMp3 error: %w", err)
	}
	defer enc.Close()
	r := bufio.NewReader(pcmInput)
	if _, err := r.WriteTo(enc); err != nil {
		return fmt.Errorf("PcmToMp3 error: %w", err)
	}
	return nil
}

func MidiToPcm(soundFont *meltysynth.SoundFont, midiInput io.Reader) (bytes.Buffer, error) {
	// Create the synthesizer.
	settings := meltysynth.NewSynthesizerSettings(44100)
	synthesizer, err := meltysynth.NewSynthesizer(soundFont, settings)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("MidiToPcm error: %w", err)
	}

	// Load the MIDI data.
	midiFile, err := meltysynth.NewMidiFile(midiInput)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("MidiToPcm error: %w", err)
	}

	// Create the MIDI sequencer.
	sequencer := meltysynth.NewMidiFileSequencer(synthesizer)
	sequencer.Play(midiFile, true)

	// The output buffer.
	length := int(float64(settings.SampleRate) * float64(midiFile.GetLength()) / float64(time.Second))
	left := make([]float32, length)
	right := make([]float32, length)

	// Render the waveform.
	sequencer.Render(left, right)
	pcmOutput, err := writePCMInterleavedInt16(left, right)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("MidiToPcm error: %w", err)
	}
	return pcmOutput, nil
}

func writePCMInterleavedInt16(left []float32, right []float32) (bytes.Buffer, error) {
	length := len(left)
	var maxS float64

	for i := 0; i < length; i++ {
		absLeft := math.Abs(float64(left[i]))
		absRight := math.Abs(float64(right[i]))
		maxS = max(maxS, absLeft)
		maxS = max(maxS, absRight)
	}

	a := 32768 * float32(0.99/maxS)

	data := make([]int16, 2*length)

	for i := 0; i < length; i++ {
		data[2*i] = int16(a * left[i])
		data[2*i+1] = int16(a * right[i])
	}

	pcm := bytes.Buffer{}
	if err := binary.Write(&pcm, binary.LittleEndian, data); err != nil {
		return bytes.Buffer{}, err
	}
	return pcm, nil
}

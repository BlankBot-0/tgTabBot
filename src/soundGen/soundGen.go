package soundGen

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/sinshu/go-meltysynth/meltysynth"
	"github.com/viert/go-lame"
	"io"
	"math"
	"os"
	"time"
)

func MidiToMp3(midiInput io.Reader, sf2Path, mp3FileName string) error {
	if len(sf2Path) == 0 {
		return fmt.Errorf("missing sf2 path")
	}
	sf2, err := os.Open(sf2Path)
	if err != nil {
		return err
	}
	soundFont, err := meltysynth.NewSoundFont(sf2)
	sf2.Close()
	if err != nil {
		return err
	}

	pcm := bytes.Buffer{}
	MidiToPcm(soundFont, midiInput, &pcm)
	pcmToMp3(&pcm, mp3FileName)
	return nil
}

func pcmToMp3(pcmInput io.Reader, mp3FileName string) {
	mp3File, _ := os.OpenFile(mp3FileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	defer mp3File.Close()
	enc := lame.NewEncoder(mp3File)
	enc.SetMode(2)
	defer enc.Close()
	r := bufio.NewReader(pcmInput)
	r.WriteTo(enc)
}

func MidiToPcm(soundFont *meltysynth.SoundFont, midiInput io.Reader, PcmOutput io.Writer) error {
	// Create the synthesizer.
	settings := meltysynth.NewSynthesizerSettings(44100)
	synthesizer, err := meltysynth.NewSynthesizer(soundFont, settings)
	if err != nil {
		return err
	}

	// Load the MIDI data.
	midiFile, err := meltysynth.NewMidiFile(midiInput)
	if err != nil {
		return err
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

	return writePCMInterleavedInt16(left, right, PcmOutput)
}

func writePCMInterleavedInt16(left []float32, right []float32, pcm io.Writer) error {
	length := len(left)
	var maxS float64

	for i := 0; i < length; i++ {
		absLeft := math.Abs(float64(left[i]))
		absRight := math.Abs(float64(right[i]))
		if maxS < absLeft {
			maxS = absLeft
		}
		if maxS < absRight {
			maxS = absRight
		}
	}

	a := 32768 * float32(0.99/maxS)

	data := make([]int16, 2*length)

	for i := 0; i < length; i++ {
		data[2*i] = int16(a * left[i])
		data[2*i+1] = int16(a * right[i])
	}

	return binary.Write(pcm, binary.LittleEndian, data)
}

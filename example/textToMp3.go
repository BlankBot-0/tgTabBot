package main

import (
	"bytes"
	"fmt"
	"github.com/sinshu/go-meltysynth/meltysynth"
	"tgScoreBot/src/TabProcessor/midi"
	"tgScoreBot/src/soundGen"
	"tgScoreBot/src/textParser"
)

func TextToMp3(input string, sf *meltysynth.SoundFont) (bytes.Buffer, error) {
	// Set up processors - one midi processor in this case
	midiBuf := new(bytes.Buffer)
	midiProc := midi.NewTabProcessorMidi(midiBuf, 480)

	// Parse text
	parseErrs := textParser.Parse(input, []textParser.TabProcessor{midiProc})

	// Extract processors' errors
	midiErrs := midiProc.Errors()

	// Write mp3 from midi
	mp3Output := bytes.Buffer{}
	if err := soundGen.MidiToMp3(midiBuf, sf, &mp3Output); err != nil {
		return bytes.Buffer{}, err
	}

	// Print errors
	fmt.Println("Parse errors:")
	for _, e := range parseErrs {
		fmt.Println(e)
	}
	fmt.Println("\nMidi errors:")
	for _, e := range midiErrs {
		fmt.Println(e)
	}

	return mp3Output, nil
}

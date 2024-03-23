package main

import (
	"bytes"
	"fmt"
	"github.com/sinshu/go-meltysynth/meltysynth"
	"os"
	"tgScoreBot/src/TabProcessor/midi"
	"tgScoreBot/src/soundGen"
	"tgScoreBot/src/textParser"
)

func TextToMp3(input, mp3Path string) error {
	// Open necessary files
	mp3File, err := os.OpenFile(mp3Path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer mp3File.Close()

	sfFile, err := os.Open("TimGM6mb.sf2")
	if err != nil {
		return err
	}
	defer sfFile.Close()

	sf, err := meltysynth.NewSoundFont(sfFile)
	if err != nil {
		return err
	}

	// Set up processors - one midi processor in this case
	midiBuf := new(bytes.Buffer)
	midiProc := midi.NewTabProcessorMidi(midiBuf, 480)

	// Parse text
	parseErrs := textParser.Parse(input, []*textParser.TabProcessor{&midiProc})

	// Extract processors' errors
	midiErrs := midiProc.Errors()

	// Write mp3 from midi
	if err = soundGen.MidiToMp3(midiBuf, sf, mp3File); err != nil {
		return err
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

	return nil
}

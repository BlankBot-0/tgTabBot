package main

import (
	"bytes"
	"errors"
	"github.com/sinshu/go-meltysynth/meltysynth"
	"tgScoreBot/src/tabs"
)

func TextToMp3(input string, sf *meltysynth.SoundFont) (bytes.Buffer, error) {
	// Set up processors - one midi processor in this case
	midiBuf := new(bytes.Buffer)
	midiProc := tabs.NewTabProcessorMidi(midiBuf, 480)

	// Parse text
	parseErrs := tabs.Parse(input, []tabs.TabProcessor{midiProc})
	if errs := errors.Join(parseErrs...); errs != nil {
		return bytes.Buffer{}, errs
	}

	// Extract processor's errors
	midiErrs := midiProc.Errors()

	// Write mp3 from midi
	mp3Output := bytes.Buffer{}
	if err := tabs.MidiToMp3(midiBuf, sf, &mp3Output); err != nil {
		return bytes.Buffer{}, err
	}

	return mp3Output, errors.Join(midiErrs...)
}

package main

import (
	"tgScoreBot/src/soundGen"
	"tgScoreBot/src/textParser"
)

func main() {
	textInput := "BPM 120\nE2 A2 D3 G3 B3 E4 \n4([65 57]) 16(65 65 65 65) 4([65 57]) 16(65 65 65 65)"
	midi := textParser.TabToMidi(textInput)
	soundGen.MidiToMp3(midi, "TimGM6mb.sf2", "testOutput.mp3")
}

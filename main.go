package main

import (
	"github.com/sinshu/go-meltysynth/meltysynth"
	"os"
	"tgScoreBot/src/telegram"
)

func main() {
	cfg := telegram.MustLoad()

	sfFile, _ := os.Open("TimGM6mb.sf2")
	sf, _ := meltysynth.NewSoundFont(sfFile)
	sfFile.Close()

	bot := telegram.NewTgProcessor(cfg, sf)
	bot.Start()
}

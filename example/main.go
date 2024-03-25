package main

import (
	"github.com/sinshu/go-meltysynth/meltysynth"
	"os"
)

func main() {
	// get mp3 files from text files
	example1, _ := os.ReadFile("./example/textExample1")
	example2, _ := os.ReadFile("./example/textExampleBend")
	example3, _ := os.ReadFile("./example/textExampleChord")
	sfFile, _ := os.Open("TimGM6mb.sf2")
	defer sfFile.Close()
	sf, _ := meltysynth.NewSoundFont(sfFile)
	mp3File1, _ := os.OpenFile("mp3Example1.mp3", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	defer mp3File1.Close()
	mp3File2, _ := os.OpenFile("mp3ExampleBend.mp3", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	defer mp3File2.Close()
	mp3File3, _ := os.OpenFile("mp3ExampleChord.mp3", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	defer mp3File3.Close()
	buf1, _ := TextToMp3(string(example1), sf)
	buf2, _ := TextToMp3(string(example2), sf)
	buf3, _ := TextToMp3(string(example3), sf)
	mp3File1.Write(buf1.Bytes())
	mp3File2.Write(buf2.Bytes())
	mp3File3.Write(buf3.Bytes())
}

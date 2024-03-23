package main

import (
	"os"
)

func main() {
	// get mp3 files from text files
	example1, _ := os.ReadFile("./example/textExample1")
	example2, _ := os.ReadFile("./example/textExampleBend")
	TextToMp3(string(example1), "mp3Example1.mp3")
	TextToMp3(string(example2), "mp3ExampleBend.mp3")
}

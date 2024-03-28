package telegram

const (
	HelpCmd         = "/help"
	StartCmd        = "/start"
	ExampleCmd      = "/example"
	InstructionsCmd = "/instructions"
)

const msgHelp = `I can convert your tabs to music if they follow these simple rules:

First, you must specify bpm and guitar tuning like this (currently only 6-string tabs are supported):
BPM 120 E2 A2 D3 G3 B3 E4
The rest of the message is build from blocks of the following form:
'duration'('some tabs of said duration')

You can see /example where all current features are demonstrated or get detailed instructions with /instructions
Don't feel intimidated - it's easier to write than to read!
Currently I support dotted and/or tuplet durations, simple polyphony and bends.
`

const msgError = "There is an error somewhere in the input 🤔. Use /help for instructions"

const msgHello = "Hi there! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command 🤔\nUse /help for instructions"
	msgExample        = `Send this to me to see how I work:
BPM 80 E2 A2 D3 G3 B3 E4

32(116 216 116 118) 16(b(118, 1)) 32(118) 16(b(119, 2)) 16(119)
32(118 119 118 116 219 116 119 116 118 119 118 116)
32(219 116) 16._4(b(219, 2)) 32(116) 16._4(b(219, 2)) 32(116) 4(b(219, 2))`
	msgInstructions = "placeholder"
)

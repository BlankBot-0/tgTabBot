package textParser

import (
	"bytes"
	"github.com/antlr4-go/antlr/v4"
	"tgScoreBot/src/textParser/parser"
)

func Tab(input string) *bytes.Buffer {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewTabLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTabParser(stream)

	// Finally parse the expression (by walking the tree)
	output := new(bytes.Buffer)
	listener := newTabListener(output, 480)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start_())

	if err := listener.enc.Write(); err != nil {
		panic(err)
	}

	return output
}

//func TestTabListener(t *testing.T) {
//	input := `BPM: 120
//   			  E2 A2 D3 G3 B3 E4
//			  4([65 57]) 16(65 65 65 65) 4([65 57]) 16(65 65 65 65)`
//}

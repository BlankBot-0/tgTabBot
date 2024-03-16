package textParser

import (
	"github.com/antlr4-go/antlr/v4"
	"io"
	"tgScoreBot/src/textParser/parser"
)

func TabToMidi(input string, output io.Writer) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewTabLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTabParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := newTabListener(output, 480)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start_())

	if err := listener.enc.Write(); err != nil {
		panic(err)
	}
}

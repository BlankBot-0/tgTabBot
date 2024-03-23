package textParser

import (
	"github.com/antlr4-go/antlr/v4"
	"tgScoreBot/src/textParser/parser"
)

// TabToMidi writes midi to output given a string input which satisfies the grammar Tab.g4
func Parse(input string, processors []*TabProcessor) []error {
	// Set up the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewTabLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTabParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := newTabListener(processors)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start_())

	return listener.errs
}

// Code generated from Tab.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TabLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TabLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tablexerLexerInit() {
	staticData := &TabLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'BPM'", "'('", "')'", "'['", "']'", "'p'", "'b('", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "INTEGER", "DURMOD", "NOTELETTER",
		"WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "INTEGER",
		"DURMOD", "NOTELETTER", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 68, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		4, 8, 46, 8, 8, 11, 8, 12, 8, 47, 1, 9, 4, 9, 51, 8, 9, 11, 9, 12, 9, 52,
		1, 10, 1, 10, 1, 10, 3, 10, 58, 8, 10, 3, 10, 60, 8, 10, 1, 11, 4, 11,
		63, 8, 11, 11, 11, 12, 11, 64, 1, 11, 1, 11, 0, 0, 12, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 1, 0, 5,
		1, 0, 48, 57, 2, 0, 46, 46, 95, 95, 1, 0, 65, 71, 1, 0, 97, 103, 3, 0,
		9, 10, 13, 13, 32, 32, 72, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 29, 1, 0, 0, 0,
		5, 31, 1, 0, 0, 0, 7, 33, 1, 0, 0, 0, 9, 35, 1, 0, 0, 0, 11, 37, 1, 0,
		0, 0, 13, 39, 1, 0, 0, 0, 15, 42, 1, 0, 0, 0, 17, 45, 1, 0, 0, 0, 19, 50,
		1, 0, 0, 0, 21, 59, 1, 0, 0, 0, 23, 62, 1, 0, 0, 0, 25, 26, 5, 66, 0, 0,
		26, 27, 5, 80, 0, 0, 27, 28, 5, 77, 0, 0, 28, 2, 1, 0, 0, 0, 29, 30, 5,
		40, 0, 0, 30, 4, 1, 0, 0, 0, 31, 32, 5, 41, 0, 0, 32, 6, 1, 0, 0, 0, 33,
		34, 5, 91, 0, 0, 34, 8, 1, 0, 0, 0, 35, 36, 5, 93, 0, 0, 36, 10, 1, 0,
		0, 0, 37, 38, 5, 112, 0, 0, 38, 12, 1, 0, 0, 0, 39, 40, 5, 98, 0, 0, 40,
		41, 5, 40, 0, 0, 41, 14, 1, 0, 0, 0, 42, 43, 5, 44, 0, 0, 43, 16, 1, 0,
		0, 0, 44, 46, 7, 0, 0, 0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 45,
		1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 18, 1, 0, 0, 0, 49, 51, 7, 1, 0, 0,
		50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1,
		0, 0, 0, 53, 20, 1, 0, 0, 0, 54, 60, 7, 2, 0, 0, 55, 57, 7, 3, 0, 0, 56,
		58, 5, 35, 0, 0, 57, 56, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0,
		0, 0, 59, 54, 1, 0, 0, 0, 59, 55, 1, 0, 0, 0, 60, 22, 1, 0, 0, 0, 61, 63,
		7, 4, 0, 0, 62, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0,
		64, 65, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 6, 11, 0, 0, 67, 24, 1,
		0, 0, 0, 6, 0, 47, 52, 57, 59, 64, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TabLexerInit initializes any static state used to implement TabLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTabLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TabLexerInit() {
	staticData := &TabLexerLexerStaticData
	staticData.once.Do(tablexerLexerInit)
}

// NewTabLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTabLexer(input antlr.CharStream) *TabLexer {
	TabLexerInit()
	l := new(TabLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TabLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Tab.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TabLexer tokens.
const (
	TabLexerT__0       = 1
	TabLexerT__1       = 2
	TabLexerT__2       = 3
	TabLexerT__3       = 4
	TabLexerT__4       = 5
	TabLexerT__5       = 6
	TabLexerT__6       = 7
	TabLexerT__7       = 8
	TabLexerINTEGER    = 9
	TabLexerDURMOD     = 10
	TabLexerNOTELETTER = 11
	TabLexerWHITESPACE = 12
)

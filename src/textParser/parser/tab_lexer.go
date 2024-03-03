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
		"", "'('", "')'", "'['", "']'", "'p'", "'b('", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "STR", "INTEGER", "DURMOD", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "STR", "INTEGER",
		"DURMOD", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 57, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 42, 8, 8, 11, 8,
		12, 8, 43, 1, 9, 4, 9, 47, 8, 9, 11, 9, 12, 9, 48, 1, 10, 4, 10, 52, 8,
		10, 11, 10, 12, 10, 53, 1, 10, 1, 10, 0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 1, 0, 4, 1, 0, 49, 57,
		1, 0, 48, 57, 2, 0, 46, 46, 95, 95, 3, 0, 9, 10, 13, 13, 32, 32, 59, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 1, 23, 1, 0, 0,
		0, 3, 25, 1, 0, 0, 0, 5, 27, 1, 0, 0, 0, 7, 29, 1, 0, 0, 0, 9, 31, 1, 0,
		0, 0, 11, 33, 1, 0, 0, 0, 13, 36, 1, 0, 0, 0, 15, 38, 1, 0, 0, 0, 17, 41,
		1, 0, 0, 0, 19, 46, 1, 0, 0, 0, 21, 51, 1, 0, 0, 0, 23, 24, 5, 40, 0, 0,
		24, 2, 1, 0, 0, 0, 25, 26, 5, 41, 0, 0, 26, 4, 1, 0, 0, 0, 27, 28, 5, 91,
		0, 0, 28, 6, 1, 0, 0, 0, 29, 30, 5, 93, 0, 0, 30, 8, 1, 0, 0, 0, 31, 32,
		5, 112, 0, 0, 32, 10, 1, 0, 0, 0, 33, 34, 5, 98, 0, 0, 34, 35, 5, 40, 0,
		0, 35, 12, 1, 0, 0, 0, 36, 37, 5, 44, 0, 0, 37, 14, 1, 0, 0, 0, 38, 39,
		7, 0, 0, 0, 39, 16, 1, 0, 0, 0, 40, 42, 7, 1, 0, 0, 41, 40, 1, 0, 0, 0,
		42, 43, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 18, 1,
		0, 0, 0, 45, 47, 7, 2, 0, 0, 46, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48,
		46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 20, 1, 0, 0, 0, 50, 52, 7, 3, 0,
		0, 51, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 6, 10, 0, 0, 56, 22, 1, 0, 0, 0,
		4, 0, 43, 48, 53, 1, 6, 0, 0,
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
	TabLexerSTR        = 8
	TabLexerINTEGER    = 9
	TabLexerDURMOD     = 10
	TabLexerWHITESPACE = 11
)

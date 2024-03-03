// Code generated from Tab.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tab

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TabParser struct {
	*antlr.BaseParser
}

var TabParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tabParserInit() {
	staticData := &TabParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'['", "']'", "'p'", "'b('", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "STR", "INTEGER", "DURMOD", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"start", "durationGroups", "durationGroup", "duration", "noteGroup",
		"chord", "note",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 59, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 22, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		33, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 44,
		8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 57, 8, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 57, 0, 14,
		1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 23, 1, 0, 0, 0, 6, 32, 1, 0, 0, 0, 8,
		43, 1, 0, 0, 0, 10, 45, 1, 0, 0, 0, 12, 56, 1, 0, 0, 0, 14, 15, 3, 2, 1,
		0, 15, 16, 5, 0, 0, 1, 16, 1, 1, 0, 0, 0, 17, 18, 3, 4, 2, 0, 18, 19, 3,
		2, 1, 0, 19, 22, 1, 0, 0, 0, 20, 22, 3, 4, 2, 0, 21, 17, 1, 0, 0, 0, 21,
		20, 1, 0, 0, 0, 22, 3, 1, 0, 0, 0, 23, 24, 3, 6, 3, 0, 24, 25, 5, 1, 0,
		0, 25, 26, 3, 8, 4, 0, 26, 27, 5, 2, 0, 0, 27, 5, 1, 0, 0, 0, 28, 29, 5,
		9, 0, 0, 29, 30, 5, 10, 0, 0, 30, 33, 5, 9, 0, 0, 31, 33, 5, 9, 0, 0, 32,
		28, 1, 0, 0, 0, 32, 31, 1, 0, 0, 0, 33, 7, 1, 0, 0, 0, 34, 35, 3, 12, 6,
		0, 35, 36, 3, 8, 4, 0, 36, 44, 1, 0, 0, 0, 37, 38, 5, 3, 0, 0, 38, 39,
		3, 10, 5, 0, 39, 40, 5, 4, 0, 0, 40, 41, 3, 8, 4, 0, 41, 44, 1, 0, 0, 0,
		42, 44, 3, 12, 6, 0, 43, 34, 1, 0, 0, 0, 43, 37, 1, 0, 0, 0, 43, 42, 1,
		0, 0, 0, 44, 9, 1, 0, 0, 0, 45, 46, 3, 8, 4, 0, 46, 11, 1, 0, 0, 0, 47,
		48, 5, 8, 0, 0, 48, 57, 5, 9, 0, 0, 49, 57, 5, 5, 0, 0, 50, 51, 5, 6, 0,
		0, 51, 52, 3, 12, 6, 0, 52, 53, 5, 7, 0, 0, 53, 54, 5, 9, 0, 0, 54, 55,
		5, 2, 0, 0, 55, 57, 1, 0, 0, 0, 56, 47, 1, 0, 0, 0, 56, 49, 1, 0, 0, 0,
		56, 50, 1, 0, 0, 0, 57, 13, 1, 0, 0, 0, 4, 21, 32, 43, 56,
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

// TabParserInit initializes any static state used to implement TabParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTabParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TabParserInit() {
	staticData := &TabParserStaticData
	staticData.once.Do(tabParserInit)
}

// NewTabParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTabParser(input antlr.TokenStream) *TabParser {
	TabParserInit()
	this := new(TabParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TabParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tab.g4"

	return this
}

// TabParser tokens.
const (
	TabParserEOF        = antlr.TokenEOF
	TabParserT__0       = 1
	TabParserT__1       = 2
	TabParserT__2       = 3
	TabParserT__3       = 4
	TabParserT__4       = 5
	TabParserT__5       = 6
	TabParserT__6       = 7
	TabParserSTR        = 8
	TabParserINTEGER    = 9
	TabParserDURMOD     = 10
	TabParserWHITESPACE = 11
)

// TabParser rules.
const (
	TabParserRULE_start          = 0
	TabParserRULE_durationGroups = 1
	TabParserRULE_durationGroup  = 2
	TabParserRULE_duration       = 3
	TabParserRULE_noteGroup      = 4
	TabParserRULE_chord          = 5
	TabParserRULE_note           = 6
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DurationGroups() IDurationGroupsContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) DurationGroups() IDurationGroupsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationGroupsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationGroupsContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TabParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *TabParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TabParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.DurationGroups()
	}
	{
		p.SetState(15)
		p.Match(TabParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDurationGroupsContext is an interface to support dynamic dispatch.
type IDurationGroupsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DurationGroup() IDurationGroupContext
	DurationGroups() IDurationGroupsContext

	// IsDurationGroupsContext differentiates from other interfaces.
	IsDurationGroupsContext()
}

type DurationGroupsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationGroupsContext() *DurationGroupsContext {
	var p = new(DurationGroupsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_durationGroups
	return p
}

func InitEmptyDurationGroupsContext(p *DurationGroupsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_durationGroups
}

func (*DurationGroupsContext) IsDurationGroupsContext() {}

func NewDurationGroupsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationGroupsContext {
	var p = new(DurationGroupsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_durationGroups

	return p
}

func (s *DurationGroupsContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationGroupsContext) DurationGroup() IDurationGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationGroupContext)
}

func (s *DurationGroupsContext) DurationGroups() IDurationGroupsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationGroupsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationGroupsContext)
}

func (s *DurationGroupsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationGroupsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationGroupsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterDurationGroups(s)
	}
}

func (s *DurationGroupsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitDurationGroups(s)
	}
}

func (p *TabParser) DurationGroups() (localctx IDurationGroupsContext) {
	localctx = NewDurationGroupsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TabParserRULE_durationGroups)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.DurationGroup()
		}
		{
			p.SetState(18)
			p.DurationGroups()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.DurationGroup()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDurationGroupContext is an interface to support dynamic dispatch.
type IDurationGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Duration() IDurationContext
	NoteGroup() INoteGroupContext

	// IsDurationGroupContext differentiates from other interfaces.
	IsDurationGroupContext()
}

type DurationGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationGroupContext() *DurationGroupContext {
	var p = new(DurationGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_durationGroup
	return p
}

func InitEmptyDurationGroupContext(p *DurationGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_durationGroup
}

func (*DurationGroupContext) IsDurationGroupContext() {}

func NewDurationGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationGroupContext {
	var p = new(DurationGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_durationGroup

	return p
}

func (s *DurationGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationGroupContext) Duration() IDurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
}

func (s *DurationGroupContext) NoteGroup() INoteGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoteGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoteGroupContext)
}

func (s *DurationGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterDurationGroup(s)
	}
}

func (s *DurationGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitDurationGroup(s)
	}
}

func (p *TabParser) DurationGroup() (localctx IDurationGroupContext) {
	localctx = NewDurationGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TabParserRULE_durationGroup)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Duration()
	}
	{
		p.SetState(24)
		p.Match(TabParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)
		p.NoteGroup()
	}
	{
		p.SetState(26)
		p.Match(TabParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDurationContext is an interface to support dynamic dispatch.
type IDurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDurationContext differentiates from other interfaces.
	IsDurationContext()
}

type DurationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationContext() *DurationContext {
	var p = new(DurationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_duration
	return p
}

func InitEmptyDurationContext(p *DurationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_duration
}

func (*DurationContext) IsDurationContext() {}

func NewDurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationContext {
	var p = new(DurationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_duration

	return p
}

func (s *DurationContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationContext) CopyAll(ctx *DurationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DurDescribedContext struct {
	DurationContext
}

func NewDurDescribedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DurDescribedContext {
	var p = new(DurDescribedContext)

	InitEmptyDurationContext(&p.DurationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DurationContext))

	return p
}

func (s *DurDescribedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurDescribedContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(TabParserINTEGER)
}

func (s *DurDescribedContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(TabParserINTEGER, i)
}

func (s *DurDescribedContext) DURMOD() antlr.TerminalNode {
	return s.GetToken(TabParserDURMOD, 0)
}

func (s *DurDescribedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterDurDescribed(s)
	}
}

func (s *DurDescribedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitDurDescribed(s)
	}
}

type DurDefaultContext struct {
	DurationContext
}

func NewDurDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DurDefaultContext {
	var p = new(DurDefaultContext)

	InitEmptyDurationContext(&p.DurationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DurationContext))

	return p
}

func (s *DurDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurDefaultContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TabParserINTEGER, 0)
}

func (s *DurDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterDurDefault(s)
	}
}

func (s *DurDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitDurDefault(s)
	}
}

func (p *TabParser) Duration() (localctx IDurationContext) {
	localctx = NewDurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TabParserRULE_duration)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDurDescribedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Match(TabParserDURMOD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDurDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INoteGroupContext is an interface to support dynamic dispatch.
type INoteGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Note() INoteContext
	NoteGroup() INoteGroupContext
	Chord() IChordContext

	// IsNoteGroupContext differentiates from other interfaces.
	IsNoteGroupContext()
}

type NoteGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoteGroupContext() *NoteGroupContext {
	var p = new(NoteGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_noteGroup
	return p
}

func InitEmptyNoteGroupContext(p *NoteGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_noteGroup
}

func (*NoteGroupContext) IsNoteGroupContext() {}

func NewNoteGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteGroupContext {
	var p = new(NoteGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_noteGroup

	return p
}

func (s *NoteGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteGroupContext) Note() INoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoteContext)
}

func (s *NoteGroupContext) NoteGroup() INoteGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoteGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoteGroupContext)
}

func (s *NoteGroupContext) Chord() IChordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChordContext)
}

func (s *NoteGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterNoteGroup(s)
	}
}

func (s *NoteGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitNoteGroup(s)
	}
}

func (p *TabParser) NoteGroup() (localctx INoteGroupContext) {
	localctx = NewNoteGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TabParserRULE_noteGroup)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Note()
		}
		{
			p.SetState(35)
			p.NoteGroup()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(TabParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Chord()
		}
		{
			p.SetState(39)
			p.Match(TabParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.NoteGroup()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.Note()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChordContext is an interface to support dynamic dispatch.
type IChordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsChordContext differentiates from other interfaces.
	IsChordContext()
}

type ChordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChordContext() *ChordContext {
	var p = new(ChordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_chord
	return p
}

func InitEmptyChordContext(p *ChordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_chord
}

func (*ChordContext) IsChordContext() {}

func NewChordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChordContext {
	var p = new(ChordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_chord

	return p
}

func (s *ChordContext) GetParser() antlr.Parser { return s.parser }

func (s *ChordContext) CopyAll(ctx *ChordContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ChordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleChordContext struct {
	ChordContext
}

func NewSimpleChordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleChordContext {
	var p = new(SimpleChordContext)

	InitEmptyChordContext(&p.ChordContext)
	p.parser = parser
	p.CopyAll(ctx.(*ChordContext))

	return p
}

func (s *SimpleChordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleChordContext) NoteGroup() INoteGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoteGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoteGroupContext)
}

func (s *SimpleChordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterSimpleChord(s)
	}
}

func (s *SimpleChordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitSimpleChord(s)
	}
}

func (p *TabParser) Chord() (localctx IChordContext) {
	localctx = NewChordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TabParserRULE_chord)
	localctx = NewSimpleChordContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.NoteGroup()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INoteContext is an interface to support dynamic dispatch.
type INoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNoteContext differentiates from other interfaces.
	IsNoteContext()
}

type NoteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoteContext() *NoteContext {
	var p = new(NoteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_note
	return p
}

func InitEmptyNoteContext(p *NoteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_note
}

func (*NoteContext) IsNoteContext() {}

func NewNoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteContext {
	var p = new(NoteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_note

	return p
}

func (s *NoteContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteContext) CopyAll(ctx *NoteContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *NoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PlayFretContext struct {
	NoteContext
}

func NewPlayFretContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlayFretContext {
	var p = new(PlayFretContext)

	InitEmptyNoteContext(&p.NoteContext)
	p.parser = parser
	p.CopyAll(ctx.(*NoteContext))

	return p
}

func (s *PlayFretContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlayFretContext) STR() antlr.TerminalNode {
	return s.GetToken(TabParserSTR, 0)
}

func (s *PlayFretContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TabParserINTEGER, 0)
}

func (s *PlayFretContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterPlayFret(s)
	}
}

func (s *PlayFretContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitPlayFret(s)
	}
}

type PauseContext struct {
	NoteContext
}

func NewPauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PauseContext {
	var p = new(PauseContext)

	InitEmptyNoteContext(&p.NoteContext)
	p.parser = parser
	p.CopyAll(ctx.(*NoteContext))

	return p
}

func (s *PauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterPause(s)
	}
}

func (s *PauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitPause(s)
	}
}

type BendContext struct {
	NoteContext
}

func NewBendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BendContext {
	var p = new(BendContext)

	InitEmptyNoteContext(&p.NoteContext)
	p.parser = parser
	p.CopyAll(ctx.(*NoteContext))

	return p
}

func (s *BendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BendContext) Note() INoteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoteContext)
}

func (s *BendContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TabParserINTEGER, 0)
}

func (s *BendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterBend(s)
	}
}

func (s *BendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitBend(s)
	}
}

func (p *TabParser) Note() (localctx INoteContext) {
	localctx = NewNoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TabParserRULE_note)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TabParserSTR:
		localctx = NewPlayFretContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(TabParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TabParserT__4:
		localctx = NewPauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(TabParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TabParserT__5:
		localctx = NewBendContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(TabParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Note()
		}
		{
			p.SetState(52)
			p.Match(TabParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(TabParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

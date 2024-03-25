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
		"", "'BPM'", "'('", "')'", "'['", "']'", "'p'", "'b('", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "INTEGER", "DURMOD", "NOTELETTER",
		"WHITESPACE",
	}
	staticData.RuleNames = []string{
		"start", "bpm", "tuning", "tuningImpl", "durationGroup", "duration",
		"entry",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 76, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 41, 8, 4, 10, 4,
		12, 4, 44, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 51, 8, 5, 10, 5, 12,
		5, 54, 9, 5, 1, 5, 3, 5, 57, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 62, 8, 6, 10,
		6, 12, 6, 65, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 74,
		8, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 76, 0, 14, 1, 0, 0, 0,
		2, 24, 1, 0, 0, 0, 4, 27, 1, 0, 0, 0, 6, 34, 1, 0, 0, 0, 8, 37, 1, 0, 0,
		0, 10, 56, 1, 0, 0, 0, 12, 73, 1, 0, 0, 0, 14, 15, 3, 2, 1, 0, 15, 19,
		3, 4, 2, 0, 16, 18, 3, 8, 4, 0, 17, 16, 1, 0, 0, 0, 18, 21, 1, 0, 0, 0,
		19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 22, 1, 0, 0, 0, 21, 19, 1,
		0, 0, 0, 22, 23, 5, 0, 0, 1, 23, 1, 1, 0, 0, 0, 24, 25, 5, 1, 0, 0, 25,
		26, 5, 9, 0, 0, 26, 3, 1, 0, 0, 0, 27, 28, 3, 6, 3, 0, 28, 29, 3, 6, 3,
		0, 29, 30, 3, 6, 3, 0, 30, 31, 3, 6, 3, 0, 31, 32, 3, 6, 3, 0, 32, 33,
		3, 6, 3, 0, 33, 5, 1, 0, 0, 0, 34, 35, 5, 11, 0, 0, 35, 36, 5, 9, 0, 0,
		36, 7, 1, 0, 0, 0, 37, 38, 3, 10, 5, 0, 38, 42, 5, 2, 0, 0, 39, 41, 3,
		12, 6, 0, 40, 39, 1, 0, 0, 0, 41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42,
		43, 1, 0, 0, 0, 43, 45, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 46, 5, 3, 0,
		0, 46, 9, 1, 0, 0, 0, 47, 48, 5, 9, 0, 0, 48, 52, 5, 10, 0, 0, 49, 51,
		5, 9, 0, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0,
		52, 53, 1, 0, 0, 0, 53, 57, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 57, 5,
		9, 0, 0, 56, 47, 1, 0, 0, 0, 56, 55, 1, 0, 0, 0, 57, 11, 1, 0, 0, 0, 58,
		74, 5, 9, 0, 0, 59, 63, 5, 4, 0, 0, 60, 62, 5, 9, 0, 0, 61, 60, 1, 0, 0,
		0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 66,
		1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 74, 5, 5, 0, 0, 67, 74, 5, 6, 0, 0,
		68, 69, 5, 7, 0, 0, 69, 70, 5, 9, 0, 0, 70, 71, 5, 8, 0, 0, 71, 72, 5,
		9, 0, 0, 72, 74, 5, 3, 0, 0, 73, 58, 1, 0, 0, 0, 73, 59, 1, 0, 0, 0, 73,
		67, 1, 0, 0, 0, 73, 68, 1, 0, 0, 0, 74, 13, 1, 0, 0, 0, 6, 19, 42, 52,
		56, 63, 73,
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
	TabParserT__7       = 8
	TabParserINTEGER    = 9
	TabParserDURMOD     = 10
	TabParserNOTELETTER = 11
	TabParserWHITESPACE = 12
)

// TabParser rules.
const (
	TabParserRULE_start         = 0
	TabParserRULE_bpm           = 1
	TabParserRULE_tuning        = 2
	TabParserRULE_tuningImpl    = 3
	TabParserRULE_durationGroup = 4
	TabParserRULE_duration      = 5
	TabParserRULE_entry         = 6
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Bpm() IBpmContext
	Tuning() ITuningContext
	EOF() antlr.TerminalNode
	AllDurationGroup() []IDurationGroupContext
	DurationGroup(i int) IDurationGroupContext

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

func (s *StartContext) Bpm() IBpmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBpmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBpmContext)
}

func (s *StartContext) Tuning() ITuningContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITuningContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITuningContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TabParserEOF, 0)
}

func (s *StartContext) AllDurationGroup() []IDurationGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDurationGroupContext); ok {
			len++
		}
	}

	tst := make([]IDurationGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDurationGroupContext); ok {
			tst[i] = t.(IDurationGroupContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) DurationGroup(i int) IDurationGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationGroupContext)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Bpm()
	}
	{
		p.SetState(15)
		p.Tuning()
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TabParserINTEGER {
		{
			p.SetState(16)
			p.DurationGroup()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(22)
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

// IBpmContext is an interface to support dynamic dispatch.
type IBpmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode

	// IsBpmContext differentiates from other interfaces.
	IsBpmContext()
}

type BpmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBpmContext() *BpmContext {
	var p = new(BpmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_bpm
	return p
}

func InitEmptyBpmContext(p *BpmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_bpm
}

func (*BpmContext) IsBpmContext() {}

func NewBpmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BpmContext {
	var p = new(BpmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_bpm

	return p
}

func (s *BpmContext) GetParser() antlr.Parser { return s.parser }

func (s *BpmContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TabParserINTEGER, 0)
}

func (s *BpmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BpmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BpmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterBpm(s)
	}
}

func (s *BpmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitBpm(s)
	}
}

func (p *TabParser) Bpm() (localctx IBpmContext) {
	localctx = NewBpmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TabParserRULE_bpm)
	p.EnterOuterAlt(localctx, 1)
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
		p.Match(TabParserINTEGER)
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

// ITuningContext is an interface to support dynamic dispatch.
type ITuningContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTuningImpl() []ITuningImplContext
	TuningImpl(i int) ITuningImplContext

	// IsTuningContext differentiates from other interfaces.
	IsTuningContext()
}

type TuningContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTuningContext() *TuningContext {
	var p = new(TuningContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_tuning
	return p
}

func InitEmptyTuningContext(p *TuningContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_tuning
}

func (*TuningContext) IsTuningContext() {}

func NewTuningContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TuningContext {
	var p = new(TuningContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_tuning

	return p
}

func (s *TuningContext) GetParser() antlr.Parser { return s.parser }

func (s *TuningContext) AllTuningImpl() []ITuningImplContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITuningImplContext); ok {
			len++
		}
	}

	tst := make([]ITuningImplContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITuningImplContext); ok {
			tst[i] = t.(ITuningImplContext)
			i++
		}
	}

	return tst
}

func (s *TuningContext) TuningImpl(i int) ITuningImplContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITuningImplContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITuningImplContext)
}

func (s *TuningContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TuningContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TuningContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterTuning(s)
	}
}

func (s *TuningContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitTuning(s)
	}
}

func (p *TabParser) Tuning() (localctx ITuningContext) {
	localctx = NewTuningContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TabParserRULE_tuning)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.TuningImpl()
	}
	{
		p.SetState(28)
		p.TuningImpl()
	}
	{
		p.SetState(29)
		p.TuningImpl()
	}
	{
		p.SetState(30)
		p.TuningImpl()
	}
	{
		p.SetState(31)
		p.TuningImpl()
	}
	{
		p.SetState(32)
		p.TuningImpl()
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

// ITuningImplContext is an interface to support dynamic dispatch.
type ITuningImplContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOTELETTER() antlr.TerminalNode
	INTEGER() antlr.TerminalNode

	// IsTuningImplContext differentiates from other interfaces.
	IsTuningImplContext()
}

type TuningImplContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTuningImplContext() *TuningImplContext {
	var p = new(TuningImplContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_tuningImpl
	return p
}

func InitEmptyTuningImplContext(p *TuningImplContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_tuningImpl
}

func (*TuningImplContext) IsTuningImplContext() {}

func NewTuningImplContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TuningImplContext {
	var p = new(TuningImplContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_tuningImpl

	return p
}

func (s *TuningImplContext) GetParser() antlr.Parser { return s.parser }

func (s *TuningImplContext) NOTELETTER() antlr.TerminalNode {
	return s.GetToken(TabParserNOTELETTER, 0)
}

func (s *TuningImplContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TabParserINTEGER, 0)
}

func (s *TuningImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TuningImplContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TuningImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.EnterTuningImpl(s)
	}
}

func (s *TuningImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TabListener); ok {
		listenerT.ExitTuningImpl(s)
	}
}

func (p *TabParser) TuningImpl() (localctx ITuningImplContext) {
	localctx = NewTuningImplContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TabParserRULE_tuningImpl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(TabParserNOTELETTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Match(TabParserINTEGER)
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

// IDurationGroupContext is an interface to support dynamic dispatch.
type IDurationGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Duration() IDurationContext
	AllEntry() []IEntryContext
	Entry(i int) IEntryContext

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

func (s *DurationGroupContext) AllEntry() []IEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEntryContext); ok {
			len++
		}
	}

	tst := make([]IEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEntryContext); ok {
			tst[i] = t.(IEntryContext)
			i++
		}
	}

	return tst
}

func (s *DurationGroupContext) Entry(i int) IEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntryContext)
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
	p.EnterRule(localctx, 8, TabParserRULE_durationGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Duration()
	}
	{
		p.SetState(38)
		p.Match(TabParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&720) != 0 {
		{
			p.SetState(39)
			p.Entry()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(45)
		p.Match(TabParserT__2)
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
	p.EnterRule(localctx, 10, TabParserRULE_duration)
	var _la int

	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDurDescribedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(TabParserDURMOD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TabParserINTEGER {
			{
				p.SetState(49)
				p.Match(TabParserINTEGER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewDurDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
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

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_entry
	return p
}

func InitEmptyEntryContext(p *EntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TabParserRULE_entry
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TabParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) CopyAll(ctx *EntryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PlayFretContext struct {
	EntryContext
}

func NewPlayFretContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlayFretContext {
	var p = new(PlayFretContext)

	InitEmptyEntryContext(&p.EntryContext)
	p.parser = parser
	p.CopyAll(ctx.(*EntryContext))

	return p
}

func (s *PlayFretContext) GetRuleContext() antlr.RuleContext {
	return s
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
	EntryContext
}

func NewPauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PauseContext {
	var p = new(PauseContext)

	InitEmptyEntryContext(&p.EntryContext)
	p.parser = parser
	p.CopyAll(ctx.(*EntryContext))

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
	EntryContext
}

func NewBendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BendContext {
	var p = new(BendContext)

	InitEmptyEntryContext(&p.EntryContext)
	p.parser = parser
	p.CopyAll(ctx.(*EntryContext))

	return p
}

func (s *BendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BendContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(TabParserINTEGER)
}

func (s *BendContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(TabParserINTEGER, i)
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

type SimpleChordContext struct {
	EntryContext
}

func NewSimpleChordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleChordContext {
	var p = new(SimpleChordContext)

	InitEmptyEntryContext(&p.EntryContext)
	p.parser = parser
	p.CopyAll(ctx.(*EntryContext))

	return p
}

func (s *SimpleChordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleChordContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(TabParserINTEGER)
}

func (s *SimpleChordContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(TabParserINTEGER, i)
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

func (p *TabParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TabParserRULE_entry)
	var _la int

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TabParserINTEGER:
		localctx = NewPlayFretContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TabParserT__3:
		localctx = NewSimpleChordContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(TabParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TabParserINTEGER {
			{
				p.SetState(60)
				p.Match(TabParserINTEGER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(66)
			p.Match(TabParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TabParserT__5:
		localctx = NewPauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(TabParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TabParserT__6:
		localctx = NewBendContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Match(TabParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(TabParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(TabParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(TabParserT__2)
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

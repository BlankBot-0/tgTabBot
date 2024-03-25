// Code generated from Tab.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tab

import "github.com/antlr4-go/antlr/v4"

// BaseTabListener is a complete listener for a parse tree produced by TabParser.
type BaseTabListener struct{}

var _ TabListener = &BaseTabListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTabListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTabListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTabListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTabListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTabListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTabListener) ExitStart(ctx *StartContext) {}

// EnterBpm is called when production bpm is entered.
func (s *BaseTabListener) EnterBpm(ctx *BpmContext) {}

// ExitBpm is called when production bpm is exited.
func (s *BaseTabListener) ExitBpm(ctx *BpmContext) {}

// EnterTuning is called when production tuning is entered.
func (s *BaseTabListener) EnterTuning(ctx *TuningContext) {}

// ExitTuning is called when production tuning is exited.
func (s *BaseTabListener) ExitTuning(ctx *TuningContext) {}

// EnterTuningImpl is called when production tuningImpl is entered.
func (s *BaseTabListener) EnterTuningImpl(ctx *TuningImplContext) {}

// ExitTuningImpl is called when production tuningImpl is exited.
func (s *BaseTabListener) ExitTuningImpl(ctx *TuningImplContext) {}

// EnterDurationGroup is called when production durationGroup is entered.
func (s *BaseTabListener) EnterDurationGroup(ctx *DurationGroupContext) {}

// ExitDurationGroup is called when production durationGroup is exited.
func (s *BaseTabListener) ExitDurationGroup(ctx *DurationGroupContext) {}

// EnterDurDescribed is called when production DurDescribed is entered.
func (s *BaseTabListener) EnterDurDescribed(ctx *DurDescribedContext) {}

// ExitDurDescribed is called when production DurDescribed is exited.
func (s *BaseTabListener) ExitDurDescribed(ctx *DurDescribedContext) {}

// EnterDurDefault is called when production DurDefault is entered.
func (s *BaseTabListener) EnterDurDefault(ctx *DurDefaultContext) {}

// ExitDurDefault is called when production DurDefault is exited.
func (s *BaseTabListener) ExitDurDefault(ctx *DurDefaultContext) {}

// EnterPlayFret is called when production PlayFret is entered.
func (s *BaseTabListener) EnterPlayFret(ctx *PlayFretContext) {}

// ExitPlayFret is called when production PlayFret is exited.
func (s *BaseTabListener) ExitPlayFret(ctx *PlayFretContext) {}

// EnterSimpleChord is called when production SimpleChord is entered.
func (s *BaseTabListener) EnterSimpleChord(ctx *SimpleChordContext) {}

// ExitSimpleChord is called when production SimpleChord is exited.
func (s *BaseTabListener) ExitSimpleChord(ctx *SimpleChordContext) {}

// EnterPause is called when production Pause is entered.
func (s *BaseTabListener) EnterPause(ctx *PauseContext) {}

// ExitPause is called when production Pause is exited.
func (s *BaseTabListener) ExitPause(ctx *PauseContext) {}

// EnterBend is called when production Bend is entered.
func (s *BaseTabListener) EnterBend(ctx *BendContext) {}

// ExitBend is called when production Bend is exited.
func (s *BaseTabListener) ExitBend(ctx *BendContext) {}

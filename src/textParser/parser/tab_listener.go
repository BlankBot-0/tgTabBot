// Code generated from Tab.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tab

import "github.com/antlr4-go/antlr/v4"

// TabListener is a complete listener for a parse tree produced by TabParser.
type TabListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterBpm is called when entering the bpm production.
	EnterBpm(c *BpmContext)

	// EnterTuning is called when entering the tuning production.
	EnterTuning(c *TuningContext)

	// EnterTuningImpl is called when entering the tuningImpl production.
	EnterTuningImpl(c *TuningImplContext)

	// EnterDurationGroup is called when entering the durationGroup production.
	EnterDurationGroup(c *DurationGroupContext)

	// EnterDurDescribed is called when entering the DurDescribed production.
	EnterDurDescribed(c *DurDescribedContext)

	// EnterDurDefault is called when entering the DurDefault production.
	EnterDurDefault(c *DurDefaultContext)

	// EnterPlayFret is called when entering the PlayFret production.
	EnterPlayFret(c *PlayFretContext)

	// EnterSimpleChord is called when entering the SimpleChord production.
	EnterSimpleChord(c *SimpleChordContext)

	// EnterPause is called when entering the Pause production.
	EnterPause(c *PauseContext)

	// EnterBend is called when entering the Bend production.
	EnterBend(c *BendContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitBpm is called when exiting the bpm production.
	ExitBpm(c *BpmContext)

	// ExitTuning is called when exiting the tuning production.
	ExitTuning(c *TuningContext)

	// ExitTuningImpl is called when exiting the tuningImpl production.
	ExitTuningImpl(c *TuningImplContext)

	// ExitDurationGroup is called when exiting the durationGroup production.
	ExitDurationGroup(c *DurationGroupContext)

	// ExitDurDescribed is called when exiting the DurDescribed production.
	ExitDurDescribed(c *DurDescribedContext)

	// ExitDurDefault is called when exiting the DurDefault production.
	ExitDurDefault(c *DurDefaultContext)

	// ExitPlayFret is called when exiting the PlayFret production.
	ExitPlayFret(c *PlayFretContext)

	// ExitSimpleChord is called when exiting the SimpleChord production.
	ExitSimpleChord(c *SimpleChordContext)

	// ExitPause is called when exiting the Pause production.
	ExitPause(c *PauseContext)

	// ExitBend is called when exiting the Bend production.
	ExitBend(c *BendContext)
}

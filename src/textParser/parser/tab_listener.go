// Code generated from Tab.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tab

import "github.com/antlr4-go/antlr/v4"

// TabListener is a complete listener for a parse tree produced by TabParser.
type TabListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterDurationGroups is called when entering the durationGroups production.
	EnterDurationGroups(c *DurationGroupsContext)

	// EnterDurationGroup is called when entering the durationGroup production.
	EnterDurationGroup(c *DurationGroupContext)

	// EnterDurDescribed is called when entering the DurDescribed production.
	EnterDurDescribed(c *DurDescribedContext)

	// EnterDurDefault is called when entering the DurDefault production.
	EnterDurDefault(c *DurDefaultContext)

	// EnterNoteGroup is called when entering the noteGroup production.
	EnterNoteGroup(c *NoteGroupContext)

	// EnterSimpleChord is called when entering the SimpleChord production.
	EnterSimpleChord(c *SimpleChordContext)

	// EnterPlayFret is called when entering the PlayFret production.
	EnterPlayFret(c *PlayFretContext)

	// EnterPause is called when entering the Pause production.
	EnterPause(c *PauseContext)

	// EnterBend is called when entering the Bend production.
	EnterBend(c *BendContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitDurationGroups is called when exiting the durationGroups production.
	ExitDurationGroups(c *DurationGroupsContext)

	// ExitDurationGroup is called when exiting the durationGroup production.
	ExitDurationGroup(c *DurationGroupContext)

	// ExitDurDescribed is called when exiting the DurDescribed production.
	ExitDurDescribed(c *DurDescribedContext)

	// ExitDurDefault is called when exiting the DurDefault production.
	ExitDurDefault(c *DurDefaultContext)

	// ExitNoteGroup is called when exiting the noteGroup production.
	ExitNoteGroup(c *NoteGroupContext)

	// ExitSimpleChord is called when exiting the SimpleChord production.
	ExitSimpleChord(c *SimpleChordContext)

	// ExitPlayFret is called when exiting the PlayFret production.
	ExitPlayFret(c *PlayFretContext)

	// ExitPause is called when exiting the Pause production.
	ExitPause(c *PauseContext)

	// ExitBend is called when exiting the Bend production.
	ExitBend(c *BendContext)
}

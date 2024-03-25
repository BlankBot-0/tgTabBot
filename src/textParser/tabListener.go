package textParser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strings"
	"tgScoreBot/src/textParser/parser"
)

var _ parser.TabListener = (*tabListener)(nil)

type TabProcessor interface {
	// Bpm sets bpm
	SetBpm(bpm string)

	// Tuning sets guitar tuning
	SetTuning(tuning []string)

	// Duration sets current note duration
	SetDuration(duration []string)

	// PlayFret processes tab entry
	PlayFret(tab string)

	// PlayChord processes a chord
	PlayChord(chord []string)

	// Pause processes pause
	Pause()

	// ModelBend models string bending
	ModelBend(tab string, bendValue string)

	// Errors returns accumulated errors
	Errors() []error

	// Finish marks the end of parsing
	Finish()
}

type tabListener struct {
	*parser.BaseTabListener
	processors []TabProcessor
	errs       []error
}

// newTabListener returns a new tabListener given an output and midi clock resolution
func newTabListener(processors []TabProcessor) *tabListener {
	return &tabListener{
		processors: processors,
	}
}

// Implement antlr ErrorListener interface for syntax errors.
func (s *tabListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	s.errs = append(s.errs, fmt.Errorf("(%d:%d) %s", line, column, msg))
}

func (s *tabListener) ExitBpm(ctx *parser.BpmContext) {
	bpmStr := ctx.GetChild(1).(antlr.ParseTree).GetText()
	for _, p := range s.processors {
		p.SetBpm(bpmStr)
	}
}

func (s *tabListener) ExitTuning(ctx *parser.TuningContext) {
	tuning := make([]string, ctx.GetChildCount())
	for i, child := range ctx.GetChildren() {
		tuning[i] = strings.ToUpper(child.(antlr.ParseTree).GetText())
	}
	for _, p := range s.processors {
		p.SetTuning(tuning)
	}
}

func (s *tabListener) ExitDurDefault(ctx *parser.DurDefaultContext) {
	for _, p := range s.processors {
		p.SetDuration([]string{ctx.GetText()})
	}
}

func (s *tabListener) ExitDurDescribed(ctx *parser.DurDescribedContext) {
	durAtoms := make([]string, 3)
	for i, ch := range ctx.GetChildren() {
		durAtoms[i] = ch.(antlr.ParseTree).GetText()
	}
	for _, p := range s.processors {
		p.SetDuration(durAtoms)
	}
}

// ExitSimpleChord is called when production SimpleChord is exited.
func (s *tabListener) ExitSimpleChord(ctx *parser.SimpleChordContext) {
	tabs := make([]string, ctx.GetChildCount()-2)
	for i := 1; i != ctx.GetChildCount()-1; i++ {
		ch := ctx.GetChild(i)
		tabs[i-1] = ch.(antlr.ParseTree).GetText()
	}
	for _, p := range s.processors {
		p.PlayChord(tabs)
	}
}

func (s *tabListener) ExitPlayFret(ctx *parser.PlayFretContext) {
	tab := ctx.GetText()
	for _, p := range s.processors {
		// PlayChord is used intentionally as PlayFret might be not needed
		p.PlayChord([]string{tab})
	}
}

func (s *tabListener) ExitBend(ctx *parser.BendContext) {
	tab := ctx.GetChild(1).(antlr.ParseTree).GetText()
	bendVal := ctx.GetChild(3).(antlr.ParseTree).GetText()
	for _, p := range s.processors {
		p.ModelBend(tab, bendVal)
	}
}

func (s *tabListener) ExitPause(ctx *parser.PauseContext) {
	for _, p := range s.processors {
		p.Pause()
	}
}

func (s *tabListener) ExitStart(ctx *parser.StartContext) {
	for _, p := range s.processors {
		p.Finish()
	}
}

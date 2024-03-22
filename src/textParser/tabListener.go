package textParser

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/go-audio/midi"
	"io"
	"strconv"
	"strings"
	"tgScoreBot/src/textParser/parser"
)

var _ parser.TabListener = (*tabListener)(nil)

const (
	minKey         = 0
	maxKey         = 127
	octaveDistance = 12
	stringsNumber  = 6
)

var noteMap = map[string]uint8{
	"C":  12,
	"C#": 13,
	"D":  14,
	"D#": 15,
	"E":  16,
	"F":  17,
	"F#": 18,
	"G":  19,
	"G#": 20,
	"A":  21,
	"A#": 22,
	"B":  23,
}

var velocity = 60

type tabListener struct {
	*parser.BaseTabListener
	enc         *midi.Encoder
	track       *midi.Track
	tuning      [6]uint8
	curDuration uint32
	errs        []error
}

// newTabListener returns a new tabListener given an output and midi clock resolution
func newTabListener(outputMidi io.Writer, opts TabOpts) *tabListener {
	enc := midi.NewEncoder(outputMidi, midi.SingleTrack, opts.ticksPerQuarterNote)
	return &tabListener{enc: enc,
		track:       enc.NewTrack(),
		opts:        opts,
		curDuration: uint32(opts.ticksPerQuarterNote)}
}

// Implement antlr ErrorListener interface for syntax errors.
func (s *tabListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	s.errs = append(s.errs, fmt.Errorf("(%d:%d) %s", line, column, msg))
}

func (s *tabListener) ExitBpm(ctx *parser.BpmContext) {
	bpmStr := ctx.GetChild(1).(antlr.ParseTree).GetText()
	bpm, err := strconv.ParseFloat(bpmStr, 64)
	if err != nil {
		s.errs = append(s.errs, fmt.Errorf("ExitBpm error: %w", err))
	}
	s.track.Add(0, midi.TempoEvent(bpm))
}

func (s *tabListener) ExitTuning(ctx *parser.TuningContext) {
	if n := ctx.GetChildCount(); n != 6 {
		s.errs = append(s.errs, fmt.Errorf("only 6-string tuning is currently supported"))
	}
	for i, child := range ctx.GetChildren() {
		child := child.(antlr.ParseTree)
		fret0 := strings.ToUpper(child.GetText())
		note := noteMap[fret0[:len(fret0)-1]]
		octave, err := strconv.ParseUint(fret0[len(fret0)-1:], 10, 8)
		if err != nil {
			s.errs = append(s.errs, fmt.Errorf("ExitTuning error: %w", err))
		}

		key := note + octaveDistance*uint8(octave)
		if minKey > key || key > maxKey {
			s.errs = append(s.errs, fmt.Errorf("tuning out of range on string %v, C(-1) set instead", i+1))
			key = minKey
		}
		s.opts.tuning[len(s.opts.tuning)-1-i] = key
	}
}

func (s *tabListener) ExitDurDefault(ctx *parser.DurDefaultContext) {
	newDur, err := strconv.ParseUint(ctx.GetText(), 10, 16)
	if err != nil {
		s.errs = append(s.errs, fmt.Errorf("ExitDurDefault error: %w", err))
	}
	s.curDuration = uint32(s.enc.TicksPerQuarterNote) * 4 / uint32(newDur)
}

func (s *tabListener) ExitDurDescribed(ctx *parser.DurDescribedContext) {
	if n := ctx.GetChildCount(); n != 3 {
		panic(errors.New("incorrect duration description"))
	}
	durationStr := ctx.GetChild(0).(antlr.ParseTree).GetText()
	newDur, err := strconv.ParseUint(durationStr, 10, 16)
	if err != nil {
		s.errs = append(s.errs, fmt.Errorf("ExitDurDescribed error: %w", err))
	}

	tupleTypeStr := ctx.GetChild(2).(antlr.ParseTree).GetText()
	tupleType, err := strconv.ParseUint(tupleTypeStr, 10, 16)
	if err != nil {
		s.errs = append(s.errs, fmt.Errorf("ExitDurDescribed error: %w", err))
	}

	newDur = uint64(s.enc.TicksPerQuarterNote) * 8 / tupleType / newDur
	if ctx.GetChild(1).(antlr.ParseTree).GetText() == "._" {
		newDur = newDur + newDur/2
	}

	s.curDuration = uint32(newDur)
}

// ExitSimpleChord is called when production SimpleChord is exited.
func (s *tabListener) ExitSimpleChord(ctx *parser.SimpleChordContext) {
	// The first NoteOn event must appear after correct time
	// and the following must appear instantaneously after it

	for i := 1; i != ctx.GetChildCount()-1; i++ {
		tab := ctx.GetChild(i).(antlr.ParseTree).GetText()
		note := s.tabToNote(tab)
		s.track.AddAfterDelta(0, midi.NoteOn(0, note, velocity))
	}

	// first NoteOff event must appear after correct time
	// and following must appear instantaneously
	tab := ctx.GetChild(1).(antlr.ParseTree).GetText()
	note := s.tabToNote(tab)
	dur := s.curDuration
	s.track.AddAfterDelta(dur, midi.NoteOff(0, note))
	for i := 2; i != ctx.GetChildCount()-1; i++ {
		tab = ctx.GetChild(i).(antlr.ParseTree).GetText()
		note = s.tabToNote(tab)
		s.track.AddAfterDelta(0, midi.NoteOff(0, note))
	}
}

func (s *tabListener) ExitPlayFret(ctx *parser.PlayFretContext) {
	tab := ctx.GetText()
	note := s.tabToNote(tab)
	s.track.AddAfterDelta(0, midi.NoteOn(0, note, velocity))
	s.track.AddAfterDelta(s.curDuration, midi.NoteOff(0, note))
}

func (s *tabListener) ExitBend(ctx *parser.BendContext) {
	note := s.tabToNote(ctx.GetChild(1).(antlr.ParseTree).GetText())
	bend := ctx.GetChild(3).(antlr.ParseTree).GetText()
	s.track.AddAfterDelta(0, midi.NoteOn(0, note, velocity))

	pitchBendEvents := int(s.curDuration / 2)
	var maxPitch int
	switch bend {
	case "1":
		maxPitch = 12288
	default:
		maxPitch = 16383
	}
	step := (maxPitch - 8192) / int(s.curDuration) * 2
	for i := range pitchBendEvents {
		s.track.AddAfterDelta(1, midi.PitchWheelChange(0, 0, 8192+(i+1)*step))
	}
	s.track.AddAfterDelta(s.curDuration-s.curDuration/2, midi.NoteOff(0, note))
	s.track.AddAfterDelta(0, midi.PitchWheelChange(0, 8192, 8192))
}

func (s *tabListener) ExitPause(ctx *parser.PauseContext) {
	s.track.AddAfterDelta(0, midi.NoteOn(0, 0, 0))
	s.track.AddAfterDelta(s.curDuration, midi.NoteOff(0, 0))
}

func (s *tabListener) ExitStart(ctx *parser.StartContext) {
	s.track.AddAfterDelta(s.curDuration, midi.EndOfTrack())
}

// tabToNote converts tab entry to a note number
func (s *tabListener) tabToNote(tab string) int {
	str, err := strconv.ParseInt(tab[:1], 10, 8)
	if err != nil {
		s.errs = append(s.errs, fmt.Errorf("couldn't convert tab to note: %w", err))
		str = 1
	}
	fret, err := strconv.ParseInt(tab[1:], 10, 8)
	if err != nil {
		s.errs = append(s.errs, fmt.Errorf("couldn't convert tab to note: %w", err))
		fret = 0
	}

	note := s.opts.tuning[str-1] + uint8(fret)
	if minKey > note || note > maxKey {
		s.errs = append(s.errs, fmt.Errorf("note out of range: %v", tab))
		note = minKey
	}
	return int(note)
}

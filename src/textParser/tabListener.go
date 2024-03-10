package textParser

import (
	"bytes"
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"github.com/go-audio/midi"
	"strconv"
	"strings"
	"tgScoreBot/src/textParser/parser"
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
	enc                    *midi.Encoder
	tuning                 [6]uint8
	prevDuration           uint32
	curDuration            uint32
	durationGroupBeginning bool
}

func newTabListener(outputMidi *bytes.Buffer, ticksPerQuarterNote uint16) *tabListener {
	return &tabListener{enc: midi.NewEncoder(outputMidi, midi.SingleTrack, ticksPerQuarterNote),
		tuning:       [6]uint8{64, 59, 55, 50, 45, 40},
		prevDuration: uint32(ticksPerQuarterNote),
		curDuration:  uint32(ticksPerQuarterNote)}
}

func (s *tabListener) ExitBpm(ctx *parser.BpmContext) {
	bpmStr := ctx.GetChild(1).(antlr.ParseTree).GetText()
	bpm, err := strconv.ParseFloat(bpmStr, 64)
	if err != nil {
		panic(err)
	}
	s.enc.Tracks[0].Add(0, midi.TempoEvent(bpm))
}

func (s *tabListener) ExitTuning(ctx *parser.TuningContext) {
	for i, child := range ctx.GetChildren() {
		child := child.(antlr.ParseTree)
		fret0 := strings.ToUpper(child.GetText())
		note := noteMap[fret0[:len(fret0)-1]]
		octave, err := strconv.ParseUint(fret0[len(fret0)-1:], 10, 8)
		if err != nil {
			panic(err)
		}
		if key := note + 12*uint8(octave); 0 > key || key > 127 {
			panic(errors.New("note out of range"))
		} else {
			s.tuning[5-i] = key
		}
	}
}

func (s *tabListener) ExitDurDefault(ctx *parser.DurDefaultContext) {
	newDur, err := strconv.ParseUint(ctx.GetText(), 10, 16)
	if err != nil {
		panic(err)
	}
	s.prevDuration = s.curDuration
	s.curDuration = uint32(s.enc.TicksPerQuarterNote) * 4 / uint32(newDur)
	s.durationGroupBeginning = true
}

func (s *tabListener) tabToNote(tab string) int {
	str, err := strconv.ParseInt(tab[:1], 10, 8)
	if err != nil {
		panic(err)
	}
	fret, err := strconv.ParseInt(tab[1:], 10, 8)
	if err != nil {
		panic(err)
	}
	if note := s.tuning[str-1] * uint8(fret); 0 > note || note > 127 {
		panic(errors.New("note out of range"))
	} else {
		return int(note)
	}
}

func (s *tabListener) duration() uint32 {
	if s.durationGroupBeginning {
		return s.prevDuration
	}
	return s.curDuration
}

// ExitSimpleChord is called when production SimpleChord is exited.
func (s *tabListener) ExitSimpleChord(ctx *parser.SimpleChordContext) {
	// The first NoteOn event must appear after correct time
	// and the following must appear instantaneously after it
	tab := ctx.GetChild(1).(antlr.ParseTree).GetText()
	note := s.tabToNote(tab)
	dur := s.duration()
	s.enc.Tracks[0].AddAfterDelta(dur, midi.NoteOn(0, note, velocity))

	for i := 2; i != ctx.GetChildCount()-1; i++ {
		tab = ctx.GetChild(i).(antlr.ParseTree).GetText()
		note = s.tabToNote(tab)
		s.enc.Tracks[0].AddAfterDelta(0, midi.NoteOn(0, note, velocity))
	}

	s.durationGroupBeginning = false

	// first NoteOff event must appear after correct time
	// and following must appear instantaneously
	tab = ctx.GetChild(1).(antlr.ParseTree).GetText()
	note = s.tabToNote(tab)
	dur = s.duration()
	s.enc.Tracks[0].AddAfterDelta(dur, midi.NoteOff(0, note))
	for i := 2; i != ctx.GetChildCount()-1; i++ {
		tab = ctx.GetChild(i).(antlr.ParseTree).GetText()
		note = s.tabToNote(tab)
		s.enc.Tracks[0].AddAfterDelta(0, midi.NoteOff(0, note))
	}
}

func (s *tabListener) ExitPlayFret(ctx parser.PlayFretContext) {
	tab := ctx.GetText()
	note := s.tabToNote(tab)
	s.enc.Tracks[0].AddAfterDelta(s.duration(), midi.NoteOn(0, note, velocity))

	s.durationGroupBeginning = false

	s.enc.Tracks[0].AddAfterDelta(s.duration(), midi.NoteOff(0, note))
}

func (s *tabListener) ExitStart(ctx parser.StartContext) {
	s.enc.Tracks[0].AddAfterDelta(s.curDuration, midi.EndOfTrack())
}

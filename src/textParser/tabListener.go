package textParser

import (
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"github.com/go-audio/midi"
	"io"
	"strconv"
	"strings"
	"tgScoreBot/src/textParser/parser"
)

var _ parser.TabListener = (*tabListener)(nil)

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
}

func newTabListener(outputMidi io.Writer, ticksPerQuarterNote uint16) *tabListener {
	enc := midi.NewEncoder(outputMidi, midi.SingleTrack, ticksPerQuarterNote)

	return &tabListener{enc: enc,
		track:       enc.NewTrack(),
		tuning:      [6]uint8{64, 59, 55, 50, 45, 40},
		curDuration: uint32(ticksPerQuarterNote)}
}

func (s *tabListener) ExitBpm(ctx *parser.BpmContext) {
	bpmStr := ctx.GetChild(1).(antlr.ParseTree).GetText()
	bpm, err := strconv.ParseFloat(bpmStr, 64)
	if err != nil {
		panic(err)
	}
	s.track.Add(0, midi.TempoEvent(bpm))
}

func (s *tabListener) ExitTuning(ctx *parser.TuningContext) {
	if n := ctx.GetChildCount(); n != 6 {
		panic(errors.New("number of strings must be 6"))
	}
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
	s.curDuration = uint32(s.enc.TicksPerQuarterNote) * 4 / uint32(newDur)
}

func (s *tabListener) ExitDurDescribed(ctx *parser.DurDescribedContext) {
	if n := ctx.GetChildCount(); n != 3 {
		panic(errors.New("incorrect duration description"))
	}
	durationStr := ctx.GetChild(0).(antlr.ParseTree).GetText()
	newDur, err := strconv.ParseUint(durationStr, 10, 16)
	if err != nil {
		panic(err)
	}

	tupleTypeStr := ctx.GetChild(2).(antlr.ParseTree).GetText()
	tupleType, err := strconv.ParseUint(tupleTypeStr, 10, 16)
	if err != nil {
		panic(err)
	}

	newDur = uint64(s.enc.TicksPerQuarterNote) * 8 / tupleType / newDur
	if ctx.GetChild(1).(antlr.ParseTree).GetText() == "._" {
		newDur = newDur + newDur/2
	}

	s.curDuration = uint32(newDur)
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
	if note := s.tuning[str-1] + uint8(fret); 0 > note || note > 127 {
		panic(errors.New("note out of range"))
	} else {
		return int(note)
	}
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

func (s *tabListener) ExitPause(ctx *parser.PauseContext) {
	s.track.AddAfterDelta(0, midi.NoteOn(0, 0, 0))
	s.track.AddAfterDelta(s.curDuration, midi.NoteOff(0, 0))
}

func (s *tabListener) ExitStart(ctx *parser.StartContext) {
	s.track.AddAfterDelta(s.curDuration, midi.EndOfTrack())
}

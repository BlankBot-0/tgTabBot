package tabs

import (
	"fmt"
	"github.com/go-audio/midi"
	"io"
	"slices"
	"strconv"
)

var _ TabProcessor = (*TabProcessorMidi)(nil)

type TabProcessorMidi struct {
	enc                 *midi.Encoder
	track               *midi.Track
	curDuration         uint32
	ticksPerQuarterNote uint16
	tuning              []uint8
	errs                []error
}

func NewTabProcessorMidi(outputMidi io.Writer, tpqn uint16) *TabProcessorMidi {
	enc := midi.NewEncoder(outputMidi, midi.SingleTrack, tpqn)
	return &TabProcessorMidi{
		enc:                 enc,
		track:               enc.NewTrack(),
		ticksPerQuarterNote: tpqn,
	}
}

func (p *TabProcessorMidi) SetBpm(bpm string) {
	bpmNumeric, err := strconv.ParseFloat(bpm, 64)
	if err != nil {
		p.errs = append(p.errs, fmt.Errorf("midi bpm error: %w", err))
	}
	p.track.Add(0, midi.TempoEvent(bpmNumeric))
}

func (p *TabProcessorMidi) SetTuning(tuning []string) {
	for i, tune := range tuning {
		note := noteMap[tune[:len(tune)-1]]
		octave, err := strconv.ParseInt(tune[len(tune)-1:], 10, 8)
		if err != nil {
			p.errs = append(p.errs, fmt.Errorf("midi tuning error: %w", err))
			octave = minOctave
		}
		key := uint8(int64(note) + octaveDistance*octave)

		if minKey > key || key > maxKey {
			p.errs = append(p.errs, fmt.Errorf("tuning out of range on string %v, C(-1) set instead", i+1))
			key = minKey
		}
		p.tuning = append(p.tuning, key)
	}
	slices.Reverse(p.tuning)
}

func (p *TabProcessorMidi) SetDuration(duration []string) {
	newDur, err := strconv.ParseUint(duration[0], 10, 16)
	if err != nil {
		p.errs = append(p.errs, fmt.Errorf("midi duration error: %w", err))
		newDur = uint64(p.ticksPerQuarterNote)
	}
	tupleType := uint64(2)
	if len(duration) == 3 {
		tupleType, err = strconv.ParseUint(duration[2], 10, 16)
		if err != nil {
			p.errs = append(p.errs, fmt.Errorf("midi duration tupleType error: %w", err))
			tupleType = TupleDefault
		}
		// TODO: test
		if duration[1] == "._" {
			newDur = newDur - newDur/2
		}
	}
	if p.ticksPerQuarterNote%uint16(tupleType) != 0 || p.ticksPerQuarterNote%uint16(newDur) != 0 {
		p.errs = append(p.errs, fmt.Errorf("midi duration error: unsuitable ticksPerQuarterNote"))
	}
	p.curDuration = uint32(uint64(p.ticksPerQuarterNote) * 8 / tupleType / newDur)
}

func (p *TabProcessorMidi) PlayFret(tab string) {
	// Might be not needed
	return
}

func (p *TabProcessorMidi) PlayChord(chord []string) {
	if len(chord) > len(p.tuning) {
		p.errs = append(p.errs, fmt.Errorf("midi error: too many notes listed"))
	}
	notesToPlay := min(len(p.tuning), len(chord))
	notes := make([]int, notesToPlay)
	for i := range notesToPlay {
		note := p.tabToNote(chord[i])
		notes[i] = note
		p.track.AddAfterDelta(0, midi.NoteOn(MidiChannelDefault, note, VelocityDefault))
	}

	// first NoteOff event must appear after correct time
	// and following must appear instantaneously
	p.track.AddAfterDelta(p.curDuration, midi.NoteOff(MidiChannelDefault, notes[0]))
	for i := 1; i < notesToPlay; i++ {
		p.track.AddAfterDelta(0, midi.NoteOff(MidiChannelDefault, notes[i]))
	}
}

func (p *TabProcessorMidi) Pause() {
	p.track.AddAfterDelta(0, midi.NoteOn(MidiChannelDefault, 0, 0))
	p.track.AddAfterDelta(p.curDuration, midi.NoteOff(MidiChannelDefault, 0))
}

func (p *TabProcessorMidi) ModelBend(tab, bendValue string) {
	note := p.tabToNote(tab)
	p.track.AddAfterDelta(0, midi.NoteOn(MidiChannelDefault, note, VelocityDefault))

	bendVal, err := strconv.ParseFloat(bendValue, 64)
	if err != nil {
		p.errs = append(p.errs, fmt.Errorf("midi bend error: %w", err))
		bendVal = BendValueDefault
	}
	bendTo, err := BendToPitchDefault(bendVal)
	if err != nil {
		p.errs = append(p.errs, err)
	}
	bendDuration := int(p.curDuration / 2)
	for _, val := range BendModelLinear(bendDuration, bendTo) {
		p.track.AddAfterDelta(1, midi.PitchWheelChange(MidiChannelDefault, 0, val))
	}
	p.track.AddAfterDelta(p.curDuration-uint32(bendDuration), midi.NoteOff(MidiChannelDefault, note))
	p.track.AddAfterDelta(0, midi.PitchWheelChange(MidiChannelDefault, 0, PitchBendZeroValue))
}

func (p *TabProcessorMidi) tabToNote(tab string) int {
	str, err := strconv.ParseInt(tab[:1], 10, 8)
	if err != nil {
		p.errs = append(p.errs, fmt.Errorf("couldn't convert tab to note: %w", err))
		str = StringDefault
	}
	fret, err := strconv.ParseInt(tab[1:], 10, 8)
	if err != nil {
		p.errs = append(p.errs, fmt.Errorf("couldn't convert tab to note: %w", err))
		fret = FretDefault
	}

	note := p.tuning[str-1] + uint8(fret)
	if minKey > note || note > maxKey {
		p.errs = append(p.errs, fmt.Errorf("note out of range: %v", tab))
		note = minKey
	}
	return int(note)
}

func (p *TabProcessorMidi) Errors() []error {
	return p.errs
}

func (p *TabProcessorMidi) Finish() {
	p.track.AddAfterDelta(0, midi.EndOfTrack())
	if err := p.enc.Write(); err != nil {
		p.errs = append(p.errs, fmt.Errorf("midi error: %w", err))
	}
}

const (
	minKey             = 0
	maxKey             = 127
	minOctave          = -1
	octaveDistance     = 12
	stringsNumber      = 6
	StringDefault      = 1
	FretDefault        = 0
	TupleDefault       = 4
	MidiChannelDefault = 0
	VelocityDefault    = 60
	BendValueDefault   = 0
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

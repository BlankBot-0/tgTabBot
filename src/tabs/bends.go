package tabs

import (
	"fmt"
	"math"
)

// BendModelLinear makes a slice of pitches in a linear fashion
func BendModelLinear(outputSize, bendTo int) []int {
	res := make([]int, outputSize)
	step := (bendTo - PitchBendZeroValue) / outputSize
	for i := range res {
		res[i] = PitchBendZeroValue + i*step
	}
	return res
}

// BendToPitchDefault is a default map which assumes that
// stringBend is measured in semitones,
// +2 semitones maps to PitchBendMaxValue
// -2 semitones maps to PitchBendMinValue
func BendToPitchDefault(stringBend float64) (int, error) {
	if math.Abs(stringBend) > 2 {
		return PitchBendZeroValue, fmt.Errorf("BendToPitchDefault: incorrect bend value")
	}
	return int(float64(PitchBendMaxValue) / 4 * (stringBend + 2)), nil
}

const (
	PitchBendMaxValue  = 16383
	PitchBendZeroValue = 8192
	PitchBendMinValue  = 0
)

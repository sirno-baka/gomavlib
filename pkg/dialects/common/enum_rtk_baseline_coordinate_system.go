//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// RTK GPS baseline coordinate system, used for RTK corrections
type RTK_BASELINE_COORDINATE_SYSTEM uint32

const (
	// Earth-centered, Earth-fixed
	RTK_BASELINE_COORDINATE_SYSTEM_ECEF RTK_BASELINE_COORDINATE_SYSTEM = 0
	// RTK basestation centered, north, east, down
	RTK_BASELINE_COORDINATE_SYSTEM_NED RTK_BASELINE_COORDINATE_SYSTEM = 1
)

var labels_RTK_BASELINE_COORDINATE_SYSTEM = map[RTK_BASELINE_COORDINATE_SYSTEM]string{
	RTK_BASELINE_COORDINATE_SYSTEM_ECEF: "RTK_BASELINE_COORDINATE_SYSTEM_ECEF",
	RTK_BASELINE_COORDINATE_SYSTEM_NED:  "RTK_BASELINE_COORDINATE_SYSTEM_NED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e RTK_BASELINE_COORDINATE_SYSTEM) MarshalText() ([]byte, error) {
	if l, ok := labels_RTK_BASELINE_COORDINATE_SYSTEM[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_RTK_BASELINE_COORDINATE_SYSTEM = map[string]RTK_BASELINE_COORDINATE_SYSTEM{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *RTK_BASELINE_COORDINATE_SYSTEM) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_RTK_BASELINE_COORDINATE_SYSTEM[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e RTK_BASELINE_COORDINATE_SYSTEM) String() string {
	if l, ok := labels_RTK_BASELINE_COORDINATE_SYSTEM[e]; ok {
		return l
	}
	return "invalid value"
}

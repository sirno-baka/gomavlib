//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Enable axes that will be tuned via autotuning. Used in MAV_CMD_DO_AUTOTUNE_ENABLE.
type AUTOTUNE_AXIS uint32

const (
	// Flight stack tunes axis according to its default settings.
	AUTOTUNE_AXIS_DEFAULT AUTOTUNE_AXIS = 0
	// Autotune roll axis.
	AUTOTUNE_AXIS_ROLL AUTOTUNE_AXIS = 1
	// Autotune pitch axis.
	AUTOTUNE_AXIS_PITCH AUTOTUNE_AXIS = 2
	// Autotune yaw axis.
	AUTOTUNE_AXIS_YAW AUTOTUNE_AXIS = 4
)

var labels_AUTOTUNE_AXIS = map[AUTOTUNE_AXIS]string{
	AUTOTUNE_AXIS_DEFAULT: "AUTOTUNE_AXIS_DEFAULT",
	AUTOTUNE_AXIS_ROLL:    "AUTOTUNE_AXIS_ROLL",
	AUTOTUNE_AXIS_PITCH:   "AUTOTUNE_AXIS_PITCH",
	AUTOTUNE_AXIS_YAW:     "AUTOTUNE_AXIS_YAW",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e AUTOTUNE_AXIS) MarshalText() ([]byte, error) {
	if l, ok := labels_AUTOTUNE_AXIS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_AUTOTUNE_AXIS = map[string]AUTOTUNE_AXIS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *AUTOTUNE_AXIS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_AUTOTUNE_AXIS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e AUTOTUNE_AXIS) String() string {
	if l, ok := labels_AUTOTUNE_AXIS[e]; ok {
		return l
	}
	return "invalid value"
}

//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Sequence that motors are tested when using MAV_CMD_DO_MOTOR_TEST.
type MOTOR_TEST_ORDER uint32

const (
	// Default autopilot motor test method.
	MOTOR_TEST_ORDER_DEFAULT MOTOR_TEST_ORDER = 0
	// Motor numbers are specified as their index in a predefined vehicle-specific sequence.
	MOTOR_TEST_ORDER_SEQUENCE MOTOR_TEST_ORDER = 1
	// Motor numbers are specified as the output as labeled on the board.
	MOTOR_TEST_ORDER_BOARD MOTOR_TEST_ORDER = 2
)

var labels_MOTOR_TEST_ORDER = map[MOTOR_TEST_ORDER]string{
	MOTOR_TEST_ORDER_DEFAULT:  "MOTOR_TEST_ORDER_DEFAULT",
	MOTOR_TEST_ORDER_SEQUENCE: "MOTOR_TEST_ORDER_SEQUENCE",
	MOTOR_TEST_ORDER_BOARD:    "MOTOR_TEST_ORDER_BOARD",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MOTOR_TEST_ORDER) MarshalText() ([]byte, error) {
	if l, ok := labels_MOTOR_TEST_ORDER[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MOTOR_TEST_ORDER = map[string]MOTOR_TEST_ORDER{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MOTOR_TEST_ORDER) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MOTOR_TEST_ORDER[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MOTOR_TEST_ORDER) String() string {
	if l, ok := labels_MOTOR_TEST_ORDER[e]; ok {
		return l
	}
	return "invalid value"
}

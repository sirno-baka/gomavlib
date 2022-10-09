//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"errors"
)

type MAV_ODID_ARM_STATUS uint32

const (
	// Passing arming checks.
	MAV_ODID_ARM_STATUS_GOOD_TO_ARM MAV_ODID_ARM_STATUS = 0
	// Generic arming failure, see error string for details.
	MAV_ODID_ARM_STATUS_PRE_ARM_FAIL_GENERIC MAV_ODID_ARM_STATUS = 1
)

var labels_MAV_ODID_ARM_STATUS = map[MAV_ODID_ARM_STATUS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_ARM_STATUS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_ARM_STATUS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_ARM_STATUS = map[string]MAV_ODID_ARM_STATUS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_ARM_STATUS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_ARM_STATUS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_ARM_STATUS) String() string {
	if l, ok := labels_MAV_ODID_ARM_STATUS[e]; ok {
		return l
	}
	return "invalid value"
}

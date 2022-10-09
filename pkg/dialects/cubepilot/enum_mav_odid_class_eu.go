//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"errors"
)

type MAV_ODID_CLASS_EU uint32

const (
	// The class for the UA, according to the EU specification, is undeclared.
	MAV_ODID_CLASS_EU_UNDECLARED MAV_ODID_CLASS_EU = 0
	// The class for the UA, according to the EU specification, is Class 0.
	MAV_ODID_CLASS_EU_CLASS_0 MAV_ODID_CLASS_EU = 1
	// The class for the UA, according to the EU specification, is Class 1.
	MAV_ODID_CLASS_EU_CLASS_1 MAV_ODID_CLASS_EU = 2
	// The class for the UA, according to the EU specification, is Class 2.
	MAV_ODID_CLASS_EU_CLASS_2 MAV_ODID_CLASS_EU = 3
	// The class for the UA, according to the EU specification, is Class 3.
	MAV_ODID_CLASS_EU_CLASS_3 MAV_ODID_CLASS_EU = 4
	// The class for the UA, according to the EU specification, is Class 4.
	MAV_ODID_CLASS_EU_CLASS_4 MAV_ODID_CLASS_EU = 5
	// The class for the UA, according to the EU specification, is Class 5.
	MAV_ODID_CLASS_EU_CLASS_5 MAV_ODID_CLASS_EU = 6
	// The class for the UA, according to the EU specification, is Class 6.
	MAV_ODID_CLASS_EU_CLASS_6 MAV_ODID_CLASS_EU = 7
)

var labels_MAV_ODID_CLASS_EU = map[MAV_ODID_CLASS_EU]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_CLASS_EU) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_CLASS_EU[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_CLASS_EU = map[string]MAV_ODID_CLASS_EU{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_CLASS_EU) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_CLASS_EU[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_CLASS_EU) String() string {
	if l, ok := labels_MAV_ODID_CLASS_EU[e]; ok {
		return l
	}
	return "invalid value"
}

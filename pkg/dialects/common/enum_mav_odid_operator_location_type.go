//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

type MAV_ODID_OPERATOR_LOCATION_TYPE uint32

const (
	// The location/altitude of the operator is the same as the take-off location.
	MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF MAV_ODID_OPERATOR_LOCATION_TYPE = 0
	// The location/altitude of the operator is dynamic. E.g. based on live GNSS data.
	MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS MAV_ODID_OPERATOR_LOCATION_TYPE = 1
	// The location/altitude of the operator are fixed values.
	MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED MAV_ODID_OPERATOR_LOCATION_TYPE = 2
)

var labels_MAV_ODID_OPERATOR_LOCATION_TYPE = map[MAV_ODID_OPERATOR_LOCATION_TYPE]string{
	MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF:   "MAV_ODID_OPERATOR_LOCATION_TYPE_TAKEOFF",
	MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS: "MAV_ODID_OPERATOR_LOCATION_TYPE_LIVE_GNSS",
	MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED:     "MAV_ODID_OPERATOR_LOCATION_TYPE_FIXED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_OPERATOR_LOCATION_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_OPERATOR_LOCATION_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_OPERATOR_LOCATION_TYPE = map[string]MAV_ODID_OPERATOR_LOCATION_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_OPERATOR_LOCATION_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_OPERATOR_LOCATION_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_OPERATOR_LOCATION_TYPE) String() string {
	if l, ok := labels_MAV_ODID_OPERATOR_LOCATION_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}

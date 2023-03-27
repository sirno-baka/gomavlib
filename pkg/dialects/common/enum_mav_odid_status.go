//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

type MAV_ODID_STATUS uint32

const (
	// The status of the (UA) Unmanned Aircraft is undefined.
	MAV_ODID_STATUS_UNDECLARED MAV_ODID_STATUS = 0
	// The UA is on the ground.
	MAV_ODID_STATUS_GROUND MAV_ODID_STATUS = 1
	// The UA is in the air.
	MAV_ODID_STATUS_AIRBORNE MAV_ODID_STATUS = 2
	// The UA is having an emergency.
	MAV_ODID_STATUS_EMERGENCY MAV_ODID_STATUS = 3
	// The remote ID system is failing or unreliable in some way.
	MAV_ODID_STATUS_REMOTE_ID_SYSTEM_FAILURE MAV_ODID_STATUS = 4
)

var labels_MAV_ODID_STATUS = map[MAV_ODID_STATUS]string{
	MAV_ODID_STATUS_UNDECLARED:               "MAV_ODID_STATUS_UNDECLARED",
	MAV_ODID_STATUS_GROUND:                   "MAV_ODID_STATUS_GROUND",
	MAV_ODID_STATUS_AIRBORNE:                 "MAV_ODID_STATUS_AIRBORNE",
	MAV_ODID_STATUS_EMERGENCY:                "MAV_ODID_STATUS_EMERGENCY",
	MAV_ODID_STATUS_REMOTE_ID_SYSTEM_FAILURE: "MAV_ODID_STATUS_REMOTE_ID_SYSTEM_FAILURE",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_STATUS) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_ODID_STATUS {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_STATUS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_ODID_STATUS
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_ODID_STATUS {
			if l == label {
				mask |= value
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_STATUS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

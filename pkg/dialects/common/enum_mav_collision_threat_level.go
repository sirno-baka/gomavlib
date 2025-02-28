//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Aircraft-rated danger from this threat.
type MAV_COLLISION_THREAT_LEVEL uint32

const (
	// Not a threat
	MAV_COLLISION_THREAT_LEVEL_NONE MAV_COLLISION_THREAT_LEVEL = 0
	// Craft is mildly concerned about this threat
	MAV_COLLISION_THREAT_LEVEL_LOW MAV_COLLISION_THREAT_LEVEL = 1
	// Craft is panicking, and may take actions to avoid threat
	MAV_COLLISION_THREAT_LEVEL_HIGH MAV_COLLISION_THREAT_LEVEL = 2
)

var labels_MAV_COLLISION_THREAT_LEVEL = map[MAV_COLLISION_THREAT_LEVEL]string{
	MAV_COLLISION_THREAT_LEVEL_NONE: "MAV_COLLISION_THREAT_LEVEL_NONE",
	MAV_COLLISION_THREAT_LEVEL_LOW:  "MAV_COLLISION_THREAT_LEVEL_LOW",
	MAV_COLLISION_THREAT_LEVEL_HIGH: "MAV_COLLISION_THREAT_LEVEL_HIGH",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_COLLISION_THREAT_LEVEL) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_COLLISION_THREAT_LEVEL[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_COLLISION_THREAT_LEVEL = map[string]MAV_COLLISION_THREAT_LEVEL{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_COLLISION_THREAT_LEVEL) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_COLLISION_THREAT_LEVEL[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_COLLISION_THREAT_LEVEL) String() string {
	if l, ok := labels_MAV_COLLISION_THREAT_LEVEL[e]; ok {
		return l
	}
	return "invalid value"
}

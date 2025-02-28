//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type GOPRO_CHARGING uint32

const (
	// Charging disabled.
	GOPRO_CHARGING_DISABLED GOPRO_CHARGING = 0
	// Charging enabled.
	GOPRO_CHARGING_ENABLED GOPRO_CHARGING = 1
)

var labels_GOPRO_CHARGING = map[GOPRO_CHARGING]string{
	GOPRO_CHARGING_DISABLED: "GOPRO_CHARGING_DISABLED",
	GOPRO_CHARGING_ENABLED:  "GOPRO_CHARGING_ENABLED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_CHARGING) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_CHARGING[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_CHARGING = map[string]GOPRO_CHARGING{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_CHARGING) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_CHARGING[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_CHARGING) String() string {
	if l, ok := labels_GOPRO_CHARGING[e]; ok {
		return l
	}
	return "invalid value"
}

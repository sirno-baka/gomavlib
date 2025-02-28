//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Enumeration of the ADSB altimeter types
type ADSB_ALTITUDE_TYPE uint32

const (
	// Altitude reported from a Baro source using QNH reference
	ADSB_ALTITUDE_TYPE_PRESSURE_QNH ADSB_ALTITUDE_TYPE = 0
	// Altitude reported from a GNSS source
	ADSB_ALTITUDE_TYPE_GEOMETRIC ADSB_ALTITUDE_TYPE = 1
)

var labels_ADSB_ALTITUDE_TYPE = map[ADSB_ALTITUDE_TYPE]string{
	ADSB_ALTITUDE_TYPE_PRESSURE_QNH: "ADSB_ALTITUDE_TYPE_PRESSURE_QNH",
	ADSB_ALTITUDE_TYPE_GEOMETRIC:    "ADSB_ALTITUDE_TYPE_GEOMETRIC",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ADSB_ALTITUDE_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_ADSB_ALTITUDE_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ADSB_ALTITUDE_TYPE = map[string]ADSB_ALTITUDE_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ADSB_ALTITUDE_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ADSB_ALTITUDE_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ADSB_ALTITUDE_TYPE) String() string {
	if l, ok := labels_ADSB_ALTITUDE_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}

//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// Enumeration of battery functions
type MAV_BATTERY_FUNCTION uint32

const (
	// Battery function is unknown
	MAV_BATTERY_FUNCTION_UNKNOWN MAV_BATTERY_FUNCTION = 0
	// Battery supports all flight systems
	MAV_BATTERY_FUNCTION_ALL MAV_BATTERY_FUNCTION = 1
	// Battery for the propulsion system
	MAV_BATTERY_FUNCTION_PROPULSION MAV_BATTERY_FUNCTION = 2
	// Avionics battery
	MAV_BATTERY_FUNCTION_AVIONICS MAV_BATTERY_FUNCTION = 3
	// Payload battery
	MAV_BATTERY_FUNCTION_PAYLOAD MAV_BATTERY_FUNCTION = 4
)

var labels_MAV_BATTERY_FUNCTION = map[MAV_BATTERY_FUNCTION]string{
	MAV_BATTERY_FUNCTION_UNKNOWN:    "MAV_BATTERY_FUNCTION_UNKNOWN",
	MAV_BATTERY_FUNCTION_ALL:        "MAV_BATTERY_FUNCTION_ALL",
	MAV_BATTERY_FUNCTION_PROPULSION: "MAV_BATTERY_FUNCTION_PROPULSION",
	MAV_BATTERY_FUNCTION_AVIONICS:   "MAV_BATTERY_FUNCTION_AVIONICS",
	MAV_BATTERY_FUNCTION_PAYLOAD:    "MAV_BATTERY_FUNCTION_PAYLOAD",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_FUNCTION) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_BATTERY_FUNCTION {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_FUNCTION) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_BATTERY_FUNCTION
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_BATTERY_FUNCTION {
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
func (e MAV_BATTERY_FUNCTION) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// Actions following geofence breach.
type FENCE_ACTION uint32

const (
	// Disable fenced mode. If used in a plan this would mean the next fence is disabled.
	FENCE_ACTION_NONE FENCE_ACTION = 0
	// Fly to geofence MAV_CMD_NAV_FENCE_RETURN_POINT in GUIDED mode. Note: This action is only supported by ArduPlane, and may not be supported in all versions.
	FENCE_ACTION_GUIDED FENCE_ACTION = 1
	// Report fence breach, but don't take action
	FENCE_ACTION_REPORT FENCE_ACTION = 2
	// Fly to geofence MAV_CMD_NAV_FENCE_RETURN_POINT with manual throttle control in GUIDED mode. Note: This action is only supported by ArduPlane, and may not be supported in all versions.
	FENCE_ACTION_GUIDED_THR_PASS FENCE_ACTION = 3
	// Return/RTL mode.
	FENCE_ACTION_RTL FENCE_ACTION = 4
	// Hold at current location.
	FENCE_ACTION_HOLD FENCE_ACTION = 5
	// Termination failsafe. Motors are shut down (some flight stacks may trigger other failsafe actions).
	FENCE_ACTION_TERMINATE FENCE_ACTION = 6
	// Land at current location.
	FENCE_ACTION_LAND FENCE_ACTION = 7
)

var labels_FENCE_ACTION = map[FENCE_ACTION]string{
	FENCE_ACTION_NONE:            "FENCE_ACTION_NONE",
	FENCE_ACTION_GUIDED:          "FENCE_ACTION_GUIDED",
	FENCE_ACTION_REPORT:          "FENCE_ACTION_REPORT",
	FENCE_ACTION_GUIDED_THR_PASS: "FENCE_ACTION_GUIDED_THR_PASS",
	FENCE_ACTION_RTL:             "FENCE_ACTION_RTL",
	FENCE_ACTION_HOLD:            "FENCE_ACTION_HOLD",
	FENCE_ACTION_TERMINATE:       "FENCE_ACTION_TERMINATE",
	FENCE_ACTION_LAND:            "FENCE_ACTION_LAND",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e FENCE_ACTION) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_FENCE_ACTION {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *FENCE_ACTION) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask FENCE_ACTION
	for _, label := range labels {
		found := false
		for value, l := range labels_FENCE_ACTION {
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
func (e FENCE_ACTION) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

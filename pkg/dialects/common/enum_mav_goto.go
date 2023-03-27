//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// Actions that may be specified in MAV_CMD_OVERRIDE_GOTO to override mission execution.
type MAV_GOTO uint32

const (
	// Hold at the current position.
	MAV_GOTO_DO_HOLD MAV_GOTO = 0
	// Continue with the next item in mission execution.
	MAV_GOTO_DO_CONTINUE MAV_GOTO = 1
	// Hold at the current position of the system
	MAV_GOTO_HOLD_AT_CURRENT_POSITION MAV_GOTO = 2
	// Hold at the position specified in the parameters of the DO_HOLD action
	MAV_GOTO_HOLD_AT_SPECIFIED_POSITION MAV_GOTO = 3
)

var labels_MAV_GOTO = map[MAV_GOTO]string{
	MAV_GOTO_DO_HOLD:                    "MAV_GOTO_DO_HOLD",
	MAV_GOTO_DO_CONTINUE:                "MAV_GOTO_DO_CONTINUE",
	MAV_GOTO_HOLD_AT_CURRENT_POSITION:   "MAV_GOTO_HOLD_AT_CURRENT_POSITION",
	MAV_GOTO_HOLD_AT_SPECIFIED_POSITION: "MAV_GOTO_HOLD_AT_SPECIFIED_POSITION",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_GOTO) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_GOTO {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_GOTO) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_GOTO
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_GOTO {
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
func (e MAV_GOTO) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

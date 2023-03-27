//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package matrixpilot

import (
	"fmt"
	"strings"
)

// Action required when performing CMD_PREFLIGHT_STORAGE
type MAV_PREFLIGHT_STORAGE_ACTION uint32

const (
	// Read all parameters from storage
	MAV_PFS_CMD_READ_ALL MAV_PREFLIGHT_STORAGE_ACTION = 0
	// Write all parameters to storage
	MAV_PFS_CMD_WRITE_ALL MAV_PREFLIGHT_STORAGE_ACTION = 1
	// Clear all  parameters in storage
	MAV_PFS_CMD_CLEAR_ALL MAV_PREFLIGHT_STORAGE_ACTION = 2
	// Read specific parameters from storage
	MAV_PFS_CMD_READ_SPECIFIC MAV_PREFLIGHT_STORAGE_ACTION = 3
	// Write specific parameters to storage
	MAV_PFS_CMD_WRITE_SPECIFIC MAV_PREFLIGHT_STORAGE_ACTION = 4
	// Clear specific parameters in storage
	MAV_PFS_CMD_CLEAR_SPECIFIC MAV_PREFLIGHT_STORAGE_ACTION = 5
	// do nothing
	MAV_PFS_CMD_DO_NOTHING MAV_PREFLIGHT_STORAGE_ACTION = 6
)

var labels_MAV_PREFLIGHT_STORAGE_ACTION = map[MAV_PREFLIGHT_STORAGE_ACTION]string{
	MAV_PFS_CMD_READ_ALL:       "MAV_PFS_CMD_READ_ALL",
	MAV_PFS_CMD_WRITE_ALL:      "MAV_PFS_CMD_WRITE_ALL",
	MAV_PFS_CMD_CLEAR_ALL:      "MAV_PFS_CMD_CLEAR_ALL",
	MAV_PFS_CMD_READ_SPECIFIC:  "MAV_PFS_CMD_READ_SPECIFIC",
	MAV_PFS_CMD_WRITE_SPECIFIC: "MAV_PFS_CMD_WRITE_SPECIFIC",
	MAV_PFS_CMD_CLEAR_SPECIFIC: "MAV_PFS_CMD_CLEAR_SPECIFIC",
	MAV_PFS_CMD_DO_NOTHING:     "MAV_PFS_CMD_DO_NOTHING",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_PREFLIGHT_STORAGE_ACTION) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_PREFLIGHT_STORAGE_ACTION {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_PREFLIGHT_STORAGE_ACTION) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_PREFLIGHT_STORAGE_ACTION
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_PREFLIGHT_STORAGE_ACTION {
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
func (e MAV_PREFLIGHT_STORAGE_ACTION) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

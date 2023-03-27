//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strings"
)

// Possible remote log data block statuses.
type MAV_REMOTE_LOG_DATA_BLOCK_STATUSES uint32

const (
	// This block has NOT been received.
	MAV_REMOTE_LOG_DATA_BLOCK_NACK MAV_REMOTE_LOG_DATA_BLOCK_STATUSES = 0
	// This block has been received.
	MAV_REMOTE_LOG_DATA_BLOCK_ACK MAV_REMOTE_LOG_DATA_BLOCK_STATUSES = 1
)

var labels_MAV_REMOTE_LOG_DATA_BLOCK_STATUSES = map[MAV_REMOTE_LOG_DATA_BLOCK_STATUSES]string{
	MAV_REMOTE_LOG_DATA_BLOCK_NACK: "MAV_REMOTE_LOG_DATA_BLOCK_NACK",
	MAV_REMOTE_LOG_DATA_BLOCK_ACK:  "MAV_REMOTE_LOG_DATA_BLOCK_ACK",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_REMOTE_LOG_DATA_BLOCK_STATUSES) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_REMOTE_LOG_DATA_BLOCK_STATUSES {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_REMOTE_LOG_DATA_BLOCK_STATUSES) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_REMOTE_LOG_DATA_BLOCK_STATUSES
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_REMOTE_LOG_DATA_BLOCK_STATUSES {
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
func (e MAV_REMOTE_LOG_DATA_BLOCK_STATUSES) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

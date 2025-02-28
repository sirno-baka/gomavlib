//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Flags to indicate usage for a particular storage (see STORAGE_INFORMATION.storage_usage and MAV_CMD_SET_STORAGE_USAGE).
type STORAGE_USAGE_FLAG uint32

const (
	// Always set to 1 (indicates STORAGE_INFORMATION.storage_usage is supported).
	STORAGE_USAGE_FLAG_SET STORAGE_USAGE_FLAG = 1
	// Storage for saving photos.
	STORAGE_USAGE_FLAG_PHOTO STORAGE_USAGE_FLAG = 2
	// Storage for saving videos.
	STORAGE_USAGE_FLAG_VIDEO STORAGE_USAGE_FLAG = 4
	// Storage for saving logs.
	STORAGE_USAGE_FLAG_LOGS STORAGE_USAGE_FLAG = 8
)

var labels_STORAGE_USAGE_FLAG = map[STORAGE_USAGE_FLAG]string{
	STORAGE_USAGE_FLAG_SET:   "STORAGE_USAGE_FLAG_SET",
	STORAGE_USAGE_FLAG_PHOTO: "STORAGE_USAGE_FLAG_PHOTO",
	STORAGE_USAGE_FLAG_VIDEO: "STORAGE_USAGE_FLAG_VIDEO",
	STORAGE_USAGE_FLAG_LOGS:  "STORAGE_USAGE_FLAG_LOGS",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e STORAGE_USAGE_FLAG) MarshalText() ([]byte, error) {
	if l, ok := labels_STORAGE_USAGE_FLAG[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_STORAGE_USAGE_FLAG = map[string]STORAGE_USAGE_FLAG{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *STORAGE_USAGE_FLAG) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_STORAGE_USAGE_FLAG[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e STORAGE_USAGE_FLAG) String() string {
	if l, ok := labels_STORAGE_USAGE_FLAG[e]; ok {
		return l
	}
	return "invalid value"
}

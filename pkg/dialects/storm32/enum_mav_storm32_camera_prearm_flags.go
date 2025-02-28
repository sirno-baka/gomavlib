//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"errors"
)

// STorM32 camera prearm check flags.
type MAV_STORM32_CAMERA_PREARM_FLAGS uint32

const (
	// The camera has been found and is connected.
	MAV_STORM32_CAMERA_PREARM_FLAGS_CONNECTED MAV_STORM32_CAMERA_PREARM_FLAGS = 1
)

var labels_MAV_STORM32_CAMERA_PREARM_FLAGS = map[MAV_STORM32_CAMERA_PREARM_FLAGS]string{
	MAV_STORM32_CAMERA_PREARM_FLAGS_CONNECTED: "MAV_STORM32_CAMERA_PREARM_FLAGS_CONNECTED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STORM32_CAMERA_PREARM_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_STORM32_CAMERA_PREARM_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_STORM32_CAMERA_PREARM_FLAGS = map[string]MAV_STORM32_CAMERA_PREARM_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STORM32_CAMERA_PREARM_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_STORM32_CAMERA_PREARM_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_STORM32_CAMERA_PREARM_FLAGS) String() string {
	if l, ok := labels_MAV_STORM32_CAMERA_PREARM_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}

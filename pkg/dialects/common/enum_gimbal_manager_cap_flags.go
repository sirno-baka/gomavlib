//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Gimbal manager high level capability flags (bitmap). The first 16 bits are identical to the GIMBAL_DEVICE_CAP_FLAGS. However, the gimbal manager does not need to copy the flags from the gimbal but can also enhance the capabilities and thus add flags.
type GIMBAL_MANAGER_CAP_FLAGS uint32

const (
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT GIMBAL_MANAGER_CAP_FLAGS = 1
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL GIMBAL_MANAGER_CAP_FLAGS = 2
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS GIMBAL_MANAGER_CAP_FLAGS = 4
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 8
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK GIMBAL_MANAGER_CAP_FLAGS = 16
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS GIMBAL_MANAGER_CAP_FLAGS = 32
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 64
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK GIMBAL_MANAGER_CAP_FLAGS = 128
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS GIMBAL_MANAGER_CAP_FLAGS = 256
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = 512
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK GIMBAL_MANAGER_CAP_FLAGS = 1024
	// Based on GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW.
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW GIMBAL_MANAGER_CAP_FLAGS = 2048
	// Based on GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME.
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME GIMBAL_MANAGER_CAP_FLAGS = 4096
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_RC_INPUTS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_RC_INPUTS GIMBAL_MANAGER_CAP_FLAGS = 8192
	// Gimbal manager supports to point to a local position.
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL GIMBAL_MANAGER_CAP_FLAGS = 65536
	// Gimbal manager supports to point to a global latitude, longitude, altitude position.
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL GIMBAL_MANAGER_CAP_FLAGS = 131072
)

var labels_GIMBAL_MANAGER_CAP_FLAGS = map[GIMBAL_MANAGER_CAP_FLAGS]string{
	GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT:                 "GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL:                 "GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS:               "GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW:             "GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK:               "GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS:              "GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW:            "GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK:              "GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS:                "GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW:              "GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK:                "GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK",
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW:       "GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW",
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME: "GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME",
	GIMBAL_MANAGER_CAP_FLAGS_HAS_RC_INPUTS:               "GIMBAL_MANAGER_CAP_FLAGS_HAS_RC_INPUTS",
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL:    "GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL",
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL:   "GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GIMBAL_MANAGER_CAP_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_GIMBAL_MANAGER_CAP_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GIMBAL_MANAGER_CAP_FLAGS = map[string]GIMBAL_MANAGER_CAP_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GIMBAL_MANAGER_CAP_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GIMBAL_MANAGER_CAP_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GIMBAL_MANAGER_CAP_FLAGS) String() string {
	if l, ok := labels_GIMBAL_MANAGER_CAP_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}

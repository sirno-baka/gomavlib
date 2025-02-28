//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

// The type of parameter for the OSD parameter editor.
type OSD_PARAM_CONFIG_TYPE uint32

const (
	OSD_PARAM_NONE              OSD_PARAM_CONFIG_TYPE = 0
	OSD_PARAM_SERIAL_PROTOCOL   OSD_PARAM_CONFIG_TYPE = 1
	OSD_PARAM_SERVO_FUNCTION    OSD_PARAM_CONFIG_TYPE = 2
	OSD_PARAM_AUX_FUNCTION      OSD_PARAM_CONFIG_TYPE = 3
	OSD_PARAM_FLIGHT_MODE       OSD_PARAM_CONFIG_TYPE = 4
	OSD_PARAM_FAILSAFE_ACTION   OSD_PARAM_CONFIG_TYPE = 5
	OSD_PARAM_FAILSAFE_ACTION_1 OSD_PARAM_CONFIG_TYPE = 6
	OSD_PARAM_FAILSAFE_ACTION_2 OSD_PARAM_CONFIG_TYPE = 7
	OSD_PARAM_NUM_TYPES         OSD_PARAM_CONFIG_TYPE = 8
)

var labels_OSD_PARAM_CONFIG_TYPE = map[OSD_PARAM_CONFIG_TYPE]string{
	OSD_PARAM_NONE:              "OSD_PARAM_NONE",
	OSD_PARAM_SERIAL_PROTOCOL:   "OSD_PARAM_SERIAL_PROTOCOL",
	OSD_PARAM_SERVO_FUNCTION:    "OSD_PARAM_SERVO_FUNCTION",
	OSD_PARAM_AUX_FUNCTION:      "OSD_PARAM_AUX_FUNCTION",
	OSD_PARAM_FLIGHT_MODE:       "OSD_PARAM_FLIGHT_MODE",
	OSD_PARAM_FAILSAFE_ACTION:   "OSD_PARAM_FAILSAFE_ACTION",
	OSD_PARAM_FAILSAFE_ACTION_1: "OSD_PARAM_FAILSAFE_ACTION_1",
	OSD_PARAM_FAILSAFE_ACTION_2: "OSD_PARAM_FAILSAFE_ACTION_2",
	OSD_PARAM_NUM_TYPES:         "OSD_PARAM_NUM_TYPES",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e OSD_PARAM_CONFIG_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_OSD_PARAM_CONFIG_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_OSD_PARAM_CONFIG_TYPE = map[string]OSD_PARAM_CONFIG_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *OSD_PARAM_CONFIG_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_OSD_PARAM_CONFIG_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e OSD_PARAM_CONFIG_TYPE) String() string {
	if l, ok := labels_OSD_PARAM_CONFIG_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}

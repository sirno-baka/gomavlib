//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

// The error type for the OSD parameter editor.
type OSD_PARAM_CONFIG_ERROR uint32

const (
	OSD_PARAM_SUCCESS                 OSD_PARAM_CONFIG_ERROR = 0
	OSD_PARAM_INVALID_SCREEN          OSD_PARAM_CONFIG_ERROR = 1
	OSD_PARAM_INVALID_PARAMETER_INDEX OSD_PARAM_CONFIG_ERROR = 2
	OSD_PARAM_INVALID_PARAMETER       OSD_PARAM_CONFIG_ERROR = 3
)

var labels_OSD_PARAM_CONFIG_ERROR = map[OSD_PARAM_CONFIG_ERROR]string{
	OSD_PARAM_SUCCESS:                 "OSD_PARAM_SUCCESS",
	OSD_PARAM_INVALID_SCREEN:          "OSD_PARAM_INVALID_SCREEN",
	OSD_PARAM_INVALID_PARAMETER_INDEX: "OSD_PARAM_INVALID_PARAMETER_INDEX",
	OSD_PARAM_INVALID_PARAMETER:       "OSD_PARAM_INVALID_PARAMETER",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e OSD_PARAM_CONFIG_ERROR) MarshalText() ([]byte, error) {
	if l, ok := labels_OSD_PARAM_CONFIG_ERROR[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_OSD_PARAM_CONFIG_ERROR = map[string]OSD_PARAM_CONFIG_ERROR{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *OSD_PARAM_CONFIG_ERROR) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_OSD_PARAM_CONFIG_ERROR[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e OSD_PARAM_CONFIG_ERROR) String() string {
	if l, ok := labels_OSD_PARAM_CONFIG_ERROR[e]; ok {
		return l
	}
	return "invalid value"
}

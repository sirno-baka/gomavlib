//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"errors"
)

// Reason for an event error response.
type MAV_EVENT_ERROR_REASON uint32

const (
	// The requested event is not available (anymore).
	MAV_EVENT_ERROR_REASON_UNAVAILABLE MAV_EVENT_ERROR_REASON = 0
)

var labels_MAV_EVENT_ERROR_REASON = map[MAV_EVENT_ERROR_REASON]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_EVENT_ERROR_REASON) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_EVENT_ERROR_REASON[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_EVENT_ERROR_REASON = map[string]MAV_EVENT_ERROR_REASON{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_EVENT_ERROR_REASON) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_EVENT_ERROR_REASON[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_EVENT_ERROR_REASON) String() string {
	if l, ok := labels_MAV_EVENT_ERROR_REASON[e]; ok {
		return l
	}
	return "invalid value"
}

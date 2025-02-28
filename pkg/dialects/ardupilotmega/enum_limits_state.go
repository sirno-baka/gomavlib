//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type LIMITS_STATE uint32

const (
	// Pre-initialization.
	LIMITS_INIT LIMITS_STATE = 0
	// Disabled.
	LIMITS_DISABLED LIMITS_STATE = 1
	// Checking limits.
	LIMITS_ENABLED LIMITS_STATE = 2
	// A limit has been breached.
	LIMITS_TRIGGERED LIMITS_STATE = 3
	// Taking action e.g. Return/RTL.
	LIMITS_RECOVERING LIMITS_STATE = 4
	// We're no longer in breach of a limit.
	LIMITS_RECOVERED LIMITS_STATE = 5
)

var labels_LIMITS_STATE = map[LIMITS_STATE]string{
	LIMITS_INIT:       "LIMITS_INIT",
	LIMITS_DISABLED:   "LIMITS_DISABLED",
	LIMITS_ENABLED:    "LIMITS_ENABLED",
	LIMITS_TRIGGERED:  "LIMITS_TRIGGERED",
	LIMITS_RECOVERING: "LIMITS_RECOVERING",
	LIMITS_RECOVERED:  "LIMITS_RECOVERED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e LIMITS_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_LIMITS_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_LIMITS_STATE = map[string]LIMITS_STATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *LIMITS_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_LIMITS_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e LIMITS_STATE) String() string {
	if l, ok := labels_LIMITS_STATE[e]; ok {
		return l
	}
	return "invalid value"
}

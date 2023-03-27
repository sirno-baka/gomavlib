//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strings"
)

// These flags are used to diagnose the failure state of CELLULAR_STATUS
type CELLULAR_NETWORK_FAILED_REASON uint32

const (
	// No error
	CELLULAR_NETWORK_FAILED_REASON_NONE CELLULAR_NETWORK_FAILED_REASON = 0
	// Error state is unknown
	CELLULAR_NETWORK_FAILED_REASON_UNKNOWN CELLULAR_NETWORK_FAILED_REASON = 1
	// SIM is required for the modem but missing
	CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING CELLULAR_NETWORK_FAILED_REASON = 2
	// SIM is available, but not usable for connection
	CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR CELLULAR_NETWORK_FAILED_REASON = 3
)

var labels_CELLULAR_NETWORK_FAILED_REASON = map[CELLULAR_NETWORK_FAILED_REASON]string{
	CELLULAR_NETWORK_FAILED_REASON_NONE:        "CELLULAR_NETWORK_FAILED_REASON_NONE",
	CELLULAR_NETWORK_FAILED_REASON_UNKNOWN:     "CELLULAR_NETWORK_FAILED_REASON_UNKNOWN",
	CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING: "CELLULAR_NETWORK_FAILED_REASON_SIM_MISSING",
	CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR:   "CELLULAR_NETWORK_FAILED_REASON_SIM_ERROR",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CELLULAR_NETWORK_FAILED_REASON) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_CELLULAR_NETWORK_FAILED_REASON {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CELLULAR_NETWORK_FAILED_REASON) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask CELLULAR_NETWORK_FAILED_REASON
	for _, label := range labels {
		found := false
		for value, l := range labels_CELLULAR_NETWORK_FAILED_REASON {
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
func (e CELLULAR_NETWORK_FAILED_REASON) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

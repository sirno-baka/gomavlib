//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"errors"
)

// Possible parameter transaction actions.
type PARAM_TRANSACTION_ACTION uint32

const (
	// Commit the current parameter transaction.
	PARAM_TRANSACTION_ACTION_START PARAM_TRANSACTION_ACTION = 0
	// Commit the current parameter transaction.
	PARAM_TRANSACTION_ACTION_COMMIT PARAM_TRANSACTION_ACTION = 1
	// Cancel the current parameter transaction.
	PARAM_TRANSACTION_ACTION_CANCEL PARAM_TRANSACTION_ACTION = 2
)

var labels_PARAM_TRANSACTION_ACTION = map[PARAM_TRANSACTION_ACTION]string{
	PARAM_TRANSACTION_ACTION_START:  "PARAM_TRANSACTION_ACTION_START",
	PARAM_TRANSACTION_ACTION_COMMIT: "PARAM_TRANSACTION_ACTION_COMMIT",
	PARAM_TRANSACTION_ACTION_CANCEL: "PARAM_TRANSACTION_ACTION_CANCEL",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PARAM_TRANSACTION_ACTION) MarshalText() ([]byte, error) {
	if l, ok := labels_PARAM_TRANSACTION_ACTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_PARAM_TRANSACTION_ACTION = map[string]PARAM_TRANSACTION_ACTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PARAM_TRANSACTION_ACTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_PARAM_TRANSACTION_ACTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e PARAM_TRANSACTION_ACTION) String() string {
	if l, ok := labels_PARAM_TRANSACTION_ACTION[e]; ok {
		return l
	}
	return "invalid value"
}

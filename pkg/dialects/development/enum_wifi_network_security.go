//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"errors"
)

// WiFi wireless security protocols.
type WIFI_NETWORK_SECURITY uint32

const (
	// Undefined or unknown security protocol.
	WIFI_NETWORK_SECURITY_UNDEFINED WIFI_NETWORK_SECURITY = 0
	// Open network, no security.
	WIFI_NETWORK_SECURITY_OPEN WIFI_NETWORK_SECURITY = 1
	// WEP.
	WIFI_NETWORK_SECURITY_WEP WIFI_NETWORK_SECURITY = 2
	// WPA1.
	WIFI_NETWORK_SECURITY_WPA1 WIFI_NETWORK_SECURITY = 3
	// WPA2.
	WIFI_NETWORK_SECURITY_WPA2 WIFI_NETWORK_SECURITY = 4
	// WPA3.
	WIFI_NETWORK_SECURITY_WPA3 WIFI_NETWORK_SECURITY = 5
)

var labels_WIFI_NETWORK_SECURITY = map[WIFI_NETWORK_SECURITY]string{
	WIFI_NETWORK_SECURITY_UNDEFINED: "WIFI_NETWORK_SECURITY_UNDEFINED",
	WIFI_NETWORK_SECURITY_OPEN:      "WIFI_NETWORK_SECURITY_OPEN",
	WIFI_NETWORK_SECURITY_WEP:       "WIFI_NETWORK_SECURITY_WEP",
	WIFI_NETWORK_SECURITY_WPA1:      "WIFI_NETWORK_SECURITY_WPA1",
	WIFI_NETWORK_SECURITY_WPA2:      "WIFI_NETWORK_SECURITY_WPA2",
	WIFI_NETWORK_SECURITY_WPA3:      "WIFI_NETWORK_SECURITY_WPA3",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e WIFI_NETWORK_SECURITY) MarshalText() ([]byte, error) {
	if l, ok := labels_WIFI_NETWORK_SECURITY[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_WIFI_NETWORK_SECURITY = map[string]WIFI_NETWORK_SECURITY{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *WIFI_NETWORK_SECURITY) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_WIFI_NETWORK_SECURITY[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e WIFI_NETWORK_SECURITY) String() string {
	if l, ok := labels_WIFI_NETWORK_SECURITY[e]; ok {
		return l
	}
	return "invalid value"
}

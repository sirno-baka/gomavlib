//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"fmt"
	"strings"
)

type AVSS_M300_OPERATION_MODE uint32

const (
	// In manual control mode
	MODE_M300_MANUAL_CTRL AVSS_M300_OPERATION_MODE = 0
	// In attitude mode
	MODE_M300_ATTITUDE AVSS_M300_OPERATION_MODE = 1
	// In GPS mode
	MODE_M300_P_GPS AVSS_M300_OPERATION_MODE = 6
	// In hotpoint mode
	MODE_M300_HOTPOINT_MODE AVSS_M300_OPERATION_MODE = 9
	// In assisted takeoff mode
	MODE_M300_ASSISTED_TAKEOFF AVSS_M300_OPERATION_MODE = 10
	// In auto takeoff mode
	MODE_M300_AUTO_TAKEOFF AVSS_M300_OPERATION_MODE = 11
	// In auto landing mode
	MODE_M300_AUTO_LANDING AVSS_M300_OPERATION_MODE = 12
	// In go home mode
	MODE_M300_NAVI_GO_HOME AVSS_M300_OPERATION_MODE = 15
	// In sdk control mode
	MODE_M300_NAVI_SDK_CTRL AVSS_M300_OPERATION_MODE = 17
	// In sport mode
	MODE_M300_S_SPORT AVSS_M300_OPERATION_MODE = 31
	// In force auto landing mode
	MODE_M300_FORCE_AUTO_LANDING AVSS_M300_OPERATION_MODE = 33
	// In tripod mode
	MODE_M300_T_TRIPOD AVSS_M300_OPERATION_MODE = 38
	// In search mode
	MODE_M300_SEARCH_MODE AVSS_M300_OPERATION_MODE = 40
	// In engine mode
	MODE_M300_ENGINE_START AVSS_M300_OPERATION_MODE = 41
)

var labels_AVSS_M300_OPERATION_MODE = map[AVSS_M300_OPERATION_MODE]string{
	MODE_M300_MANUAL_CTRL:        "MODE_M300_MANUAL_CTRL",
	MODE_M300_ATTITUDE:           "MODE_M300_ATTITUDE",
	MODE_M300_P_GPS:              "MODE_M300_P_GPS",
	MODE_M300_HOTPOINT_MODE:      "MODE_M300_HOTPOINT_MODE",
	MODE_M300_ASSISTED_TAKEOFF:   "MODE_M300_ASSISTED_TAKEOFF",
	MODE_M300_AUTO_TAKEOFF:       "MODE_M300_AUTO_TAKEOFF",
	MODE_M300_AUTO_LANDING:       "MODE_M300_AUTO_LANDING",
	MODE_M300_NAVI_GO_HOME:       "MODE_M300_NAVI_GO_HOME",
	MODE_M300_NAVI_SDK_CTRL:      "MODE_M300_NAVI_SDK_CTRL",
	MODE_M300_S_SPORT:            "MODE_M300_S_SPORT",
	MODE_M300_FORCE_AUTO_LANDING: "MODE_M300_FORCE_AUTO_LANDING",
	MODE_M300_T_TRIPOD:           "MODE_M300_T_TRIPOD",
	MODE_M300_SEARCH_MODE:        "MODE_M300_SEARCH_MODE",
	MODE_M300_ENGINE_START:       "MODE_M300_ENGINE_START",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e AVSS_M300_OPERATION_MODE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_AVSS_M300_OPERATION_MODE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *AVSS_M300_OPERATION_MODE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask AVSS_M300_OPERATION_MODE
	for _, label := range labels {
		found := false
		for value, l := range labels_AVSS_M300_OPERATION_MODE {
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
func (e AVSS_M300_OPERATION_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strings"
)

type GOPRO_CAPTURE_MODE uint32

const (
	// Video mode.
	GOPRO_CAPTURE_MODE_VIDEO GOPRO_CAPTURE_MODE = 0
	// Photo mode.
	GOPRO_CAPTURE_MODE_PHOTO GOPRO_CAPTURE_MODE = 1
	// Burst mode, Hero 3+ only.
	GOPRO_CAPTURE_MODE_BURST GOPRO_CAPTURE_MODE = 2
	// Time lapse mode, Hero 3+ only.
	GOPRO_CAPTURE_MODE_TIME_LAPSE GOPRO_CAPTURE_MODE = 3
	// Multi shot mode, Hero 4 only.
	GOPRO_CAPTURE_MODE_MULTI_SHOT GOPRO_CAPTURE_MODE = 4
	// Playback mode, Hero 4 only, silver only except when LCD or HDMI is connected to black.
	GOPRO_CAPTURE_MODE_PLAYBACK GOPRO_CAPTURE_MODE = 5
	// Playback mode, Hero 4 only.
	GOPRO_CAPTURE_MODE_SETUP GOPRO_CAPTURE_MODE = 6
	// Mode not yet known.
	GOPRO_CAPTURE_MODE_UNKNOWN GOPRO_CAPTURE_MODE = 255
)

var labels_GOPRO_CAPTURE_MODE = map[GOPRO_CAPTURE_MODE]string{
	GOPRO_CAPTURE_MODE_VIDEO:      "GOPRO_CAPTURE_MODE_VIDEO",
	GOPRO_CAPTURE_MODE_PHOTO:      "GOPRO_CAPTURE_MODE_PHOTO",
	GOPRO_CAPTURE_MODE_BURST:      "GOPRO_CAPTURE_MODE_BURST",
	GOPRO_CAPTURE_MODE_TIME_LAPSE: "GOPRO_CAPTURE_MODE_TIME_LAPSE",
	GOPRO_CAPTURE_MODE_MULTI_SHOT: "GOPRO_CAPTURE_MODE_MULTI_SHOT",
	GOPRO_CAPTURE_MODE_PLAYBACK:   "GOPRO_CAPTURE_MODE_PLAYBACK",
	GOPRO_CAPTURE_MODE_SETUP:      "GOPRO_CAPTURE_MODE_SETUP",
	GOPRO_CAPTURE_MODE_UNKNOWN:    "GOPRO_CAPTURE_MODE_UNKNOWN",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_CAPTURE_MODE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_GOPRO_CAPTURE_MODE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_CAPTURE_MODE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask GOPRO_CAPTURE_MODE
	for _, label := range labels {
		found := false
		for value, l := range labels_GOPRO_CAPTURE_MODE {
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
func (e GOPRO_CAPTURE_MODE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

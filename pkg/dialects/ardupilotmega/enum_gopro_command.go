//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strings"
)

type GOPRO_COMMAND uint32

const (
	// (Get/Set).
	GOPRO_COMMAND_POWER GOPRO_COMMAND = 0
	// (Get/Set).
	GOPRO_COMMAND_CAPTURE_MODE GOPRO_COMMAND = 1
	// (___/Set).
	GOPRO_COMMAND_SHUTTER GOPRO_COMMAND = 2
	// (Get/___).
	GOPRO_COMMAND_BATTERY GOPRO_COMMAND = 3
	// (Get/___).
	GOPRO_COMMAND_MODEL GOPRO_COMMAND = 4
	// (Get/Set).
	GOPRO_COMMAND_VIDEO_SETTINGS GOPRO_COMMAND = 5
	// (Get/Set).
	GOPRO_COMMAND_LOW_LIGHT GOPRO_COMMAND = 6
	// (Get/Set).
	GOPRO_COMMAND_PHOTO_RESOLUTION GOPRO_COMMAND = 7
	// (Get/Set).
	GOPRO_COMMAND_PHOTO_BURST_RATE GOPRO_COMMAND = 8
	// (Get/Set).
	GOPRO_COMMAND_PROTUNE GOPRO_COMMAND = 9
	// (Get/Set) Hero 3+ Only.
	GOPRO_COMMAND_PROTUNE_WHITE_BALANCE GOPRO_COMMAND = 10
	// (Get/Set) Hero 3+ Only.
	GOPRO_COMMAND_PROTUNE_COLOUR GOPRO_COMMAND = 11
	// (Get/Set) Hero 3+ Only.
	GOPRO_COMMAND_PROTUNE_GAIN GOPRO_COMMAND = 12
	// (Get/Set) Hero 3+ Only.
	GOPRO_COMMAND_PROTUNE_SHARPNESS GOPRO_COMMAND = 13
	// (Get/Set) Hero 3+ Only.
	GOPRO_COMMAND_PROTUNE_EXPOSURE GOPRO_COMMAND = 14
	// (Get/Set).
	GOPRO_COMMAND_TIME GOPRO_COMMAND = 15
	// (Get/Set).
	GOPRO_COMMAND_CHARGING GOPRO_COMMAND = 16
)

var labels_GOPRO_COMMAND = map[GOPRO_COMMAND]string{
	GOPRO_COMMAND_POWER:                 "GOPRO_COMMAND_POWER",
	GOPRO_COMMAND_CAPTURE_MODE:          "GOPRO_COMMAND_CAPTURE_MODE",
	GOPRO_COMMAND_SHUTTER:               "GOPRO_COMMAND_SHUTTER",
	GOPRO_COMMAND_BATTERY:               "GOPRO_COMMAND_BATTERY",
	GOPRO_COMMAND_MODEL:                 "GOPRO_COMMAND_MODEL",
	GOPRO_COMMAND_VIDEO_SETTINGS:        "GOPRO_COMMAND_VIDEO_SETTINGS",
	GOPRO_COMMAND_LOW_LIGHT:             "GOPRO_COMMAND_LOW_LIGHT",
	GOPRO_COMMAND_PHOTO_RESOLUTION:      "GOPRO_COMMAND_PHOTO_RESOLUTION",
	GOPRO_COMMAND_PHOTO_BURST_RATE:      "GOPRO_COMMAND_PHOTO_BURST_RATE",
	GOPRO_COMMAND_PROTUNE:               "GOPRO_COMMAND_PROTUNE",
	GOPRO_COMMAND_PROTUNE_WHITE_BALANCE: "GOPRO_COMMAND_PROTUNE_WHITE_BALANCE",
	GOPRO_COMMAND_PROTUNE_COLOUR:        "GOPRO_COMMAND_PROTUNE_COLOUR",
	GOPRO_COMMAND_PROTUNE_GAIN:          "GOPRO_COMMAND_PROTUNE_GAIN",
	GOPRO_COMMAND_PROTUNE_SHARPNESS:     "GOPRO_COMMAND_PROTUNE_SHARPNESS",
	GOPRO_COMMAND_PROTUNE_EXPOSURE:      "GOPRO_COMMAND_PROTUNE_EXPOSURE",
	GOPRO_COMMAND_TIME:                  "GOPRO_COMMAND_TIME",
	GOPRO_COMMAND_CHARGING:              "GOPRO_COMMAND_CHARGING",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_COMMAND) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_GOPRO_COMMAND {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_COMMAND) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask GOPRO_COMMAND
	for _, label := range labels {
		found := false
		for value, l := range labels_GOPRO_COMMAND {
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
func (e GOPRO_COMMAND) String() string {
	val, _ := e.MarshalText()
	return string(val)
}

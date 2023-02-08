//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Radio channels. Supports up to 24 channels. Channel values are in centerd 13 bit format. Range is [-4096,4096], center is 0. Conversion to PWM is x * 5/32 + 1500. Should be emitted only by components with component id MAV_COMP_ID_TELEMETRY_RADIO.
type MessageRadioRcChannels struct {
	// Total number of RC channels being received. This can be larger than 24, indicating that more channels are available but not given in this message.
	Count uint8
	// Radio channels status flags.
	Flags RADIO_RC_CHANNELS_FLAGS `mavenum:"uint8"`
	// RC channels. Channels above count should be set to 0, to benefit from MAVLink's zero padding.
	Channels [24]int16 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageRadioRcChannels) GetID() uint32 {
	return 60045
}

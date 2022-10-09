//autogenerated:yes
//nolint:revive,misspell,govet,lll
package cubepilot

// Configure cellular modems.
// This message is re-emitted as an acknowledgement by the modem.
// The message may also be explicitly requested using MAV_CMD_REQUEST_MESSAGE.
type MessageCellularConfig struct {
	// Enable/disable LTE. 0: setting unchanged, 1: disabled, 2: enabled. Current setting when sent back as a response.
	EnableLte uint8
	// Enable/disable PIN on the SIM card. 0: setting unchanged, 1: disabled, 2: enabled. Current setting when sent back as a response.
	EnablePin uint8
	// PIN sent to the SIM card. Blank when PIN is disabled. Empty when message is sent back as a response.
	Pin string `mavlen:"16"`
	// New PIN when changing the PIN. Blank to leave it unchanged. Empty when message is sent back as a response.
	NewPin string `mavlen:"16"`
	// Name of the cellular APN. Blank to leave it unchanged. Current APN when sent back as a response.
	Apn string `mavlen:"32"`
	// Required PUK code in case the user failed to authenticate 3 times with the PIN. Empty when message is sent back as a response.
	Puk string `mavlen:"16"`
	// Enable/disable roaming. 0: setting unchanged, 1: disabled, 2: enabled. Current setting when sent back as a response.
	Roaming uint8
	// Message acceptance response (sent back to GS).
	Response CELLULAR_CONFIG_RESPONSE `mavenum:"uint8"`
}

// GetID implements the message.Message interface.
func (*MessageCellularConfig) GetID() uint32 {
	return 336
}

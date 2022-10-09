//autogenerated:yes
//nolint:revive,misspell,govet,lll
package cubepilot

// Reply to LOG_REQUEST_LIST
type MessageLogEntry struct {
	// Log id
	Id uint16
	// Total number of logs
	NumLogs uint16
	// High log number
	LastLogNum uint16
	// UTC timestamp of log since 1970, or 0 if not available
	TimeUtc uint32
	// Size of the log (may be approximate)
	Size uint32
}

// GetID implements the message.Message interface.
func (*MessageLogEntry) GetID() uint32 {
	return 118
}

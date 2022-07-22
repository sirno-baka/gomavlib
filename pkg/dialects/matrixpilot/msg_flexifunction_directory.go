//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Acknowldge success or failure of a flexifunction command
type MessageFlexifunctionDirectory struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// 0=inputs, 1=outputs
	DirectoryType uint8
	// index of first directory entry to write
	StartIndex uint8
	// count of directory entries to write
	Count uint8
	// Settings data
	DirectoryData [48]int8
}

// GetID implements the message.Message interface.
func (*MessageFlexifunctionDirectory) GetID() uint32 {
	return 155
}

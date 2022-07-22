//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Publishes the GPS coordinates of the vehicle local origin (0,0,0) position. Emitted whenever a new GPS-Local position mapping is requested or set - e.g. following SET_GPS_GLOBAL_ORIGIN message.
type MessageGpsGlobalOrigin struct {
	// Latitude (WGS84)
	Latitude int32
	// Longitude (WGS84)
	Longitude int32
	// Altitude (MSL). Positive for up.
	Altitude int32
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageGpsGlobalOrigin) GetID() uint32 {
	return 49
}

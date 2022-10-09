//autogenerated:yes
//nolint:revive,misspell,govet,lll
package cubepilot

// High level message to control a gimbal's attitude. This message is to be sent to the gimbal manager (e.g. from a ground station). Angles and rates can be set to NaN according to use case.
type MessageGimbalManagerSetAttitude struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// High level gimbal manager flags to use.
	Flags GIMBAL_MANAGER_FLAGS `mavenum:"uint32"`
	// Component ID of gimbal device to address (or 1-6 for non-MAVLink gimbal), 0 for all gimbal device components. Send command multiple times for more than one gimbal (but not all gimbals).
	GimbalDeviceId uint8
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation, the frame is depends on whether the flag GIMBAL_MANAGER_FLAGS_YAW_LOCK is set)
	Q [4]float32
	// X component of angular velocity, positive is rolling to the right, NaN to be ignored.
	AngularVelocityX float32
	// Y component of angular velocity, positive is pitching up, NaN to be ignored.
	AngularVelocityY float32
	// Z component of angular velocity, positive is yawing to the right, NaN to be ignored.
	AngularVelocityZ float32
}

// GetID implements the message.Message interface.
func (*MessageGimbalManagerSetAttitude) GetID() uint32 {
	return 282
}

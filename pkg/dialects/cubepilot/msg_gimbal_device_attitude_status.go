//autogenerated:yes
//nolint:revive,misspell,govet,lll
package cubepilot

// Message reporting the status of a gimbal device. This message should be broadcasted by a gimbal device component. The angles encoded in the quaternion are relative to absolute North if the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set (roll: positive is rolling to the right, pitch: positive is pitching up, yaw is turn to the right) or relative to the vehicle heading if the flag is not set. This message should be broadcast at a low regular rate (e.g. 10Hz).
type MessageGimbalDeviceAttitudeStatus struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Current gimbal flags set.
	Flags GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation, the frame is depends on whether the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set)
	Q [4]float32
	// X component of angular velocity (NaN if unknown)
	AngularVelocityX float32
	// Y component of angular velocity (NaN if unknown)
	AngularVelocityY float32
	// Z component of angular velocity (NaN if unknown)
	AngularVelocityZ float32
	// Failure flags (0 for no failure)
	FailureFlags GIMBAL_DEVICE_ERROR_FLAGS `mavenum:"uint32"`
}

// GetID implements the message.Message interface.
func (*MessageGimbalDeviceAttitudeStatus) GetID() uint32 {
	return 285
}

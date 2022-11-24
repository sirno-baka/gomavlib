//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Low level message to control a gimbal device's attitude.
// This message is to be sent from the gimbal manager to the gimbal device component.
// The quaternion and angular velocities can be set to NaN according to use case.
// For the angles encoded in the quaternion and the angular velocities holds:
// If the flag GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME is set, then they are relative to the vehicle heading (vehicle frame).
// If the flag GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME is set, then they are relative to absolute North (earth frame).
// If neither of these flags are set, then (for backwards compatibility) it holds:
// If the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set, then they are relative to absolute North (earth frame),
// else they are relative to the vehicle heading (vehicle frame).
// Setting both GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME and GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME is not allowed.
// These rules are to ensure backwards compatibility.
// New implementations should always set either GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME or GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME.
type MessageGimbalDeviceSetAttitude struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Low level gimbal flags.
	Flags GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation). The frame is described in the message description. Set fields to NaN to be ignored.
	Q [4]float32
	// X component of angular velocity (positive: rolling to the right). The frame is described in the message description. NaN to be ignored.
	AngularVelocityX float32
	// Y component of angular velocity (positive: pitching up). The frame is described in the message description. NaN to be ignored.
	AngularVelocityY float32
	// Z component of angular velocity (positive: yawing to the right). The frame is described in the message description. NaN to be ignored.
	AngularVelocityZ float32
}

// GetID implements the message.Message interface.
func (*MessageGimbalDeviceSetAttitude) GetID() uint32 {
	return 284
}

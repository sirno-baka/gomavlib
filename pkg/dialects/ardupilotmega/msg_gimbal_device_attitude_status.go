//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// Message reporting the status of a gimbal device.
// This message should be broadcast by a gimbal device component at a low regular rate (e.g. 5 Hz).
// For the angles encoded in the quaternion and the angular velocities holds:
// If the flag GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME is set, then they are relative to the vehicle heading (vehicle frame).
// If the flag GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME is set, then they are relative to absolute North (earth frame).
// If neither of these flags are set, then (for backwards compatibility) it holds:
// If the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set, then they are relative to absolute North (earth frame),
// else they are relative to the vehicle heading (vehicle frame).
// Other conditions of the flags are not allowed.
// The quaternion and angular velocities in the other frame can be calculated from delta_yaw and delta_yaw_velocity as
// q_earth = q_delta_yaw * q_vehicle and w_earth = w_delta_yaw_velocity + w_vehicle (if not NaN).
// If neither the GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME nor the GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME flag is set,
// then (for backwards compatibility) the data in the delta_yaw and delta_yaw_velocity fields are to be ignored.
// New implementations should always set either GIMBAL_DEVICE_FLAGS_YAW_IN_VEHICLE_FRAME or GIMBAL_DEVICE_FLAGS_YAW_IN_EARTH_FRAME,
// and always should set delta_yaw and delta_yaw_velocity either to the proper value or NaN.
type MessageGimbalDeviceAttitudeStatus struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Current gimbal flags set.
	Flags GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation). The frame is described in the message description.
	Q [4]float32
	// X component of angular velocity (positive: rolling to the right). The frame is described in the message description. NaN if unknown.
	AngularVelocityX float32
	// Y component of angular velocity (positive: pitching up). The frame is described in the message description. NaN if unknown.
	AngularVelocityY float32
	// Z component of angular velocity (positive: yawing to the right). The frame is described in the message description. NaN if unknown.
	AngularVelocityZ float32
	// Failure flags (0 for no failure)
	FailureFlags GIMBAL_DEVICE_ERROR_FLAGS `mavenum:"uint32"`
	// Yaw angle relating the quaternions in earth and body frames (see message description). NaN if unknown.
	DeltaYaw float32 `mavext:"true"`
	// Yaw angular velocity relating the angular velocities in earth and body frames (see message description). NaN if unknown.
	DeltaYawVelocity float32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageGimbalDeviceAttitudeStatus) GetID() uint32 {
	return 285
}

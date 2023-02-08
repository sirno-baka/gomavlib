//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Message to a gimbal manager to control the gimbal tilt and pan angles. Angles and rates can be set to NaN according to use case. A gimbal device is never to react to this message.
type MessageStorm32GimbalManagerControlPitchyaw struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Gimbal ID of the gimbal manager to address (component ID or 1-6 for non-MAVLink gimbal, 0 for all gimbals). Send command multiple times for more than one but not all gimbals.
	GimbalId uint8
	// Client which is contacting the gimbal manager (must be set).
	Client MAV_STORM32_GIMBAL_MANAGER_CLIENT `mavenum:"uint8"`
	// Gimbal device flags to be applied (UINT16_MAX to be ignored). Same flags as used in GIMBAL_DEVICE_SET_ATTITUDE.
	DeviceFlags GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Gimbal manager flags to be applied (0 to be ignored).
	ManagerFlags MAV_STORM32_GIMBAL_MANAGER_FLAGS `mavenum:"uint16"`
	// Pitch/tilt angle (positive: tilt up). NaN to be ignored.
	Pitch float32
	// Yaw/pan angle (positive: pan the right). NaN to be ignored. The frame is determined by the GIMBAL_DEVICE_FLAGS_YAW_IN_xxx_FRAME flags.
	Yaw float32
	// Pitch/tilt angular rate (positive: tilt up). NaN to be ignored.
	PitchRate float32
	// Yaw/pan angular rate (positive: pan to the right). NaN to be ignored. The frame is determined by the GIMBAL_DEVICE_FLAGS_YAW_IN_xxx_FRAME flags.
	YawRate float32
}

// GetID implements the message.Message interface.
func (*MessageStorm32GimbalManagerControlPitchyaw) GetID() uint32 {
	return 60013
}

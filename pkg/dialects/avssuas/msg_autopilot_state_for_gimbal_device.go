//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Low level message containing autopilot state relevant for a gimbal device. This message is to be sent from the autopilot to the gimbal device component. The data of this message are for the gimbal device's estimator corrections, in particular horizon compensation, as well as indicates autopilot control intentions, e.g. feed forward angular control in the z-axis.
type MessageAutopilotStateForGimbalDevice struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Timestamp (time since system boot).
	TimeBootUs uint64
	// Quaternion components of autopilot attitude: w, x, y, z (1 0 0 0 is the null-rotation, Hamilton convention).
	Q [4]float32
	// Estimated delay of the attitude data.
	QEstimatedDelayUs uint32
	// X Speed in NED (North, East, Down).
	Vx float32
	// Y Speed in NED (North, East, Down).
	Vy float32
	// Z Speed in NED (North, East, Down).
	Vz float32
	// Estimated delay of the speed data.
	VEstimatedDelayUs uint32
	// Feed forward Z component of angular velocity (positive: yawing to the right). NaN to be ignored. This is to indicate if the autopilot is actively yawing.
	FeedForwardAngularVelocityZ float32
	// Bitmap indicating which estimator outputs are valid.
	EstimatorStatus ESTIMATOR_STATUS_FLAGS `mavenum:"uint16"`
	// The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
	LandedState MAV_LANDED_STATE `mavenum:"uint8"`
	// Z component of angular velocity in NED (North, East, Down). NaN if unknown.
	AngularVelocityZ float32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageAutopilotStateForGimbalDevice) GetID() uint32 {
	return 286
}

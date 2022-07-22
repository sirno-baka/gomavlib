//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package storm32

import (
	"errors"
)

// Enumeration of possible shot modes.
type MAV_QSHOT_MODE uint32

const (
	// Undefined shot mode. Can be used to determine if qshots should be used or not.
	MAV_QSHOT_MODE_UNDEFINED MAV_QSHOT_MODE = 0
	// Start normal gimbal operation. Is usually used to return back from a shot.
	MAV_QSHOT_MODE_DEFAULT MAV_QSHOT_MODE = 1
	// Load and keep safe gimbal position and stop stabilization.
	MAV_QSHOT_MODE_GIMBAL_RETRACT MAV_QSHOT_MODE = 2
	// Load neutral gimbal position and keep it while stabilizing.
	MAV_QSHOT_MODE_GIMBAL_NEUTRAL MAV_QSHOT_MODE = 3
	// Start mission with gimbal control.
	MAV_QSHOT_MODE_GIMBAL_MISSION MAV_QSHOT_MODE = 4
	// Start RC gimbal control.
	MAV_QSHOT_MODE_GIMBAL_RC_CONTROL MAV_QSHOT_MODE = 5
	// Start gimbal tracking the point specified by Lat, Lon, Alt.
	MAV_QSHOT_MODE_POI_TARGETING MAV_QSHOT_MODE = 6
	// Start gimbal tracking the system with specified system ID.
	MAV_QSHOT_MODE_SYSID_TARGETING MAV_QSHOT_MODE = 7
	// Start 2-point cable cam quick shot.
	MAV_QSHOT_MODE_CABLECAM_2POINT MAV_QSHOT_MODE = 8
	// Start gimbal tracking the home location.
	MAV_QSHOT_MODE_HOME_TARGETING MAV_QSHOT_MODE = 9
)

var labels_MAV_QSHOT_MODE = map[MAV_QSHOT_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_QSHOT_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_QSHOT_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_QSHOT_MODE = map[string]MAV_QSHOT_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_QSHOT_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_QSHOT_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_QSHOT_MODE) String() string {
	if l, ok := labels_MAV_QSHOT_MODE[e]; ok {
		return l
	}
	return "invalid value"
}

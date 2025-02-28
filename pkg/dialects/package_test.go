//autogenerated:yes
//nolint:revive
package dialects

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bluenviron/gomavlib/v2/pkg/dialect"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/all"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/ardupilotmega"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/asluav"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/avssuas"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/common"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/cubepilot"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/development"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/icarous"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/matrixpilot"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/minimal"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/paparazzi"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/pythonarraytest"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/standard"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/storm32"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/test"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/ualberta"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/uavionix"
)

func TestDialects(t *testing.T) {
	func() {
		_, err := dialect.NewReadWriter(asluav.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(avssuas.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(all.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(ardupilotmega.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(common.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(cubepilot.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(development.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(icarous.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(matrixpilot.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(minimal.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(paparazzi.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(pythonarraytest.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(standard.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(storm32.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(test.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(uavionix.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewReadWriter(ualberta.Dialect)
		require.NoError(t, err)
	}()
}

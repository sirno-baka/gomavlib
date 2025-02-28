package gomavlib

import (
	"context"
	"io"

	"github.com/tarm/serial"

	"github.com/bluenviron/gomavlib/v2/pkg/autoreconnector"
)

var serialOpenFunc = func(device string, baud int) (io.ReadWriteCloser, error) {
	return serial.OpenPort(&serial.Config{
		Name: device,
		Baud: baud,
	})
}

// EndpointSerial sets up a endpoint that works with a serial port.
type EndpointSerial struct {
	// the name of the device of the serial port (i.e: /dev/ttyUSB0)
	Device string

	// baud rate (i.e: 57600)
	Baud int
}

type endpointSerial struct {
	conf EndpointConf
	io.ReadWriteCloser
}

func (conf EndpointSerial) init(node *Node) (Endpoint, error) {
	// check device existence
	test, err := serialOpenFunc(conf.Device, conf.Baud)
	if err != nil {
		return nil, err
	}
	test.Close()

	t := &endpointSerial{
		conf: conf,
		ReadWriteCloser: autoreconnector.New(
			func(ctx context.Context) (io.ReadWriteCloser, error) {
				return serialOpenFunc(conf.Device, conf.Baud)
			},
		),
	}

	return t, nil
}

func (t *endpointSerial) isEndpoint() {}

func (t *endpointSerial) Conf() EndpointConf {
	return t.conf
}

func (t *endpointSerial) label() string {
	return "serial"
}

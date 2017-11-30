// Code generated by cdpgen. DO NOT EDIT.

// Package deviceorientation implements the DeviceOrientation domain.
package deviceorientation

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the DeviceOrientation domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the DeviceOrientation domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// ClearDeviceOrientationOverride invokes the DeviceOrientation method. Clears the overridden Device Orientation.
func (d *domainClient) ClearDeviceOrientationOverride(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "DeviceOrientation.clearDeviceOrientationOverride", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DeviceOrientation", Op: "ClearDeviceOrientationOverride", Err: err}
	}
	return
}

// SetDeviceOrientationOverride invokes the DeviceOrientation method. Overrides the Device Orientation.
func (d *domainClient) SetDeviceOrientationOverride(ctx context.Context, args *SetDeviceOrientationOverrideArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "DeviceOrientation.setDeviceOrientationOverride", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "DeviceOrientation.setDeviceOrientationOverride", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DeviceOrientation", Op: "SetDeviceOrientationOverride", Err: err}
	}
	return
}

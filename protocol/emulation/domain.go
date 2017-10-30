// Code generated by cdpgen. DO NOT EDIT.

// Package emulation implements the Emulation domain. This domain emulates different environments for the page.
package emulation

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Emulation domain. This domain emulates different environments for the page.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Emulation domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// SetDeviceMetricsOverride invokes the Emulation method. Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (d *domainClient) SetDeviceMetricsOverride(ctx context.Context, args *SetDeviceMetricsOverrideArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setDeviceMetricsOverride", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setDeviceMetricsOverride", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetDeviceMetricsOverride", Err: err}
	}
	return
}

// ClearDeviceMetricsOverride invokes the Emulation method. Clears the overridden device metrics.
func (d *domainClient) ClearDeviceMetricsOverride(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Emulation.clearDeviceMetricsOverride", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "ClearDeviceMetricsOverride", Err: err}
	}
	return
}

// ResetPageScaleFactor invokes the Emulation method. Requests that page scale factor is reset to initial values.
func (d *domainClient) ResetPageScaleFactor(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Emulation.resetPageScaleFactor", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "ResetPageScaleFactor", Err: err}
	}
	return
}

// SetPageScaleFactor invokes the Emulation method. Sets a specified page scale factor.
func (d *domainClient) SetPageScaleFactor(ctx context.Context, args *SetPageScaleFactorArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setPageScaleFactor", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setPageScaleFactor", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetPageScaleFactor", Err: err}
	}
	return
}

// SetVisibleSize invokes the Emulation method. Resizes the frame/viewport of the page. Note that this does not affect the frame's container (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported on Android.
func (d *domainClient) SetVisibleSize(ctx context.Context, args *SetVisibleSizeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setVisibleSize", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setVisibleSize", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetVisibleSize", Err: err}
	}
	return
}

// SetScriptExecutionDisabled invokes the Emulation method. Switches script execution in the page.
func (d *domainClient) SetScriptExecutionDisabled(ctx context.Context, args *SetScriptExecutionDisabledArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setScriptExecutionDisabled", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setScriptExecutionDisabled", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetScriptExecutionDisabled", Err: err}
	}
	return
}

// SetGeolocationOverride invokes the Emulation method. Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (d *domainClient) SetGeolocationOverride(ctx context.Context, args *SetGeolocationOverrideArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setGeolocationOverride", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setGeolocationOverride", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetGeolocationOverride", Err: err}
	}
	return
}

// ClearGeolocationOverride invokes the Emulation method. Clears the overridden Geolocation Position and Error.
func (d *domainClient) ClearGeolocationOverride(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Emulation.clearGeolocationOverride", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "ClearGeolocationOverride", Err: err}
	}
	return
}

// SetTouchEmulationEnabled invokes the Emulation method. Enables touch on platforms which do not support them.
func (d *domainClient) SetTouchEmulationEnabled(ctx context.Context, args *SetTouchEmulationEnabledArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setTouchEmulationEnabled", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setTouchEmulationEnabled", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetTouchEmulationEnabled", Err: err}
	}
	return
}

// SetEmitTouchEventsForMouse invokes the Emulation method.
func (d *domainClient) SetEmitTouchEventsForMouse(ctx context.Context, args *SetEmitTouchEventsForMouseArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setEmitTouchEventsForMouse", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setEmitTouchEventsForMouse", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetEmitTouchEventsForMouse", Err: err}
	}
	return
}

// SetEmulatedMedia invokes the Emulation method. Emulates the given media for CSS media queries.
func (d *domainClient) SetEmulatedMedia(ctx context.Context, args *SetEmulatedMediaArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setEmulatedMedia", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setEmulatedMedia", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetEmulatedMedia", Err: err}
	}
	return
}

// SetCPUThrottlingRate invokes the Emulation method. Enables CPU throttling to emulate slow CPUs.
func (d *domainClient) SetCPUThrottlingRate(ctx context.Context, args *SetCPUThrottlingRateArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setCPUThrottlingRate", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setCPUThrottlingRate", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetCPUThrottlingRate", Err: err}
	}
	return
}

// CanEmulate invokes the Emulation method. Tells whether emulation is supported.
func (d *domainClient) CanEmulate(ctx context.Context) (reply *CanEmulateReply, err error) {
	reply = new(CanEmulateReply)
	err = rpcc.Invoke(ctx, "Emulation.canEmulate", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "CanEmulate", Err: err}
	}
	return
}

// SetVirtualTimePolicy invokes the Emulation method. Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget.
func (d *domainClient) SetVirtualTimePolicy(ctx context.Context, args *SetVirtualTimePolicyArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setVirtualTimePolicy", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setVirtualTimePolicy", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetVirtualTimePolicy", Err: err}
	}
	return
}

// SetNavigatorOverrides invokes the Emulation method. Overrides value returned by the javascript navigator object.
func (d *domainClient) SetNavigatorOverrides(ctx context.Context, args *SetNavigatorOverridesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setNavigatorOverrides", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setNavigatorOverrides", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetNavigatorOverrides", Err: err}
	}
	return
}

// SetDefaultBackgroundColorOverride invokes the Emulation method. Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one.
func (d *domainClient) SetDefaultBackgroundColorOverride(ctx context.Context, args *SetDefaultBackgroundColorOverrideArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Emulation.setDefaultBackgroundColorOverride", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Emulation.setDefaultBackgroundColorOverride", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Emulation", Op: "SetDefaultBackgroundColorOverride", Err: err}
	}
	return
}

func (d *domainClient) VirtualTimeBudgetExpired(ctx context.Context) (VirtualTimeBudgetExpiredClient, error) {
	s, err := rpcc.NewStream(ctx, "Emulation.virtualTimeBudgetExpired", d.conn)
	if err != nil {
		return nil, err
	}
	return &virtualTimeBudgetExpiredClient{Stream: s}, nil
}

type virtualTimeBudgetExpiredClient struct{ rpcc.Stream }

func (c *virtualTimeBudgetExpiredClient) Recv() (*VirtualTimeBudgetExpiredReply, error) {
	event := new(VirtualTimeBudgetExpiredReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Emulation", Op: "VirtualTimeBudgetExpired Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) VirtualTimeAdvanced(ctx context.Context) (VirtualTimeAdvancedClient, error) {
	s, err := rpcc.NewStream(ctx, "Emulation.virtualTimeAdvanced", d.conn)
	if err != nil {
		return nil, err
	}
	return &virtualTimeAdvancedClient{Stream: s}, nil
}

type virtualTimeAdvancedClient struct{ rpcc.Stream }

func (c *virtualTimeAdvancedClient) Recv() (*VirtualTimeAdvancedReply, error) {
	event := new(VirtualTimeAdvancedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Emulation", Op: "VirtualTimeAdvanced Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) VirtualTimePaused(ctx context.Context) (VirtualTimePausedClient, error) {
	s, err := rpcc.NewStream(ctx, "Emulation.virtualTimePaused", d.conn)
	if err != nil {
		return nil, err
	}
	return &virtualTimePausedClient{Stream: s}, nil
}

type virtualTimePausedClient struct{ rpcc.Stream }

func (c *virtualTimePausedClient) Recv() (*VirtualTimePausedReply, error) {
	event := new(VirtualTimePausedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Emulation", Op: "VirtualTimePaused Recv", Err: err}
	}
	return event, nil
}

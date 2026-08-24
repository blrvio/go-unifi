package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

// DeviceRaw is a device's full JSON as a mutable field map. Unlike the typed
// Device struct it preserves EVERY field the controller returns — including ones
// the struct does not model and zero-valued fields the struct would drop via
// omitempty. Use it for radio_table updates on WiFi6+ APs, where the controller
// rejects a PUT (api.err.DeviceNotSupport5gException / InvalidPayload) that omits
// per-radio capability fields it echoed on the GET.
type DeviceRaw = map[string]json.RawMessage

//go:generate go run golang.org/x/tools/cmd/stringer -trimprefix DeviceState -type DeviceState
type DeviceState int

const (
	DeviceStateUnknown          DeviceState = 0
	DeviceStateConnected        DeviceState = 1
	DeviceStatePending          DeviceState = 2
	DeviceStateFirmwareMismatch DeviceState = 3
	DeviceStateUpgrading        DeviceState = 4
	DeviceStateProvisioning     DeviceState = 5
	DeviceStateHeartbeatMissed  DeviceState = 6
	DeviceStateAdopting         DeviceState = 7
	DeviceStateDeleting         DeviceState = 8
	DeviceStateInformError      DeviceState = 9
	DeviceStateAdoptFailed      DeviceState = 10
	DeviceStateIsolated         DeviceState = 11
)

func (c *client) ListDevice(ctx context.Context, site string) ([]Device, error) {
	return c.listDevice(ctx, site)
}

func (c *client) GetDeviceByMAC(ctx context.Context, site, mac string) (*Device, error) {
	return c.getDevice(ctx, site, mac)
}

func (c *client) DeleteDevice(ctx context.Context, site, id string) error {
	return c.deleteDevice(ctx, site, id)
}

func (c *client) CreateDevice(ctx context.Context, site string, d *Device) (*Device, error) {
	return c.createDevice(ctx, site, d)
}

func (c *client) UpdateDevice(ctx context.Context, site string, d *Device) (*Device, error) {
	return c.updateDevice(ctx, site, d)
}

func (c *client) GetDevice(ctx context.Context, site, id string) (*Device, error) {
	devices, err := c.ListDevice(ctx, site)
	if err != nil {
		return nil, err
	}

	for _, d := range devices {
		if d.ID == id {
			return &d, nil
		}
	}

	return nil, ErrNotFound
}

// GetDeviceRaw returns the device's full JSON as a field map, preserving every
// field. Devices are matched by _id via the list endpoint, mirroring GetDevice
// (the by-id stat route expects a MAC, not the _id).
func (c *client) GetDeviceRaw(ctx context.Context, site, id string) (DeviceRaw, error) {
	var respBody struct {
		Data []json.RawMessage `json:"data"`
	}

	err := c.Get(ctx, fmt.Sprintf("s/%s/stat/device", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	for _, raw := range respBody.Data {
		var probe struct {
			ID string `json:"_id"`
		}
		if err := json.Unmarshal(raw, &probe); err != nil {
			return nil, fmt.Errorf("decoding device id: %w", err)
		}
		if probe.ID == id {
			var m DeviceRaw
			if err := json.Unmarshal(raw, &m); err != nil {
				return nil, fmt.Errorf("decoding raw device: %w", err)
			}
			return m, nil
		}
	}

	return nil, ErrNotFound
}

// UpdateDeviceRaw PUTs a raw device object, preserving all fields byte-for-byte,
// and returns the persisted device. The controller echoes an empty body on 10.x,
// so an empty response is treated as success and the device is re-read.
func (c *client) UpdateDeviceRaw(ctx context.Context, site, id string, raw DeviceRaw) (DeviceRaw, error) {
	var respBody struct {
		Data []json.RawMessage `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/device/%s", site, id), raw, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) == 1 {
		var m DeviceRaw
		if err := json.Unmarshal(respBody.Data[0], &m); err != nil {
			return nil, fmt.Errorf("decoding raw device: %w", err)
		}
		return m, nil
	}

	// Empty (or unexpected) echo: re-read the persisted device by id.
	return c.GetDeviceRaw(ctx, site, id)
}

func (c *client) AdoptDevice(ctx context.Context, site, mac string) error {
	reqBody := struct {
		Cmd string `json:"cmd"`
		MAC string `json:"mac"`
	}{
		Cmd: "adopt",
		MAC: mac,
	}

	var respBody struct {
		Meta Meta `json:"Meta"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/cmd/devmgr", site), reqBody, &respBody)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) ForgetDevice(ctx context.Context, site, mac string) error {
	reqBody := struct {
		Cmd  string   `json:"cmd"`
		MACs []string `json:"macs"`
	}{
		Cmd:  "delete-device",
		MACs: []string{mac},
	}

	var respBody struct {
		Meta Meta     `json:"Meta"`
		Data []Device `json:"data"`
	}

	err := c.Post(ctx, fmt.Sprintf("s/%s/cmd/sitemgr", site), reqBody, &respBody)
	if err != nil {
		return err
	}

	return nil
}

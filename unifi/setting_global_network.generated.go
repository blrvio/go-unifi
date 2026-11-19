// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

// just to fix compile issues with the import.
var (
	_ context.Context
	_ fmt.Formatter
	_ json.Marshaler
)

const SettingGlobalNetworkKey = "global_network"

// Self-register this setting's fields factory so the settingFactories registry
// in setting_registry.go stays a 1:1 reflection of the generated catalog and
// can never drift from it by hand.
func init() { //nolint:gochecknoinits
	registerSetting(SettingGlobalNetworkKey, func() any { return &SettingGlobalNetwork{} })
}

type SettingGlobalNetwork struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	DefaultSecurityPosture string `json:"default_security_posture,omitempty"`
}

func (dst *SettingGlobalNetwork) UnmarshalJSON(b []byte) error {
	type Alias SettingGlobalNetwork
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	return nil
}

// GetSettingGlobalNetwork Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingGlobalNetwork(ctx context.Context, site string) (*SettingGlobalNetwork, error) {
	s, f, err := c.GetSetting(ctx, site, SettingGlobalNetworkKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingGlobalNetworkKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingGlobalNetworkKey, s.Key)
	}
	resource, ok := f.(*SettingGlobalNetwork)
	if !ok {
		return nil, fmt.Errorf("unexpected type for setting value. expected: *SettingGlobalNetwork, received: %T", f)
	}
	return resource, nil
}

// UpdateSettingGlobalNetwork Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingGlobalNetwork(ctx context.Context, site string, s *SettingGlobalNetwork) (*SettingGlobalNetwork, error) {
	s.Key = SettingGlobalNetworkKey
	result, err := c.SetSetting(ctx, site, SettingGlobalNetworkKey, s)
	if err != nil {
		return nil, err
	}
	updatedResource, ok := result.(*SettingGlobalNetwork)
	if !ok {
		return nil, fmt.Errorf("unexpected type for setting value. expected: *SettingGlobalNetwork, received: %T", result)
	}
	return updatedResource, nil
}

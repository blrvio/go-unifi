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

const SettingUsgGeoKey = "usg_geo"

// Self-register this setting's fields factory so the settingFactories registry
// in setting_registry.go stays a 1:1 reflection of the generated catalog and
// can never drift from it by hand.
func init() { //nolint:gochecknoinits
	registerSetting(SettingUsgGeoKey, func() any { return &SettingUsgGeo{} })
}

type SettingUsgGeo struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	IPFiltering SettingUsgGeoIPFiltering `json:"ip_filtering,omitempty"`
}

func (dst *SettingUsgGeo) UnmarshalJSON(b []byte) error {
	type Alias SettingUsgGeo
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

type SettingUsgGeoIPFiltering struct {
	Action           string `json:"action,omitempty" validate:"omitempty,oneof=block allow"` // block|allow
	Countries        string `json:"countries,omitempty"`                                     // ^([A-Z]{2})?(,[A-Z]{2}){0,149}$
	Enabled          bool   `json:"enabled"`
	TrafficDirection string `json:"traffic_direction,omitempty" validate:"omitempty,oneof=both ingress egress"` // ^(both|ingress|egress)$
}

func (dst *SettingUsgGeoIPFiltering) UnmarshalJSON(b []byte) error {
	type Alias SettingUsgGeoIPFiltering
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

// GetSettingUsgGeo Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingUsgGeo(ctx context.Context, site string) (*SettingUsgGeo, error) {
	s, f, err := c.GetSetting(ctx, site, SettingUsgGeoKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingUsgGeoKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingUsgGeoKey, s.Key)
	}
	resource, ok := f.(*SettingUsgGeo)
	if !ok {
		return nil, fmt.Errorf("unexpected type for setting value. expected: *SettingUsgGeo, received: %T", f)
	}
	return resource, nil
}

// UpdateSettingUsgGeo Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingUsgGeo(ctx context.Context, site string, s *SettingUsgGeo) (*SettingUsgGeo, error) {
	s.Key = SettingUsgGeoKey
	result, err := c.SetSetting(ctx, site, SettingUsgGeoKey, s)
	if err != nil {
		return nil, err
	}
	updatedResource, ok := result.(*SettingUsgGeo)
	if !ok {
		return nil, fmt.Errorf("unexpected type for setting value. expected: *SettingUsgGeo, received: %T", result)
	}
	return updatedResource, nil
}

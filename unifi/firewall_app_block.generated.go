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

type FirewallAppBlock struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	AppCategoryIDs []int                    `json:"app_category_ids"`
	AppIDs         []int                    `json:"app_ids"`
	ClientMACs     []string                 `json:"client_macs" validate:"omitempty,dive,mac"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Name           string                   `json:"name,omitempty"`
	NetworkIDs     []string                 `json:"network_ids"`
	Schedule       FirewallAppBlockSchedule `json:"schedule,omitempty"`
	TargetType     string                   `json:"target_type,omitempty" validate:"omitempty,oneof=APP APP_CATEGORY"` // APP|APP_CATEGORY
	Type           string                   `json:"type,omitempty" validate:"omitempty,oneof=DEVICE NETWORK"`          // DEVICE|NETWORK
}

func (dst *FirewallAppBlock) UnmarshalJSON(b []byte) error {
	type Alias FirewallAppBlock
	aux := &struct {
		*Alias

		AppCategoryIDs []emptyStringInt `json:"app_category_ids"`
		AppIDs         []emptyStringInt `json:"app_ids"`
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.AppCategoryIDs = make([]int, len(aux.AppCategoryIDs))
	for i, v := range aux.AppCategoryIDs {
		dst.AppCategoryIDs[i] = int(v)
	}
	dst.AppIDs = make([]int, len(aux.AppIDs))
	for i, v := range aux.AppIDs {
		dst.AppIDs[i] = int(v)
	}

	return nil
}

type FirewallAppBlockSchedule struct {
	Date           string   `json:"date,omitempty"`                                                                             // ^$|^(20[0-9]{2})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$
	DateEnd        string   `json:"date_end,omitempty"`                                                                         // ^$|^(20[0-9]{2})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$
	DateStart      string   `json:"date_start,omitempty"`                                                                       // ^$|^(20[0-9]{2})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$
	Mode           string   `json:"mode,omitempty" validate:"omitempty,oneof=ALWAYS EVERY_DAY EVERY_WEEK ONE_TIME_ONLY CUSTOM"` // ALWAYS|EVERY_DAY|EVERY_WEEK|ONE_TIME_ONLY|CUSTOM
	RepeatOnDays   []string `json:"repeat_on_days,omitempty" validate:"omitempty,dive,oneof=mon tue wed thu fri sat sun"`       // mon|tue|wed|thu|fri|sat|sun
	TimeAllDay     bool     `json:"time_all_day"`
	TimeRangeEnd   string   `json:"time_range_end,omitempty"`   // ^[0-9][0-9]:[0-9][0-9]$
	TimeRangeStart string   `json:"time_range_start,omitempty"` // ^[0-9][0-9]:[0-9][0-9]$
}

func (dst *FirewallAppBlockSchedule) UnmarshalJSON(b []byte) error {
	type Alias FirewallAppBlockSchedule
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

func (c *client) listFirewallAppBlock(ctx context.Context, site string) ([]FirewallAppBlock, error) {
	var respBody []FirewallAppBlock

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/firewall-app-blocks", c.apiPaths.ApiV2Path, site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *client) getFirewallAppBlock(ctx context.Context, site, id string) (*FirewallAppBlock, error) {
	var respBody FirewallAppBlock

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/firewall-app-blocks/%s", c.apiPaths.ApiV2Path, site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}
	if respBody.ID == "" {
		return nil, ErrNotFound
	}
	return &respBody, nil
}

func (c *client) deleteFirewallAppBlock(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("%s/site/%s/firewall-app-blocks/%s", c.apiPaths.ApiV2Path, site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createFirewallAppBlock(ctx context.Context, site string, d *FirewallAppBlock) (*FirewallAppBlock, error) {
	var respBody FirewallAppBlock

	err := c.Post(ctx, fmt.Sprintf("%s/site/%s/firewall-app-blocks", c.apiPaths.ApiV2Path, site), d, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *client) updateFirewallAppBlock(ctx context.Context, site string, d *FirewallAppBlock) (*FirewallAppBlock, error) {
	var respBody FirewallAppBlock

	err := c.Put(ctx, fmt.Sprintf("%s/site/%s/firewall-app-blocks/%s", c.apiPaths.ApiV2Path, site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

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

type QOSRule struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Destination     QOSRuleDestination `json:"destination,omitempty"`
	DownloadBurst   string             `json:"download_burst,omitempty" validate:"omitempty,oneof=OFF ON"` // OFF|ON
	Enabled         bool               `json:"enabled"`
	Index           int                `json:"index,omitempty"` // ^[0-9][0-9]?$|^
	Name            string             `json:"name,omitempty"`
	Objective       string             `json:"objective,omitempty"`
	Schedule        QOSRuleSchedule    `json:"schedule,omitempty"`
	Source          QOSRuleSource      `json:"source,omitempty"`
	UploadBurst     string             `json:"upload_burst,omitempty" validate:"omitempty,oneof=OFF ON"` // OFF|ON
	WANOrVPNNetwork string             `json:"wan_or_vpn_network,omitempty"`
}

func (dst *QOSRule) UnmarshalJSON(b []byte) error {
	type Alias QOSRule
	aux := &struct {
		*Alias

		Index emptyStringInt `json:"index"`
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.Index = int(aux.Index)

	return nil
}

type QOSRuleDestination struct {
	AppCategoryIDs   []int    `json:"app_category_ids,omitempty"`                                                                      // ^[0-9][0-9]?$|^
	AppIDs           []int    `json:"app_ids,omitempty"`                                                                               // ^[0-9][0-9]?$|^
	IPs              []string `json:"ips,omitempty" validate:"omitempty,dive,ipv4"`                                                    // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	MatchingTarget   string   `json:"matching_target,omitempty" validate:"omitempty,oneof=ANY APP APP_CATEGORY IP NETWORK REGION WEB"` // ANY|APP|APP_CATEGORY|IP|NETWORK|REGION|WEB
	NetworkIDs       []string `json:"network_ids,omitempty"`
	PortMatchingType string   `json:"port_matching_type,omitempty" validate:"omitempty,oneof=ANY SPECIFIC OBJECT"` // ANY|SPECIFIC|OBJECT
	Regions          []string `json:"regions,omitempty"`
}

func (dst *QOSRuleDestination) UnmarshalJSON(b []byte) error {
	type Alias QOSRuleDestination
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

type QOSRuleSchedule struct {
	Date           string   `json:"date,omitempty"`                                                                             // ^$|^(20[0-9]{2})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$
	DateEnd        string   `json:"date_end,omitempty"`                                                                         // ^$|^(20[0-9]{2})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$
	DateStart      string   `json:"date_start,omitempty"`                                                                       // ^$|^(20[0-9]{2})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$
	Mode           string   `json:"mode,omitempty" validate:"omitempty,oneof=ALWAYS EVERY_DAY EVERY_WEEK ONE_TIME_ONLY CUSTOM"` // ALWAYS|EVERY_DAY|EVERY_WEEK|ONE_TIME_ONLY|CUSTOM
	RepeatOnDays   []string `json:"repeat_on_days,omitempty" validate:"omitempty,dive,oneof=mon tue wed thu fri sat sun"`       // mon|tue|wed|thu|fri|sat|sun
	TimeAllDay     bool     `json:"time_all_day"`
	TimeRangeEnd   string   `json:"time_range_end,omitempty"`   // ^[0-9][0-9]:[0-9][0-9]$
	TimeRangeStart string   `json:"time_range_start,omitempty"` // ^[0-9][0-9]:[0-9][0-9]$
}

func (dst *QOSRuleSchedule) UnmarshalJSON(b []byte) error {
	type Alias QOSRuleSchedule
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

type QOSRuleSource struct {
	ClientMACs       []string `json:"client_macs,omitempty" validate:"omitempty,dive,mac"`                                   // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	IPs              []string `json:"ips,omitempty" validate:"omitempty,dive,ipv4"`                                          // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	MatchingTarget   string   `json:"matching_target,omitempty" validate:"omitempty,oneof=ANY CLIENT NETWORK IP MAC REGION"` // ANY|CLIENT|NETWORK|IP|MAC|REGION
	NetworkIDs       []string `json:"network_ids,omitempty"`
	PortMatchingType string   `json:"port_matching_type,omitempty" validate:"omitempty,oneof=ANY SPECIFIC OBJECT"` // ANY|SPECIFIC|OBJECT
	Regions          []string `json:"regions,omitempty"`
}

func (dst *QOSRuleSource) UnmarshalJSON(b []byte) error {
	type Alias QOSRuleSource
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

func (c *client) listQOSRule(ctx context.Context, site string) ([]QOSRule, error) {
	var respBody []QOSRule

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/qos-rules", c.apiPaths.ApiV2Path, site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *client) getQOSRule(ctx context.Context, site, id string) (*QOSRule, error) {
	var respBody QOSRule

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/qos-rules/%s", c.apiPaths.ApiV2Path, site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}
	if respBody.ID == "" {
		return nil, ErrNotFound
	}
	return &respBody, nil
}

func (c *client) deleteQOSRule(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("%s/site/%s/qos-rules/%s", c.apiPaths.ApiV2Path, site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createQOSRule(ctx context.Context, site string, d *QOSRule) (*QOSRule, error) {
	var respBody QOSRule

	err := c.Post(ctx, fmt.Sprintf("%s/site/%s/qos-rules", c.apiPaths.ApiV2Path, site), d, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *client) updateQOSRule(ctx context.Context, site string, d *QOSRule) (*QOSRule, error) {
	var respBody QOSRule

	err := c.Put(ctx, fmt.Sprintf("%s/site/%s/qos-rules/%s", c.apiPaths.ApiV2Path, site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

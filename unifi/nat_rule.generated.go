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

type NatRule struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Description           string                   `json:"description,omitempty"`
	DestinationFilter     NatRuleDestinationFilter `json:"destination_filter,omitempty"`
	Enabled               bool                     `json:"enabled"`
	Exclude               bool                     `json:"exclude"`
	IPAddress             string                   `json:"ip_address,omitempty"`
	IPVersion             string                   `json:"ip_version,omitempty" validate:"omitempty,oneof=IPV4 IPV6"` // IPV4|IPV6
	InInterface           string                   `json:"in_interface,omitempty"`
	IsPredefined          bool                     `json:"is_predefined"`
	Logging               bool                     `json:"logging"`
	OutInterface          string                   `json:"out_interface,omitempty"`
	Port                  string                   `json:"port,omitempty"`
	PppoeUseBaseInterface bool                     `json:"pppoe_use_base_interface"`
	Protocol              string                   `json:"protocol,omitempty" validate:"omitempty,oneof=all tcp udp tcp_udp"` // all|tcp|udp|tcp_udp
	RuleIndex             int                      `json:"rule_index,omitempty"`
	SettingPreference     string                   `json:"setting_preference,omitempty" validate:"omitempty,oneof=auto manual"` // auto|manual
	SourceFilter          NatRuleSourceFilter      `json:"source_filter,omitempty"`
	Type                  string                   `json:"type,omitempty" validate:"omitempty,oneof=MASQUERADE SNAT DNAT"` // MASQUERADE|SNAT|DNAT
}

func (dst *NatRule) UnmarshalJSON(b []byte) error {
	type Alias NatRule
	aux := &struct {
		*Alias

		RuleIndex emptyStringInt `json:"rule_index"`
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.RuleIndex = int(aux.RuleIndex)

	return nil
}

type NatRuleDestinationFilter struct {
	Address          string   `json:"address,omitempty"`
	FilterType       string   `json:"filter_type,omitempty" validate:"omitempty,oneof=NONE ADDRESS_AND_PORT NETWORK_CONF FIREWALL_GROUPS IID_AND_PORT"` // NONE|ADDRESS_AND_PORT|NETWORK_CONF|FIREWALL_GROUPS|IID_AND_PORT
	FirewallGroupIDs []string `json:"firewall_group_ids,omitempty"`
	InvertAddress    bool     `json:"invert_address"`
	InvertPort       bool     `json:"invert_port"`
	NetworkConfID    string   `json:"network_conf_id,omitempty"`
	Port             string   `json:"port,omitempty"`
}

func (dst *NatRuleDestinationFilter) UnmarshalJSON(b []byte) error {
	type Alias NatRuleDestinationFilter
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

type NatRuleSourceFilter struct {
	Address          string   `json:"address,omitempty"`
	FilterType       string   `json:"filter_type,omitempty" validate:"omitempty,oneof=NONE ADDRESS_AND_PORT NETWORK_CONF FIREWALL_GROUPS IID_AND_PORT"` // NONE|ADDRESS_AND_PORT|NETWORK_CONF|FIREWALL_GROUPS|IID_AND_PORT
	FirewallGroupIDs []string `json:"firewall_group_ids,omitempty"`
	InvertAddress    bool     `json:"invert_address"`
	InvertPort       bool     `json:"invert_port"`
	NetworkConfID    string   `json:"network_conf_id,omitempty"`
	Port             string   `json:"port,omitempty"`
}

func (dst *NatRuleSourceFilter) UnmarshalJSON(b []byte) error {
	type Alias NatRuleSourceFilter
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

func (c *client) listNatRule(ctx context.Context, site string) ([]NatRule, error) {
	var respBody []NatRule

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/nat", c.apiPaths.ApiV2Path, site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *client) getNatRule(ctx context.Context, site, id string) (*NatRule, error) {
	var respBody NatRule

	err := c.Get(ctx, fmt.Sprintf("%s/site/%s/nat/%s", c.apiPaths.ApiV2Path, site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}
	if respBody.ID == "" {
		return nil, ErrNotFound
	}
	return &respBody, nil
}

func (c *client) deleteNatRule(ctx context.Context, site, id string) error {
	err := c.Delete(ctx, fmt.Sprintf("%s/site/%s/nat/%s", c.apiPaths.ApiV2Path, site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) createNatRule(ctx context.Context, site string, d *NatRule) (*NatRule, error) {
	var respBody NatRule

	err := c.Post(ctx, fmt.Sprintf("%s/site/%s/nat", c.apiPaths.ApiV2Path, site), d, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *client) updateNatRule(ctx context.Context, site string, d *NatRule) (*NatRule, error) {
	var respBody NatRule

	err := c.Put(ctx, fmt.Sprintf("%s/site/%s/nat/%s", c.apiPaths.ApiV2Path, site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

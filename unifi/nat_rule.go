package unifi

import (
	"context"
)

func (c *client) ListNatRule(ctx context.Context, site string) ([]NatRule, error) {
	return c.listNatRule(ctx, site)
}

func (c *client) GetNatRule(ctx context.Context, site, id string) (*NatRule, error) {
	// The v2 firewall-app-rules collection has no GET-by-id (it returns HTTP 405), so filter client-side.
	rules, err := c.listNatRule(ctx, site)
	if err != nil {
		return nil, err
	}

	for i := range rules {
		if rules[i].ID == id {
			return &rules[i], nil
		}
	}
	return nil, ErrNotFound
}

func (c *client) CreateNatRule(ctx context.Context, site string, d *NatRule) (*NatRule, error) {
	return c.createNatRule(ctx, site, d)
}

func (c *client) UpdateNatRule(ctx context.Context, site string, d *NatRule) (*NatRule, error) {
	return c.updateNatRule(ctx, site, d)
}

func (c *client) DeleteNatRule(ctx context.Context, site, id string) error {
	return c.deleteNatRule(ctx, site, id)
}

package unifi

import (
	"context"
)

func (c *client) ListFirewallAppBlock(ctx context.Context, site string) ([]FirewallAppBlock, error) {
	return c.listFirewallAppBlock(ctx, site)
}

func (c *client) GetFirewallAppBlock(ctx context.Context, site, id string) (*FirewallAppBlock, error) {
	// The v2 firewall-app-blocks collection has no GET-by-id (it returns HTTP 405), so filter client-side.
	blocks, err := c.listFirewallAppBlock(ctx, site)
	if err != nil {
		return nil, err
	}

	for i := range blocks {
		if blocks[i].ID == id {
			return &blocks[i], nil
		}
	}
	return nil, ErrNotFound
}

func (c *client) CreateFirewallAppBlock(ctx context.Context, site string, d *FirewallAppBlock) (*FirewallAppBlock, error) {
	return c.createFirewallAppBlock(ctx, site, d)
}

func (c *client) UpdateFirewallAppBlock(ctx context.Context, site string, d *FirewallAppBlock) (*FirewallAppBlock, error) {
	return c.updateFirewallAppBlock(ctx, site, d)
}

func (c *client) DeleteFirewallAppBlock(ctx context.Context, site, id string) error {
	return c.deleteFirewallAppBlock(ctx, site, id)
}

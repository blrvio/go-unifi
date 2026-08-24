package unifi

import (
	"context"
)

func (c *client) ListQOSRule(ctx context.Context, site string) ([]QOSRule, error) {
	return c.listQOSRule(ctx, site)
}

func (c *client) GetQOSRule(ctx context.Context, site, id string) (*QOSRule, error) {
	// The v2 qos-rules collection has no GET-by-id (it returns HTTP 405), so
	// filter client-side. Mirrors GetAPGroup / GetDNSRecord.
	rules, err := c.listQOSRule(ctx, site)
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

func (c *client) CreateQOSRule(ctx context.Context, site string, d *QOSRule) (*QOSRule, error) {
	return c.createQOSRule(ctx, site, d)
}

func (c *client) UpdateQOSRule(ctx context.Context, site string, d *QOSRule) (*QOSRule, error) {
	return c.updateQOSRule(ctx, site, d)
}

func (c *client) DeleteQOSRule(ctx context.Context, site, id string) error {
	return c.deleteQOSRule(ctx, site, id)
}

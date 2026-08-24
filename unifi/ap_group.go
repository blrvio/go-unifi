package unifi

import (
	"context"
)

func (c *client) ListAPGroup(ctx context.Context, site string) ([]APGroup, error) {
	return c.listAPGroup(ctx, site)
}

func (c *client) CreateAPGroup(ctx context.Context, site string, d *APGroup) (*APGroup, error) {
	return c.createAPGroup(ctx, site, d)
}

func (c *client) GetAPGroup(ctx context.Context, site, id string) (*APGroup, error) {
	// The v2 apgroups collection has no GET-by-id (it returns HTTP 405), so filter client-side.
	groups, err := c.listAPGroup(ctx, site)
	if err != nil {
		return nil, err
	}

	for i := range groups {
		if groups[i].ID == id {
			return &groups[i], nil
		}
	}
	return nil, ErrNotFound
}

func (c *client) DeleteAPGroup(ctx context.Context, site, id string) error {
	return c.deleteAPGroup(ctx, site, id)
}

func (c *client) UpdateAPGroup(ctx context.Context, site string, d *APGroup) (*APGroup, error) {
	return c.updateAPGroup(ctx, site, d)
}

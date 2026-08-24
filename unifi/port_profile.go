package unifi

import (
	"context"
	"fmt"
)

func (c *client) ListPortProfile(ctx context.Context, site string) ([]PortProfile, error) {
	return c.listPortProfile(ctx, site)
}

func (c *client) GetPortProfile(ctx context.Context, site, id string) (*PortProfile, error) {
	return c.getPortProfile(ctx, site, id)
}

func (c *client) DeletePortProfile(ctx context.Context, site, id string) error {
	return c.deletePortProfile(ctx, site, id)
}

func (c *client) CreatePortProfile(ctx context.Context, site string, d *PortProfile) (*PortProfile, error) {
	return c.createPortProfile(ctx, site, d)
}

// UpdatePortProfile wraps the generated update to tolerate the empty write-echo
// that UniFi Network 10.x returns for a successful PUT to rest/portconf/{id}.
// The generated updatePortProfile rejects a zero-length data array as
// "expected 1 PortProfile, got 0"; here we treat rc=ok with no echo as success
// and re-fetch the profile by ID so callers still get the persisted object.
func (c *client) UpdatePortProfile(ctx context.Context, site string, d *PortProfile) (*PortProfile, error) {
	var respBody struct {
		Meta Meta          `json:"meta"`
		Data []PortProfile `json:"data"`
	}

	err := c.Put(ctx, fmt.Sprintf("s/%s/rest/portconf/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	switch len(respBody.Data) {
	case 1:
		updated := respBody.Data[0]
		return &updated, nil
	case 0:
		// Empty but successful echo: re-read the persisted resource by ID.
		return c.getPortProfile(ctx, site, d.ID)
	default:
		return nil, fmt.Errorf("unexpected response: expected 1 PortProfile, got %d", len(respBody.Data))
	}
}

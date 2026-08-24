package unifi

import (
	"context"
	"fmt"
)

func (c *client) ListContentFiltering(ctx context.Context, site string) ([]ContentFiltering, error) {
	return c.listContentFiltering(ctx, site)
}

func (c *client) DeleteContentFiltering(ctx context.Context, site, id string) error {
	return c.deleteContentFiltering(ctx, site, id)
}

// CreateContentFiltering overrides the generated create, which POSTs to the
// collection (`content-filtering`). On UniFi Network 10.x that path is not a
// creation endpoint (HTTP 405); new content-filtering rules are created via a
// dedicated `content-filtering/create` action that returns the created rule as a
// raw (non-enveloped) object.
func (c *client) CreateContentFiltering(ctx context.Context, site string, d *ContentFiltering) (*ContentFiltering, error) {
	var respBody ContentFiltering

	err := c.Post(ctx, fmt.Sprintf("%s/site/%s/content-filtering/create", c.apiPaths.ApiV2Path, site), d, &respBody)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *client) UpdateContentFiltering(ctx context.Context, site string, d *ContentFiltering) (*ContentFiltering, error) {
	return c.updateContentFiltering(ctx, site, d)
}

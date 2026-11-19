package unifi

import (
	"context"
	"fmt"
)

// WANLoadBalancing is the site-wide WAN mode (Settings > Internet): failover-only
// or load-balancing, with per-WAN priority and weight. It is a v2 singleton read
// and written in full at wan/load-balancing/configuration.
type WANLoadBalancing struct {
	Mode          string                      `json:"mode"` // FAILOVER_ONLY|DISTRIBUTED
	WANInterfaces []WANLoadBalancingInterface `json:"wan_interfaces"`
}

type WANLoadBalancingInterface struct {
	Mode            string `json:"mode"` // FAILOVER_ONLY|DISTRIBUTED
	Name            string `json:"name,omitempty"`
	Priority        int    `json:"priority"`
	WANNetworkgroup string `json:"wan_networkgroup"` // WAN, WAN2, ...
	Weight          int    `json:"weight"`
}

func (c *client) wanLoadBalancingPath(site string) string {
	return fmt.Sprintf("%s/site/%s/wan/load-balancing/configuration", c.apiPaths.ApiV2Path, site)
}

func (c *client) GetWANLoadBalancing(ctx context.Context, site string) (*WANLoadBalancing, error) {
	var respBody WANLoadBalancing
	if err := c.Get(ctx, c.wanLoadBalancingPath(site), nil, &respBody); err != nil {
		return nil, err
	}
	if respBody.Mode == "" {
		return nil, ErrNotFound
	}
	return &respBody, nil
}

func (c *client) UpdateWANLoadBalancing(ctx context.Context, site string, d *WANLoadBalancing) (*WANLoadBalancing, error) {
	var respBody WANLoadBalancing
	if err := c.Put(ctx, c.wanLoadBalancingPath(site), d, &respBody); err != nil {
		return nil, err
	}
	if respBody.Mode == "" {
		// Some v2 singletons answer a PUT with an empty body; re-read to return the stored value.
		return c.GetWANLoadBalancing(ctx, site)
	}
	return &respBody, nil
}

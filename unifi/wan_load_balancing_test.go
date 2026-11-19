package unifi //nolint: testpackage

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Shape observed on a UniFi Network 11.0.99 console in failover mode.
const wanLBBody = `{"mode":"FAILOVER_ONLY","wan_interfaces":[` +
	`{"mode":"DISTRIBUTED","name":"Internet Principal 1","priority":1,"wan_networkgroup":"WAN","weight":50},` +
	`{"mode":"FAILOVER_ONLY","name":"Internet 2","priority":2,"wan_networkgroup":"WAN2","weight":50}]}`

func TestGetWANLoadBalancing(t *testing.T) {
	t.Parallel()

	path := apiV2("site/default/wan/load-balancing/configuration")
	cs := newControllerServer(t, route{path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		_, _ = w.Write([]byte(wanLBBody))
	}})

	got, err := cs.client().GetWANLoadBalancing(context.Background(), "default")
	require.NoError(t, err)
	assert.Equal(t, "FAILOVER_ONLY", got.Mode)
	require.Len(t, got.WANInterfaces, 2)
	assert.Equal(t, "WAN2", got.WANInterfaces[1].WANNetworkgroup)
	assert.Equal(t, 2, got.WANInterfaces[1].Priority)
}

func TestUpdateWANLoadBalancing(t *testing.T) {
	t.Parallel()

	path := apiV2("site/default/wan/load-balancing/configuration")
	cases := map[string]struct {
		putEcho      string
		wantRequests int
	}{
		"echoed body":                {putEcho: wanLBBody, wantRequests: 1},
		"empty body re-reads config": {putEcho: ``, wantRequests: 2},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			cs := newControllerServer(t, route{path, func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPut {
					var sent WANLoadBalancing
					assert.NoError(t, json.NewDecoder(r.Body).Decode(&sent))
					assert.Equal(t, "FAILOVER_ONLY", sent.Mode)
					_, _ = w.Write([]byte(tc.putEcho))
					return
				}
				_, _ = w.Write([]byte(wanLBBody))
			}})

			got, err := cs.client().UpdateWANLoadBalancing(context.Background(), "default", &WANLoadBalancing{Mode: "FAILOVER_ONLY"})
			require.NoError(t, err)
			assert.Equal(t, "FAILOVER_ONLY", got.Mode)
			assert.Equal(t, tc.wantRequests, cs.countRequestsTo(path))
		})
	}
}

package unifi //nolint: testpackage

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNatRuleRoutes(t *testing.T) {
	t.Parallel()

	const site = "default"
	coll := apiV2("site/" + site + "/nat")
	item := coll + "/r1"
	one := `{"_id":"r1","description":"dnat","enabled":true,"type":"DNAT","ip_version":"IPV4","in_interface":"wan1",` +
		`"protocol":"tcp","port":"8443","ip_address":"10.0.0.5","rule_index":1,"logging":false,"exclude":false,` +
		`"source_filter":{"filter_type":"NONE"},"destination_filter":{"filter_type":"ADDRESS_AND_PORT","port":"443","invert_port":false}}`

	cases := map[string]struct {
		path, method, body string
		call               func(c *client) (*NatRule, error)
	}{
		"get filters the list": {coll, http.MethodGet, "[" + one + "]", func(c *client) (*NatRule, error) {
			return c.GetNatRule(context.Background(), site, "r1")
		}},
		"create posts to the collection": {coll, http.MethodPost, one, func(c *client) (*NatRule, error) {
			return c.CreateNatRule(context.Background(), site, &NatRule{Type: "DNAT"})
		}},
		"update puts to the item": {item, http.MethodPut, one, func(c *client) (*NatRule, error) {
			return c.UpdateNatRule(context.Background(), site, &NatRule{ID: "r1", Type: "DNAT"})
		}},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			cs := newControllerServer(t, route{tc.path, func(w http.ResponseWriter, _ *http.Request) {
				_, _ = w.Write([]byte(tc.body))
			}})
			got, err := tc.call(cs.client())
			require.NoError(t, err)
			assert.Equal(t, tc.method, cs.lastRequest().Method)
			assert.Equal(t, tc.path, cs.lastRequest().Path)
			assert.Equal(t, "DNAT", got.Type)
			assert.Equal(t, "443", got.DestinationFilter.Port)
			assert.Equal(t, 1, got.RuleIndex)
		})
	}
}

// Like the UI, only the keys the filter_type uses are sent; booleans that disable a rule are always sent.
func TestNatRuleMarshalOmitsUnusedFilterKeys(t *testing.T) {
	t.Parallel()

	b, err := json.Marshal(NatRule{Type: "MASQUERADE", SourceFilter: NatRuleSourceFilter{FilterType: "NONE"}})
	require.NoError(t, err)
	var m map[string]any
	require.NoError(t, json.Unmarshal(b, &m))
	assert.Equal(t, false, m["enabled"])
	src, ok := m["source_filter"].(map[string]any)
	require.True(t, ok)
	assert.NotContains(t, src, "network_conf_id")
	assert.NotContains(t, src, "firewall_group_ids")
}

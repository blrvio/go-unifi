package unifi //nolint: testpackage

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Shape observed on a UniFi Network 11.0.99 console (CyberSecure > Simple App Blocking).
const appBlockListBody = `[{"_id":"ab1","app_category_ids":[13,6,11],"app_ids":[],"client_macs":["6c:3c:7c:a7:ab:32"],` +
	`"name":"[Block:internet] Impressora","network_ids":[],"schedule":{"mode":"ALWAYS","repeat_on_days":[],"time_all_day":false},` +
	`"target_type":"APP_CATEGORY","type":"DEVICE"}]`

func TestGetFirewallAppBlockFiltersList(t *testing.T) {
	t.Parallel()

	const site = "default"
	listPath := apiV2("site/" + site + "/firewall-app-blocks")

	cases := map[string]struct {
		id      string
		wantErr error
	}{
		"found":   {id: "ab1"},
		"missing": {id: "nope", wantErr: ErrNotFound},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			cs := newControllerServer(t, route{listPath, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				_, _ = w.Write([]byte(appBlockListBody))
			}})

			got, err := cs.client().GetFirewallAppBlock(context.Background(), site, tc.id)
			assert.Equal(t, listPath, cs.lastRequest().Path, "GET-by-id is 405 on the controller; must read the list")
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "[Block:internet] Impressora", got.Name)
			assert.Equal(t, []int{13, 6, 11}, got.AppCategoryIDs)
			assert.Equal(t, []string{"6c:3c:7c:a7:ab:32"}, got.ClientMACs)
			assert.Equal(t, "ALWAYS", got.Schedule.Mode)
			assert.Equal(t, "APP_CATEGORY", got.TargetType)
			assert.Equal(t, "DEVICE", got.Type)
		})
	}
}

func TestFirewallAppBlockWriteRoutes(t *testing.T) {
	t.Parallel()

	const site = "default"
	coll := apiV2("site/" + site + "/firewall-app-blocks")
	item := coll + "/ab1"
	echo := `{"_id":"ab1","name":"x","type":"DEVICE","target_type":"APP","app_ids":[1],"app_category_ids":[],"client_macs":[],"network_ids":[]}`

	cases := map[string]struct {
		path, method string
		call         func(c *client) error
	}{
		"create posts to the collection": {coll, http.MethodPost, func(c *client) error {
			_, err := c.CreateFirewallAppBlock(context.Background(), site, &FirewallAppBlock{Name: "x"})
			return err
		}},
		"update puts to the item": {item, http.MethodPut, func(c *client) error {
			_, err := c.UpdateFirewallAppBlock(context.Background(), site, &FirewallAppBlock{ID: "ab1", Name: "x"})
			return err
		}},
		"delete deletes the item": {item, http.MethodDelete, func(c *client) error {
			return c.DeleteFirewallAppBlock(context.Background(), site, "ab1")
		}},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			cs := newControllerServer(t, route{tc.path, func(w http.ResponseWriter, _ *http.Request) {
				_, _ = w.Write([]byte(echo))
			}})
			require.NoError(t, tc.call(cs.client()))
			assert.Equal(t, tc.method, cs.lastRequest().Method)
			assert.Equal(t, tc.path, cs.lastRequest().Path)
		})
	}
}

// Empty lists must go over the wire as [] so an update can clear them.
func TestFirewallAppBlockMarshalsEmptyLists(t *testing.T) {
	t.Parallel()

	b, err := json.Marshal(FirewallAppBlock{Name: "x", AppCategoryIDs: []int{}, AppIDs: []int{}, ClientMACs: []string{}, NetworkIDs: []string{}})
	require.NoError(t, err)
	var m map[string]any
	require.NoError(t, json.Unmarshal(b, &m))
	for _, k := range []string{"app_category_ids", "app_ids", "client_macs", "network_ids"} {
		assert.Contains(t, m, k)
	}
}

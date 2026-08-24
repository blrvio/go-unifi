package unifi //nolint: testpackage

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCreateContentFilteringUsesCreateRoute verifies that creating a content
// filtering rule POSTs to the dedicated `content-filtering/create` action, not
// the collection. On UniFi Network 10.x the collection POST returns HTTP 405;
// the create action returns the created rule as a raw (non-enveloped) object.
func TestCreateContentFilteringUsesCreateRoute(t *testing.T) {
	t.Parallel()

	const site = "default"
	createPath := apiV2("site/" + site + "/content-filtering/create")

	cs := newControllerServer(t, route{createPath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		// Raw object, as returned by the 10.5.67 controller.
		_, _ = w.Write([]byte(`{"_id":"cf1","name":"vlan-iot","enabled":true,"categories":["FAMILY"],"network_ids":["n1"]}`))
	}})
	c := cs.client()

	got, err := c.CreateContentFiltering(context.Background(), site, &ContentFiltering{
		Name:       "vlan-iot",
		Enabled:    true,
		Categories: []string{"FAMILY"},
		NetworkIDs: []string{"n1"},
	})
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, http.MethodPost, cs.lastRequest().Method)
	assert.Equal(t, createPath, cs.lastRequest().Path)
	assert.Equal(t, "cf1", got.ID)
	assert.Equal(t, "vlan-iot", got.Name)
}

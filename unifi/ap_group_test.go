package unifi //nolint: testpackage

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetAPGroupUsesListRoute verifies that GetAPGroup reads via the apgroups
// LIST collection and filters client-side. The v2 collection has no GET-by-id
// (it returns HTTP 405 on 10.5.67), so a per-id GET would fail; the list route
// is the only one the controller serves.
func TestGetAPGroupUsesListRoute(t *testing.T) {
	t.Parallel()

	const site = "default"
	listPath := apiV2("site/" + site + "/apgroups")
	byIDPath := listPath + "/g1"

	cs := newControllerServer(t,
		route{listPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			_, _ = w.Write([]byte(`[{"_id":"g1","name":"All APs","attr_no_delete":true,"device_macs":["aa:bb:cc:dd:ee:ff"]}]`))
		}},
		// GET-by-id must never be hit; if it is, fail loudly like the real 405.
		route{byIDPath, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}},
	)
	c := cs.client()

	got, err := c.GetAPGroup(context.Background(), site, "g1")
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, http.MethodGet, cs.lastRequest().Method)
	assert.Equal(t, listPath, cs.lastRequest().Path, "GetAPGroup must read via the LIST route, not GET-by-id")
	assert.Equal(t, "g1", got.ID)
	assert.Equal(t, "All APs", got.Name)
}

// TestGetAPGroupNotFound verifies an absent id yields ErrNotFound (not a nil,nil).
func TestGetAPGroupNotFound(t *testing.T) {
	t.Parallel()

	const site = "default"
	listPath := apiV2("site/" + site + "/apgroups")

	cs := newControllerServer(t, route{listPath, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`[{"_id":"g1","name":"All APs"}]`))
	}})
	c := cs.client()

	got, err := c.GetAPGroup(context.Background(), site, "does-not-exist")
	assert.Nil(t, got)
	assert.ErrorIs(t, err, ErrNotFound)
}

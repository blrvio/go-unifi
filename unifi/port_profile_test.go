package unifi //nolint: testpackage

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUpdatePortProfileEmptyEcho reproduces the UniFi Network 10.x behavior where
// a successful PUT to rest/portconf/{id} returns rc=ok with an EMPTY data array.
// The generated update rejects that as "expected 1 PortProfile, got 0"; the
// hand-written wrapper must instead treat it as success and re-fetch by ID.
func TestUpdatePortProfileEmptyEcho(t *testing.T) {
	t.Parallel()

	const (
		site = "default"
		id   = "pp1"
	)
	path := apiV1Path("s/" + site + "/rest/portconf/" + id)

	cs := newControllerServer(t,
		route{path, func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPut {
				// Empty write-echo, as observed on 10.5.67.
				_, _ = w.Write([]byte(`{"meta":{"rc":"ok"},"data":[]}`))
				return
			}
			// GET re-fetch returns the persisted profile.
			_, _ = w.Write([]byte(`{"meta":{"rc":"ok"},"data":[{"_id":"` + id + `","name":"TF-Disabled","forward":"all"}]}`))
		}},
	)
	c := cs.client()

	got, err := c.UpdatePortProfile(context.Background(), site, &PortProfile{ID: id, Name: "TF-Disabled"})
	require.NoError(t, err, "empty write-echo must not surface as an error")
	require.NotNil(t, got)
	assert.Equal(t, id, got.ID)
	assert.Equal(t, "TF-Disabled", got.Name)
}

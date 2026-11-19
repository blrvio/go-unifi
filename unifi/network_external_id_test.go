package unifi //nolint: testpackage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// On 11.x every networkconf carries the Official API UUID as external_id (G5).
func TestNetworkExternalID(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		body, want string
		wantSent   bool
	}{
		"present on read": {body: `{"_id":"n1","purpose":"remote-user-vpn","external_id":"663bc70c-85bc-415f-9316-19dbdc2353c0"}`, want: "663bc70c-85bc-415f-9316-19dbdc2353c0", wantSent: true},
		"absent (create)": {body: `{"name":"new"}`, want: "", wantSent: false},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var n Network
			require.NoError(t, json.Unmarshal([]byte(tc.body), &n))
			assert.Equal(t, tc.want, n.ExternalID)

			b, err := json.Marshal(n)
			require.NoError(t, err)
			var m map[string]any
			require.NoError(t, json.Unmarshal(b, &m))
			_, sent := m["external_id"]
			assert.Equal(t, tc.wantSent, sent, "an empty external_id must never be sent")
		})
	}
}

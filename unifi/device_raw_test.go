package unifi //nolint: testpackage

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetDeviceRawPreservesUnknownFields verifies GetDeviceRaw returns EVERY
// field the controller sends — including radio_table capability fields the typed
// Device struct would drop — so a subsequent PUT can echo them back and avoid
// api.err.DeviceNotSupport5gException.
func TestGetDeviceRawPreservesUnknownFields(t *testing.T) {
	t.Parallel()

	const site = "default"
	path := apiV1Path("s/" + site + "/stat/device")

	// A radio object with a field (some_future_cap) the typed struct does NOT model.
	body := `{"meta":{"rc":"ok"},"data":[` +
		`{"_id":"dev1","model":"UAP6MP","radio_table":[` +
		`{"radio":"na","channel":40,"ht":"80","nss":4,"radio_caps":251805700,"min_rssi_enabled":false,"current_antenna_gain":0,"some_future_cap":123}` +
		`]}]}`

	cs := newControllerServer(t, route{path, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(body))
	}})
	c := cs.client()

	raw, err := c.GetDeviceRaw(context.Background(), site, "dev1")
	require.NoError(t, err)
	require.NotNil(t, raw)

	var radios []map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(raw["radio_table"], &radios))
	require.Len(t, radios, 1)
	na := radios[0]

	// Every original key survives, including the unmodeled one and zero-valued ones.
	for _, k := range []string{"nss", "radio_caps", "min_rssi_enabled", "current_antenna_gain", "some_future_cap", "ht", "channel"} {
		_, ok := na[k]
		assert.Truef(t, ok, "raw radio must preserve %q", k)
	}
	assert.Equal(t, "123", string(na["some_future_cap"]), "unmodeled field value must be byte-preserved")
	assert.Equal(t, "0", string(na["current_antenna_gain"]), "zero-valued field must be preserved")
	assert.Equal(t, "false", string(na["min_rssi_enabled"]), "false bool must be preserved")
}

// TestUpdateDeviceRawEmptyEchoReReads verifies UpdateDeviceRaw treats the empty
// PUT echo (returned by 10.x) as success and re-reads the device by id.
func TestUpdateDeviceRawEmptyEchoReReads(t *testing.T) {
	t.Parallel()

	const site = "default"
	putPath := apiV1Path("s/" + site + "/rest/device/dev1")
	listPath := apiV1Path("s/" + site + "/stat/device")

	cs := newControllerServer(t,
		route{putPath, func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"meta":{"rc":"ok"},"data":[]}`)) // empty echo
		}},
		route{listPath, func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"meta":{"rc":"ok"},"data":[{"_id":"dev1","name":"AP","radio_table":[{"radio":"ng","channel":1}]}]}`))
		}},
	)
	c := cs.client()

	got, err := c.UpdateDeviceRaw(context.Background(), site, "dev1", DeviceRaw{"_id": json.RawMessage(`"dev1"`)})
	require.NoError(t, err, "empty PUT echo must not surface as an error")
	require.NotNil(t, got)
	assert.Equal(t, `"AP"`, string(got["name"]), "device must be re-read after the empty echo")
}

package unifi //nolint: testpackage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// realUAP6MP5GHzRadio is the actual `na` (5 GHz) radio object returned by a
// UAP6MP (U6 Pro) on UniFi Network 10.5.67. It carries capability fields
// (nss, radio_caps, radio_caps2, has_ht160, max/min_txpower, ...) that the
// controller REQUIRES to be echoed back on a radio_table PUT.
const realUAP6MP5GHzRadio = `{"antenna_gain":6,"builtin_antenna":true,"has_dfs":true,"vwire_enabled":true,` +
	`"channel":40,"max_txpower":26,"min_rssi_enabled":false,"is_11ac":true,"builtin_ant_gain":6,` +
	`"ht":"80","has_ht160":true,"radio":"na","nss":4,"tx_power_mode":"auto","is_11ax":true,` +
	`"name":"wifi1","min_txpower":6,"has_fccdfs":true,"radio_caps":251805700,"radio_caps2":31,` +
	`"antenna_id":-1,"current_antenna_gain":0}`

// TestDeviceRadioTable5GFieldsRoundTrip locks the fix for
// `400 api.err.DeviceNotSupport5gException`. Before the fix, DeviceRadioTable did
// not model the UAP6MP's 5 GHz capability fields, so json.Unmarshal DROPPED them;
// a subsequent radio_table PUT re-sent the `na` radio stripped of nss/radio_caps/
// has_ht160/etc. and the controller rejected it. This test proves those fields
// now survive an unmarshal -> marshal round-trip.
func TestDeviceRadioTable5GFieldsRoundTrip(t *testing.T) {
	t.Parallel()

	var radio DeviceRadioTable
	require.NoError(t, json.Unmarshal([]byte(realUAP6MP5GHzRadio), &radio))

	// The capability fields the controller requires must be populated.
	assert.Equal(t, 4, radio.Nss, "nss must be preserved")
	assert.Equal(t, 251805700, radio.RadioCaps, "radio_caps must be preserved")
	assert.Equal(t, 31, radio.RadioCaps2, "radio_caps2 must be preserved")
	assert.True(t, radio.HasHt160, "has_ht160 must be preserved")
	assert.True(t, radio.Is11Ac, "is_11ac must be preserved")
	assert.True(t, radio.Is11Ax, "is_11ax must be preserved")
	assert.Equal(t, 26, radio.MaxTxpower, "max_txpower must be preserved")
	assert.Equal(t, 6, radio.MinTxpower, "min_txpower must be preserved")

	// And they must be re-emitted on marshal (not dropped) so a PUT keeps them.
	out, err := json.Marshal(&radio)
	require.NoError(t, err)

	var m map[string]any
	require.NoError(t, json.Unmarshal(out, &m))
	for _, k := range []string{"nss", "radio_caps", "radio_caps2", "has_ht160", "max_txpower", "min_txpower", "is_11ac", "is_11ax"} {
		_, ok := m[k]
		assert.Truef(t, ok, "field %q must be present in the marshaled radio (not stripped)", k)
	}
}

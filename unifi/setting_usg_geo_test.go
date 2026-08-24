package unifi //nolint: testpackage

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetSettingUsgGeo verifies that Region Blocking is read from the separate
// `usg_geo` setting key (nested `ip_filtering`), which is where UniFi Network
// 10.x stores it — not from flat fields on the `usg` setting.
func TestGetSettingUsgGeo(t *testing.T) {
	t.Parallel()

	const site = "default"
	path := apiV1Path("s/" + site + "/get/setting")

	// Real shape observed on a 10.5.67 controller: usg_geo alongside usg, with a
	// nested ip_filtering object; the usg object carries no geo fields.
	response := `{"meta":{"rc":"ok"},"data":[` +
		`{"_id":"a","key":"usg"},` +
		`{"_id":"b","key":"usg_geo","ip_filtering":{"action":"block","countries":"RU","enabled":true,"traffic_direction":"both"}}` +
		`]}`

	cs := newControllerServer(t, route{path, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(response))
	}})
	c := cs.client()

	got, err := c.GetSettingUsgGeo(context.Background(), site)
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, http.MethodGet, cs.lastRequest().Method)
	assert.Equal(t, path, cs.lastRequest().Path)

	assert.Equal(t, SettingUsgGeoKey, got.Key)
	assert.Equal(t, "block", got.IPFiltering.Action)
	assert.Equal(t, "RU", got.IPFiltering.Countries)
	assert.True(t, got.IPFiltering.Enabled)
	assert.Equal(t, "both", got.IPFiltering.TrafficDirection)
}

// TestUpdateSettingUsgGeo verifies the wrapper PUTs to set/setting/usg_geo with
// the nested ip_filtering payload and decodes the echoed result.
func TestUpdateSettingUsgGeo(t *testing.T) {
	t.Parallel()

	const site = "default"
	path := apiV1Path("s/" + site + "/set/setting/" + SettingUsgGeoKey)

	response := `{"meta":{"rc":"ok"},"data":[` +
		`{"_id":"b","key":"usg_geo","ip_filtering":{"action":"block","countries":"RU","enabled":true,"traffic_direction":"both"}}` +
		`]}`

	cs := newControllerServer(t, route{path, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(response))
	}})
	c := cs.client()

	got, err := c.UpdateSettingUsgGeo(context.Background(), site, &SettingUsgGeo{
		IPFiltering: SettingUsgGeoIPFiltering{
			Action:           "block",
			Countries:        "RU",
			Enabled:          true,
			TrafficDirection: "both",
		},
	})
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, http.MethodPut, cs.lastRequest().Method)
	assert.Equal(t, path, cs.lastRequest().Path)

	assert.Equal(t, SettingUsgGeoKey, got.Key)
	assert.True(t, got.IPFiltering.Enabled)
	assert.Equal(t, "RU", got.IPFiltering.Countries)
}

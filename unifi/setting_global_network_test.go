package unifi //nolint: testpackage

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettingGlobalNetworkRoundTrip(t *testing.T) {
	t.Parallel()

	const site = "default"
	getPath := apiV1Path("s/" + site + "/get/setting")
	setPath := apiV1Path("s/" + site + "/set/setting/" + SettingGlobalNetworkKey)
	body := `{"meta":{"rc":"ok"},"data":[{"_id":"g","key":"global_network","default_security_posture":"ALLOW_ALL"}]}`

	cs := newControllerServer(t,
		route{getPath, func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte(body)) }},
		route{setPath, func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte(body)) }},
	)
	c := cs.client()

	got, err := c.GetSettingGlobalNetwork(context.Background(), site)
	require.NoError(t, err)
	assert.Equal(t, "ALLOW_ALL", got.DefaultSecurityPosture)

	_, err = c.UpdateSettingGlobalNetwork(context.Background(), site, got)
	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, cs.lastRequest().Method)
	assert.Equal(t, setPath, cs.lastRequest().Path)
	assert.Contains(t, string(cs.lastRequest().Body), `"key":"global_network"`)
}

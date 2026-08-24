package unifi //nolint: testpackage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNetworkWireguardServerFields verifies the wireguard-server (remote-user-vpn)
// fields added for G4 unmarshal and round-trip on a networkconf.
func TestNetworkWireguardServerFields(t *testing.T) {
	t.Parallel()

	const body = `{"_id":"n1","name":"WireGuard Server 1","purpose":"remote-user-vpn",` +
		`"vpn_type":"wireguard-server","wireguard_id":1,` +
		`"wireguard_interface_binding_mode_ip_version":"v4","vpn_binding_mode":"interface",` +
		`"mss_clamp":"auto","interface_mtu_enabled":false,"local_port":51820}`

	var n Network
	require.NoError(t, json.Unmarshal([]byte(body), &n))
	assert.Equal(t, "remote-user-vpn", n.Purpose)
	assert.Equal(t, "wireguard-server", n.VPNType)
	assert.Equal(t, 1, n.WireguardID)
	assert.Equal(t, "v4", n.WireguardInterfaceBindingModeIPVersion)
	assert.Equal(t, "interface", n.VPNBindingMode)
	assert.Equal(t, "auto", n.MssClamp)

	b, err := json.Marshal(n)
	require.NoError(t, err)
	var back Network
	require.NoError(t, json.Unmarshal(b, &back))
	assert.Equal(t, n.WireguardID, back.WireguardID)
	assert.Equal(t, n.VPNBindingMode, back.VPNBindingMode)
	assert.Equal(t, n.MssClamp, back.MssClamp)
}

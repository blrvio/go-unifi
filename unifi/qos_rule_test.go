package unifi //nolint: testpackage

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The Critical Apps Prioritization rule as returned by the 10.5.67 controller.
const qosRuleListBody = `[{"_id":"q1","name":"Critical Apps Prioritization","enabled":true,"index":10001,` +
	`"objective":"PRIORITIZE","download_burst":"OFF","upload_burst":"OFF","wan_or_vpn_network":"net1",` +
	`"schedule":{"mode":"ALWAYS"},"source":{"matching_target":"ANY"},` +
	`"destination":{"matching_target":"APP","port_matching_type":"ANY","app_ids":[393220,1114124]}}]`

// TestGetQOSRuleUsesListRoute verifies GetQOSRule reads via the qos-rules LIST
// collection and filters client-side — the v2 collection has no GET-by-id.
func TestGetQOSRuleUsesListRoute(t *testing.T) {
	t.Parallel()

	const site = "default"
	listPath := apiV2("site/" + site + "/qos-rules")

	cs := newControllerServer(t,
		route{listPath, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			_, _ = w.Write([]byte(qosRuleListBody))
		}},
		route{listPath + "/q1", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}},
	)
	c := cs.client()

	got, err := c.GetQOSRule(context.Background(), site, "q1")
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, listPath, cs.lastRequest().Path, "GetQOSRule must read via LIST, not GET-by-id")
	assert.Equal(t, "q1", got.ID)
	assert.Equal(t, "Critical Apps Prioritization", got.Name)
	assert.Equal(t, 10001, got.Index)
	assert.Equal(t, "PRIORITIZE", got.Objective)
	assert.Equal(t, "APP", got.Destination.MatchingTarget)
	assert.Equal(t, []int{393220, 1114124}, got.Destination.AppIDs)
	assert.Equal(t, "ALWAYS", got.Schedule.Mode)

	_, err = c.GetQOSRule(context.Background(), site, "nope")
	assert.ErrorIs(t, err, ErrNotFound)
}

// TestQOSRuleRoundTrip ensures the model (un)marshals the nested objects without
// loss — app_ids stays []int, index stays int, bursts stay strings.
func TestQOSRuleRoundTrip(t *testing.T) {
	t.Parallel()

	var rules []QOSRule
	require.NoError(t, json.Unmarshal([]byte(qosRuleListBody), &rules))
	require.Len(t, rules, 1)

	b, err := json.Marshal(rules[0])
	require.NoError(t, err)

	var back QOSRule
	require.NoError(t, json.Unmarshal(b, &back))
	assert.Equal(t, rules[0], back)
	assert.Equal(t, []int{393220, 1114124}, back.Destination.AppIDs)
}

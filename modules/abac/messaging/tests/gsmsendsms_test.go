// GsmSendSmsWithProviderAction (/gsmProvider/send/sms) black-box tests, following
// webpushconfig_test.go's exact conventions. Reuses createSampleGsmProvider/
// deleteGsmProvider from gsmprovider_test.go and postJSON from sendemail_test.go (same
// package) - needs a real, "terminal"-type GsmProviderEntity to resolve against
// (GsmSendSMSByTerminal just logs and returns "", so no real gateway is ever needed).
//
// Unlike SendEmailAction/SendEmailWithProviderAction, this action already validated and
// used its request fields correctly before this change - these tests are purely new
// coverage (tools/checkendpointtests), not regression tests for a fix.
package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	abactests "github.com/torabian/fireback/modules/abac/tests"
)

func TestGsmSendSmsWithProvider_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	cases := []struct {
		name    string
		payload map[string]any
	}{
		{"missing toNumber", map[string]any{"body": "hello"}},
		{"missing body", map[string]any{"toNumber": "+10000000000"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := postJSON(t, cfg, "/gsmProvider/send/sms", tc.payload, true)
			if resp.StatusCode == http.StatusOK {
				t.Fatalf("expected %s to be rejected, got 200 OK: %s", tc.name, body)
			}
		})
	}
}

func TestGsmSendSmsWithProvider_HTTP_RequiresAuth(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)

	resp, body := postJSON(t, cfg, "/gsmProvider/send/sms", map[string]any{
		"toNumber": "+10000000000",
		"body":     "hello",
	}, false)

	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unauthenticated send to be rejected, got 200 OK: %s", body)
	}
}

func TestGsmSendSmsWithProvider_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	provider := createSampleGsmProvider(t, cfg)
	defer deleteGsmProvider(t, cfg, provider.UniqueId)

	resp, body := postJSON(t, cfg, "/gsmProvider/send/sms", map[string]any{
		"gsmProviderId": provider.UniqueId,
		"toNumber":      "+10000000000",
		"body":          "hello from checkendpointtests",
	}, true)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK sending sms, got %d: %s", resp.StatusCode, body)
	}

	// GsmSendSMSByTerminal (the "terminal" provider type's send path) returns an empty
	// queueId on success - just decode to confirm the response is at least well-formed
	// JSON matching the declared {"queueId": string} shape.
	var out sendEmailRes
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode send response: %v\nbody: %s", err, body)
	}
}

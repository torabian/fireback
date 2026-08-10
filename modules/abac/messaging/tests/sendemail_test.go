// SendEmailAction (/email/send) and SendEmailWithProviderAction (/emailProvider/send)
// black-box tests, following webpushconfig_test.go's exact conventions. Reuses
// createSampleEmailProvider/deleteEmailProvider from emailprovider_test.go (same
// package) - both hand actions need a real, "terminal"-type EmailProviderEntity to
// resolve against (SendMail's "terminal" case just logs and returns nil, so no real SMTP
// is ever needed).
package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	abactests "github.com/torabian/fireback/modules/abac/tests"
)

// sendEmailRes mirrors the bare (non-GResponse-wrapped) {"queueId": "..."} payload both
// SendEmailAction and SendEmailWithProviderAction respond with - see
// SendEmailActionImplementation.go/SendEmailWithProviderActionImplementation.go, which
// set resp.Payload directly to a *SendEmail(WithProvider)ActionRes rather than wrapping
// it via fireback.GResponseSingleItem.
type sendEmailRes struct {
	QueueId string `json:"queueId"`
}

func postJSON(t *testing.T, cfg abactests.TestConfig, url string, payload map[string]any, authed bool) (*http.Response, []byte) {
	t.Helper()

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL(url), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if authed {
		req.Header.Set("Authorization", cfg.CliToken)
		req.Header.Set("Workspace-id", cfg.WorkspaceID)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request to %s failed: %v", req.URL, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	return resp, respBody
}

// --- SendEmailAction (/email/send) ---

func TestSendEmail_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	cases := []struct {
		name    string
		payload map[string]any
	}{
		{"missing toAddress", map[string]any{"body": "hello"}},
		{"missing body", map[string]any{"toAddress": "invitee@example.com"}},
		{"empty request", map[string]any{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := postJSON(t, cfg, "/email/send", tc.payload, true)
			if resp.StatusCode == http.StatusOK {
				t.Fatalf("expected %s to be rejected, got 200 OK: %s", tc.name, body)
			}
		})
	}
}

// TestSendEmail_HTTP_Succeeds covers the fix making SendEmailAction actually validate
// and use req.ToAddress/req.Body (it used to hardcode a "test@test.com"/"Hello :)"
// message regardless of what was posted, and never returned a queueId).
func TestSendEmail_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	provider := createSampleEmailProvider(t, cfg)
	defer deleteEmailProvider(t, cfg, provider.UniqueId)

	resp, body := postJSON(t, cfg, "/email/send", map[string]any{
		"providerId": provider.UniqueId,
		"toAddress":  "invitee@example.com",
		"body":       "hello from checkendpointtests",
	}, true)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK sending email, got %d: %s", resp.StatusCode, body)
	}

	var out sendEmailRes
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode send response: %v\nbody: %s", err, body)
	}
	if out.QueueId == "" {
		t.Errorf("expected a non-empty queueId, got none: %s", body)
	}
}

// --- SendEmailWithProviderAction (/emailProvider/send) ---

func TestSendEmailWithProvider_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	cases := []struct {
		name    string
		payload map[string]any
	}{
		{"missing toAddress", map[string]any{"body": "hello"}},
		{"missing body", map[string]any{"toAddress": "invitee@example.com"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := postJSON(t, cfg, "/emailProvider/send", tc.payload, true)
			if resp.StatusCode == http.StatusOK {
				t.Fatalf("expected %s to be rejected, got 200 OK: %s", tc.name, body)
			}
		})
	}
}

// TestSendEmailWithProvider_HTTP_RejectsUnknownProvider covers the fix implementing
// SendEmailWithProviderAction for real - it used to be a pure no-op that always
// returned 200 OK regardless of the (previously never resolved) emailProvider.
func TestSendEmailWithProvider_HTTP_RejectsUnknownProvider(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postJSON(t, cfg, "/emailProvider/send", map[string]any{
		"emailProvider": map[string]any{"uniqueId": "does-not-exist"},
		"toAddress":     "invitee@example.com",
		"body":          "hello",
	}, true)

	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected sending through an unknown emailProvider to be rejected, got 200 OK: %s", body)
	}
}

func TestSendEmailWithProvider_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	provider := createSampleEmailProvider(t, cfg)
	defer deleteEmailProvider(t, cfg, provider.UniqueId)

	resp, body := postJSON(t, cfg, "/emailProvider/send", map[string]any{
		"emailProvider": map[string]any{"uniqueId": provider.UniqueId},
		"toAddress":     "invitee@example.com",
		"body":          "hello from checkendpointtests",
	}, true)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK sending email with provider, got %d: %s", resp.StatusCode, body)
	}

	var out sendEmailRes
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode send response: %v\nbody: %s", err, body)
	}
	if out.QueueId == "" {
		t.Errorf("expected a non-empty queueId, got none: %s", body)
	}
}

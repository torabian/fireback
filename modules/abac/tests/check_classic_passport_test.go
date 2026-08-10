// CheckClassicPassportAction (/workspace/passport/check) black-box tests - the
// "does this email/phone already have an account" step the /welcome screen calls
// before showing either a password-entry or account-creation form. Follows this
// package's established conventions; reuses freshSignupTarget (core_session_test.go)
// and googleResponseEnvelope (passport_methods_http_test.go).
package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

type checkClassicPassportRes struct {
	Next  []string `json:"next"`
	Flags []string `json:"flags"`
}

func postCheckClassicPassport(t *testing.T, cfg TestConfig, value string) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(map[string]any{"value": value})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/passport/check"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build check request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("check request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

// TestCheckClassicPassport_HTTP_NonExistingAccount covers checkStepsForNonExistingAccount
// - with no WorkspaceConfig (this test DB's default state, no otp enabled/forced), the
// only next step for an email nobody has signed up with yet is "create-with-password".
func TestCheckClassicPassport_HTTP_NonExistingAccount(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	value := fmt.Sprintf("checkendpointtests-ccp-new-%d@example.com", time.Now().UnixNano())
	resp, body := postCheckClassicPassport(t, cfg, value)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK checking a non-existing passport, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[checkClassicPassportRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode check response: %v\nbody: %s", err, body)
	}
	found := false
	for _, n := range out.Data.Item.Next {
		if n == "create-with-password" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected next steps to include %q for a non-existing account, got %+v", "create-with-password", out.Data.Item.Next)
	}
}

// TestCheckClassicPassport_HTTP_ExistingAccountWithPassword covers
// checkStepsForExistingAccount - an email that already has a password-based passport
// (from freshSignupTarget's signup) should be offered "signin-with-password".
func TestCheckClassicPassport_HTTP_ExistingAccountWithPassword(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	email, _, _ := freshSignupTarget(t, cfg, "ccp-existing")

	resp, body := postCheckClassicPassport(t, cfg, email)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK checking an existing passport, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[checkClassicPassportRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode check response: %v\nbody: %s", err, body)
	}
	found := false
	for _, n := range out.Data.Item.Next {
		if n == "signin-with-password" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected next steps to include %q for an existing account with a password, got %+v", "signin-with-password", out.Data.Item.Next)
	}
}

// TestCheckClassicPassport_HTTP_RejectsInvalidFormat covers validateValueFormat - a
// value that's neither a valid email nor a valid phone number (and not the special
// "anonymous_" prefix) should be rejected outright, before ever touching the database.
func TestCheckClassicPassport_HTTP_RejectsInvalidFormat(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	resp, body := postCheckClassicPassport(t, cfg, "not-a-valid-email-or-phone")
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an invalid-format value to be rejected, got 200 OK: %s", body)
	}
}

// Core session lifecycle black-box tests: Whoami, ClassicSignin, ClassicSignup,
// Signout, ChangePassword. Follows this package's established conventions (see
// testconfig.go's doc comment) and reuses signupClassic/createRoleAndWorkspaceType/
// ensureEmailPassportMethod/googleResponseEnvelope/whoamiResult
// (workspace_invite_accept_test.go) - same package, no need to redeclare them.
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

// freshSignupTarget creates a throwaway role+workspaceType (no capabilities needed -
// these tests only exercise session actions, not anything capability-gated) and signs a
// brand new account up against it, returning its email/token/workspaceId.
func freshSignupTarget(t *testing.T, cfg TestConfig, label string) (email, token, workspaceId string) {
	t.Helper()
	ensureEmailPassportMethod(t, cfg)
	// A role needs at least one capability (see RoleCreate's RoleNeedsOneCapability
	// check) - this one is inert for what these tests actually exercise (session
	// lifecycle, not access control), same harmless pick workspace-invite.cy.ts uses
	// for its invitee-own role.
	workspaceTypeId := createRoleAndWorkspaceType(t, cfg, fmt.Sprintf("checkendpointtests %s role", label), []string{"root.abac.email-confirmation.query"})
	email = fmt.Sprintf("checkendpointtests-%s-%d@example.com", label, time.Now().UnixNano())
	token, workspaceId = signupClassic(t, cfg, email, workspaceTypeId)
	return email, token, workspaceId
}

func TestWhoami_HTTP_ReturnsOwnWorkspaces(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	_, token, workspaceId := freshSignupTarget(t, cfg, "whoami")

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/whoami"), nil)
	if err != nil {
		t.Fatalf("failed to build whoami request: %v", err)
	}
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("whoami request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK from whoami, got %d: %s", resp.StatusCode, body)
	}

	var out whoamiResult
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode whoami response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UserId == "" {
		t.Errorf("expected a non-empty userId, got none: %+v", out.Data.Item)
	}
	found := false
	for _, w := range out.Data.Item.Workspaces {
		if w.UniqueId == workspaceId {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected whoami's workspaces to include %s, got %+v", workspaceId, out.Data.Item.Workspaces)
	}
}

func TestWhoami_HTTP_RequiresAuth(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/whoami"), nil)
	if err != nil {
		t.Fatalf("failed to build whoami request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("whoami request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unauthenticated whoami to be rejected, got 200 OK")
	}
}

func postSignin(t *testing.T, cfg TestConfig, value, password string) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(map[string]any{"value": value, "password": password})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passports/signin/classic"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build signin request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("signin request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

func TestClassicSignin_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	email, _, _ := freshSignupTarget(t, cfg, "signin")

	resp, body := postSignin(t, cfg, email, "checkendpointtests-pass-123")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK signing in with correct credentials, got %d: %s", resp.StatusCode, body)
	}
	var out signupResult
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode signin response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.Session.Token == "" {
		t.Errorf("expected a session token, got none: %s", body)
	}
}

func TestClassicSignin_HTTP_RejectsWrongPassword(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	email, _, _ := freshSignupTarget(t, cfg, "signin-wrongpass")

	resp, body := postSignin(t, cfg, email, "not-the-right-password")
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected signin with a wrong password to be rejected, got 200 OK: %s", body)
	}
}

func TestClassicSignin_HTTP_RejectsUnknownValue(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	resp, body := postSignin(t, cfg, fmt.Sprintf("checkendpointtests-nobody-%d@example.com", time.Now().UnixNano()), "whatever-password")
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected signin for an unknown passport to be rejected, got 200 OK: %s", body)
	}
}

func postSignup(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passports/signup/classic"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build signup request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("signup request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

func TestClassicSignup_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	_, token, workspaceId := freshSignupTarget(t, cfg, "signup")
	if token == "" || workspaceId == "" {
		t.Errorf("expected a session token and workspaceId from signup, got token=%q workspaceId=%q", token, workspaceId)
	}
}

func TestClassicSignup_HTTP_RequiresWorkspaceTypeId(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	resp, body := postSignup(t, cfg, map[string]any{
		"value":     fmt.Sprintf("checkendpointtests-notype-%d@example.com", time.Now().UnixNano()),
		"type":      "email",
		"password":  "checkendpointtests-pass-123",
		"firstName": "Checkendpointtests",
		"lastName":  "NoType",
	})
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected signup with no workspaceTypeId to be rejected, got 200 OK: %s", body)
	}
}

func TestClassicSignup_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	workspaceTypeId := createRoleAndWorkspaceType(t, cfg, "checkendpointtests signup validation role", []string{"root.abac.email-confirmation.query"})

	for _, tc := range []struct {
		name    string
		payload map[string]any
	}{
		{"missing firstName", map[string]any{"value": fmt.Sprintf("checkendpointtests-nofn-%d@example.com", time.Now().UnixNano()), "type": "email", "password": "checkendpointtests-pass-123", "lastName": "X", "workspaceTypeId": workspaceTypeId}},
		{"missing lastName", map[string]any{"value": fmt.Sprintf("checkendpointtests-noln-%d@example.com", time.Now().UnixNano()), "type": "email", "password": "checkendpointtests-pass-123", "firstName": "X", "workspaceTypeId": workspaceTypeId}},
		{"missing password", map[string]any{"value": fmt.Sprintf("checkendpointtests-nopw-%d@example.com", time.Now().UnixNano()), "type": "email", "firstName": "X", "lastName": "Y", "workspaceTypeId": workspaceTypeId}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := postSignup(t, cfg, tc.payload)
			if resp.StatusCode == http.StatusOK {
				t.Fatalf("expected signup with %s to be rejected, got 200 OK: %s", tc.name, body)
			}
		})
	}
}

func TestSignout_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	_, token, _ := freshSignupTarget(t, cfg, "signout")

	client := cfg.NewHTTPClient()
	// The generated Gin handler always calls ShouldBindJSON, even though SignoutActionReq's
	// only field (Clear) is optional - an empty/absent body fails with "invalid JSON: EOF",
	// so an explicit "{}" is required here.
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/signout"), bytes.NewReader([]byte("{}")))
	if err != nil {
		t.Fatalf("failed to build signout request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("signout request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK signing out, got %d: %s", resp.StatusCode, body)
	}

	// Unlike almost every other action in this package, SignoutActionImplementation.go
	// returns its SignoutActionRes payload bare - not wrapped in
	// fireback.GResponseSingleItem - so the response is flat {"okay":true}, not the
	// usual {"data":{"item":{...}}} envelope.
	var out struct {
		Okay bool `json:"okay"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode signout response: %v\nbody: %s", err, body)
	}
	if !out.Okay {
		t.Errorf("expected okay:true, got %+v", out)
	}
}

func TestChangePassword_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	email, token, _ := freshSignupTarget(t, cfg, "changepass")

	// uniqueId is validate:"required" on ChangePasswordActionReq, but
	// ChangePasswordActionImplementation.go never actually reads it - it always
	// resolves the target passport itself from the authenticated session's own userId
	// (passports[0], from a query scoped to q.UserId). Any non-empty placeholder
	// satisfies validation without affecting which passport gets updated - itself a
	// safer design than trusting a client-supplied id would be, just worth a comment
	// since it makes the field's required-ness a little surprising from outside.
	body, _ := json.Marshal(map[string]any{"password": "new-checkendpointtests-pass-456", "uniqueId": "placeholder-unused-by-implementation"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/change-password"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build change-password request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("change-password request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK changing password, got %d: %s", resp.StatusCode, respBody)
	}

	// The old password must no longer work, and the new one must.
	oldResp, oldBody := postSignin(t, cfg, email, "checkendpointtests-pass-123")
	if oldResp.StatusCode == http.StatusOK {
		t.Errorf("expected signin with the old password to be rejected after change, got 200 OK: %s", oldBody)
	}
	newResp, newBody := postSignin(t, cfg, email, "new-checkendpointtests-pass-456")
	if newResp.StatusCode != http.StatusOK {
		t.Errorf("expected signin with the new password to succeed, got %d: %s", newResp.StatusCode, newBody)
	}
}

func TestChangePassword_HTTP_RejectsShortPassword(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	_, token, _ := freshSignupTarget(t, cfg, "changepass-short")

	// uniqueId included (see TestChangePassword_HTTP_Succeeds's comment) so this
	// actually proves rejection is because of the too-short password, not an
	// incidental missing-required-field error on uniqueId.
	body, _ := json.Marshal(map[string]any{"password": "abc", "uniqueId": "placeholder-unused-by-implementation"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/change-password"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build change-password request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("change-password request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected a too-short password to be rejected, got 200 OK: %s", respBody)
	}
}

func TestChangePassword_HTTP_RequiresAuth(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	body, _ := json.Marshal(map[string]any{"password": "some-new-password-123"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/change-password"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build change-password request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("change-password request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unauthenticated change-password to be rejected, got 200 OK")
	}
}

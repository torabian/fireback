// Black-box tests for CreatePassportForUserAction
// (modules/abac/CreatePassportForUserActionImplementation.go) - lets root manually
// create a new email/phone passport for an existing user with a password already set,
// for provisioning a working login when the user can't complete signup/confirmation
// themselves (e.g. no mail server configured). Root only (AllowOnRoot).
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

type createPassportForUserRes struct {
	UniqueId string `json:"uniqueId"`
	UserId   string `json:"userId"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type classicSigninResultRes struct {
	Data struct {
		Item struct {
			Session struct {
				Token string `json:"token"`
			} `json:"session"`
		} `json:"item"`
	} `json:"data"`
}

// signinClassic hits POST /passports/signin/classic (no auth header needed - it's how a
// browser's own signin form would call it) and returns the raw response, so callers
// decide what status they expect.
func signinClassic(t *testing.T, cfg TestConfig, value, password string) (*http.Response, []byte) {
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

// TestCreatePassportForUser_HTTP_Succeeds covers the actual point of this action: the
// created passport isn't just a row in the database - it has to really work as a login,
// with the exact plaintext password that was given (hashed server-side, never round-
// tripped), against the real /passports/signin/classic endpoint.
func TestCreatePassportForUser_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	value := fmt.Sprintf("checkendpointtests-create-passport-%d@example.com", time.Now().UnixNano())
	password := "checkendpointtests-secret-1"

	resp, body := postToRootAction(t, cfg, "/passport/create-for-user", "root", map[string]any{
		"userId":   user.UniqueId,
		"type":     "email",
		"value":    value,
		"password": password,
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating passport for user, got %d: %s", resp.StatusCode, body)
	}
	var out struct {
		Data struct {
			Item createPassportForUserRes `json:"item"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", out.Data.Item)
	}
	defer deletePassport(t, cfg, out.Data.Item.UniqueId)

	if out.Data.Item.Type != "email" {
		t.Errorf("expected type %q, got %q", "email", out.Data.Item.Type)
	}
	if out.Data.Item.Value != value {
		t.Errorf("expected value %q, got %q", value, out.Data.Item.Value)
	}
	if out.Data.Item.UserId != user.UniqueId {
		t.Errorf("expected userId %q, got %q", user.UniqueId, out.Data.Item.UserId)
	}

	// The real point: sign in with exactly the password that was set, no session/
	// token/other admin action involved.
	signinResp, signinRespBody := signinClassic(t, cfg, value, password)
	if signinResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK signing in with the newly created passport, got %d: %s", signinResp.StatusCode, signinRespBody)
	}
	var signinOut classicSigninResultRes
	if err := json.Unmarshal(signinRespBody, &signinOut); err != nil {
		t.Fatalf("failed to decode signin response: %v\nbody: %s", err, signinRespBody)
	}
	if signinOut.Data.Item.Session.Token == "" {
		t.Errorf("expected a session token after signing in with the created passport, got none: %s", signinRespBody)
	}
}

// TestCreatePassportForUser_HTTP_RejectsInvalidType covers type validation - only
// PASSPORT_METHOD_EMAIL ("email") and PASSPORT_METHOD_PHONE ("phone") are accepted.
func TestCreatePassportForUser_HTTP_RejectsInvalidType(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	resp, body := postToRootAction(t, cfg, "/passport/create-for-user", "root", map[string]any{
		"userId":   user.UniqueId,
		"type":     "checkendpointtests-bogus-type",
		"value":    fmt.Sprintf("checkendpointtests-invalid-type-%d@example.com", time.Now().UnixNano()),
		"password": "checkendpointtests-secret-1",
	})
	if resp.StatusCode == http.StatusOK {
		var out struct {
			Data struct {
				Item createPassportForUserRes `json:"item"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &out); err == nil {
			deletePassport(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected an invalid type to be rejected, got 200 OK: %s", body)
	}
}

// TestCreatePassportForUser_HTTP_RejectsShortPassword covers the same minimum-length rule
// SetPassportPasswordAction/ChangePasswordAction enforce.
func TestCreatePassportForUser_HTTP_RejectsShortPassword(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	resp, body := postToRootAction(t, cfg, "/passport/create-for-user", "root", map[string]any{
		"userId":   user.UniqueId,
		"type":     "email",
		"value":    fmt.Sprintf("checkendpointtests-short-pw-%d@example.com", time.Now().UnixNano()),
		"password": "123",
	})
	if resp.StatusCode == http.StatusOK {
		var out struct {
			Data struct {
				Item createPassportForUserRes `json:"item"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &out); err == nil {
			deletePassport(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected a too-short password to be rejected, got 200 OK: %s", body)
	}
}

// TestCreatePassportForUser_HTTP_RejectsUnknownUser covers userId validation - it must be
// an existing user.
func TestCreatePassportForUser_HTTP_RejectsUnknownUser(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postToRootAction(t, cfg, "/passport/create-for-user", "root", map[string]any{
		"userId":   "checkendpointtests-does-not-exist",
		"type":     "email",
		"value":    fmt.Sprintf("checkendpointtests-unknown-user-%d@example.com", time.Now().UnixNano()),
		"password": "checkendpointtests-secret-1",
	})
	if resp.StatusCode == http.StatusOK {
		var out struct {
			Data struct {
				Item createPassportForUserRes `json:"item"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &out); err == nil {
			deletePassport(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected an unknown userId to be rejected, got 200 OK: %s", body)
	}
}

// TestCreatePassportForUser_HTTP_RejectsDuplicateValue covers PassportEntity.Value's
// db-level unique constraint - two passports (even for two different users) can't share
// the same email/phone value.
func TestCreatePassportForUser_HTTP_RejectsDuplicateValue(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	user1 := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user1.UniqueId)
	user2 := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user2.UniqueId)

	value := fmt.Sprintf("checkendpointtests-duplicate-%d@example.com", time.Now().UnixNano())

	resp1, body1 := postToRootAction(t, cfg, "/passport/create-for-user", "root", map[string]any{
		"userId": user1.UniqueId, "type": "email", "value": value, "password": "checkendpointtests-secret-1",
	})
	if resp1.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating the first passport, got %d: %s", resp1.StatusCode, body1)
	}
	var out1 struct {
		Data struct {
			Item createPassportForUserRes `json:"item"`
		} `json:"data"`
	}
	_ = json.Unmarshal(body1, &out1)
	defer deletePassport(t, cfg, out1.Data.Item.UniqueId)

	resp2, body2 := postToRootAction(t, cfg, "/passport/create-for-user", "root", map[string]any{
		"userId": user2.UniqueId, "type": "email", "value": value, "password": "checkendpointtests-secret-2",
	})
	if resp2.StatusCode == http.StatusOK {
		var out2 struct {
			Data struct {
				Item createPassportForUserRes `json:"item"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body2, &out2); err == nil {
			deletePassport(t, cfg, out2.Data.Item.UniqueId)
		}
		t.Fatalf("expected a duplicate passport value to be rejected, got 200 OK: %s", body2)
	}
}

// TestCreatePassportForUser_HTTP_RootOnly covers AllowOnRoot.
func TestCreatePassportForUser_HTTP_RootOnly(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws, wt, backingRole := createSampleWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, backingRole.UniqueId)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	resp, body := postToRootAction(t, cfg, "/passport/create-for-user", ws.UniqueId, map[string]any{
		"userId":   user.UniqueId,
		"type":     "email",
		"value":    fmt.Sprintf("checkendpointtests-root-only-%d@example.com", time.Now().UnixNano()),
		"password": "checkendpointtests-secret-1",
	})
	if resp.StatusCode == http.StatusOK {
		var out struct {
			Data struct {
				Item createPassportForUserRes `json:"item"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &out); err == nil {
			deletePassport(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected create-for-user with a non-root Workspace-Id header to be rejected, got 200 OK: %s", body)
	}
}

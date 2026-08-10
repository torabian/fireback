// Alternate auth method black-box tests: ClassicPassportRequestOtp, ClassicPassportOtp,
// ConfirmClassicPassportTotp, OauthAuthenticate, OsLoginAuthenticate. Follows this
// package's established conventions; reuses freshSignupTarget (core_session_test.go),
// ensureInviteEmailSendingConfigured/signupClassic/googleResponseEnvelope/
// googleResponseListEnvelope (workspace_invite_accept_test.go).
package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
)

// requestOtpAndFetchCode calls /workspace/passport/request-otp for value, then, as root,
// browses /publicAuthentication (see PublicAuthenticationActions.go - root's own CLI
// token has the wildcard "root.*", which covers this "root.manage" scoped entity same as
// everything else) to read back the plaintext Otp code that was "sent" - there's no
// other observable way to retrieve it in a black-box HTTP test, since email delivery
// here goes through a "terminal" provider that only logs.
func requestOtpAndFetchCode(t *testing.T, cfg TestConfig, value string) string {
	t.Helper()

	body, _ := json.Marshal(map[string]any{"value": value})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/passport/request-otp"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request-otp request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request-otp request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK requesting otp for %q, got %d: %s", value, resp.StatusCode, respBody)
	}

	browseReq, err := http.NewRequest(http.MethodGet, cfg.URL("/publicAuthentication/browse?itemsPerPage=1000"), nil)
	if err != nil {
		t.Fatalf("failed to build publicAuthentication browse request: %v", err)
	}
	browseReq.Header.Set("Authorization", cfg.CliToken)
	browseReq.Header.Set("Workspace-id", cfg.WorkspaceID)
	browseResp, err := client.Do(browseReq)
	if err != nil {
		t.Fatalf("publicAuthentication browse request failed: %v", err)
	}
	defer browseResp.Body.Close()
	browseBody, _ := io.ReadAll(browseResp.Body)
	if browseResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK browsing publicAuthentication, got %d: %s", browseResp.StatusCode, browseBody)
	}

	var list googleResponseListEnvelope[struct {
		PassportValue string `json:"passportValue"`
		Otp           string `json:"otp"`
	}]
	if err := json.Unmarshal(browseBody, &list); err != nil {
		t.Fatalf("failed to decode publicAuthentication browse response: %v\nbody: %s", err, browseBody)
	}

	// Most recent (highest id, so last in insertion order) matching row - a value can
	// accumulate more than one PublicAuthentication row across a test run.
	code := ""
	for _, item := range list.Data.Items {
		if item.PassportValue == value {
			code = item.Otp
		}
	}
	if code == "" {
		t.Fatalf("no publicAuthentication row found for passportValue %q after requesting an otp: %s", value, browseBody)
	}
	return code
}

func TestClassicPassportRequestOtp_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ensureInviteEmailSendingConfigured(t, cfg)
	email, _, _ := freshSignupTarget(t, cfg, "otp-request")

	body, _ := json.Marshal(map[string]any{"value": email})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/passport/request-otp"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request-otp request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request-otp request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d: %s", resp.StatusCode, respBody)
	}

	var out googleResponseEnvelope[struct {
		SecondsToUnblock int64 `json:"secondsToUnblock"`
	}]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.SecondsToUnblock <= 0 {
		t.Errorf("expected a positive secondsToUnblock, got %d", out.Data.Item.SecondsToUnblock)
	}
}

// TestClassicPassportRequestOtp_HTTP_BlocksRepeatedRequests covers the
// BlockedUntil/OtaRequestBlockedUntil throttling - a second request for the same value
// before the first's block window elapses must be rejected.
func TestClassicPassportRequestOtp_HTTP_BlocksRepeatedRequests(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ensureInviteEmailSendingConfigured(t, cfg)
	email, _, _ := freshSignupTarget(t, cfg, "otp-throttle")

	postOnce := func() (*http.Response, []byte) {
		body, _ := json.Marshal(map[string]any{"value": email})
		client := cfg.NewHTTPClient()
		req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/passport/request-otp"), bytes.NewReader(body))
		if err != nil {
			t.Fatalf("failed to build request-otp request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("request-otp request failed: %v", err)
		}
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		return resp, respBody
	}

	first, firstBody := postOnce()
	if first.StatusCode != http.StatusOK {
		t.Fatalf("expected first request-otp to succeed, got %d: %s", first.StatusCode, firstBody)
	}

	second, secondBody := postOnce()
	if second.StatusCode == http.StatusOK {
		t.Errorf("expected a second, immediate request-otp for the same value to be blocked, got 200 OK: %s", secondBody)
	}
}

func TestClassicPassportOtp_HTTP_SigninWithValidCode(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ensureInviteEmailSendingConfigured(t, cfg)
	email, _, _ := freshSignupTarget(t, cfg, "otp-signin")
	code := requestOtpAndFetchCode(t, cfg, email)

	body, _ := json.Marshal(map[string]any{"value": email, "otp": code})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/passport/otp"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build otp confirm request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("otp confirm request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK confirming a valid otp, got %d: %s", resp.StatusCode, respBody)
	}

	var out signupResult
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.Session.Token == "" {
		t.Errorf("expected a session token from a valid otp confirm, got none: %s", respBody)
	}
}

func TestClassicPassportOtp_HTTP_RejectsInvalidCode(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ensureInviteEmailSendingConfigured(t, cfg)
	email, _, _ := freshSignupTarget(t, cfg, "otp-wrong")
	requestOtpAndFetchCode(t, cfg, email) // sends a real otp, deliberately not used below

	body, _ := json.Marshal(map[string]any{"value": email, "otp": "000000"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/passport/otp"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build otp confirm request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("otp confirm request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected a wrong otp code to be rejected, got 200 OK: %s", respBody)
	}
}

// setWorkspaceConfigForceTotp toggles root's WorkspaceConfig.forceTotp - see
// WorkspaceConfigDistinctUpdateAction (upserts by workspaceId=ROOT_VAR, ignoring any
// :uniqueId). Returns a cleanup func that restores it to false, since this is a single
// global row shared by every other test/spec that signs in during the same run -
// leaving forceTotp on would silently break every other signin.
func setWorkspaceConfigForceTotp(t *testing.T, cfg TestConfig, enabled bool) func() {
	t.Helper()
	apply := func(v bool) {
		body, _ := json.Marshal(map[string]any{"forceTotp": v})
		client := cfg.NewHTTPClient()
		req, err := http.NewRequest(http.MethodPatch, cfg.URL("/workspace-config/distinct"), bytes.NewReader(body))
		if err != nil {
			t.Fatalf("failed to build workspace-config/distinct request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", cfg.CliToken)
		req.Header.Set("Workspace-id", "root")
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("workspace-config/distinct request failed: %v", err)
		}
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200 OK setting forceTotp=%v, got %d: %s", v, resp.StatusCode, respBody)
		}
	}
	apply(enabled)
	return func() { apply(false) }
}

func TestConfirmClassicPassportTotp_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	email, _, _ := freshSignupTarget(t, cfg, "totp-confirm")

	restore := setWorkspaceConfigForceTotp(t, cfg, true)
	defer restore()

	// The first signin attempt with forceTotp on assigns a totpSecret to the passport
	// and returns {next:["setup-totp"]} instead of a session - see
	// ClassicSigninActionImplementation.go. It doesn't return the raw secret in a form
	// this test can feed straight to pquerna/otp/totp (only a full otpauth:// URL), so
	// read the persisted secret back via PassportBrowse instead (root's CLI token
	// covers it, same as publicAuthentication above).
	signinResp, signinBody := postSignin(t, cfg, email, "checkendpointtests-pass-123")
	if signinResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK on the totp-setup-triggering signin, got %d: %s", signinResp.StatusCode, signinBody)
	}

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/passport/browse?itemsPerPage=1000"), nil)
	if err != nil {
		t.Fatalf("failed to build passport browse request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("passport browse request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK browsing passport, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[struct {
		Value      string `json:"value"`
		TotpSecret string `json:"totpSecret"`
	}]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode passport browse response: %v\nbody: %s", err, body)
	}
	secret := ""
	for _, item := range list.Data.Items {
		if item.Value == email {
			secret = item.TotpSecret
		}
	}
	if secret == "" {
		t.Fatalf("expected the signed-up passport %q to have a totpSecret after the forced-totp signin, found none", email)
	}

	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		t.Fatalf("failed to generate a totp code from the persisted secret: %v", err)
	}

	confirmBody, _ := json.Marshal(map[string]any{
		"value":    email,
		"password": "checkendpointtests-pass-123",
		"totpCode": code,
	})
	confirmReq, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/totp/confirm"), bytes.NewReader(confirmBody))
	if err != nil {
		t.Fatalf("failed to build totp confirm request: %v", err)
	}
	confirmReq.Header.Set("Content-Type", "application/json")
	confirmResp, err := client.Do(confirmReq)
	if err != nil {
		t.Fatalf("totp confirm request failed: %v", err)
	}
	defer confirmResp.Body.Close()
	confirmRespBody, _ := io.ReadAll(confirmResp.Body)
	if confirmResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK confirming a valid totp code, got %d: %s", confirmResp.StatusCode, confirmRespBody)
	}

	var out signupResult
	if err := json.Unmarshal(confirmRespBody, &out); err != nil {
		t.Fatalf("failed to decode totp confirm response: %v\nbody: %s", err, confirmRespBody)
	}
	if out.Data.Item.Session.Token == "" {
		t.Errorf("expected a session token from a confirmed totp signin, got none: %s", confirmRespBody)
	}
}

func TestConfirmClassicPassportTotp_HTTP_RejectsWrongCode(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	email, _, _ := freshSignupTarget(t, cfg, "totp-wrong")

	restore := setWorkspaceConfigForceTotp(t, cfg, true)
	defer restore()

	// Trigger totpSecret assignment the same way as the success case above.
	if resp, body := postSignin(t, cfg, email, "checkendpointtests-pass-123"); resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK on the totp-setup-triggering signin, got %d: %s", resp.StatusCode, body)
	}

	confirmBody, _ := json.Marshal(map[string]any{
		"value":    email,
		"password": "checkendpointtests-pass-123",
		"totpCode": "000000",
	})
	client := cfg.NewHTTPClient()
	confirmReq, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/totp/confirm"), bytes.NewReader(confirmBody))
	if err != nil {
		t.Fatalf("failed to build totp confirm request: %v", err)
	}
	confirmReq.Header.Set("Content-Type", "application/json")
	confirmResp, err := client.Do(confirmReq)
	if err != nil {
		t.Fatalf("totp confirm request failed: %v", err)
	}
	defer confirmResp.Body.Close()
	confirmRespBody, _ := io.ReadAll(confirmResp.Body)
	if confirmResp.StatusCode == http.StatusOK {
		t.Fatalf("expected a wrong totp code to be rejected, got 200 OK: %s", confirmRespBody)
	}
}

func TestOauthAuthenticate_HTTP_RejectsUnsupportedService(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	body, _ := json.Marshal(map[string]any{"service": "not-a-real-provider", "token": "whatever"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/via-oauth"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build via-oauth request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("via-oauth request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unsupported oauth service to be rejected, got 200 OK: %s", respBody)
	}
}

// TestOauthAuthenticate_HTTP_RejectsInvalidGoogleToken makes a real call out to
// Google's tokeninfo endpoint (authenticateWithGoogle has no test-seam to mock it) - a
// garbage access token must come back rejected either way, so this only needs outbound
// network access to succeed, not a real Google account.
func TestOauthAuthenticate_HTTP_RejectsInvalidGoogleToken(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	body, _ := json.Marshal(map[string]any{"service": "google", "token": fmt.Sprintf("checkendpointtests-not-a-real-token-%d", time.Now().UnixNano())})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/via-oauth"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build via-oauth request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("via-oauth request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an invalid google token to be rejected, got 200 OK: %s", respBody)
	}
}

// TestOsLoginAuthenticate_HTTP_Succeeds covers GET /passports/os/login
// (SigninWithOsUser2), which signs in as a user derived from the server process's own
// OS account (see UserCli.go's GetOsHostUserRoleWorkspaceDef) - no request body, no
// prior fixture needed, and no auth required to call it (nil security model).
func TestOsLoginAuthenticate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/passports/os/login"), nil)
	if err != nil {
		t.Fatalf("failed to build os/login request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("os/login request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK from os/login, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[struct {
		Token string `json:"token"`
	}]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode os/login response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.Token == "" {
		t.Errorf("expected a session token from os/login, got none: %s", body)
	}
}

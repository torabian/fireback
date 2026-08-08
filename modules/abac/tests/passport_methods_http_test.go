package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

// checkPassportMethodsRes mirrors abac.CheckPassportMethodsActionRes (see
// modules/abac/CheckPassportMethodsAction.dyno.go) - duplicated here rather than
// imported, since these are black-box tests against the HTTP surface, not the Go
// package itself.
type checkPassportMethodsRes struct {
	Email                bool   `json:"email"`
	Phone                bool   `json:"phone"`
	Google               bool   `json:"google"`
	Facebook             bool   `json:"facebook"`
	GoogleOAuthClientKey string `json:"googleOAuthClientKey"`
	FacebookAppId        string `json:"facebookAppId"`
	EnabledRecaptcha2    bool   `json:"enabledRecaptcha2"`
	Recaptcha2ClientKey  string `json:"recaptcha2ClientKey"`
}

// googleResponseEnvelope mirrors fireback.GoogleResponse[T]/GResponseSingleItem's
// {"data":{"item": ...}} wrapper every fireback action response uses.
type googleResponseEnvelope[T any] struct {
	Data struct {
		Item T `json:"item"`
	} `json:"data"`
}

// TestCheckPassportMethods_HTTP hits GET /passports/available-methods (abac's
// CheckPassportMethodsAction) the same way a browser's signin screen would: no
// authorization header, since PassportMethod discovery is intentionally public (see
// CheckPassportMethodsAction.go - fireback.ResolveActionContext(c, nil), a nil security
// model).
func TestCheckPassportMethods_HTTP(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/passports/available-methods"), nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request to %s failed: %v", req.URL, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d %s\nbody: %s", resp.StatusCode, resp.Status, body)
	}

	var out googleResponseEnvelope[checkPassportMethodsRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode JSON response: %v\nbody: %s", err, body)
	}

	item := out.Data.Item

	// At least one passport method needs to be seeded/enabled for a working signin
	// screen - if none are, PassportMethodSyncSeeders() likely hasn't run (see
	// SeedersSync in cmd/fireback/main.go) or every passportMethod row got deleted.
	if !item.Email && !item.Phone && !item.Google && !item.Facebook {
		t.Errorf("expected at least one authentication method to be enabled, got none: %+v", item)
	}

	// Field-level sanity: an enabled OAuth method should always carry its client
	// key/app id alongside it - a true flag with an empty key means the signin screen
	// would render a broken OAuth button.
	if item.Google && item.GoogleOAuthClientKey == "" {
		t.Errorf("google is enabled but googleOAuthClientKey is empty")
	}
	if item.Facebook && item.FacebookAppId == "" {
		t.Errorf("facebook is enabled but facebookAppId is empty")
	}
	if item.EnabledRecaptcha2 && item.Recaptcha2ClientKey == "" {
		t.Errorf("enabledRecaptcha2 is true but recaptcha2ClientKey is empty")
	}
}

// TestCheckPassportMethods_HTTP_NoAuthRequired confirms this endpoint stays public even
// without any authorization header - a regression here would lock every unauthenticated
// visitor out of the signin screen before they can even see what methods exist.
func TestCheckPassportMethods_HTTP_NoAuthRequired(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	// Deliberately no authorization/workspace-id headers.
	resp, err := client.Get(cfg.URL("/passports/available-methods"))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected the public methods endpoint to work without auth, got %d: %s", resp.StatusCode, body)
	}
}

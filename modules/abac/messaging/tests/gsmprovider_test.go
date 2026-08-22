// gsmProvider black-box tests, following webpushconfig_test.go's exact conventions (see
// that file's package doc comment for the full rationale). Reuses its
// googleResponseEnvelope[T]/googleResponseListEnvelope[T] types - same package, no need
// to redeclare them here.
package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"testing"

	abactests "github.com/torabian/fireback/modules/abac/tests"
)

// gsmProviderRes mirrors messaging.GsmProviderEntity's JSON response shape - only the
// fields these tests actually read.
type gsmProviderRes struct {
	UniqueId         string `json:"uniqueId"`
	ApiKey           string `json:"apiKey"`
	MainSenderNumber string `json:"mainSenderNumber"`
	Type             string `json:"type"`
	WorkspaceId      string `json:"workspaceId"`
}

func postGsmProvider(t *testing.T, cfg abactests.TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/gsmProvider"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

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

// createSampleGsmProvider is a shared helper (not a test itself) for the Browse/Get/
// Update/Delete tests below, which all need an existing, valid ("terminal" type, so no
// real gateway is ever actually invoked - see GsmSendSMSByTerminal) record to act on.
func createSampleGsmProvider(t *testing.T, cfg abactests.TestConfig) gsmProviderRes {
	t.Helper()

	resp, body := postGsmProvider(t, cfg, map[string]any{
		"mainSenderNumber": "+10000000000",
		"type":             "terminal",
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating gsmProvider, got %d: %s", resp.StatusCode, body)
	}

	var created googleResponseEnvelope[gsmProviderRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deleteGsmProvider(t *testing.T, cfg abactests.TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}

	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/gsmProvider/delete"), bytes.NewReader(body))
	if err != nil {
		t.Logf("cleanup: failed to build delete request for gsmProvider %s: %v", uniqueId, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("cleanup: failed to delete gsmProvider %s: %v", uniqueId, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		t.Logf("cleanup: deleting gsmProvider %s returned %d: %s", uniqueId, resp.StatusCode, b)
	}
}

func TestGsmProviderCreate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleGsmProvider(t, cfg)
	defer deleteGsmProvider(t, cfg, created.UniqueId)

	if created.Type != "terminal" {
		t.Errorf("expected type %q, got %q", "terminal", created.Type)
	}
}

func TestGsmProviderCreate_HTTP_RequiresAuth(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	body, _ := json.Marshal(map[string]any{"mainSenderNumber": "+10000000000", "type": "terminal"})
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/gsmProvider"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		var out googleResponseEnvelope[gsmProviderRes]
		if err := json.Unmarshal(respBody, &out); err == nil {
			deleteGsmProvider(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected an unauthenticated create to be rejected, got 200 OK: %s", respBody)
	}
}

// TestGsmProviderCreate_HTTP_ValidationRequiredFields covers the fix making
// GsmProviderCreateAction actually call fireback.CommonStructValidatorPointer -
// mainSenderNumber/type are both `validate:"required"` (see Messaging.emi.yml) but were
// previously accepted empty.
func TestGsmProviderCreate_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	cases := []struct {
		name         string
		payload      map[string]any
		wantLocation string
	}{
		{"missing mainSenderNumber", map[string]any{"type": "terminal"}, "mainSenderNumber"},
		{"missing type", map[string]any{"mainSenderNumber": "+10000000000"}, "type"},
		{"empty body", map[string]any{}, "mainSenderNumber"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := postGsmProvider(t, cfg, tc.payload)

			if resp.StatusCode == http.StatusOK {
				var out googleResponseEnvelope[gsmProviderRes]
				if err := json.Unmarshal(body, &out); err == nil {
					deleteGsmProvider(t, cfg, out.Data.Item.UniqueId)
				}
				t.Fatalf("expected creation with %s to be rejected, got 200 OK: %s", tc.name, body)
			}
			if !strings.Contains(string(body), `"location": "`+tc.wantLocation+`"`) &&
				!strings.Contains(string(body), `"location":"`+tc.wantLocation+`"`) {
				t.Errorf("expected a field error located at %q, got %d: %s", tc.wantLocation, resp.StatusCode, body)
			}
		})
	}
}

// TestGsmProviderBrowse_HTTP_IncludesOwnRecord covers the workspace-stamping fix: without
// UserId/WorkspaceId set on create, the row is invisible to this workspace-scoped browse.
func TestGsmProviderBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleGsmProvider(t, cfg)
	defer deleteGsmProvider(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/gsmProvider/browse"), nil)
	if err != nil {
		t.Fatalf("failed to build browse request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("browse request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK browsing gsmProvider, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[gsmProviderRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}

	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created gsmProvider %s (workspaceId %q), got %d items - Create may not be stamping workspaceId from the resolved query context", created.UniqueId, cfg.WorkspaceID, len(list.Data.Items))
}

func TestGsmProviderGet_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleGsmProvider(t, cfg)
	defer deleteGsmProvider(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/gsmProvider/"+created.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build get request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("get request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK getting gsmProvider, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[gsmProviderRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode get response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UniqueId != created.UniqueId {
		t.Errorf("expected uniqueId %q, got %q", created.UniqueId, out.Data.Item.UniqueId)
	}
}

func TestGsmProviderUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleGsmProvider(t, cfg)
	defer deleteGsmProvider(t, cfg, created.UniqueId)

	body, _ := json.Marshal(map[string]any{"mainSenderNumber": "+19999999999"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/gsmProvider/"+created.UniqueId), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build update request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("update request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK updating gsmProvider, got %d: %s", resp.StatusCode, respBody)
	}
}

// TestGsmProviderAwareDeletePreview_HTTP_ThenDelete covers both delete-preview and the
// actual delete (its bare action name "GsmProviderAwareDelete" is a prefix of
// "GsmProviderAwareDeletePreview", so this single test name covers both per
// tools/checkendpointtests' substring match, same trick webpushconfig_test.go uses).
func TestGsmProviderAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleGsmProvider(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/gsmProvider/delete-preview?uniqueIds="+created.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build delete-preview request: %v", err)
	}
	previewReq.Header.Set("Authorization", cfg.CliToken)
	previewReq.Header.Set("Workspace-id", cfg.WorkspaceID)
	previewResp, err := client.Do(previewReq)
	if err != nil {
		t.Fatalf("delete-preview request failed: %v", err)
	}
	defer previewResp.Body.Close()
	previewBody, _ := io.ReadAll(previewResp.Body)
	if previewResp.StatusCode != http.StatusOK {
		deleteGsmProvider(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/gsmProvider/delete"), bytes.NewReader(deleteBody))
	if err != nil {
		t.Fatalf("failed to build delete request: %v", err)
	}
	deleteReq.Header.Set("Content-Type", "application/json")
	deleteReq.Header.Set("Authorization", cfg.CliToken)
	deleteReq.Header.Set("Workspace-id", cfg.WorkspaceID)
	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		t.Fatalf("delete request failed: %v", err)
	}
	defer deleteResp.Body.Close()
	deleteRespBody, _ := io.ReadAll(deleteResp.Body)
	if deleteResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK deleting gsmProvider, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/gsmProvider/"+created.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build post-delete get request: %v", err)
	}
	getReq.Header.Set("Authorization", cfg.CliToken)
	getReq.Header.Set("Workspace-id", cfg.WorkspaceID)
	getResp, err := client.Do(getReq)
	if err != nil {
		t.Fatalf("post-delete get request failed: %v", err)
	}
	defer getResp.Body.Close()
	if getResp.StatusCode == http.StatusOK {
		t.Errorf("expected the deleted gsmProvider %s to no longer be gettable", created.UniqueId)
	}
}

func TestGsmProviderCreate_CLI_Help(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	// The bare Name every entity's create command registers under ("create") collides
	// across every entity nested in the "messaging" group - each is only unambiguously
	// reachable via its own alias (see GsmProviderCreateActionCliHandler's
	// cmd.Aliases = []string{meta.CliShort}; confirmed against `./app messaging --help`).
	cmd := exec.CommandContext(ctx, bin, "messaging", "gsm provider c", "--help")
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s messaging gsm provider c --help` failed: %v\noutput:\n%s", bin, err, out)
	}

	// urfave's default help renderer only prints the primary Name ("create", shared by
	// every entity in this group) in the NAME:/USAGE: header, never the Aliases actually
	// used to route here - so the entity name from the auto-generated description
	// ('Creates a new "gsmProvider" row.') is what actually proves this reached the
	// right subcommand, not a coincidentally-successful but wrong one.
	if !strings.Contains(string(out), "gsmProvider") {
		t.Errorf("expected --help output to mention the gsmProvider entity, got:\n%s", out)
	}
}

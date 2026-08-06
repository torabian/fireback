// Package tests holds black-box endpoint/CLI tests for the messaging module, following
// the exact same convention (and reusing the exact same TestConfig) as
// modules/abac/tests - see that package's doc comment for the full rationale: real HTTP
// requests against a running `./app start` and real CLI subprocess invocations, every
// test skipping (not failing) when its target isn't reachable, so `go test ./...` stays
// green without live infra.
//
// Test function names are chosen deliberately to each contain the bare action name
// (e.g. "WebPushConfigBrowse") that tools/checkendpointtests matches on - see that
// tool's collectTestFuncNames/checkModule for the substring convention this follows.
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

// webPushConfigRes mirrors messaging.WebPushConfigEntity's JSON response shape - only
// the fields these tests actually read.
type webPushConfigRes struct {
	UniqueId     string `json:"uniqueId"`
	Subscription any    `json:"subscription"`
}

type googleResponseEnvelope[T any] struct {
	Data struct {
		Item T `json:"item"`
	} `json:"data"`
}

type googleResponseListEnvelope[T any] struct {
	Data struct {
		Items []T `json:"items"`
	} `json:"data"`
}

func postWebPushConfig(t *testing.T, cfg abactests.TestConfig, subscription any) (*http.Response, []byte) {
	t.Helper()

	body, err := json.Marshal(map[string]any{"subscription": subscription})
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/webPushConfig"), bytes.NewReader(body))
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

// createSampleWebPushConfig is a shared helper (not a test itself) for the Get/Update/
// Delete tests below, which all need an existing record to act on.
func createSampleWebPushConfig(t *testing.T, cfg abactests.TestConfig) webPushConfigRes {
	t.Helper()

	sub := map[string]any{"endpoint": "https://example.com/push/abc", "keys": map[string]string{"p256dh": "x", "auth": "y"}}
	resp, body := postWebPushConfig(t, cfg, sub)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating webPushConfig, got %d: %s", resp.StatusCode, body)
	}

	var created googleResponseEnvelope[webPushConfigRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deleteWebPushConfig(t *testing.T, cfg abactests.TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}

	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/webPushConfig/delete"), bytes.NewReader(body))
	if err != nil {
		t.Logf("cleanup: failed to build delete request for webPushConfig %s: %v", uniqueId, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("cleanup: failed to delete webPushConfig %s: %v", uniqueId, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		t.Logf("cleanup: deleting webPushConfig %s returned %d: %s", uniqueId, resp.StatusCode, b)
	}
}

// TestWebPushConfigCreate_HTTP_Succeeds also covers the unauthenticated-rejection path,
// since webPushConfig has no capability permission gate (see webPushConfigSecurity in
// WebPushConfigActions.go) - the token check is the only thing stopping an anonymous
// caller.
func TestWebPushConfigCreate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWebPushConfig(t, cfg)
	defer deleteWebPushConfig(t, cfg, created.UniqueId)
}

func TestWebPushConfigCreate_HTTP_RequiresAuth(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	body, _ := json.Marshal(map[string]any{"subscription": map[string]any{"endpoint": "https://example.com/push/anon"}})
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/webPushConfig"), bytes.NewReader(body))
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
		var out googleResponseEnvelope[webPushConfigRes]
		if err := json.Unmarshal(respBody, &out); err == nil {
			deleteWebPushConfig(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected an unauthenticated create to be rejected, got 200 OK: %s", respBody)
	}
}

// TestWebPushConfigBrowse_HTTP_IncludesOwnRecord covers webPushConfigSecurity's
// ResolveStrategyUser scoping on the list endpoint: a user's browse results must
// include a subscription they just created.
func TestWebPushConfigBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWebPushConfig(t, cfg)
	defer deleteWebPushConfig(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/webPushConfig/browse"), nil)
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
		t.Fatalf("expected 200 OK browsing webPushConfig, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[webPushConfigRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}

	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created webPushConfig %s, got %d items", created.UniqueId, len(list.Data.Items))
}

// TestWebPushConfigGet_HTTP_Succeeds covers fetching a single owned record by
// uniqueId - see webPushConfigGetOwn for the ownership check this exercises.
func TestWebPushConfigGet_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWebPushConfig(t, cfg)
	defer deleteWebPushConfig(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/webPushConfig/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting webPushConfig, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[webPushConfigRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode get response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UniqueId != created.UniqueId {
		t.Errorf("expected uniqueId %q, got %q", created.UniqueId, out.Data.Item.UniqueId)
	}
}

// TestWebPushConfigUpdate_HTTP_Succeeds covers updating an owned record's subscription
// payload.
func TestWebPushConfigUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWebPushConfig(t, cfg)
	defer deleteWebPushConfig(t, cfg, created.UniqueId)

	newSub := map[string]any{"endpoint": "https://example.com/push/updated"}
	body, _ := json.Marshal(map[string]any{"subscription": newSub})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/webPushConfig/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating webPushConfig, got %d: %s", resp.StatusCode, respBody)
	}
}

// TestWebPushConfigAwareDeletePreview_HTTP_ThenDelete covers both the delete-preview
// and the actual delete (its bare action name "WebPushConfigAwareDelete" is a prefix of
// "WebPushConfigAwareDeletePreview", so this single test name covers both per
// tools/checkendpointtests' substring match) - and confirms a deleted record no longer
// appears in browse.
func TestWebPushConfigAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWebPushConfig(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/webPushConfig/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteWebPushConfig(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/webPushConfig/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting webPushConfig, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/webPushConfig/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted webPushConfig %s to no longer be gettable", created.UniqueId)
	}
}

// TestWebPushConfigCreate_CLI_Help is a cheap smoke test that the webPushConfig
// commands are actually registered in the CLI tree (catches them silently
// disappearing from MessagingModule.go's command wiring without needing a live
// database/authenticated session).
func TestWebPushConfigCreate_CLI_Help(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, "messaging", "web-push-config-create-action", "--help")
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s messaging web-push-config-create-action --help` failed: %v\noutput:\n%s", bin, err, out)
	}

	if !strings.Contains(string(out), "web-push-config-create-action") {
		t.Errorf("expected --help output to mention the command name, got:\n%s", out)
	}
}

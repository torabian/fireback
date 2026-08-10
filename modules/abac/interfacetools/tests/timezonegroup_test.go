// timezoneGroup black-box tests, following appmenu_test.go's exact conventions (same
// package, reuses its googleResponseEnvelope/googleResponseListEnvelope types).
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

// timezoneGroupRes mirrors interfacetools.TimezoneGroupEntity's JSON response shape -
// only the fields these tests actually read.
type timezoneGroupRes struct {
	UniqueId string `json:"uniqueId"`
	Title    string `json:"title"`
}

func postTimezoneGroup(t *testing.T, cfg abactests.TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/timezoneGroup"), bytes.NewReader(body))
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

// createSampleTimezoneGroup is a shared helper (not a test itself) for the Browse/Get/
// Update/Delete tests, which all need an existing record to act on.
func createSampleTimezoneGroup(t *testing.T, cfg abactests.TestConfig) timezoneGroupRes {
	t.Helper()

	resp, body := postTimezoneGroup(t, cfg, map[string]any{
		"title": "checkendpointtests timezone",
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating timezoneGroup, got %d: %s", resp.StatusCode, body)
	}

	var created googleResponseEnvelope[timezoneGroupRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deleteTimezoneGroup(t *testing.T, cfg abactests.TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}

	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/timezoneGroup/delete"), bytes.NewReader(body))
	if err != nil {
		t.Logf("cleanup: failed to build delete request for timezoneGroup %s: %v", uniqueId, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("cleanup: failed to delete timezoneGroup %s: %v", uniqueId, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		t.Logf("cleanup: deleting timezoneGroup %s returned %d: %s", uniqueId, resp.StatusCode, b)
	}
}

func TestTimezoneGroupCreate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTimezoneGroup(t, cfg)
	defer deleteTimezoneGroup(t, cfg, created.UniqueId)

	if created.Title != "checkendpointtests timezone" {
		t.Errorf("expected title %q, got %q", "checkendpointtests timezone", created.Title)
	}
}

func TestTimezoneGroupCreate_HTTP_RequiresAuth(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	body, _ := json.Marshal(map[string]any{"title": "no auth"})
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/timezoneGroup"), bytes.NewReader(body))
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
		var out googleResponseEnvelope[timezoneGroupRes]
		if err := json.Unmarshal(respBody, &out); err == nil {
			deleteTimezoneGroup(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected an unauthenticated create to be rejected, got 200 OK: %s", respBody)
	}
}

func TestTimezoneGroupBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTimezoneGroup(t, cfg)
	defer deleteTimezoneGroup(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/timezoneGroup/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing timezoneGroup, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[timezoneGroupRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}

	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created timezoneGroup %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestTimezoneGroupGet_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTimezoneGroup(t, cfg)
	defer deleteTimezoneGroup(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/timezoneGroup/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting timezoneGroup, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[timezoneGroupRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode get response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UniqueId != created.UniqueId {
		t.Errorf("expected uniqueId %q, got %q", created.UniqueId, out.Data.Item.UniqueId)
	}
}

func TestTimezoneGroupUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTimezoneGroup(t, cfg)
	defer deleteTimezoneGroup(t, cfg, created.UniqueId)

	body, _ := json.Marshal(map[string]any{"title": "renamed by test"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/timezoneGroup/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating timezoneGroup, got %d: %s", resp.StatusCode, respBody)
	}

	var out googleResponseEnvelope[timezoneGroupRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.Title != "renamed by test" {
		t.Errorf("expected title %q after update, got %q", "renamed by test", out.Data.Item.Title)
	}
}

// TestTimezoneGroupAwareDeletePreview_HTTP_ThenDelete covers both delete-preview and the
// actual delete, same substring trick as appmenu_test.go's equivalent.
func TestTimezoneGroupAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTimezoneGroup(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/timezoneGroup/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteTimezoneGroup(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/timezoneGroup/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting timezoneGroup, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/timezoneGroup/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted timezoneGroup %s to no longer be gettable", created.UniqueId)
	}
}

func TestTimezoneGroupCreate_CLI_Help(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	// "timezone" (alias "tz") is the entity group nested under "interface" - see
	// InterfaceToolsModule.go.
	cmd := exec.CommandContext(ctx, bin, "interface", "timezone", "timezoneGroup-c", "--help")
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s interface timezone timezoneGroup-c --help` failed: %v\noutput:\n%s", bin, err, out)
	}
	if !strings.Contains(string(out), "timezoneGroup") {
		t.Errorf("expected --help output to mention the timezoneGroup entity, got:\n%s", out)
	}
}

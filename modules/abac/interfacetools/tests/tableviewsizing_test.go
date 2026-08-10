// tableViewSizing black-box tests, following appmenu_test.go's exact conventions (same
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

// tableViewSizingRes mirrors interfacetools.TableViewSizingEntity's JSON response shape -
// only the fields these tests actually read.
type tableViewSizingRes struct {
	UniqueId  string `json:"uniqueId"`
	TableName string `json:"tableName"`
	Sizes     string `json:"sizes"`
}

func postTableViewSizing(t *testing.T, cfg abactests.TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/tableViewSizing"), bytes.NewReader(body))
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

// createSampleTableViewSizing is a shared helper (not a test itself) for the Browse/Get/
// Update/Delete tests, which all need an existing record to act on.
func createSampleTableViewSizing(t *testing.T, cfg abactests.TestConfig) tableViewSizingRes {
	t.Helper()

	resp, body := postTableViewSizing(t, cfg, map[string]any{
		"tableName": "checkendpointtests-table",
		"sizes":     `[{"columnName":"uniqueId","width":120}]`,
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating tableViewSizing, got %d: %s", resp.StatusCode, body)
	}

	var created googleResponseEnvelope[tableViewSizingRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deleteTableViewSizing(t *testing.T, cfg abactests.TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}

	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/tableViewSizing/delete"), bytes.NewReader(body))
	if err != nil {
		t.Logf("cleanup: failed to build delete request for tableViewSizing %s: %v", uniqueId, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("cleanup: failed to delete tableViewSizing %s: %v", uniqueId, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		t.Logf("cleanup: deleting tableViewSizing %s returned %d: %s", uniqueId, resp.StatusCode, b)
	}
}

func TestTableViewSizingCreate_HTTP_ValidationRequiresTableName(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postTableViewSizing(t, cfg, map[string]any{"sizes": "[]"})
	if resp.StatusCode == http.StatusOK {
		var out googleResponseEnvelope[tableViewSizingRes]
		if err := json.Unmarshal(body, &out); err == nil {
			deleteTableViewSizing(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected creation with a missing tableName to be rejected, got 200 OK: %s", body)
	}
}

func TestTableViewSizingCreate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTableViewSizing(t, cfg)
	defer deleteTableViewSizing(t, cfg, created.UniqueId)

	if created.TableName != "checkendpointtests-table" {
		t.Errorf("expected tableName %q, got %q", "checkendpointtests-table", created.TableName)
	}
}

func TestTableViewSizingBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTableViewSizing(t, cfg)
	defer deleteTableViewSizing(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/tableViewSizing/browse"), nil)
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
		t.Fatalf("expected 200 OK browsing tableViewSizing, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[tableViewSizingRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}

	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created tableViewSizing %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestTableViewSizingGet_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTableViewSizing(t, cfg)
	defer deleteTableViewSizing(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/tableViewSizing/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting tableViewSizing, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[tableViewSizingRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode get response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UniqueId != created.UniqueId {
		t.Errorf("expected uniqueId %q, got %q", created.UniqueId, out.Data.Item.UniqueId)
	}
}

func TestTableViewSizingUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTableViewSizing(t, cfg)
	defer deleteTableViewSizing(t, cfg, created.UniqueId)

	body, _ := json.Marshal(map[string]any{"sizes": `[{"columnName":"uniqueId","width":240}]`})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/tableViewSizing/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating tableViewSizing, got %d: %s", resp.StatusCode, respBody)
	}

	var out googleResponseEnvelope[tableViewSizingRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if !strings.Contains(out.Data.Item.Sizes, "240") {
		t.Errorf("expected sizes to contain the updated width 240, got %q", out.Data.Item.Sizes)
	}
}

// TestTableViewSizingUpdate_HTTP_UpsertsWhenMissing covers the upsert-by-caller-chosen-
// uniqueId behavior TableViewSizingUpdateAction falls back to on a 404 (see
// TableViewSizingActions.go) - CommonListManager.tsx addresses this entity by a
// per-table, per-user key it picks itself, so the very first PATCH for a given key must
// succeed by creating the row, not fail because nothing exists yet at that uniqueId.
func TestTableViewSizingUpdate_HTTP_UpsertsWhenMissing(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	missingId := "checkendpointtests-upsert-key"
	defer deleteTableViewSizing(t, cfg, missingId)

	body, _ := json.Marshal(map[string]any{
		"tableName": "checkendpointtests-upsert-table",
		"sizes":     `[{"columnName":"uniqueId","width":100}]`,
	})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/tableViewSizing/"+missingId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK upserting a not-yet-existing tableViewSizing, got %d: %s", resp.StatusCode, respBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/tableViewSizing/"+missingId), nil)
	if err != nil {
		t.Fatalf("failed to build post-upsert get request: %v", err)
	}
	getReq.Header.Set("Authorization", cfg.CliToken)
	getReq.Header.Set("Workspace-id", cfg.WorkspaceID)
	getResp, err := client.Do(getReq)
	if err != nil {
		t.Fatalf("post-upsert get request failed: %v", err)
	}
	defer getResp.Body.Close()
	if getResp.StatusCode != http.StatusOK {
		getBody, _ := io.ReadAll(getResp.Body)
		t.Fatalf("expected the upserted tableViewSizing to now be gettable, got %d: %s", getResp.StatusCode, getBody)
	}
}

// TestTableViewSizingAwareDeletePreview_HTTP_ThenDelete covers both delete-preview and
// the actual delete, same substring trick as appmenu_test.go's equivalent.
func TestTableViewSizingAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleTableViewSizing(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/tableViewSizing/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteTableViewSizing(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/tableViewSizing/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting tableViewSizing, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/tableViewSizing/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted tableViewSizing %s to no longer be gettable", created.UniqueId)
	}
}

func TestTableViewSizingCreate_CLI_Help(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	// "tableviewsizing" (alias "tvs") is the entity group nested under "interface" -
	// see InterfaceToolsModule.go.
	cmd := exec.CommandContext(ctx, bin, "interface", "tableviewsizing", "tableViewSizing-c", "--help")
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s interface tableviewsizing tableViewSizing-c --help` failed: %v\noutput:\n%s", bin, err, out)
	}
	if !strings.Contains(string(out), "tableViewSizing") {
		t.Errorf("expected --help output to mention the tableViewSizing entity, got:\n%s", out)
	}
}

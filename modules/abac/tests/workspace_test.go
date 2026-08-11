// workspace black-box tests, following capability_test.go's conventions (same package,
// reuses googleResponseEnvelope/googleResponseListEnvelope). Create/Update/AwareDelete/
// AwareDeletePreview all have AllowOnRoot:true (see WorkspaceActions.go), so every write
// here must run in the literal "root" workspace. Create requires an existing typeId
// (validate:"required") - reuses workspace_type_test.go's createSampleWorkspaceType/
// deleteWorkspaceType helpers (same package) for that, which in turn reuses
// role_crud_test.go's createSampleRole/deleteRole.
//
// Note: CreateWorkspaceAction (a hand-declared action, tested separately in
// misc_account_workspace_test.go) is a different, higher-level endpoint - it wraps
// CreateWorkspaceAndAssignUser to also enroll the caller as a member. WorkspaceCreateAction
// tested here is the plain entity CRUD endpoint (POST /workspace), which just inserts the
// row with no membership side effects.
package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type workspaceRes struct {
	UniqueId string `json:"uniqueId"`
	Name     string `json:"name"`
	TypeId   string `json:"typeId"`
}

func postWorkspace(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

// createSampleWorkspace also returns the backing workspaceType/role it created, so
// callers can clean all three up.
func createSampleWorkspace(t *testing.T, cfg TestConfig) (workspaceRes, workspaceTypeRes, roleRes) {
	t.Helper()
	wt, role := createSampleWorkspaceType(t, cfg)
	resp, body := postWorkspace(t, cfg, map[string]any{
		"name":   "checkendpointtests workspace",
		"typeId": wt.UniqueId,
	})
	if resp.StatusCode != http.StatusOK {
		deleteWorkspaceType(t, cfg, wt.UniqueId)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("expected 200 OK creating workspace, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[workspaceRes]
	if err := json.Unmarshal(body, &created); err != nil {
		deleteWorkspaceType(t, cfg, wt.UniqueId)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		deleteWorkspaceType(t, cfg, wt.UniqueId)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item, wt, role
}

func deleteWorkspace(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/delete"), bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func TestWorkspaceCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, wt, role := createSampleWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, created.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	if created.TypeId != wt.UniqueId {
		t.Errorf("expected typeId %q, got %q", wt.UniqueId, created.TypeId)
	}
}

func TestWorkspaceBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, wt, role := createSampleWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, created.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace/browse?itemsPerPage=1000"), nil)
	if err != nil {
		t.Fatalf("failed to build browse request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("browse request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK browsing workspace, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[workspaceRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created workspace %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestWorkspaceGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, wt, role := createSampleWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, created.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace/"+created.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build get request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("get request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK getting workspace, got %d: %s", resp.StatusCode, body)
	}
}

func TestWorkspaceUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, wt, role := createSampleWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, created.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	body, _ := json.Marshal(map[string]any{"name": "checkendpointtests renamed workspace"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/workspace/"+created.UniqueId), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build update request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("update request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK updating workspace, got %d: %s", resp.StatusCode, respBody)
	}
	var out googleResponseEnvelope[workspaceRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.Name != "checkendpointtests renamed workspace" {
		t.Errorf("expected name %q after update, got %q", "checkendpointtests renamed workspace", out.Data.Item.Name)
	}
}

func TestWorkspaceAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, wt, role := createSampleWorkspace(t, cfg)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace/delete-preview?uniqueIds="+created.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build delete-preview request: %v", err)
	}
	previewReq.Header.Set("Authorization", cfg.CliToken)
	previewReq.Header.Set("Workspace-id", "root")
	previewResp, err := client.Do(previewReq)
	if err != nil {
		t.Fatalf("delete-preview request failed: %v", err)
	}
	defer previewResp.Body.Close()
	previewBody, _ := io.ReadAll(previewResp.Body)
	if previewResp.StatusCode != http.StatusOK {
		deleteWorkspace(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/delete"), bytes.NewReader(deleteBody))
	if err != nil {
		t.Fatalf("failed to build delete request: %v", err)
	}
	deleteReq.Header.Set("Content-Type", "application/json")
	deleteReq.Header.Set("Authorization", cfg.CliToken)
	deleteReq.Header.Set("Workspace-id", "root")
	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		t.Fatalf("delete request failed: %v", err)
	}
	defer deleteResp.Body.Close()
	deleteRespBody, _ := io.ReadAll(deleteResp.Body)
	if deleteResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK deleting workspace, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace/"+created.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build post-delete get request: %v", err)
	}
	getReq.Header.Set("Authorization", cfg.CliToken)
	getReq.Header.Set("Workspace-id", "root")
	getResp, err := client.Do(getReq)
	if err != nil {
		t.Fatalf("post-delete get request failed: %v", err)
	}
	defer getResp.Body.Close()
	if getResp.StatusCode == http.StatusOK {
		t.Errorf("expected the deleted workspace %s to no longer be gettable", created.UniqueId)
	}
}

// getWorkspace is a thin GET helper, used below to confirm the root workspace (and any
// other workspace caught up in a rejected batch delete) is still actually there.
func getWorkspace(t *testing.T, cfg TestConfig, uniqueId string) (*http.Response, []byte) {
	t.Helper()
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace/"+uniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build get request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("get request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return resp, body
}

// TestWorkspaceAwareDelete_HTTP_RejectsRootWorkspace is the regression guard for
// WorkspaceAwareDeleteAction/WorkspaceAwareDeletePreviewAction's root-workspace guard
// (WorkspaceActions.go) - every other workspace/role/permission in the install
// bootstraps from "root" (see RepairTheWorkspaces), so deleting it would leave the
// whole install unusable. Before that guard existed, this request would have succeeded
// with a 200 OK.
func TestWorkspaceAwareDelete_HTTP_RejectsRootWorkspace(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	client := cfg.NewHTTPClient()

	// The preview step rejects it too, so an admin sees the problem before ever
	// reaching the confirm step, not just when the delete itself is attempted.
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace/delete-preview?uniqueIds=root"), nil)
	if err != nil {
		t.Fatalf("failed to build delete-preview request: %v", err)
	}
	previewReq.Header.Set("Authorization", cfg.CliToken)
	previewReq.Header.Set("Workspace-id", "root")
	previewResp, err := client.Do(previewReq)
	if err != nil {
		t.Fatalf("delete-preview request failed: %v", err)
	}
	defer previewResp.Body.Close()
	if previewResp.StatusCode == http.StatusOK {
		t.Errorf("expected delete-preview of the root workspace to be rejected, got 200 OK")
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {"root"}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/delete"), bytes.NewReader(deleteBody))
	if err != nil {
		t.Fatalf("failed to build delete request: %v", err)
	}
	deleteReq.Header.Set("Content-Type", "application/json")
	deleteReq.Header.Set("Authorization", cfg.CliToken)
	deleteReq.Header.Set("Workspace-id", "root")
	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		t.Fatalf("delete request failed: %v", err)
	}
	defer deleteResp.Body.Close()
	deleteRespBody, _ := io.ReadAll(deleteResp.Body)
	if deleteResp.StatusCode == http.StatusOK {
		t.Fatalf("expected deleting the root workspace to be rejected, got 200 OK: %s", deleteRespBody)
	}

	// Regardless of the rejection above, root must still actually be there.
	getResp, getBody := getWorkspace(t, cfg, "root")
	if getResp.StatusCode != http.StatusOK {
		t.Errorf("expected the root workspace to still exist after the rejected delete, got status %d: %s", getResp.StatusCode, getBody)
	}
}

// TestWorkspaceAwareDelete_HTTP_RejectsBatchContainingRootWorkspace confirms the guard
// rejects the *whole* batch - including a perfectly deletable workspace alongside "root"
// - rather than silently dropping "root" from the list and deleting the rest, which
// would succeed with a 200 OK and leave the caller thinking everything they asked for
// was deleted.
func TestWorkspaceAwareDelete_HTTP_RejectsBatchContainingRootWorkspace(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, wt, role := createSampleWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, created.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId, "root"}})
	client := cfg.NewHTTPClient()
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/delete"), bytes.NewReader(deleteBody))
	if err != nil {
		t.Fatalf("failed to build delete request: %v", err)
	}
	deleteReq.Header.Set("Content-Type", "application/json")
	deleteReq.Header.Set("Authorization", cfg.CliToken)
	deleteReq.Header.Set("Workspace-id", "root")
	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		t.Fatalf("delete request failed: %v", err)
	}
	defer deleteResp.Body.Close()
	deleteRespBody, _ := io.ReadAll(deleteResp.Body)
	if deleteResp.StatusCode == http.StatusOK {
		t.Fatalf("expected a batch delete including the root workspace to be rejected entirely, got 200 OK: %s", deleteRespBody)
	}

	// Neither workspace should have been deleted - not root (never deletable), and not
	// the otherwise-deletable one either (the whole batch was rejected).
	if getResp, getBody := getWorkspace(t, cfg, "root"); getResp.StatusCode != http.StatusOK {
		t.Errorf("expected the root workspace to still exist, got status %d: %s", getResp.StatusCode, getBody)
	}
	if getResp, getBody := getWorkspace(t, cfg, created.UniqueId); getResp.StatusCode != http.StatusOK {
		t.Errorf("expected the sample workspace %s to still exist (whole batch should have been rejected), got status %d: %s", created.UniqueId, getResp.StatusCode, getBody)
	}
}

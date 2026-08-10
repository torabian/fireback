// userWorkspace black-box tests, following capability_test.go's conventions (same
// package, reuses googleResponseEnvelope/googleResponseListEnvelope, and
// workspace_invite_accept_test.go's whoamiResult).
//
// Unlike every other entity in this batch, UserWorkspaceActions.go's security uses
// ResolveStrategy: "user" (not AllowOnRoot's workspace-based one) - QueryEntitiesPointer
// (CrudCoreActions.go) filters ResolveStrategyUser browse/query by `user_id =
// query.UserId`, i.e. the *caller's own* userId, not by workspace at all. So a
// UserWorkspace row attached to some other, newly-created user would never show up in
// root's own /userWorkspace/browse - every test below attaches root's own userId
// (fetched via whoami) instead, same as a real "which workspaces do I belong to" caller
// would see.
package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type userWorkspaceRes struct {
	UniqueId    string `json:"uniqueId"`
	UserId      string `json:"userId"`
	WorkspaceId string `json:"workspaceId"`
}

// rootUserId fetches the currently authenticated caller's own userId via /whoami -
// needed because UserWorkspace/WorkspaceRole browse (ResolveStrategy: "user") only ever
// shows rows belonging to the caller's own userId.
func rootUserId(t *testing.T, cfg TestConfig) string {
	t.Helper()
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/whoami"), nil)
	if err != nil {
		t.Fatalf("failed to build whoami request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
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
		t.Fatalf("expected a non-empty userId from whoami, got none: %+v", out.Data.Item)
	}
	return out.Data.Item.UserId
}

func postUserWorkspace(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/userWorkspace"), bytes.NewReader(body))
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

// createSampleUserWorkspace attaches root's own userId to a freshly-created workspace
// (with its own backing workspaceType/role) and returns everything created, so callers
// can clean it all up.
func createSampleUserWorkspace(t *testing.T, cfg TestConfig) (userWorkspaceRes, workspaceRes, workspaceTypeRes, roleRes) {
	t.Helper()
	ws, wt, role := createSampleWorkspace(t, cfg)
	userId := rootUserId(t, cfg)
	resp, body := postUserWorkspace(t, cfg, map[string]any{
		"userId":      userId,
		"workspaceId": ws.UniqueId,
	})
	if resp.StatusCode != http.StatusOK {
		deleteWorkspace(t, cfg, ws.UniqueId)
		deleteWorkspaceType(t, cfg, wt.UniqueId)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("expected 200 OK creating userWorkspace, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[userWorkspaceRes]
	if err := json.Unmarshal(body, &created); err != nil {
		deleteWorkspace(t, cfg, ws.UniqueId)
		deleteWorkspaceType(t, cfg, wt.UniqueId)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		deleteWorkspace(t, cfg, ws.UniqueId)
		deleteWorkspaceType(t, cfg, wt.UniqueId)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item, ws, wt, role
}

func deleteUserWorkspace(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/userWorkspace/delete"), bytes.NewReader(body))
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

func cleanupSampleUserWorkspace(t *testing.T, cfg TestConfig, uw userWorkspaceRes, ws workspaceRes, wt workspaceTypeRes, role roleRes) {
	deleteUserWorkspace(t, cfg, uw.UniqueId)
	deleteWorkspace(t, cfg, ws.UniqueId)
	deleteWorkspaceType(t, cfg, wt.UniqueId)
	deleteRole(t, cfg, role.UniqueId)
}

func TestUserWorkspaceCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, ws, wt, role := createSampleUserWorkspace(t, cfg)
	defer cleanupSampleUserWorkspace(t, cfg, created, ws, wt, role)

	if created.WorkspaceId != ws.UniqueId {
		t.Errorf("expected workspaceId %q, got %q", ws.UniqueId, created.WorkspaceId)
	}
}

func TestUserWorkspaceBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, ws, wt, role := createSampleUserWorkspace(t, cfg)
	defer cleanupSampleUserWorkspace(t, cfg, created, ws, wt, role)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/userWorkspace/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing userWorkspace, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[userWorkspaceRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created userWorkspace %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestUserWorkspaceGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, ws, wt, role := createSampleUserWorkspace(t, cfg)
	defer cleanupSampleUserWorkspace(t, cfg, created, ws, wt, role)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/userWorkspace/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting userWorkspace, got %d: %s", resp.StatusCode, body)
	}
}

func TestUserWorkspaceUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, ws, wt, role := createSampleUserWorkspace(t, cfg)
	defer cleanupSampleUserWorkspace(t, cfg, created, ws, wt, role)

	// Re-point the same userWorkspace row at a second freshly-created workspace.
	ws2, wt2, role2 := createSampleWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, ws2.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt2.UniqueId)
	defer deleteRole(t, cfg, role2.UniqueId)

	body, _ := json.Marshal(map[string]any{"workspaceId": ws2.UniqueId})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/userWorkspace/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating userWorkspace, got %d: %s", resp.StatusCode, respBody)
	}
	var out googleResponseEnvelope[userWorkspaceRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.WorkspaceId != ws2.UniqueId {
		t.Errorf("expected workspaceId %q after update, got %q", ws2.UniqueId, out.Data.Item.WorkspaceId)
	}
}

func TestUserWorkspaceAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, ws, wt, role := createSampleUserWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/userWorkspace/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteUserWorkspace(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/userWorkspace/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting userWorkspace, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/userWorkspace/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted userWorkspace %s to no longer be gettable", created.UniqueId)
	}
}

// workspaceRole black-box tests, following capability_test.go's conventions (same
// package, reuses googleResponseEnvelope/googleResponseListEnvelope). Unlike most other
// entities in this batch, WorkspaceRoleActions.go's security has neither AllowOnRoot nor
// an explicit ResolveStrategy, so it falls back to the plain workspace_id = query.
// WorkspaceId filter (CrudCoreActions.go's QueryEntitiesPointer) - every write here runs
// in the literal "root" workspace, same as its userWorkspace/role dependencies.
package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type workspaceRoleRes struct {
	UniqueId        string `json:"uniqueId"`
	UserWorkspaceId string `json:"userWorkspaceId"`
	RoleId          string `json:"roleId"`
}

func postWorkspaceRole(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceRole"), bytes.NewReader(body))
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

// sampleWorkspaceRoleChain bundles every dependency createSampleWorkspaceRole below
// creates, so a single deferred cleanup call can tear all of it down in the right order.
type sampleWorkspaceRoleChain struct {
	workspaceRole workspaceRoleRes
	userWorkspace userWorkspaceRes
	role          roleRes
	workspace     workspaceRes
	workspaceType workspaceTypeRes
	backingRole   roleRes
}

func createSampleWorkspaceRole(t *testing.T, cfg TestConfig) sampleWorkspaceRoleChain {
	t.Helper()
	uw, ws, wt, backingRole := createSampleUserWorkspace(t, cfg)
	role := createSampleRole(t, cfg)

	resp, body := postWorkspaceRole(t, cfg, map[string]any{
		"userWorkspaceId": uw.UniqueId,
		"roleId":          role.UniqueId,
	})
	if resp.StatusCode != http.StatusOK {
		cleanupSampleUserWorkspace(t, cfg, uw, ws, wt, backingRole)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("expected 200 OK creating workspaceRole, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[workspaceRoleRes]
	if err := json.Unmarshal(body, &created); err != nil {
		cleanupSampleUserWorkspace(t, cfg, uw, ws, wt, backingRole)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		cleanupSampleUserWorkspace(t, cfg, uw, ws, wt, backingRole)
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return sampleWorkspaceRoleChain{
		workspaceRole: created.Data.Item,
		userWorkspace: uw,
		role:          role,
		workspace:     ws,
		workspaceType: wt,
		backingRole:   backingRole,
	}
}

func deleteWorkspaceRole(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceRole/delete"), bytes.NewReader(body))
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

func (c sampleWorkspaceRoleChain) cleanup(t *testing.T, cfg TestConfig) {
	deleteWorkspaceRole(t, cfg, c.workspaceRole.UniqueId)
	deleteRole(t, cfg, c.role.UniqueId)
	cleanupSampleUserWorkspace(t, cfg, c.userWorkspace, c.workspace, c.workspaceType, c.backingRole)
}

func TestWorkspaceRoleCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	chain := createSampleWorkspaceRole(t, cfg)
	defer chain.cleanup(t, cfg)

	if chain.workspaceRole.RoleId != chain.role.UniqueId {
		t.Errorf("expected roleId %q, got %q", chain.role.UniqueId, chain.workspaceRole.RoleId)
	}
}

// TestWorkspaceRoleBrowse_HTTP_IncludesOwnRecord is also the regression guard for
// WorkspaceRoleCreateAction's workspace-stamping fix (it only ever set UserWorkspaceId/
// RoleId on the created entity, never WorkspaceId, so the row's workspace_id stayed
// null and never matched the caller's own workspace-scoped browse).
func TestWorkspaceRoleBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	chain := createSampleWorkspaceRole(t, cfg)
	defer chain.cleanup(t, cfg)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceRole/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing workspaceRole, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[workspaceRoleRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == chain.workspaceRole.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created workspaceRole %s, got %d items", chain.workspaceRole.UniqueId, len(list.Data.Items))
}

func TestWorkspaceRoleGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	chain := createSampleWorkspaceRole(t, cfg)
	defer chain.cleanup(t, cfg)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceRole/"+chain.workspaceRole.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting workspaceRole, got %d: %s", resp.StatusCode, body)
	}
}

func TestWorkspaceRoleUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	chain := createSampleWorkspaceRole(t, cfg)
	defer chain.cleanup(t, cfg)

	role2 := createSampleRole(t, cfg)
	defer deleteRole(t, cfg, role2.UniqueId)

	body, _ := json.Marshal(map[string]any{"roleId": role2.UniqueId})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/workspaceRole/"+chain.workspaceRole.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating workspaceRole, got %d: %s", resp.StatusCode, respBody)
	}
	var out googleResponseEnvelope[workspaceRoleRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.RoleId != role2.UniqueId {
		t.Errorf("expected roleId %q after update, got %q", role2.UniqueId, out.Data.Item.RoleId)
	}
}

func TestWorkspaceRoleAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	chain := createSampleWorkspaceRole(t, cfg)
	defer deleteRole(t, cfg, chain.role.UniqueId)
	defer cleanupSampleUserWorkspace(t, cfg, chain.userWorkspace, chain.workspace, chain.workspaceType, chain.backingRole)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceRole/delete-preview?uniqueIds="+chain.workspaceRole.UniqueId), nil)
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
		deleteWorkspaceRole(t, cfg, chain.workspaceRole.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {chain.workspaceRole.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceRole/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting workspaceRole, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceRole/"+chain.workspaceRole.UniqueId), nil)
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
		t.Errorf("expected the deleted workspaceRole %s to no longer be gettable", chain.workspaceRole.UniqueId)
	}
}

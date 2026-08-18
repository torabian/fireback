// role Browse/Get/Update/AwareDelete/AwareDeletePreview black-box tests - Create is
// already covered by role_create_test.go (TestRoleCreate_HTTP_Succeeds et al.), whose
// postRole/roleDtoPayload/roleRes/deleteRole helpers this file reuses (same package).
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

func createSampleRole(t *testing.T, cfg TestConfig) roleRes {
	t.Helper()
	resp, body := postRole(t, cfg, roleDtoPayload{
		Name:               fmt.Sprintf("checkendpointtests role crud %d", time.Now().UnixNano()),
		CapabilitiesListId: []string{"root.abac.email-confirmation.query"},
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating role, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	return created.Data.Item
}

func TestRoleBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRole(t, cfg)
	defer deleteRole(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/role/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing role, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[roleRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created role %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestRoleGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRole(t, cfg)
	defer deleteRole(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/role/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting role, got %d: %s", resp.StatusCode, body)
	}
}

func TestRoleUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRole(t, cfg)
	defer deleteRole(t, cfg, created.UniqueId)

	// Unlike Name (emigo.Nullable[string], checked via .Get() so an omitted field
	// leaves the existing value alone), RoleOptionalDto.CapabilitiesListId is a plain
	// complexes.JSON with no "was this key present" tracking - RoleUpdateAction always
	// takes it from the body verbatim (see RoleActions.go), so an update that omits it
	// looks identical to one clearing it, and gets rejected by the same
	// RoleNeedsOneCapability check RoleCreateAction enforces. A real caller (e.g. the
	// manage UI's role edit form) always resubmits the full list on every save for the
	// same reason.
	body, _ := json.Marshal(map[string]any{
		"name":               fmt.Sprintf("checkendpointtests role renamed %d", time.Now().UnixNano()),
		"capabilitiesListId": []string{"root.abac.email-confirmation.query"},
	})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/role/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating role, got %d: %s", resp.StatusCode, respBody)
	}
}

func TestRoleAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRole(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/role/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteRole(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/role/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting role, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/role/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted role %s to no longer be gettable", created.UniqueId)
	}
}

// TestRoleAwareDelete_HTTP_RejectsRootRole is the regression test for
// RoleAwareDeleteAction/RoleAwareDeletePreviewAction's root-role guard
// (RoleActions.go) - mirrors TestWorkspaceAwareDelete_HTTP_RejectsRootWorkspace in
// workspace_test.go for the equivalent root-role case. RoleActions.RemoveEnqueue
// already excluded "root" from the older generic "remove" query path, but
// AwareDeletePreview/AwareDelete go straight to abacdefs.RoleEntityActions,
// bypassing that guard entirely - before RoleAwareDeletePreviewAction/
// RoleAwareDeleteAction's own checks existed, this request would have succeeded
// with a 200 OK and actually deleted the root role.
func TestRoleAwareDelete_HTTP_RejectsRootRole(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	client := cfg.NewHTTPClient()

	// The preview step rejects it too, so an admin sees the problem before ever
	// reaching the confirm step, not just when the delete itself is attempted.
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/role/delete-preview?uniqueIds=root"), nil)
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
		t.Errorf("expected delete-preview of the root role to be rejected, got 200 OK")
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {"root"}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/role/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected deleting the root role to be rejected, got 200 OK: %s", deleteRespBody)
	}

	// The root role must still be there and still gettable - the guard rejected the
	// request outright rather than deleting it and merely returning an error.
	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/role/root"), nil)
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
	if getResp.StatusCode != http.StatusOK {
		t.Errorf("expected the root role to still exist after a rejected delete, got %d", getResp.StatusCode)
	}
}

// TestRoleAwareDelete_HTTP_RejectsBatchContainingRootRole confirms the guard rejects
// the *whole* batch - including a perfectly deletable role alongside "root" - rather
// than silently dropping "root" from the list and deleting the rest, which would
// succeed with a 200 OK and leave the caller thinking everything they asked for was
// deleted. Mirrors TestWorkspaceAwareDelete_HTTP_RejectsBatchContainingRootWorkspace.
func TestRoleAwareDelete_HTTP_RejectsBatchContainingRootRole(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRole(t, cfg)
	defer deleteRole(t, cfg, created.UniqueId)

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId, "root"}})
	client := cfg.NewHTTPClient()
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/role/delete"), bytes.NewReader(deleteBody))
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
	if deleteResp.StatusCode == http.StatusOK {
		t.Fatalf("expected a batch delete including the root role to be rejected entirely, got 200 OK: %s", deleteRespBody)
	}

	// Neither role should have been deleted - not root (never deletable), and not the
	// otherwise-deletable one either (the whole batch was rejected).
	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/role/"+created.UniqueId), nil)
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
	if getResp.StatusCode != http.StatusOK {
		t.Errorf("expected the otherwise-deletable role %s to survive a rejected batch delete, got %d", created.UniqueId, getResp.StatusCode)
	}
}

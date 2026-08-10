// workspaceType black-box tests, following capability_test.go's conventions (same
// package, reuses googleResponseEnvelope/googleResponseListEnvelope). Create/Update/
// AwareDelete/AwareDeletePreview all have AllowOnRoot:true (see WorkspaceTypeActions.go's
// workspaceTypeSecurity), so every write here must run in the literal "root" workspace.
// Create also requires an existing, root-owned role (roleId is validate:"required" and
// checked for existence/ownership by ValidateRoleAndItsExistence) - reuses
// role_crud_test.go's createSampleRole/deleteRole helpers (same package) for that.
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

type workspaceTypeRes struct {
	UniqueId string `json:"uniqueId"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	RoleId   string `json:"roleId"`
}

func postWorkspaceType(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceType"), bytes.NewReader(body))
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

// createSampleWorkspaceType creates its own backing role (root.abac.email-confirmation.
// query - same capability role_crud_test.go's createSampleRole already grants itself, so
// no extra ValidateRoleAndItsExistence-driven capability checks come into play here) and
// returns both so callers can clean up either independently.
func createSampleWorkspaceType(t *testing.T, cfg TestConfig) (workspaceTypeRes, roleRes) {
	t.Helper()
	role := createSampleRole(t, cfg)
	// ValidateTheWorkspaceTypeEntity/ValidateWorkspaceTypeSlug restricts the slug to
	// lowercase a-z and dashes only, no digits - lettersFromInt (workspace_invite_
	// accept_test.go, same package) base-26-encodes the timestamp into letters instead.
	slug := "/checkendpointtests-wt-" + lettersFromInt(time.Now().UnixNano())
	resp, body := postWorkspaceType(t, cfg, map[string]any{
		"title":  "checkendpointtests workspace type",
		"slug":   slug,
		"roleId": role.UniqueId,
	})
	if resp.StatusCode != http.StatusOK {
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("expected 200 OK creating workspaceType, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[workspaceTypeRes]
	if err := json.Unmarshal(body, &created); err != nil {
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		deleteRole(t, cfg, role.UniqueId)
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item, role
}

func deleteWorkspaceType(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceType/delete"), bytes.NewReader(body))
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

func TestWorkspaceTypeCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, role := createSampleWorkspaceType(t, cfg)
	defer deleteWorkspaceType(t, cfg, created.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	if created.RoleId != role.UniqueId {
		t.Errorf("expected roleId %q, got %q", role.UniqueId, created.RoleId)
	}
}

// TestWorkspaceTypeCreate_HTTP_RejectsMissingRole covers WorkspaceTypeActions.go's
// ValidateRoleAndItsExistence - roleId is validate:"required" and must reference an
// existing, root-owned role.
func TestWorkspaceTypeCreate_HTTP_RejectsMissingRole(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postWorkspaceType(t, cfg, map[string]any{
		"title": "checkendpointtests no role",
		"slug":  fmt.Sprintf("/checkendpointtests-norole-%d", time.Now().UnixNano()),
	})
	if resp.StatusCode == http.StatusOK {
		var out googleResponseEnvelope[workspaceTypeRes]
		if err := json.Unmarshal(body, &out); err == nil {
			deleteWorkspaceType(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected creation with no roleId to be rejected, got 200 OK: %s", body)
	}
}

func TestWorkspaceTypeBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, role := createSampleWorkspaceType(t, cfg)
	defer deleteWorkspaceType(t, cfg, created.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceType/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing workspaceType, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[workspaceTypeRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created workspaceType %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestWorkspaceTypeGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, role := createSampleWorkspaceType(t, cfg)
	defer deleteWorkspaceType(t, cfg, created.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceType/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting workspaceType, got %d: %s", resp.StatusCode, body)
	}
}

func TestWorkspaceTypeUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, role := createSampleWorkspaceType(t, cfg)
	defer deleteWorkspaceType(t, cfg, created.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	body, _ := json.Marshal(map[string]any{"title": "checkendpointtests renamed"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/workspaceType/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating workspaceType, got %d: %s", resp.StatusCode, respBody)
	}
	var out googleResponseEnvelope[workspaceTypeRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.Title != "checkendpointtests renamed" {
		t.Errorf("expected title %q after update, got %q", "checkendpointtests renamed", out.Data.Item.Title)
	}
}

func TestWorkspaceTypeAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, role := createSampleWorkspaceType(t, cfg)
	defer deleteRole(t, cfg, role.UniqueId)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceType/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteWorkspaceType(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceType/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting workspaceType, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceType/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted workspaceType %s to no longer be gettable", created.UniqueId)
	}
}

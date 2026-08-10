// publicJoinKey black-box tests, following capability_test.go's conventions (same
// package, reuses googleResponseEnvelope/googleResponseListEnvelope, plus
// postRole/roleDtoPayload/deleteRole from role_create_test.go for the roleId fixture
// every publicJoinKey needs).
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

type publicJoinKeyRes struct {
	UniqueId string `json:"uniqueId"`
	RoleId   string `json:"roleId"`
}

func postPublicJoinKey(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/publicJoinKey"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

// createSamplePublicJoinKey creates a throwaway role first (publicJoinKey's only
// meaningful field), returning both records so the caller can clean up either.
func createSamplePublicJoinKey(t *testing.T, cfg TestConfig) (publicJoinKeyRes, string) {
	t.Helper()

	roleResp, roleBody := postRole(t, cfg, roleDtoPayload{
		Name:               fmt.Sprintf("checkendpointtests public-join-key role %d", time.Now().UnixNano()),
		CapabilitiesListId: []string{"root.abac.email-confirmation.query"},
	})
	if roleResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating the role fixture, got %d: %s", roleResp.StatusCode, roleBody)
	}
	var role googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(roleBody, &role); err != nil {
		t.Fatalf("failed to decode role create response: %v\nbody: %s", err, roleBody)
	}

	// Like publicAuthentication, PublicJoinKeyCreateAction trusts workspaceId from the
	// request body verbatim rather than deriving it from the resolved auth context.
	resp, body := postPublicJoinKey(t, cfg, map[string]any{
		"roleId":      role.Data.Item.UniqueId,
		"workspaceId": cfg.WorkspaceID,
	})
	if resp.StatusCode != http.StatusOK {
		deleteRole(t, cfg, role.Data.Item.UniqueId)
		t.Fatalf("expected 200 OK creating publicJoinKey, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[publicJoinKeyRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item, role.Data.Item.UniqueId
}

func deletePublicJoinKey(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/publicJoinKey/delete"), bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func TestPublicJoinKeyCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, roleId := createSamplePublicJoinKey(t, cfg)
	defer deletePublicJoinKey(t, cfg, created.UniqueId)
	defer deleteRole(t, cfg, roleId)

	if created.RoleId != roleId {
		t.Errorf("expected roleId %q, got %q", roleId, created.RoleId)
	}
}

func TestPublicJoinKeyBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, roleId := createSamplePublicJoinKey(t, cfg)
	defer deletePublicJoinKey(t, cfg, created.UniqueId)
	defer deleteRole(t, cfg, roleId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/publicJoinKey/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing publicJoinKey, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[publicJoinKeyRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created publicJoinKey %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestPublicJoinKeyGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, roleId := createSamplePublicJoinKey(t, cfg)
	defer deletePublicJoinKey(t, cfg, created.UniqueId)
	defer deleteRole(t, cfg, roleId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/publicJoinKey/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting publicJoinKey, got %d: %s", resp.StatusCode, body)
	}
}

func TestPublicJoinKeyUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, roleId := createSamplePublicJoinKey(t, cfg)
	defer deletePublicJoinKey(t, cfg, created.UniqueId)
	defer deleteRole(t, cfg, roleId)

	secondRoleResp, secondRoleBody := postRole(t, cfg, roleDtoPayload{
		Name:               fmt.Sprintf("checkendpointtests public-join-key second role %d", time.Now().UnixNano()),
		CapabilitiesListId: []string{"root.abac.email-confirmation.query"},
	})
	if secondRoleResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating the second role fixture, got %d: %s", secondRoleResp.StatusCode, secondRoleBody)
	}
	var secondRole googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(secondRoleBody, &secondRole); err != nil {
		t.Fatalf("failed to decode second role create response: %v\nbody: %s", err, secondRoleBody)
	}
	defer deleteRole(t, cfg, secondRole.Data.Item.UniqueId)

	body, _ := json.Marshal(map[string]any{"roleId": secondRole.Data.Item.UniqueId})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/publicJoinKey/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating publicJoinKey, got %d: %s", resp.StatusCode, respBody)
	}
	var out googleResponseEnvelope[publicJoinKeyRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.RoleId != secondRole.Data.Item.UniqueId {
		t.Errorf("expected roleId %q after update, got %q", secondRole.Data.Item.UniqueId, out.Data.Item.RoleId)
	}
}

func TestPublicJoinKeyAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created, roleId := createSamplePublicJoinKey(t, cfg)
	defer deleteRole(t, cfg, roleId)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/publicJoinKey/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deletePublicJoinKey(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/publicJoinKey/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting publicJoinKey, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/publicJoinKey/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted publicJoinKey %s to no longer be gettable", created.UniqueId)
	}
}

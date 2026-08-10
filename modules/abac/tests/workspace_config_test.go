// workspaceConfig black-box tests, following capability_test.go's conventions (same
// package, reuses googleResponseEnvelope/googleResponseListEnvelope). Every action has
// AllowOnRoot:true with ResolveStrategy "workspace" (see WorkspaceConfigActions.go), so
// every write here must run in the literal "root" workspace.
//
// The two /workspace-config/distinct endpoints always operate on the single, real
// workspace_id='root' row shared by the whole running server (not a row this test can
// create/delete of its own) - TestWorkspaceConfigDistinctGet/Update below capture
// whatever value is already there before touching it and restore it afterwards, so they
// don't leave global side effects for the rest of the suite.
package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type workspaceConfigRes struct {
	UniqueId   string `json:"uniqueId"`
	EnableTotp *bool  `json:"enableTotp"`
}

func postWorkspaceConfig(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceConfig"), bytes.NewReader(body))
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

func createSampleWorkspaceConfig(t *testing.T, cfg TestConfig) workspaceConfigRes {
	t.Helper()
	resp, body := postWorkspaceConfig(t, cfg, map[string]any{"enableTotp": true})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating workspaceConfig, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[workspaceConfigRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deleteWorkspaceConfig(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceConfig/delete"), bytes.NewReader(body))
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

func TestWorkspaceConfigCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWorkspaceConfig(t, cfg)
	defer deleteWorkspaceConfig(t, cfg, created.UniqueId)
	if created.EnableTotp == nil || !*created.EnableTotp {
		t.Errorf("expected enableTotp true, got %+v", created.EnableTotp)
	}
}

func TestWorkspaceConfigBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWorkspaceConfig(t, cfg)
	defer deleteWorkspaceConfig(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceConfig/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing workspaceConfig, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[workspaceConfigRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created workspaceConfig %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestWorkspaceConfigGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWorkspaceConfig(t, cfg)
	defer deleteWorkspaceConfig(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceConfig/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting workspaceConfig, got %d: %s", resp.StatusCode, body)
	}
}

// TestWorkspaceConfigUpdate_HTTP_Succeeds covers PATCH /workspaceConfig/:uniqueId, which
// is actually workspaceConfigUpsertByWorkspace under the hood - it finds-or-creates by
// the caller's *current* workspace (root here), ignoring the uniqueId path param beyond
// routing. Since that means it always targets the very same real "root" row the
// /distinct endpoints below use, this captures and restores enableTotp too.
func TestWorkspaceConfigUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	client := cfg.NewHTTPClient()
	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace-config/distinct"), nil)
	if err != nil {
		t.Fatalf("failed to build distinct-get request: %v", err)
	}
	getReq.Header.Set("Authorization", cfg.CliToken)
	getReq.Header.Set("Workspace-id", "root")
	getResp, err := client.Do(getReq)
	if err != nil {
		t.Fatalf("distinct-get request failed: %v", err)
	}
	defer getResp.Body.Close()
	getBody, _ := io.ReadAll(getResp.Body)
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK on distinct-get, got %d: %s", getResp.StatusCode, getBody)
	}
	var original googleResponseEnvelope[workspaceConfigRes]
	if err := json.Unmarshal(getBody, &original); err != nil {
		t.Fatalf("failed to decode distinct-get response: %v\nbody: %s", err, getBody)
	}
	originalEnableTotp := false
	if original.Data.Item.EnableTotp != nil {
		originalEnableTotp = *original.Data.Item.EnableTotp
	}
	defer func() {
		restoreBody, _ := json.Marshal(map[string]any{"enableTotp": originalEnableTotp})
		restoreReq, _ := http.NewRequest(http.MethodPatch, cfg.URL("/workspaceConfig/root"), bytes.NewReader(restoreBody))
		restoreReq.Header.Set("Content-Type", "application/json")
		restoreReq.Header.Set("Authorization", cfg.CliToken)
		restoreReq.Header.Set("Workspace-id", "root")
		if restoreResp, err := client.Do(restoreReq); err == nil {
			restoreResp.Body.Close()
		}
	}()

	toggled := !originalEnableTotp
	body, _ := json.Marshal(map[string]any{"enableTotp": toggled})
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/workspaceConfig/root"), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating workspaceConfig, got %d: %s", resp.StatusCode, respBody)
	}
	var out googleResponseEnvelope[workspaceConfigRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.EnableTotp == nil || *out.Data.Item.EnableTotp != toggled {
		t.Errorf("expected enableTotp %v after update, got %+v", toggled, out.Data.Item.EnableTotp)
	}
}

func TestWorkspaceConfigAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleWorkspaceConfig(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceConfig/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteWorkspaceConfig(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaceConfig/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting workspaceConfig, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq2, err := http.NewRequest(http.MethodGet, cfg.URL("/workspaceConfig/"+created.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build post-delete get request: %v", err)
	}
	getReq2.Header.Set("Authorization", cfg.CliToken)
	getReq2.Header.Set("Workspace-id", "root")
	getResp2, err := client.Do(getReq2)
	if err != nil {
		t.Fatalf("post-delete get request failed: %v", err)
	}
	defer getResp2.Body.Close()
	if getResp2.StatusCode == http.StatusOK {
		t.Errorf("expected the deleted workspaceConfig %s to no longer be gettable", created.UniqueId)
	}
}

// TestWorkspaceConfigDistinct_HTTP_GetAndUpdate covers the two hand-declared
// /workspace-config/distinct actions together, since Update's find-or-create and Get's
// read need to agree on the same real root row - restores enableTotp to whatever it was
// before this test ran, same reasoning as TestWorkspaceConfigUpdate_HTTP_Succeeds above.
func TestWorkspaceConfigDistinct_HTTP_GetAndUpdate(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	client := cfg.NewHTTPClient()

	doGet := func() workspaceConfigRes {
		req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace-config/distinct"), nil)
		if err != nil {
			t.Fatalf("failed to build distinct-get request: %v", err)
		}
		req.Header.Set("Authorization", cfg.CliToken)
		req.Header.Set("Workspace-id", "root")
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("distinct-get request failed: %v", err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200 OK on distinct-get, got %d: %s", resp.StatusCode, body)
		}
		var out googleResponseEnvelope[workspaceConfigRes]
		if err := json.Unmarshal(body, &out); err != nil {
			t.Fatalf("failed to decode distinct-get response: %v\nbody: %s", err, body)
		}
		return out.Data.Item
	}

	doUpdate := func(enableTotp bool) workspaceConfigRes {
		body, _ := json.Marshal(map[string]any{"enableTotp": enableTotp})
		req, err := http.NewRequest(http.MethodPatch, cfg.URL("/workspace-config/distinct"), bytes.NewReader(body))
		if err != nil {
			t.Fatalf("failed to build distinct-update request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", cfg.CliToken)
		req.Header.Set("Workspace-id", "root")
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("distinct-update request failed: %v", err)
		}
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200 OK on distinct-update, got %d: %s", resp.StatusCode, respBody)
		}
		var out googleResponseEnvelope[workspaceConfigRes]
		if err := json.Unmarshal(respBody, &out); err != nil {
			t.Fatalf("failed to decode distinct-update response: %v\nbody: %s", err, respBody)
		}
		return out.Data.Item
	}

	original := doGet()
	originalEnableTotp := false
	if original.EnableTotp != nil {
		originalEnableTotp = *original.EnableTotp
	}
	defer doUpdate(originalEnableTotp)

	toggled := !originalEnableTotp
	updated := doUpdate(toggled)
	if updated.EnableTotp == nil || *updated.EnableTotp != toggled {
		t.Errorf("expected enableTotp %v after distinct-update, got %+v", toggled, updated.EnableTotp)
	}

	reGet := doGet()
	if reGet.EnableTotp == nil || *reGet.EnableTotp != toggled {
		t.Errorf("expected distinct-get to reflect the update (enableTotp %v), got %+v", toggled, reGet.EnableTotp)
	}
}

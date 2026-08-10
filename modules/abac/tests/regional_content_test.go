// regionalContent black-box tests, following capability_test.go's conventions (same
// package, reuses googleResponseEnvelope/googleResponseListEnvelope). Create/Update/
// AwareDelete/AwareDeletePreview all have AllowOnRoot:true, so every write here must run
// in the literal "root" workspace.
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

type regionalContentRes struct {
	UniqueId string `json:"uniqueId"`
	Content  string `json:"content"`
}

func postRegionalContent(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/regionalContent"), bytes.NewReader(body))
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

func createSampleRegionalContent(t *testing.T, cfg TestConfig) regionalContentRes {
	t.Helper()
	resp, body := postRegionalContent(t, cfg, map[string]any{
		"content":    "checkendpointtests content",
		"region":     "any",
		"languageId": "en",
		"keyGroup":   fmt.Sprintf("CHECKENDPOINTTESTS_%d", time.Now().UnixNano()),
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating regionalContent, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[regionalContentRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deleteRegionalContent(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/regionalContent/delete"), bytes.NewReader(body))
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

func TestRegionalContentCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRegionalContent(t, cfg)
	defer deleteRegionalContent(t, cfg, created.UniqueId)
}

// TestRegionalContentCreate_HTTP_ValidationRequiredFields covers the fix making
// RegionalContentCreateAction actually call fireback.CommonStructValidatorPointer -
// content/region/languageId/keyGroup are validate:"required" but were previously
// accepted empty.
func TestRegionalContentCreate_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postRegionalContent(t, cfg, map[string]any{"region": "any"})
	if resp.StatusCode == http.StatusOK {
		var out googleResponseEnvelope[regionalContentRes]
		if err := json.Unmarshal(body, &out); err == nil {
			deleteRegionalContent(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected creation with missing required fields to be rejected, got 200 OK: %s", body)
	}
}

// TestRegionalContentBrowse_HTTP_IncludesOwnRecord is also the regression guard for
// RegionalContentCreateAction's workspace-stamping fix.
func TestRegionalContentBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRegionalContent(t, cfg)
	defer deleteRegionalContent(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/regionalContent/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing regionalContent, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[regionalContentRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created regionalContent %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestRegionalContentGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRegionalContent(t, cfg)
	defer deleteRegionalContent(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/regionalContent/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting regionalContent, got %d: %s", resp.StatusCode, body)
	}
}

func TestRegionalContentUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRegionalContent(t, cfg)
	defer deleteRegionalContent(t, cfg, created.UniqueId)

	body, _ := json.Marshal(map[string]any{"content": "checkendpointtests updated content"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/regionalContent/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating regionalContent, got %d: %s", resp.StatusCode, respBody)
	}
	var out googleResponseEnvelope[regionalContentRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.Content != "checkendpointtests updated content" {
		t.Errorf("expected content %q after update, got %q", "checkendpointtests updated content", out.Data.Item.Content)
	}
}

func TestRegionalContentAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleRegionalContent(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/regionalContent/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteRegionalContent(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/regionalContent/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting regionalContent, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/regionalContent/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted regionalContent %s to no longer be gettable", created.UniqueId)
	}
}

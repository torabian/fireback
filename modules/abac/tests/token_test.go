// token black-box tests, following capability_test.go's conventions (same package,
// reuses googleResponseEnvelope/googleResponseListEnvelope). Create/Update/AwareDelete/
// AwareDeletePreview all have AllowOnRoot:true, so every write here must run in the
// literal "root" workspace.
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

type tokenRes struct {
	UniqueId string `json:"uniqueId"`
	Token    string `json:"token"`
}

func postToken(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/token"), bytes.NewReader(body))
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

func createSampleToken(t *testing.T, cfg TestConfig) tokenRes {
	t.Helper()
	resp, body := postToken(t, cfg, map[string]any{
		"token": fmt.Sprintf("checkendpointtests-token-%d", time.Now().UnixNano()),
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating token, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[tokenRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deleteToken(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/token/delete"), bytes.NewReader(body))
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

func TestTokenCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleToken(t, cfg)
	defer deleteToken(t, cfg, created.UniqueId)
}

// TestTokenBrowse_HTTP_IncludesOwnRecord is also the regression guard for
// TokenCreateAction's workspace-stamping fix.
func TestTokenBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleToken(t, cfg)
	defer deleteToken(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/token/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing token, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[tokenRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created token %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestTokenGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleToken(t, cfg)
	defer deleteToken(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/token/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting token, got %d: %s", resp.StatusCode, body)
	}
}

func TestTokenUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleToken(t, cfg)
	defer deleteToken(t, cfg, created.UniqueId)

	newToken := fmt.Sprintf("checkendpointtests-token-updated-%d", time.Now().UnixNano())
	body, _ := json.Marshal(map[string]any{"token": newToken})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/token/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating token, got %d: %s", resp.StatusCode, respBody)
	}
	var out googleResponseEnvelope[tokenRes]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode update response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.Token != newToken {
		t.Errorf("expected token %q after update, got %q", newToken, out.Data.Item.Token)
	}
}

func TestTokenAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleToken(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/token/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteToken(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/token/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting token, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/token/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted token %s to no longer be gettable", created.UniqueId)
	}
}

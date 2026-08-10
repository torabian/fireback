// passportMethod black-box tests, following capability_test.go's conventions (same
// package, reuses googleResponseEnvelope/googleResponseListEnvelope). Every action
// requires the literal "root" workspace (see passportMethodSecurity's AllowOnRoot).
// Uses "google"/"facebook" (not "email") to avoid the type+region duplicate-check
// colliding with ensureEmailPassportMethod's "email"/"global" row, which many other
// tests in this package (and the e2e specs) rely on already existing.
package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type passportMethodRes struct {
	UniqueId string `json:"uniqueId"`
	Type     string `json:"type"`
	Region   string `json:"region"`
}

func postPassportMethod(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passportMethod"), bytes.NewReader(body))
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

func deletePassportMethod(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passportMethod/delete"), bytes.NewReader(body))
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

func TestPassportMethodCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postPassportMethod(t, cfg, map[string]any{"type": "google", "region": "global"})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating passportMethod, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[passportMethodRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	defer deletePassportMethod(t, cfg, created.Data.Item.UniqueId)
	if created.Data.Item.Type != "google" {
		t.Errorf("expected type %q, got %q", "google", created.Data.Item.Type)
	}
}

func TestPassportMethodCreate_HTTP_RejectsDuplicateTypeRegion(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	first, firstBody := postPassportMethod(t, cfg, map[string]any{"type": "facebook", "region": "global"})
	if first.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating the first facebook/global passportMethod, got %d: %s", first.StatusCode, firstBody)
	}
	var created googleResponseEnvelope[passportMethodRes]
	if err := json.Unmarshal(firstBody, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, firstBody)
	}
	defer deletePassportMethod(t, cfg, created.Data.Item.UniqueId)

	second, secondBody := postPassportMethod(t, cfg, map[string]any{"type": "facebook", "region": "global"})
	if second.StatusCode == http.StatusOK {
		var dup googleResponseEnvelope[passportMethodRes]
		if err := json.Unmarshal(secondBody, &dup); err == nil {
			deletePassportMethod(t, cfg, dup.Data.Item.UniqueId)
		}
		t.Fatalf("expected a duplicate type+region passportMethod to be rejected, got 200 OK: %s", secondBody)
	}
}

func TestPassportMethodCreate_HTTP_ValidationRejectsUnknownType(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postPassportMethod(t, cfg, map[string]any{"type": "not-a-real-type", "region": "global"})
	if resp.StatusCode == http.StatusOK {
		var out googleResponseEnvelope[passportMethodRes]
		if err := json.Unmarshal(body, &out); err == nil {
			deletePassportMethod(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected an unknown type to be rejected, got 200 OK: %s", body)
	}
}

func TestPassportMethodBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postPassportMethod(t, cfg, map[string]any{"type": "google", "region": "global"})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating passportMethod, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[passportMethodRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	defer deletePassportMethod(t, cfg, created.Data.Item.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/passportMethod/browse?itemsPerPage=1000"), nil)
	if err != nil {
		t.Fatalf("failed to build browse request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	browseResp, err := client.Do(req)
	if err != nil {
		t.Fatalf("browse request failed: %v", err)
	}
	defer browseResp.Body.Close()
	browseBody, _ := io.ReadAll(browseResp.Body)
	if browseResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK browsing passportMethod, got %d: %s", browseResp.StatusCode, browseBody)
	}
	var list googleResponseListEnvelope[passportMethodRes]
	if err := json.Unmarshal(browseBody, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, browseBody)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.Data.Item.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created passportMethod %s, got %d items", created.Data.Item.UniqueId, len(list.Data.Items))
}

func TestPassportMethodGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postPassportMethod(t, cfg, map[string]any{"type": "google", "region": "global"})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating passportMethod, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[passportMethodRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	defer deletePassportMethod(t, cfg, created.Data.Item.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/passportMethod/"+created.Data.Item.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build get request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	getResp, err := client.Do(req)
	if err != nil {
		t.Fatalf("get request failed: %v", err)
	}
	defer getResp.Body.Close()
	getBody, _ := io.ReadAll(getResp.Body)
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK getting passportMethod, got %d: %s", getResp.StatusCode, getBody)
	}
}

func TestPassportMethodUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postPassportMethod(t, cfg, map[string]any{"type": "google", "region": "global"})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating passportMethod, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[passportMethodRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	defer deletePassportMethod(t, cfg, created.Data.Item.UniqueId)

	updateBody, _ := json.Marshal(map[string]any{"clientKey": "checkendpointtests-client-key"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/passportMethod/"+created.Data.Item.UniqueId), bytes.NewReader(updateBody))
	if err != nil {
		t.Fatalf("failed to build update request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", "root")
	updateResp, err := client.Do(req)
	if err != nil {
		t.Fatalf("update request failed: %v", err)
	}
	defer updateResp.Body.Close()
	updateRespBody, _ := io.ReadAll(updateResp.Body)
	if updateResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK updating passportMethod, got %d: %s", updateResp.StatusCode, updateRespBody)
	}
}

func TestPassportMethodAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postPassportMethod(t, cfg, map[string]any{"type": "google", "region": "global"})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating passportMethod, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[passportMethodRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/passportMethod/delete-preview?uniqueIds="+created.Data.Item.UniqueId), nil)
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
		deletePassportMethod(t, cfg, created.Data.Item.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.Data.Item.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/passportMethod/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting passportMethod, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/passportMethod/"+created.Data.Item.UniqueId), nil)
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
		t.Errorf("expected the deleted passportMethod %s to no longer be gettable", created.Data.Item.UniqueId)
	}
}

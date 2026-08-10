// passport black-box tests, following capability_test.go's conventions (same package,
// reuses googleResponseEnvelope/googleResponseListEnvelope). Create/Update/AwareDelete/
// AwareDeletePreview all have AllowOnRoot:true (see PassportActions.go), so every write
// here must run in the literal "root" workspace - matches "Passport and user always
// belong to the root workspace" elsewhere in this package (UnsafeGenerateUser).
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

type passportRes struct {
	UniqueId string `json:"uniqueId"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

func postPassport(t *testing.T, cfg TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passport"), bytes.NewReader(body))
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

func createSamplePassport(t *testing.T, cfg TestConfig) passportRes {
	t.Helper()
	value := fmt.Sprintf("checkendpointtests-passport-%d@example.com", time.Now().UnixNano())
	resp, body := postPassport(t, cfg, map[string]any{"type": "email", "value": value})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating passport, got %d: %s", resp.StatusCode, body)
	}
	var created googleResponseEnvelope[passportRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deletePassport(t *testing.T, cfg TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}
	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/delete"), bytes.NewReader(body))
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

func TestPassportCreate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSamplePassport(t, cfg)
	defer deletePassport(t, cfg, created.UniqueId)
	if created.Type != "email" {
		t.Errorf("expected type %q, got %q", "email", created.Type)
	}
}

// TestPassportCreate_HTTP_ValidationRequiredFields covers the fix making
// PassportCreateAction actually call fireback.CommonStructValidatorPointer - type/value
// are validate:"required" (see PassportDto.go) but were previously accepted empty.
func TestPassportCreate_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postPassport(t, cfg, map[string]any{"type": "email"})
	if resp.StatusCode == http.StatusOK {
		var out googleResponseEnvelope[passportRes]
		if err := json.Unmarshal(body, &out); err == nil {
			deletePassport(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected creation with a missing value to be rejected, got 200 OK: %s", body)
	}
}

// TestPassportBrowse_HTTP_IncludesOwnRecord is also the regression guard for
// PassportCreateAction's workspace-stamping fix - before it, every created row's
// workspaceId stayed null, making it invisible to this same workspace-scoped browse.
func TestPassportBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSamplePassport(t, cfg)
	defer deletePassport(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/passport/browse?itemsPerPage=1000"), nil)
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
		t.Fatalf("expected 200 OK browsing passport, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[passportRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created passport %s, got %d items", created.UniqueId, len(list.Data.Items))
}

func TestPassportGet_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSamplePassport(t, cfg)
	defer deletePassport(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/passport/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting passport, got %d: %s", resp.StatusCode, body)
	}
	var out googleResponseEnvelope[passportRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode get response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UniqueId != created.UniqueId {
		t.Errorf("expected uniqueId %q, got %q", created.UniqueId, out.Data.Item.UniqueId)
	}
}

func TestPassportUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSamplePassport(t, cfg)
	defer deletePassport(t, cfg, created.UniqueId)

	body, _ := json.Marshal(map[string]any{"thirdPartyVerifier": "checkendpointtests-verifier"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/passport/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating passport, got %d: %s", resp.StatusCode, respBody)
	}
}

func TestPassportAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSamplePassport(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/passport/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deletePassport(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/passport/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting passport, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/passport/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted passport %s to no longer be gettable", created.UniqueId)
	}
}

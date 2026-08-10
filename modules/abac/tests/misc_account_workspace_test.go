// Misc standalone account/workspace action black-box tests: CreateWorkspace,
// QueryUserRoleWorkspaces, QueryWorkspaceTypesPublicly, UserInvitations, UserPassports,
// GsmSendSms. Follows this package's established conventions; reuses freshSignupTarget
// (core_session_test.go), googleResponseEnvelope/googleResponseListEnvelope
// (workspace_invite_accept_test.go / passport_methods_http_test.go).
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

func TestCreateWorkspace_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	_, token, _ := freshSignupTarget(t, cfg, "create-workspace")

	body, _ := json.Marshal(map[string]any{"name": "checkendpointtests new workspace"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaces/create"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build workspaces/create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("workspaces/create request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating a workspace, got %d: %s", resp.StatusCode, respBody)
	}

	var out googleResponseEnvelope[struct {
		WorkspaceId string `json:"workspaceId"`
	}]
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, respBody)
	}
	newWorkspaceId := out.Data.Item.WorkspaceId
	if newWorkspaceId == "" {
		t.Fatalf("expected a non-empty workspaceId, got none: %s", respBody)
	}

	// Confirm the creator was actually linked to it (see
	// CreateWorkspaceAndAssignUser/QueryUserRoleWorkspacesAction).
	urwReq, err := http.NewRequest(http.MethodGet, cfg.URL("/urw/query"), nil)
	if err != nil {
		t.Fatalf("failed to build urw/query request: %v", err)
	}
	urwReq.Header.Set("Authorization", token)
	urwResp, err := client.Do(urwReq)
	if err != nil {
		t.Fatalf("urw/query request failed: %v", err)
	}
	defer urwResp.Body.Close()
	urwBody, _ := io.ReadAll(urwResp.Body)
	if urwResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK from urw/query, got %d: %s", urwResp.StatusCode, urwBody)
	}
	var urwList googleResponseListEnvelope[struct {
		UniqueId string `json:"uniqueId"`
	}]
	if err := json.Unmarshal(urwBody, &urwList); err != nil {
		t.Fatalf("failed to decode urw/query response: %v\nbody: %s", err, urwBody)
	}
	found := false
	for _, w := range urwList.Data.Items {
		if w.UniqueId == newWorkspaceId {
			found = true
		}
	}
	if !found {
		t.Errorf("expected the creator's own urw/query to include the just-created workspace %s, got %+v", newWorkspaceId, urwList.Data.Items)
	}
}

func TestCreateWorkspace_HTTP_RequiresAuth(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	body, _ := json.Marshal(map[string]any{"name": "checkendpointtests anon workspace"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspaces/create"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build workspaces/create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("workspaces/create request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unauthenticated workspaces/create to be rejected, got 200 OK")
	}
}

func TestQueryUserRoleWorkspaces_HTTP_ReturnsOwnWorkspace(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	_, token, workspaceId := freshSignupTarget(t, cfg, "urw-query")

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/urw/query"), nil)
	if err != nil {
		t.Fatalf("failed to build urw/query request: %v", err)
	}
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("urw/query request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK from urw/query, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[struct {
		UniqueId string `json:"uniqueId"`
	}]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode urw/query response: %v\nbody: %s", err, body)
	}
	found := false
	for _, w := range list.Data.Items {
		if w.UniqueId == workspaceId {
			found = true
		}
	}
	if !found {
		t.Errorf("expected urw/query to include the signup's own workspace %s, got %+v", workspaceId, list.Data.Items)
	}
}

func TestQueryWorkspaceTypesPublicly_HTTP_ExcludesRootIncludesCustom(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	customId := createRoleAndWorkspaceType(t, cfg, fmt.Sprintf("checkendpointtests public-types %d", time.Now().UnixNano()), []string{"root.abac.email-confirmation.query"})

	client := cfg.NewHTTPClient()
	// No auth header at all - see QueryWorkspaceTypesPubliclyActionImplementation.go's
	// nil SecurityModel: "so the signup screen can get it" before the caller has any
	// session.
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/workspace/public/types?itemsPerPage=1000"), nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK from workspace/public/types with no auth, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[struct {
		UniqueId string `json:"uniqueId"`
	}]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, body)
	}

	foundCustom, foundRoot := false, false
	for _, wt := range list.Data.Items {
		if wt.UniqueId == customId {
			foundCustom = true
		}
		if wt.UniqueId == "root" {
			foundRoot = true
		}
	}
	if !foundCustom {
		t.Errorf("expected the public list to include the just-created workspace type %s, got %+v", customId, list.Data.Items)
	}
	if foundRoot {
		t.Errorf("expected the public list to never include the seeded \"root\" workspace type, got %+v", list.Data.Items)
	}
}

func TestUserInvitations_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	_, token, _ := freshSignupTarget(t, cfg, "user-invitations")

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/users/invitations"), nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK from users/invitations (even with none pending), got %d: %s", resp.StatusCode, body)
	}
}

func TestUserInvitations_HTTP_RequiresAuth(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/users/invitations"), nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unauthenticated users/invitations to be rejected, got 200 OK")
	}
}

func TestUserPassports_HTTP_ReturnsOwnPassport(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	email, token, _ := freshSignupTarget(t, cfg, "user-passports")

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/user/passports"), nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK from user/passports, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	}]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, body)
	}
	found := false
	for _, p := range list.Data.Items {
		if p.Value == email && p.Type == "email" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected user/passports to include the signed-up email %q, got %+v", email, list.Data.Items)
	}
}

func TestGsmSendSms_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	// No NotificationConfig.generalGsmProviderId configured in this test's DB state -
	// GsmSendSMSUsingNotificationConfig gracefully falls back to a "print-to-terminal"
	// queue instead of erroring (see GsmProviderActions.go), so this needs no fixture
	// setup at all.
	body, _ := json.Marshal(map[string]any{"toNumber": "+15550000000", "body": "checkendpointtests sms"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/gsm/send/sms"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK sending an sms (falls back to a terminal queue with no gsm provider configured), got %d: %s", resp.StatusCode, respBody)
	}
}

func TestGsmSendSms_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	for _, tc := range []struct {
		name    string
		payload map[string]any
	}{
		{"missing toNumber", map[string]any{"body": "checkendpointtests sms"}},
		{"missing body", map[string]any{"toNumber": "+15550000000"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)
			client := cfg.NewHTTPClient()
			req, err := http.NewRequest(http.MethodPost, cfg.URL("/gsm/send/sms"), bytes.NewReader(body))
			if err != nil {
				t.Fatalf("failed to build request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			defer resp.Body.Close()
			respBody, _ := io.ReadAll(resp.Body)
			if resp.StatusCode == http.StatusOK {
				t.Fatalf("expected %s to be rejected, got 200 OK: %s", tc.name, respBody)
			}
		})
	}
}

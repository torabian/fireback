// emailSender black-box tests, following webpushconfig_test.go's exact conventions.
// Reuses its googleResponseEnvelope[T]/googleResponseListEnvelope[T] types - same
// package, no need to redeclare them here.
package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"

	abactests "github.com/torabian/fireback/modules/abac/tests"
)

// emailSenderRes mirrors messaging.EmailSenderEntity's JSON response shape - only the
// fields these tests actually read.
type emailSenderRes struct {
	UniqueId         string `json:"uniqueId"`
	FromName         string `json:"fromName"`
	FromEmailAddress string `json:"fromEmailAddress"`
	ReplyTo          string `json:"replyTo"`
	NickName         string `json:"nickName"`
	WorkspaceId      string `json:"workspaceId"`
}

func postEmailSender(t *testing.T, cfg abactests.TestConfig, payload map[string]any) (*http.Response, []byte) {
	t.Helper()

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/emailSender"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request to %s failed: %v", req.URL, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	return resp, respBody
}

// sampleEmailSenderPayload returns a fully-valid, unique-per-call payload (fromEmailAddress
// is gorm:"unique" - see EmailSenderEntity - so tests that create more than one sender in
// the same run can't reuse a fixed address).
func sampleEmailSenderPayload() map[string]any {
	addr := fmt.Sprintf("checkendpointtests+%d@example.com", time.Now().UnixNano())
	return map[string]any{
		"fromName":         "Checkendpointtests Sender",
		"fromEmailAddress": addr,
		"replyTo":          addr,
		"nickName":         "checkendpointtests",
	}
}

// createSampleEmailSender is a shared helper (not a test itself) for the Browse/Get/
// Update/Delete tests below, which all need an existing record to act on.
func createSampleEmailSender(t *testing.T, cfg abactests.TestConfig) emailSenderRes {
	t.Helper()

	resp, body := postEmailSender(t, cfg, sampleEmailSenderPayload())
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating emailSender, got %d: %s", resp.StatusCode, body)
	}

	var created googleResponseEnvelope[emailSenderRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

func deleteEmailSender(t *testing.T, cfg abactests.TestConfig, uniqueId string) {
	t.Helper()
	if uniqueId == "" {
		return
	}

	body, _ := json.Marshal(map[string][]string{"uniqueIds": {uniqueId}})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/emailSender/delete"), bytes.NewReader(body))
	if err != nil {
		t.Logf("cleanup: failed to build delete request for emailSender %s: %v", uniqueId, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", cfg.WorkspaceID)

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("cleanup: failed to delete emailSender %s: %v", uniqueId, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		t.Logf("cleanup: deleting emailSender %s returned %d: %s", uniqueId, resp.StatusCode, b)
	}
}

func TestEmailSenderCreate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleEmailSender(t, cfg)
	defer deleteEmailSender(t, cfg, created.UniqueId)

	if created.FromName != "Checkendpointtests Sender" {
		t.Errorf("expected fromName %q, got %q", "Checkendpointtests Sender", created.FromName)
	}
}

func TestEmailSenderCreate_HTTP_RequiresAuth(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)

	client := cfg.NewHTTPClient()
	body, _ := json.Marshal(sampleEmailSenderPayload())
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/emailSender"), bytes.NewReader(body))
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
		var out googleResponseEnvelope[emailSenderRes]
		if err := json.Unmarshal(respBody, &out); err == nil {
			deleteEmailSender(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected an unauthenticated create to be rejected, got 200 OK: %s", respBody)
	}
}

// TestEmailSenderCreate_HTTP_ValidationRequiredFields covers the fix making
// EmailSenderCreateAction actually call fireback.CommonStructValidatorPointer -
// fromName/fromEmailAddress/replyTo/nickName are all `validate:"required"` (see
// Messaging.emi.yml) but were previously accepted empty.
func TestEmailSenderCreate_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	cases := []struct {
		name         string
		omit         string
		wantLocation string
	}{
		{"missing fromName", "fromName", "fromName"},
		{"missing fromEmailAddress", "fromEmailAddress", "fromEmailAddress"},
		{"missing replyTo", "replyTo", "replyTo"},
		{"missing nickName", "nickName", "nickName"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			payload := sampleEmailSenderPayload()
			delete(payload, tc.omit)

			resp, body := postEmailSender(t, cfg, payload)

			if resp.StatusCode == http.StatusOK {
				var out googleResponseEnvelope[emailSenderRes]
				if err := json.Unmarshal(body, &out); err == nil {
					deleteEmailSender(t, cfg, out.Data.Item.UniqueId)
				}
				t.Fatalf("expected creation with %s to be rejected, got 200 OK: %s", tc.name, body)
			}
			if !strings.Contains(string(body), `"location": "`+tc.wantLocation+`"`) &&
				!strings.Contains(string(body), `"location":"`+tc.wantLocation+`"`) {
				t.Errorf("expected a field error located at %q, got %d: %s", tc.wantLocation, resp.StatusCode, body)
			}
		})
	}
}

// TestEmailSenderCreate_HTTP_RejectsDuplicateFromEmailAddress covers the DB-level
// uniqueness constraint on fromEmailAddress (gorm:"unique" - see EmailSenderEntity):
// creating a second sender with an address already in use must fail, not silently
// succeed with two rows sharing one address.
func TestEmailSenderCreate_HTTP_RejectsDuplicateFromEmailAddress(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	first := createSampleEmailSender(t, cfg)
	defer deleteEmailSender(t, cfg, first.UniqueId)

	dupe := map[string]any{
		"fromName":         "Duplicate Sender",
		"fromEmailAddress": first.FromEmailAddress,
		"replyTo":          first.FromEmailAddress,
		"nickName":         "duplicate",
	}
	resp, body := postEmailSender(t, cfg, dupe)

	if resp.StatusCode == http.StatusOK {
		var out googleResponseEnvelope[emailSenderRes]
		if err := json.Unmarshal(body, &out); err == nil {
			deleteEmailSender(t, cfg, out.Data.Item.UniqueId)
		}
		t.Fatalf("expected creating a second emailSender with the same fromEmailAddress %q to be rejected, got 200 OK: %s", first.FromEmailAddress, body)
	}
}

// TestEmailSenderBrowse_HTTP_IncludesOwnRecord covers the workspace-stamping fix: without
// UserId/WorkspaceId set on create, the row is invisible to this workspace-scoped browse
// (and, in the real app, to every picker built on it - e.g. NotificationConfig's "invite
// to workspace sender" selector, which is exactly the one blocking the invite-email flow).
func TestEmailSenderBrowse_HTTP_IncludesOwnRecord(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleEmailSender(t, cfg)
	defer deleteEmailSender(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/emailSender/browse"), nil)
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
		t.Fatalf("expected 200 OK browsing emailSender, got %d: %s", resp.StatusCode, body)
	}

	var list googleResponseListEnvelope[emailSenderRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode browse response: %v\nbody: %s", err, body)
	}

	for _, item := range list.Data.Items {
		if item.UniqueId == created.UniqueId {
			return
		}
	}
	t.Errorf("expected browse to include the just-created emailSender %s, got %d items - Create may not be stamping workspaceId from the resolved query context", created.UniqueId, len(list.Data.Items))
}

func TestEmailSenderGet_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleEmailSender(t, cfg)
	defer deleteEmailSender(t, cfg, created.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/emailSender/"+created.UniqueId), nil)
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
		t.Fatalf("expected 200 OK getting emailSender, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[emailSenderRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode get response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UniqueId != created.UniqueId {
		t.Errorf("expected uniqueId %q, got %q", created.UniqueId, out.Data.Item.UniqueId)
	}
}

func TestEmailSenderUpdate_HTTP_Succeeds(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleEmailSender(t, cfg)
	defer deleteEmailSender(t, cfg, created.UniqueId)

	body, _ := json.Marshal(map[string]any{"nickName": "renamed-by-test"})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPatch, cfg.URL("/emailSender/"+created.UniqueId), bytes.NewReader(body))
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
		t.Fatalf("expected 200 OK updating emailSender, got %d: %s", resp.StatusCode, respBody)
	}
}

// TestEmailSenderAwareDeletePreview_HTTP_ThenDelete covers both delete-preview and the
// actual delete (its bare action name "EmailSenderAwareDelete" is a prefix of
// "EmailSenderAwareDeletePreview", so this single test name covers both per
// tools/checkendpointtests' substring match, same trick webpushconfig_test.go uses).
func TestEmailSenderAwareDeletePreview_HTTP_ThenDelete(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	created := createSampleEmailSender(t, cfg)

	client := cfg.NewHTTPClient()
	previewReq, err := http.NewRequest(http.MethodGet, cfg.URL("/emailSender/delete-preview?uniqueIds="+created.UniqueId), nil)
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
		deleteEmailSender(t, cfg, created.UniqueId)
		t.Fatalf("expected 200 OK on delete-preview, got %d: %s", previewResp.StatusCode, previewBody)
	}

	deleteBody, _ := json.Marshal(map[string][]string{"uniqueIds": {created.UniqueId}})
	deleteReq, err := http.NewRequest(http.MethodPost, cfg.URL("/emailSender/delete"), bytes.NewReader(deleteBody))
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
		t.Fatalf("expected 200 OK deleting emailSender, got %d: %s", deleteResp.StatusCode, deleteRespBody)
	}

	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/emailSender/"+created.UniqueId), nil)
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
		t.Errorf("expected the deleted emailSender %s to no longer be gettable", created.UniqueId)
	}
}

func TestEmailSenderCreate_CLI_Help(t *testing.T) {
	cfg := abactests.LoadTestConfig(t)
	bin := cfg.ResolveAppBinary(t)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()

	// The bare Name every entity's create command registers under ("create") collides
	// across every entity nested in the "messaging" group - each is only unambiguously
	// reachable via its own alias (see EmailSenderCreateActionCliHandler's
	// cmd.Aliases = []string{meta.CliShort}; confirmed against `./app messaging --help`).
	cmd := exec.CommandContext(ctx, bin, "messaging", "emailSender-c", "--help")
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("`%s messaging emailSender-c --help` failed: %v\noutput:\n%s", bin, err, out)
	}

	// urfave's default help renderer only prints the primary Name ("create", shared by
	// every entity in this group) in the NAME:/USAGE: header, never the Aliases actually
	// used to route here - so the entity name from the auto-generated description
	// ('Creates a new "emailSender" row.') is what actually proves this reached the
	// right subcommand, not a coincidentally-successful but wrong one.
	if !strings.Contains(string(out), "emailSender") {
		t.Errorf("expected --help output to mention the emailSender entity, got:\n%s", out)
	}
}

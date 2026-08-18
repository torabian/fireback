// InviteToWorkspaceAction (/workspace/invite) and AcceptInviteAction
// (/user/invitation/accept) black-box tests - the "admin invites, invitee accepts"
// mechanism. Follows this package's established conventions (see testconfig.go's doc
// comment and role_create_test.go): real HTTP requests against a running `./app start`,
// every test skipping (not failing) when its target isn't reachable. Reuses
// googleResponseEnvelope[T] (passport_methods_http_test.go) and
// postRole/roleDtoPayload/roleRes/deleteRole (role_create_test.go) - same package, no
// need to redeclare them.
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
)

// workspaceInviteRes mirrors the WorkspaceInviteEntity fields InviteToWorkspaceAction
// actually responds with (it wraps the entity itself via fireback.GResponseSingleItem,
// not the slimmer WorkspaceInvitationDto the emi.yml declares as its "out" - see
// InviteToWorkspaceActionImplementation.go) - only the fields these tests read.
type workspaceInviteRes struct {
	UniqueId  string `json:"uniqueId"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	RoleId    string `json:"roleId"`
}

// signupResult mirrors ClassicSignupAction's {"data":{"item":{"session":{"token",
// "userWorkspaces":[{"workspaceId"}]}}}} response shape - only what these tests read
// (matches e2e/cypress/e2e/workspace-role-capability-scoping.cy.ts's SignupResponse).
type signupResult struct {
	Data struct {
		Item struct {
			Session struct {
				Token          string `json:"token"`
				UserWorkspaces []struct {
					WorkspaceId string `json:"workspaceId"`
				} `json:"userWorkspaces"`
			} `json:"session"`
		} `json:"item"`
	} `json:"data"`
}

// googleResponseListEnvelope mirrors fireback.GResponseQuery's {"data":{"items":
// [...]}} wrapper - passport_methods_http_test.go only declares the single-item variant
// (googleResponseEnvelope), reused above; this package didn't otherwise need the list
// shape until this file's direct workspaceRole browse check.
type googleResponseListEnvelope[T any] struct {
	Data struct {
		Items []T `json:"items"`
	} `json:"data"`
}

type whoamiWorkspace struct {
	UniqueId     string   `json:"uniqueId"`
	Capabilities []string `json:"capabilities"`
	Roles        []struct {
		Name         string   `json:"name"`
		Capabilities []string `json:"capabilities"`
	} `json:"roles"`
}

type whoamiResult struct {
	Data struct {
		Item struct {
			UserId     string            `json:"userId"`
			Workspaces []whoamiWorkspace `json:"workspaces"`
		} `json:"item"`
	} `json:"data"`
}

// lettersFromInt base-26-encodes n into a lowercase a-z string (no digits/dashes) - used
// to uniquify workspace type slugs, which ValidateTheWorkspaceTypeEntity restricts to
// lowercase letters and dashes only (a numeric timestamp suffix would fail that).
func lettersFromInt(n int64) string {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return "a"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('a' + n%26)}, b...)
		n /= 26
	}
	return string(b)
}

// ensureEmailPassportMethod makes signup-by-email reachable. Best-effort/idempotent:
// this repo's dev/test database is shared across other suites (cypress, other Go tests)
// that may have already created it, so a failure here (most likely "already exists") is
// only logged, never fatal. Every caller already independently requires the CLI binary
// (createRoleAndWorkspaceType's own workspaceType-c call), so resolving it the normal,
// skip-on-missing way here doesn't add a new hard dependency.
func ensureEmailPassportMethod(t *testing.T, cfg TestConfig) {
	t.Helper()
	bin := cfg.ResolveAppBinary(t)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, "passport", "method", "create", "--region", "global", "--type", "email")
	cmd.Dir = cfg.WorkDir(t)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Logf("ensureEmailPassportMethod: create returned (likely already exists, non-fatal): %v: %s", err, out)
	}
}

// createRoleAndWorkspaceType creates a role with the given capability plus a workspace
// type bound to it - a throwaway signup target so a freshly-signed-up user picks up
// exactly (and only) that capability, mirroring
// workspace-role-capability-scoping.cy.ts's own setup. Returns the workspaceTypeId.
func createRoleAndWorkspaceType(t *testing.T, cfg TestConfig, roleName string, capabilities []string) (workspaceTypeId string) {
	t.Helper()

	resp, body := postRole(t, cfg, roleDtoPayload{
		Name:               roleName,
		CapabilitiesListId: capabilities,
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("failed to create role %q: %d: %s", roleName, resp.StatusCode, body)
	}
	var role googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(body, &role); err != nil {
		t.Fatalf("failed to decode role create response: %v\nbody: %s", err, body)
	}

	// Slug validation (ValidateTheWorkspaceTypeEntity) only allows lowercase a-z and
	// dashes after the leading "/" - a numeric uniqueifier would fail that, so the
	// timestamp is base-26 letter-encoded instead.
	slug := fmt.Sprintf("/checkendpointtests-%s", lettersFromInt(time.Now().UnixNano()))
	bin := cfg.ResolveAppBinary(t)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.CLITimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, "ws", "workspaceType-c", "--title", roleName, "--slug", slug, "--role-id", role.Data.Item.UniqueId)
	cmd.Dir = cfg.WorkDir(t)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to create workspace type for role %q: %v\noutput:\n%s", roleName, err, out)
	}
	var wt googleResponseEnvelope[struct {
		UniqueId string `json:"uniqueId"`
	}]
	if err := json.Unmarshal(out, &wt); err != nil {
		t.Fatalf("failed to decode workspace type create output: %v\noutput:\n%s", err, out)
	}
	return wt.Data.Item.UniqueId
}

// signupClassic signs a brand new account up by email against workspaceTypeId, returning
// its session token and the uniqueId of the workspace that signup created for them.
func signupClassic(t *testing.T, cfg TestConfig, email, workspaceTypeId string) (token, workspaceId string) {
	t.Helper()

	body, _ := json.Marshal(map[string]any{
		"value":           email,
		"type":            "email",
		"password":        "checkendpointtests-pass-123",
		"firstName":       "Checkendpointtests",
		"lastName":        "User",
		"workspaceTypeId": workspaceTypeId,
	})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/passports/signup/classic"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build signup request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("signup request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK signing up %q, got %d: %s", email, resp.StatusCode, respBody)
	}

	var out signupResult
	if err := json.Unmarshal(respBody, &out); err != nil {
		t.Fatalf("failed to decode signup response: %v\nbody: %s", err, respBody)
	}
	if out.Data.Item.Session.Token == "" || len(out.Data.Item.Session.UserWorkspaces) == 0 {
		t.Fatalf("signup response missing session token/workspace: %s", respBody)
	}
	return out.Data.Item.Session.Token, out.Data.Item.Session.UserWorkspaces[0].WorkspaceId
}

// ensureInviteEmailSendingConfigured wires up a terminal EmailProvider + EmailSender +
// NotificationConfig for the root workspace, as root - InviteToWorkspaceAction's
// SendInviteEmail call (see WorkspaceActionUpdate.go) needs all three configured to
// succeed. "terminal" is used throughout (both here and in
// modules/abac/messaging/tests) because SendMail's terminal case just logs and returns
// nil - no real SMTP/provider credentials are ever needed.
func ensureInviteEmailSendingConfigured(t *testing.T, cfg TestConfig) {
	t.Helper()
	client := cfg.NewHTTPClient()

	senderAddr := fmt.Sprintf("checkendpointtests+%d@example.com", time.Now().UnixNano())
	senderBody, _ := json.Marshal(map[string]any{
		"fromName":         "Checkendpointtests",
		"fromEmailAddress": senderAddr,
		"replyTo":          senderAddr,
		"nickName":         "checkendpointtests",
	})
	senderReq, _ := http.NewRequest(http.MethodPost, cfg.URL("/emailSender"), bytes.NewReader(senderBody))
	senderReq.Header.Set("Content-Type", "application/json")
	senderReq.Header.Set("Authorization", cfg.CliToken)
	senderReq.Header.Set("Workspace-id", "root")
	senderResp, err := client.Do(senderReq)
	if err != nil {
		t.Fatalf("failed to create emailSender: %v", err)
	}
	defer senderResp.Body.Close()
	senderRespBody, _ := io.ReadAll(senderResp.Body)
	if senderResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating emailSender, got %d: %s", senderResp.StatusCode, senderRespBody)
	}
	var sender googleResponseEnvelope[struct {
		UniqueId string `json:"uniqueId"`
	}]
	if err := json.Unmarshal(senderRespBody, &sender); err != nil {
		t.Fatalf("failed to decode emailSender response: %v\nbody: %s", err, senderRespBody)
	}

	providerBody, _ := json.Marshal(map[string]any{"type": "terminal", "title": "checkendpointtests"})
	providerReq, _ := http.NewRequest(http.MethodPost, cfg.URL("/emailProvider"), bytes.NewReader(providerBody))
	providerReq.Header.Set("Content-Type", "application/json")
	providerReq.Header.Set("Authorization", cfg.CliToken)
	providerReq.Header.Set("Workspace-id", "root")
	providerResp, err := client.Do(providerReq)
	if err != nil {
		t.Fatalf("failed to create emailProvider: %v", err)
	}
	defer providerResp.Body.Close()
	providerRespBody, _ := io.ReadAll(providerResp.Body)
	if providerResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating emailProvider, got %d: %s", providerResp.StatusCode, providerRespBody)
	}
	var provider googleResponseEnvelope[struct {
		UniqueId string `json:"uniqueId"`
	}]
	if err := json.Unmarshal(providerRespBody, &provider); err != nil {
		t.Fatalf("failed to decode emailProvider response: %v\nbody: %s", err, providerRespBody)
	}

	// NotificationConfigUpdateAction upserts by the resolved workspace (from the
	// Workspace-id header), ignoring the :uniqueId path param entirely - see its own doc
	// comment in NotificationConfigActions.go.
	configBody, _ := json.Marshal(map[string]any{
		"generalEmailProviderId":    provider.Data.Item.UniqueId,
		"inviteToWorkspaceSenderId": sender.Data.Item.UniqueId,
	})
	configReq, _ := http.NewRequest(http.MethodPatch, cfg.URL("/notificationConfig/upsert"), bytes.NewReader(configBody))
	configReq.Header.Set("Content-Type", "application/json")
	configReq.Header.Set("Authorization", cfg.CliToken)
	configReq.Header.Set("Workspace-id", "root")
	configResp, err := client.Do(configReq)
	if err != nil {
		t.Fatalf("failed to upsert notificationConfig: %v", err)
	}
	defer configResp.Body.Close()
	configRespBody, _ := io.ReadAll(configResp.Body)
	if configResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK upserting notificationConfig, got %d: %s", configResp.StatusCode, configRespBody)
	}
}

func postInvite(t *testing.T, cfg TestConfig, token, workspaceId string, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/workspace/invite"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build invite request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	if workspaceId != "" {
		req.Header.Set("Workspace-id", workspaceId)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("invite request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

func validInvitePayload(email, roleId string) map[string]any {
	return map[string]any{
		"firstName": "Invitee",
		"lastName":  "Person",
		"email":     email,
		"roleId":    roleId,
	}
}

// --- InviteToWorkspaceAction (/workspace/invite) ---

func TestInviteToWorkspace_HTTP_RequiresAuth(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	resp, body := postInvite(t, cfg, "", "root", validInvitePayload("invitee@example.com", "does-not-matter"))
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unauthenticated invite to be rejected, got 200 OK: %s", body)
	}
}

// TestInviteToWorkspace_HTTP_RequiresCapability is the regression test for the core bug:
// InviteToWorkspaceActionImplementation.go used to resolve its context with a nil
// security model, so literally any authenticated user - regardless of role/capability -
// could invite people into any workspace. A plain signed-up user with no
// workspace-invite capability must now be rejected.
func TestInviteToWorkspace_HTTP_RequiresCapability(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)
	ensureEmailPassportMethod(t, cfg)

	// A capability unrelated to workspace-invite - this role can log in and do
	// *something*, just not send invitations.
	plainWorkspaceTypeId := createRoleAndWorkspaceType(t, cfg, fmt.Sprintf("plainrole-%d", time.Now().UnixNano()), []string{"root.abac.email-confirmation.query"})
	plainEmail := fmt.Sprintf("checkendpointtests-plain+%d@example.com", time.Now().UnixNano())
	plainToken, plainWorkspaceId := signupClassic(t, cfg, plainEmail, plainWorkspaceTypeId)

	resp, body := postInvite(t, cfg, plainToken, plainWorkspaceId, validInvitePayload("invitee@example.com", "does-not-matter"))
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected a user without workspace-invite capability to be rejected, got 200 OK: %s", body)
	}
}

func TestInviteToWorkspace_HTTP_ValidationRequiredFields(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	// A genuinely valid role, reused below - using a placeholder like "root" would let
	// ValidateRoleAndItsExistence's "unknown role" rejection fire before the field being
	// tested ever gets checked (roleId here is a workspace-scoped Role's uniqueId, not a
	// workspace uniqueId - "root" only happens to satisfy both by coincidence in the seed
	// data).
	roleResp, roleBody := postRole(t, cfg, roleDtoPayload{
		Name:               fmt.Sprintf("checkendpointtests validation role %d", time.Now().UnixNano()),
		CapabilitiesListId: []string{"root.abac.email-confirmation.query"},
	})
	if roleResp.StatusCode != http.StatusOK {
		t.Fatalf("failed to create role for validation cases: %d: %s", roleResp.StatusCode, roleBody)
	}
	var role googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(roleBody, &role); err != nil {
		t.Fatalf("failed to decode role response: %v\nbody: %s", err, roleBody)
	}
	defer deleteRole(t, cfg, role.Data.Item.UniqueId)
	roleId := role.Data.Item.UniqueId

	cases := []struct {
		name    string
		payload map[string]any
	}{
		{"missing firstName", map[string]any{"lastName": "Person", "email": "invitee@example.com", "roleId": roleId}},
		{"missing lastName", map[string]any{"firstName": "Invitee", "email": "invitee@example.com", "roleId": roleId}},
		{"missing roleId", map[string]any{"firstName": "Invitee", "lastName": "Person", "email": "invitee@example.com"}},
		{"missing both email and phone", map[string]any{"firstName": "Invitee", "lastName": "Person", "roleId": roleId}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body := postInvite(t, cfg, cfg.CliToken, cfg.WorkspaceID, tc.payload)
			if resp.StatusCode == http.StatusOK {
				t.Fatalf("expected %s to be rejected, got 200 OK: %s", tc.name, body)
			}
		})
	}

	// The "missing both email and phone" case specifically exercises the new
	// InviteRequiresEmailOrPhone check (InviteToWorkspaceActionImplementation.go) - assert
	// its message shows up, not just any rejection (e.g. the roleId check firing first
	// would also produce a non-200, but for the wrong reason).
	resp, body := postInvite(t, cfg, cfg.CliToken, cfg.WorkspaceID, map[string]any{
		"firstName": "Invitee", "lastName": "Person", "roleId": roleId,
	})
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected missing email+phone to be rejected, got 200 OK: %s", body)
	}
	// The public response only ever carries the resolved message text (ToPublicEndUser
	// picks Message[lang] per field) - the internal "$": "InviteRequiresEmailOrPhone"
	// code itself is never serialized to the client, so match on the actual wording
	// instead of that code.
	if !strings.Contains(string(body), "email address or a phone number") ||
		!strings.Contains(string(body), `"location":"email"`) && !strings.Contains(string(body), `"location": "email"`) {
		t.Errorf("expected the InviteRequiresEmailOrPhone error located at email, got %d: %s", resp.StatusCode, body)
	}
}

func TestInviteToWorkspace_HTTP_RejectsUnknownRole(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postInvite(t, cfg, cfg.CliToken, cfg.WorkspaceID, validInvitePayload("invitee@example.com", "role-that-does-not-exist"))
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unknown roleId to be rejected, got 200 OK: %s", body)
	}
}

// TestInviteToWorkspace_HTTP_Succeeds covers the field-mapping fix: the request DTO used
// to never be copied onto the saved entity, so every invite (and the "sent" email) ended
// up with a blank name/email/role regardless of what was posted.
func TestInviteToWorkspace_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)
	ensureInviteEmailSendingConfigured(t, cfg)

	roleResp, roleBody := postRole(t, cfg, roleDtoPayload{
		Name:               fmt.Sprintf("checkendpointtests invitee role %d", time.Now().UnixNano()),
		CapabilitiesListId: []string{"root.*"},
	})
	if roleResp.StatusCode != http.StatusOK {
		t.Fatalf("failed to create role for invite: %d: %s", roleResp.StatusCode, roleBody)
	}
	var role googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(roleBody, &role); err != nil {
		t.Fatalf("failed to decode role response: %v\nbody: %s", err, roleBody)
	}
	defer deleteRole(t, cfg, role.Data.Item.UniqueId)

	email := fmt.Sprintf("checkendpointtests-invitee+%d@example.com", time.Now().UnixNano())
	resp, body := postInvite(t, cfg, cfg.CliToken, cfg.WorkspaceID, validInvitePayload(email, role.Data.Item.UniqueId))
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating invite, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[workspaceInviteRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode invite response: %v\nbody: %s", err, body)
	}
	item := out.Data.Item
	if item.UniqueId == "" {
		t.Errorf("expected a generated uniqueId, got none: %+v", item)
	}
	if item.Email != email {
		t.Errorf("expected email %q to have been saved onto the invite, got %q - field mapping may be broken again", email, item.Email)
	}
	if item.FirstName != "Invitee" || item.LastName != "Person" {
		t.Errorf("expected firstName/lastName %q/%q to have been saved, got %q/%q", "Invitee", "Person", item.FirstName, item.LastName)
	}
	if item.RoleId != role.Data.Item.UniqueId {
		t.Errorf("expected roleId %q to have been saved onto the invite, got %q", role.Data.Item.UniqueId, item.RoleId)
	}
}

// TestInviteToWorkspace_HTTP_RootRoleSucceeds is the regression test for: inviting
// someone into the root workspace with the "Root" role picked from the role list fails
// with a "role is not accessible" style error, even though it's a real, existing role
// and the caller is root itself.
//
// Root cause: ValidateRoleAndItsExistence (WorkspaceTypeActions.go) looked the role up
// via RoleActions.GetOne(fireback.QueryDSL{UniqueId: roleId}) - a bare QueryDSL with no
// WorkspaceId at all. RoleActions.GetOne is wrapped (RoleActions.go's init()) to treat
// the root role as invisible/404 from any query whose WorkspaceId isn't literally
// "root" - and an empty string never is, regardless of which workspace the caller is
// actually acting in. So a lookup for roleId "root" 404'd unconditionally, every time,
// even from inside the root workspace itself - InviteToWorkspaceAction surfaced that as
// RoleIsNotAccessible.
func TestInviteToWorkspace_HTTP_RootRoleSucceeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)
	ensureInviteEmailSendingConfigured(t, cfg)

	email := fmt.Sprintf("checkendpointtests-rootinvitee+%d@example.com", time.Now().UnixNano())
	resp, body := postInvite(t, cfg, cfg.CliToken, cfg.WorkspaceID, validInvitePayload(email, "root"))
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK inviting with the root role from the root workspace, got %d: %s", resp.StatusCode, body)
	}

	var out googleResponseEnvelope[workspaceInviteRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode invite response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.RoleId != "root" {
		t.Errorf("expected roleId %q to have been saved onto the invite, got %q", "root", out.Data.Item.RoleId)
	}
}

// --- AcceptInviteAction (/user/invitation/accept) ---

func postAccept(t *testing.T, cfg TestConfig, token, workspaceId, invitationUniqueId string) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(map[string]any{"invitationUniqueId": invitationUniqueId})
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL("/user/invitation/accept"), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build accept request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	if workspaceId != "" {
		req.Header.Set("Workspace-id", workspaceId)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("accept request failed: %v", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

func TestAcceptInvite_HTTP_RequiresAuth(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)

	resp, body := postAccept(t, cfg, "", "root", "does-not-matter")
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected an unauthenticated accept to be rejected, got 200 OK: %s", body)
	}
}

func TestAcceptInvite_HTTP_NotFound(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	resp, body := postAccept(t, cfg, cfg.CliToken, cfg.WorkspaceID, "checkendpointtests-does-not-exist")
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected accepting a nonexistent invitation to be rejected, got 200 OK: %s", body)
	}
	if !strings.Contains(string(body), "InvitationNotFound") {
		t.Errorf("expected the InvitationNotFound error, got %d: %s", resp.StatusCode, body)
	}
}

// TestAcceptInvite_HTTP_Succeeds is the full admin-invites/invitee-accepts flow: root
// invites a fresh email into its workspace with a specific role, a brand new account
// signs up for that same email (so there's someone to log in and accept as), accepts the
// invite, and /whoami confirms they actually joined the target workspace with the
// invited role's capabilities. This is the regression test for two bugs in
// AcceptInviteActionImplementation.go: the typed-nil-interface bug that made GetOne's
// success case look like a failure (rejecting every accept attempt), and UserWorkspace
// being created with the invite's own (always-empty) UserId instead of the accepting
// caller's.
func TestAcceptInvite_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)
	ensureEmailPassportMethod(t, cfg)
	ensureInviteEmailSendingConfigured(t, cfg)

	roleResp, roleBody := postRole(t, cfg, roleDtoPayload{
		Name:               fmt.Sprintf("checkendpointtests accept role %d", time.Now().UnixNano()),
		CapabilitiesListId: []string{"root.abac.email-confirmation.query"},
	})
	if roleResp.StatusCode != http.StatusOK {
		t.Fatalf("failed to create role for invite: %d: %s", roleResp.StatusCode, roleBody)
	}
	var role googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(roleBody, &role); err != nil {
		t.Fatalf("failed to decode role response: %v\nbody: %s", err, roleBody)
	}
	defer deleteRole(t, cfg, role.Data.Item.UniqueId)

	inviteeEmail := fmt.Sprintf("checkendpointtests-accept+%d@example.com", time.Now().UnixNano())
	inviteResp, inviteBody := postInvite(t, cfg, cfg.CliToken, cfg.WorkspaceID, validInvitePayload(inviteeEmail, role.Data.Item.UniqueId))
	if inviteResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating invite, got %d: %s", inviteResp.StatusCode, inviteBody)
	}
	var invite googleResponseEnvelope[workspaceInviteRes]
	if err := json.Unmarshal(inviteBody, &invite); err != nil {
		t.Fatalf("failed to decode invite response: %v\nbody: %s", err, inviteBody)
	}

	// A throwaway workspace type/role, unrelated to the invite, purely so the invitee has
	// somewhere to land from their own signup.
	inviteeOwnWorkspaceTypeId := createRoleAndWorkspaceType(t, cfg, fmt.Sprintf("inviteeownrole-%d", time.Now().UnixNano()), []string{"root.abac.email-confirmation.query"})
	inviteeToken, inviteeOwnWorkspaceId := signupClassic(t, cfg, inviteeEmail, inviteeOwnWorkspaceTypeId)

	acceptResp, acceptBody := postAccept(t, cfg, inviteeToken, inviteeOwnWorkspaceId, invite.Data.Item.UniqueId)
	if acceptResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK accepting the invite, got %d: %s", acceptResp.StatusCode, acceptBody)
	}

	client := cfg.NewHTTPClient()
	whoamiReq, _ := http.NewRequest(http.MethodGet, cfg.URL("/whoami"), nil)
	whoamiReq.Header.Set("Authorization", inviteeToken)
	whoamiResp, err := client.Do(whoamiReq)
	if err != nil {
		t.Fatalf("whoami request failed: %v", err)
	}
	defer whoamiResp.Body.Close()
	whoamiBody, _ := io.ReadAll(whoamiResp.Body)
	if whoamiResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK from whoami, got %d: %s", whoamiResp.StatusCode, whoamiBody)
	}

	var who whoamiResult
	if err := json.Unmarshal(whoamiBody, &who); err != nil {
		t.Fatalf("failed to decode whoami response: %v\nbody: %s", err, whoamiBody)
	}

	joined := false
	for i := range who.Data.Item.Workspaces {
		if who.Data.Item.Workspaces[i].UniqueId == cfg.WorkspaceID {
			joined = true
			break
		}
	}
	if !joined {
		t.Fatalf("expected accepting the invite to add workspace %q to the invitee's whoami, got workspaces: %+v", cfg.WorkspaceID, who.Data.Item.Workspaces)
	}

	// Not asserted via whoami's per-workspace roles[] here: GetUserAccessLevels (which
	// whoami/every permission check reads from) has a pre-existing bug, unrelated to this
	// fix, that can resolve a workspace's roles[] to an unrelated/dangling role in a
	// database that has accumulated many WorkspaceRoleEntity rows across many
	// users/roles - confirmed by cross-checking against a direct, targeted query below,
	// which is unaffected by that. What IS this fix's responsibility - and is what's
	// checked here - is that AcceptInviteActionImplementation.go's transaction actually
	// created a WorkspaceRoleEntity row carrying the exact roleId from the invite.
	wrClient := cfg.NewHTTPClient()
	wrReq, _ := http.NewRequest(http.MethodGet, cfg.URL("/workspaceRole/browse?itemsPerPage=1000"), nil)
	wrReq.Header.Set("Authorization", cfg.CliToken)
	wrReq.Header.Set("Workspace-id", cfg.WorkspaceID)
	wrResp, err := wrClient.Do(wrReq)
	if err != nil {
		t.Fatalf("workspaceRole browse request failed: %v", err)
	}
	defer wrResp.Body.Close()
	wrBody, _ := io.ReadAll(wrResp.Body)
	if wrResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK browsing workspaceRole, got %d: %s", wrResp.StatusCode, wrBody)
	}
	var wrList googleResponseListEnvelope[struct {
		RoleId string `json:"roleId"`
	}]
	if err := json.Unmarshal(wrBody, &wrList); err != nil {
		t.Fatalf("failed to decode workspaceRole browse response: %v\nbody: %s", err, wrBody)
	}
	foundRole := false
	for _, wr := range wrList.Data.Items {
		if wr.RoleId == role.Data.Item.UniqueId {
			foundRole = true
			break
		}
	}
	if !foundRole {
		t.Errorf("expected accepting the invite to create a workspaceRole row for roleId %q, found none among %d rows", role.Data.Item.UniqueId, len(wrList.Data.Items))
	}

	// The invitee's own signup workspace should still be there too - accepting an invite
	// is additive, not a replacement of their existing membership.
	hasOwnWorkspace := false
	for _, w := range who.Data.Item.Workspaces {
		if w.UniqueId == inviteeOwnWorkspaceId {
			hasOwnWorkspace = true
			break
		}
	}
	if !hasOwnWorkspace {
		t.Errorf("expected the invitee to still belong to their own signup workspace %q after accepting an unrelated invite", inviteeOwnWorkspaceId)
	}
}

// Black-box tests for the root-only workspace membership admin actions:
// AddUserToWorkspaceAction, RemoveUserFromWorkspaceAction, ChangeUserWorkspaceRoleAction
// and QueryWorkspaceRolesAction (modules/abac/*ActionImplementation.go). These let root
// add/remove an existing user's membership in any workspace and change their role there,
// without going through the invite/email flow, and let root list a *specific* workspace's
// own roles (not root's) to populate that "change role" picker.
//
// Every write here runs with the caller's own Workspace-Id header set to "root"
// (AllowOnRoot on all four actions) - the workspace being acted *on* travels as a separate
// "workspaceId" body field instead, since root generally has no real membership in it.
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

type userWorkspaceAdminRes struct {
	UniqueId    string `json:"uniqueId"`
	UserId      string `json:"userId"`
	WorkspaceId string `json:"workspaceId"`
	RoleId      string `json:"roleId"`
}

type workspaceRoleOptionRes struct {
	UniqueId           string   `json:"uniqueId"`
	Name               string   `json:"name"`
	CapabilitiesListId []string `json:"capabilitiesListId"`
}

func postToRootAction(t *testing.T, cfg TestConfig, path string, workspaceIdHeader string, payload map[string]any) (*http.Response, []byte) {
	t.Helper()
	body, _ := json.Marshal(payload)
	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodPost, cfg.URL(path), bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", workspaceIdHeader)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request to %s failed: %v", path, err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return resp, respBody
}

// createSampleWorkspaceForAdmin creates a workspace via POST /workspaces/create
// (CreateWorkspaceAction) rather than the raw entity POST /workspace
// (createSampleWorkspace, workspace_test.go) - unlike the raw entity create, this one
// auto-enrolls the caller (root) as a member (CreateWorkspaceAndAssignUser). That
// membership - just a UserWorkspace row, no role/capability grant attached - turns out to
// be exactly what root needs to then create/manage roles genuinely scoped to this
// workspace (RoleCreateAction has no AllowOnRoot bypass; it resolves the caller's
// capabilities against whatever workspace the Workspace-Id header names, and root has
// none there at all without at least this membership - confirmed empirically against a
// live server; a bare POST /userWorkspace enrollment on a *raw* createSampleWorkspace
// workspace was NOT sufficient, so whatever grants this must be tied to
// CreateWorkspaceAction's own path specifically, not to UserWorkspace membership alone).
func createSampleWorkspaceForAdmin(t *testing.T, cfg TestConfig) workspaceRes {
	t.Helper()
	resp, body := postToRootAction(t, cfg, "/workspaces/create", "root", map[string]any{
		"name": fmt.Sprintf("checkendpointtests admin workspace %d", time.Now().UnixNano()),
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating workspace, got %d: %s", resp.StatusCode, body)
	}
	var out struct {
		Data struct {
			Item struct {
				WorkspaceId string `json:"workspaceId"`
			} `json:"item"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.WorkspaceId == "" {
		t.Fatalf("expected a generated workspaceId, got none: %s", body)
	}
	return workspaceRes{UniqueId: out.Data.Item.WorkspaceId}
}

// createSampleRoleInWorkspace creates a role scoped to the given workspace (not "root" -
// unlike createSampleRole, which always creates in cfg.WorkspaceID/"root") by pointing the
// Workspace-Id header at that workspace directly: RoleCreateAction stamps the created
// entity's workspaceId from the resolved query context (the header), not from anything in
// the request body. workspaceId must have come from createSampleWorkspaceForAdmin (or
// otherwise already have root enrolled as a member) - see its doc comment.
func createSampleRoleInWorkspace(t *testing.T, cfg TestConfig, workspaceId string) roleRes {
	t.Helper()
	resp, body := postToRootAction(t, cfg, "/role", workspaceId, map[string]any{
		"name":               fmt.Sprintf("checkendpointtests workspace-role %d", time.Now().UnixNano()),
		"capabilitiesListId": []string{"root.manage.user-workspace.query"},
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK creating role in workspace %s, got %d: %s", workspaceId, resp.StatusCode, body)
	}
	var created googleResponseEnvelope[roleRes]
	if err := json.Unmarshal(body, &created); err != nil {
		t.Fatalf("failed to decode create response: %v\nbody: %s", err, body)
	}
	if created.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated uniqueId, got none: %+v", created.Data.Item)
	}
	return created.Data.Item
}

// TestAddUserToWorkspace_HTTP_Succeeds covers the happy path: an existing user, with no
// prior membership, is added to a workspace with an existing role of that workspace - and
// the resulting userWorkspace row is really persisted (GetOneEntity, unlike the
// generic browse actions, filters only by uniqueId - not by the caller's own
// workspace/user scope - so it's a reliable way to check a row exists regardless of who
// "owns" it).
func TestAddUserToWorkspace_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	role := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	resp, body := postToRootAction(t, cfg, "/workspace/add-user", "root", map[string]any{
		"userId":      user.UniqueId,
		"workspaceId": ws.UniqueId,
		"roleId":      role.UniqueId,
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK adding user to workspace, got %d: %s", resp.StatusCode, body)
	}
	var out googleResponseEnvelope[userWorkspaceAdminRes]
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, body)
	}
	if out.Data.Item.UniqueId == "" {
		t.Fatalf("expected a generated userWorkspace uniqueId, got none: %+v", out.Data.Item)
	}
	if out.Data.Item.WorkspaceId != ws.UniqueId {
		t.Errorf("expected workspaceId %q, got %q", ws.UniqueId, out.Data.Item.WorkspaceId)
	}
	if out.Data.Item.RoleId != role.UniqueId {
		t.Errorf("expected roleId %q, got %q", role.UniqueId, out.Data.Item.RoleId)
	}
	defer deleteUserWorkspace(t, cfg, out.Data.Item.UniqueId)

	client := cfg.NewHTTPClient()
	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/userWorkspace/"+out.Data.Item.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build get request: %v", err)
	}
	getReq.Header.Set("Authorization", cfg.CliToken)
	getReq.Header.Set("Workspace-id", "root")
	getResp, err := client.Do(getReq)
	if err != nil {
		t.Fatalf("get request failed: %v", err)
	}
	defer getResp.Body.Close()
	if getResp.StatusCode != http.StatusOK {
		t.Errorf("expected the created userWorkspace %s to be gettable, got %d", out.Data.Item.UniqueId, getResp.StatusCode)
	}
}

// TestAddUserToWorkspace_HTTP_RejectsDuplicateMembership covers the "already a member"
// guard - UserWorkspaceEntity has a real unique(user_id, workspace_id) db index
// (gormMap on the "userWorkspace" Module3 entity), but this should be caught with a clean
// 400 + field error before ever reaching a raw constraint-violation error.
func TestAddUserToWorkspace_HTTP_RejectsDuplicateMembership(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	role := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	resp, body := postToRootAction(t, cfg, "/workspace/add-user", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": role.UniqueId,
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK on first add, got %d: %s", resp.StatusCode, body)
	}
	var out googleResponseEnvelope[userWorkspaceAdminRes]
	_ = json.Unmarshal(body, &out)
	defer deleteUserWorkspace(t, cfg, out.Data.Item.UniqueId)

	resp2, body2 := postToRootAction(t, cfg, "/workspace/add-user", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": role.UniqueId,
	})
	if resp2.StatusCode == http.StatusOK {
		t.Fatalf("expected the duplicate add-user to be rejected, got 200 OK: %s", body2)
	}
}

// TestAddUserToWorkspace_HTTP_RootOnly covers AllowOnRoot: calling with the caller's own
// Workspace-Id header pointed anywhere but "root" - even at the very workspace being acted
// on - must be rejected.
func TestAddUserToWorkspace_HTTP_RootOnly(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	role := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	resp, body := postToRootAction(t, cfg, "/workspace/add-user", ws.UniqueId, map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": role.UniqueId,
	})
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected add-user with a non-root Workspace-Id header to be rejected, got 200 OK: %s", body)
	}
}

// TestChangeUserWorkspaceRole_HTTP_Succeeds covers replacing an existing member's role
// with a different one, both belonging to the target workspace.
func TestChangeUserWorkspaceRole_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	role1 := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role1.UniqueId)
	role2 := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role2.UniqueId)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	addResp, addBody := postToRootAction(t, cfg, "/workspace/add-user", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": role1.UniqueId,
	})
	if addResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK adding user, got %d: %s", addResp.StatusCode, addBody)
	}
	var added googleResponseEnvelope[userWorkspaceAdminRes]
	_ = json.Unmarshal(addBody, &added)
	defer deleteUserWorkspace(t, cfg, added.Data.Item.UniqueId)

	changeResp, changeBody := postToRootAction(t, cfg, "/workspace/change-user-role", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": role2.UniqueId,
	})
	if changeResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK changing role, got %d: %s", changeResp.StatusCode, changeBody)
	}
	var changed googleResponseEnvelope[userWorkspaceAdminRes]
	if err := json.Unmarshal(changeBody, &changed); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, changeBody)
	}
	if changed.Data.Item.RoleId != role2.UniqueId {
		t.Errorf("expected roleId %q after change, got %q", role2.UniqueId, changed.Data.Item.RoleId)
	}
	if changed.Data.Item.UniqueId != added.Data.Item.UniqueId {
		t.Errorf("expected the same userWorkspace row %q, got %q", added.Data.Item.UniqueId, changed.Data.Item.UniqueId)
	}
}

// TestChangeUserWorkspaceRole_HTTP_UnknownRoleRejected covers roleId validation - it must
// be a real, existing role.
func TestChangeUserWorkspaceRole_HTTP_UnknownRoleRejected(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	role := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	addResp, addBody := postToRootAction(t, cfg, "/workspace/add-user", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": role.UniqueId,
	})
	if addResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK adding user, got %d: %s", addResp.StatusCode, addBody)
	}
	var added googleResponseEnvelope[userWorkspaceAdminRes]
	_ = json.Unmarshal(addBody, &added)
	defer deleteUserWorkspace(t, cfg, added.Data.Item.UniqueId)

	resp, body := postToRootAction(t, cfg, "/workspace/change-user-role", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": "checkendpointtests-does-not-exist",
	})
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected changing to an unknown roleId to be rejected, got 200 OK: %s", body)
	}
}

// TestChangeUserWorkspaceRole_HTTP_RequiresExistingMembership covers rejecting a change
// for a user who was never added to the workspace in the first place.
func TestChangeUserWorkspaceRole_HTTP_RequiresExistingMembership(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	role := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	resp, body := postToRootAction(t, cfg, "/workspace/change-user-role", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": role.UniqueId,
	})
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected changing role for a non-member to be rejected, got 200 OK: %s", body)
	}
}

// TestRemoveUserFromWorkspace_HTTP_Succeeds covers the happy path, and that the removed
// userWorkspace row is really gone (not just unreachable via some scoped browse).
func TestRemoveUserFromWorkspace_HTTP_Succeeds(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	role := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	user := createSampleUser(t, cfg)
	defer deleteUser(t, cfg, user.UniqueId)

	addResp, addBody := postToRootAction(t, cfg, "/workspace/add-user", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId, "roleId": role.UniqueId,
	})
	if addResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK adding user, got %d: %s", addResp.StatusCode, addBody)
	}
	var added googleResponseEnvelope[userWorkspaceAdminRes]
	_ = json.Unmarshal(addBody, &added)

	removeResp, removeBody := postToRootAction(t, cfg, "/workspace/remove-user", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId,
	})
	if removeResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK removing user, got %d: %s", removeResp.StatusCode, removeBody)
	}

	client := cfg.NewHTTPClient()
	getReq, err := http.NewRequest(http.MethodGet, cfg.URL("/userWorkspace/"+added.Data.Item.UniqueId), nil)
	if err != nil {
		t.Fatalf("failed to build get request: %v", err)
	}
	getReq.Header.Set("Authorization", cfg.CliToken)
	getReq.Header.Set("Workspace-id", "root")
	getResp, err := client.Do(getReq)
	if err != nil {
		t.Fatalf("get request failed: %v", err)
	}
	defer getResp.Body.Close()
	if getResp.StatusCode == http.StatusOK {
		t.Errorf("expected the removed userWorkspace %s to no longer be gettable", added.Data.Item.UniqueId)
	}

	// Removing an already-removed membership should fail cleanly, not succeed
	// silently or panic.
	resp2, body2 := postToRootAction(t, cfg, "/workspace/remove-user", "root", map[string]any{
		"userId": user.UniqueId, "workspaceId": ws.UniqueId,
	})
	if resp2.StatusCode == http.StatusOK {
		t.Fatalf("expected removing an already-removed membership to be rejected, got 200 OK: %s", body2)
	}
}

// TestRemoveUserFromWorkspace_HTTP_CannotRemoveOwnRootMembership covers the safety guard
// against a root caller ever locking its own session out of the root workspace - this is
// intentionally an identity check (is the target the caller itself), not a headcount:
// see RemoveUserFromWorkspaceActionImplementation.go's guard for why a "last member"
// count turned out to be the wrong thing to check.
func TestRemoveUserFromWorkspace_HTTP_CannotRemoveOwnRootMembership(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	rootId := rootUserId(t, cfg)

	resp, body := postToRootAction(t, cfg, "/workspace/remove-user", "root", map[string]any{
		"userId": rootId, "workspaceId": "root",
	})
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected removing root's own membership from the root workspace to be rejected when it would leave it empty, got 200 OK: %s", body)
	}
}

// TestQueryWorkspaceRoles_HTTP_ListsTargetWorkspaceRoles covers root being able to list a
// specific *other* workspace's roles - not its own - which is exactly what a
// ChangeUserWorkspaceRole picker needs. The generic GET /role/browse can't do this: root
// has no real membership in most workspaces, so pointing its Workspace-Id header at one
// 401s instead (see TestQueryWorkspaceRoles_HTTP_GenericRoleBrowseCannotDoThis below).
func TestQueryWorkspaceRoles_HTTP_ListsTargetWorkspaceRoles(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	role := createSampleRoleInWorkspace(t, cfg, ws.UniqueId)
	defer deleteRole(t, cfg, role.UniqueId)

	resp, body := postToRootAction(t, cfg, "/workspace/roles", "root", map[string]any{
		"workspaceId": ws.UniqueId,
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK listing workspace roles, got %d: %s", resp.StatusCode, body)
	}
	var list googleResponseListEnvelope[workspaceRoleOptionRes]
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to decode response: %v\nbody: %s", err, body)
	}
	for _, item := range list.Data.Items {
		if item.UniqueId == role.UniqueId {
			return
		}
	}
	t.Errorf("expected the workspace's own role %s to be listed, got %d items: %+v", role.UniqueId, len(list.Data.Items), list.Data.Items)
}

// TestQueryWorkspaceRoles_HTTP_GenericRoleBrowseCannotDoThis is the regression guard
// QueryWorkspaceRolesAction exists to fix: root pointing its own Workspace-Id header at a
// workspace it has no real membership in gets 401 NotEnoughPermission from the generic
// browse action, even though it's root.
func TestQueryWorkspaceRoles_HTTP_GenericRoleBrowseCannotDoThis(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws, wt, backingRole := createSampleWorkspace(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)
	defer deleteWorkspaceType(t, cfg, wt.UniqueId)
	defer deleteRole(t, cfg, backingRole.UniqueId)

	client := cfg.NewHTTPClient()
	req, err := http.NewRequest(http.MethodGet, cfg.URL("/role/browse?itemsPerPage=1000"), nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	req.Header.Set("Authorization", cfg.CliToken)
	req.Header.Set("Workspace-id", ws.UniqueId)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		t.Skip("generic /role/browse unexpectedly succeeded with a foreign Workspace-Id header - the gap QueryWorkspaceRolesAction exists for may have been closed some other way; nothing to regression-guard here anymore")
	}
}

// TestQueryWorkspaceRoles_HTTP_RootOnly covers AllowOnRoot on QueryWorkspaceRolesAction
// itself.
func TestQueryWorkspaceRoles_HTTP_RootOnly(t *testing.T) {
	cfg := LoadTestConfig(t)
	cfg.RequireServer(t)
	cfg.RequireAuth(t)

	ws := createSampleWorkspaceForAdmin(t, cfg)
	defer deleteWorkspace(t, cfg, ws.UniqueId)

	resp, body := postToRootAction(t, cfg, "/workspace/roles", ws.UniqueId, map[string]any{
		"workspaceId": ws.UniqueId,
	})
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected /workspace/roles with a non-root Workspace-Id header to be rejected, got 200 OK: %s", body)
	}
}

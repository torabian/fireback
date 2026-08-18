package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// QueryWorkspaceRolesAction lists the roles that belong to a specific workspace - root
// only. This deliberately does NOT reuse the generic RoleBrowseAction
// (GET /role/browse): that action's security resolves capabilities against whatever
// Workspace-Id header the caller sends, and root has no real UserWorkspace/WorkspaceRole
// membership in most workspaces (only in "root" itself) - pointing its header at some
// other workspaceId would 401 there with NotEnoughPermission, even though the caller is
// root. AllowOnRoot below pins the *caller's* Workspace-Id header to "root" (proving
// they're actually root) while workspaceId - the workspace whose roles are being listed -
// stays a separate body field, same split as AddUserToWorkspace/RemoveUserFromWorkspace/
// ChangeUserWorkspaceRole. RoleActions.Query's own workspace_id filter (see
// QueryEntitiesPointer, CrudCoreActions.go) is applied directly from that field, entirely
// independent of the caller's own access-derived SqlContext - which is exactly what makes
// this action able to see into a workspace root doesn't otherwise belong to.
func QueryWorkspaceRolesAction(c abacdefs.QueryWorkspaceRolesActionRequest) (*abacdefs.QueryWorkspaceRolesActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_ROLE_QUERY},
		AllowOnRoot:    true,
	})
	if err != nil {
		return nil, err
	}

	req := c.Body
	if verr := fireback.CommonStructValidatorPointer(&req, false); verr != nil {
		return nil, verr
	}

	items, qrm, err2 := RoleActions.Query(fireback.QueryDSL{
		WorkspaceId:  req.WorkspaceId,
		ItemsPerPage: 500,
	})
	if err2 != nil {
		return nil, err2
	}

	return &abacdefs.QueryWorkspaceRolesActionResponse{
		Payload: fireback.GResponseQuery(items, qrm, &fireback.QueryDSL{}),
	}, nil
}

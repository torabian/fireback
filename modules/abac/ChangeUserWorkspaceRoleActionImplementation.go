package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// ChangeUserWorkspaceRoleAction replaces an existing member's role assignment(s) in a
// workspace with a single new role. Root only (AllowOnRoot, same idiom as
// AddUserToWorkspaceAction/RemoveUserFromWorkspaceAction).
//
// roleId is validated against workspaceId (the target workspace), not against root's own
// roles - list the real options with a `role browse`/`GET /role/browse` call using a
// Workspace-Id header set to workspaceId instead of root, exactly like
// WorkspaceAddUserDrawer's role picker already does on the frontend.
func ChangeUserWorkspaceRoleAction(c abacdefs.ChangeUserWorkspaceRoleActionRequest) (*abacdefs.ChangeUserWorkspaceRoleActionResponse, error) {
	queryPtr, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_CREATE},
		AllowOnRoot:    true,
	})
	if err != nil {
		return nil, err
	}

	req := c.Body
	if verr := fireback.CommonStructValidatorPointer(&req, false); verr != nil {
		return nil, verr
	}

	// The membership/role rows belong to req.WorkspaceId, not query.WorkspaceId -
	// AllowOnRoot above already pinned query.WorkspaceId to "root".
	targetQuery := *queryPtr
	targetQuery.WorkspaceId = req.WorkspaceId

	// RealEscape, not a struct-by-example Where - see AddUserToWorkspaceActionImplementation
	// .go's matching lookup for why.
	membership := &abacdefs.UserWorkspaceEntity{}
	findErr := fireback.GetDbRef().
		Where(fireback.RealEscape("user_id = ?", req.UserId)).
		Where(fireback.RealEscape("workspace_id = ?", req.WorkspaceId)).
		First(membership).Error
	if findErr != nil {
		return nil, &fireback.IError{
			HttpCode: 404,
			Errors: []*fireback.IErrorItem{
				{
					Location: "userId",
					Message: &fireback.ErrorItem{
						"$":  "UserNotInWorkspace",
						"en": "This user is not a member of the given workspace - use AddUserToWorkspace instead.",
					},
				},
			},
		}
	}

	if _, roleErrors := ValidateRoleAndItsExistence(emigo.NullableOf(req.RoleId), targetQuery); len(roleErrors) != 0 {
		return nil, &fireback.IError{Errors: roleErrors}
	}

	// Replace, not append: every existing role assignment for this membership is
	// dropped first, so the caller ends up with exactly the one role they asked for -
	// matching the "change his role" (singular) mental model, and AddUserToWorkspace's
	// own one-role-per-membership behavior.
	if _, err2 := WorkspaceRoleActions.RemoveEnqueue(fireback.DeleteRequest{
		Query:          "user_workspace_id = " + membership.UniqueId,
		ForceImmediate: true,
	}, targetQuery); err2 != nil {
		return nil, err2
	}

	if _, err3 := WorkspaceRoleActions.Create(&abacdefs.WorkspaceRoleEntity{
		UserWorkspaceId: emigo.NullableOf(membership.UniqueId),
		RoleId:          emigo.NullableOf(req.RoleId),
		WorkspaceId:     emigo.NullableOf(req.WorkspaceId),
	}, targetQuery); err3 != nil {
		return nil, err3
	}

	return &abacdefs.ChangeUserWorkspaceRoleActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.ChangeUserWorkspaceRoleActionRes{
			UniqueId:    membership.UniqueId,
			UserId:      req.UserId,
			WorkspaceId: req.WorkspaceId,
			RoleId:      req.RoleId,
		}),
	}, nil
}

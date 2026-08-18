package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// RemoveUserFromWorkspaceAction removes an existing user's membership - and every role
// assignment they hold in it - from a workspace. Root only (AllowOnRoot, same idiom as
// AddUserToWorkspaceAction, which this undoes): the caller's own Workspace-Id header is
// forced to "root", so workspaceId (the workspace being removed *from*) has to be a
// separate body field.
func RemoveUserFromWorkspaceAction(c abacdefs.RemoveUserFromWorkspaceActionRequest) (*abacdefs.RemoveUserFromWorkspaceActionResponse, error) {
	queryPtr, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_WORKSPACE_DELETE},
		AllowOnRoot:    true,
	})
	if err != nil {
		return nil, err
	}

	req := c.Body
	if verr := fireback.CommonStructValidatorPointer(&req, false); verr != nil {
		return nil, verr
	}

	// The membership row (and its roles) belongs to req.WorkspaceId, not
	// query.WorkspaceId - AllowOnRoot above already pinned query.WorkspaceId to "root".
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
						"en": "This user is not a member of the given workspace.",
					},
				},
			},
		}
	}

	// Safety guard: never let a root caller remove *its own* membership from the root
	// workspace - that's the specific, concrete way this action could permanently
	// lock an admin out with no way back in (a stray leftover member elsewhere in
	// root doesn't help the session that's mid-request right now).
	//
	// This intentionally checks identity, not a workspace-wide headcount: an earlier
	// version counted every userWorkspace row with workspace_id = "root" and
	// rejected only when that count was <= 1 - which sounds equivalent but isn't.
	// Once any *other* user has ever been added to root (an entirely normal thing to
	// happen over the life of a real deployment, or even within a single test suite
	// run that invites a second root member along the way), the count guard waves
	// the removal through - it doesn't matter that those other rows have nothing to
	// do with the caller: the caller's own session still loses root access the
	// moment its own row is gone. Verified the hard way once already.
	if req.WorkspaceId == ROOT_VAR && req.UserId == queryPtr.UserId {
		return nil, &fireback.IError{
			HttpCode: 400,
			Errors: []*fireback.IErrorItem{
				{
					Location: "userId",
					Message: &fireback.ErrorItem{
						"$":  "CannotRemoveOwnRootMembership",
						"en": "You cannot remove your own membership from the root workspace.",
					},
				},
			},
		}
	}

	if _, err2 := WorkspaceRoleActions.RemoveEnqueue(fireback.DeleteRequest{
		Query:          "user_workspace_id = " + membership.UniqueId,
		ForceImmediate: true,
	}, targetQuery); err2 != nil {
		return nil, err2
	}

	if _, err3 := UserWorkspaceActions.RemoveEnqueue(fireback.DeleteRequest{
		Query:          "unique_id = " + membership.UniqueId,
		ForceImmediate: true,
	}, targetQuery); err3 != nil {
		return nil, err3
	}

	return &abacdefs.RemoveUserFromWorkspaceActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.RemoveUserFromWorkspaceActionRes{
			UniqueId:    membership.UniqueId,
			UserId:      req.UserId,
			WorkspaceId: req.WorkspaceId,
		}),
	}, nil
}

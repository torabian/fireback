package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"gorm.io/gorm"
)

// AddUserToWorkspaceAction adds an existing user straight into a workspace with one of
// that workspace's existing roles - no invitation, no email/SMS involved. It's the
// direct counterpart to InviteToWorkspaceAction for the case where the target person
// already has an account: root picks a workspace (workspace archive list "add user"
// button), searches for the user, picks one of that workspace's roles, and submits.
//
// AllowOnRoot mirrors every other action in this file tree documented as "(root only)"
// (UserCreateAction, UserWorkspaceCreateAction, ...): it forces the caller's currently
// selected Workspace-Id header to be "root" (see WithAuthorizationPureDefault), so only
// an actual root user can call this - the workspaceId being added *to* is a separate
// field on the request body, since the caller's own context workspace is always root.
func AddUserToWorkspaceAction(c abacdefs.AddUserToWorkspaceActionRequest) (*abacdefs.AddUserToWorkspaceActionResponse, error) {
	queryPtr, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_WORKSPACE_CREATE},
		AllowOnRoot:    true,
	})
	if err != nil {
		return nil, err
	}

	req := c.Body

	if verr := fireback.CommonStructValidatorPointer(&req, false); verr != nil {
		return nil, verr
	}

	// The membership row (and its role) belongs to req.WorkspaceId, not query.WorkspaceId
	// - AllowOnRoot above already pinned query.WorkspaceId to "root".
	targetQuery := *queryPtr
	targetQuery.WorkspaceId = req.WorkspaceId

	if ws := GetWorkspaceByUniqueId(req.WorkspaceId); ws == nil || ws.UniqueId != req.WorkspaceId {
		return nil, &fireback.IError{
			HttpCode: 404,
			Errors: []*fireback.IErrorItem{
				{
					Location: "workspaceId",
					Message: &fireback.ErrorItem{
						"$":  "WorkspaceNotFound",
						"en": "The given workspace does not exist.",
					},
				},
			},
		}
	}

	if _, userErr := UserActions.GetOne(fireback.QueryDSL{UniqueId: req.UserId}); userErr != nil {
		return nil, &fireback.IError{
			HttpCode: 404,
			Errors: []*fireback.IErrorItem{
				{
					Location: "userId",
					Message: &fireback.ErrorItem{
						"$":  "UserNotFound",
						"en": "The given user does not exist.",
					},
				},
			},
		}
	}

	if _, roleErrors := ValidateRoleAndItsExistence(emigo.NullableOf(req.RoleId), targetQuery); len(roleErrors) != 0 {
		return nil, &fireback.IError{Errors: roleErrors}
	}

	// Give a clean, explicit error instead of a raw duplicate-row/constraint-violation
	// surprise (UserWorkspaceEntity does carry a db-level unique(user_id,
	// workspace_id) index - see the "userWorkspace" Module3 entity's gormMap - but a
	// dedicated field error here is a much better experience than whatever
	// GormErrorToIError makes of a raw constraint violation).
	//
	// RealEscape, not .Where(&abacdefs.UserWorkspaceEntity{...}) (a struct passed by
	// example): confirmed the hard way (see RemoveUserFromWorkspaceActionImplementation
	// .go's root-member-count guard) that GORM's struct-condition builder doesn't
	// reliably filter on emigo.Nullable[string]-typed fields - matching the raw-SQL
	// convention every other WorkspaceId/UserId-scoped query in this codebase already
	// uses (QueryEntitiesPointer, GetOneByWorkspaceEntity, ...) avoids that risk
	// entirely rather than depending on it happening to work.
	existing := &abacdefs.UserWorkspaceEntity{}
	existingErr := fireback.GetDbRef().
		Where(fireback.RealEscape("user_id = ?", req.UserId)).
		Where(fireback.RealEscape("workspace_id = ?", req.WorkspaceId)).
		First(existing).Error
	if existingErr == nil {
		return nil, &fireback.IError{
			HttpCode: 400,
			Errors: []*fireback.IErrorItem{
				{
					Location: "userId",
					Message: &fireback.ErrorItem{
						"$":  "UserAlreadyInWorkspace",
						"en": "This user already belongs to the given workspace.",
					},
				},
			},
		}
	} else if !gormRecordNotFound(existingErr) {
		return nil, fireback.GormErrorToIError(existingErr)
	}

	createdUserWorkspace, err2 := UserWorkspaceActions.Create(&abacdefs.UserWorkspaceEntity{
		UniqueId:    fireback.UUID(),
		UserId:      emigo.NullableOf(req.UserId),
		WorkspaceId: emigo.NullableOf(req.WorkspaceId),
	}, targetQuery)
	if err2 != nil {
		return nil, err2
	}

	if _, err3 := WorkspaceRoleActions.Create(&abacdefs.WorkspaceRoleEntity{
		UserWorkspaceId: emigo.NullableOf(createdUserWorkspace.UniqueId),
		RoleId:          emigo.NullableOf(req.RoleId),
		WorkspaceId:     emigo.NullableOf(req.WorkspaceId),
	}, targetQuery); err3 != nil {
		return nil, err3
	}

	return &abacdefs.AddUserToWorkspaceActionResponse{
		Payload: fireback.GResponseSingleItem(abacdefs.AddUserToWorkspaceActionRes{
			UniqueId:    createdUserWorkspace.UniqueId,
			UserId:      req.UserId,
			WorkspaceId: req.WorkspaceId,
			RoleId:      req.RoleId,
		}),
	}, nil
}

func gormRecordNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}

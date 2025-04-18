package abac

import (
	"github.com/torabian/fireback/modules/workspaces"
	"gorm.io/gorm"
)

func init() {
	// Override the implementation with our actual code.
	AcceptInviteActionImp = AcceptInviteAction
}
func AcceptInviteAction(
	req *AcceptInviteActionReqDto,
	q workspaces.QueryDSL) (string,
	*workspaces.IError,
) {

	// First of all, we will find the invitation and gather some information.
	q.UniqueId = req.InvitationUniqueId
	q.Deep = true
	invite, err := WorkspaceInviteActions.GetOne(q)

	if err != nil {
		return "", err
	}

	if invite == nil {
		return "", workspaces.Create401Error(&AbacMessages.InvitationNotFound, []string{})
	}

	err2d := workspaces.GetDbRef().Transaction(func(tx *gorm.DB) error {

		q.Tx = tx
		// In order to add a user to a workspace, we need to know the role which he will have.
		// there for, adding a workspaceuser entity in the database, and deleting the invitation is enough.

		q.WorkspaceId = invite.WorkspaceId.String
		uw, uwErr := UserWorkspaceActions.Create(&UserWorkspaceEntity{
			WorkspaceId: invite.WorkspaceId,

			// Think about here, user maybe exists or maybe doesn't exist.
			UserId: invite.UserId,
		}, q)

		if uwErr != nil {
			return uwErr
		}

		wre := &WorkspaceRoleEntity{
			UserWorkspaceId: workspaces.NewString(uw.UniqueId),
			RoleId:          invite.RoleId,
			WorkspaceId:     invite.WorkspaceId,
		}

		q.WorkspaceId = invite.WorkspaceId.String
		if _, wrErr := WorkspaceRoleActions.Create(wre, q); err != wrErr {
			return wrErr
		}

		q.UniqueId = req.InvitationUniqueId
		q.Query = "unique_id = " + q.UniqueId
		_, errRemove := WorkspaceInviteActions.Remove(q)

		if errRemove != nil {
			return errRemove
		}

		return nil
	})

	if err2d != nil {
		return "", err2d.(*workspaces.IError)
	}

	// Implement the logic here.
	return "", nil
}

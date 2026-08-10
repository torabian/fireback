package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
)

func AcceptInviteAction(c abacdefs.AcceptInviteActionRequest) (*abacdefs.AcceptInviteActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

	req := c.Body

	// First of all, we will find the invitation and gather some information.
	//
	// Bug fix: this used to reuse the outer `err` (declared above as the plain `error`
	// interface ResolveActionContext returns) for GetOne's *fireback.IError result too.
	// Reusing one variable name across two differently-typed fallible calls in the same
	// scope doesn't change its static type - `err` stayed `error`, so GetOne's nil
	// *fireback.IError success value got boxed into a non-nil `error` interface (the
	// classic typed-nil-in-interface gotcha), making `if err != nil` true unconditionally
	// and rejecting every single accept attempt right after successfully finding the
	// invite. A fresh, concretely-typed variable (as every other *fireback.IError-returning
	// call in this file already uses - see uwErr/wrErr/errRemove below) avoids the boxing
	// entirely.
	//
	// Bug fix: GetOneEntity's not-found path returns (&<zero value>, GormErrorToIError(err))
	// - a non-nil *WorkspaceInviteEntity pointing at an empty struct, paired with a
	// generic "ResourceNotFound" error - never a nil entity. That made the "if invite ==
	// nil" friendly-message check below dead code: a bad/expired invitationUniqueId
	// always hit "if ierr != nil" first and surfaced the generic message instead of the
	// specific AbacMessages.InvitationNotFound one. Both conditions now map to the same,
	// intended response.
	q.UniqueId = req.InvitationUniqueId
	q.Deep = true
	invite, ierr := WorkspaceInviteActions.GetOne(q)

	if ierr != nil || invite == nil {
		return nil, fireback.Create401Error(&AbacMessages.InvitationNotFound, []string{})
	}

	err2d := fireback.GetDbRef().Transaction(func(tx *gorm.DB) error {

		q.Tx = tx
		// In order to add a user to a workspace, we need to know the role which he will have.
		// there for, adding a workspaceuser entity in the database, and deleting the invitation is enough.

		q.WorkspaceId = invite.WorkspaceId.OrDefault("")

		// Bug fix: this used invite.UserId (WorkspaceInviteEntity's own doc comment: "the
		// unique-id of the user which created/owns the record", i.e. the inviter - and in
		// practice always null anyway, since InviteToWorkspaceActionImplementation.go never
		// sets it). The person who should actually be joining the workspace is whoever is
		// authenticated and calling this action right now to accept it - q.UserId, resolved
		// by ResolveActionContext above from their own session, regardless of whether they
		// already existed or just signed up to accept this invite.
		uw, uwErr := UserWorkspaceActions.Create(&abacdefs.UserWorkspaceEntity{
			WorkspaceId: invite.WorkspaceId,
			UserId:      emigo.NullableOf(q.UserId),
		}, q)

		if uwErr != nil {
			return uwErr
		}

		wre := &abacdefs.WorkspaceRoleEntity{
			UserWorkspaceId: emigo.NullableOf(uw.UniqueId),
			RoleId:          invite.RoleId,
			WorkspaceId:     invite.WorkspaceId,
		}

		q.WorkspaceId = invite.WorkspaceId.OrDefault("")
		if _, wrErr := WorkspaceRoleActions.Create(wre, q); wrErr != nil {
			return wrErr
		}

		q.UniqueId = req.InvitationUniqueId
		q.Query = "unique_id = " + q.UniqueId
		_, errRemove := WorkspaceInviteActions.RemoveEnqueue(fireback.DeleteRequest{
			Query:          "unique_id = " + q.UniqueId,
			ForceImmediate: true,
		}, q)

		if errRemove != nil {
			return errRemove
		}

		return nil
	})

	if err2d != nil {
		return nil, err2d.(*fireback.IError)
	}

	return &abacdefs.AcceptInviteActionResponse{
		Payload: fireback.GResponseSingleItem(&abacdefs.AcceptInviteActionRes{Accepted: true}),
	}, nil
}

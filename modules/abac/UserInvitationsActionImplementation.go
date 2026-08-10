package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func UserInvitationsAction(c abacdefs.UserInvitationsActionRequest) (*abacdefs.UserInvitationsActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

	invitations, qrm, err3 := UserInvitationsQuery(q)

	if err3 != nil {
		return nil, fireback.CastToIError(err3)
	}

	return &abacdefs.UserInvitationsActionResponse{
		Payload: fireback.GResponseQuery(invitations, qrm, &q),
	}, nil
}

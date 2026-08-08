package abac

import "github.com/torabian/fireback/modules/fireback"

func UserInvitationsAction(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error) {
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

	return &UserInvitationsActionResponse{
		Payload: fireback.GResponseQuery(invitations, qrm, &q),
	}, nil
}

package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	UserInvitationsImpl = UserInvitationsAction
}

func UserInvitationsAction(c UserInvitationsActionRequest, q fireback.QueryDSL) (*UserInvitationsActionResponse, error) {

	invitations, qrm, err3 := UserInvitationsQuery(q)

	if err3 != nil {
		return nil, fireback.CastToIError(err3)
	}

	return &UserInvitationsActionResponse{
		Payload: fireback.GResponseQuery(invitations, qrm, &q),
	}, nil
}

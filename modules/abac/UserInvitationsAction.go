package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	UserInvitationsImpl = UserInvitationsAction
}
func UserInvitationsAction(c UserInvitationsActionRequest, query fireback.QueryDSL) (*UserInvitationsActionResponse, error) {

	invitations, qrm, err := UserInvitationsQuery(query)
	if err != nil {
		return nil, fireback.CastToIError(err)
	}

	return &UserInvitationsActionResponse{
		Payload: fireback.GResponseQuery(invitations, qrm, &query),
	}, nil
}

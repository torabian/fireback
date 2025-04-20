package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	UserInvitationsActionImp = UserInvitationsAction
}
func UserInvitationsAction(
	q fireback.QueryDSL) ([]*UserInvitationsQueryColumns,
	*fireback.QueryResultMeta,
	*fireback.IError,
) {

	invitations, _, err3 := UserInvitationsQuery(q)

	if err3 != nil {
		return nil, nil, fireback.CastToIError(err3)
	}

	return invitations, nil, nil
}

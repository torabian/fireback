package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	UserInvitationsActionImp = UserInvitationsAction
}
func UserInvitationsAction(
	q workspaces.QueryDSL) ([]*UserInvitationsQueryColumns,
	*workspaces.QueryResultMeta,
	*workspaces.IError,
) {

	invitations, _, err3 := UserInvitationsQuery(q)

	if err3 != nil {
		return nil, nil, workspaces.CastToIError(err3)
	}

	return invitations, nil, nil
}

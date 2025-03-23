package workspaces

func init() {
	// Override the implementation with our actual code.
	UserInvitationsActionImp = UserInvitationsAction
}
func UserInvitationsAction(
	q QueryDSL) ([]*UserInvitationsQueryColumns,
	*QueryResultMeta,
	*IError,
) {

	invitations, _, err3 := UserInvitationsQuery(q)

	if err3 != nil {
		return nil, nil, CastToIError(err3)
	}

	return invitations, nil, nil
}

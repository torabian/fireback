package workspaces

func init() {
	// Override the implementation with our actual code.

	UserPassportsActionImp = UserPassportsAction
}
func UserPassportsAction(q QueryDSL) ([]*UserPassportsActionResDto, *QueryResultMeta, *IError) {
	// Implement the logic here.

	// Passports all belong to root workspace, so we need to query that
	// thats why it's changed manually here. Passport needs to belong to current user.
	q.Query = "user_id = " + q.UserId
	q.WorkspaceId = ROOT_VAR
	passports, _, err := PassportActions.Query(q)
	if err != nil {
		return nil, nil, CastToIError(err)
	}

	result := []*UserPassportsActionResDto{}
	for _, item := range passports {
		result = append(result, &UserPassportsActionResDto{
			Value:         item.Value,
			Type:          item.Type,
			UniqueId:      item.UniqueId,
			TotpConfirmed: item.TotpConfirmed,
		})
	}

	return result, nil, nil
}

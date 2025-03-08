package workspaces

func init() {
	// Override the implementation with our actual code.
	ChangePasswordActionImp = ChangePasswordAction
}

func ChangePasswordAction(req *ChangePasswordActionReqDto, q QueryDSL) (string, *IError) {

	if err := ChangePasswordActionReqValidator(req); err != nil {
		return "", err
	}

	// We need to give the administrator option to set rules to figure out what passwords are safe
	// and what are weak. By default, a 6 character password is enough.
	if len(req.Password) < 6 {
		return "", Create401Error(&WorkspacesMessages.PasswordDoesNotMeetTheSecurityRequirement, []string{})
	}

	// Passports all belong to root workspace, so we need to query that
	// thats why it's changed manually here. Passport needs to belong to current user.
	q.Query = "value = " + req.Value + " and user_id = " + q.UserId
	q.WorkspaceId = ROOT_VAR
	passports, _, err := PassportActions.Query(q)
	if err != nil {
		return "", CastToIError(err)
	}

	if len(passports) == 0 {
		return "", Create401Error(&WorkspacesMessages.PassportNotFound, []string{})
	}

	previousPassword := passports[0].Password

	passwordHashed, err1 := HashPassword(req.Password)
	if err1 != nil {
		return "", CastToIError(err)
	}

	updated, err2 := PassportActions.Update(q, &PassportEntity{
		Value:    req.Value,
		Password: passwordHashed,
		UniqueId: passports[0].UniqueId,
	})
	if err2 != nil {
		return "", CastToIError(err)
	}

	if updated.Password == previousPassword {
		return "", Create401Error(&WorkspacesMessages.PasswordDidNotUpdated, []string{})
	}

	return "Password changed", nil
}

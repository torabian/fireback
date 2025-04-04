package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	ChangePasswordActionImp = ChangePasswordAction
}

func ChangePasswordAction(req *ChangePasswordActionReqDto, q workspaces.QueryDSL) (string, *workspaces.IError) {

	if err := ChangePasswordActionReqValidator(req); err != nil {
		return "", err
	}

	// We need to give the administrator option to set rules to figure out what passwords are safe
	// and what are weak. By default, a 6 character password is enough.
	if len(req.Password) < 6 {
		return "", workspaces.Create401Error(&AbacMessages.PasswordDoesNotMeetTheSecurityRequirement, []string{})
	}

	// Passports all belong to root workspace, so we need to query that
	// thats why it's changed manually here. Passport needs to belong to current user.
	passports := []PassportEntity{}
	err := workspaces.GetRef(q).Where(PassportEntity{UserId: workspaces.NewString(q.UserId)}).Find(&passports).Error
	if err != nil {
		return "", workspaces.CastToIError(err)
	}

	if len(passports) == 0 {
		return "", workspaces.Create401Error(&AbacMessages.PassportNotFound, []string{})
	}

	previousPassword := passports[0].Password

	passwordHashed, err1 := workspaces.HashPassword(req.Password)
	if err1 != nil {
		return "", workspaces.CastToIError(err)
	}

	updated, err2 := PassportActions.Update(q, &PassportEntity{
		Password: passwordHashed,
		UniqueId: passports[0].UniqueId,
	})
	if err2 != nil {
		return "", workspaces.CastToIError(err)
	}

	if updated.Password == previousPassword {
		return "", workspaces.Create401Error(&AbacMessages.PasswordDidNotUpdated, []string{})
	}

	return "Password changed", nil
}

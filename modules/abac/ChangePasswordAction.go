package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	ChangePasswordActionImp = ChangePasswordAction
}

func ChangePasswordAction(req *ChangePasswordActionReqDto, q fireback.QueryDSL) (string, *fireback.IError) {

	if err := ChangePasswordActionReqValidator(req); err != nil {
		return "", err
	}

	// We need to give the administrator option to set rules to figure out what passwords are safe
	// and what are weak. By default, a 6 character password is enough.
	if len(req.Password) < 6 {
		return "", fireback.Create401Error(&AbacMessages.PasswordDoesNotMeetTheSecurityRequirement, []string{})
	}

	// Passports all belong to root workspace, so we need to query that
	// thats why it's changed manually here. Passport needs to belong to current user.
	passports := []PassportEntity{}
	err := fireback.GetRef(q).Where(PassportEntity{UserId: fireback.NewString(q.UserId)}).Find(&passports).Error
	if err != nil {
		return "", fireback.CastToIError(err)
	}

	if len(passports) == 0 {
		return "", fireback.Create401Error(&AbacMessages.PassportNotFound, []string{})
	}

	previousPassword := passports[0].Password

	passwordHashed, err1 := fireback.HashPassword(req.Password)
	if err1 != nil {
		return "", fireback.CastToIError(err)
	}

	updated, err2 := PassportActions.Update(q, &PassportEntity{
		Password: passwordHashed,
		UniqueId: passports[0].UniqueId,
	})
	if err2 != nil {
		return "", fireback.CastToIError(err)
	}

	if updated.Password == previousPassword {
		return "", fireback.Create401Error(&AbacMessages.PasswordDidNotUpdated, []string{})
	}

	return "Password changed", nil
}

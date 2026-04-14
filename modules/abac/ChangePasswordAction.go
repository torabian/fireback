package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	ChangePasswordImpl = ChangePasswordAction
}

func ChangePasswordAction(c ChangePasswordActionRequest, q fireback.QueryDSL) (*ChangePasswordActionResponse, error) {
	req := c.Body

	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	// We need to give the administrator option to set rules to figure out what passwords are safe
	// and what are weak. By default, a 6 character password is enough.
	if len(req.Password) < 6 {
		return nil, fireback.Create401Error(&AbacMessages.PasswordDoesNotMeetTheSecurityRequirement, []string{})
	}

	// Passports all belong to root workspace, so we need to query that
	// thats why it's changed manually here. Passport needs to belong to current user.
	passports := []PassportEntity{}
	err := fireback.GetRef(q).Where(PassportEntity{UserId: fireback.NewString(q.UserId)}).Find(&passports).Error
	if err != nil {
		return nil, fireback.CastToIError(err)
	}

	if len(passports) == 0 {
		return nil, fireback.Create401Error(&AbacMessages.PassportNotFound, []string{})
	}

	previousPassword := passports[0].Password

	passwordHashed, err1 := fireback.HashPassword(req.Password)
	if err1 != nil {
		return nil, fireback.CastToIError(err)
	}

	updated, err2 := PassportActions.Update(q, &PassportEntity{
		Password: passwordHashed,
		UniqueId: passports[0].UniqueId,
	})
	if err2 != nil {
		return nil, fireback.CastToIError(err)
	}

	if updated.Password == previousPassword {
		return nil, fireback.Create401Error(&AbacMessages.PasswordDidNotUpdated, []string{})
	}

	return &ChangePasswordActionResponse{
		Payload: fireback.GResponseSingleItem(ChangePasswordActionRes{Changed: true}),
	}, nil
}

package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/security"
)

func ChangePasswordAction(c ChangePasswordActionRequest) (*ChangePasswordActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

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
	err2 := fireback.GetRef(q).Where(PassportEntity{UserId: emigo.NullableOf(q.UserId)}).Find(&passports).Error
	if err2 != nil {
		return nil, fireback.CastToIError(err2)
	}

	if len(passports) == 0 {
		return nil, fireback.Create401Error(&AbacMessages.PassportNotFound, []string{})
	}

	previousPassword := passports[0].Password

	passwordHashed, err1 := security.HashPassword(req.Password)
	if err1 != nil {
		return nil, fireback.CastToIError(err1)
	}

	updated, err3 := PassportActions.Update(q, &PassportEntity{
		Password: passwordHashed,
		UniqueId: passports[0].UniqueId,
	})
	if err3 != nil {
		return nil, fireback.CastToIError(err3)
	}

	if updated.Password == previousPassword {
		return nil, fireback.Create401Error(&AbacMessages.PasswordDidNotUpdated, []string{})
	}

	return &ChangePasswordActionResponse{
		Payload: fireback.GResponseSingleItem(ChangePasswordActionRes{Changed: true}),
	}, nil
}

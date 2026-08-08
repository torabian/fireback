package abac

import (
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/security"
)

func PassportAppendEmailToUser(dto *ClassicAuthDto, query fireback.QueryDSL) (*PassportEntity, *fireback.IError) {
	passwordHashed, _ := security.HashPassword(dto.Password)

	if iError := UserWithEmailAndPasswordValidator(dto, false); iError != nil {
		return nil, iError
	}

	entity := &PassportEntity{
		Value: dto.Value,
		// Confirmed: 1,
		Password: passwordHashed,
		Type:     PassportTypes.EmailPassword,
	}

	return PassportActions.Create(entity, query)

}

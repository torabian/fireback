package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/security"
)

func PassportAppendEmailToUser(dto *abacdefs.ClassicAuthDto, query fireback.QueryDSL) (*abacdefs.PassportEntity, *fireback.IError) {
	passwordHashed, _ := security.HashPassword(dto.Password)

	if iError := UserWithEmailAndPasswordValidator(dto, false); iError != nil {
		return nil, iError
	}

	entity := &abacdefs.PassportEntity{
		Value: dto.Value,
		// Confirmed: 1,
		Password: passwordHashed,
		Type:     PassportTypes.EmailPassword,
	}

	return PassportActions.Create(entity, query)

}

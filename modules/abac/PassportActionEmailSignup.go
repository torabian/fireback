package abac

import "github.com/torabian/fireback/modules/workspaces"

func PassportAppendEmailToUser(dto *ClassicAuthDto, query workspaces.QueryDSL) (*PassportEntity, *workspaces.IError) {
	passwordHashed, _ := workspaces.HashPassword(dto.Password)

	if iError := UserWithEmailAndPasswordValidator(dto, false); iError != nil {
		return nil, iError
	}

	entity := &PassportEntity{
		Value: dto.Value,
		// Confirmed: 1,
		Password: passwordHashed,
		Type:     PassportTypes.EmailPassword,
	}

	return PassportActionCreateFn(entity, query)

}

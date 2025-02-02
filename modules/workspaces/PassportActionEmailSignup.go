package workspaces

func PassportAppendEmailToUser(dto *ClassicAuthDto, query QueryDSL) (*PassportEntity, *IError) {
	passwordHashed, _ := HashPassword(*dto.Password)

	if iError := UserWithEmailAndPasswordValidator(dto, false); iError != nil {
		return nil, iError
	}

	entity := &PassportEntity{
		Value: dto.Value,
		// Confirmed: 1,
		Password: &passwordHashed,
		Type:     &PassportTypes.EmailPassword,
	}

	return PassportActionCreateFn(entity, query)

}

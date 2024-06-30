package workspaces

func UserSigninEmailAndPasswordValidator(dto *EmailAccountSigninDto, isPatch bool) *IError {
	return CommonStructValidatorPointer(dto, isPatch)
}

func UserWithEmailAndPasswordValidator(dto *ClassicAuthDto, isPatch bool) *IError {
	return CommonStructValidatorPointer(dto, isPatch)
}

func UserWithPhoneValidator(dto PhoneNumberAccountCreationDto, isPatch bool) *IError {
	return CommonStructValidatorPointer(&dto, isPatch)
}

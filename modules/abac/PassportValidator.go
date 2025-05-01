package abac

import "github.com/torabian/fireback/modules/fireback"

func UserSigninEmailAndPasswordValidator(dto *EmailAccountSigninDto, isPatch bool) *fireback.IError {
	return fireback.CommonStructValidatorPointer(dto, isPatch)
}

func UserWithEmailAndPasswordValidator(dto *ClassicAuthDto, isPatch bool) *fireback.IError {
	return fireback.CommonStructValidatorPointer(dto, isPatch)
}

func UserWithPhoneValidator(dto PhoneNumberAccountCreationDto, isPatch bool) *fireback.IError {
	return fireback.CommonStructValidatorPointer(&dto, isPatch)
}

package abac

import "github.com/torabian/fireback/modules/workspaces"

func UserSigninEmailAndPasswordValidator(dto *EmailAccountSigninDto, isPatch bool) *workspaces.IError {
	return workspaces.CommonStructValidatorPointer(dto, isPatch)
}

func UserWithEmailAndPasswordValidator(dto *ClassicAuthDto, isPatch bool) *workspaces.IError {
	return workspaces.CommonStructValidatorPointer(dto, isPatch)
}

func UserWithPhoneValidator(dto PhoneNumberAccountCreationDto, isPatch bool) *workspaces.IError {
	return workspaces.CommonStructValidatorPointer(&dto, isPatch)
}

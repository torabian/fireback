package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func UserSigninEmailAndPasswordValidator(dto *abacdefs.EmailAccountSigninDto, isPatch bool) *fireback.IError {
	return fireback.CommonStructValidatorPointer(dto, isPatch)
}

func UserWithEmailAndPasswordValidator(dto *abacdefs.ClassicAuthDto, isPatch bool) *fireback.IError {
	return fireback.CommonStructValidatorPointer(dto, isPatch)
}

func UserWithPhoneValidator(dto abacdefs.PhoneNumberAccountCreationDto, isPatch bool) *fireback.IError {
	return fireback.CommonStructValidatorPointer(&dto, isPatch)
}

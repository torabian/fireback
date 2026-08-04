package fireback

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func CommonStructValidatorPointer[T any](dto *T, isPatch bool) *IError {

	if dto == nil {
		return Create401Error(&FirebackMessages.BodyIsMissing, []string{})
	}

	var validate *validator.Validate = validator.New()

	err := validate.Struct(dto)

	errors := []*IErrorItem{}
	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {

			// Required fields when updating an entity are not required
			// to be validated
			if isPatch && err.ActualTag() == "required" {
				continue
			}

			t := strings.Replace(err.Error(), "Key: '", "", -1)
			t = t[0:strings.Index(t, "'")]

			t = t[strings.Index(t, ".")+1:]
			t = strings.ToLower(t[0:1]) + t[1:]
			t = dotToCamelCase(t)

			// Find out a way that I can translate messages, for example if the ActualTag is oneof,
			// in a way I can translate all

			errors = append(errors, &IErrorItem{
				Location:   t,
				ErrorParam: err.Param(),
				Message:    CastFieldErrorToErrorItem(err),
				Type:       err.Type().String(),
			})
		}

	}

	if len(errors) > 0 {
		var result IError = IError{
			Message:  FirebackMessages.ValidationFailedOnSomeFields,
			Errors:   errors,
			HttpCode: 403,
		}
		return &result
	}

	return nil
}

func SliceValidator[T any](items []*T, isPatch bool, prefix string) []*IErrorItem {
	errItems := []*IErrorItem{}

	for index, item := range items {
		err := CommonStructValidatorPointer(item, isPatch)

		if err != nil {
			for _, subError := range err.Errors {
				errItems = append(errItems, &IErrorItem{
					Location: prefix + "[" + fmt.Sprint(index) + "]." + subError.Location,
					Message:  subError.Message,
					Type:     subError.Type,
				})
			}
		}

	}

	return errItems
}

// FieldError is from validator library
// We need to complete this with translation somehow and I have no idea how
func CastFieldErrorToErrorItem(fe validator.FieldError) *ErrorItem {
	switch fe.Tag() {
	case "required":
		return &FirebackMessages.FieldRequired
	case "email":
		return &FirebackMessages.FieldInvalidEmail
	case "oneof":
		return &FirebackMessages.FieldOneOf
	}

	return &ErrorItem{
		"en": fe.Error(),
	}
}

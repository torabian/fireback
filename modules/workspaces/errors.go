package workspaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	// "github.com/flamego/validator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CreateIErrorString(msg string, list []string, code int32) *IError {

	er := []*IErrorItem{}

	for _, item := range list {
		er = append(er, &IErrorItem{Location: item})
	}

	result := IError{
		Message:  msg,
		Errors:   er,
		HttpCode: code,
	}

	return &result
}

func CastToIError(err error) *IError {
	if err == nil {
		return nil
	}

	return err.(*IError)
}

func IErrorFromString(err string) *IError {
	if err == "" {
		return nil
	}

	body := &IError{}
	uncastErr := json.Unmarshal([]byte(err), &body)

	if uncastErr != nil {
		return nil
	}

	return body
}
func IResponseFromString[T any](err string) *IResponse[T] {
	if err == "" {
		return nil
	}

	body := &IResponse[T]{}
	uncastErr := json.Unmarshal([]byte(err), &body)

	if uncastErr != nil {
		return nil
	}

	return body
}

func GormErrorToIError(err error) *IError {

	if err == nil {
		return nil
	}

	// Implement this, and translate all of the error message if needed.
	// Some messages should not go out, at all.
	msg := err.Error()
	var code int32 = 0

	if err == gorm.ErrRecordNotFound {
		msg = "NOT_FOUND"
		code = http.StatusNotFound
	}

	// var sqliteErr sqlite3.Error
	// errors.As(err, &sqliteErr)

	result := IError{
		Message:  msg,
		HttpCode: code,
	}

	return &result
}

func (r *IError) Error() string {
	str, _ := json.MarshalIndent(r, "", "  ")
	return string(str)
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

func AppendSliceErrors[T any](items []*T, isPatch bool, prefix string, err *IError) {
	if items == nil {
		return
	}

	subErrors := SliceValidator(items, isPatch, prefix)
	if len(subErrors) > 0 {
		if err == nil {
			err = &IError{}
		}

		err.Errors = append(err.Errors, subErrors...)
	}
}

func IsNilish(val any) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return false
}

func CommonStructValidatorPointer[T any](dto *T, isPatch bool) *IError {

	if dto == nil {
		return CreateIErrorString("ENTITY_NEEDED", []string{}, 403)
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

			errors = append(errors, &IErrorItem{
				Location:   t,
				ErrorParam: err.Param(),
				Message:    err.Tag(),
				Type:       err.Type().String(),
			})
		}

	}

	if len(errors) > 0 {
		var result IError = IError{
			Message: "VALIDATION_FAILED_ON_SOME_FIELDS",
			Errors:  errors,
		}
		return &result
	}

	return nil
}

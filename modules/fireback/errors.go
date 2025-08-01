package fireback

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/xeipuuv/gojsonschema"
	"gorm.io/gorm"
)

func Create401Error(msg *ErrorItem, list []string) *IError {
	result := IError{
		Message:  *msg,
		HttpCode: 401,
		MessageParams: map[string]interface{}{
			"list": list,
		},
	}

	return &result
}

func Create401ParamOnly(msg *ErrorItem, params map[string]interface{}) *IError {
	result := IError{
		Message:       *msg,
		HttpCode:      401,
		MessageParams: params,
	}

	return &result
}

func Create401ErrorWithItems(msg *ErrorItem, items []*IErrorItem) *IError {

	result := IError{
		Message:  *msg,
		HttpCode: 401,
		Errors:   items,
	}

	return &result
}

func CastToIError(err error) *IError {
	if err == nil {
		return nil
	}

	if ierr, ok := err.(*mysql.MySQLError); ok {
		return &IError{
			Message: ErrorItem{
				"en": ierr.Message,
			},
		}
	}

	if ierr, ok := err.(*IError); ok {
		return ierr
	}

	return &IError{
		Message: ErrorItem{
			"en": err.Error(),
		},
	}
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

	httpCode := http.StatusInternalServerError
	params := map[string]interface{}{}

	// For not found errors, we are going to tell it to client no matter production or not
	if errors.Is(err, gorm.ErrRecordNotFound) {
		httpCode = http.StatusNotFound
		return &IError{
			Message:  FirebackMessages.ResourceNotFound,
			HttpCode: int32(httpCode),
		}
	}

	// For other gorm errors, its necessary, that to understand there is a database level error,
	// and we hide the details on production

	if !config.Production {
		params["nativeError"] = err.Error()
	}

	httpCode = http.StatusBadRequest

	return &IError{
		Message:       FirebackMessages.DatabaseOperationError,
		MessageParams: params,
		HttpCode:      int32(httpCode),
	}

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

func toCamelCase(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// dotToCamelCase converts "person.FirstName.LastName" to "person.firstName.lastName"
func dotToCamelCase(input string) string {
	parts := strings.Split(input, ".")
	for i, part := range parts {
		parts[i] = toCamelCase(part)
	}
	return strings.Join(parts, ".")
}

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
			Message: FirebackMessages.ValidationFailedOnSomeFields,
			Errors:  errors,
		}
		return &result
	}

	return nil
}

func JsonSchemaToIError(result *gojsonschema.Result) *IError {
	if result.Valid() || len(result.Errors()) == 0 {
		return nil
	}

	err := &IError{}
	for _, er := range result.Errors() {
		d := er.Details()

		location := ""
		if msg, ok := d["property"].(string); ok {
			location = msg
		} else if msg, ok := d["field"].(string); ok {
			location = msg
		}
		err.Errors = append(err.Errors, &IErrorItem{
			Location: location,
			// Message:  er.Description(),
		})
	}

	return err
}

package workspaces

import (
	"encoding/json"
	"fmt"
	"log"
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
		MessageTranslated: err.Error(),
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

	message := map[string]string{
		"$": "Unknown error occured",
	}
	// Implement this, and translate all of the error message if needed.
	// Some messages should not go out, at all.
	// msg := err.Error()
	var code int32 = 500

	if err == gorm.ErrRecordNotFound {
		code = http.StatusNotFound
	} else if strings.Contains(err.Error(), "UNIQUE constraint") {

		code = 501
		message = map[string]string{
			"$": "Unique key violation, you cannot have duplicate unique keys",
		}
	} else {
		log.Default().Println("Unhandled gorm error", err)
	}

	result := IError{
		Message:  message,
		HttpCode: code,
	}

	return &result
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
func CastFieldErrorToErrorItem(err validator.FieldError) *ErrorItem {
	return &ErrorItem{
		"en": err.Tag(),
	}
}

func CommonStructValidatorPointer[T any](dto *T, isPatch bool) *IError {

	if dto == nil {
		return Create401Error(&WorkspacesMessages.BodyIsMissing, []string{})
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
				Message:    CastFieldErrorToErrorItem(err),
				Type:       err.Type().String(),
			})
		}

	}

	if len(errors) > 0 {
		var result IError = IError{
			Message: WorkspacesMessages.ValidationFailedOnSomeFields,
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

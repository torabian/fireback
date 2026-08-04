// Package errors provides a two-level error handling mechanism, featuring type-safe error messages
// and a common envelope for errors. At its core is the IError struct, which actions can create
// instead of a standard error. This struct can contain nested errors with detailed field validations.
// This file gathers all related error types and definitions.
//
// Note: This file has no dependencies.
package fireback

import (
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/torabian/fireback/modules/fireback/ferror"
	"gorm.io/gorm"
)

/*
* This could be used instead of return 'error'
* It's based on google json styleguide, and we can append more details if it's required
* Modification is, each string representation should be translated based on accept-language
 */
// Same-name aliases: every existing fireback.IError / fireback.ErrorItem
// reference across modules/abac's .dyno.go files, etc. keeps compiling
// unchanged, since these aren't new types — just other names for
// ferror.Error / ferror.ErrorItem.
type IError = ferror.Error
type IErrorItem = ferror.FieldError
type IPublicError = ferror.PublicError
type IPublicErrorItem = ferror.PublicFieldError
type ErrorItem = ferror.ErrorItem

var Create401Error = ferror.Create401Error
var Create401ParamOnly = ferror.Create401ParamOnly
var Create401ErrorWithItems = ferror.Create401ErrorWithItems

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

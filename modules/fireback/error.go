// Package errors provides a two-level error handling mechanism, featuring type-safe error messages
// and a common envelope for errors. At its core is the IError struct, which actions can create
// instead of a standard error. This struct can contain nested errors with detailed field validations.
// This file gathers all related error types and definitions.
//
// Note: This file has no dependencies.
package fireback

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/*
* This could be used instead of return 'error'
* It's based on google json styleguide, and we can append more details if it's required
* Modification is, each string representation should be translated based on accept-language
 */

type IError struct {
	Message           ErrorItem              `json:"message,omitempty"`
	MessageParams     map[string]interface{} `json:"messageParams,omitempty"`
	MessageTranslated string                 `json:"messageTranslated,omitempty"`
	Errors            []*IErrorItem          `json:"errors,omitempty"`
	HttpCode          int32                  `json:"httpCode,omitempty"`
}

func (x *IError) Is(value string) bool {
	if x.Message == nil {
		return false
	}

	code := x.Message["$"]

	if code == value {
		return true
	}

	return false
}

func (x *IError) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}

	return ""
}

/*
* A public representation of the IError. IError should not be printed to the public,
* users, due to possibility of it's sensetive information. Therefor, we cast IError to IPublicError
* before sending it through http, cli, etc.
 */
type IPublicError struct {
	Message           string                 `json:"message,omitempty"`
	MessageTranslated string                 `json:"messageTranslated,omitempty"`
	MessageParams     map[string]interface{} `json:"messageParams,omitempty"`
	Errors            []*IPublicErrorItem    `json:"errors,omitempty"`
	HttpCode          int32                  `json:"httpCode,omitempty"`
}

// IPublicErrorItem represents an error item that can be used to convey specific
// error details to the end user. For example, in a form, the location could point to
// `users[0].name`, indicating that the first user's name field is invalid.
// This structure is part of the IPublicError, which is used to safely expose error
// information to the public, as opposed to internal error handling structures like IErrorItem.
type IPublicErrorItem struct {
	// Location of the error in the request data, such as a specific field or parameter.
	Location string `json:"location,omitempty"`
	// Message is the default error message.
	Message string `json:"message,omitempty"`
	// MessageTranslated is the error message translated to the user's language.
	MessageTranslated string `json:"messageTranslated,omitempty"`
	// ErrorParam contains any additional parameters related to the error.
	ErrorParam string `json:"errorParam,omitempty"`
	// Type categorizes the error for easier identification and handling.
	Type string `json:"type,omitempty"`
}

type IErrorItem struct {
	Location   string     `json:"location,omitempty"`
	Message    *ErrorItem `json:"message,omitempty"`
	ErrorParam string     `json:"errorParam,omitempty"`
	Type       string     `json:"type,omitempty"`
}

func ReplacePlaceholders(input string, values map[string]interface{}) string {
	if values == nil {
		return input
	}
	for key, val := range values {
		placeholder := "%" + key

		var strVal string
		switch v := val.(type) {
		case string:
			strVal = v
		case int, int64, float64, float32:
			strVal = fmt.Sprintf("%v", v)
		case bool:
			strVal = strconv.FormatBool(v)
		default:
			strVal = fmt.Sprintf("%v", v)
		}

		input = strings.ReplaceAll(input, placeholder, strVal)
	}
	return input
}

// ToPublicEndUser converts an IError to an IPublicError.
// Ensure you do not return IError directly to the public to avoid exposing internal error details.
//
// Parameters:
//   - q: An interface with a GetLanguage method that returns a string representing the language code. QueryDSL already has this interface so you can pass that instead
//
// Returns:
//   - A pointer to an IPublicError containing translated messages and the HTTP code.
func (r *IError) ToPublicEndUser(q interface {
	GetLanguage() string
}) *IPublicError {

	// Retrieve the language code
	lang := q.GetLanguage()
	if lang == "" {
		lang = "en" // or any default language code
	}

	err := &IPublicError{}
	err.HttpCode = r.HttpCode
	err.MessageTranslated = ReplacePlaceholders(r.Message[lang], r.MessageParams)
	err.MessageParams = r.MessageParams
	err.Message = r.Message["$"]

	for _, item := range r.Errors {
		msg := (*item.Message)[lang]

		if item.ErrorParam != "" {
			msg = strings.ReplaceAll(msg, "%s", item.ErrorParam)
		}
		err.Errors = append(err.Errors, &IPublicErrorItem{
			Location:          item.Location,
			ErrorParam:        item.ErrorParam,
			Type:              item.Type,
			MessageTranslated: msg,
			Message:           msg,
		})
	}

	return err
}

/*
* Convert it into the string
 */
func (r *IError) Error() string {
	str, _ := json.MarshalIndent(r, "", "  ")
	return string(str)
}

func FindActionByName(actions []Module3Action, name string) *Module3Action {
	for _, action := range actions {
		if action.Name == name {
			return &action
		}
	}

	return nil
}

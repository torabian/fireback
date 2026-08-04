// Package ferror is a standalone copy of modules/fireback's IError family,
// under shorter names and with zero dependency on the core fireback
// package - matching the original (since-drifted) intent of Error.go's own
// "this file has no dependencies" comment.
//
// This is a copy, not the live implementation: modules/fireback/Error.go is
// untouched and still defines the real IError used everywhere today. See
// the bottom of this file for how modules/fireback would alias onto this
// package if/when that copy is retired in favor of this one.
//
// Later (not yet): consider wrapping/rebuilding this on top of
// github.com/rotisserie/eris for stack traces and richer wrapping. Eris
// error values still implement the standard `error` interface, so adopting
// it should only touch construction sites (Create401Error, CastToError,
// etc.), not the Error/FieldError struct shapes or ToPublicEndUser's output
// - the parts everything else depends on.
package ferror

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ErrorItem is a language-code -> message-text map, e.g.
// {"en": "Not found", "fa": "..."}. Renamed from nothing (same name) -
// already short.
type ErrorItem map[string]string

// Error is the renamed IError: the internal, detailed error envelope
// actions build instead of a plain `error`. Same fields, same semantics.
type Error struct {
	Message           ErrorItem              `json:"message,omitempty"`
	MessageParams     map[string]interface{} `json:"messageParams,omitempty"`
	MessageTranslated string                 `json:"messageTranslated,omitempty"`
	Errors            []*FieldError          `json:"errors,omitempty"`
	HttpCode          int32                  `json:"httpCode,omitempty"`
}

// FieldError is the renamed IErrorItem: one entry in Error.Errors, pointing
// at a specific field/location.
type FieldError struct {
	Location   string     `json:"location,omitempty"`
	Message    *ErrorItem `json:"message,omitempty"`
	ErrorParam string     `json:"errorParam,omitempty"`
	Type       string     `json:"type,omitempty"`
}

// PublicError is the renamed IPublicError: the safe-to-expose projection of
// Error, translated to one language, with no internal-only fields.
type PublicError struct {
	Message           string                 `json:"message,omitempty"`
	MessageTranslated string                 `json:"messageTranslated,omitempty"`
	MessageParams     map[string]interface{} `json:"messageParams,omitempty"`
	Errors            []*PublicFieldError    `json:"errors,omitempty"`
	HttpCode          int32                  `json:"httpCode,omitempty"`
}

// PublicFieldError is the renamed IPublicErrorItem.
type PublicFieldError struct {
	Location          string `json:"location,omitempty"`
	Message           string `json:"message,omitempty"`
	MessageTranslated string `json:"messageTranslated,omitempty"`
	ErrorParam        string `json:"errorParam,omitempty"`
	Type              string `json:"type,omitempty"`
}

func (x *Error) Is(value string) bool {
	if x.Message == nil {
		return false
	}

	code := x.Message["$"]

	if code == value {
		return true
	}

	return false
}

func (x *Error) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}

	return ""
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

// ToPublicEndUser converts an Error to a PublicError.
// Ensure you do not return Error directly to the public to avoid exposing internal error details.
//
// Parameters:
//   - q: An interface with a GetLanguage method that returns a string representing the language code. QueryDSL already has this interface so you can pass that instead
//
// Returns:
//   - A pointer to a PublicError containing translated messages and the HTTP code.
func (r *Error) ToPublicEndUser(q interface {
	GetLanguage() string
}) *PublicError {

	// Retrieve the language code
	lang := q.GetLanguage()
	if lang == "" {
		lang = "en" // or any default language code
	}

	err := &PublicError{}
	err.HttpCode = r.HttpCode
	err.MessageTranslated = ReplacePlaceholders(r.Message[lang], r.MessageParams)
	err.MessageParams = r.MessageParams
	err.Message = r.Message["$"]

	for _, item := range r.Errors {
		msg := (*item.Message)[lang]

		if item.ErrorParam != "" {
			msg = strings.ReplaceAll(msg, "%s", item.ErrorParam)
		}
		err.Errors = append(err.Errors, &PublicFieldError{
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
func (r *Error) Error() string {
	str, _ := json.MarshalIndent(r, "", "  ")
	return string(str)
}

func Create401Error(msg *ErrorItem, list []string) *Error {
	result := Error{
		Message:  *msg,
		HttpCode: 401,
		MessageParams: map[string]interface{}{
			"list": list,
		},
	}

	return &result
}

func Create401ParamOnly(msg *ErrorItem, params map[string]interface{}) *Error {
	result := Error{
		Message:       *msg,
		HttpCode:      401,
		MessageParams: params,
	}

	return &result
}

func Create401ErrorWithItems(msg *ErrorItem, items []*FieldError) *Error {

	result := Error{
		Message:  *msg,
		HttpCode: 401,
		Errors:   items,
	}

	return &result
}

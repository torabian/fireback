package fireback

import (
	"encoding/json"
	"errors"
	reflect "reflect"
	"strings"
)

func getLineAndCharFromOffset(body []byte, offset int64) (line int, col int) {
	line = 1
	col = 1
	for i := int64(0); i < offset && i < int64(len(body)); i++ {
		if body[i] == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return
}

func BindJsonStringWithDetails(jsonInput []byte, target any) *IError {
	var syntaxErr *json.SyntaxError
	var unmarshalTypeErr *json.UnmarshalTypeError
	var unsupportedTypeErr *json.UnsupportedTypeError
	var unmarshalFieldError *json.UnmarshalFieldError
	var marshallerError *json.MarshalerError

	err := json.Unmarshal(jsonInput, target)
	if err == nil {
		return nil
	}

	switch {
	case errors.As(err, &unmarshalTypeErr):
		fieldPath := unmarshalTypeErr.Field
		if fieldPath == "" {
			fieldPath = findFieldName(target, unmarshalTypeErr.Struct)
		}

		line, col := getLineAndCharFromOffset(jsonInput, unmarshalTypeErr.Offset)

		return Create401ParamOnly(&FirebackMessages.JsonInvalidFieldType, map[string]interface{}{
			"field":    fieldPath,
			"expected": unmarshalTypeErr.Type.String(),
			"offset":   unmarshalTypeErr.Offset,
			"actual":   unmarshalTypeErr.Value,
			"line":     line,
			"col":      col,
		})

	case errors.As(err, &syntaxErr):
		line, col := getLineAndCharFromOffset(jsonInput, syntaxErr.Offset)

		return Create401ParamOnly(&FirebackMessages.JsonMalformed, map[string]interface{}{
			"offset": syntaxErr.Offset,
			"line":   line,
			"col":    col,
		})
	case errors.As(err, &unmarshalFieldError):
		line, col := getLineAndCharFromOffset(jsonInput, int64(unmarshalFieldError.Field.Offset))

		return Create401ParamOnly(&FirebackMessages.JsonMalformed, map[string]interface{}{
			"offset": syntaxErr.Offset,
			"line":   line,
			"col":    col,
		})
	case errors.As(err, &marshallerError):
		line, col := getLineAndCharFromOffset(jsonInput, int64(unmarshalFieldError.Field.Offset))

		return Create401ParamOnly(&FirebackMessages.JsonMalformed, map[string]interface{}{
			"offset": syntaxErr.Offset,
			"line":   line,
			"col":    col,
		})
	case errors.As(err, &unsupportedTypeErr):
		return Create401ParamOnly(&FirebackMessages.JsonUnmarshalUnsupportedType, map[string]interface{}{
			"type": unsupportedTypeErr.Type,
		})

	default:

		errx := Create401ParamOnly(&FirebackMessages.JsonDecodingError, nil)
		errx.Errors = append(errx.Errors, &IErrorItem{
			Message: &ErrorItem{"en": err.Error()},
		})
		return errx
	}

}

func findFieldName(target any, structName string) string {
	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if strings.Contains(f.Type.String(), structName) {
			return f.Name
		}
	}
	return ""
}

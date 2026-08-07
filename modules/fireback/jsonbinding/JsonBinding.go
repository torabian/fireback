package jsonbinding

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	reflect "reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/complexes"
	"github.com/torabian/fireback/modules/fireback/ferror"
	"github.com/torabian/fireback/modules/fireback/gintools"
	"gopkg.in/yaml.v3"
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

func BindJsonStringWithDetails(jsonInput []byte, target any) *ferror.Error {
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

		return ferror.Create401ParamOnly(&JsonMessages.JsonInvalidFieldType, map[string]interface{}{
			"field":    fieldPath,
			"expected": unmarshalTypeErr.Type.String(),
			"offset":   unmarshalTypeErr.Offset,
			"actual":   unmarshalTypeErr.Value,
			"line":     line,
			"col":      col,
		})

	case errors.As(err, &syntaxErr):
		line, col := getLineAndCharFromOffset(jsonInput, syntaxErr.Offset)

		return ferror.Create401ParamOnly(&JsonMessages.JsonMalformed, map[string]interface{}{
			"offset": syntaxErr.Offset,
			"line":   line,
			"col":    col,
		})
	case errors.As(err, &unmarshalFieldError):
		line, col := getLineAndCharFromOffset(jsonInput, int64(unmarshalFieldError.Field.Offset))

		return ferror.Create401ParamOnly(&JsonMessages.JsonMalformed, map[string]interface{}{
			"offset": syntaxErr.Offset,
			"line":   line,
			"col":    col,
		})
	case errors.As(err, &marshallerError):
		line, col := getLineAndCharFromOffset(jsonInput, int64(unmarshalFieldError.Field.Offset))

		return ferror.Create401ParamOnly(&JsonMessages.JsonMalformed, map[string]interface{}{
			"offset": syntaxErr.Offset,
			"line":   line,
			"col":    col,
		})
	case errors.As(err, &unsupportedTypeErr):
		return ferror.Create401ParamOnly(&JsonMessages.JsonUnmarshalUnsupportedType, map[string]interface{}{
			"type": unsupportedTypeErr.Type,
		})

	default:

		errx := ferror.Create401ParamOnly(&JsonMessages.JsonDecodingError, nil)
		errx.Errors = append(errx.Errors, &ferror.FieldError{
			Message: &ferror.ErrorItem{"en": err.Error()},
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

type jsonCode string

const (
	BodyIsEmptyEof               jsonCode = "BodyIsEmptyEof"
	BodyIsMissing                jsonCode = "BodyIsMissing"
	BodyReadAfterClose           jsonCode = "BodyReadAfterClose"
	BodyUnexpectedEof            jsonCode = "BodyUnexpectedEof"
	DatabaseOperationError       jsonCode = "DatabaseOperationError"
	FieldInvalidEmail            jsonCode = "FieldInvalidEmail"
	FieldOneOf                   jsonCode = "FieldOneOf"
	FieldRequired                jsonCode = "FieldRequired"
	InvalidContent               jsonCode = "InvalidContent"
	InvalidFormDataContentType   jsonCode = "InvalidFormDataContentType"
	JsonDecodingError            jsonCode = "JsonDecodingError"
	JsonInvalidFieldType         jsonCode = "JsonInvalidFieldType"
	JsonMalformed                jsonCode = "JsonMalformed"
	JsonUnmarshalUnsupportedType jsonCode = "JsonUnmarshalUnsupportedType"
	UnknownErrorReadingBody      jsonCode = "UnknownErrorReadingBody"
	XmlDecodingError             jsonCode = "XmlDecodingError"
	XmlMalformed                 jsonCode = "XmlMalformed"
	XmlUnmarshalError            jsonCode = "XmlUnmarshalError"
	YamlDecodingError            jsonCode = "YamlDecodingError"
	YamlTypeError                jsonCode = "YamlTypeError"
	FormDataMalformed            jsonCode = "FormDataMalformed"
)

var JsonMessages = newJsonMessageCode()

func newJsonMessageCode() *JsonMsgs {
	return &JsonMsgs{
		BodyIsEmptyEof: ferror.ErrorItem{
			"$":    "BodyIsEmptyEof",
			"$key": "io.EOF",
			"en":   "Body is empty. Please provide the necessary data and try again.",
		},
		BodyIsMissing: ferror.ErrorItem{
			"$":  "BodyIsMissing",
			"en": "Body content is not correct. You need a valid json.",
		},
		BodyReadAfterClose: ferror.ErrorItem{
			"$":    "BodyReadAfterClose",
			"$key": "http.ErrBodyReadAfterClose",
			"en":   "Body is read after closed. The request might have been processed incorrectly.",
		},
		BodyUnexpectedEof: ferror.ErrorItem{
			"$":    "BodyUnexpectedEof",
			"$key": "io.ErrUnexpectedEOF",
			"en":   "Body unexpected EOF. The data you sent appears incomplete. Please check your request and try again.",
		},
		JsonDecodingError: ferror.ErrorItem{
			"$":  "JsonDecodingError",
			"en": "Unknown error happened upon decoding.",
		},
		JsonInvalidFieldType: ferror.ErrorItem{
			"$":  "JsonInvalidFieldType",
			"en": "Expected type '%expected' but got a different type '%actual' on %offset (line %line, col %col)",
		},
		FormDataMalformed: ferror.ErrorItem{
			"$":  "FormDataMalformed",
			"en": "The form data submitted is malformed or contains invalid fields. Please check the form and ensure all required fields are properly filled out.",
		},
		JsonMalformed: ferror.ErrorItem{
			"$":  "JsonMalformed",
			"en": "Json is malformed. Check your commas, braces, etc.",
		},
		JsonUnmarshalUnsupportedType: ferror.ErrorItem{
			"$":  "JsonUnmarshalUnsupportedType",
			"en": "Unsupported type when unmarshalling json",
		},
		XmlDecodingError: ferror.ErrorItem{
			"$":  "XmlDecodingError",
			"en": "Something went wrong while processing the XML. Please check the content or try again later.",
		},
		XmlMalformed: ferror.ErrorItem{
			"$":  "XmlMalformed",
			"en": "The XML format is broken or incomplete. Please make sure all tags are properly opened and closed.",
		},
		XmlUnmarshalError: ferror.ErrorItem{
			"$":  "XmlUnmarshalError",
			"en": "The XML structure doesn’t match the expected format. Some elements may be missing or in the wrong place.",
		},
		YamlDecodingError: ferror.ErrorItem{
			"$":  "YamlDecodingError",
			"en": "There’s something wrong with the format of your YAML. Please check indentation, colons, and line breaks to fix the formatting.",
		},
		UnknownErrorReadingBody: ferror.ErrorItem{
			"$":  "UnknownErrorReadingBody",
			"en": "We cannot read the body of your request.",
		},
		YamlTypeError: ferror.ErrorItem{
			"$":  "YamlTypeError",
			"en": "One of the values is in the wrong format. For example, you might have entered text instead of a number or used quotes incorrectly.",
		},
	}
}

type JsonMsgs struct {
	BodyIsEmptyEof               ferror.ErrorItem
	BodyIsMissing                ferror.ErrorItem
	BodyReadAfterClose           ferror.ErrorItem
	BodyUnexpectedEof            ferror.ErrorItem
	JsonDecodingError            ferror.ErrorItem
	JsonInvalidFieldType         ferror.ErrorItem
	JsonMalformed                ferror.ErrorItem
	JsonUnmarshalUnsupportedType ferror.ErrorItem
	XmlDecodingError             ferror.ErrorItem
	XmlMalformed                 ferror.ErrorItem
	UnknownErrorReadingBody      ferror.ErrorItem

	XmlUnmarshalError ferror.ErrorItem
	YamlDecodingError ferror.ErrorItem
	FormDataMalformed ferror.ErrorItem
	YamlTypeError     ferror.ErrorItem
}

// Use it for requests which are kinda having body, such as post, put, patch, etc.
// It would read the body (either if it's json, form-data, yaml, etc, based on headers)
// and cast it to the 'body'. Make sure calling this with &body, not body
// Extend this function if you want to support different formats.
func ReadGinRequestBodyAndCastToGoStruct(c *gin.Context, body any) (aborted bool) {

	// Only following request methods do have a body
	if c.Request.Method != "POST" && c.Request.Method != "PATCH" && c.Request.Method != "PUT" {
		return false
	}

	bodyBytes, err := ginBodyToBytes(c)
	if err != nil {
		return abortWithError(c, err)
	}

	switch gintools.DetectGinContentType(c) {
	case gintools.ContentTypeYAML:
		if err := BindYamlStringWithDetails(bodyBytes, body); err != nil {
			return abortWithError(c, err)
		}
	case gintools.ContentTypeFormData:
		if err := BindMultiPartFormDataWithDetails(c, body); err != nil {
			return abortWithError(c, err)
		}
	case gintools.ContentTypeURLEncoded:
		if err := BindFormUrlEncodedWithDetails(c, body); err != nil {
			return abortWithError(c, err)
		}
	case gintools.ContentTypeXML:
		if err := BindXmlStringWithDetails(bodyBytes, body); err != nil {
			return abortWithError(c, err)
		}
	default:
		if err := BindJsonStringWithDetails(bodyBytes, body); err != nil {
			return abortWithError(c, err)
		}
	}

	return false
}

func abortWithError(c *gin.Context, err error) bool {
	accept := c.GetHeader("Accept")
	isYAML := accept == "application/x-yaml" || accept == "application/yaml" || accept == "text/yaml"

	if isYAML {
		c.Header("Content-Type", "application/x-yaml")
		yamlData, marshalErr := yaml.Marshal(err)
		if marshalErr != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to marshal yaml"})
			return true
		}
		c.Writer.WriteHeader(500)
		c.Writer.Write(yamlData)
	} else {
		c.AbortWithStatusJSON(500, err)
	}
	return true
}

func ginBodyToBytes(c *gin.Context) ([]byte, *ferror.Error) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, ferror.Create401Error(&JsonMessages.BodyIsEmptyEof, []string{})
		} else if errors.Is(err, io.ErrUnexpectedEOF) {
			return nil, ferror.Create401Error(&JsonMessages.BodyUnexpectedEof, []string{})
		} else if errors.Is(err, http.ErrBodyReadAfterClose) {
			return nil, ferror.Create401Error(&JsonMessages.BodyReadAfterClose, []string{})
		} else {
			return nil, ferror.Create401Error(&JsonMessages.UnknownErrorReadingBody, []string{})
		}
	}

	// Reset the body so it can be read again later
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

func BindYamlStringWithDetails(yamlInput []byte, target any) *ferror.Error {
	var node yaml.Node
	err := yaml.Unmarshal(yamlInput, &node)
	if err != nil {
		if syntaxErr, ok := err.(*yaml.TypeError); ok && len(syntaxErr.Errors) > 0 {
			return ferror.Create401ParamOnly(&JsonMessages.YamlTypeError, map[string]interface{}{
				"errors": syntaxErr.Errors,
			})
		}
		return ferror.Create401ParamOnly(&JsonMessages.YamlDecodingError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = node.Decode(target)

	if err != nil {
		var yamlNodeErr *yaml.TypeError
		if errors.As(err, &yamlNodeErr) && len(yamlNodeErr.Errors) > 0 {
			return ferror.Create401ParamOnly(&JsonMessages.YamlTypeError, map[string]interface{}{
				"errors": yamlNodeErr.Errors,
			})
		}

		return ferror.Create401ParamOnly(&JsonMessages.YamlDecodingError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return nil
}

func BindMultiPartFormDataWithDetails(c *gin.Context, target any) *ferror.Error {
	// For URL-encoded forms
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB limit
		return ferror.Create401ParamOnly(&JsonMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Prepare a map to hold the form data
	formData := c.Request.MultipartForm
	formMap := make(map[string]interface{})

	for fieldName, files := range formData.File {
		for _, fileHeader := range files {
			xfile, err := complexes.ConvertToXFile(fileHeader)
			if err != nil {
				// return ferror.CastToIError(err)
				return &ferror.Error{MessageTranslated: err.Error()}
			}

			formMap[fieldName] = xfile
		}
	}

	// Iterate over the form data and populate the map
	for key, values := range formData.Value {
		// If a key has multiple values, we keep it as a slice
		if len(values) > 1 {
			formMap[key] = values
		} else {
			// If it has only one value, store it as a single value
			formMap[key] = values[0]
		}
	}

	// This is very inefficient way of formatting the data, to marshal to json and
	// and unmarshall it specially if there is a file uploaded.

	// Convert the form map to a JSON string
	formJSON, err := json.Marshal(formMap)
	if err != nil {
		return ferror.Create401ParamOnly(&JsonMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Now unmarshal the JSON into the struct
	if err := json.Unmarshal(formJSON, target); err != nil {
		return ferror.Create401ParamOnly(&JsonMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return nil

}

func BindFormUrlEncodedWithDetails(c *gin.Context, target any) *ferror.Error {
	// For URL-encoded forms
	if err := c.Request.ParseForm(); err != nil {
		return ferror.Create401ParamOnly(&JsonMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Prepare a map to hold the form data
	formData := c.Request.Form
	formMap := make(map[string]interface{})

	// Iterate over the form data and populate the map
	for key, values := range formData {
		// If a key has multiple values, we keep it as a slice
		if len(values) > 1 {
			formMap[key] = values
		} else {
			// If it has only one value, store it as a single value
			formMap[key] = values[0]
		}
	}

	// Convert the form map to a JSON string
	formJSON, err := json.Marshal(formMap)
	if err != nil {
		return ferror.Create401ParamOnly(&JsonMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
		return nil
	}

	// Now unmarshal the JSON into the struct
	if err := json.Unmarshal(formJSON, target); err != nil {
		return ferror.Create401ParamOnly(&JsonMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
		return nil
	}

	return nil

}

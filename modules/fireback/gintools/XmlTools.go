package gintools

import (
	"encoding/xml"
	"errors"

	"github.com/torabian/fireback/modules/fireback/ferror"
)

func BindXmlStringWithDetails(xmlInput []byte, target any) *ferror.Error {
	var syntaxErr *xml.SyntaxError
	var unmarshalErr *xml.UnmarshalError

	err := xml.Unmarshal(xmlInput, target)
	if err == nil {
		return nil
	}

	switch {
	case errors.As(err, &syntaxErr):
		return ferror.Create401ParamOnly(&JsonMessages.XmlMalformed, map[string]interface{}{
			"line": syntaxErr.Line,
		})

	case errors.As(err, &unmarshalErr):
		return ferror.Create401ParamOnly(&JsonMessages.XmlUnmarshalError, map[string]interface{}{
			"error": unmarshalErr.Error(),
		})

	default:
		return ferror.Create401ParamOnly(&JsonMessages.XmlDecodingError, map[string]interface{}{
			"error": err.Error(),
		})
	}
}

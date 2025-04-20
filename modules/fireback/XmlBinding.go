package fireback

import (
	"encoding/xml"
	"errors"
)

func BindXmlStringWithDetails(xmlInput []byte, target any) *IError {
	var syntaxErr *xml.SyntaxError
	var unmarshalErr *xml.UnmarshalError

	err := xml.Unmarshal(xmlInput, target)
	if err == nil {
		return nil
	}

	switch {
	case errors.As(err, &syntaxErr):
		return Create401ParamOnly(&WorkspacesMessages.XmlMalformed, map[string]interface{}{
			"line": syntaxErr.Line,
		})

	case errors.As(err, &unmarshalErr):
		return Create401ParamOnly(&WorkspacesMessages.XmlUnmarshalError, map[string]interface{}{
			"error": unmarshalErr.Error(),
		})

	default:
		return Create401ParamOnly(&WorkspacesMessages.XmlDecodingError, map[string]interface{}{
			"error": err.Error(),
		})
	}
}

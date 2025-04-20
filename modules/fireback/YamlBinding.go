package fireback

import (
	"errors"

	"gopkg.in/yaml.v3"
)

func BindYamlStringWithDetails(yamlInput []byte, target any) *IError {
	var node yaml.Node
	err := yaml.Unmarshal(yamlInput, &node)
	if err != nil {
		if syntaxErr, ok := err.(*yaml.TypeError); ok && len(syntaxErr.Errors) > 0 {
			return Create401ParamOnly(&FirebackMessages.YamlTypeError, map[string]interface{}{
				"errors": syntaxErr.Errors,
			})
		}
		return Create401ParamOnly(&FirebackMessages.YamlDecodingError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = node.Decode(target)

	if err != nil {
		var yamlNodeErr *yaml.TypeError
		if errors.As(err, &yamlNodeErr) && len(yamlNodeErr.Errors) > 0 {
			return Create401ParamOnly(&FirebackMessages.YamlTypeError, map[string]interface{}{
				"errors": yamlNodeErr.Errors,
			})
		}

		return Create401ParamOnly(&FirebackMessages.YamlDecodingError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return nil
}

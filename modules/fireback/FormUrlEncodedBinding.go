package fireback

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func BindFormUrlEncodedWithDetails(c *gin.Context, target any) *IError {
	// For URL-encoded forms
	if err := c.Request.ParseForm(); err != nil {
		return Create401ParamOnly(&FirebackMessages.FormDataMalformed, map[string]interface{}{
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
		return Create401ParamOnly(&FirebackMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
		return nil
	}

	// Now unmarshal the JSON into the struct
	if err := json.Unmarshal(formJSON, target); err != nil {
		return Create401ParamOnly(&FirebackMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
		return nil
	}

	return nil

}

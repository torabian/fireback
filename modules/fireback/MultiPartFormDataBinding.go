package fireback

import (
	"encoding/json"
	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func BindMultiPartFormDataWithDetails(c *gin.Context, target any) *IError {
	// For URL-encoded forms
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB limit
		return Create401ParamOnly(&FirebackMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Prepare a map to hold the form data
	formData := c.Request.MultipartForm
	formMap := make(map[string]interface{})

	for fieldName, files := range formData.File {
		for _, fileHeader := range files {
			xfile, err := ConvertToXFile(fileHeader)
			if err != nil {
				return CastToIError(err)
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
		return Create401ParamOnly(&FirebackMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Now unmarshal the JSON into the struct
	if err := json.Unmarshal(formJSON, target); err != nil {
		return Create401ParamOnly(&FirebackMessages.FormDataMalformed, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return nil

}

func ConvertToXFile(fileHeader *multipart.FileHeader) (*XFile, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	blob, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	mime := fileHeader.Header.Get("Content-Type")

	return &XFile{
		Meta: XFileMeta{
			FileName: fileHeader.Filename,
			Mime:     mime,
		},
		Filesize: uint64(fileHeader.Size),
		Blob:     blob,
		// FileID, //URL: populate after upload or save
	}, nil
}

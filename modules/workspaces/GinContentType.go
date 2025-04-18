package workspaces

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type ContentType string

const (
	ContentTypeJSON       ContentType = "json"
	ContentTypeURLEncoded ContentType = "urlencoded"
	ContentTypeFormData   ContentType = "form-data"
	ContentTypeYAML       ContentType = "yaml"
	ContentTypeXML        ContentType = "xml"
	ContentTypeText       ContentType = "text"
	ContentTypeBinary     ContentType = "binary"
	ContentTypeUnknown    ContentType = "unknown"
)

func DetectGinContentType(c *gin.Context) ContentType {
	contentType := c.GetHeader("Content-Type")

	switch {
	case strings.HasPrefix(contentType, "application/json"):
		return ContentTypeJSON
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		return ContentTypeURLEncoded
	case strings.HasPrefix(contentType, "multipart/form-data"):
		return ContentTypeFormData
	case strings.HasPrefix(contentType, "application/yaml"), strings.HasPrefix(contentType, "application/x-yaml"), strings.HasPrefix(contentType, "text/yaml"):
		return ContentTypeYAML
	case strings.HasPrefix(contentType, "application/xml"), strings.HasPrefix(contentType, "text/xml"):
		return ContentTypeXML
	case strings.HasPrefix(contentType, "text/plain"):
		return ContentTypeText
	case strings.HasPrefix(contentType, "application/octet-stream"):
		return ContentTypeBinary
	default:
		return ContentTypeUnknown
	}
}

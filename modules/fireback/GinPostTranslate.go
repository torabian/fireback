package fireback

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAcceptFromGinHeaders(c *gin.Context) string {

	acceptLang := c.GetHeader("accept-language")
	if acceptLang != "" && len(acceptLang) == 2 {
		return strings.ToLower(acceptLang)
	}

	return "en"
}

func GinPostTranslateErrorMessages(dic map[string]map[string]string) gin.HandlerFunc {

	return func(c *gin.Context) {
		wb := &toolBodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
			status:         Origin,
		}

		lang := GetExactLanguageFromAcceptLanguage(c.GetHeader("accept-language"), []string{})
		if lang != "fa" && lang != "en" && lang != "pl" && lang != "ru" && lang != "ua" {
			lang = "en"
		}
		c.Writer = wb
		c.Next()
		statusCode := c.Writer.Status()

		wb.status = Replace
		originBytes := wb.body
		wb.body = &bytes.Buffer{}

		if statusCode >= 400 {
			res := IResponseFromString[any](originBytes.String())
			data, _ := json.MarshalIndent(res, "", "  ")
			wb.Write(data)
		} else {
			wb.Write(originBytes.Bytes())
		}
	}
}

type toolBodyWriter struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status byte
}

const (
	Origin  byte = 0x0
	Replace      = 0x1
)

func (r toolBodyWriter) Write(b []byte) (int, error) {
	if r.status == 0x1 {
		r.body.Write(b)
		return r.ResponseWriter.Write(b)
	} else {
		return r.body.Write(b) //r.ResponseWriter.Write(b)
	}
}

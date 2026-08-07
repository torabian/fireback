package fireback

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/gintools"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v2"
)

func HttpUpdateEntity[T any, V any](c *gin.Context, fn func(QueryDSL, T) (V, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	var body T

	if gintools.ReadGinRequestBodyAndCastToGoStruct(c, &body) {
		return
	}

	if entity, err := fn(f, body); err != nil {
		GinWriteContent(c, int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&f)})
	} else {
		GinWriteContent(c, 200, gin.H{
			"data": entity,
		})
	}
}

func isYaml(headerValue string) bool {

	return slices.Contains([]string{
		"application/x-yaml",
		"application/yaml",
		"text/yaml",
		"yaml",
		"yml",
	}, headerValue)

}
func isCsv(headerValue string) bool {

	return slices.Contains([]string{
		"text/csv",
		"csv",
	}, headerValue)

}

func IsYaml(c *gin.Context) bool {
	return isYaml(c.GetHeader("Accept"))
}

func IsYamlCli(c *cli.Command) bool {
	return isYaml(c.String("x-accept"))
}

func IsCsvCli(c *cli.Command) bool {
	return isCsv(c.String("x-accept"))
}

// When done with a http handler, you can use this to write the content
// Use it for successful operations
func GinWriteContent(c *gin.Context, code int, content gin.H) {

	isYAML := IsYaml(c)

	if isYAML {
		c.Header("Content-Type", "application/x-yaml")
		c.Status(code)
		yamlData, err := yaml.Marshal(content)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to marshal yaml"})
			return
		}
		c.Writer.Write(yamlData)

		return
	}

	c.JSON(code, content)
}

func HttpGetEntity[T any](
	c *gin.Context,
	fn func(QueryDSL) (T, *IError),
) {
	id := c.Param("uniqueId")
	f := ExtractQueryDslFromGinContext(c)
	f.UniqueId = id

	if item, err := fn(f); err != nil {
		code := http.StatusBadRequest
		if err.HttpCode > 0 {
			code = int(err.HttpCode)
		}

		GinWriteContent(c, code, gin.H{
			"error": err.ToPublicEndUser(&f),
		})
	} else {

		data := PolyglotQueryHandler(item, &f)
		GinWriteContent(c, 200, gin.H{
			"data": gin.H{
				"item": data,
			},
		})
	}
}

func toBytes(v any) []byte {
	switch t := v.(type) {
	case []byte:
		return t
	case string:
		return []byte(t)
	default:
		b, _ := json.Marshal(t)
		return b
	}
}

func WriteResponse(c *gin.Context, status int, resp emigo.EmiActionResult) {
	payload := resp.GetPayload()

	if payload == nil {
		c.Status(status)
		return
	}

	headers := resp.GetRespHeaders()
	if headers == nil {
		c.JSON(status, payload)

		return
	}
	switch p := payload.(type) {
	case func(io.Writer) error:
		// template streaming
		if err := p(c.Writer); err != nil {
			c.Error(err)
			return
		}

	case io.Reader:
		// file or image streaming
		_, err := io.Copy(c.Writer, p)
		if err != nil {
			c.Error(err)
			return
		}

	default:
		// fallback based on content-type
		ct := headers["Content-Type"]
		c.Data(status, ct, toBytes(payload))

	}
}

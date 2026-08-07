package fireback

import (
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

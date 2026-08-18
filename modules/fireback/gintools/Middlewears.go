package gintools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinAllowEverything() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.Writer.Header()

		h.Set("Access-Control-Allow-Origin", "*")
		h.Set("Access-Control-Allow-Methods", "*")
		h.Set("Access-Control-Allow-Headers", "*")
		h.Set("Access-Control-Expose-Headers", "*")
		h.Set("Access-Control-Max-Age", "86400")

		// Only enable this if you're NOT using "*"
		// h.Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

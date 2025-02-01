package workspaces

import (
	"github.com/gin-gonic/gin"
)

func HttpRequestAuthorize2(c *gin.Context) {
	HttpPostEntity(c, PassportActionAuthorize2)
}

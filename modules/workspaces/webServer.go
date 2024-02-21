package workspaces

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func Body(c *gin.Context) map[string]string {
	content, _ := c.GetRawData()
	body := map[string]string{}
	json.Unmarshal(content, &body)

	return body
}

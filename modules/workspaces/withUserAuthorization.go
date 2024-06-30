package workspaces

import (
	"github.com/gin-gonic/gin"
)

func WithUserAuthorization(scopedToUser bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		user, err := GetUserFromToken(token)

		if err != nil {

			c.AbortWithStatusJSON(401, gin.H{
				"error": gin.H{
					"message": "Authorization is not valid.",
				},
			})
		}

		// Means user would be only query things that is related to himself, not other users
		if scopedToUser {
			c.Set("internal_sql", `user_id = "`+user.UniqueId+`"`)
		}

		c.Set("user_id", &user.UniqueId)
		c.Set("user", user)
	}
}

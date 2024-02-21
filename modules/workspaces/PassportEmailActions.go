package workspaces

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func httpConfirmEmail(c *gin.Context) {

	id := c.Param("uniqueId")

	confirmed := ConfirmEmailAddress(id)

	if confirmed {
		c.JSON(200, gin.H{
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"error": gin.H{"code": PassportMessageCode.AlreadyConfirmed},
		})
	}
}

func HttpRequestAuthorize2(c *gin.Context) {
	HttpPostEntity(c, PassportActionAuthorize2)
}

// func HttpRequestResetMailPassword(c *gin.Context) {

// }

func HttpGetResetMailPasswordInfo(c *gin.Context) {

	id := c.Param("uniqueId")

	info := GetResetPasswordInfo(id)

	if info.UniqueId == "" {
		c.JSON(403, gin.H{
			"error": gin.H{
				"code": PassportMessageCode.InvitationExpired,
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"item": info,
		},
	})
}

func HttpResetMailPassword(c *gin.Context) {
	content, _ := c.GetRawData()
	body := &ResetEmailDto{}
	json.Unmarshal(content, &body)
	id := c.Param("uniqueId")

	if body.Password == nil {
		c.JSON(402, gin.H{
			"error": gin.H{
				"code": PassportMessageCode.PasswordRequired,
				"errors": []gin.H{
					{
						"location": "password",
					},
				},
			},
		})
		return
	}

	if id == "" {
		c.JSON(402, gin.H{
			"error": gin.H{
				"code": PassportMessageCode.ResetNotFound,
			},
		})
		return
	}

	changed, email := ResetUserPasswordWithRequest(id, *body.Password)

	if changed {
		user, token, err := SigninUserWithEmailAndPassword(email, *body.Password)

		if err == nil {
			// ws := GetUserWorkspaces(user.UniqueId)

			c.JSON(200, gin.H{
				"data": gin.H{

					"user":  user,
					"token": token,
					// "workspaces":  ws,
					"exchangeKey": PutTokenInExchangePool(token),
				},
			})
			return
		}
	} else {
		c.JSON(403, gin.H{
			"error": gin.H{
				"code": PassportMessageCode.InvitationExpired,
			},
		})
		return
	}

}

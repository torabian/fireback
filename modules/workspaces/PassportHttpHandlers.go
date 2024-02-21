package workspaces

// func WorkspaceInviteGetOne(query QueryDSL) (*WorkspaceInviteEntity, *IError) {
// 	refl := reflect.ValueOf(&WorkspaceInviteEntity{})
// 	invite, err := GetOneEntity[WorkspaceInviteEntity](query, refl)

// 	if invite != nil {
// 		pass := &PassportEntity{}
// 		if err := GetDbRef().Debug().Model(&PassportEntity{}).Where("value = ?", invite.Email).First(&pass).Error; err == nil {
// 			if pass != nil && pass.UserId != "" {
// 				invite.InviteeUserId = &pass.UserId
// 			}
// 		} else {
// 			invite.InviteeUserId = nil
// 		}

// 	}

// 	return invite, err
// }

// func HttpGetInvitePublicInformation(c *gin.Context) {
// 	HttpGetEntity(c, WorkspaceInviteGetOne)
// }

// func HttpPostPassport(c *gin.Context) {
// 	HttpPostEntity(c, PassportActionCreate)
// }

// func HttpQueryPassports(c *gin.Context) {
// 	HttpQueryEntity(c, PassportActionQuery)
// }

// func HttpGetOnePassport(c *gin.Context) {
// 	HttpGetEntity(c, PassportActionGetOne)
// }

// func HttpRemovePassport(c *gin.Context) {
// 	HttpRemoveEntity(c, PassportActionRemove)
// }

// func HttpUpdatePassport(c *gin.Context) {
// 	HttpUpdateEntity(c, PassportActionUpdate)
// }

// func HttpPostWorkspaceJoinKey(c *gin.Context) {
// 	HttpPostEntity(c, PublicJoinKeyActionCreate)
// }

// func HttpGetOneJoinKey(c *gin.Context) {
// 	HttpGetEntity(c, PublicJoinKeyActionGetOne)
// }

// func HttpGetOneJoinKeyPublicInfo(c *gin.Context) {
// 	HttpGetEntity(c, PublicJoinKeyActionGetOnePublic)
// }

// func HttpPatchJoinKey(c *gin.Context) {
// 	HttpUpdateEntity(c, PublicJoinKeyActionUpdate)
// }

// func HttpRemovePublicJoinKey(c *gin.Context) {
// 	HttpRemoveEntity(c, PublicJoinKeyActionRemove)
// }

// func HttpGetWorkspaceJoinKeys(c *gin.Context) {
// 	HttpQueryEntity(c, PublicJoinKeyActionQuery)
// }

// func HttpPostUserPhoneNumberSignupConfirm(c *gin.Context) {
// 	HttpPostEntity(c, UserActionCreateByPhoneNumberConfirm)
// }

// func HttpPostUserPhoneNumberSignup(c *gin.Context) {

// 	var body PhoneNumberAccountCreationDto

// 	if err := c.BindJSON(&body); err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	if entity, err := UserActionCreateByPhoneNumber(body, QueryDSL{}); err != nil {
// 		c.AbortWithStatusJSON(500, err)
// 	} else {

// 		if os.Getenv("ENV") == "production" {
// 			c.JSON(200, gin.H{
// 				"data": gin.H{},
// 			})
// 		} else {
// 			c.JSON(200, gin.H{
// 				"data": gin.H{"code": entity, "warning$": "You Production environemnt, this code should be sent through sms or phone call"},
// 			})
// 		}
// 	}

// }

// func HttpPostUserPhoneNumberSignin(c *gin.Context) {
// 	HttpPostEntity(c, PassportActionCreateSessionByPhoneNumber)
// }

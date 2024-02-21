package workspaces

// func GetPassportModule2Actions() []Module2Action {
// 	routes := []Module2Action{
// 		{
// 			Method: "GET",
// 			Url:    "/passports",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpQueryPassports,
// 			},
// 			ResponseEntity: &[]PassportEntity{},
// 		},
// 		{
// 			Method: "GET",
// 			Url:    "/passport/:uniqueId",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpGetOnePassport,
// 			},
// 			ResponseEntity: &PassportEntity{},
// 		},
// 		{
// 			Method: "POST",
// 			Url:    "/passport",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpPostPassport,
// 			},
// 			RequestEntity:  &PassportEntity{},
// 			ResponseEntity: &PassportEntity{},
// 		},
// 		{
// 			Method: "PATCH",
// 			Url:    "/passport/:uniqueId",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpUpdatePassport,
// 			},
// 			ResponseEntity: &PassportEntity{},
// 			RequestEntity:  &PassportEntity{},
// 		},
// 		{
// 			Method: "DELETE",
// 			Url:    "/passport",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpRemovePassport,
// 			},
// 			RequestEntity:  &DeleteRequest{},
// 			ResponseEntity: &DeleteResponse{},
// 			TargetEntity:   &PassportEntity{},
// 		},
// 		{
// 			Method:         "POST",
// 			Url:            ("/passport/signup/email"),
// 			Handlers:       []gin.HandlerFunc{HttpPostUserEmailSignup},
// 			RequestEntity:  &ClassicAuthDto{},
// 			ResponseEntity: &UserSessionDto{},
// 		},
// 		{
// 			Method: "GET",
// 			Url:    "/workspace/publicjoinkeys",
// 			Handlers: []gin.HandlerFunc{
// 				WithAuthorization([]string{}),
// 				HttpGetWorkspaceJoinKeys,
// 			},
// 			ResponseEntity: &[]PublicJoinKey{},
// 		},
// 		{
// 			Method: "POST",
// 			Url:    "/workspace/publicjoinkey",
// 			Handlers: []gin.HandlerFunc{
// 				WithAuthorization([]string{}),
// 				HttpPostWorkspaceJoinKey,
// 			},
// 			ResponseEntity: &PublicJoinKey{},
// 			RequestEntity:  &PublicJoinKey{},
// 		},
// 		{
// 			Method: "GET",
// 			Url:    "/workspace/publicjoinkey/:uniqueId",
// 			Handlers: []gin.HandlerFunc{
// 				WithAuthorization([]string{}),
// 				HttpGetOneJoinKey,
// 			},
// 			ResponseEntity: &PublicJoinKey{},
// 		},
// 		{
// 			Method: "GET",
// 			Url:    "/workspace/publicJoinKeyInfo/:uniqueId",
// 			Handlers: []gin.HandlerFunc{
// 				HttpGetOneJoinKeyPublicInfo,
// 			},
// 			ResponseEntity: &PublicJoinKey{},
// 		},
// 		{
// 			Method: "DELETE",
// 			Url:    "/workspace/publicjoinkey",
// 			Handlers: []gin.HandlerFunc{
// 				WithAuthorization([]string{}),
// 				HttpRemovePublicJoinKey,
// 			},
// 			TargetEntity:   &PublicJoinKey{},
// 			RequestEntity:  &DeleteRequest{},
// 			ResponseEntity: &DeleteResponse{},
// 		},
// 		{
// 			Method: "PATCH",
// 			Url:    "/workspace/publicjoinkey",
// 			Handlers: []gin.HandlerFunc{
// 				WithAuthorization([]string{}),
// 				HttpPatchJoinKey,
// 			},
// 			RequestEntity:  &PublicJoinKey{},
// 			ResponseEntity: &PublicJoinKey{},
// 		},

// 		{
// 			Method: "GET",
// 			Url:    "/workspaces/invite/:uniqueId",
// 			Handlers: []gin.HandlerFunc{
// 				HttpGetInvitePublicInformation,
// 			},
// 			ResponseEntity: &WorkspaceInviteEntity{},
// 		},
// 		{
// 			Method:         "POST",
// 			Url:            ("/passport/signup/phoneNumber"),
// 			Handlers:       []gin.HandlerFunc{HttpPostUserPhoneNumberSignup},
// 			RequestEntity:  &PhoneNumberAccountCreationDto{},
// 			ResponseEntity: &OkayResponse{},
// 		},
// 		{
// 			Method:         "POST",
// 			Url:            ("/passport/signup/phoneNumberConfirm"),
// 			Handlers:       []gin.HandlerFunc{HttpPostUserPhoneNumberSignup},
// 			RequestEntity:  &PhoneNumberAccountCreationDto{},
// 			ResponseEntity: &OkayResponse{},
// 		},
// 		{
// 			Method:         "POST",
// 			Url:            ("/passport/signin/phoneNumber"),
// 			Handlers:       []gin.HandlerFunc{HttpPostUserPhoneNumberSignin},
// 			RequestEntity:  &PhoneNumberAccountCreationDto{},
// 			ResponseEntity: &UserSessionDto{},
// 		},
// 		{
// 			Method:         "POST",
// 			Url:            ("/confirm/email/:uniqueId"),
// 			Handlers:       []gin.HandlerFunc{httpConfirmEmail},
// 			RequestEntity:  &EmptyRequest{},
// 			ResponseEntity: &OkayResponse{},
// 		},

// 		{
// 			Method:         "POST",
// 			Url:            ("/passport/authorize2"),
// 			Handlers:       []gin.HandlerFunc{HttpRequestAuthorize2},
// 			RequestEntity:  &OtpAuthenticateDto{},
// 			ResponseEntity: &EmailOtpResponse{},
// 		},
// 		{
// 			Method:         "POST",
// 			Url:            ("/passport/reset-mail-password/:uniqueId"),
// 			Handlers:       []gin.HandlerFunc{HttpResetMailPassword},
// 			RequestEntity:  &ResetEmailDto{},
// 			ResponseEntity: &UserSessionDto{},
// 		},
// 		{
// 			Method:   "GET",
// 			Url:      ("/passport/request-reset-mail-password"),
// 			Handlers: []gin.HandlerFunc{httpConfirmEmail},
// 			// RequestEntity:  &EmptyRequest{},
// 			ResponseEntity: &OkayResponse{},
// 		},
// 		{
// 			Method:   "GET",
// 			Url:      ("/passport/reset-mail-password-info/:uniqueId"),
// 			Handlers: []gin.HandlerFunc{HttpGetResetMailPasswordInfo},
// 			// RequestEntity:  &EmptyRequest{},
// 			ResponseEntity: &OkayResponse{},
// 		},
// 	}

// 	return routes
// }

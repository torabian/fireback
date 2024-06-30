package workspaces

var UserEntityFields map[string]string = map[string]string{
	"firstName": "firstName",
	"lastName":  "lastName",
	"photo":     "photo",
}

// Many of these routes are needed. I just temporarily commented them
// func CreateUserRouter(r *gin.Engine) {

// 	httpRoutes := []Module2Action{
// 		{
// 			Method: "PATCH",
// 			Url:    "/user",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpUpdateUser,
// 			},
// 			RequestEntity:  &UserEntity{},
// 			ResponseEntity: &UserEntity{},
// 		},
// 		{
// 			Method: "POST",
// 			Url:    "/user",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpPostUser,
// 			},
// 			RequestEntity:  &UserEntity{},
// 			ResponseEntity: &UserEntity{},
// 		},
// 		{
// 			Method: "DELETE",
// 			Url:    "/user",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpRemoveUser,
// 			},
// 			RequestEntity:  &DeleteRequest{},
// 			ResponseEntity: &DeleteResponse{},
// 			TargetEntity:   &UserEntity{},
// 		},
// 		{
// 			Method: "GET",
// 			Url:    "/users",
// 			Handlers: []gin.HandlerFunc{
// 				WithAuthorization([]string{PERM_ROOT_WORKSPACES_USER_QUERY}),
// 				HttpQueryUsers,
// 			},
// 			ResponseEntity: &[]UserEntity{},
// 		},
// 		{
// 			Method: "GET",
// 			Url:    "/user/:uniqueId",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpGetOneUser,
// 			},
// 			ResponseEntity: &UserEntity{},
// 		},
// 		{

// 			Method: "GET",
// 			Url:    "/exchangeKey/:uniqueId",
// 			Handlers: []gin.HandlerFunc{
// 				HttpGetExchangeKey,
// 			},
// 			ResponseEntity: &ExchangeKeyInformationDto{},
// 		},
// 		{

// 			Method: "POST",
// 			Url:    "/authorization",
// 			Handlers: []gin.HandlerFunc{
// 				HttpGetAuthorizationInformation,
// 			},
// 			RequestEntity:  &EmptyRequest{},
// 			ResponseEntity: &AuthResult{},
// 		},
// 		{

// 			Method: "GET",
// 			Url:    "profile",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpGetUserProfile,
// 			},
// 			ResponseEntity: &UserEntity{},
// 		},
// 		{

// 			Method: "POST",
// 			Url:    "profile",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpPostUpdateUserProfile,
// 			},
// 			RequestEntity:  &UserProfileEntity{},
// 			ResponseEntity: &UserProfileEntity{},
// 		},
// 		{

// 			Method: "DELETE",
// 			Url:    "/auth/revoke",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpRemoveToken,
// 			},
// 			RequestEntity:  &DeleteRequest{},
// 			ResponseEntity: &DeleteResponse{},
// 			TargetEntity:   &Token{},
// 		},
// 		{

// 			Method: "POST",
// 			Url:    "/preferences",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				HttpPostPreferences,
// 			},
// 			RequestEntity:  map[string]interface{}{},
// 			ResponseEntity: map[string]interface{}{},
// 		},
// 		{

// 			Method: "GET",
// 			Url:    "/preferences",
// 			Handlers: []gin.HandlerFunc{
// 				WithUserAuthorization(true),
// 				httpGetPreferences,
// 			},
// 			ResponseEntity: map[string]interface{}{},
// 		},
// 	}

// 	CastRoutes(httpRoutes, r)
// 	WriteHttpInformationToFile(&httpRoutes, []EntityJsonField{}, "user-http", "workspaces")

// }

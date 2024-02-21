package workspaces

import (
	context "context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var ROOT_VAR = "root"

// Bare minimum handler
func WithAuthorizationPure(context *AuthContextDto) (*AuthResultDto, *IError) {
	result := &AuthResultDto{}

	// workspaceId := context.WorkspaceId
	token := context.Token

	if token == nil || *token == "" {
		return nil, CreateIErrorString(WorkspacesMessageCode.ProvideTokenInAuthorization, []string{}, 401)
	}

	user, err := GetUserFromToken(*token)
	if err != nil {
		return nil, CreateIErrorString(WorkspacesMessageCode.UserWhichHasThisTokenDoesNotExist, []string{}, 401)
	}

	if user == nil {
		return nil, CreateIErrorString(WorkspacesMessageCode.UserNotFoundOrDeleted, []string{}, 401)
	}

	access, _ := GetUserAccessLevels(QueryDSL{UserId: user.UniqueId})

	if len(access.Workspaces) == 1 {
		result.WorkspaceId = &access.Workspaces[0]
		if *result.WorkspaceId == "*" {
			result.WorkspaceId = context.WorkspaceId
			if *result.WorkspaceId == "" {
				result.WorkspaceId = &ROOT_VAR
			}
		}

	} else {
		result.WorkspaceId = context.WorkspaceId
	}

	query := QueryDSL{
		UserHas:        access.Capabilities,
		ActionRequires: context.Capabilities,
	}

	meets, missing := MeetsAccessLevel(query, false)

	if err != nil || !meets {
		return nil, CreateIErrorString("NOT_ENOUGH_PERMISSION", missing, 401)
	}

	result.AccessLevel = access
	result.InternalSql = access.SQL
	result.UserId = &user.UniqueId
	result.User = user
	result.UserHas = access.Capabilities

	return result, nil
}

// For go http package
// byPassGetMethod means, that the get method does not need to be authorized.
// used for accessing files on the disk, because the unique id is as long as the token
// itself
func WithAuthorizationHttp(next http.Handler, byPassGetMethod bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With, Workspace, Workspace-Id, role-id, workspace-id")

		if r.Method == http.MethodOptions || (r.Method == http.MethodGet && byPassGetMethod) {
			next.ServeHTTP(w, r)
			return
		}

		wi := r.Header.Get("Workspace-id")
		tk := r.Header.Get("authorization")

		context := &AuthContextDto{
			WorkspaceId:  &wi,
			Token:        &tk,
			Capabilities: []string{},
		}

		_, err := WithAuthorizationPure(context)
		if err != nil {
			f, _ := json.MarshalIndent(gin.H{"error": err}, "", "  ")
			http.Error(w, string(f), int(err.HttpCode))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Combine all for gin
func WithAuthorization(capabilities []string) gin.HandlerFunc {
	return WithAuthorizationFn(capabilities, false)
}
func WithAuthorizationSkip(capabilities []string) gin.HandlerFunc {
	return WithAuthorizationFn(capabilities, true)
}

var USER_SYSTEM = "system"

func WithSocketAuthorization(capabilities []string, skipWorkspaceId bool) gin.HandlerFunc {
	if os.Getenv("BYPASS_WORKSPACES") == "YES" {
		return func(c *gin.Context) {
			c.Set("user_id", &USER_SYSTEM)
			c.Set("workspaceId", "SYSTEM")
		}
	}

	return func(c *gin.Context) {

		wsURL := c.Request.URL
		wsURLParam, err3 := url.ParseQuery(wsURL.RawQuery)

		workspaceId := c.Request.Header.Get("Workspace-id")
		token := c.Request.Header.Get("authorization")

		if err3 == nil && wsURLParam["token"] != nil && len(wsURLParam["token"]) == 1 {

			token = wsURLParam["token"][0]
		}

		if err3 == nil && wsURLParam["workspaceId"] != nil && len(wsURLParam["workspaceId"]) == 1 {

			workspaceId = wsURLParam["workspaceId"][0]
		}

		context := &AuthContextDto{
			WorkspaceId:  &workspaceId,
			Token:        &token,
			Capabilities: capabilities,
		}

		result, err := WithAuthorizationPure(context)

		if err != nil {
			c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err})
			return
		}

		c.Set("user_has", result.UserHas)
		c.Set("internal_sql", *result.InternalSql)
		c.Set("user_id", result.UserId)
		c.Set("user", result.User)
		c.Set("authResult", result)
		c.Set("workspaceId", result.WorkspaceId)

	}
}

func WithAuthorizationFn(capabilities []string, skipWorkspaceId bool) gin.HandlerFunc {
	if os.Getenv("BYPASS_WORKSPACES") == "YES" {
		return func(c *gin.Context) {
			c.Set("user_id", &USER_SYSTEM)
			c.Set("workspaceId", "SYSTEM")
		}
	}

	return func(c *gin.Context) {
		wi := c.GetHeader("Workspace-id")
		tk := c.GetHeader("Authorization")
		context := &AuthContextDto{
			WorkspaceId:     &wi,
			Token:           &tk,
			Capabilities:    capabilities,
			SkipWorkspaceId: &skipWorkspaceId,
		}
		result, err := WithAuthorizationPure(context)

		if err != nil {
			c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err})
			return
		}

		c.Set("user_has", result.UserHas)
		c.Set("internal_sql", *result.AccessLevel.SQL)
		c.Set("user_id", *result.UserId)
		c.Set("user", result.User)
		c.Set("authResult", result)
		c.Set("workspaceId", result.WorkspaceId)

	}
}

func GrpcContextToAuthContext(ctx *context.Context, capabilities []string) *AuthContextDto {
	data, _ := metadata.FromIncomingContext(*ctx)

	authList := data.Get("authorization")
	authValue := ""
	if len(authList) > 0 {
		authValue = authList[0]
	}

	workspaceList := data.Get("workspace-id")
	workspaceValue := ""
	if len(workspaceList) > 0 {
		workspaceValue = workspaceList[0]
	}

	context := AuthContextDto{
		WorkspaceId:  &workspaceValue,
		Token:        &authValue,
		Capabilities: capabilities,
	}

	return &context
}

func GrpcWithAuthorization(ctx *context.Context, capabilities []string) (QueryDSL, *IError) {

	auth, err := WithAuthorizationPure(GrpcContextToAuthContext(ctx, capabilities))

	if err == nil {
		return ExtractQueryDslFromGrpcContext(ctx, auth), nil
	}

	return QueryDSL{}, err

}

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	data, _ := metadata.FromIncomingContext(ctx)

	authList := data.Get("authorization")
	authValue := ""
	if len(authList) > 0 {
		authValue = authList[0]
	}

	workspaceList := data.Get("workspace-id")
	workspaceValue := ""
	if len(workspaceList) > 0 {
		workspaceValue = workspaceList[0]
	}

	context := &AuthContextDto{
		WorkspaceId:  &workspaceValue,
		Token:        &authValue,
		Capabilities: []string{},
	}
	result, err2 := WithAuthorizationPure(context)

	fmt.Println(result)

	if err2 != nil {
		return nil, errors.New("Error")
	} else {

		fmt.Println("Authorized")
		// 	// c.Set("internal_sql", result.InternalSql)
		// 	// c.Set("user_id", result.UserId)
		// 	// c.Set("user", result.User)
		// 	// c.Set("workspaceId", result.WorkspaceId)
	}

	h, err := handler(ctx, req)

	return h, err
}

func WithServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

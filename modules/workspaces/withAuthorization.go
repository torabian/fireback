package workspaces

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
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

	access, accessError := GetUserAccessLevels(QueryDSL{UserId: user.UniqueId})

	if accessError != nil {
		return nil, accessError
	}

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
	result.UserRoleWorkspacePermissions = access.UserRoleWorkspacePermissions

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
			Capabilities: []PermissionInfo{},
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
func WithAuthorization(securityModel *SecurityModel) gin.HandlerFunc {
	return WithAuthorizationFn(securityModel, false)
}
func WithAuthorizationSkip(securityModel *SecurityModel) gin.HandlerFunc {
	return WithAuthorizationFn(securityModel, true)
}

var USER_SYSTEM = "system"

func WithSocketAuthorization(securityModel *SecurityModel, skipWorkspaceId bool) gin.HandlerFunc {
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
		uniqueId := c.Request.Header.Get("uniqueId")

		if err3 == nil && wsURLParam["token"] != nil && len(wsURLParam["token"]) == 1 {

			token = wsURLParam["token"][0]
		}

		if err3 == nil && wsURLParam["workspaceId"] != nil && len(wsURLParam["workspaceId"]) == 1 {

			workspaceId = wsURLParam["workspaceId"][0]
		}

		if err3 == nil && wsURLParam["uniqueId"] != nil && len(wsURLParam["uniqueId"]) == 1 {
			uniqueId = wsURLParam["uniqueId"][0]
		}

		context := &AuthContextDto{
			WorkspaceId:  &workspaceId,
			Token:        &token,
			Capabilities: securityModel.ActionRequires,
		}

		result, err := WithAuthorizationPure(context)

		if err != nil {
			c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err})
			return
		}

		c.Set("urw", result.UserRoleWorkspacePermissions)
		c.Set("user_has", result.UserHas)
		c.Set("internal_sql", *result.InternalSql)
		c.Set("user_id", result.UserId)
		c.Set("user", result.User)
		c.Set("uniqueId", uniqueId)
		c.Set("authResult", result)
		c.Set("workspaceId", result.WorkspaceId)

	}
}

func WithAuthorizationFn(securityModel *SecurityModel, skipWorkspaceId bool) gin.HandlerFunc {
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
			Capabilities:    securityModel.ActionRequires,
			SkipWorkspaceId: &skipWorkspaceId,
		}
		result, err := WithAuthorizationPure(context)

		if err != nil {
			c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err})
			return
		}

		c.Set("urw", result.UserRoleWorkspacePermissions)
		c.Set("resolveStrategy", securityModel.ResolveStrategy)
		c.Set("user_has", result.UserHas)
		c.Set("internal_sql", *result.AccessLevel.SQL)
		c.Set("user_id", *result.UserId)
		c.Set("user", result.User)
		c.Set("authResult", result)
		c.Set("workspaceId", result.WorkspaceId)

	}
}

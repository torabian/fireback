package abac

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
)

var ROOT_VAR = "root"

// Bare minimum handler

type WithAuthorizationPureImpl func(context *fireback.AuthContextDto) (*fireback.AuthResultDto, *fireback.IError)

func maskToken(token string) string {
	if len(token) <= 6 {
		return token // too short to mask meaningfully
	}
	return token[:2] + "***" + token[len(token)-4:]
}

// Default authorization compute function for Fireback ABAC.
// You can get inspired by this function and make the authorization
func WithAuthorizationPureDefault(context *fireback.AuthContextDto) (*fireback.AuthResultDto, *fireback.IError) {
	result := &fireback.AuthResultDto{}

	// workspaceId := context.WorkspaceId
	token := context.Token

	if token == "" {
		return nil, fireback.Create401Error(&AbacMessages.ProvideTokenInAuthorization, []string{})
	}

	user, err := GetUserFromToken(token)

	if err != nil {
		return nil, fireback.Create401Error(&AbacMessages.TokenNotFound, []string{
			maskToken(token),
		})
	}

	if user == nil {
		return nil, fireback.Create401Error(&AbacMessages.UserNotFoundOrDeleted, []string{})
	}

	access, accessError := GetUserAccessLevels(fireback.QueryDSL{UserId: user.UniqueId})

	if accessError != nil {
		return nil, accessError
	}

	query := fireback.QueryDSL{
		UserAccessPerWorkspace: access.UserAccessPerWorkspace,
		ActionRequires:         context.Capabilities,
	}

	meets, missing := fireback.MeetsAccessLevel(query, false)

	if !meets {
		return nil, fireback.Create401Error(&AbacMessages.NotEnoughPermission, missing)
	}

	result.UserId = fireback.NewString(user.UniqueId)
	result.User = user
	result.UserAccessPerWorkspace = access.UserAccessPerWorkspace
	result.SqlContext = GetSqlContext(access.UserAccessPerWorkspace, context.WorkspaceId, context.AllowCascade)

	// some actions could be restricted to happen only on root workspaces
	// here we check if user belongs to root, and the workspace selected needs to be root
	// as well. In some cases, user is in root workspace, but also exploring
	// another workspace for maintenance, should not be able to create root level content
	// in another workspace.

	// Fix this allow only on root.
	if context.Security != nil && context.Security.AllowOnRoot {
		if context.WorkspaceId != ROOT_VAR {
			return nil, &fireback.IError{
				HttpCode: 400,
				Message:  AbacMessages.ActionOnlyInRoot,
			}
		}
	}

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

		context := &fireback.AuthContextDto{
			WorkspaceId:  wi,
			Token:        tk,
			Capabilities: []fireback.PermissionInfo{},
		}

		_, err := WithAuthorizationPureDefault(context)
		if err != nil {
			f, _ := json.MarshalIndent(gin.H{"error": err}, "", "  ")
			http.Error(w, string(f), int(err.HttpCode))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func WithAuthorization(securityModel *fireback.SecurityModel) gin.HandlerFunc {
	return WithAuthorizationFn(securityModel)
}

var USER_SYSTEM = "system"

func WithSocketAuthorization(securityModel *fireback.SecurityModel) gin.HandlerFunc {

	return func(c *gin.Context) {

		wsURL := c.Request.URL
		wsURLParam, err3 := url.ParseQuery(wsURL.RawQuery)

		workspaceId := c.Request.Header.Get("Workspace-id")
		token := c.Request.Header.Get("authorization")
		uniqueId := c.Request.Header.Get("uniqueId")

		if err3 == nil && wsURLParam["token"] != nil && len(wsURLParam["token"]) == 1 {

			token = wsURLParam["token"][0]
		}

		if err3 == nil && wsURLParam["authorization"] != nil && len(wsURLParam["authorization"]) == 1 {

			token = wsURLParam["authorization"][0]
		}

		if err3 == nil && wsURLParam["workspaceId"] != nil && len(wsURLParam["workspaceId"]) == 1 {

			workspaceId = wsURLParam["workspaceId"][0]
		}

		if err3 == nil && wsURLParam["uniqueId"] != nil && len(wsURLParam["uniqueId"]) == 1 {
			uniqueId = wsURLParam["uniqueId"][0]
		}

		context := &fireback.AuthContextDto{
			WorkspaceId:  workspaceId,
			Token:        token,
			Capabilities: securityModel.ActionRequires,
			Security:     securityModel,
		}

		result, err := WithAuthorizationPureDefault(context)

		if err != nil {
			c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err})
			return
		}

		c.Set("internal_sql", result.SqlContext)
		c.Set("urw", result.UserAccessPerWorkspace)
		c.Set("user_id", result.UserId.String)
		c.Set("uniqueId", uniqueId)
		c.Set("workspaceId", workspaceId)
		c.Set("workspace-id", workspaceId)

	}
}

func WithAuthorizationFn(securityModel *fireback.SecurityModel) gin.HandlerFunc {

	return func(c *gin.Context) {

		q := fireback.ExtractQueryDslFromGinContext(c)
		wi := c.GetHeader("Workspace-id")
		ri := c.GetHeader("Role-id")
		tk := c.GetHeader("Authorization")
		ck, ckerr := c.Cookie("authorization")

		if ckerr == nil && ck != "" {
			// If on secure cookie we have the authorization, we prefer that one.
			tk = ck
		}

		context := &fireback.AuthContextDto{
			WorkspaceId:  wi,
			Token:        tk,
			Capabilities: securityModel.ActionRequires,
			Security:     securityModel,
		}

		result, err := WithAuthorizationPureDefault(context)

		if err != nil {
			fmt.Println("Aborting: ", err.ToPublicEndUser(&q))
			c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&q)})
			fmt.Println("Aborted")
			return
		}

		c.Set("urw", result.UserAccessPerWorkspace)
		c.Set("resolveStrategy", securityModel.ResolveStrategy)
		c.Set("internal_sql", result.SqlContext)
		c.Set("role_id", ri)
		c.Set("user_id", result.UserId.String)
		c.Set("authResult", result)
		c.Set("workspaceId", wi)
	}
}

// It would convert the current selected role_id and workspace_id into a sql
// with given permissions to make the queries do not need check that again
func GetSqlContext(x *fireback.UserAccessPerWorkspaceDto, activeWorkspaceId string, allowCascade bool) string {
	conditions := []string{

		// Visibility A means that the content is accessible across the entire project.
		// It's a public content.
		`visibility = "A"`,
	}

	// Let's allow the user to see everything which they belong to
	// but usually it's not necessary, because they are focused on one workspace at the moment
	if allowCascade {
		for workspaceId := range *x {
			conditions = append(conditions, "workspace_id in (\""+workspaceId+"\")")
		}
	} else {
		userBelongsToWorkspace := false
		for workspaceId := range *x {
			if workspaceId == activeWorkspaceId {
				userBelongsToWorkspace = true

				// Important to break, otherwise can show other workspaces
				break
			}
		}

		if userBelongsToWorkspace {
			conditions = append(conditions, "workspace_id in (\""+activeWorkspaceId+"\")")
		}
	}

	sql := strings.Join(conditions, " or ")

	return sql
}

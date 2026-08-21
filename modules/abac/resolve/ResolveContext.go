package resolve

import (
	"encoding/json"
	"log"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	queries "github.com/torabian/fireback/modules/abac/queries"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

var ROOT_ALL_ACCESS = "root.*"
var ROOT_VAR = "root"

func ResolveActionContext(request emigo.EmiRequestContexts, securityModel *SecurityModel) (*AuthResultDto, error) {
	GinCtx := request.GetGinCtx()
	CliCtx := request.GetCliCtx()

	if value, ok := GinCtx.(*gin.Context); ok {
		if securityModel != nil && !fireback.IsNilish(GinCtx) {
			// Fireback no longer plans to check security anymore. Do it manually,
			// on every action. This is fr transparency.
			return AuthorizeRequest(securityModel, value)

		}
	}

	if !fireback.IsNilish(CliCtx) {
		// Fireback no longer plans to check security anymore. Do it manually,
		// on every action. This is fr transparency.
		if value, ok := CliCtx.(*cli.Command); ok {
			return CommonCliQueryDSLBuilderAuthorize(value, securityModel)
		}
	}

	return nil, nil
}
func ExtractQueryDslFromGinContext(c *gin.Context) QueryDSL {
	workspaceId := c.GetString("workspaceId")
	id := c.Param("uniqueId")

	resolveStrategy := c.GetString("resolveStrategy")
	userHas := c.GetStringSlice("user_has")
	workspaceHas := c.GetStringSlice("workspace_has")

	user, isUserSet := c.Get("user_id")
	var userId string

	if isUserSet {
		value, ok := user.(string)
		if ok {
			userId = value
		} else if value2, ok2 := user.(*string); ok2 {
			userId = *value2
		}
	}

	uniqueID := c.GetString("uniqueId")
	if uniqueID != "" {
		id = uniqueID
	}

	var f QueryDSL = QueryDSL{

		UserHas:      userHas,
		WorkspaceHas: workspaceHas,
		WorkspaceId:  workspaceId,

		Language:      "en",
		Region:        "us",
		UniqueId:      id,
		Authorization: c.GetHeader("Authorization"),
	}

	if resolveStrategy != "" {
		f.ResolveStrategy = resolveStrategy
	} else {
		f.ResolveStrategy = ResolveStrategyWorkspace
	}

	f.UserId = userId

	acceptLang := c.GetHeader("accept-language")
	if acceptLang != "" && len(acceptLang) == 2 {
		f.Language = strings.ToLower(acceptLang)
	}

	// The language set in the header has higher priority
	language := c.Query("acceptLanguage")
	if language != "" {
		f.Language = language
	}

	return f
}

type AccessState struct {
}

func AuthorizeRequest(securityModel *SecurityModel, c *gin.Context) (*AuthResultDto, error) {
	q := fireback.ExtractQueryDslFromGinContext(c)

	// A WebSocket handshake can't carry custom headers from the browser, so
	// clients of a reactive action (e.g. eventbus's /ws, opened by the UI's
	// useFirebackSocket) send token/workspace-id/role-id as query params
	// instead. And by the time a reactive action's factory calls into this
	// function (see EventBusSubscriptionActionSig), gorilla-websocket has
	// already hijacked the connection to complete the upgrade - gin's
	// response writer can no longer send an HTTP response onto it, so the
	// c.AbortWithStatusJSON below would crash with "response.Write on
	// hijacked connection" if reached. Both quirks trace back to the same
	// signal, checked once here: is this request a websocket upgrade.
	isWebsocketUpgrade := websocket.IsWebSocketUpgrade(c.Request)

	wi := c.GetHeader("Workspace-id")
	ri := c.GetHeader("Role-id")
	tk := c.GetHeader("Authorization")

	if isWebsocketUpgrade {
		if wi == "" {
			wi = c.Query("workspaceId")
		}
		if ri == "" {
			ri = c.Query("roleId")
		}
		if tk == "" {
			tk = c.Query("token")
		}
	}

	ck, ckerr := c.Cookie("authorization")

	if ckerr == nil && ck != "" && tk == "" {
		// If on secure cookie we have the authorization, and it's not defined on headers or query params.
		tk = ck
	}

	context := &AuthContextDto{
		WorkspaceId:  wi,
		Token:        tk,
		Capabilities: securityModel.ActionRequires,
		Security:     securityModel,
	}

	result, err := WithAuthorizationPureDefault(context)

	if err != nil {
		if isWebsocketUpgrade {
			// Already hijacked - can't send an HTTP response. The caller
			// (the reactive action's generated Gin handler) reports the
			// failure as a websocket frame instead, using the real
			// connection rather than gin's now-unusable writer.
			return nil, nil
		}
		c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&q)})
		return nil, nil
	}

	return result, nil
}

func WithAuthorizationFn(securityModel *SecurityModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthorizeRequest(securityModel, c)
	}
}

type Translatable struct{}

func (x Translatable) GetLanguage() string {
	lang := "en"
	if fireback.GetConfig().CliLanguage != "" {
		lang = fireback.GetConfig().CliLanguage
	}

	return lang
}

func CommonCliQueryDSLBuilderAuthorize(c *cli.Command, security *SecurityModel) (*AuthResultDto, error) {

	t := Translatable{}

	if security != nil && security.ResolveStrategy != ResolveStrategyPublic {

		context := &AuthContextDto{
			WorkspaceId:  fireback.GetConfig().CliWorkspace,
			Token:        fireback.GetConfig().CliToken,
			Capabilities: []application.PermissionInfo{},
			Security:     security,
		}

		result, err := WithAuthorizationPureDefault(context)

		if err != nil {

			if err.ToPublicEndUser(t).Message != err.ToPublicEndUser(t).MessageTranslated {
				log.Fatalf("%s", err.ToPublicEndUser(t).Message)
			}
			log.Default().Printf("%s", err.ToPublicEndUser(t).MessageTranslated)
		}

		return result, nil

	}

	return nil, nil
}

func maskToken(token string) string {
	if len(token) <= 6 {
		return token // too short to mask meaningfully
	}
	return token[:2] + "***" + token[len(token)-4:]
}

func GetUserFromToken(tokenString string) (*abacdefs.UserEntity, error) {

	var item abacdefs.TokenEntity

	if err := fireback.GetDbRef().Where(fireback.RealEscape("token = ?", tokenString)).First(&item).Error; err != nil {
		return &abacdefs.UserEntity{}, err
	}

	// Not workspace-scoped (see UserBrowseAction's own comment) - abacdefs.UserEntityActions.Get
	// is the entity's own generated, unscoped lookup.
	user, _ := abacdefs.UserEntityActions.Get(fireback.GetDbRef(), item.UserId.OrDefault(""))

	// HydrateUserPrimaryAddress(user)
	return user, nil
}

func WithAuthorizationPureDefault(context *AuthContextDto) (*AuthResultDto, *fireback.IError) {
	result := &AuthResultDto{}

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

	access, accessError := GetUserAccessLevels(QueryDSL{UserId: user.UniqueId})

	if accessError != nil {
		return nil, accessError
	}

	query := QueryDSL{
		UserAccessPerWorkspace: access.UserAccessPerWorkspace,
		ActionRequires:         context.Capabilities,
		WorkspaceId:            context.WorkspaceId,
	}

	// MeetsAccessLevel checks query.UserHas/WorkspaceHas directly - they don't
	// come from UserAccessPerWorkspace automatically, so they need to be flattened
	// out of it for the *active* workspace (context.WorkspaceId) first. Without
	// this, both are always empty, which - combined with the now-fixed
	// MeetsAccessLevel actually enforcing its verdict instead of always returning
	// true - would deny every capability-gated action for everyone, root included.
	query.WorkspaceHas, query.UserHas = GetWorkspaceAndUserAccesses(query)

	meets, missing := MeetsAccessLevel(query, false)

	if !meets {
		return nil, fireback.Create401Error(&AbacMessages.NotEnoughPermission, missing)
	}

	result.UserId = emigo.NullableOf(user.UniqueId)
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
		if context.WorkspaceId != fireback.ROOT_VAR {
			return nil, &fireback.IError{
				HttpCode: 400,
				Message:  AbacMessages.ActionOnlyInRoot,
			}
		}
	}

	return result, nil
}

// It would convert the current selected role_id and workspace_id into a sql
// with given permissions to make the queries do not need check that again
func GetSqlContext(x *UserAccessPerWorkspaceDto, activeWorkspaceId string, allowCascade bool) string {
	conditions := []string{}

	// Let's allow the user to see everything which they belong to
	// but usually it's not necessary, because they are focused on one workspace at the moment
	if allowCascade {
		for workspaceId := range *x {
			conditions = append(conditions, fireback.RealEscape("workspace_id in (?)", workspaceId))
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
			conditions = append(conditions, fireback.RealEscape("workspace_id in (?)", activeWorkspaceId))
		}
	}

	sql := strings.Join(conditions, " or ")

	return sql
}

func MeetsAccessLevel(query QueryDSL, onlyRoot bool) (bool, []string) {
	if onlyRoot && (query.WorkspaceId != ROOT_VAR && query.WorkspaceId != "system") {
		return false, []string{"SYSTEM_OR_ROOT_ALLOWED"}
	}

	missingPerms := []string{}

	// A user (and their active workspace) holding the full root.* wildcard always
	// meets every requirement - no need to check individual capabilities. This was
	// previously inverted (returned false, i.e. denied, for exactly this case) and
	// the real verdict below was discarded entirely (always returned true) - see
	// the fix note on the final return.
	if slices.Contains(query.UserHas, ROOT_ALL_ACCESS) && slices.Contains(query.WorkspaceHas, ROOT_ALL_ACCESS) {
		return true, missingPerms
	}

	meetsUser := MeetsCheck(query.ActionRequires, query.UserHas)
	meetsWorkspace := MeetsCheck(query.ActionRequires, query.WorkspaceHas)

	if !meetsUser || !meetsWorkspace {
		for _, perm := range query.ActionRequires {
			if slices.Contains(query.UserHas, perm.CompleteKey) {
				continue
			}

			missingPerms = append(missingPerms, perm.CompleteKey)
		}
	}

	// Bug fix: this always returned true regardless of meetsUser/meetsWorkspace,
	// so every action requiring a capability the caller didn't actually have was
	// silently allowed through - the only real gate left was the onlyRoot check
	// above. missingPerms was still computed correctly, just never acted on.
	return meetsUser && meetsWorkspace, missingPerms
}

func MeetsCheck(actionRequires []application.PermissionInfo, perms []string) bool {
	meets := true
	for _, requiredPermission := range actionRequires {

		// Two things needs to be checked, first if it does contain exact capability
		hasExactKey := slices.Contains(perms, requiredPermission.CompleteKey)
		hasParentalKey := false

		for _, a := range perms {
			if strings.Contains(requiredPermission.CompleteKey, strings.ReplaceAll(a, "*", "")) {
				hasParentalKey = true
				continue
			}
		}

		if !hasExactKey && !hasParentalKey {
			meets = false
		}
	}
	return meets
}

func GetWorkspaceAndUserAccesses(query QueryDSL) ([]string, []string) {

	if query.UserAccessPerWorkspace == nil {
		return []string{}, []string{}
	}

	data := *query.UserAccessPerWorkspace
	workspaceAccesses := []string{}
	rolesPermission := []string{}
	if data[query.WorkspaceId] != nil {
		workspaceAccesses = data[query.WorkspaceId].WorkspacesAccesses

		// Now we are checking with all the roles user has, but need to have access to role id
		// and only look for that.
		for _, role := range data[query.WorkspaceId].UserRoles {
			rolesPermission = append(rolesPermission, role.Accesses...)
		}
	}

	return workspaceAccesses, rolesPermission
}

func GetUserAccessLevels(query QueryDSL) (*UserAccessLevelDto, *fireback.IError) {

	access := &UserAccessLevelDto{}
	query.ItemsPerPage = 1000

	items, _, err := fireback.UnsafeQuerySqlFromFs[UserRoleWorkspacePermissionDto](
		&queries.QueriesFs, "UserRolePermission", fireback.QueryDSL{
			UserId:        query.UserId,
			WorkspaceId:   query.WorkspaceId,
			InternalQuery: query.InternalQuery,
		},
	)

	if err != nil {
		return nil, fireback.CastToIError(err)
	}

	ws := UserAccessPerWorkspaceDto{}

	for _, item := range items {
		if ws[item.WorkspaceId] == nil {
			ws[item.WorkspaceId] = &struct {
				Name               string
				WorkspacesAccesses []string
				UserRoles          map[string]*struct {
					Name     string
					Accesses []string
				}
			}{}
		}

		ws[item.WorkspaceId].Name = item.WorkspaceName

		if item.Type == "account_restrict" {
			// Bug fix: this used to re-init the whole UserRoles map (wiping every
			// other role already accumulated for this same workspace) whenever it hit
			// a *second* role_id it hadn't seen yet - the map-is-nil check and the
			// key-is-nil check were conflated into one condition. In practice this
			// meant a user holding more than one role in the same workspace (e.g. via
			// a Workspaceabacdefs.RoleEntity granting an extra role alongside their normal one)
			// silently lost every role but the last one processed - for root, that
			// could drop the seeded "root.*" wildcard role itself, causing
			// MeetsAccessLevel to reject actions root should always be allowed.
			if ws[item.WorkspaceId].UserRoles == nil {
				ws[item.WorkspaceId].UserRoles = map[string]*struct {
					Name     string
					Accesses []string
				}{}
			}
			if ws[item.WorkspaceId].UserRoles[item.RoleId] == nil {
				ws[item.WorkspaceId].UserRoles[item.RoleId] = &struct {
					Name     string
					Accesses []string
				}{}
			}
			ws[item.WorkspaceId].UserRoles[item.RoleId].Accesses = append(ws[item.WorkspaceId].UserRoles[item.RoleId].Accesses, item.CapabilityId)
			ws[item.WorkspaceId].UserRoles[item.RoleId].Name = item.RoleName
		}

		if item.Type == "workspace_restrict" {
			ws[item.WorkspaceId].WorkspacesAccesses = append(ws[item.WorkspaceId].WorkspacesAccesses, item.CapabilityId)
		}
	}

	access.UserAccessPerWorkspace = &ws

	return access, nil
}

var ResolveStrategyPublic = "public"
var ResolveStrategyUser = "user"
var ResolveStrategyWorkspace = "workspace"

// Used for actions generally
type SecurityModel struct {
	// Only users which belong to root and actively selected the root workspace can
	// write to this entity from Fireback default functionality
	AllowOnRoot bool `json:"allowOnRoot,omitempty" yaml:"allowOnRoot,omitempty"`

	// Set of permissions which are required for this service.
	ActionRequires []application.PermissionInfo `json:"requires,omitempty" yaml:"requires,omitempty"`

	// Resolve strategy is by default on the workspace, you can change it by user
	// also. Be sure of the consequences
	ResolveStrategy string `json:"resolveStrategy,omitempty" yaml:"resolveStrategy,omitempty"`
}

type UserAccessPerWorkspaceDto map[string]*struct {
	Name string
	// The access which are available to this workspace, not to the specific user.
	// Even a user has access to many things, these accesses need to reduce those
	WorkspacesAccesses []string

	// The permissions which user has access to
	UserRoles map[string]*struct {
		Name     string
		Accesses []string
	}
}

func (x UserAccessPerWorkspaceDto) Json() string {
	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))

}

type UserRoleWorkspacePermissionDto struct {
	WorkspaceName string `json:"workspaceName" yaml:"workspaceName"        `
	WorkspaceId   string `json:"workspaceId" yaml:"workspaceId"        `
	RoleName      string `json:"roleName" yaml:"roleName"        `
	UserId        string `json:"userId" yaml:"userId"        `
	RoleId        string `json:"roleId" yaml:"roleId"        `
	CapabilityId  string `json:"capabilityId" yaml:"capabilityId"        `
	Type          string `json:"type" yaml:"type"        `
}
type UserRoleWorkspacePermissionDtoList struct {
	Items []*UserRoleWorkspacePermissionDto
}

type UserAccessLevelDto struct {
	UserAccessPerWorkspace   *UserAccessPerWorkspaceDto `json:"userAccessPerWorkspace" yaml:"userAccessPerWorkspace"    gorm:"foreignKey:UserAccessPerWorkspaceId;references:UniqueId"      `
	UserAccessPerWorkspaceId emigo.Nullable[string]     `json:"userAccessPerWorkspaceId" yaml:"userAccessPerWorkspaceId"`
}
type UserAccessLevelDtoList struct {
	Items []*UserAccessLevelDto
}

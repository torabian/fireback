package resolve

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/interfacetools/queries"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/ferror"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

var ROOT_ALL_ACCESS = "root.*"
var ROOT_VAR = "root"

func ExtractQueryDslFromGinContext(c *gin.Context) QueryDSL {
	workspaceId := c.GetString("workspaceId")
	id := c.Param("uniqueId")
	sort := c.Query("sort")

	resolveStrategy := c.GetString("resolveStrategy")
	linkerId := c.Param("linkerId")
	queryString, _ := c.GetQuery("query")
	withPreloads, _ := c.GetQuery("withPreloads")
	isDeep, _ := c.GetQuery("deep")

	searchPhrase := c.Query("searchPhrase")

	o, _ := c.GetQuery("startIndex")
	startIndex, _ := strconv.Atoi(o)

	cursor, _ := c.GetQuery("cursor")

	l, _ := c.GetQuery("itemsPerPage")
	itemsPerPage, _ := strconv.Atoi(l)

	if startIndex < 0 {
		startIndex = 0
	}

	switch {
	case itemsPerPage > 1000:
		itemsPerPage = 1000
	case itemsPerPage <= 0:
		itemsPerPage = 20
	}

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

	var urw *UserAccessPerWorkspaceDto
	if value, exists := c.Get("urw"); exists {
		if casted, ok := value.(*UserAccessPerWorkspaceDto); ok {
			urw = casted
		}
	}

	var f QueryDSL = QueryDSL{
		Query:        queryString,
		StartIndex:   startIndex,
		ItemsPerPage: itemsPerPage,

		G:                      c,
		UserAccessPerWorkspace: urw,
		Cursor:                 &cursor,
		UserHas:                userHas,
		WorkspaceHas:           workspaceHas,
		Sort:                   sort,
		SearchPhrase:           searchPhrase,
		LinkerId:               linkerId,
		WorkspaceId:            workspaceId,

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

	if len(withPreloads) > 0 {
		f.WithPreloads = strings.Split(strings.Trim(withPreloads, " "), ",")
	}

	deep := c.GetHeader("deep")

	if deep == "true" || deep == "yes" || deep == "1" || isDeep == "true" || isDeep == "yes" || isDeep == "1" {
		f.Deep = true
	}

	query := c.GetHeader("query")
	if query != "" && f.Query == "" {
		f.Query = query
	}

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

// Keeps the functionality to resolve access level.

// ResolveActionContextFromGinContext resolves the QueryDSL for a plain
// *gin.Context, enforcing the security model if one is given. This holds the
// gin-specific logic shared by ResolveActionContext so it isn't duplicated
// wherever only a *gin.Context (and no emigo.EmiRequestContexts) is available.
func ResolveActionContextFromGinContext(ginCtx *gin.Context, securityModel *SecurityModel) (*QueryDSL, error) {
	qsdl := ExtractQueryDslFromGinContext(ginCtx)

	// Only handles the gin security, and this is a problem needs to handle cli as well
	// perfectly
	if securityModel != nil && !fireback.IsNilish(ginCtx) {
		// Fireback no longer plans to check security anymore. Do it manually,
		// on every action. This is fr transparency.
		if !AuthorizeRequest(securityModel, ginCtx) {
			return nil, errors.New("Authorization general failed")
		}

		// Important because now we have more details of security
		qsdl = ExtractQueryDslFromGinContext(ginCtx)
	}

	return &qsdl, nil
}

func ResolveActionContext(request emigo.EmiRequestContexts, securityModel *SecurityModel) (*QueryDSL, error) {
	GinCtx := request.GetGinCtx()
	CliCtx := request.GetCliCtx()

	var qsdl QueryDSL

	if value, ok := GinCtx.(*gin.Context); ok {
		resolved, err := ResolveActionContextFromGinContext(value, securityModel)
		if err != nil {
			return nil, err
		}
		qsdl = *resolved
	}

	if !fireback.IsNilish(CliCtx) {
		// Fireback no longer plans to check security anymore. Do it manually,
		// on every action. This is fr transparency.
		if value, ok := CliCtx.(*cli.Command); ok {
			qsdl = CommonCliQueryDSLBuilderAuthorize(value, securityModel)
		}
	}

	return &qsdl, nil
}

func AuthorizeRequest(securityModel *SecurityModel, c *gin.Context) bool {
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
			return false
		}
		c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&q)})
		return false
	}

	c.Set("urw", result.UserAccessPerWorkspace)
	c.Set("resolveStrategy", securityModel.ResolveStrategy)
	c.Set("role_id", ri)
	c.Set("user_id", result.UserId.OrDefault(""))
	c.Set("authResult", result)
	c.Set("workspaceId", wi)

	return true
}

func WithAuthorizationFn(securityModel *SecurityModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthorizeRequest(securityModel, c)
	}
}

func CliAuth(security *SecurityModel) (*AuthResultDto, *ferror.Error) {
	context := &AuthContextDto{
		WorkspaceId:  fireback.GetConfig().CliWorkspace,
		Token:        fireback.GetConfig().CliToken,
		Capabilities: []application.PermissionInfo{},
		Security:     security,
	}

	return WithAuthorizationPureDefault(context)
}

func CommonCliQueryDSLBuilderAuthorize(c *cli.Command, security *SecurityModel) QueryDSL {
	q := CommonCliQueryDSLBuilder(c)

	if security != nil && security.ResolveStrategy != ResolveStrategyPublic {
		result, err := CliAuth(security)

		if err != nil {

			if err.ToPublicEndUser(&q).Message != err.ToPublicEndUser(&q).MessageTranslated {
				log.Fatalf("%s", err.ToPublicEndUser(&q).Message)
			}
			log.Default().Printf("%s", err.ToPublicEndUser(&q).MessageTranslated)
		}

		q.ResolveStrategy = security.ResolveStrategy
		q.InternalQuery = result.SqlContext
		if result.UserId.IsSet() && result.UserId.OrDefault("") != "" {
			q.UserId = result.UserId.OrDefault("")
		}
		q.UserAccessPerWorkspace = result.UserAccessPerWorkspace

	}

	return q
}

func CommonCliQueryDSLBuilder(c *cli.Command) QueryDSL {

	queryString := c.String("query")
	startIndex := c.Int("offset")
	var cursor *string = nil
	if c.IsSet("cursor") {
		val := c.String("cursor")
		cursor = &val
	}

	itemsPerPage := c.Int("limit")

	if startIndex < 0 {
		startIndex = 0
	}

	switch {
	case itemsPerPage > 1000:
		itemsPerPage = 1000
	case itemsPerPage <= 0:
		itemsPerPage = 20
	}

	lang := "en"
	region := "US"
	workspaceId := fireback.GetConfig().CliWorkspace

	if fireback.GetConfig().CliLanguage != "" {
		lang = fireback.GetConfig().CliLanguage
	}

	if fireback.GetConfig().CliRegion != "" {
		region = fireback.GetConfig().CliRegion
	}

	withPreloads := c.String("wp")

	var f QueryDSL = QueryDSL{
		Query:        queryString,
		StartIndex:   startIndex,
		C:            c,
		Cursor:       cursor,
		WorkspaceId:  workspaceId,
		Language:     lang,
		Region:       strings.ToUpper(region),
		ItemsPerPage: itemsPerPage,
	}

	if len(withPreloads) > 0 {
		f.WithPreloads = strings.Split(strings.Trim(withPreloads, " "), ",")
	}

	if c.IsSet("lang") {
		f.Language = c.String("lang")
	}

	if c.IsSet("deep") {
		f.Deep = c.Bool("deep")
	}
	if c.IsSet("sort") {
		f.Sort = c.String("sort")
	}

	if c.IsSet("workspaceId") {
		f.WorkspaceId = c.String("workspaceId")
	}

	if c.IsSet("userId") {
		f.UserId = c.String("userId")
	}

	if c.IsSet("id") {
		f.UniqueId = c.String("id")
		fmt.Println(f.UniqueId)
	}

	return f
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
				Message:  abac.AbacMessages.ActionOnlyInRoot,
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

type QueryDSL struct {

	// It's a common string query againt database in text format.
	Query string `json:"query"`

	// Usefull for the paginated queries, it would add the start index
	// and in SQL becomes as offset
	StartIndex int `json:"startIndex"`

	// Numeric cursor
	Cursor *string `json:"cursor"`

	// Useful for paginated queries, similar to limit in SQL queries
	ItemsPerPage int `json:"itemsPerPage"`

	// It's gonna make left joins on the query, if the entity has
	// objects or arrays. It would slow down the query dramatically.
	Deep bool `json:"deep"`

	// It would indicate how the result to be sorted on SQL queries
	Sort string `json:"sort"`

	// Useful when querying against a single element, and by passing
	// uniqueId you can retrieve the record
	UniqueId string `json:"uniqueId"`

	// Sometimes you need to access the raw socket connection, here there is :)
	RawSocketConnection *websocket.Conn

	// It would indicate to the Gorm orm which tables to be included in the
	// SQL search.
	WithPreloads []string `json:"withPreloads"`

	// this is gin context upon the request, which is being attached to the dsl
	// regularly, should not be accessed directly but in reality many times we need
	// to work low level and there is no reason framework do not allow it.
	G *gin.Context `json:"-" yaml:"-"`

	// this is cli context upon the request, which is being attached to the dsl
	// regularly, should not be accessed directly but in reality many times we need
	// to work low level and there is no reason framework do not allow it.
	C *cli.Command `json:"-" yaml:"-"`

	// The gorm transaction object. By setting the query Tx, you can connect
	// few Fireback actions to be done as transaction. Fireback also uses this
	// Object between it's own operations
	Tx *gorm.DB

	// This event will be trigged in the system, if that action is done
	TriggerEventName string `json:"-"`

	// The header authorization, the encrypted token is availble
	// in every request
	Authorization string `json:"authorization"`

	// Automatically assigned UserId to the request after analising the token
	// This will be used to save each entity and determine the owner of the record
	UserId string `json:"-"`

	ResolveStrategy string `json:"-"`

	LinkerId string `json:"-"`

	/// Extra where sql generated in the filter process.
	FilterQuery string `json:"-" yaml:"-"`

	/// JQ Json query generated
	JqQuery string `json:"-" yaml:"-"`

	// This is the person who is requesting, regardless of the workspace
	SearchPhrase string `json:"searchPhrase"`

	// This is the workspace which user is working inside, usually data belongs there
	WorkspaceId string `json:"-"`

	// Those capabilities which user has
	ActionRequires []application.PermissionInfo `json:"-"`

	// This is the capabilities that user has
	UserHas []string `json:"-"`

	UserAccessPerWorkspace *UserAccessPerWorkspaceDto `json:"-" yaml:"-"`

	// This is limitation of that workspace
	WorkspaceHas []string `json:"-"`

	InternalQuery string   `json:"-"`
	Language      string   `json:"-"`
	Region        string   `json:"-"`
	Preloads      []string `json:"-"`
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

func (x QueryDSL) Json() string {
	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))

}

func (x QueryDSL) GetLanguage() string {
	return x.Language
}

// GetWorkspaceId implements owner.Owner.
func (x QueryDSL) GetWorkspaceId() string {
	return x.WorkspaceId
}

// GetUserId implements owner.Owner.
func (x QueryDSL) GetUserId() string {
	return x.UserId
}

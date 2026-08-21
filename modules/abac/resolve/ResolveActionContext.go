package resolve

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

var ROOT_ALL_ACCESS = "root.*"
var ROOT_VAR = "root"

var ResolveStrategyPublic = "public"
var ResolveStrategyUser = "user"
var ResolveStrategyWorkspace = "workspace"

func ResolveActionContext(
	request emigo.EmiRequestContexts,
	securityModel *SecurityModel,
) (*AuthResultDto, error) {
	if securityModel == nil {
		return nil, nil
	}

	if ctx, ok := request.GetGinCtx().(*gin.Context); ok {
		return resolveGinActionContext(ctx, securityModel)
	}

	if _, ok := request.GetCliCtx().(*cli.Command); ok {
		return resolveCLIActionContext(securityModel)
	}

	return nil, nil
}

func resolveGinActionContext(
	ctx *gin.Context,
	securityModel *SecurityModel,
) (*AuthResultDto, error) {
	authContext, isWebsocketUpgrade := CreateContextFromGin(securityModel, ctx)

	result, err := WithAuthorizationPureDefault(authContext)
	if err == nil {
		return result, nil
	}

	if isWebsocketUpgrade {
		// Connection has already been hijacked. The generated websocket
		// handler will report the error through the websocket connection.
		return nil, nil
	}

	t := Translatable{
		Langauge: ctx.GetHeader("accept-language"),
	}
	ctx.AbortWithStatusJSON(
		int(err.HttpCode),
		gin.H{"error": err.ToPublicEndUser(t)},
	)

	return nil, err
}

func resolveCLIActionContext(
	securityModel *SecurityModel,
) (*AuthResultDto, error) {
	if securityModel.ResolveStrategy == ResolveStrategyPublic {
		return nil, nil
	}

	authContext := &AuthContextDto{
		WorkspaceId:  fireback.GetConfig().CliWorkspace,
		Token:        fireback.GetConfig().CliToken,
		Capabilities: []application.PermissionInfo{},
		Security:     securityModel,
	}

	result, err := WithAuthorizationPureDefault(authContext)
	if err == nil {
		return result, nil
	}

	lang := "en"
	if fireback.GetConfig().CliLanguage != "" {
		lang = fireback.GetConfig().CliLanguage
	}
	t := Translatable{Langauge: lang}

	publicError := err.ToPublicEndUser(t)

	if publicError.Message != publicError.MessageTranslated {
		log.Fatalf("%s", publicError.Message)
	}

	log.Printf("%s", publicError.MessageTranslated)

	return nil, nil
}

func WithAuthorizationFn(securityModel *SecurityModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthorizeRequest(securityModel, c)
	}
}

type Translatable struct {
	Langauge string
}

func (x Translatable) GetLanguage() string {

	return x.Langauge
}

func maskToken(token string) string {
	if len(token) <= 6 {
		return token // too short to mask meaningfully
	}
	return token[:2] + "***" + token[len(token)-4:]
}

func GetWorkspaceAndUserAccesses(query MsgContext) ([]string, []string) {

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

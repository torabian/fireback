package resolve

import (
	"encoding/json"
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

type AccessState struct {
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

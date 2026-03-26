package abac

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/abac/migrations"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func AppMenuWriteQueryCteMock(ctx fireback.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := fireback.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := AppMenuActions.CteQuery(f)
		result := fireback.QueryEntitySuccessResult(f, items, count)
		fireback.WriteMockDataToFile(lang, "", "AppMenu", result)
	}
}

func workspaceModuleCore(module *fireback.ModuleProvider) {

	module.ProvidePermissionHandler(
		ALL_WORKSPACE_CONFIG_PERMISSIONS,
		ALL_WORKSPACE_TYPE_PERMISSIONS,
		ALL_EMAIL_SENDER_PERMISSIONS,
		ALL_EMAIL_PROVIDER_PERMISSIONS,
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
		ALL_GSM_PROVIDER_PERMISSIONS,
		ALL_WORKSPACE_INVITE_PERMISSIONS,
		ALL_BACKUP_TABLE_META_PERMISSIONS,
		ALL_TABLE_VIEW_SIZING_PERMISSIONS,
		ALL_APP_MENU_PERMISSIONS,
		ALL_REGIONAL_CONTENT_PERMISSIONS,
		ALL_USER_WORKSPACE_PERMISSIONS,
		ALL_USER_PERMISSIONS,
		ALL_ROLE_PERMISSIONS,
		ALL_WORKSPACE_ROLE_PERMISSIONS,
		ALL_WORKSPACE_PERMISSIONS,
		ALL_PERM_ABAC_MODULE,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		items := []interface{}{
			&UserEntity{},
			&TokenEntity{},
			&PreferenceEntity{},
			&RoleEntity{},
			&WorkspaceEntity{},
			&WorkspaceInviteEntity{},
			&WorkspaceConfigEntity{},
			&WorkspaceTypeEntity{},
			&WorkspaceTypeEntityPolyglot{},
			&GsmProviderEntity{},
			&BackupTableMetaEntity{},
			&WorkspaceRoleEntity{},
			&UserWorkspaceEntity{},
			&RegionalContentEntity{},
			&TableViewSizingEntity{},
			&AppMenuEntity{},
			&AppMenuEntityPolyglot{},
			&TimezoneGroupEntity{},
			&TimezoneGroupEntityPolyglot{},
			&TimezoneGroupUtcItems{},
		}

		items2 := []interface{}{}
		items2 = append(items2, items...)

		for _, item := range items2 {

			if err := dbref.Debug().AutoMigrate(item); err != nil {
				fmt.Println("Migrating entity issue:", fireback.GetInterfaceName(item))
				return err
			}
		}

		// This is an important function, to create the root workspace.
		// root workspaces is the only, main workspace, which has every other workspace under it.
		return RepairTheWorkspaces()
	})

}

type MicroserviceSetupConfig struct {
	AuthorizationResolver WithAuthorizationPureImpl
}

// Inject this into any project as a complete solution
func AbacCompleteModules() []*fireback.ModuleProvider {
	return []*fireback.ModuleProvider{
		WorkspaceModuleSetup(),
		DriveModuleSetup(),
		NotificationModuleSetup(),
		PassportsModuleSetup(),
	}
}

func WorkspaceModuleSetup() *fireback.ModuleProvider {

	// Default Fireback authorization. You can Override this on microservices
	fireback.WithAuthorizationPure = WithAuthorizationPureDefault
	fireback.WithAuthorizationFn = WithAuthorizationFn
	fireback.WithSocketAuthorization = WithSocketAuthorization

	module := &fireback.ModuleProvider{
		Name:               "abac",
		Definitions:        &Module3Definitions,
		OnEnvInit:          OnInitEnvHook,
		GoMigrateDirectory: &migrations.MigrationsFs,
	}

	workspaceModuleCore(module)

	module.ProvideMockWriterHandler(func(languages []string) {
		// WorkspaceTypeWriteQueryMock(MockQueryContext{Languages: languages})
		// GsmProviderWriteQueryMock(MockQueryContext{Languages: languages})
		// AppMenuWriteQueryCteMock(MockQueryContext{Languages: languages})
	})

	module.ProvideTests(fireback.UserTests,
		[]fireback.Test{
			fireback.TestNewModuleProjectGen,
		},
		AppMenuTests,
		fireback.IntelisenseTest,
	)

	module.ProvideSeederImportHandler(func() {
		// We do not use syncing here.
		// Because fireback is being imported by other modules,
		// they might want their own unique menu items
		// sync items in the fireback/main or desktop one manually for this project.
		// for other projects extending fireback you can use here.
		TimezoneGroupSyncSeeders()
	})

	module.ProvideMockImportHandler(func() {
		// GsmProviderImportMocks()
	})

	module.Actions = [][]fireback.Module3Action{
		GetUserModule3Actions(),
		GetWorkspaceModule3Actions(),
		GetRoleModule3Actions(),
		GetWorkspaceTypeModule3Actions(),
		GetGsmProviderModule3Actions(),
		GetWorkspaceInviteModule3Actions(),
		GetBackupTableMetaModule3Actions(),
		GetTableViewSizingModule3Actions(),
		GetAppMenuModule3Actions(),
		GetEmailConfirmationModule3Actions(),
		AbacCustomActions(),
		GetUserWorkspaceModule3Actions(),
		GetWorkspaceRoleModule3Actions(),
		GetTimezoneGroupModule3Actions(),
		GetWorkspaceConfigModule3Actions(),
		GetRegionalContentModule3Actions(),
		{
			AS_FIREBACK_ACTION,
		},
	}

	module.ProvideCliHandlers([]cli.Command{
		RoleCliFn(),
		UserCliFn(),
		WorkspaceCliFn(),
		MiscCli,
		AS_FIREBACK_ACTION.ToCli(),
	})

	return module
}

func WrapData(v any) any {
	return map[string]any{
		"data": map[string]any{
			"item": v,
		},
	}
}

// func CreateGinCommand() func(g *gin.RouterGroup, x *fireback.FirebackApp) error {
// 	return func(g *gin.RouterGroup, x *fireback.FirebackApp) error {
// 		method, url, h := CheckPassportMethods2ActionHandler(CheckPassportMethods2ActionImpl)
// 		g.Handle(method, url, h)
// 		return nil
// 	}
// }
// var CheckPassportMethods2ActionImpl = func(c CheckPassportMethods2ActionRequest) (*CheckPassportMethods2ActionResponse, error) {
// 	var query fireback.QueryDSL
// 	if c.GinCtx == nil {
// 		query = fireback.CommonCliQueryDSLBuilderAuthorize(c.CliCtx, CheckPassportMethodsSecurityModel)
// 	} else {
// 		query = fireback.ExtractQueryDslFromGinContext(c.GinCtx)
// 	}

// 	return CheckPassportMethods2Impl(c, query)
// }

// func CreateCliCommand() cli.Command {
// 	return cli.Command{
// 		Name:  "check-passport-methods2",
// 		Usage: `Publicly available information to create the authentication form, and show users how they can signin or signup to the system. Based on the PassportMethod entities, it will compute the available methods for the user, considering their region (IP for example)`,
// 		Action: func(c *cli.Context) {
// 			result, err := CheckPassportMethods2ActionImpl(CheckPassportMethods2ActionRequest{CliCtx: c})
// 			fireback.HandleActionInCli2(c, result, err, map[string]map[string]string{})
// 		},
// 	}
// }

// Developer per action starts from here.
var CheckPassportMethods2Impl = func(c CheckPassportMethods2ActionRequest, query fireback.QueryDSL) (*CheckPassportMethods2ActionResponse, error) {

	resp, err := CheckPassportMethodsAction(query)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &CheckPassportMethods2ActionResponse{
		Payload: WrapData(resp),
	}, nil
}

/**
*	Each result from an action, either can directly access to Gin or Cli
* Context and handle things over there, or can return an EmiAction Result
** Which is standard for a quick result.
**/
type EmiActionResult interface {
	GetStatusCode() int
	GetRespHeaders() map[string]string
	GetPayload() interface{}
}

func WriteActionResponseToGin(m *gin.Context, resp EmiActionResult, err error) {
	if err != nil {
		m.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If the handler returned nil (and no error), it means the response was handled manually.
	if resp == nil {
		return
	}

	// Apply headers
	for k, v := range resp.GetRespHeaders() {
		m.Header(k, v)
	}

	// Apply status and payload
	status := resp.GetStatusCode()
	if status == 0 {
		status = http.StatusOK
	}
	if resp.GetPayload() != nil {
		m.JSON(status, resp.GetPayload())
	} else {
		m.Status(status)
	}
}

var AS_FIREBACK_ACTION = fireback.Module3Action{
	Method: CheckPassportMethods2ActionMeta().Method,
	Url:    CheckPassportMethods2ActionMeta().URL,
	Handlers: []gin.HandlerFunc{
		func(m *gin.Context) {
			req := CheckPassportMethods2ActionRequest{
				QueryParams: m.Request.URL.Query(),
				Headers:     m.Request.Header,
				GinCtx:      m,
			}

			var query fireback.QueryDSL
			query = fireback.ExtractQueryDslFromGinContext(m)
			resp, err := CheckPassportMethods2Impl(req, query)
			WriteActionResponseToGin(m, resp, err)
		},
	},
	CliAction: func(c *cli.Context, security *fireback.SecurityModel) error {
		query := fireback.CommonCliQueryDSLBuilderAuthorize(c, CheckPassportMethodsSecurityModel)
		req := CheckPassportMethods2ActionRequest{
			CliCtx: c,
			// Can be casted from --qp-x-a, --h-aa
			// QueryParams: ,
			// Headers: ,
		}

		resp, err := CheckPassportMethods2Impl(req, query)
		fireback.HandleActionInCli2(c, resp.Payload, err, map[string]map[string]string{})

		return nil
	},
	CliName:     CheckPassportMethods2ActionMeta().CliName,
	Name:        CheckPassportMethods2ActionMeta().Name,
	Description: "new",
}

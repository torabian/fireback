package abac

import (
	"embed"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/abac/migrations"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
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

type MicroserviceSetupConfig struct {
	AuthorizationResolver WithAuthorizationPureImpl
}

// Inject this into any project as a complete solution
func AbacCompleteModules() []*fireback.ModuleProvider {
	return []*fireback.ModuleProvider{
		WorkspaceModuleSetup(),
		NotificationModuleSetup(),
		PassportsModuleSetup(),
	}
}

func WorkspaceModuleSetup() *fireback.ModuleProvider {

	// Default Fireback authorization. You can Override this on microservices
	fireback.WithAuthorizationPure = WithAuthorizationPureDefault
	fireback.WithAuthorizationFn = WithAuthorizationFn
	fireback.AuthorizeRequest = AuthorizeRequest
	fireback.WithSocketAuthorization = WithSocketAuthorization

	module := &fireback.ModuleProvider{
		Name:               "abac",
		Definitions:        &Module3Definitions,
		OnEnvInit:          OnInitEnvHook,
		GoMigrateDirectory: &migrations.MigrationsFs,

		// Actions declared in Abac.emi.yml (moved out of AbacModule3.yml's old actions:
		// section) are wired directly here, the same way FirebackModuleSetup wires the
		// Capability* actions - rather than through the legacy Module3Action/Impl glue.
		GinWebServerInitHooks: []func(g *gin.RouterGroup, x *fireback.FirebackApp) error{
			func(g *gin.RouterGroup, x *fireback.FirebackApp) error {
				QueryUserRoleWorkspacesActionGin(g, QueryUserRoleWorkspacesAction)
				InviteToWorkspaceActionGin(g, InviteToWorkspaceAction)
				UserInvitationsActionGin(g, UserInvitationsAction)
				SignoutActionGin(g, SignoutAction)
				OauthAuthenticateActionGin(g, OauthAuthenticateAction)
				AcceptInviteActionGin(g, AcceptInviteAction)
				ConfirmClassicPassportTotpActionGin(g, ConfirmClassicPassportTotpAction)
				ChangePasswordActionGin(g, ChangePasswordAction)
				UserPassportsActionGin(g, UserPassportsAction)
				CreateWorkspaceActionGin(g, CreateWorkspaceAction)
				ClassicPassportRequestOtpActionGin(g, ClassicPassportRequestOtpAction)
				ClassicPassportOtpActionGin(g, ClassicPassportOtpAction)
				CheckClassicPassportActionGin(g, CheckClassicPassportAction)
				ClassicSignupActionGin(g, ClassicSignupAction)
				ClassicSigninActionGin(g, ClassicSigninAction)
				QueryWorkspaceTypesPubliclyActionGin(g, QueryWorkspaceTypesPubliclyAction)
				CheckPassportMethodsActionGin(g, CheckPassportMethodsAction)
				OsLoginAuthenticateActionGin(g, OsLoginAuthenticateAction)
				SendEmailActionGin(g, SendEmailAction)
				SendEmailWithProviderActionGin(g, SendEmailWithProviderAction)
				GsmSendSmsActionGin(g, GsmSendSmsAction)
				GsmSendSmsWithProviderActionGin(g, GsmSendSmsWithProviderAction)

				return nil
			},
		},
	}

	module.ProvidePermissionHandler(
		ALL_WORKSPACE_CONFIG_PERMISSIONS,
		ALL_WORKSPACE_TYPE_PERMISSIONS,
		ALL_EMAIL_SENDER_PERMISSIONS,
		ALL_EMAIL_PROVIDER_PERMISSIONS,
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
		ALL_GSM_PROVIDER_PERMISSIONS,
		ALL_WORKSPACE_INVITE_PERMISSIONS,
		ALL_TABLE_VIEW_SIZING_PERMISSIONS,
		ALL_APP_MENU_PERMISSIONS,
		ALL_REGIONAL_CONTENT_PERMISSIONS,
		ALL_USER_WORKSPACE_PERMISSIONS,
		ALL_USER_PERMISSIONS,
		ALL_ROLE_PERMISSIONS,
		ALL_WORKSPACE_ROLE_PERMISSIONS,
		ALL_WORKSPACE_PERMISSIONS,
		ALL_PERM_ABAC_MODULE,
		ALL_TIMEZONE_GROUP_PERMISSIONS,
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
			&WorkspaceRoleEntity{},
			&UserWorkspaceEntity{},
			&RegionalContentEntity{},
			&TableViewSizingEntity{},
			&AppMenuEntity{},
			&AppMenuEntityPolyglot{},
			&TimezoneGroupEntity{},
			&TimezoneGroupEntityPolyglot{},
		}

		items2 := []interface{}{}
		items2 = append(items2, items...)

		for _, item := range items2 {

			if err := dbref.AutoMigrate(item); err != nil {
				fmt.Println("Migrating entity issue:", fireback.GetInterfaceName(item))
				return err
			}
		}

		// This is an important function, to create the root workspace.
		// root workspaces is the only, main workspace, which has every other workspace under it.
		return RepairTheWorkspaces()
	})

	module.ProvideMockWriterHandler(func(languages []string) {
		// WorkspaceTypeWriteQueryMock(MockQueryContext{Languages: languages})
		// GsmProviderWriteQueryMock(MockQueryContext{Languages: languages})
		// AppMenuWriteQueryCteMock(MockQueryContext{Languages: languages})
	})

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
		GetTableViewSizingModule3Actions(),
		GetAppMenuModule3Actions(),
		GetEmailConfirmationModule3Actions(),
		GetUserWorkspaceModule3Actions(),
		GetWorkspaceRoleModule3Actions(),
		GetTimezoneGroupModule3Actions(),
		GetWorkspaceConfigModule3Actions(),
		GetRegionalContentModule3Actions(),
		// {
		// 	AS_FIREBACK_ACTION,
		// },
	}

	module.ProvideCliHandlers([]*cli.Command{
		RoleCliFn(),
		UserCliFn(),
		WorkspaceCliFn(),
		&MiscCli,
		TimezoneGroupCliFn(),
		// The actions below moved to Abac.emi.yml and don't have a home yet among the
		// entity-scoped cli groups above (AcceptInvite/UserInvitations live under UserCliFn,
		// QueryUserRoleWorkspaces/QueryWorkspaceTypesPublicly under WorkspaceCliFn,
		// OsLoginAuthenticate/CheckPassportMethods/UserPassports/OauthAuthenticate under
		// PassportCliFn - see UserEntity.go/WorkspaceCli.go/PassportCli.go).
		SignoutActionCliHandler(SignoutAction),
		InviteToWorkspaceActionCliHandler(InviteToWorkspaceAction),
		ConfirmClassicPassportTotpActionCliHandler(ConfirmClassicPassportTotpAction),
		ChangePasswordActionCliHandler(ChangePasswordAction),
		CreateWorkspaceActionCliHandler(CreateWorkspaceAction),
		ClassicPassportRequestOtpActionCliHandler(ClassicPassportRequestOtpAction),
		ClassicPassportOtpActionCliHandler(ClassicPassportOtpAction),
		CheckClassicPassportActionCliHandler(CheckClassicPassportAction),
		ClassicSignupActionCliHandler(ClassicSignupAction),
		ClassicSigninActionCliHandler(ClassicSigninAction),
		SendEmailActionCliHandler(SendEmailAction),
		SendEmailWithProviderActionCliHandler(SendEmailWithProviderAction),
		GsmSendSmsActionCliHandler(GsmSendSmsAction),
		GsmSendSmsWithProviderActionCliHandler(GsmSendSmsWithProviderAction),
	})

	module.ProvideCliHandlers([]*cli.Command{&AuthFlow, &AbacActions})

	return module
}

var AbacActions cli.Command = cli.Command{
	Name:  "abac",
	Usage: "All actions which are available for abac module",
	Commands: append(
		[]*cli.Command{
			{
				Name:  "internal",
				Usage: "Internal entities which are used for processes. Manipulating these requires deep internal knowledge",
				Commands: []*cli.Command{
					PublicJoinKeyCliFn(),
					PublicAuthenticationCliFn(),
				},
			},
		},
		GetAbacActionsCli()...,
	),
}

package abac

import (
	"embed"
	"fmt"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func AppMenuWriteQueryCteMock(ctx workspaces.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := workspaces.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := AppMenuActions.CteQuery(f)
		result := workspaces.QueryEntitySuccessResult(f, items, count)
		workspaces.WriteMockDataToFile(lang, "", "AppMenu", result)
	}
}

var EverRunEntities []interface{} = []interface{}{
	&CapabilityEntity{},
	&CapabilityEntityPolyglot{},
}

func workspaceModuleCore(module *workspaces.ModuleProvider) {

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
		ALL_CAPABILITY_PERMISSIONS,
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
		items2 = append(items2, EverRunEntities...)
		items2 = append(items2, items...)

		for _, item := range items2 {
			if err := dbref.AutoMigrate(item); err != nil {
				fmt.Println("Migrating entity issue:", workspaces.GetInterfaceName(item))
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

// This authorization would by pass the security requirement and by default is used in microservice
// setup. Basically, everything is allowed. You can make a authorization with jwt or other methods
// you want, even sending a http request to thirdparty.
// context contains the token, workspace-id in header, and you need to decide if the action is allowed
// Maybe it's gonna be a good practise to provide headers here as well, in case only token is not enough
// to decide but for now it's only context and workspaceId
func MicroserviceWithoutContextResolver(context *AuthContextDto) (*AuthResultDto, *workspaces.IError) {
	return &AuthResultDto{}, nil
}

func FirebackMicroService(config *MicroserviceSetupConfig) *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name: "workspaces",

		EntityProvider: func(d *gorm.DB) error {

			for _, item := range EverRunEntities {
				if err := workspaces.GetDbRef().AutoMigrate(item); err != nil {
					fmt.Println("Migrating entity issue:", workspaces.GetInterfaceName(item))
					return err
				}
			}

			return nil
		},
	}

	module.ProvidePermissionHandler(
		ALL_CAPABILITY_PERMISSIONS,
	)

	module.ProvideCliHandlers([]cli.Command{
		CapabilityCliFn(),
	})

	// This is a little bit funcy, but no other option at the moment, we change the function
	// globally.
	if config != nil && config.AuthorizationResolver != nil {
		WithAuthorizationPure = config.AuthorizationResolver
	} else {
		WithAuthorizationPure = MicroserviceWithoutContextResolver
	}

	return module
}

func WorkspaceModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "workspaces",
		Definitions: &Module3Definitions,
	}

	workspaceModuleCore(module)

	module.ProvideMockWriterHandler(func(languages []string) {
		// WorkspaceTypeWriteQueryMock(MockQueryContext{Languages: languages})
		// GsmProviderWriteQueryMock(MockQueryContext{Languages: languages})
		// AppMenuWriteQueryCteMock(MockQueryContext{Languages: languages})
	})

	module.ProvideTests(workspaces.UserTests,
		[]workspaces.Test{
			workspaces.TestNewModuleProjectGen,
		},
		AppMenuTests,
		workspaces.IntelisenseTest,
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

	module.Actions = [][]workspaces.Module3Action{
		GetUserModule3Actions(),
		GetWorkspaceModule3Actions(),
		GetRoleModule3Actions(),
		GetCapabilityModule3Actions(),
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
	}

	module.ProvideCliHandlers([]cli.Command{
		CapabilityCliFn(),
		RoleCliFn(),
		UserCliFn(),
		WorkspaceCliFn(),
		MiscCli,
	})

	return module
}

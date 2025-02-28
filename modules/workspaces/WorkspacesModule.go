package workspaces

import (
	"embed"
	"fmt"
	"log"

	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func UpsertPermission(permInfo *PermissionInfo, hasChildren bool, db *gorm.DB) {
	var entity *CapabilityEntity = nil
	perm := permInfo.CompleteKey

	if hasChildren {
		perm = perm + ".*"
	}

	system := "system"

	if (db.Where(CapabilityEntity{UniqueId: perm}).First(&entity).Error != nil) {
		err := db.Create(&CapabilityEntity{
			UniqueId:    perm,
			WorkspaceId: &system,
			Description: &permInfo.Description,
			Name:        &permInfo.Name,
		}).Error

		if err != nil {
			log.Fatalln("Cannot start the app because a permission creation failed.", perm, err)
		}
	}
}

func AppMenuWriteQueryCteMock(ctx MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := AppMenuActionCteQuery(f)
		result := QueryEntitySuccessResult(f, items, count)
		WriteMockDataToFile(lang, "", "AppMenu", result)
	}
}
func workspaceModuleCore(module *ModuleProvider) {

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
		ALL_PERM_WORKSPACES_MODULE,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		items := []interface{}{
			&CapabilityEntity{},
			&CapabilityEntityPolyglot{},
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

		for _, item := range items {
			if err := dbref.AutoMigrate(item); err != nil {
				fmt.Println("Migrating entity issue:", GetInterfaceName(item))
				return err
			}
		}

		// This is an important function, to create the root workspace.
		// root workspaces is the only, main workspace, which has every other workspace under it.
		return RepairTheWorkspaces()
	})

}

func WorkspaceModuleMicroServiceSetup() *ModuleProvider {
	module := &ModuleProvider{
		Name:        "workspaces",
		Definitions: &Module3Definitions,
	}

	workspaceModuleCore(module)

	return module
}

func WorkspaceModuleSetup() *ModuleProvider {
	module := &ModuleProvider{
		Name:        "workspaces",
		Definitions: &Module3Definitions,
	}

	workspaceModuleCore(module)

	module.ProvideMockWriterHandler(func(languages []string) {
		// WorkspaceTypeWriteQueryMock(MockQueryContext{Languages: languages})
		// GsmProviderWriteQueryMock(MockQueryContext{Languages: languages})
		// AppMenuWriteQueryCteMock(MockQueryContext{Languages: languages})
	})

	module.ProvideTests(UserTests,
		[]Test{
			TestNewModuleProjectGen,
		},
		WorkspaceCreationTests,
		AppMenuTests,
		IntelisenseTest,
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

	module.Actions = [][]Module3Action{
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
		WorkspacesCustomActions(),
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

package workspaces

import (
	"embed"
	"log"

	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func UpsertPermission(perm string, hasChildren bool, db *gorm.DB) {
	var entity *CapabilityEntity = nil

	if hasChildren {
		perm = perm + "/*"
	}

	system := "system"

	if (db.Where(CapabilityEntity{UniqueId: perm}).First(&entity).Error != nil) {
		err := db.Create(&CapabilityEntity{UniqueId: perm, WorkspaceId: &system}).Error

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

func WorkspaceModuleSetup() *ModuleProvider {
	module := &ModuleProvider{
		Name:        "workspaces",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockWriterHandler(func(languages []string) {
		WorkspaceTypeWriteQueryMock(MockQueryContext{Languages: languages})
		GsmProviderWriteQueryMock(MockQueryContext{Languages: languages})
		AppMenuWriteQueryCteMock(MockQueryContext{Languages: languages})
	})

	module.ProvideSeederImportHandler(func() {
		RegionalContentSyncSeeders()
	})

	module.ProvideMockImportHandler(func() {
		GsmProviderImportMocks()
	})

	module.ProvidePermissionHandler(
		ALL_WORKSPACES_PERMISSIONS,
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
		ALL_WORKSPACE_ROLE_PERMISSIONS,
	)
	module.ProvideTranslationList(WorkspacesTranslations)

	module.Actions = [][]Module2Action{
		GetUserModule2Actions(),
		GetWorkspaceModule2Actions(),
		GetRoleModule2Actions(),
		GetCapabilityModule2Actions(),
		GetWorkspaceTypeModule2Actions(),
		GetGsmProviderModule2Actions(),
		GetWorkspaceInviteModule2Actions(),
		GetBackupTableMetaModule2Actions(),
		GetTableViewSizingModule2Actions(),
		GetAppMenuModule2Actions(),
		GetEmailConfirmationModule2Actions(),
		WorkspacesCustomActions(),
		GetUserWorkspaceModule2Actions(),
		GetWorkspaceRoleModule2Actions(),
		GetRegionalContentModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		dbref.AutoMigrate(&CapabilityEntity{})
		dbref.AutoMigrate(&CapabilityEntityPolyglot{})
		dbref.AutoMigrate(&UserEntity{})
		dbref.AutoMigrate(&TokenEntity{})
		dbref.AutoMigrate(&PreferenceEntity{})
		dbref.AutoMigrate(&RoleEntity{})
		dbref.AutoMigrate(&WorkspaceEntity{})
		dbref.AutoMigrate(&WorkspaceInviteEntity{})
		dbref.AutoMigrate(&WorkspaceConfigEntity{})
		dbref.AutoMigrate(&WorkspaceTypeEntity{})
		dbref.AutoMigrate(&WorkspaceTypeEntityPolyglot{})
		dbref.AutoMigrate(&GsmProviderEntity{})
		dbref.AutoMigrate(&BackupTableMetaEntity{})
		dbref.AutoMigrate(&WorkspaceRoleEntity{})
		dbref.AutoMigrate(&UserWorkspaceEntity{})
		dbref.AutoMigrate(&RegionalContentEntity{})

		dbref.AutoMigrate(&TableViewSizingEntity{})
		dbref.AutoMigrate(&AppMenuEntity{}, &AppMenuEntityPolyglot{})

		// This is an important function, to create the root workspace.
		// root workspaces is the only, main workspace, which has every other workspace under it.
		if err := RepairTheWorkspaces(); err == nil {
			// fmt.Println("✓ Root role seems to be healthy")
		}

	})

	module.ProvideCliHandlers([]cli.Command{
		CapabilityCliFn(),
		RoleCliFn(),
		UserCliFn(),
		WorkspaceCliFn(),
		WorkspaceInviteCliFn(),
		BackupTableMetaCliFn(),
		TableViewSizingCliFn(),
		AppMenuCliFn(),
		RegionalContentCliFn(),
	})

	return module
}

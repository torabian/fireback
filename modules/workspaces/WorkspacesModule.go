package workspaces

import (
	"embed"
	"log"

	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func UpsertPermission(permInfo *PermissionInfo, hasChildren bool, db *gorm.DB) {
	var entity *CapabilityEntity = nil
	perm := permInfo.CompleteKey

	if hasChildren {
		perm = perm + "/*"
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

func WorkspaceModuleSetup() *ModuleProvider {
	module := &ModuleProvider{
		Name:        "workspaces",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockWriterHandler(func(languages []string) {
		// WorkspaceTypeWriteQueryMock(MockQueryContext{Languages: languages})
		// GsmProviderWriteQueryMock(MockQueryContext{Languages: languages})
		// AppMenuWriteQueryCteMock(MockQueryContext{Languages: languages})
	})

	module.ProvideTests(UserTests, []Test{TestNewModuleProjectGen})

	module.ProvideSeederImportHandler(func() {
		// We do not use syncing here.
		// Because fireback is being imported by other modules,
		// they might want their own unique menu items
		// sync items in the fireback/main or desktop one manually for this project.
		// for other projects extending fireback you can use here.
	})

	module.ProvideMockImportHandler(func() {
		// GsmProviderImportMocks()
	})

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
		ALL_WORKSPACE_ROLE_PERMISSIONS,
		ALL_PERM_WORKSPACES_MODULE,
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

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		if err := dbref.AutoMigrate(
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
		); err != nil {
			return err
		}

		// This is an important function, to create the root workspace.
		// root workspaces is the only, main workspace, which has every other workspace under it.
		return RepairTheWorkspaces()
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

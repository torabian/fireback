package main

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/cmd/fireback-server/seeders"
	fbseeders "github.com/torabian/fireback/cmd/fireback-server/seeders"
	"github.com/torabian/fireback/modules/commonprofile"
	"github.com/torabian/fireback/modules/drive"
	"github.com/torabian/fireback/modules/keyboardActions"
	"github.com/torabian/fireback/modules/widget"
	"github.com/torabian/fireback/modules/workspaces"
)

func baseService() *workspaces.XWebServer {

	xapp := &workspaces.XWebServer{
		Title:              "Fireback desktop boilerplate",
		SupportedLanguages: []string{"fa", "en"},
		RunTus: func() {
			// drive.LiftTusServer()
		},
		SearchProviders: []workspaces.SearchProviderFn{
			workspaces.QueryMenusReact,
			workspaces.QueryRolesReact,
		},
		RunSearch: workspaces.InjectReactiveSearch,
		RunSocket: func(e *gin.Engine) {
			workspaces.HandleSocket(e)
		},
		PublicFolders: []workspaces.PublicFolderInfo{},
		SeedersSync: func() {
			workspaces.AppMenuSyncSeederFromFs(&seeders.ViewsFs, []string{})
			workspaces.AppMenuSyncSeederFromFs(&fbseeders.ViewsFs, []string{"personal-menu.yml", "fireback-menu-common.yml"})
		},
	}

	xapp.Modules = []*workspaces.ModuleProvider{
		// Important to setup the workspaces at first, so the capabilties module is there
		workspaces.WorkspaceModuleSetup(),
		keyboardActions.KeyboardActionsModuleSetup(),
		drive.DriveModuleSetup(),
		workspaces.NotificationModuleSetup(),
		workspaces.PassportsModuleSetup(),
		commonprofile.CommonProfileModuleSetup(),
		widget.WidgetModuleSetup(),
	}

	db, err := workspaces.CreateDatabasePool()

	if db != nil && err == nil {
		workspaces.SyncDatabase(xapp, db)
		workspaces.SyncPermissionsInDatabase(xapp, db)
		workspaces.ExecuteSeederImport(xapp)
	}

	return xapp
}

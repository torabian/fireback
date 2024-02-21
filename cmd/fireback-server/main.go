package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"pixelplux.com/fireback/modules/commonprofile"
	"pixelplux.com/fireback/modules/currency"
	"pixelplux.com/fireback/modules/drive"
	"pixelplux.com/fireback/modules/geo"
	"pixelplux.com/fireback/modules/keyboardActions"
	"pixelplux.com/fireback/modules/licenses"
	"pixelplux.com/fireback/modules/widget"
	"pixelplux.com/fireback/modules/workspaces"
	"pixelplux.com/fireback/modules/worldtimezone"
)

//REMOVEMEgo:embed all:ui
// var ui embed.FS

var PRODUCT_NAMESPACENAME = "fireback"
var PRODUCT_DESCRIPTION = "Fireback core microservice"
var PRODUCT_LANGUAGES = []string{"fa", "en"}

var xapp = &workspaces.XWebServer{
	Title:              PRODUCT_DESCRIPTION,
	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []workspaces.SearchProviderFn{
		workspaces.QueryMenusReact,
		workspaces.QueryRolesReact,
	},
	RunTus: func() {
		drive.LiftTusServer()
	},
	RunSocket: func(e *gin.Engine) {
		workspaces.HandleSocket(e)
	},
	RunSearch:     workspaces.InjectReactiveSearch,
	PublicFolders: []workspaces.PublicFolderInfo{
		// {Fs: &ui, Folder: "ui"},
	},
	Modules: []*workspaces.ModuleProvider{
		// Important to setup the workspaces at first, so the capabilties module is there
		workspaces.WorkspaceModuleSetup(),
		geo.GeoModuleSetup(),
		keyboardActions.KeyboardActionsModuleSetup(),
		drive.DriveModuleSetup(),
		workspaces.NotificationModuleSetup(),
		workspaces.PassportsModuleSetup(),
		widget.WidgetModuleSetup(),
		commonprofile.CommonProfileModuleSetup(),
		currency.CurrencyModuleSetup(),
		licenses.LicensesModuleSetup(),
		worldtimezone.LicensesModuleSetup(),
	},
}

func main() {

	os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)

	db, dbErr := workspaces.CreateDatabasePool()

	if db != nil && dbErr == nil {
		workspaces.SyncDatabase(xapp, db)
		workspaces.SyncPermissionsInDatabase(xapp, db)
	} else {
		log.Fatalln("Database error", dbErr)
	}

	xapp.CliActions = func() []cli.Command {
		return workspaces.GetCommonWebServerCliActions(xapp)
	}

	workspaces.RunApp(xapp)
}

package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/cms"
	"github.com/torabian/fireback/modules/commonprofile"
	"github.com/torabian/fireback/modules/currency"
	"github.com/torabian/fireback/modules/geo"
	"github.com/torabian/fireback/modules/keyboardActions"
	"github.com/torabian/fireback/modules/licenses"
	"github.com/torabian/fireback/modules/shop"
	"github.com/torabian/fireback/modules/widget"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/torabian/fireback/modules/worldtimezone"
)

var PRODUCT_NAMESPACENAME = "{{ .ctx.Name }}"
var PRODUCT_DESCRIPTION = "{{ .ctx.Description }}"
var PRODUCT_LANGUAGES = []string{"en"}

var xapp = &workspaces.XWebServer{
	Title: PRODUCT_DESCRIPTION,

	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []workspaces.SearchProviderFn{
		workspaces.QueryMenusReact,
		workspaces.QueryRolesReact,
	},
	RunTus: func() {
		workspaces.LiftTusServer()
	},
	RunSocket: func(e *gin.Engine) {
		workspaces.HandleSocket(e)
	},
	RunSearch:     workspaces.InjectReactiveSearch,
	PublicFolders: []workspaces.PublicFolderInfo{
		// You can set a series of static folders to be served along with fireback.
		// This is only for static content. For advanced MVX render templates, you need to
		// Bootstrap those themes
		// Add these two lines on the top of the file
		/////go:embed all:ui
		// var ui embed.FS
		// and then uncomment this, for example to serve static react or angular content
		// {Fs: &ui, Folder: "ui"},
	},
	SetupWebServerHook: func(e *gin.Engine, xs *workspaces.XWebServer) {
		// You can uncomment the sample theme for shop here
		// zayshop.Bootstrap(e)

	},
	Modules: []*workspaces.ModuleProvider{
		// Important to setup the workspaces at first, so the capabilties module is there
		workspaces.WorkspaceModuleSetup(),
		workspaces.DriveModuleSetup(),
		workspaces.NotificationModuleSetup(),
		workspaces.PassportsModuleSetup(),

		// These are optional packages that might be used or might not be needed.
		geo.GeoModuleSetup(),
		keyboardActions.KeyboardActionsModuleSetup(),
		widget.WidgetModuleSetup(),
		commonprofile.CommonProfileModuleSetup(),
		cms.CmsModuleSetup(),
		currency.CurrencyModuleSetup(),
		licenses.LicensesModuleSetup(),
		shop.ShopModuleSetup(),
		worldtimezone.LicensesModuleSetup(),
	},
}

func main() {

	// This is an important setting for some kind of app which will be installed
	// it makes it easier for fireback to find the configuration.
	os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)

	// This AppStart function is a wrapper for few things commonly can handle entire backend project
	// startup. For mobile or desktop might other functionality be used.
	xapp.CommonHeadlessAppStart(func() {
		// If anything needs to be done after database initialized
	})
}

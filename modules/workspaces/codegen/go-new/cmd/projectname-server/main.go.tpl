package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"{{ .ctx.ModuleName }}/cmd/{{ .ctx.Name }}-server/menu"

	{{ if .ctx.IsMonolith }}
	"github.com/torabian/fireback/modules/commonprofile"
	"github.com/torabian/fireback/modules/currency"
	"github.com/torabian/fireback/modules/geo"
	"github.com/torabian/fireback/modules/accessibility"
	"github.com/torabian/fireback/modules/licenses"
	"github.com/torabian/fireback/modules/widget"
	"github.com/torabian/fireback/modules/worldtimezone"
	{{ end }}
	"github.com/torabian/fireback/modules/workspaces"
)

var PRODUCT_NAMESPACENAME = "{{ .ctx.Name }}"
var PRODUCT_DESCRIPTION = "{{ .ctx.Description }}"
var PRODUCT_LANGUAGES = []string{"en"}

var xapp = &workspaces.FirebackApp{
	Title: PRODUCT_DESCRIPTION,

	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []workspaces.SearchProviderFn{
		{{ if .ctx.IsMonolith }}
		workspaces.QueryMenusReact,
		workspaces.QueryRolesReact,
		{{ end }}
	},
	SeedersSync: func() {
		// Sample menu item to make it easier for demos
		workspaces.AppMenuSyncSeederFromFs(&menu.Menu, []string{"new-menu.yml"})
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
	SetupWebServerHook: func(e *gin.Engine, xs *workspaces.FirebackApp) {

	},
	Modules: []*workspaces.ModuleProvider{
		{{ if .ctx.IsMonolith }}
		// Important to setup the workspaces at first, so the capabilties module is there
		workspaces.WorkspaceModuleSetup(),
		workspaces.DriveModuleSetup(),
		workspaces.NotificationModuleSetup(),
		workspaces.PassportsModuleSetup(),

		// These are optional packages that might be used or might not be needed.
		geo.GeoModuleSetup(),
		accessibility.AccessibilityModuleSetup(),
		widget.WidgetModuleSetup(),
		commonprofile.CommonProfileModuleSetup(),
		currency.CurrencyModuleSetup(),
		licenses.LicensesModuleSetup(),
		worldtimezone.LicensesModuleSetup(),
		{{ else }}

		workspaces.WorkspaceModuleMicroServiceSetup(),
		{{ end }}

		// do not remove this comment line - it's used by fireback to append new modules

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

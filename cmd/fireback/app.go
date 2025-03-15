package main

import (
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/workspaces"
	FbSelfService "github.com/torabian/fireback/modules/workspaces/codegen/selfservice"
)

var PRODUCT_NAMESPACENAME = "fireback"
var PRODUCT_DESCRIPTION = "Fireback core microservice - v" + workspaces.FIREBACK_VERSION
var PRODUCT_LANGUAGES = []string{"fa", "en"}

//go:embed all:ui
var ui embed.FS

var xapp = &workspaces.FirebackApp{
	Title:              PRODUCT_DESCRIPTION,
	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []workspaces.SearchProviderFn{
		workspaces.QueryMenusReact,
		workspaces.QueryRolesReact,
		// shop.QueryProductSubmissionsReact,
	},
	RunTus: func() {
		workspaces.LiftTusServer()
	},
	RunSocket: func(e *gin.Engine) {
		workspaces.HandleSocket(e)
	},
	RunSearch: workspaces.InjectReactiveSearch,
	PublicFolders: []workspaces.PublicFolderInfo{
		// You can set a series of static folders to be served along with fireback.
		// This is only for static content. For advanced MVX render templates, you need to
		// Bootstrap those themes
		// Add these two lines on the top of the file
		/////go:embed all:ui
		// var ui embed.FS
		// and then uncomment this, for example to serve static react or angular content

		// 		//go:embed all:selfservice
		// var selfservice embed.FS
		{Fs: &FbSelfService.FbSelfService, Folder: ".", Prefix: "/selfservice"},
		{Fs: &ui, Folder: "ui"},
	},
	SetupWebServerHook: func(e *gin.Engine, xs *workspaces.FirebackApp) {
		// You can uncomment the sample theme for shop here
		// zayshop.Bootstrap(e)

	},
	Modules: []*workspaces.ModuleProvider{
		// Important to setup the workspaces at first, so the capabilties module is there
		workspaces.WorkspaceModuleSetup(),
		workspaces.DriveModuleSetup(),
		workspaces.NotificationModuleSetup(),
		workspaces.PassportsModuleSetup(),
	},
}

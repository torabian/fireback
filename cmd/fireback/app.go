package main

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/abac"
	"github.com/torabian/fireback/modules/fireback"
	FBManage "github.com/torabian/fireback/modules/fireback/codegen/fireback-manage"
	FbSelfService "github.com/torabian/fireback/modules/fireback/codegen/selfservice"
	"github.com/urfave/cli"
)

var PRODUCT_NAMESPACENAME = "fireback"
var PRODUCT_DESCRIPTION = "Fireback core microservice - v" + fireback.FIREBACK_VERSION
var PRODUCT_LANGUAGES = []string{"fa", "en"}

// Fireback doesn't come with default ui in the cmd anymore.
// Fireback itself has 2 uis: Manage and SelfService.
// Developer needs to build them if necessary and put the static files in workspaces
// Folder. Fireback serves them on /manage and /selfservice, similarly child projects
// Can serve those react projects if they wanted to.
// //go:embed all:ui
// var ui embed.FS

var xapp = &fireback.FirebackApp{
	Title:              PRODUCT_DESCRIPTION,
	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []fireback.SearchProviderFn{
		abac.QueryMenusReact,
		abac.QueryRolesReact,
	},
	SeedersSync: func() {
		abac.PassportMethodSyncSeeders()
		abac.AppMenuSyncSeeders()
	},
	RunTus: func() {
		abac.LiftTusServer()
	},
	InjectSearchEndpoint: fireback.InjectReactiveSearch,
	PublicFolders: []fireback.PublicFolderInfo{
		// You can set a series of static folders to be served along with fireback.
		// This is only for static content. For advanced MVX render templates, you need to
		// Bootstrap those themes
		// Add these two lines on the top of the file
		/////go:embed all:ui
		// var ui embed.FS
		// and then uncomment this, for example to serve static react or angular content

		// 		//go:embed all:selfservice
		// var selfservice embed.FS
		{Fs: &FBManage.FirebackManageTmpl, Folder: ".", Prefix: "/manage"},
		{Fs: &FbSelfService.FbSelfService, Folder: ".", Prefix: "/selfservice"},
	},
	SetupWebServerHook: func(e *gin.Engine, xs *fireback.FirebackApp) {

	},
	Modules: append([]*fireback.ModuleProvider{
		// Add the very core module, such as capabilities
		fireback.FirebackModuleSetup(nil),

		{
			CliHandlers: []cli.Command{
				fireback.NewProjectCli(),
			},
		},
	}, abac.AbacCompleteModules()...),
}

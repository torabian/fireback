package main

import (
	"os"

	"github.com/gin-gonic/gin"

	{{ if .ctx.IsMonolith }}
	"{{ .ctx.ModuleName }}/cmd/{{ .ctx.Name }}-server/menu"
	{{ end }}

	"github.com/torabian/fireback/modules/workspaces"

	{{ if or (.ctx.SelfService) }}
	FbSelfService "github.com/torabian/fireback/modules/workspaces/codegen/selfservice"
	{{ end }}
	{{ if or (.ctx.FirebackManage)}}

	FBManage "github.com/torabian/fireback/modules/workspaces/codegen/fireback-manage"
	{{ end }}

	{{ if or (.ctx.FirebackManage) (.ctx.CreateReactProject) }}
	"embed"
	{{ end }}
)

var PRODUCT_NAMESPACENAME = "{{ .ctx.Name }}"
var PRODUCT_DESCRIPTION = "{{ .ctx.Description }}"
var PRODUCT_LANGUAGES = []string{"en"}

{{ if .ctx.FirebackManage }}

//go:embed all:manage
var manageui embed.FS

{{ end }}

{{ if .ctx.CreateReactProject }}

//go:embed all:ui
var ui embed.FS

{{ end }}

var xapp = &workspaces.FirebackApp{
	Title: PRODUCT_DESCRIPTION,
	{{ if ne .ctx.IsMonolith true }}
	MicroService:       true,
	{{ end }}

	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []workspaces.SearchProviderFn{
		{{ if .ctx.IsMonolith }}
		workspaces.QueryMenusReact,
		workspaces.QueryRolesReact,
		{{ end }}
	},
	SeedersSync: func() {
		{{ if .ctx.IsMonolith }}
		// Sample menu item to make it easier for demos
		workspaces.AppMenuSyncSeederFromFs(&menu.Menu, []string{"new-menu.yml"}, workspaces.QueryDSL{
			WorkspaceId: "system",
		})
		{{ end }}
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



		{{ if .ctx.CreateReactProject }}
			{Fs: &ui, Folder: "ui" },

		{{ end }}

		{{ if or (.ctx.SelfService) }}
			{Fs: &FbSelfService.FbSelfService, Folder: ".", Prefix: "/selfservice"},
		{{ end }}

		{{ if .ctx.FirebackManage }}
			// You can change the Prefix to something else for more security,
			// or make it only available internally over vpn
			{Fs: &FBManage.FirebackManageTmpl, Folder: ".", Prefix: "/manage"},
		{{ end }}
	},
	SetupWebServerHook: func(e *gin.Engine, xs *workspaces.FirebackApp) {

	},
	Modules: []*workspaces.ModuleProvider{
		{{ if ne .ctx.IsMonolith true }}
		/*
		// Projects generated as microservice, will not include the following modules,
		// and that's all the difference between microservice and monolith in fireback
		{{ end }}
		workspaces.WorkspaceModuleSetup(),
		workspaces.DriveModuleSetup(),
		workspaces.NotificationModuleSetup(),
		workspaces.PassportsModuleSetup(),
		
		{{ if ne .ctx.IsMonolith true }}
		*/

		// Instead of few *ModuleSetup above, we are adding microservice module,
		// which essentially changes the Authorization resolver to allow everything,
		// and adds Capability* tables into the database.
		// You can uncomment the WorkspaceModuleSetup or other default Modules and go back to normal.
		workspaces.FirebackMicroService(nil),
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

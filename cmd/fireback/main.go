package main

import (
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/torabian/fireback/modules/abac"
	"github.com/torabian/emi/lib/gorunner"
	"github.com/urfave/cli/v3"

	"github.com/torabian/fireback/modules/abac"
	"github.com/torabian/fireback/modules/fireback"
	// clitools registers every terminal/CLI-interactive fireback feature
	// (promptui/bubbletea prompts, os/exec service management, asynq
	// workers, graceful HTTP shutdown) via init(). It's tagged !wasm and
	// deliberately not imported by cmd/fireback-wasm.
	_ "github.com/torabian/fireback/modules/fireback/clitools"
	"github.com/torabian/fireback/modules/fireback/gintools"
	FBManage "github.com/torabian/fireback/modules/interfaces/fireback-manage"
	FbSelfService "github.com/torabian/fireback/modules/interfaces/selfservice"
	project_generator "github.com/torabian/fireback/modules/project-generator"
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

func main() {

	emiCommand := &cli.Command{
		Name:     "emi",
		Usage:    "Emi compiler - Backend-for-Frontend with automatic SDK generation.",
		Commands: gorunner.BuildCommands(),
	}

	modules := []*fireback.ModuleProvider{
		fireback.FirebackModuleSetup(nil),
		{
			CliHandlers: []*cli.Command{
				project_generator.NewProjectCli(),
				emiCommand,
			},
		},
	}

	// For fireback we have abac module added.
	modules = append(modules, abac.AbacCompleteModules()...)

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

		PublicFolders: []gintools.PublicFolderInfo{
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
		Modules: modules,
	}

	// This is an important setting for some kind of app which will be installed
	// it makes it easier for fireback to find the configuration.
	os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)

	// This AppStart function is a wrapper for few things commonly can handle entire backend project
	// startup. For mobile or desktop might other functionality be used.
	xapp.CommonHeadlessAppStart(func() {
		// If anything needs to be done after database initialized
		// fireback.RegionalContentSyncSeeders()
		// fireback.AppMenuSyncSeeders()
	})
}

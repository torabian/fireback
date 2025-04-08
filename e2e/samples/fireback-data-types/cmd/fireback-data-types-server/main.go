package main

import (
	"github.com/torabian/fireback/fireback-data-types/modules/product/tags"
	"github.com/urfave/cli"

	"github.com/torabian/fireback/fireback-data-types/modules/product"

	"os"

	"github.com/gin-gonic/gin"

	"github.com/torabian/fireback/modules/abac"
	"github.com/torabian/fireback/modules/workspaces"
)

var PRODUCT_NAMESPACENAME = "fireback-data-types"
var PRODUCT_DESCRIPTION = "Sample project which uses all datatypes avaialble in fireback"
var PRODUCT_LANGUAGES = []string{"en"}

var xapp = &workspaces.FirebackApp{
	Title: PRODUCT_DESCRIPTION,

	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders:    []workspaces.SearchProviderFn{},
	SeedersSync: func() {

	},
	RunTus: func() {

	},
	RunSocket: func(e *gin.Engine) {

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

	Modules: append([]*workspaces.ModuleProvider{
		// Add the very core module, such as capabilities
		tags.TagsModuleSetup(nil),
		product.ProductModuleSetup(nil),
		workspaces.FirebackModuleSetup(nil),
		{
			CliHandlers: []cli.Command{
				workspaces.NewProjectCli(),
			},
		},
	}, abac.AbacCompleteModules()...),
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

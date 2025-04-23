package main

import (
	"context"
	"embed"
	"log"
	"os"

	rt "runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/gin-gonic/gin"

	"github.com/torabian/fireback/modules/fireback"

	{{ if .ctx.IsMonolith }}
	"github.com/torabian/fireback/modules/abac"
	{{ end }}

)


//go:embed all:frontend/dist
var assets embed.FS

// App struct
type AppIPCBridge struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *AppIPCBridge {
	os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)
	return &AppIPCBridge{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *AppIPCBridge) startup(ctx context.Context) {
	os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)


	f, err := os.OpenFile(fireback.GetEnvironmentUris().AppLogDirectory+"desktop.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer f.Close()

		log.SetOutput(f)
		log.Println("This is a test log entry")

	}

	fireback.Doctor()
	fireback.CreateHttpServer(
		fireback.SetupHttpServer(xapp, fireback.HttpServerInstanceConfig{}),
		fireback.HttpServerInstanceConfig{
			Port:    61210,
			Monitor: false,
			SSL:     false,
			Slow:    false,
		},
	)
	a.ctx = ctx
}

func main() {

	// Create an instance of the app structure
	app := NewApp()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
		os.Exit(0)
	})
	FileMenu.AddText("Close Window", keys.CmdOrCtrl("w"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
		os.Exit(0)
	})
	FileMenu.AddText("New", keys.Key("n"), func(_ *menu.CallbackData) {
		// runtime.Quit(app.ctx)
		// os.Exit(0)
	})
	FileMenu.AddText("Search in the app", keys.Key("s"), func(_ *menu.CallbackData) {
		// runtime.Quit(app.ctx)
		// os.Exit(0)
	})
	FileMenu.AddText("New", keys.Key("n"), func(_ *menu.CallbackData) {
		// runtime.Quit(app.ctx)
		// os.Exit(0)
	})
	FileMenu.AddText("Edit", keys.Key("e"), func(_ *menu.CallbackData) {
		// runtime.Quit(app.ctx)
		// os.Exit(0)
	})

	FileMenu.AddText("Select 1st", keys.Key("1"), func(_ *menu.CallbackData) {
	})
	FileMenu.AddText("Select 2nd", keys.Key("2"), func(_ *menu.CallbackData) {
	})
	FileMenu.AddText("Select 3rd", keys.Key("3"), func(_ *menu.CallbackData) {
	})
	FileMenu.AddText("Select 4th", keys.Key("4"), func(_ *menu.CallbackData) {
	})
	FileMenu.AddText("Select 5th", keys.Key("5"), func(_ *menu.CallbackData) {
	})
	FileMenu.AddText("Select 6th", keys.Key("6"), func(_ *menu.CallbackData) {
	})
	FileMenu.AddText("Select 7th", keys.Key("7"), func(_ *menu.CallbackData) {
	})
	FileMenu.AddText("Select 8th", keys.Key("8"), func(_ *menu.CallbackData) {
	})
	FileMenu.AddText("Select 9th", keys.Key("9"), func(_ *menu.CallbackData) {
	})

	if rt.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	}

	// This is an important setting for some kind of app which will be installed
	// it makes it easier for fireback to find the configuration.
	os.Setenv("PRODUCT_UNIQUE_NAME", PRODUCT_NAMESPACENAME)

	// This AppStart function is a wrapper for few things commonly can handle entire backend project
	// startup. For mobile or desktop might other functionality be used.
	xapp.CommonHeadlessAppStart(func() {
		// If anything needs to be done after database initialized
	})

	// // Create application with options
	err := wails.Run(&options.App{
		Width:  1024,
		Height: 768,
		Menu:   AppMenu, // reference the menu above
		AssetServer: &assetserver.Options{
			Assets: assets,
			// Handler: NewFileLoader(),
		},

		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
		Logger:             nil,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				FullSizeContent:            false,
			},
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

 

var PRODUCT_NAMESPACENAME = "{{ .ctx.Name }}"
var PRODUCT_DESCRIPTION = "{{ .ctx.Description }}"
var PRODUCT_LANGUAGES = []string{"en"}



var xapp = &fireback.FirebackApp{
	Title: PRODUCT_DESCRIPTION,
	{{ if ne .ctx.IsMonolith true }}
	MicroService:       true,
	{{ end }}

	SupportedLanguages: PRODUCT_LANGUAGES,
	SearchProviders: []fireback.SearchProviderFn{
		{{ if .ctx.IsMonolith }}
		abac.QueryMenusReact,
		abac.QueryRolesReact,
		{{ end }}
	},
	SeedersSync: func() {
		 
	},

	{{ if ne .ctx.IsMonolith true }}
	/* File uploader is a part of drive module in abac module
	{{ end }}
	RunTus: func() {
		abac.LiftTusServer()
	},
	{{ if ne .ctx.IsMonolith true }}
	*/
	RunSocket: func(e *gin.Engine) {
		fireback.HandleSocket(e)
	},
	{{ end }}
	InjectSearchEndpoint:     fireback.InjectReactiveSearch,
	PublicFolders: []fireback.PublicFolderInfo{
		// You can set a series of static folders to be served along with fireback.
		// This is only for static content. For advanced MVX render templates, you need to
		// Bootstrap those themes
		// Add these two lines on the top of the file
		/////go:embed all:ui
		// var ui embed.FS
		// and then uncomment this, for example to serve static react or angular content
		// {Fs: &ui, Folder: "ui"},

	},
	SetupWebServerHook: func(e *gin.Engine, xs *fireback.FirebackApp) {

	},
	Modules: []*fireback.ModuleProvider{
		{{ if ne .ctx.IsMonolith true }}
		/*
		// Projects generated as microservice, will not include the following modules,
		// and that's all the difference between microservice and monolith in fireback
		{{ end }}
		abac.WorkspaceModuleSetup(),
		abac.DriveModuleSetup(),
		abac.NotificationModuleSetup(),
		abac.PassportsModuleSetup(),
		
		{{ if ne .ctx.IsMonolith true }}
		*/

		// Instead of few *ModuleSetup above, we are adding microservice module,
		// which essentially changes the Authorization resolver to allow everything,
		// and adds Capability* tables into the database.
		// You can uncomment the WorkspaceModuleSetup or other default Modules and go back to normal.
		fireback.FirebackModuleSetup(nil),
		{{ end }}

		// do not remove this comment line - it's used by fireback to append new modules

	},
}


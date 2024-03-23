package main

import (
	"context"
	"embed"
	"log"
	"os"

	rt "runtime"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {

	workspaces.BundledConfig = &workspaces.AppConfig{
		Name:       "EspStudioDesktop",
		SelfHosted: true,
		Drive: workspaces.Drive{
			Storage: "{appDataDirectory}storage",
			Enabled: true,
			Port:    "61211",
		},
		Headers: workspaces.Headers{
			AccessControlAllowOrigin: "*",
			AccessControlAllowHeaders: `Accept, Authorization, Content-Type, Content-Length,
			  X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language,
			  X-Requested-With, Workspace, Workspace-Id, deep, query, role-id`,
		},

		PublicServer: workspaces.PublicServer{
			Enabled: true,
			Port:    "53610",
			Host:    "localhost",
		},
		Database: workspaces.Database{
			Database: "{appDataDirectory}fireback-desktop-database.db",
			Vendor:   "sqlite",
		},
	}

}

// App struct
type AppIPCBridge struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *AppIPCBridge {
	os.Setenv("PRODUCT_UNIQUE_NAME", "esp-studio-desktop")
	return &AppIPCBridge{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *AppIPCBridge) startup(ctx context.Context) {
	os.Setenv("PRODUCT_UNIQUE_NAME", "esp-studio-desktop")

	xapp := baseService()

	f, err := os.OpenFile(workspaces.GetEnvironmentUris().AppLogDirectory+"esp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer f.Close()
		log.SetOutput(f)
		log.Println("This is a test log entry")

	}

	workspaces.Doctor()
	workspaces.CreateHttpServer(workspaces.SetupHttpServer(xapp))
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
	FileMenu.AddText("Back", keys.Key("escape"), func(_ *menu.CallbackData) {
		// runtime.Quit(app.ctx)
		// os.Exit(0)
	})

	if rt.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	}

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

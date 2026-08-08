package main

import (
	"os"

	"github.com/torabian/fireback/modules/fireback"
)

// Fireback doesn't come with default ui in the cmd anymore.
// Fireback itself has 2 uis: Manage and SelfService.
// Developer needs to build them if necessary and put the static files in workspaces
// Folder. Fireback serves them on /manage and /selfservice, similarly child projects
// Can serve those react projects if they wanted to.
// //go:embed all:ui
// var ui embed.FS

var xapp = &fireback.FirebackApp{
	Title: "Fireback core microservice - v" + fireback.FIREBACK_VERSION,
	Modules: []*fireback.ModuleProvider{
		// Add the very core module, such as capabilities
		fireback.FirebackModuleSetup(nil),
	},
}

func main() {

	// This is an important setting for some kind of app which will be installed
	// it makes it easier for fireback to find the configuration.
	os.Setenv("PRODUCT_UNIQUE_NAME", "fireback")

	// This AppStart function is a wrapper for few things commonly can handle entire backend project
	// startup. For mobile or desktop might other functionality be used.
	xapp.CommonHeadlessAppStart(func() {
		// If anything needs to be done after database initialized
		// fireback.RegionalContentSyncSeeders()
		// fireback.AppMenuSyncSeeders()
	})
}

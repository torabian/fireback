package main

import (
	"fmt"
	"os"

	"github.com/torabian/fireback/modules/fireback/application"
)

// Fireback doesn't come with default ui in the cmd anymore.
// Fireback itself has 2 uis: Manage and SelfService.
// Developer needs to build them if necessary and put the static files in workspaces
// Folder. Fireback serves them on /manage and /selfservice, similarly child projects
// Can serve those react projects if they wanted to.
// //go:embed all:ui
// var ui embed.FS

var xapp = &application.Application{

	Modules: []*application.ModuleProvider{},
}

func main() {

	fmt.Println(1, xapp)
	// This is an important setting for some kind of app which will be installed
	// it makes it easier for fireback to find the configuration.
	os.Setenv("PRODUCT_UNIQUE_NAME", "fireback")

}

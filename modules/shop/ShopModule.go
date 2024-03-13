package shop

import (
	"embed"
	// "github.com/urfave/cli"
	"gorm.io/gorm"
	"github.com/torabian/fireback/modules/workspaces"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func ShopModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name: "shop",
        Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {

	})

	module.ProvideSeederImportHandler(func() {

	})

	module.ProvideMockWriterHandler(func(languages []string) {

	})

	// module.ProvidePermissionHandler()

	// module.Actions = [][]workspaces.RouteDefinition{}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		// if err := dbref.AutoMigrate(& Entity{}); err != nil {
		// 	fmt.Println(err.Error())
		// }

	})

	/*
	module.ProvideCliHandlers([]cli.Command{
		{
			Name:        "",
			Usage:       "",
			Subcommands: cli.Commands{},
		},
	}) */

	return module
}

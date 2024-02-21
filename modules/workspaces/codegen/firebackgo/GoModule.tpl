package {{ .path }}

import (
	"embed"
	// "github.com/urfave/cli"
	"gorm.io/gorm"
	"pixelplux.com/fireback/modules/workspaces"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func {{ .Name }}ModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name: "{{ .name }}",
        Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {

	})

	module.ProvideSeederImportHandler(func() {

	})

	module.ProvideMockWriterHandler(func(languages []string) {

	})

	// module.ProvidePermissionHandler()

	// Add entity or actions here
	module.Actions = [][]workspaces.Module2Action{
	
	}

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

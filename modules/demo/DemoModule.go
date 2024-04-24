package demo

import (
	"embed"
	"fmt"

	// "github.com/urfave/cli"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func DemoModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "demo",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {
		CustomerImportMocks()
	})

	module.ProvideSeederImportHandler(func() {

	})

	module.ProvideMockWriterHandler(func(languages []string) {
		CustomerWriteQueryMock(workspaces.MockQueryContext{Languages: languages, ItemsPerPage: 20})
	})

	// module.ProvidePermissionHandler()

	module.Actions = [][]workspaces.Module2Action{
		GetCustomerModule2Actions(),
		DemoCustomActions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(&CustomerEntity{}, &CustomerAddress{}); err != nil {
			fmt.Println(err.Error())
		}

	})

	module.ProvideCliHandlers(append([]cli.Command{
		CustomerCliFn(),
	}, DemoCustomActionsCli...))

	return module
}

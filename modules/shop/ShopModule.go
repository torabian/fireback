package shop

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

func ShopModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "shop",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {
		// FormImportMocks()
	})

	module.ProvideSeederImportHandler(func() {

	})

	module.ProvideMockWriterHandler(func(languages []string) {

	})

	module.ProvidePermissionHandler(
		ALL_PRODUCT_PERMISSIONS,
		ALL_PRODUCTSUBMISSION_PERMISSIONS,
	)

	module.Actions = [][]workspaces.Module2Action{
		GetProductSubmissionModule2Actions(),
		GetProductModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(
			&ProductEntity{},
			&ProductFields{},
			&ProductSubmissionEntity{},
			&ProductSubmissionValues{},
		); err != nil {
			fmt.Println(err.Error())
		}

	})

	module.ProvideCliHandlers([]cli.Command{
		ProductCliFn(),
		ProductSubmissionCliFn(),
	})

	return module
}

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
		FormImportMocks()
	})

	module.ProvideSeederImportHandler(func() {

	})

	module.ProvideMockWriterHandler(func(languages []string) {

	})

	module.ProvidePermissionHandler(
		ALL_FORM_PERMISSIONS,
		ALL_FORMDATA_PERMISSIONS,
	)

	module.Actions = [][]workspaces.Module2Action{
		GetFormDataModule2Actions(),
		GetFormModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(
			&FormEntity{},
			&FormFields{},
			&FormDataEntity{},
			&FormDataValues{},
		); err != nil {
			fmt.Println(err.Error())
		}

	})

	module.ProvideCliHandlers([]cli.Command{
		FormCliFn(),
		FormDataCliFn(),
	})

	return module
}

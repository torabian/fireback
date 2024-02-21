package currency

import (
	"embed"
	"fmt"

	"github.com/urfave/cli"
	"gorm.io/gorm"
	"pixelplux.com/fireback/modules/workspaces"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func CurrencyModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "currency",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {
		CurrencySyncSeeders()
	})

	module.ProvideMockWriterHandler(func(languages []string) {
		CurrencyWriteQueryMock(workspaces.MockQueryContext{Languages: languages})
	})

	module.ProvidePermissionHandler(ALL_CURRENCY_PERMISSIONS, ALL_PRICETAG_PERMISSIONS)

	module.Actions = [][]workspaces.Module2Action{
		GetCurrencyModule2Actions(),
		GetPriceTagModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		if err := dbref.AutoMigrate(&CurrencyEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&CurrencyEntityPolyglot{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&PriceTagEntity{}); err != nil {
			fmt.Println(err.Error())
		}
		if err := dbref.AutoMigrate(&PriceTagVariations{}); err != nil {
			fmt.Println(err.Error())
		}
	})

	module.ProvideCliHandlers([]cli.Command{
		CurrencyCliFn(),
		PriceTagCliFn(),
	})

	return module
}

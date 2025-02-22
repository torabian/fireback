package currency

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func CurrencyModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "currency",
		Definitions: &Module3Definitions,
	}

	module.ProvideMockImportHandler(func() {
		CurrencySyncSeeders()
	})

	module.ProvideMockWriterHandler(func(languages []string) {
		CurrencyWriteQueryMock(workspaces.MockQueryContext{Languages: languages})
	})

	module.ProvidePermissionHandler(ALL_CURRENCY_PERMISSIONS, ALL_PRICE_TAG_PERMISSIONS)

	module.Actions = [][]workspaces.Module3Action{
		GetCurrencyModule3Actions(),
		GetPriceTagModule3Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&CurrencyEntity{},
			&CurrencyEntityPolyglot{},
			&PriceTagEntity{},
			&PriceTagVariations{},
		)
	})

	module.ProvideCliHandlers([]cli.Command{
		CurrencyCliFn(),
		PriceTagCliFn(),
	})

	return module
}

package widget

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func WidgetModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "widget",
		Definitions: &Module2Definitions,
	}

	module.ProvideMockImportHandler(func() {
		WidgetImportMocks()
		WidgetAreaImportMocks()
	})

	module.ProvideMockWriterHandler(func(languages []string) {
		WidgetAreaWriteQueryMock(workspaces.MockQueryContext{Languages: languages, WithPreloads: []string{"Widgets.Widget"}})
		WidgetWriteQueryMock(workspaces.MockQueryContext{Languages: languages})
	})

	module.ProvidePermissionHandler(
		ALL_WIDGET_AREA_PERMISSIONS,
		ALL_WIDGET_PERMISSIONS,
	)

	module.Actions = [][]workspaces.Module2Action{
		GetWidgetAreaModule2Actions(),
		GetWidgetModule2Actions(),
	}

	module.Actions = [][]workspaces.Module2Action{
		GetWidgetAreaModule2Actions(),
		GetWidgetModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&WidgetAreaEntity{},
			&WidgetAreaWidgets{},
			&WidgetAreaEntityPolyglot{},
			&WidgetEntity{},
			&WidgetAreaEntityPolyglot{},
			&WidgetEntityPolyglot{},
		)
	})

	module.ProvideCliHandlers([]cli.Command{
		WidgetAreaCliFn(),
		WidgetCliFn(),
	})

	return module
}

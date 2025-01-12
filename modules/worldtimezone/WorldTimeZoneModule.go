package worldtimezone

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func LicensesModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "worldtimezone",
		Definitions: &Module3Definitions,
	}

	module.ProvidePermissionHandler(ALL_TIMEZONE_GROUP_PERMISSIONS)

	module.Actions = [][]workspaces.Module3Action{
		GetTimezoneGroupModule3Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(
			&TimezoneGroupEntity{},
			&TimezoneGroupEntityPolyglot{},
			&TimezoneGroupUtcItems{},
		)
	})

	module.ProvideCliHandlers([]cli.Command{
		TimezoneGroupCliFn(),
	})

	return module
}

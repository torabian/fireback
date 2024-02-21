package worldtimezone

import (
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func LicensesModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name: "worldtimezone",
	}

	module.ProvidePermissionHandler(ALL_TIMEZONEGROUP_PERMISSIONS)

	module.Actions = [][]workspaces.Module2Action{
		GetTimezoneGroupModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		dbref.AutoMigrate(&TimezoneGroupEntity{})
		dbref.AutoMigrate(&TimezoneGroupEntityPolyglot{})
		dbref.AutoMigrate(&TimezoneGroupUtcItemsEntity{})
		dbref.AutoMigrate(&TimezoneGroupUtcItemsEntityPolyglot{})
	})

	module.ProvideCliHandlers([]cli.Command{
		TimezoneGroupCliFn(),
	})

	return module
}

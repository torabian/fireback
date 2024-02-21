package drive

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func DriveModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "drive",
		Definitions: &Module2Definitions,
	}

	module.ProvidePermissionHandler(ALL_FILE_PERMISSIONS)

	// Drive is not coverting route definitions, needs to be fixed
	module.Actions = [][]workspaces.Module2Action{
		GetFileModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		dbref.AutoMigrate(&FileEntity{})
	})

	return module
}

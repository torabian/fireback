package workspaces

import (
	"embed"

	"gorm.io/gorm"
)

//go:embed *Module3.yml
var DriveDefinition embed.FS

func DriveModuleSetup() *ModuleProvider {
	module := &ModuleProvider{
		Name:        "drive",
		Definitions: &DriveDefinition,
	}

	module.ProvidePermissionHandler(ALL_FILE_PERMISSIONS)

	// Drive is not coverting route definitions, needs to be fixed
	module.Actions = [][]Module2Action{
		GetFileModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) {
		dbref.AutoMigrate(&FileEntity{})
	})

	return module
}

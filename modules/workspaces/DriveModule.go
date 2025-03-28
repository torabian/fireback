package workspaces

import (
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func DriveModuleSetup() *ModuleProvider {
	module := &ModuleProvider{
		// This is also weird for me. We need a mechanism for naming module better
		// now because of react/java/swift compiler I write this the same name as folder.
		Name: "workspaces",
	}

	module.ProvidePermissionHandler(ALL_FILE_PERMISSIONS)

	// Drive is not coverting route definitions, needs to be fixed
	module.Actions = [][]Module3Action{
		GetFileModule3Actions(),
	}

	module.ProvideCliHandlers([]cli.Command{
		FileCliFn(),
	})

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		return dbref.AutoMigrate(&FileEntity{}, &FileVariations{})
	})

	return module
}

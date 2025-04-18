package abac

import (
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func DriveModuleSetup() *workspaces.ModuleProvider {

	// Overriding the uploading mechanism
	workspaces.ImportYamlFromFsResources = ImportYamlFromFsResources

	module := &workspaces.ModuleProvider{
		// This is also weird for me. We need a mechanism for naming module better
		// now because of react/java/swift compiler I write this the same name as folder.
		Name: "abac",
	}

	module.ProvidePermissionHandler(ALL_FILE_PERMISSIONS)

	// Drive is not coverting route definitions, needs to be fixed
	module.Actions = [][]workspaces.Module3Action{
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

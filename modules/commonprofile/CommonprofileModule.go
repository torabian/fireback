package commonprofile

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func CommonProfileModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "commonprofile",
		Definitions: &Module2Definitions,
	}
	module.ProvidePermissionHandler(ALL_COMMONPROFILE_PERMISSIONS)

	module.Actions = [][]workspaces.Module2Action{
		GetCommonProfileModule2Actions(),
	}

	module.ProvideCliHandlers([]cli.Command{
		CommonProfileCliFn(),
	})

	return module
}

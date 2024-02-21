package commonprofile

import (
	"embed"

	"github.com/urfave/cli"
	"pixelplux.com/fireback/modules/workspaces"
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

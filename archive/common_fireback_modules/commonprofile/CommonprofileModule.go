package commonprofile

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func CommonProfileModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "commonprofile",
		Definitions: &Module3Definitions,
	}
	module.ProvidePermissionHandler(ALL_COMMON_PROFILE_PERMISSIONS)

	module.Actions = [][]workspaces.Module3Action{
		GetCommonProfileModule3Actions(),
	}

	module.ProvideCliHandlers([]cli.Command{
		CommonProfileCliFn(),
	})

	return module
}

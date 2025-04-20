package commonprofile

import (
	"embed"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func CommonProfileModuleSetup() *fireback.ModuleProvider {
	module := &fireback.ModuleProvider{
		Name:        "commonprofile",
		Definitions: &Module3Definitions,
	}
	module.ProvidePermissionHandler(ALL_COMMON_PROFILE_PERMISSIONS)

	module.Actions = [][]fireback.Module3Action{
		GetCommonProfileModule3Actions(),
	}

	module.ProvideCliHandlers([]cli.Command{
		CommonProfileCliFn(),
	})

	return module
}

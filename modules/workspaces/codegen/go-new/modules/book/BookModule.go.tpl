package book

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func BookModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "book",
		Definitions: &Module2Definitions,
		EntityBundles: []workspaces.EntityBundle{
			BookEntityBundle,
		},
	}

	return module
}

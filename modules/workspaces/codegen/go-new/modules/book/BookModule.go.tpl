package book

import (
	"embed"

	"github.com/torabian/fireback/modules/workspaces"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func BookModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "book",
		Definitions: &Module3Definitions,
		EntityBundles: []workspaces.EntityBundle{
			BookEntityBundle,
		},
	}

	return module
}

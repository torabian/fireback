package {{ .path }}

import (
	"embed"
	"github.com/torabian/fireback/modules/workspaces"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func {{ .Name }}ModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name: "{{ .name }}",
        Definitions: &Module2Definitions,
		EntityBundles: []workspaces.EntityBundle{
			// Do not remove this comment, aef0
		},
	}

	return module
}

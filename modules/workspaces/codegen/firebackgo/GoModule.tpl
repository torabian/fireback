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
			// Insert the NameEntityBundle here.
			// each entity, has multiple features, such as permissions, events, translations
			// *EntityBundle objects are a list of them which are auto generated,
			// and by adding them here it will be automatically added.
			// we cannot add them automatically upon saving yaml for you,
			// when you add a new entity in yaml, add it manually here.
		},
	}

	return module
}

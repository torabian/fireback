//go:build !wasm

package abac

import (
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

func PassportsModuleSetup() *application.ModuleProvider {
	module := &application.ModuleProvider{

		// it must write on the workspaces instead
		Name: "abac",

		// passportMethod/publicJoinKey/emailConfirmation/phoneConfirmation moved from
		// AbacModule3.yml's old entities: section to Abac.emi.yml, so they're wired
		// directly here now, the same way FirebackModuleSetup wires Capability* actions.
		GinWebServerInitHooks: GinWebServerInitHooks,
	}

	module.ProvidePermissionHandler(PassportProvidePermissionHandler...)

	module.ProvideEntityHandlers(PassportProvideEntityHandlers)

	module.ProvideCliHandlers([]*cli.Command{
		&PassportCli,
	})

	return module
}

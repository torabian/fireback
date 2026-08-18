package fireback

import (
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/migrations"
)

type FirebackModuleConfig struct{}

func FirebackModuleSetup(setup *FirebackModuleConfig) *application.ModuleProvider {

	module := &application.ModuleProvider{
		Name:               "fireback",
		GoMigrateDirectory: &migrations.MigrationsFs,
	}

	module.ProvideCliHandlers(firebackModuleCliHandlers())

	return module
}

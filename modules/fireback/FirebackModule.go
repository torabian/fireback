package fireback

import (
	"github.com/torabian/fireback/modules/fireback/migrations"
	"github.com/urfave/cli/v3"
)

type FirebackModuleConfig struct{}

func FirebackModuleSetup(setup *FirebackModuleConfig) *ModuleProvider {

	module := &ModuleProvider{
		Name:               "fireback",
		GoMigrateDirectory: &migrations.MigrationsFs,
	}

	module.ProvideCliHandlers([]*cli.Command{
		&PushNotificationCmd,
	})

	return module
}

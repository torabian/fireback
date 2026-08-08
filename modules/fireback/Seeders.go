package fireback

import (
	"context"

	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

func GetSeeder(xapp *application.Application) *cli.Command {
	return &cli.Command{

		Name:  "seeders",
		Usage: "Imports all necessary seeders",
		Action: func(ctx context.Context, c *cli.Command) error {
			ExecuteSeederImport(xapp)
			return nil
		},
	}
}

func ExecuteSeederImport(x *application.Application) {

	for _, item := range x.Modules {
		if item.SeederHandler != nil {

			item.SeederHandler()
		}

	}

	if x.SeedersSync != nil {
		x.SeedersSync()
	}
}

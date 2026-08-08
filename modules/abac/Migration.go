package abac

import (
	"context"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

func GetMigrationCommand(xapp *application.Application) *cli.Command {
	return &cli.Command{

		Name:  "migration",
		Usage: "Database and content migration, syncing the application entities with database",
		Commands: []*cli.Command{
			GetCapabilityRefreshCommand(xapp),
			{
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:  "level",
						Usage: "Silent = 1, Error = 2, Warn = 3, Info = 4 (Default is 2, errors shown)",
						Value: 2,
					},
				},
				Name:  "apply",
				Usage: "Applies all necessary migration code on database or other infrastructure the the project.",
				Action: func(ctx context.Context, c *cli.Command) error {

					fireback.ApplyMigration(xapp, c.Int64("level"))
					SyncPermissionsInDatabase(xapp, fireback.GetDbRef())

					return nil
				},
			},
		},
	}

}

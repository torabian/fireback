package fireback

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func GetMigrationCommand(xapp *FirebackApp) *cli.Command {
	return &cli.Command{

		Name:  "migration",
		Usage: "Database and content migration, syncing the application entities with database",
		Commands: []*cli.Command{
			GetCapabilityRefreshCommand(xapp),
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Usage:    "The address of file you want the yaml be exported to",
						Required: true,
					},
				},
				Name:  "export",
				Usage: "Exports the content of the migration based on the criteria",
				Action: func(ctx context.Context, c *cli.Command) error {
					xinfo := []TableMetaData{}

					for _, module := range xapp.Modules {
						for _, item := range module.BackupTables {
							xinfo = append(xinfo, item)
						}
					}

					fmt.Println("File", c.String("file"))
					// CreateBackup(xinfo, c.String("file"))

					return nil
				},
			},
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

					ApplyMigration(xapp, c.Int64("level"))
					SyncPermissionsInDatabase(xapp, GetDbRef())

					return nil
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Usage:    "The address of file you want the yaml be exported to",
						Required: true,
					},
				},
				Name:  "import",
				Usage: "Import system data from a previous export",
				Action: func(ctx context.Context, c *cli.Command) error {
					xinfo := []TableMetaData{}
					// f := CommonCliQueryDSLBuilder(c)

					for _, module := range xapp.Modules {
						for _, item := range module.BackupTables {
							xinfo = append(xinfo, item)
						}
					}

					fmt.Println("File", c.String("file"))
					// ImportBackup(xinfo, c.String("file"), f)

					return nil
				},
			},
		},
	}

}

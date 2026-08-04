package fireback

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Fireback generates a lot of code, this file contains fireback gen command, and combines
// different operation into a single one.

var cliGlobalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "al",
		Usage: "Set's the language of the query, equal to accept-language header in http requests",
		Value: "en-us",
	},
}

func GetSeeder(xapp *FirebackApp) *cli.Command {
	return &cli.Command{

		Name:  "seeders",
		Usage: "Imports all necessary seeders",
		Action: func(ctx context.Context, c *cli.Command) error {
			ExecuteSeederImport(xapp)
			return nil
		},
	}
}

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

func GetApplicationTasks(xapp *FirebackApp) *cli.Command {
	sub := []*cli.Command{}

	for _, m := range xapp.Modules {
		for _, t := range m.Tasks {
			sub = append(sub, &cli.Command{
				Name:   t.Name,
				Flags:  t.Flags,
				Action: t.Cli,
			})
		}
	}

	return &cli.Command{
		Name:  "tasks",
		Usage: "Actions related to the project tasks, running them in background, list, etc.",
		Commands: []*cli.Command{

			{
				Name:     "enqueue",
				Usage:    "Enqueues a task to the queue so worker can pick it up",
				Commands: sub,
			},
			{
				Name:  "list",
				Usage: "Lists all of the tasks in the app",
				Action: func(ctx context.Context, c *cli.Command) error {
					for _, m := range xapp.Modules {
						for _, t := range m.Tasks {

							fmt.Println(t.Name)
						}
					}
					return nil
				},
			},
			{
				Name:  "start",
				Usage: "Starts the background worker server",
				Action: func(ctx context.Context, c *cli.Command) error {
					taskServerLifter(xapp)
					return nil
				},
			},
		},
	}
}

func taskServerLifter(xapp *FirebackApp) {

	tasks := []*TaskAction{}
	for _, m := range xapp.Modules {
		for _, t := range m.Tasks {
			tasks = append(tasks, t)
		}
	}

	LiftAsyncqWorkerServer(tasks)
}

package fireback

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Fireback generates a lot of code, this file contains fireback gen command, and combines
// different operation into a single one.

var fbGoModuleFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "name",
		Usage:    "Name of the module - name will be used to create the yaml file",
		Required: true,
	},
	&cli.StringFlag{
		Name:  "dir",
		Usage: "The directory which will the module be created - if not set, the name of module will be used",
	},
	&cli.StringFlag{
		Name:  "auto-import",
		Usage: "It would add the module, into a server or desktop main app file in fireback if file path is given, also the magic comment exists as well",
	},
}

var cliGlobalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "al",
		Usage: "Set's the language of the query, equal to accept-language header in http requests",
		Value: "en-us",
	},
}
var commonFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "sdk-dir",
		Usage: "Location of the sdk for UI projects",
	},
	&cli.StringFlag{
		Name:  "fb-ui-dir",
		Usage: "The location that fireback UI components and common hooks is located",
	},
	&cli.StringFlag{
		Name:  "path",
		Usage: "Address of the folder, which the content will be generated into",
		// Required: true,
	},
	&cli.StringFlag{
		Name:  "relative-to",
		Usage: "Address of the relative folder to the modules, for go files",
		// Required: true,
	},
	&cli.StringFlag{
		Name:  "openapi",
		Usage: "Use openapi 3 definitions to feed into the codegen",
	},
	&cli.StringFlag{
		Name:  "no-cache",
		Usage: "Ignores the cache",
	},
	&cli.StringFlag{
		Name:  "modules",
		Usage: "build only specific modules, for example --modules workspaces,iot",
	},
	&cli.StringFlag{
		Name:  "def",
		Usage: "Gets the module file from disk, and compiles it, instead of internal definition files",
	},
	&cli.StringFlag{
		Name:  "gof-module",
		Usage: "Go module name in go mod for generation",
	},
}

var reactFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "no-fbdef",
		Usage: "Skip include the fireback definition as json into dto/entity classes",
	},
	&cli.BoolFlag{
		Name:  "no-nav",
		Usage: "Skip include the navigation urls into the fireback entities",
	},
	&cli.BoolFlag{
		Name:  "no-static",
		Usage: "Skip include the static string fields in in process",
	},
}

var reactUIFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "entity-path",
		Usage:    "Address of the entity on binary with module full address (fireback.User)",
		Required: true,
	},
}

func GetReportsTool(xapp *FirebackApp) cli.Command {
	return cli.Command{

		Name:  "reports",
		Usage: "Views all the reports available in the system",
		Flags: append(CommonQueryFlags,
			&cli.StringFlag{
				Name:     "file",
				Usage:    "The address of file you want the csv/yaml/json/pdf be exported to",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "id",
				Usage:    "Report id",
				Required: false,
			},
		),
		Action: func(ctx context.Context, c *cli.Command) error {

			reports := []Report{}
			for _, m := range xapp.Modules {
				reports = append(reports, m.Reports...)
			}
			f := CommonCliQueryDSLBuilder(c)
			var report *Report
			var file string
			if c.String("id") != "" {
				report = GetReportById(c.String("id"), reports)
			} else {
				report = GetReport(reports)
			}
			if c.String("file") != "" {
				file = c.String("file")
			} else {
				file = AskForInput("Where to export the report", "report.pdf")
			}

			if report == nil {
				fmt.Println("No report has been selected")
				return nil
			}

			report.Fn(file, f, report, report.V)

			return nil
		},
	}
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

	liftAsyncqWorkerServer(tasks)
}

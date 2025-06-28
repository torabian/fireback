package fireback

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	reactui "github.com/torabian/fireback/modules/fireback/codegen/react-ui"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v3"
)

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

var reconfigFlag = []cli.Flag{
	&cli.StringFlag{
		Name:     "binary-name",
		Usage:    "Binary name that will be used to access final binary",
		Required: true,
	},

	&cli.StringFlag{
		Name:     "project",
		Usage:    "Project name on files and disks",
		Required: true,
	},

	&cli.StringFlag{
		Name:  "description",
		Usage: "Description of the project",
	},

	&cli.StringFlag{
		Name:  "languages",
		Usage: "Languages that this support",
		Value: "en, fa",
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

func GenContextFromCli(c *cli.Context, cat CodeGenCatalog) *CodeGenContext {
	tsx := &TypeScriptGenContext{
		IncludeStaticField:      true,
		IncludeFirebackDef:      true,
		IncludeStaticNavigation: true,
	}

	if c.IsSet("no-fbdef") {
		tsx.IncludeFirebackDef = false
	}

	if c.IsSet("no-static") {
		tsx.IncludeStaticField = false
	}

	if c.IsSet("no-nav") {
		tsx.IncludeStaticNavigation = false
	}

	GofModuleName := "github.com/torabian/fireback"

	if c.String("gof-module") != "" {
		GofModuleName = c.String("gof-module")
	}

	ctx := &CodeGenContext{
		Path:          c.String("path"),
		OpenApiFile:   c.String("openapi"),
		GofModuleName: GofModuleName,
		EntityPath:    c.String("entity-path"),
		Catalog:       cat,
		NoCache:       c.Bool("no-cache"),
		Modules:       strings.Split(c.String("modules"), ","),
		Ts:            *tsx,
	}

	if c.IsSet("sdk-dir") {
		ctx.UiSdkDir = c.String("sdk-dir")
	}

	if c.IsSet("fb-ui-dir") {
		ctx.FirebackUIDir = c.String("fb-ui-dir")
	}

	if c.IsSet("def") {
		ctx.ModulesOnDisk = strings.Split(c.String("def"), ",")
	}

	return ctx
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
		Action: func(c *cli.Context) error {

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

func GetSeeder(xapp *FirebackApp) cli.Command {
	return cli.Command{

		Name:  "seeders",
		Usage: "Imports all necessarys eeders",
		Action: func(c *cli.Context) error {
			ExecuteSeederImport(xapp)
			return nil
		},
	}
}

func GetMigrationCommand(xapp *FirebackApp) cli.Command {
	return cli.Command{

		Name:  "migration",
		Usage: "Database and content migration, syncing the application entities with database",
		Subcommands: cli.Commands{
			GetCapabilityRefreshCommand(xapp),
			cli.Command{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Usage:    "The address of file you want the yaml be exported to",
						Required: true,
					},
				},
				Name:  "export",
				Usage: "Exports the content of the migration based on the criteria",
				Action: func(c *cli.Context) error {
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
			cli.Command{
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:  "level",
						Usage: "Silent = 1, Error = 2, Warn = 3, Info = 4 (Default is 2, errors shown)",
						Value: 2,
					},
				},
				Name:  "apply",
				Usage: "Applies all necessary migration code on database or other infrastructure the the project.",
				Action: func(c *cli.Context) error {

					ApplyMigration(xapp, c.Int64("level"))
					SyncPermissionsInDatabase(xapp, GetDbRef())

					return nil
				},
			},
			cli.Command{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Usage:    "The address of file you want the yaml be exported to",
						Required: true,
					},
				},
				Name:  "import",
				Usage: "Import system data from a previous export",
				Action: func(c *cli.Context) error {
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

func GetApplicationTasks(xapp *FirebackApp) cli.Command {
	sub := []cli.Command{}

	for _, m := range xapp.Modules {
		for _, t := range m.Tasks {
			sub = append(sub, cli.Command{
				Name:   t.Name,
				Flags:  t.Flags,
				Action: t.Cli,
			})
		}
	}

	return cli.Command{

		Name:  "tasks",
		Usage: "Actions related to the project tasks, running them in background, list, etc.",
		Subcommands: cli.Commands{

			{
				Name:        "enqueue",
				Usage:       "Enqueues a task to the queue so worker can pick it up",
				Subcommands: sub,
			},
			{
				Name:  "list",
				Usage: "Lists all of the tasks in the app",
				Action: func(c *cli.Context) error {
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
				Action: func(c *cli.Context) error {
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

// func GetApplicationTests(xapp *FirebackApp) cli.Command {
// 	return cli.Command{
// 		Name:  "tests",
// 		Usage: "Tools and actions related to the products testing",
// 		Subcommands: cli.Commands{

// 			{
// 				Name: "run",
// 				Flags: []cli.Flag{
// 					cli.StringFlag{
// 						Name:     "n",
// 						Usage:    "Run specific test - by default we run all tests",
// 						Required: false,
// 					},
// 				},
// 				Usage: "Runs the tests on the product.",
// 				Action: func(c *cli.Context) error {
// 					query := CommonCliQueryDSLBuilder(c)
// 					ctx := TestContext{F: query}
// 					for _, m := range xapp.Modules {
// 						tests := m.Tests
// 						for _, test := range tests {
// 							if c.IsSet("n") && test.Name != c.String("n") {
// 								continue
// 							}
// 							err := test.Function(&ctx)
// 							if err == nil {
// 								c := color.New(color.FgGreen)
// 								fmt.Print("\u2713 Test \"")
// 								c.Print(test.Name)
// 								fmt.Print("\" Has passed successfully")
// 							}
// 							fmt.Println("")
// 						}
// 					}
// 					return nil
// 				},
// 			},
// 			{
// 				Name:  "list",
// 				Usage: "Lists all of the tests in the app",
// 				Action: func(c *cli.Context) error {
// 					for _, m := range xapp.Modules {
// 						for _, t := range m.Tests {
// 							fmt.Println(t.Name)
// 						}
// 					}
// 					return nil
// 				},
// 			},
// 		},
// 	}
// }

func CodeGenTools(xapp *FirebackApp) cli.Command {
	return cli.Command{
		Name:  "gen",
		Usage: "Code generation tools, both for internal codes and sdk remote files",
		Subcommands: cli.Commands{
			{
				Name: "module3spec",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "source",
						Usage: "You can pass a yaml file address on disk, to make the schema based on that. If left empty, an empty module3 file will be considered.",
					},
					cli.StringFlag{
						Name:  "out",
						Usage: "Where to write output. If not set, result will be printed to stdout",
					},
					cli.StringFlag{
						Name:  "vscode-settings",
						Usage: "Updates the .vscode/settings.json file for redhat yaml extension",
					},
				},
				Usage: "Generates json schema for module3 file.",
				Action: func(c *cli.Context) error {

					source := ""
					if c.IsSet("source") {
						source = c.String("source")
					}
					out := ""
					if c.IsSet("out") {
						out = c.String("out")
					}

					update := ""
					if c.IsSet("vscode-settings") {
						update = c.String("vscode-settings")
					}

					res := GenerateJsonSpecForModule3(source, out, update)

					if out != "" {
						fmt.Println(res)
					}

					return nil
				},
			},
			{
				Name:  "entities",
				Usage: "Lists all of the entities across the binary",

				Action: func(c *cli.Context) error {
					for _, item := range ListModule3WithEntities(xapp) {
						fmt.Println(item)
					}
					return nil
				},
			},
			{
				Name:  "openapi",
				Usage: "Writes the entire app definitions into openapi yml or json file",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "path",
						Usage: "The location that it would write the content, will print it out if left empty",
					},
				},
				Action: func(c *cli.Context) error {
					if content, err := ConvertStructToOpenAPIYaml(xapp); err != nil {
						return err
					} else {
						if strings.HasSuffix(c.String("path"), ".json") {
							jsonData, err := json.MarshalIndent(content, "", "  ")
							if err != nil {
								return err
							}

							os.WriteFile(c.String("path"), jsonData, 0644)
						} else if strings.HasSuffix(c.String("path"), ".yml") || strings.HasSuffix(c.String("path"), ".yaml") {
							yamlData, err := yaml.Marshal(content)
							if err != nil {
								return err
							}

							os.WriteFile(c.String("path"), yamlData, 0644)
						} else {
							jsonData, err := json.MarshalIndent(content, "", "  ")
							if err != nil {
								return err
							}
							fmt.Println(string(jsonData))
						}

						return nil
					}
				},
			},
			{
				Name:  "describe",
				Usage: "Writes a markdown document, explaining entities, actions, tasks, cronjobs - useful for documenting on project management softwares",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "only",
						Usage: "a list of specific modules to be included only, a whitelist of modules as string separated by comma (,)",
					},
				},
				Action: func(c *cli.Context) error {

					ctx := &DescribeContext{}

					if c.IsSet("only") {
						ctx.IncludeOnly = strings.Split(c.String("only"), ",")
					}

					fmt.Print(Describe(xapp, ctx))
					return nil
				},
			},
			{
				Name:  "module-entities",
				Usage: "Lists all of the entities that project has inside a module",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "path",
						Usage:    "Module path, you can get the list using 'list' command",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					for _, item := range ListModule3Entities(xapp, c.String("path")) {
						fmt.Println(item.Name)
					}
					return nil
				},
			},
			{
				Name:  "actions",
				Usage: "Lists all of the available actions over http calls",
				Action: func(c *cli.Context) error {
					for _, item := range xapp.Modules {
						for _, actions := range item.Actions {
							for _, action := range actions {
								fmt.Println(item.Name, action.Url, action.Method)
							}
						}
					}
					return nil
				},
			},
			cli.Command{
				Name: "csv",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "file",
						Usage:    "The address of csv file which will be used to generate",
						Required: true,
					},
				},
				Usage: "Generates Module3 definitions from a csv file, by auto detecting the fields from header",
				Action: func(c *cli.Context) error {

					fields := CastJsonFileToModule3Fields(c.String("file"))
					m2 := &Module3{
						Entities: []Module3Entity{
							{
								Name:   ToCamelCaseClean(c.String("file")),
								Fields: fields,
							},
						},
					}

					fmt.Println(m2.Yaml())

					return nil
				},
			},

			{
				Name:  "postman",
				Usage: "Generates postman collection for all actions in the product (except socket connections)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Where to write the postman collection",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {

					postman := PostmanCollection{
						Info: PostmanInfo{
							Name:   "Fireback Http endpoints",
							Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
						},
						Auth: &PostmanAuth{
							Type: "apikey",
							ApiKey: []PostmanAuthApiKey{
								{
									Key:   "value",
									Value: "{{AUTH}}",
									Type:  "string",
								},
								{
									Key:   "key",
									Value: "authorization",
									Type:  "string",
								},
							},
						},

						Variable: []PostmanVariable{
							{
								Key:   "HOST",
								Value: "localhost",
								Type:  "default",
							},
							{
								Key:         "WID",
								Value:       "root",
								Type:        "default",
								Description: "Current active workspace id",
							},
							{
								Key:         "AUTH",
								Value:       "",
								Type:        "default",
								Description: "Authorization header to authenticate",
							},
							{
								Key:         "RID",
								Value:       "",
								Type:        "default",
								Description: "Selected role which user is working on that behalf",
							},
							{
								Key:   "PORT",
								Value: "4500",
								Type:  "default",
							},
						},
					}
					for _, item := range xapp.Modules {
						for _, actions := range item.Actions {
							for _, action := range actions {

								if strings.TrimSpace(action.Url) == "" {
									continue
								}
								postman.Item = append(postman.Item, PostmanItem{
									Name: action.Url,
									Request: PostmanRequest{
										Method: action.Method,
										Body: PostmanBody{
											Mode: "raw",
											Raw:  action.RequestExample(),
											Options: PostmanBodyOption{
												Raw: PostmanBodyOptionRaw{
													Language: "json",
												},
											},
										},
										Header: []PostmanHeader{
											{
												Key:   "authorization",
												Value: "{{AUTH}}",
											},
											{
												Key:   "role-id",
												Value: "{{RID}}",
											},
											{
												Key:   "workspace-id",
												Value: "{{WID}}",
											},
										},
										Url: PostmanUrl{

											Raw:      "http://{{HOST}}:{{PORT}}" + action.Url,
											Protocol: "http",
											Host: []string{
												"{{HOST}}",
											},

											Port: "{{PORT}}",
											Path: []string{
												action.Url[1:],
											},
										},
									},
								})
							}
						}
					}

					os.WriteFile(c.String("path"), []byte(postman.Json()), 0644)
					return nil
				},
			},
			{
				Flags: commonFlags,
				Name:  "swiftui",
				Usage: "Generates the ios related codes, classes, http calls to build apps easier",
				Action: func(c *cli.Context) error {

					RunCodeGen(xapp, GenContextFromCli(c, SwiftGenCatalog))

					return nil
				},
			},
			{
				Flags: commonFlags,
				Name:  "gof",
				Usage: "Generates the fireback module as golang (backend)",
				Action: func(c *cli.Context) error {
					// gof is a little bit different. We want
					// to generate it's content just next to the module 3 file
					// to allow nested operation

					ctx := GenContextFromCli(c, FirebackGoGenCatalog)

					if len(ctx.ModulesOnDisk) > 0 {
						ctx.Path = path.Dir(ctx.ModulesOnDisk[0])
					}

					if c.IsSet("relative-to") {
						ctx.RelativePath = strings.ReplaceAll(ctx.Path, c.String("relative-to"), "")
						if strings.HasPrefix(ctx.RelativePath, "/") {
							ctx.RelativePath = ctx.RelativePath[1:]
						}
					} else {
						ctx.RelativePath = "not specified"
					}

					ctx.RelativePathDot = strings.ReplaceAll(ctx.RelativePath, "/", ".")

					RunCodeGen(xapp, ctx)

					return nil
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Translation yaml or json entry point",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "langs",
						Usage:    "The languages that you are supporting. Make them attached and separate with comma. for example --langs en,fa,pl,de ",
						Required: false,
					},
				},
				Name:  "strings",
				Usage: "Language resource translation runner",
				Action: func(c *cli.Context) error {
					ctx := TranslationResourceCatalog{
						EntryPoint: c.String("path"),
						Languages:  []string{"en"},
						FileFormat: "yml",
					}

					if c.IsSet("langs") {
						items := strings.Split(c.String("langs"), ",")
						for _, item := range items {
							n := strings.TrimSpace(item)
							n = strings.ToLower(n)

							if n == "en" {
								continue
							}
							ctx.Languages = append(ctx.Languages, n)
						}
					}

					TranslateResource(ctx)

					return nil
				},
			},
			{
				Flags: fbGoModuleFlags,
				Name:  "module",
				Usage: "Generates a new golang (fireback) module into it's own specific directory",
				Action: func(c *cli.Context) error {
					var dirname string
					var moduleName string
					var autoImport string

					if c.IsSet("name") {
						moduleName = c.String("name")
					}

					moduleName = strings.ReplaceAll(moduleName, "\\", "/")
					pathtree := strings.Split(moduleName, "/")

					if c.IsSet("dir") {
						dirname = strings.ToLower(c.String("dir"))
					} else {
						dirname = strings.ToLower(moduleName)
					}

					if len(pathtree) > 1 {
						moduleName = ToUpper(pathtree[len(pathtree)-1])
					}

					if c.IsSet("auto-import") {
						autoImport = c.String("auto-import")
					}

					fmt.Println(4, moduleName)
					return NewGoNativeModule(moduleName, dirname, autoImport)

				},
			},

			{
				Flags: commonFlags,
				Name:  "android",
				Usage: "Generates the android class definitions of the project in Java",
				Action: func(c *cli.Context) error {

					RunCodeGen(xapp, GenContextFromCli(c, JavaGenCatalog))

					return nil
				},
			},
			{
				Flags: commonFlags,
				Name:  "andkot",
				Usage: "Generate Android Kotlin client sdk",
				Action: func(c *cli.Context) error {

					RunCodeGen(xapp, GenContextFromCli(c, KotlinGenCatalog))

					return nil
				},
			},
			{
				Flags: append(commonFlags, reactFlags...),
				Name:  "react",
				Usage: "Generates the typescript definition and react tools for the product",
				Action: func(c *cli.Context) error {

					RunCodeGen(xapp, GenContextFromCli(c, TypeScriptGenCatalog))

					return nil
				},
			},
			{
				Flags: append(commonFlags, reactUIFlags...),
				Name:  "react-ui",
				Usage: "Generates the ui elements for react application, entity manger, form, etc...",
				Action: func(c *cli.Context) error {

					ReactUiCodeGen(xapp, GenContextFromCli(c, TypeScriptGenCatalog), reactui.ReactUITpl)

					return nil
				},
			},
		},
	}
}

package workspaces

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/swaggest/openapi-go/openapi3"
	"github.com/urfave/cli"
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

var commonFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "path",
		Usage:    "Address of the folder, which the content will be generated into",
		Required: true,
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

	ctx := &CodeGenContext{
		Path:          c.String("path"),
		OpenApiFile:   c.String("openapi"),
		Catalog:       cat,
		NoCache:       c.Bool("no-cache"),
		Modules:       strings.Split(c.String("modules"), ","),
		ModulesOnDisk: strings.Split(c.String("def"), ","),
		Ts:            *tsx,
	}

	return ctx
}

func CodeGenTools(xapp *XWebServer) cli.Command {
	return cli.Command{
		Name:  "gen",
		Usage: "Code generation tools, both for internal codes and sdk remote files",
		Subcommands: cli.Commands{
			{
				Name:  "list",
				Usage: "Lists all of the definition modules available in the project",
				Action: func(c *cli.Context) error {
					for _, item := range ListModule2Files(xapp) {
						fmt.Println(item.Path)
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
								fmt.Println(item.Name, action.Url, action.Method, action.ExternFuncName)
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
				Usage: "Generates module 2 definitions from a csv file, by auto detecting the fields from header",
				Action: func(c *cli.Context) error {
					fields := CastJsonFileToModule2Fields(c.String("file"))
					m2 := &Module2{
						Entities: []Module2Entity{
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
				Name:  "csharp",
				Usage: "Unofficial support for C# client",
				Action: func(c *cli.Context) error {
					fmt.Println("Support for this language and related sdk is not available at this product.")
					fmt.Println("To enable, you need to contact Ali Torabi (ali-torabian@outlook.com)")
					return nil
				},
			},
			{
				Name:  "flutter",
				Usage: "Unofficial support for flutter apps",
				Action: func(c *cli.Context) error {
					fmt.Println("Support for this language and related sdk is not available at this product.")
					fmt.Println("To enable, you need to contact Ali Torabi (ali-torabian@outlook.com)")
					return nil
				},
			},
			{
				Name:  "angular",
				Usage: "Unofficial support for Angular 2+ client",
				Action: func(c *cli.Context) error {
					fmt.Println("Support for this language and related sdk is not available at this product.")
					fmt.Println("To enable, you need to contact Ali Torabi (ali-torabian@outlook.com)")
					return nil
				},
			},
			{
				Name:  "qt6",
				Usage: "Unofficial support for QT6 C++ client",
				Action: func(c *cli.Context) error {
					fmt.Println("Support for this language and related sdk is not available at this product.")
					fmt.Println("To enable, you need to contact Ali Torabi (ali-torabian@outlook.com)")
					return nil
				},
			},
			{
				Name:  "svelte",
				Usage: "Unofficial support for svelte framework",
				Action: func(c *cli.Context) error {
					fmt.Println("Support for this language and related sdk is not available at this product.")
					fmt.Println("To enable, you need to contact Ali Torabi (ali-torabian@outlook.com)")
					return nil
				},
			},
			{
				Name:  "vuejs",
				Usage: "Unofficial support for vue framework",
				Action: func(c *cli.Context) error {
					fmt.Println("Support for this language and related sdk is not available at this product.")
					fmt.Println("To enable, you need to contact Ali Torabi (ali-torabian@outlook.com)")
					return nil
				},
			},
			{
				Name:  "oa3-fb",
				Usage: "Converts an open api 3 into fireback definition",
				Flags: append(commonFlags, &cli.StringFlag{
					Name:     "source",
					Usage:    "Where to find the openapi 3 json file",
					Required: true,
				}),
				Action: func(c *cli.Context) error {
					src := c.String("source")

					data, _ := ioutil.ReadFile(src)
					s := openapi3.Spec{}

					if err := s.UnmarshalJSON(data); err != nil {
						log.Fatal("Converting json content:", err)
					}

					app := OpenApiToFireback(s)
					os.WriteFile(c.String("path"), []byte(app.Yaml()), 0644)

					// RunCodeGen(app, GenContextFromCli(c, TypeScriptGenCatalog))

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
								postman.Item = append(postman.Item, PostmanItem{
									Name: action.Url,
									Request: PostmanRequest{
										Method: action.Method,
										Body: PostmanBody{
											Mode: "raw",
											Raw:  "{}",
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

					RunCodeGen(xapp, GenContextFromCli(c, FirebackGoGenCatalog))

					return nil
				},
			},
			{
				Flags: reconfigFlag,
				Name:  "reconfig",
				Usage: "Reconfig the project, usually used for renaming",
				Action: func(c *cli.Context) error {

					dto := ReconfigDto{ProjectSource: "fireback"}

					if c.IsSet("binary-name") {
						dto.BinaryName = c.String("binary-name")
					}
					if c.IsSet("project") {
						dto.NewProjectName = c.String("project")
					}
					if c.IsSet("description") {
						dto.Description = c.String("description")
					}
					if c.IsSet("languages") {
						dto.Languages = strings.Split(c.String("description"), ",")

					}

					return Reconfig(dto)

				},
			},
			{
				Flags: fbGoModuleFlags,
				Name:  "module",
				Usage: "Generates a new golang (fireback) module into it's own specific directory",
				Action: func(c *cli.Context) error {
					var dirname string
					var moduleName string

					if c.IsSet("name") {
						moduleName = c.String("name")
					}

					if c.IsSet("dir") {
						dirname = strings.ToLower(c.String("dir"))
					} else {
						dirname = strings.ToLower(moduleName)
					}

					return NewGoNativeModule(moduleName, dirname)

				},
			},

			{
				Flags: commonFlags,
				Name:  "android",
				Usage: "Generates the android class definitions of the project",
				Action: func(c *cli.Context) error {

					RunCodeGen(xapp, GenContextFromCli(c, JavaGenCatalog))

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
		},
	}
}

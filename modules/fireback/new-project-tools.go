package fireback

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	tmplAndroid "github.com/torabian/fireback/modules/fireback/codegen/android"
	tmplCordova "github.com/torabian/fireback/modules/fireback/codegen/capacitor"
	tmplManage "github.com/torabian/fireback/modules/fireback/codegen/fireback-manage"
	tmpl "github.com/torabian/fireback/modules/fireback/codegen/go-new"
	tmplIos "github.com/torabian/fireback/modules/fireback/codegen/ios"
	tmplReactNative "github.com/torabian/fireback/modules/fireback/codegen/react-native-new"
	tmplReact "github.com/torabian/fireback/modules/fireback/codegen/react-new"
	"github.com/urfave/cli"
)

func newProjectContentWriter(fsys embed.FS, ctx *NewProjectContext, prefix string) {

	// Walk through the embedded filesystem
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// This index.go is not needed
		if path == "index.go" {
			return nil
		}

		// Create the corresponding destination path
		destPath := filepath.Join(ctx.Path, prefix, path)

		destPath = strings.ReplaceAll(destPath, "projectname", ctx.Name)

		// Check if the entry is a directory
		if d.IsDir() {
			// Create the directory if it does not exist
			return os.MkdirAll(destPath, os.ModePerm)
		}

		// If it's a file, read its contents
		content, err := fsys.ReadFile(path)
		if err != nil {
			return err
		}

		c := string(content)
		content = []byte(strings.ReplaceAll(c, "projectname", ctx.Name))

		// Check file extension
		if filepath.Ext(path) == ".tpl" {
			t, err := template.New("").Funcs(CommonMap).Parse(c)
			if err != nil {
				return err
			}

			file, err := os.Create(strings.ReplaceAll(destPath, ".tpl", ""))
			if err != nil {
				return err
			}
			defer file.Close()

			err = t.Execute(file, gin.H{
				"ctx": ctx,
			})
			if err != nil {
				return err
			}

		} else {
			return os.WriteFile(destPath, content, os.ModePerm)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func copyFile(src string, dst string) {
	// Read all content of src to data, may cause OOM for a large file.
	data, _ := ioutil.ReadFile(src)

	ioutil.WriteFile(dst, data, 0644)

}

type NewProjectContext struct {
	Name                     string
	Path                     string
	IsMonolith               bool
	FirebackManage           bool
	ModuleName               string
	ReplaceFireback          string
	Description              string
	FirebackVersion          string
	CreateReactProject       bool
	SelfService              bool
	CreateIOSProject         bool
	CreateAndroidProject     bool
	CreateCapacitorProject   bool
	CreateReactNativeProject bool
}

func NewProjectCli() cli.Command {
	return cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "Name of the project that you want to create.",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "path",
				Usage:    "The directory that new project will be created. If not entered, project name will be used",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "micro",
				Usage:    "If the new project is a micro service - default is false, and we create monolith",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "ui",
				Usage:    "If you set --ui true, there will be a front-end project also added.",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "mobile",
				Usage:    "If you set --mobile true, there will be a application project also added for react native",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "android",
				Usage:    "If you set --android true, native android project in java will be generated",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "ios",
				Usage:    "If you set --ios true, native ios project using swiftui for xcode will be generated",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "manage",
				Usage:    "Adds the prebuilt react.js app 'manage' as a separate admin panel for the project",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "capacitor",
				Usage:    "If you set --capacitor true, fireback adds capacitor project compatible with front-end (react)",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "description",
				Usage:    "Description of the project which would appear in few places",
				Required: false,
				Value:    "Backend project built by Fireback",
			},
			&cli.StringFlag{
				Name:     "module",
				Usage:    "Module name of the go.mod - project comes with go modules. for example --module github.com/you/project",
				Required: false,
			},
			&cli.StringFlag{
				Name:  "replace-fb",
				Usage: "Replace fireback with a local copy inside or outside project (if you ejected the fireback)",
			},
		},
		Name:  "new",
		Usage: "Generate a new fireback project or microservice.",
		Action: func(c *cli.Context) error {
			ctx := &NewProjectContext{
				FirebackVersion: FIREBACK_VERSION,
				IsMonolith:      false,
			}

			if c.NumFlags() == 0 {
				ctx.Name = AskForInput("Give the project a name", "newapp")
				ctx.ModuleName = AskForInput("What is the golang module name?", "github.com/torabian/testapp")
				if r := AskForSelect("Architecture type of the project?", []string{"monolith", "microservice"}); r == "monolith" {
					ctx.IsMonolith = true
				}

				if r := AskForSelect("Do you want to use a local copy of fireback instead of mod?", []string{"yes", "no"}); r == "yes" {
					for {
						ctx.ReplaceFireback = AskForInput("Set the fireback relative folder address", "../fireback")
						if !Exists(filepath.Join(ctx.ReplaceFireback, "modules", "workspaces")) {
							if r := AskForSelect("Fireback not found in directory. Try again?", []string{"yes", "no"}); r == "no" {
								break
							}
						} else {
							break
						}
					}
				}

				ctx.Path = AskForInput("Name of the folder which will be generated", ctx.Name)
				ctx.Description = AskForInput("Any description for the project created", ctx.Description)

				if r := AskForSelect("Do you want to have front-end project in react.js?", []string{"yes", "no"}); r == "yes" {
					ctx.CreateReactProject = true

					if r := AskForSelect("Do you want to have capacitor created?", []string{"yes", "no"}); r == "yes" {
						ctx.CreateCapacitorProject = true
					}
				}

				if r := AskForSelect("Do you want to have fireback 'manage' admin panel react.js built independent in project?", []string{"yes", "no"}); r == "yes" {
					ctx.FirebackManage = true
				}

				if r := AskForSelect("Do you want to add fireback self-service ui as well into your project?", []string{"yes", "no"}); r == "yes" {
					ctx.SelfService = true
				}

				// These are not complete boilerplates, I am not deleting them
				// because over time I'll do to it, with small priority.
				// if r := AskForSelect("Do you want to have react native project?", []string{"no", "yes"}); r == "yes" {
				// 	ctx.CreateReactNativeProject = true
				// }

				// if r := AskForSelect("Do you want to have native ios boilerplate project?", []string{"no", "yes"}); r == "yes" {
				// 	ctx.CreateIOSProject = true
				// }
				// if r := AskForSelect("Do you want to have native android java boilerplate project?", []string{"no", "yes"}); r == "yes" {
				// 	ctx.CreateAndroidProject = true
				// }

			} else {
				pathd := c.String("path")
				if pathd == "" {
					pathd = c.String("name")
				}
				ctx = &NewProjectContext{
					FirebackVersion:          FIREBACK_VERSION,
					IsMonolith:               true,
					Name:                     c.String("name"),
					Description:              c.String("description"),
					ReplaceFireback:          c.String("replace-fb"),
					Path:                     pathd,
					ModuleName:               c.String("module"),
					CreateReactProject:       c.Bool("ui"),
					SelfService:              c.Bool("self-service"),
					FirebackManage:           c.Bool("manage"),
					CreateReactNativeProject: c.Bool("mobile"),
					CreateCapacitorProject:   c.Bool("capacitor"),
					CreateIOSProject:         c.Bool("ios"),
					CreateAndroidProject:     c.Bool("android"),
				}

				if c.IsSet("micro") {
					ctx.IsMonolith = !c.Bool("micro")
				}
			}

			newProjectContentWriter(tmpl.FbGoNewTemplate, ctx, "")

			if ctx.CreateReactProject {
				newProjectContentWriter(tmplReact.FbReactjsNewTemplate, ctx, "front-end")
				source := filepath.Join(ctx.Path, "front-end", "src/apps", ctx.Name, ".env.local.txt")
				uibuilt := filepath.Join(ctx.Path, "cmd", ctx.Name+"-server", "ui")
				os.MkdirAll(uibuilt, os.ModePerm)
				os.WriteFile(path.Join(uibuilt, ".gitkeep"), []byte("UI built will replace this"), 0644)
				os.WriteFile(path.Join(uibuilt, "index.html"), []byte(`<h1>UI Placeholder</h1>
<p>
  Welcome to the Fireback project. This folder can contains your html/css
  project. If you want to serve the your static project from react/angular, etc,
  make an script that would clear the content of this ui folder and replace it
  with those files.
</p>

<p>
  In Golang, you can actually build your html pages in old fasion server side
  rendering MVC, without any extra ui library, if that's what you want.
</p>

<p>
  Depending on your configuration when creating project, you might be able to
  access 'manage' and 'selfservice' already here:
</p>

<ul>
  <li>
    <a href="/selfservice"
      >Selfservice: For login and get token, and maybe redirect or postMessage
      to your own app</a
    >
  </li>
  <li><a href="/manage">Fireback management UI called 'manage'</a></li>
  <li><a href="/docs">OpenAPI spec</a></li>
</ul>

<p>
  Fireback doesn't have any special rule regarding front-end, you can build your
  own on your own stack, or even remove this folder, and just use it as API
  server.
</p>
`), 0644)
				dest := filepath.Join(ctx.Path, "front-end", "src/apps", ctx.Name, ".env.local")
				copyFile(source, dest)

			}

			if ctx.CreateReactNativeProject {
				newProjectContentWriter(tmplReactNative.FbReactNativeNewTemplate, ctx, "mobile")
			}

			if ctx.CreateIOSProject {
				newProjectContentWriter(tmplIos.IosProjectTmpl, ctx, "ios")
			}

			if ctx.FirebackManage {
				newProjectContentWriter(tmplManage.FirebackManageTmpl, ctx, path.Join("cmd", ctx.Name+"-server", "manage"))
			}

			if ctx.CreateAndroidProject {
				newProjectContentWriter(tmplAndroid.AndroidProjectTmpl, ctx, "android")
			}

			if ctx.CreateCapacitorProject {
				newProjectContentWriter(tmplCordova.FbReactCapacitorNewTemplate, ctx, "capacitor")
			}

			return nil
		},
	}
}

package workspaces

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	tmpl "github.com/torabian/fireback/modules/workspaces/codegen/go-new"
	tmplReactNative "github.com/torabian/fireback/modules/workspaces/codegen/react-native-new"
	tmplReact "github.com/torabian/fireback/modules/workspaces/codegen/react-new"
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

type NewProjectContext struct {
	Name            string
	Path            string
	IsMonolith      bool
	ModuleName      string
	ReplaceFireback string
	Description     string
	FirebackVersion string
}

func NewProjectCli() cli.Command {
	return cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "Name of the project that you want to create.",
				Required: true,
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
			&cli.StringFlag{
				Name:     "description",
				Usage:    "Description of the project which would appear in few places",
				Required: false,
				Value:    "Backend project built by Fireback",
			},
			&cli.StringFlag{
				Name:     "module",
				Usage:    "Module name of the go.mod - project comes with go modules. for example --module github.com/you/project",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "replace-fb",
				Usage: "Replace fireback with a local copy inside or outside project (if you ejected the fireback)",
			},
		},
		Name:  "new",
		Usage: "Generate a new fireback project or microservice.",
		Action: func(c *cli.Context) error {
			pathd := c.String("path")
			if pathd == "" {
				pathd = c.String("name")
			}
			ctx := &NewProjectContext{
				Name:            c.String("name"),
				Description:     c.String("description"),
				ReplaceFireback: c.String("replace-fb"),
				Path:            pathd,
				IsMonolith:      true,
				FirebackVersion: FIREBACK_VERSION,
				ModuleName:      c.String("module"),
			}

			if c.IsSet("micro") {
				ctx.IsMonolith = !c.Bool("micro")
			}

			newProjectContentWriter(tmpl.FbGoNewTemplate, ctx, "")

			if c.Bool("ui") {
				newProjectContentWriter(tmplReact.FbReactjsNewTemplate, ctx, "front-end")
				source := filepath.Join(ctx.Path, "front-end", "src/apps", ctx.Name, ".env.local.txt")
				dest := filepath.Join(ctx.Path, "front-end", "src/apps", ctx.Name, ".env.local")
				copyFile(source, dest)
			}

			if c.Bool("mobile") {
				newProjectContentWriter(tmplReactNative.FbReactNativeNewTemplate, ctx, "mobile")
				// source := filepath.Join(ctx.Path, "front-end", "src/apps", ctx.Name, ".env.local.txt")
				// dest := filepath.Join(ctx.Path, "front-end", "src/apps", ctx.Name, ".env.local")
				// copyFile(source, dest)
			}

			return nil
		},
	}
}

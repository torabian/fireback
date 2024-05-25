package workspaces

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	tmpl "github.com/torabian/fireback/modules/workspaces/codegen/go-new"
	"github.com/urfave/cli"
)

func newProjectContentWriter(ctx *NewProjectContext) {

	// Walk through the embedded filesystem
	err := fs.WalkDir(tmpl.FbGoNewTemplate, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// This index.go is not needed
		if path == "index.go" {
			return nil
		}

		// Create the corresponding destination path
		destPath := filepath.Join(ctx.Path, path)

		destPath = strings.ReplaceAll(destPath, "projectname", ctx.Name)

		// Check if the entry is a directory
		if d.IsDir() {
			// Create the directory if it does not exist
			return os.MkdirAll(destPath, os.ModePerm)
		}

		// If it's a file, read its contents
		content, err := tmpl.FbGoNewTemplate.ReadFile(path)
		if err != nil {
			return err
		}

		// Check file extension
		if filepath.Ext(path) == ".tpl" {

			t, err := template.New("").Funcs(CommonMap).Parse(string(content))
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
	ModuleName      string
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
			&cli.StringFlag{
				Name:     "description",
				Usage:    "Description of the project which would appear in few places",
				Required: false,
				Value:    "Backend project built by Fireback",
			},
			&cli.StringFlag{
				Name:     "moduleName",
				Usage:    "Module name of the go.mod - project comes with go modules. for example --moduleName github.com/you/project",
				Required: true,
			},
		},
		Name:  "new",
		Usage: "Generate a new fireback project.",
		Action: func(c *cli.Context) error {
			pathd := c.String("path")
			if pathd == "" {
				pathd = c.String("name")
			}
			ctx := &NewProjectContext{
				Name:            c.String("name"),
				Description:     c.String("description"),
				Path:            pathd,
				FirebackVersion: FIREBACK_VERSION,
				ModuleName:      c.String("moduleName"),
			}
			newProjectContentWriter(ctx)
			return nil
		},
	}
}

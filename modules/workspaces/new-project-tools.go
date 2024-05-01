package workspaces

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	tmpl "github.com/torabian/fireback/modules/workspaces/codegen/go-new"
	"github.com/urfave/cli"
)

func newProjectContentWriter(destDir string) {

	// Walk through the embedded filesystem
	err := fs.WalkDir(tmpl.FbGoNewTemplate, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Create the corresponding destination path
		destPath := filepath.Join(destDir, path)

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
			// If it's a .tpl file, render it as a template
			tmpl, err := template.New("").Parse(string(content))
			if err != nil {
				return err
			}

			// Create the file
			file, err := os.Create(strings.ReplaceAll(destPath, ".tpl", ""))
			if err != nil {
				return err
			}
			defer file.Close()

			// Execute the template and write to the file
			return tmpl.Execute(file, nil)
		}

		// If it's not a .tpl file, copy it to the destination
		return os.WriteFile(destPath, content, os.ModePerm)
	})
	if err != nil {
		panic(err)
	}
}

func NewProjectCli() cli.Command {
	return cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Usage:    "Name of the project that you want to create.",
				Required: true,
			},
		},
		Name:  "new",
		Usage: "Generate a new fireback project.",
		Action: func(c *cli.Context) error {
			newProjectContentWriter(c.String("name"))
			return nil
		},
	}
}

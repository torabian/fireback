package abac

// Css is a good tool for building UI elements. This file gives such tools to the applications

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/urfave/cli"
)

func resolveCssImports(path string, visited map[string]bool) (string, error) {
	if visited[path] {
		return "", nil // avoid circular imports
	}
	visited[path] = true

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(path)
	re := regexp.MustCompile(`@import\s+(url\()?["']?([^"')]+)["']?\)?[^;]*;`)
	content := re.ReplaceAllStringFunc(string(data), func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 3 {
			return match // skip malformed
		}

		importPath := filepath.Join(dir, matches[2])
		imported, err := resolveCssImports(importPath, visited)
		if err != nil {
			return "\r\n/* failed: " + matches[2] + "*/\r\n"
		}
		imported = "\r\n/*" + matches[2] + "*/\r\n" + imported
		return imported
	})

	return content, nil
}

func combineAndMinifyCssRecursively(sourceFilePath string, outputFilePath string, skipMinify bool) {
	cssContent, err := resolveCssImports(sourceFilePath, map[string]bool{})
	if err != nil {
		panic(err)
	}

	if skipMinify != true {

		m := minify.New()
		m.AddFunc("text/css", css.Minify)

		minified, err := m.String("text/css", cssContent)
		if err != nil {
			panic(err)
		}

		cssContent = minified
	}

	if outputFilePath == "" {
		fmt.Print(cssContent)
		return
	}

	err = ioutil.WriteFile(outputFilePath, []byte(cssContent), 0644)
	if err != nil {
		panic(err)
	}
}

func getCssMinCombineCli() cli.Command {
	return cli.Command{
		Name:        "cssx",
		Description: "Minifies css file, and resolves the @import dependencies recursively",
		Usage:       `Minifies css file, and resolves the @import dependencies recursively`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "entry-point",
				Value:    "",
				Usage:    "Absolute or relative path of the css entry point file on the disk",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "out",
				Value: "",
				Usage: "output file address. If left empty, it will write result to stdout",
			},
			&cli.BoolFlag{
				Name:  "skip-minify",
				Usage: "If true, would skip minifying the css result file",
			},
		},
		Action: func(c *cli.Context) error {
			combineAndMinifyCssRecursively(c.String("entry-point"), c.String("out"), c.Bool("skip-minify"))
			return nil
		},
	}
}

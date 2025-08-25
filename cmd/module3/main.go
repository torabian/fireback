package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/manifoldco/promptui"

	"github.com/torabian/fireback/modules/fireback/module3/m3angular"
	"github.com/torabian/fireback/modules/fireback/module3/m3js"
	"github.com/torabian/fireback/modules/fireback/module3/mcore"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func main() {
	app := &cli.App{
		Name:  "module3-cli",
		Usage: "Module3 definitions code generator",
		Commands: []cli.Command{

			GetCliActionForActionGen(
				"js:action:headers",
				"Generate the javascript class for an action headers extending Headers class",
				m3js.JsActionHeaderClass,
			),
			GetCliActionForActionGen(
				"js:action",
				"Combines multiple code gens which would make an action typesafe in javascript",
				m3js.JsActionClass,
			),

			GetCliActionForActionGen(
				"js:action:qs",
				"Generate javascript query string class library",
				m3js.JsActionQsClass,
			),
			GetCliActionForCompleteModule("js:module", "Compiles the entire javascript modules and writes them to disk", m3js.JsModuleFullVirtualFiles),
			GetCliActionForCompleteModule("angular:module", "Compiles the angular + javascript (typescript) required files to use the sdk", m3angular.AngularModuleFullVirtualFiles),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

type TypeScriptGenContext struct {
	IncludeStaticField      bool
	IncludeFirebackDef      bool
	IncludeStaticNavigation bool
}

func GenContextFromCli(c *cli.Context) *CodeGenContext {
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
		Path:    c.String("path"),
		NoCache: c.Bool("no-cache"),
	}

	return ctx
}

type CodeGenContext struct {
	// Where the content will be exported to
	Path string

	// Used in golang which indicates the relative path
	RelativePath    string
	RelativePathDot string

	// Only build specific modules
	ModulesOnDisk []string

	NoCache bool
}

func GetCliActionForCompleteModule(name string, description string, fn mcore.CompleteModuleGenerator) cli.Command {
	return cli.Command{
		Name:        name,
		Description: description,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "path",
				Usage:    "Module3 yaml definition file location for the import",
				Required: true,
			},
			cli.StringFlag{
				Name:  "output",
				Usage: "The directory which the generated files will be rewritten to",
			},

			cli.StringFlag{
				Name:  "tags",
				Usage: "A set of string flags separated by comma (,) to add or remove compile feature. Such as 'nestjs-headers-decorator'",
			},
		},
		Action: func(c *cli.Context) error {
			ctx := mcore.MicroGenContext{
				Tags:   c.String("tags"),
				Output: c.String("output"),
			}

			data, err := ReadModule3FromFile(c.String("path"))
			if err != nil {
				return err
			}

			// Let's combine the import requirements of the chunk
			files, err := fn(data, ctx)
			if err != nil {
				return err
			}

			// If there is no output, we just write the content as json into the output
			if ctx.Output == "" {
				res, _ := json.MarshalIndent(files, "", "  ")
				fmt.Println(string(res))

				return nil
			}

			// @todo bring back the caching mechanism into here from fireback
			os.MkdirAll(ctx.Output, os.ModePerm)

			for _, file := range files {
				filePath := path.Join(ctx.Output, file.Location, file.Name+file.Extension)

				fmt.Println("filePath", filePath)
				if err := os.WriteFile(filePath, []byte(file.ActualScript), 0644); err != nil {
					return fmt.Errorf("error on writing file to disk: %v, %v, %w", file.Location, file.Name, err)
				}
			}

			return nil
		},
	}
}

func GetCliActionForActionGen(name string, description string, call mcore.ActionCodeGenerator) cli.Command {
	return cli.Command{
		Name:        name,
		Description: description,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "path",
				Usage:    "Module3 yaml definition file location for the import",
				Required: true,
			},
			cli.StringFlag{
				Name:  "out",
				Usage: "If set, it would write the output into a file",
			},

			cli.StringFlag{
				Name:  "tags",
				Usage: "A set of string flags separated by comma (,) to add or remove compile feature. Such as 'nestjs-headers-decorator'",
			},
		},
		Action: func(c *cli.Context) error {
			action, err := PickActionFromModule3(c.String("path"))
			if err != nil {
				return err
			}

			ctx := mcore.MicroGenContext{
				Tags: c.String("tags"),
			}

			result, err := call(action, ctx)
			if err != nil {
				return err
			}

			// Let's combine the import requirements of the chunk
			importsList := m3js.CombineImportsJsWorld(*result)
			var finalContent string = importsList + "\r\n" + string(result.ActualScript)

			if c.IsSet("out") {
				return os.WriteFile(c.String("out"), []byte(finalContent), 0644)
			} else {
				fmt.Println(string(finalContent))
			}

			return nil
		},
	}
}

func ReadModule3FromFile(source string) (*mcore.Module3, error) {
	content, err := os.ReadFile(source)
	if err != nil {
		return nil, err
	}

	var data mcore.Module3
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func PickActionFromModule3(source string) (*mcore.Module3Action, error) {
	data, err := ReadModule3FromFile(source)
	if err != nil {
		return nil, err
	}

	var index int64 = 0
	items := data.ActionsAsList()

	if len(items) > 1 {
		index0 := AskForSelect("Select the action to generate header out of it.", items)
		index, err = ToInt64(index0)
		if err != nil {
			return nil, err
		}

	}

	action := data.Actions[index]

	if action == nil {
		return nil, errors.New("Action is not readable")
	}

	return action, nil
}

func AskForSelect(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return ""
		}
		return ""
	}

	index := strings.Index(result, ">>>")
	if index <= 0 {
		return result
	}
	return strings.Trim(result[0:index], " ")

}

// ToInt64 casts a string to int64 if it's numeric
func ToInt64(s string) (int64, error) {
	if IsNumeric(s) {
		return strconv.ParseInt(s, 10, 64)
	}
	return 0, fmt.Errorf("not a valid number")
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

package workspaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var CommonQueryFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "verbose",
		Usage: "Verbose query, show fireback columns as well such as workspace, etc",
	},
	&cli.BoolFlag{
		Name:  "minimal",
		Usage: "Make a minimal query, skips printing some of the fields",
	},
	&cli.IntFlag{
		Name:  "offset",
		Usage: "Add the start index",
		Value: 0,
	},
	&cli.IntFlag{
		Name:  "limit",
		Usage: "Items per page",
		Value: 0,
	},
	&cli.BoolFlag{
		Name:  "yaml",
		Usage: "Make result as yaml file",
	},
	&cli.StringFlag{
		Name:  "sort",
		Usage: "Sorting strategy",
	},
	&cli.BoolFlag{
		Name:  "deep",
		Usage: "Should query the arrays, nested objects, relations of the entity",
	},
	&cli.StringFlag{
		Name:  "query",
		Usage: "Query DSL which filters out the results.",
		Value: "",
	},
	&cli.StringFlag{
		Name:  "wp",
		Usage: "withPreloads The sub or nested entities to be loaded with. Comma separated",
		Value: "",
	},
	&cli.StringFlag{
		Name:  "lang",
		Usage: "define the language in 2 char code, aka: en, de",
		Value: "en",
	},
}

func GetCommonRemoveQuery(el reflect.Value, fn ActionDeleteSignature) cli.Command {

	return cli.Command{

		Name:    "remove",
		Aliases: []string{"r", "del", "delete"},
		Usage:   "Deletes an entity with given id (uniqueid)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Value:    "",
				Usage:    "id of the record to be deleted",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {

			f := QueryDSL{
				UniqueId: c.String("id"),
				Query:    "unique_id = " + c.String("id"),
			}
			fmt.Println(fn(f))
			return nil

		},
	}

}

func GetCommonQuery[T any](fn func(query QueryDSL) ([]*T, *QueryResultMeta, error)) cli.Command {

	return cli.Command{

		Name:    "query",
		Aliases: []string{"q"},
		Flags:   CommonQueryFlags,
		Usage:   "Queries all of the entities in database based on the standard query format",
		Action: func(c *cli.Context) error {
			CommonCliQueryCmd(
				c,
				fn,
			)

			return nil
		},
	}

}

// This is with security
func GetCommonQuery2[T any](
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, error),
	security *SecurityModel,
) cli.Command {

	return cli.Command{

		Name:    "query",
		Aliases: []string{"q"},
		Flags:   CommonQueryFlags,
		Usage:   "Queries all of the entities in database based on the standard query format (s+)",
		Action: func(c *cli.Context) error {
			CommonCliQueryCmd2(
				c,
				fn,
				security,
			)

			return nil
		},
	}

}

func GetCommonCteQuery[T any](fn func(query QueryDSL) ([]*T, *QueryResultMeta, error)) cli.Command {

	return cli.Command{

		Name:    "query-cte",
		Aliases: []string{"cte"},
		Flags:   CommonQueryFlags,
		Usage:   "Same as query, but in recursive manner",
		Action: func(c *cli.Context) error {
			CommonCliQueryCmd(
				c,
				fn,
			)

			return nil
		},
	}

}

func GetCommonExtendedQuery[T any](fn func(query QueryDSL) ([]*T, *QueryResultMeta, error)) cli.Command {

	return cli.Command{

		Name:    "query-extended",
		Aliases: []string{"extended"},
		Flags:   CommonQueryFlags,
		Usage:   "Extended query, provides way more details, and combines the one-to-many hirechical relations",
		Action: func(c *cli.Context) error {
			CommonCliQueryCmd(
				c,
				fn,
			)

			return nil
		},
	}

}

func ManifestTools() cli.Command {

	return cli.Command{

		Name:  "manifest",
		Usage: "Tools related to the manifest definitions",
		Subcommands: []cli.Command{
			{
				Name:  "compile",
				Usage: "Compiles the entire bundles, styles, etc related to a manifest",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Usage:    "The address the manifest file",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					CompileMvxManifest(c.String("file"))
					return nil
				},
			},
		},
	}

}

func GetCommonPivotQuery[T any](fn func(query QueryDSL) ([]*T, *QueryResultMeta, error)) cli.Command {

	return cli.Command{

		Name:    "query-pivot",
		Aliases: []string{"pivot"},
		Flags:   CommonQueryFlags,
		Usage:   "Pivots the the entire table based on conditions",
		Action: func(c *cli.Context) error {
			CommonCliQueryCmd(
				c,
				fn,
			)

			return nil
		},
	}

}

type CliInteractiveFlag struct {
	Name        string
	StructField string
	// It is required on database level
	Required bool

	// Its recommended on the cli level to make it easier.
	Recommended bool
	Usage       string
	Type        string
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

	return result
}

func AskForInputOptional(label string, defaultV string) (string, bool, error) {

	promptVariable := promptui.Prompt{
		Label:   label,
		Default: defaultV,
	}

	value, err := promptVariable.Run()
	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return "", true, err
		}
		return "", false, err
	}

	return value, false, nil
}

func AskForInput(label string, defaultV string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("this is necessary")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Default:  defaultV,
	}

	value, err := promptVariable.Run()
	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return ""
		}
		return ""
	}

	return value
}

func AskForPassword(label string, defaultV string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("this is necessary")
		}
		return nil
	}

	promptVariable := promptui.Prompt{
		Label:    label,
		Mask:     '*',
		Validate: validate,
		Default:  defaultV,
	}

	value, err := promptVariable.Run()
	if err != nil {
		if err.Error() == "^C" {
			os.Exit(35)
			return ""
		}
		return ""
	}
	return value
}

func TranslateIError(err *IError, translateDictionary map[string]map[string]string, targetLanguage string) {
	if err == nil {
		return
	}
	// if err.Message != "" && translateDictionary[err.Message][targetLanguage] != "" {
	// 	err.MessageTranslated = translateDictionary[err.Message][targetLanguage]
	// }

	// for _, errItem := range err.Errors {
	// 	// Some fields are having params, so we detect them and translate them appropriately

	// 	// min=1 means that field is required, and empty string is not accepted
	// 	if errItem.Message == "min" && errItem.ErrorParam == "1" {
	// 		errItem.Message = "required"
	// 	}

	// 	if errItem.Message != "" && translateDictionary[errItem.Message][targetLanguage] != "" {

	// 		errItem.MessageTranslated = translateDictionary[errItem.Message][targetLanguage]
	// 	}
	// }
}

func HandleActionInCli(c *cli.Context, result any, err *IError, t map[string]map[string]string) {
	f := CommonCliQueryDSLBuilder(c)

	if result != nil {
		body, _ := yaml.Marshal(result)
		fmt.Printf("%-20s:\n", string(body)) // Left-align keys, pad to 20 spaces
	}

	if err != nil {
		err2 := err.ToPublicEndUser(&f)

		if err2 == nil {
			log.Panicln("Panic on handle action, without public error: %w", err)
			return
		}

		fmt.Println("Error HttpCode:", err2.HttpCode)
		fmt.Println("Error Message:", err2.Message, err.MessageTranslated)
		for index, errItem := range err2.Errors {
			fmt.Println(index, ":",
				errItem.Message, "on", errItem.Location,
				// errItem.MessageTranslated,
			)
		}
		os.Exit(int(err2.HttpCode))
	}

}

func CommonInitSeeder[T any](format string, entity *T) {
	body := []byte{}
	var err error
	data := []*T{}
	data = append(data, entity)

	if format == "" {
		format = "yml"
	}

	if format == "yml" || format == "yaml" {
		body, err = yaml.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
	}

	if format == "json" {
		body, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println(string(body))
}

// Use the actions bundle for ease and provide it to the ModuleProvider
// and it would gather all actions in the module level, it's to make it easier
// for intelisense
type CliActionsBundle = cli.Command

// Represents both http, and cli actions in one single object
type ModuleActionsBundle struct {

	// cli.Command which has Subcommands of all actions
	CliAction *cli.Command

	// these are actions used for web generally, more general. In ideal world these
	// could be used to create cli actions
	Actions []Module2Action
}

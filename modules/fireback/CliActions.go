package fireback

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v2"
)

// Fireback actions run in cli as well. In this file, we place tools and helpers for that.

var CommonQueryFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "verbose",
		Usage: "Verbose query, show fireback columns as well such as workspace, etc",
	},
	&cli.StringFlag{
		Name:     "x-select",
		Required: false,
		Usage:    `Select only specific fields to be queried and returned`,
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
	&cli.StringFlag{
		Name:  "cursor",
		Usage: "Cursor value from the pagination",
	},
	&cli.IntFlag{
		Name:  "limit",
		Usage: "Items per page",
		Value: 0,
	},
	&cli.StringFlag{
		Name:  "x-accept",
		Usage: "Change the return results to such as as 'yaml'",
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

func GetCommonRemoveQuery(el reflect.Value, fn ActionDeleteSignature) *cli.Command {

	return &cli.Command{

		Name:    "remove",
		Aliases: []string{"r", "del", "delete"},
		Usage:   "Deletes a given record using it's unique_id string hash field from database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "uid",
				Value:    "",
				Usage:    "String unique id of the record which will be deleted.",
				Required: true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {

			f := QueryDSL{
				UniqueId: c.String("uid"),
			}

			res, err := fn(DeleteRequest{
				Query: "unique_id = " + c.String("id"),
			}, f)

			if err != nil {
				fmt.Println("Delete operation encountered an error on database level: %w", err)
				return err
			}

			if res.Data.Item.RowsAffected == 0 {
				fmt.Println("Query executed successfully, but no Rows have been deleted.")
			} else {
				fmt.Printf("Deleted rows by delete operation: %d\r\n", res.Data.Item.RowsAffected)
			}

			return nil

		},
	}

}

func GetCommonCteQuery[T any](fn func(query QueryDSL) ([]*T, *QueryResultMeta, *IError)) *cli.Command {

	return &cli.Command{

		Name:    "query-cte",
		Aliases: []string{"cte"},
		Flags:   CommonQueryFlags,
		Usage:   "Same as query, but in recursive manner",
		Action: func(ctx context.Context, c *cli.Command) error {
			CommonCliQueryCmd3(
				c,
				fn,
				nil,
				nil,
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

// AskForSelect, AskBoolean, AskForInputOptional, AskForInput and AskForPassword
// are interactive terminal prompts. Their real implementation lives in the
// modules/fireback/terminal package (tagged !wasm) and registers itself here
// via init(). Building without importing that package (e.g. cmd/fireback-wasm)
// leaves these nil - callers only reach them from CLI flows that aren't part
// of a wasm build.
var AskForSelect func(label string, items []string) string
var AskBoolean func(label string) bool
var AskForInputOptional func(label string, defaultV string) (string, bool, error)
var AskForInput func(label string, defaultV string) string
var AskForPassword func(label string, defaultV string) string

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
}

type QuerySelectionInfo struct {
	Columns  []string
	Preloads []string
}

func (x QuerySelectionInfo) Json() string {

	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))

}

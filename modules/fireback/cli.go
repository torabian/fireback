package fireback

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

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

func GetCommonCteQuery[T any](fn func(query QueryDSL) ([]*T, *QueryResultMeta, *IError)) cli.Command {

	return cli.Command{

		Name:    "query-cte",
		Aliases: []string{"cte"},
		Flags:   CommonQueryFlags,
		Usage:   "Same as query, but in recursive manner",
		Action: func(c *cli.Context) error {
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

func GetCommonExtendedQuery[T any](fn func(query QueryDSL) ([]*T, *QueryResultMeta, *IError)) cli.Command {

	return cli.Command{

		Name:    "query-extended",
		Aliases: []string{"extended"},
		Flags:   CommonQueryFlags,
		Usage:   "Extended query, provides way more details, and combines the one-to-many hirechical relations",
		Action: func(c *cli.Context) error {
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

func GetCommonPivotQuery[T any](fn func(query QueryDSL) ([]*T, *QueryResultMeta, *IError)) cli.Command {

	return cli.Command{

		Name:    "query-pivot",
		Aliases: []string{"pivot"},
		Flags:   CommonQueryFlags,
		Usage:   "Pivots the the entire table based on conditions",
		Action: func(c *cli.Context) error {
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

func AskBoolean(label string) bool {
	if r := AskForSelect(label, []string{"true", "false"}); r == "true" {
		return true
	}

	return false
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

func HandleActionInCli(c *cli.Context, result any, err *IError, t map[string]map[string]string) {
	f := CommonCliQueryDSLBuilder(c)

	resultIsNil := result == nil || (reflect.ValueOf(result).Kind() == reflect.Ptr && reflect.ValueOf(result).IsNil())

	if !resultIsNil {
		cliSuccessPrinter(c, result)
	}

	if err != nil {

		err2 := err.ToPublicEndUser(&f)

		if err2 == nil {
			log.Panicln("Panic on handle action, without public error: %w", err)
			return
		}

		body, _ := json.MarshalIndent(err2, "", "  ")
		fmt.Println(string(body))

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
	Actions []Module3Action
}

func populateQueriableFields(v reflect.Value, parent string, getValue func(field reflect.StructField) string) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := parent + field.Name
		fieldValue := v.Field(i)

		// Recursively handle embedded/nested structs
		if field.Type.Kind() == reflect.Struct {
			populateQueriableFields(fieldValue, name+".", getValue)
		}

		if fieldValue.CanSet() && fieldValue.Type() == reflect.TypeOf(QueriableField{}) {
			current := fieldValue.Interface().(QueriableField)
			current.UserInput = getValue(field)
			fieldValue.Set(reflect.ValueOf(current))
		}
	}
}

func QueriableFieldFromCliContext(v reflect.Value, parent string, c *cli.Context) {
	populateQueriableFields(v, parent, func(field reflect.StructField) string {
		return c.String(field.Tag.Get("cli"))
	})
}

func QueriableFieldFromGinContext(v reflect.Value, parent string, c *gin.Context) {
	populateQueriableFields(v, parent, func(field reflect.StructField) string {
		val, _ := c.GetQuery(field.Tag.Get("qs"))
		return val
	})
}

type QuerySelectionInfo struct {
	Columns  []string
	Preloads []string
}

func (x QuerySelectionInfo) Json() string {

	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))

}

func DatabaseColumnsResolver(
	v reflect.Value,
	parent string,
	requestedFields []string,
	value QuerySelectionInfo,
	autoInclude bool,
) QuerySelectionInfo {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return value
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		qs := field.Tag.Get("qs")
		column := field.Tag.Get("column")
		preload := field.Tag.Get("preload")
		typeof := field.Tag.Get("typeof")
		fieldValue := v.Field(i)

		fullName := parent + qs

		if qs == "" {
			continue
		}

		// Nested struct
		if typeof == "one" || typeof == "object" || typeof == "array" || typeof == "many2many" {
			if slices.Contains(requestedFields, qs) {
				if preload != "" {
					value.Preloads = append(value.Preloads, preload)
				}
			}

			value = DatabaseColumnsResolver(fieldValue, fullName+".", requestedFields, value, false)
		} else if typeof == "embed" {
			if slices.Contains(requestedFields, qs) {
				value = DatabaseColumnsResolver(fieldValue, fullName+".", requestedFields, value, true)
			}
		} else {
			if fieldValue.Type() == reflect.TypeOf(QueriableField{}) {
				if slices.Contains(requestedFields, qs) || autoInclude {
					value.Columns = append(value.Columns, column)
				}
			}
		}

	}

	return value
}

func GenerateQueryStringStyle(v reflect.Value, parent string) string {
	data := GenerateQuery(v, parent, "AsSql")

	query := strings.Join(data, " and ")
	return query
}

func GenerateQueryJqStyle(v reflect.Value, parent string) string {
	data := GenerateQuery(v, parent, "AsJq")
	if len(data) == 0 {
		return ""
	}

	query := strings.Join(data, " and ")

	return fmt.Sprintf("map(select(%s))", query)
}

func GenerateQuery(v reflect.Value, parent string, funcName string) []string {
	var values []string

	// Dereference pointer if it's a pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Only proceed if it's a struct
	if v.Kind() != reflect.Struct {
		return values
	}

	t := v.Type() // Get the struct type

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := parent + field.Name

		// Handle embedded/nested structs
		if field.Type.Kind() == reflect.Struct {
			values = append(values, GenerateQuery(v.Field(i), name+".", funcName)...)
		}

		// Skip non-QueriableField fields
		if !strings.Contains(field.Type.String(), "QueriableField") {
			continue
		}

		// Call the String method for QueriableField
		fieldValue := v.Field(i)
		if fieldValue.CanAddr() {
			method := fieldValue.Addr().MethodByName(funcName)
			if method.IsValid() {
				result := method.Call(nil) // Call String() with no arguments
				if len(result) > 0 {
					// Append the result to the values slice
					res := result[0].Interface().(string)

					if strings.TrimSpace(res) != "" {
						val := field.Tag.Get("column")
						values = append(values, strings.ReplaceAll(res, "$col", val))
					}
				}
			}
		}
	}

	return values
}

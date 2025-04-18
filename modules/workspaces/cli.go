package workspaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
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

	resultIsNil := result == nil || (reflect.ValueOf(result).Kind() == reflect.Ptr && reflect.ValueOf(result).IsNil())

	if !resultIsNil {
		body, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(body))

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
	os.Exit(0)
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

type QueriableField struct {
	Query     string
	Operation string
}

func (q QueriableField) String() string {
	if q.Query == "" {
		return ""
	}

	// Now we need to check, if there are no special signs, such as
	// =, < > ! , {", then user intended to be equal

	autoPrefix := "="
	if len(q.Query) > 0 {
		if q.Query[0] == '<' || q.Query[0] == '>' || q.Query[0] == '%' || q.Query[0] == '!' || q.Query[0] == '=' || q.Query[0] == '{' {
			// Then it's assumed user completed the expq.Queryssion,
			autoPrefix = ""
		}
	}

	// I need the reflect tag column from itself
	return "$col " + autoPrefix + q.Query
}

func QueriableFieldFromCliContext(v reflect.Value, parent string, c *cli.Context) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Dereference pointer
	}
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type() // Get the struct type

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := parent + field.Name

		cli := field.Tag.Get("cli")

		// Handle embedded/nested struct
		if field.Type.Kind() == reflect.Struct {
			QueriableFieldFromCliContext(v.Field(i), name+".", c)
		}

		if cli == "" {
			continue
		}

		// Set value if the field is a struct with a `Value` field
		fieldValue := v.Field(i)
		if fieldValue.CanSet() {
			if fieldValue.Type() == reflect.TypeOf(QueriableField{}) {
				fieldValue.Set(reflect.ValueOf(QueriableField{Query: c.String(cli)}))
			}
		}

	}
}

func QueriableFieldFromGinContext(v reflect.Value, parent string, c *gin.Context) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Dereference pointer
	}
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type() // Get the struct type

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := parent + field.Name

		qs := field.Tag.Get("qs")

		// Handle embedded/nested struct
		if field.Type.Kind() == reflect.Struct {
			QueriableFieldFromGinContext(v.Field(i), name+".", c)
		}

		if qs == "" {
			continue
		}

		// Set value if the field is a struct with a `Value` field
		fieldValue := v.Field(i)
		if fieldValue.CanSet() {
			if fieldValue.Type() == reflect.TypeOf(QueriableField{}) {

				qk, ok := c.GetQuery(qs)
				if ok {
					fieldValue.Set(reflect.ValueOf(QueriableField{Query: qk}))
				}
			}
		}

	}
}

func GenerateQueryStringStyle(v reflect.Value, parent string) string {
	data := GenerateQuery(v, parent)

	query := strings.Join(data, " and ")
	return query
}

func GenerateQuery(v reflect.Value, parent string) []string {
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
			values = append(values, GenerateQuery(v.Field(i), name+".")...)
		}

		// Skip non-QueriableField fields
		if !strings.Contains(field.Type.String(), "QueriableField") {
			continue
		}

		// Call the String method for QueriableField
		fieldValue := v.Field(i)
		if fieldValue.CanAddr() {
			method := fieldValue.Addr().MethodByName("String")
			if method.IsValid() {
				result := method.Call(nil) // Call String() with no arguments
				if len(result) > 0 {
					// Append the result to the values slice
					res := result[0].Interface().(string)

					if strings.TrimSpace(res) != "" {
						values = append(values, strings.ReplaceAll(res, "$col", field.Tag.Get("column")))
					}
				}
			}
		}
	}

	return values
}

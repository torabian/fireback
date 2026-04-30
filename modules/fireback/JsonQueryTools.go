package fireback

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/h22rana/jsonlogic2sql"
	"github.com/urfave/cli"
)

// Iterates through a QS struct, and tries to map the fields from there with a getValue callback.
// getValue can be configured, for example to read the query param and then return the value entered
// also can be used for different purpose, such as reading from cli query params.
func CaptureQueryStringStruct(v reflect.Value, parent string, getValue func(field reflect.StructField) string) {
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
			CaptureQueryStringStruct(fieldValue, name+".", getValue)
		}

		if fieldValue.CanSet() && fieldValue.Type() == reflect.TypeOf(QueriableField{}) {
			current := fieldValue.Interface().(QueriableField)
			current.UserInput = getValue(field)
			fieldValue.Set(reflect.ValueOf(current))
		}
	}
}

// Contains the field complete name, and user entered value, which can be later convereted into json-logic for example
type FlatKeyPairCondition struct {
	Field           string
	Value           string
	FinalColumnName string
}

func ConditionToJsonLogic(con []FlatKeyPairCondition) []byte {
	result := map[string]interface{}{}

	and := []interface{}{}
	for _, item := range con {

		// We might need to go deepr in nullability issues here.
		if item.Value == "" {
			continue
		}

		// In case it's a json entered, we automatically parse it, and do not assume anything.
		if strings.Contains(item.Value, "{") || strings.Contains(item.Value, "[") {
			var expression interface{}
			json.Unmarshal([]byte(item.Value), &expression)
			and = append(and, expression)
		} else {
			and = append(and, map[string]any{
				"contains": []any{
					map[string]any{
						"var": item.FinalColumnName,
					},
					item.Value,
				},
			})
		}

	}

	result["and"] = and
	a, _ := json.MarshalIndent(result, "", "  ")

	return a
}

// Converts a query string structure, into a json logic object.
func QuickOr(v reflect.Value, parent string) []FlatKeyPairCondition {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return []FlatKeyPairCondition{}
	}

	t := v.Type()

	res := []FlatKeyPairCondition{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := parent + field.Name
		fieldValue := v.Field(i)

		if field.Type.Kind() == reflect.Struct {
			add := QuickOr(fieldValue, name+".")
			res = append(res, add...)
		}

		if fieldValue.CanSet() && fieldValue.Type() == reflect.TypeOf(QueriableField{}) {
			current := fieldValue.Interface().(QueriableField)
			res = append(res, FlatKeyPairCondition{
				Field:           field.Tag.Get("qs"),
				Value:           current.UserInput,
				FinalColumnName: field.Tag.Get("column"),
			})
			fieldValue.Set(reflect.ValueOf(current))
		}
	}

	return res
}

func CaptureGinRequestIntoFilterQuery(qs any, c *gin.Context, f *QueryDSL) {
	CaptureRequestIntoFilterQuery(qs, func(field reflect.StructField) string {
		return c.Query(field.Tag.Get("qs"))
	}, f)
}

func CaptureCliRequestIntoFilterQuery(qs any, c *cli.Context, f *QueryDSL) {
	CaptureRequestIntoFilterQuery(qs, func(field reflect.StructField) string {
		return c.String(field.Tag.Get("cli"))
	}, f)
}

type FieldReader func(field reflect.StructField) string

func registerOperators(tr *jsonlogic2sql.Transpiler) {
	tr.RegisterOperatorFunc("contains", func(op string, args []interface{}) (string, error) {
		if len(args) != 2 {
			return "", fmt.Errorf("contains requires 2 arguments")
		}

		column := args[0].(string)

		value := args[1].(string)
		value = strings.Trim(value, `"'`)

		return fmt.Sprintf("%s LIKE '%%%s%%'", column, value), nil
	})
}

func CaptureRequestIntoFilterQuery(qs any, read FieldReader, f *QueryDSL) {
	v := reflect.ValueOf(qs)

	CaptureQueryStringStruct(v, "", read)

	tree := QuickOr(v, "")

	jsonLogic := ConditionToJsonLogic(tree)

	tr, _ := jsonlogic2sql.NewTranspiler(jsonlogic2sql.DialectBigQuery)

	registerOperators(tr)

	sql, err := tr.TranspileCondition(string(jsonLogic))
	if err != nil {
		return
	}

	f.FilterQuery = sql
}

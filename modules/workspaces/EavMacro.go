package workspaces

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	firebackgo "github.com/torabian/fireback/modules/workspaces/codegen/firebackgo"
	"github.com/xeipuuv/gojsonschema"
)

func prependCreateScript(name string) string {
	return `
	if fields, err := Cast` + ToUpper(name) + `FieldsFromJson(dto.JsonSchema); err != nil {
		return nil, err
	} else {
		dto.Fields = fields
	}
	`
}

func prependUpdateScript(name string) string {
	return `
	if fields2, err := Cast` + ToUpper(name) + `FieldsFromJson(fields.JsonSchema); err != nil {
		return nil, err
	} else {
		fields.Fields = fields2
	}
	`
}

func EavMacro(macro Module3Macro, x *Module3) {

	wsPrefix := "workspaces."
	if x.MetaWorkspace {
		wsPrefix = ""
	}

	key := macro.Name
	eavMacroTools, err := CompileString(&firebackgo.FbGoTpl, "EavMacro.tpl", gin.H{
		"Key":      ToUpper(key),
		"key":      key,
		"wsprefix": wsPrefix,
	})
	if err != nil {
		log.Fatalln(err)
	}

	form := Module3Entity{
		Name:                key,
		PrependScript:       eavMacroTools,
		PrependCreateScript: prependCreateScript(key),
		PrependUpdateScript: prependUpdateScript(key),
		Fields: []*Module3Field{
			{
				Name:     "name",
				Type:     "string",
				Validate: "required",
			},
			{
				Name: "description",
				Type: "string",
			},
			{
				Name: "uiSchema",
				Type: "json",
			},
			{
				Name: "jsonSchema",
				Type: "json",
			},
			{
				Name: "fields",
				Type: "array",
				Fields: []*Module3Field{
					{
						Name:   key,
						Type:   "one",
						Target: ToUpper(key) + "Entity",
					},
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "type",
						Type: "string",
					},
				},
			},
		},
	}

	submissionFields := []*Module3Field{
		{
			Name:     key,
			Type:     "one",
			Target:   ToUpper(key) + "Entity",
			Validate: "required",
		},
		{
			Name: "data",
			Type: "json",
		},
		{
			Name: "values",
			Type: "array",
			Fields: []*Module3Field{
				{
					Name:      key + "Field",
					Type:      "one",
					Target:    ToUpper(key) + "Fields",
					RootClass: ToUpper(key) + "Entity",
				},
				{
					Name: "valueInt64",
					Type: "int64",
				},
				{
					Name: "valueFloat64",
					Type: "float64",
				},
				{
					Name: "valueString",
					Type: "string",
				},
				{
					Name: "valueBoolean",
					Type: "bool",
				},
			},
		},
	}

	submissionFields = append(submissionFields, macro.Fields...)

	formSubmission := Module3Entity{
		Name:                key + "Submission",
		Fields:              submissionFields,
		PrependCreateScript: ToUpper(key) + `SubmissionCastFieldsToEavAndValidate(dto, query)`,
		PrependUpdateScript: ToUpper(key) + `SubmissionCastFieldsToEavAndValidate(fields, query)`,
	}

	x.Entities = append(x.Entities, form, formSubmission)
}

type UISchema struct {
	Field            string              `json:"ui:field,omitempty" yaml:"ui:field,omitempty"`
	Widget           string              `json:"ui:widget,omitempty" yaml:"ui:widget,omitempty"`
	Options          map[string]string   `json:"ui:options,omitempty" yaml:"ui:options,omitempty"`
	Order            []string            `json:"ui:order,omitempty" yaml:"ui:order,omitempty"`
	Fields           map[string]UISchema `json:"ui:fields,omitempty" yaml:"ui:fields,omitempty"`
	Autocomplete     string              `json:"ui:autocomplete,omitempty" yaml:"ui:autocomplete,omitempty"`
	Addable          bool                `json:"ui:addable,omitempty" yaml:"ui:addable,omitempty"`
	Removable        bool                `json:"ui:removable,omitempty" yaml:"ui:removable,omitempty"`
	Disabled         bool                `json:"ui:disabled,omitempty" yaml:"ui:disabled,omitempty"`
	ReadOnly         bool                `json:"ui:readonly,omitempty" yaml:"ui:readonly,omitempty"`
	Hidden           bool                `json:"ui:hidden,omitempty" yaml:"ui:hidden,omitempty"`
	Title            string              `json:"ui:title,omitempty" yaml:"ui:title,omitempty"`
	Description      string              `json:"ui:description,omitempty" yaml:"ui:description,omitempty"`
	Format           string              `json:"ui:format,omitempty" yaml:"ui:format,omitempty"`
	EnumNames        []string            `json:"ui:enumNames,omitempty" yaml:"ui:enumNames,omitempty"`
	Default          interface{}         `json:"ui:default,omitempty" yaml:"ui:default,omitempty"`
	Help             string              `json:"ui:help,omitempty" yaml:"ui:help,omitempty"`
	Placeholder      string              `json:"ui:placeholder,omitempty" yaml:"ui:placeholder,omitempty"`
	MinItems         int                 `json:"ui:minItems,omitempty" yaml:"ui:minItems,omitempty"`
	MaxItems         int                 `json:"ui:maxItems,omitempty" yaml:"ui:maxItems,omitempty"`
	MinLength        int                 `json:"ui:minLength,omitempty" yaml:"ui:minLength,omitempty"`
	MaxLength        int                 `json:"ui:maxLength,omitempty" yaml:"ui:maxLength,omitempty"`
	Pattern          string              `json:"ui:pattern,omitempty" yaml:"ui:pattern,omitempty"`
	MultipleOf       float64             `json:"ui:multipleOf,omitempty" yaml:"ui:multipleOf,omitempty"`
	Minimum          float64             `json:"ui:minimum,omitempty" yaml:"ui:minimum,omitempty"`
	Maximum          float64             `json:"ui:maximum,omitempty" yaml:"ui:maximum,omitempty"`
	ExclusiveMinimum bool                `json:"ui:exclusiveMinimum,omitempty" yaml:"ui:exclusiveMinimum,omitempty"`
	ExclusiveMaximum bool                `json:"ui:exclusiveMaximum,omitempty" yaml:"ui:exclusiveMaximum,omitempty"`
}

type JSONSchema struct {
	Title       string                 `json:"title,omitempty" yaml:"title,omitempty"`
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Type        string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Properties  map[string]SchemaField `json:"properties,omitempty" yaml:"properties,omitempty"`
	Required    []string               `json:"required,omitempty" yaml:"required,omitempty"`
}

func (x *JSONSchema) FromJson(schema *JSON) error {
	if schema == nil {
		return nil
	}
	k, err := schema.MarshalJSON()
	if err == nil {
		json.Unmarshal(k, &x)
		return nil
	} else {
		return err
	}
}

type SchemaField struct {
	Type        string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Format      string                 `json:"format,omitempty" yaml:"format,omitempty"`
	Pattern     string                 `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Enum        []interface{}          `json:"enum,omitempty" yaml:"enum,omitempty"`
	Properties  map[string]SchemaField `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items       *SchemaField           `json:"items,omitempty" yaml:"items,omitempty"`
}

type FieldInfo struct {
	Field string `json:"field" yaml:"field"`
	Type  string `json:"type" yaml:"type"`
}

// Keep this for now we wil see why we need it later, on parsing the data
func FlattenData(data map[string]interface{}, prefix string) map[string]interface{} {
	flatData := make(map[string]interface{})

	for key, value := range data {
		keyStr := fmt.Sprintf("%s%s", prefix, key)
		switch v := value.(type) {
		case map[string]interface{}:
			flatMap := FlattenData(v, keyStr+".")
			for k, v := range flatMap {
				flatData[k] = v
			}
		case []interface{}:
			for i, item := range v {
				switch itemValue := item.(type) {
				case map[string]interface{}:
					flatMap := FlattenData(itemValue, fmt.Sprintf("%s[%d].", keyStr, i))
					for k, v := range flatMap {
						flatData[k] = v
					}
				default:
					flatData[keyStr] = v
				}
			}
		default:
			flatData[keyStr] = value
		}
	}
	return flatData
}

func FlattenFields(prefix string, fields map[string]SchemaField) map[string]SchemaField {
	flatFields := make(map[string]SchemaField)

	for key, field := range fields {
		fieldName := key
		if prefix != "" {
			fieldName = prefix + "." + key
		}
		flatFields[fieldName] = field

		if field.Properties != nil {
			subFields := FlattenFields(fieldName, field.Properties)
			for subKey, subField := range subFields {
				flatFields[subKey] = subField
			}
		}
	}

	return flatFields
}

func ValidateEavSchema(schemaStr *JSON, dataJson *JSON) *IError {
	if schemaStr == nil {
		return nil
	}

	loader3 := gojsonschema.NewStringLoader(schemaStr.String())
	sl := gojsonschema.NewSchemaLoader()

	schema, err5 := sl.Compile(loader3)

	if err5 != nil {
		return GormErrorToIError(err5)
	}

	data := "{}"

	if dataJson != nil {
		data = dataJson.String()
	}

	documentLoader := gojsonschema.NewStringLoader(data)

	result, err6 := schema.Validate(documentLoader)
	if err6 != nil {
		return GormErrorToIError(err6)
	}

	if err9 := JsonSchemaToIError(result); err9 != nil {
		return err9
	}

	return nil
}

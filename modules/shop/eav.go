package shop

import (
	"encoding/json"
	"fmt"

	"github.com/torabian/fireback/modules/workspaces"
	"gorm.io/gorm"
)

type EavModel struct {
	JsonSchema JSONSchema                  `json:"jsonSchema,omitempty" yaml:"jsonSchema,omitempty"`
	UiSchema   map[string]UISchema         `json:"uiSchema,omitempty" yaml:"uiSchema,omitempty"`
	Data       map[interface{}]interface{} `json:"data,omitempty" yaml:"data,omitempty"`
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

func (x *EavModel) FlattenFields() map[string]SchemaField {
	return flattenFields("", x.JsonSchema.Properties)

}

func (x *EavModel) StoreInDatabase(query workspaces.QueryDSL) (*EavModel, *workspaces.IError) {

	return workspaces.RunTransaction[EavModel](x, query, func(tx *gorm.DB) error {
		query.Tx = tx

		js1, err0 := json.Marshal(x.UiSchema)
		if err0 != nil {
			return err0
		}

		uiSchema := &workspaces.JSON{}
		uiSchema.Scan(js1)

		js2, err0 := json.Marshal(x.JsonSchema)
		if err0 != nil {
			return err0
		}

		jsonSchema := &workspaces.JSON{}
		jsonSchema.Scan(js2)

		fields := []*FormFields{}
		for key, field := range x.FlattenFields() {
			key := key
			field := field
			fields = append(fields, &FormFields{
				Type: &field.Type,
				Name: &key,
			})
		}

		_, err := FormActionCreate(&FormEntity{
			Name:        &x.JsonSchema.Title,
			Description: &x.JsonSchema.Description,
			UiSchema:    uiSchema,
			JsonSchema:  jsonSchema,
			Fields:      fields,
		}, query)

		if err != nil {
			return err
		}

		return nil
	})
}

func flattenData(data map[interface{}]interface{}, prefix string) map[string]interface{} {
	flatData := make(map[string]interface{})

	for key, value := range data {
		keyStr := fmt.Sprintf("%s%s", prefix, key)
		switch v := value.(type) {
		case map[interface{}]interface{}:
			flatMap := flattenData(v, keyStr+".")
			for k, v := range flatMap {
				flatData[k] = v
			}
		case []interface{}:
			for i, item := range v {
				switch itemValue := item.(type) {
				case map[interface{}]interface{}:
					flatMap := flattenData(itemValue, fmt.Sprintf("%s[%d].", keyStr, i))
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
func flattenFields(prefix string, fields map[string]SchemaField) map[string]SchemaField {
	flatFields := make(map[string]SchemaField)

	for key, field := range fields {
		fieldName := fmt.Sprintf("%s.%s", prefix, key)
		flatFields[fieldName] = field

		if field.Properties != nil {
			subFields := flattenFields(fieldName, field.Properties)
			for subKey, subField := range subFields {
				flatFields[subKey] = subField
			}
		}
	}

	return flatFields
}

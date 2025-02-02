package workspaces

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func AppendEavCustomParams(schema *jsonschema.Schema) {

	using := orderedmap.New[string, *jsonschema.Schema]()
	using.Set("using", &jsonschema.Schema{
		Const: "eav",
	})

	for key := range schema.Definitions {
		if key == "Module3Macro" {

			jsonStr := `{
				"oneOf": [
					{
						"if": {
							"properties": {
								"using": {
									"const": "eav"
								}
							}
						},
						"then": {
							"properties": {
								"params": {
									"$ref": "#/definitions/EavMacroParams"
								}
							}
						}
					}
				]
			}`

			var schemaP jsonschema.Schema
			err := json.Unmarshal([]byte(jsonStr), &schemaP)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			schema.Definitions[key].OneOf = schemaP.OneOf
		}
		if key == "Module3ConfigField" {

			jsonStr := `{
				"anyOf": [
					{
					"if": {
						"properties": {
						"type": {
							"const": "bool"
						}
						}
					},
					"then": {
						"properties": {
						"default": {
							"type": "boolean"
						}
						}
					}
					},
					{
					"if": {
						"properties": {
						"type": {
							"const": "string"
						}
						}
					},
					"then": {
						"properties": {
						"default": {
							"type": "string"
						}
						}
					}
					}
				]
			}`

			var schemaP jsonschema.Schema
			err := json.Unmarshal([]byte(jsonStr), &schemaP)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			var schemaEmpty jsonschema.Schema
			json.Unmarshal([]byte("{}"), &schemaEmpty)

			schema.Definitions[key].Properties.Set("default", &schemaEmpty)
			schema.Definitions[key].AnyOf = schemaP.AnyOf
		}
	}

	// add the missing definitions
	reflector := jsonschema.Reflector{}
	schema2 := reflector.Reflect(&EavMacroParams{})

	for key := range schema2.Definitions {
		schema.Definitions[key] = schema2.Definitions[key]
	}

}

func GenerateJsonSpecForModule3(source string, out string, updateVsCodeSettings string) string {
	// Create a reflector
	reflector := jsonschema.Reflector{}
	schema := reflector.Reflect(&Module3{})

	AppendEavCustomParams(schema)

	// Convert the schema to JSON
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(err)
	}

	toWrite := string(schemaJSON)
	toWrite = strings.ReplaceAll(toWrite, "$defs", "definitions")
	toWrite = strings.ReplaceAll(toWrite, "https://json-schema.org/draft/2020-12/schema", "http://json-schema.org/draft-07/schema#")

	if updateVsCodeSettings != "" {
		first := strings.ReplaceAll(out, ".yml", ".json")
		second := strings.ReplaceAll(out, ".jsonschemas/", "")
		second = strings.ReplaceAll(second, ".jsonschemas\\", "")
		if err := UpdateYamlSchemas(updateVsCodeSettings, first, second); err != nil {
			fmt.Println(err)
		}
	}

	if out != "" {
		basePath := filepath.Dir(out)

		if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
			return ""
		}

		out = strings.ReplaceAll(out, ".yml", ".json")
		os.WriteFile(out, []byte(toWrite), 0644)

		return ""
	}

	return toWrite
}

func UpdateYamlSchemas(settingsPath, newSchemaKey, newSchemaValue string) error {
	// Read the existing settings.json file
	data, err := os.ReadFile(settingsPath)

	if err != nil {
		return fmt.Errorf("failed to read settings file: %w", err)
	}

	// Parse the JSON into a map
	var settings map[string]interface{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Access or create the yaml.schemas section
	yamlSchemas, ok := settings["yaml.schemas"].(map[string]interface{})
	if !ok {
		yamlSchemas = make(map[string]interface{})
		settings["yaml.schemas"] = yamlSchemas
	}

	// Add the new key-value pair to yaml.schemas
	yamlSchemas[newSchemaKey] = []string{newSchemaValue}

	// Convert the updated settings back to JSON
	updatedData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal updated settings: %w", err)
	}

	// Write the updated JSON back to the file
	if err := os.WriteFile(settingsPath, updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write updated settings file: %w", err)
	}

	return nil
}

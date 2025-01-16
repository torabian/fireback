package workspaces

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/invopop/jsonschema"
)

func GenerateJsonSpecForModule3(source string, out string, updateVsCodeSettings string) string {
	// Create a reflector
	reflector := jsonschema.Reflector{}

	// Generate the schema
	schema := reflector.Reflect(&Module3{})
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

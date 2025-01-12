package workspaces

import (
	"fmt"
	reflect "reflect"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"gopkg.in/yaml.v2"
)

// Define your struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetTypeName(v any) string {
	u := reflect.TypeOf(v).String()
	u = strings.ReplaceAll(u, "/", ".")
	return u
}

// Function to convert struct to OpenAPI 3 schema and output as YAML
func ConvertStructToOpenAPIYaml(xapp *FirebackApp) (string, error) {

	paths := &openapi3.Paths{
		Extensions: map[string]interface{}{},
	}

	components := &openapi3.Components{
		Schemas: openapi3.Schemas{},
	}

	for _, item := range xapp.Modules {
		for _, bundle := range item.EntityBundles {
			CodeItem(bundle.Actions, paths, components)
		}

		for _, actions := range item.Actions {
			CodeItem(actions, paths, components)
		}
	}
	// Create the OpenAPI 3 document
	openapi := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   "Open API Document",
			Version: "1.0.0",
		},
		Paths:      paths,
		Components: components,
	}

	// Marshal the OpenAPI document to YAML
	yamlData, err := yaml.Marshal(openapi)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}

func CodeItem(actions []Module3Action, paths *openapi3.Paths, components *openapi3.Components) error {
	for _, action := range actions {
		fmt.Println("Action:", action.Name, action.Url)
		opt := &openapi3.Operation{
			Summary:     action.Description,
			Description: action.Description,
			Responses: &openapi3.Responses{
				Extensions: map[string]interface{}{
					"200": openapi3.ResponseRef{
						Value: &openapi3.Response{
							Content: openapi3.Content{
								"application/json": &openapi3.MediaType{
									Schema: &openapi3.SchemaRef{
										Ref: "#/components/schemas/" + GetTypeName(action.ResponseEntity),
									},
								},
							},
						},
					},
				},
			},
		}

		itemPath := &openapi3.PathItem{}

		if action.Method == "GET" || action.Format == "QUERY" {
			itemPath.Get = opt
		}

		if action.Method == "POST" {
			itemPath.Post = opt
		}

		if action.Method == "PATCH" {
			itemPath.Patch = opt
		}

		if action.Method == "DELETE" {
			itemPath.Delete = opt
		}

		paths.Extensions[action.Url] = itemPath

		if action.RequestEntity != nil {
			// Generate OpenAPI 3 schema from the struct
			schemaRef, err := openapi3gen.NewSchemaRefForValue(action.RequestEntity, nil)
			if err != nil {
				return err
			}
			components.Schemas[GetTypeName(action.RequestEntity)] = schemaRef
		}

		if action.ResponseEntity != nil {
			// Generate OpenAPI 3 schema from the struct
			schemaRef, err := openapi3gen.NewSchemaRefForValue(action.ResponseEntity, nil)
			if err != nil {
				return err
			}

			components.Schemas[GetTypeName(action.ResponseEntity)] = schemaRef
		}
	}

	return nil
}

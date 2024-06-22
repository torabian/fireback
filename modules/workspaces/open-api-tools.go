package workspaces

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/swaggest/openapi-go/openapi3"
)

// func ToModule2Actions(paths *openapi3.Paths) []Module2Action {

// 	items := []Module2Action{}

// 	return items
// }

func OpenApiRefObjectToFireback(ref string) string {

	return strings.ReplaceAll(ref, "#/components/schemas/", "")
}

func OpenApiPrimitveToFireback(field string) string {
	switch field {
	case "integer":
		return "float64"
	case "string":
		return "string"
	}
	return field
}

func OpenApiPropertiesToFirebackType(t *openapi3.Schema) []Module2Field {

	items := []Module2Field{}

	for property, content := range t.Properties {
		if content.SchemaReference != nil {
			items = append(items, Module2Field{
				Name:   property,
				Type:   "one",
				Target: OpenApiRefObjectToFireback(content.SchemaReference.Ref),
			})
		}
	}

	fmt.Println("items", items)

	return items

}

func OpenApiPropertiesToFirebackDto(t *openapi3.Schema) string {

	for _, content := range t.Properties {
		if content.SchemaReference != nil {
			return OpenApiRefObjectToFireback(content.SchemaReference.Ref)
		}
	}

	return ""

}

func OpenApiTypeToFirebackType(fieldName string, t *openapi3.Schema) Module2Field {

	field := Module2Field{
		Name: ToLower(fieldName),
		Type: "unknown",
	}

	if t == nil || t.Type == nil {
		return field
	}

	v := string(*t.Type)

	switch v {
	case "string":
		field.Type = "string"
	case "number":
		field.Type = "float64"
	case "boolean":
		field.Type = "bool"
	// case "object":
	// fmt.Println("Object!")
	case "array":
		field.Type = "array"

		// This is scenario when the field is a primitive, not dto
		if t.Items.Schema != nil && t.Items.Schema.Type != nil {
			field.Type = FIELD_TYPE_ARRAYP
			field.Primitive = OpenApiPrimitveToFireback(string(*t.Items.Schema.Type))
		} else if t.Items.SchemaReferenceEns() != nil {
			ref := t.Items.SchemaReferenceEns().Ref
			field.Target = OpenApiRefObjectToFireback(ref)
		}
	}

	return field
}

func endsWithIgnoreCase(str, suffix string) bool {
	str = strings.ToLower(str)
	suffix = strings.ToLower(suffix)

	return strings.HasSuffix(str, suffix)
}

func removeLastCharacters(input string, num int) string {
	if num >= len(input) {
		return ""
	}
	return input[:len(input)-num]
}

// fireback has some conventions about the naming, needs to be respected
func NormalizeOpenApi3DtoName(field string) string {

	if endsWithIgnoreCase(field, "dto") {
		field = removeLastCharacters(field, 3)
	}

	field = ToLower(field)

	return field
}

func ConvertOpenAPIRouteToGinPathParam(url string) string {
	re := regexp.MustCompile("({([^{}]*)})")
	return re.ReplaceAllString(url, ":$2")
}

func OpenApiSchemasToFirebackDtos(ref map[string]openapi3.SchemaOrRef) []Module2DtoBase {

	dtos := []Module2DtoBase{}

	for schemaName, refFields := range ref {
		fields := []*Module2Field{}
		for field, property := range refFields.Schema.Properties {
			if property.Schema != nil && property.Schema.Type != nil {
				o := OpenApiTypeToFirebackType(field, property.Schema)
				fields = append(fields, &o)
			}
		}

		dtos = append(dtos, Module2DtoBase{
			Name:   NormalizeOpenApi3DtoName(schemaName),
			Fields: fields,
		})
	}

	return dtos
}

// This is a complex conversion
func OpenApiToFireback(s openapi3.Spec) *Module2 {

	actions := []*Module2Action{}
	for url, content := range s.Paths.MapOfPathItemValues {
		for method, opt := range content.MapOfOperationValues {
			action := Module2Action{
				Url:    ConvertOpenAPIRouteToGinPathParam(url),
				Method: strings.ToUpper(method),
			}

			if opt.RequestBody != nil {
				if c := opt.RequestBody.RequestBody.Content; c != nil {
					if v, okay := c["application/json"]; okay {
						if v.Schema != nil {
							ref := OpenApiRefObjectToFireback(string(v.Schema.SchemaReference.Ref))
							if ref != "" {
								action.In.Dto = ref
							}
						}
					}
				}
			}

			if opt.Responses.MapOfResponseOrRefValues != nil {

				for resType, response := range opt.Responses.MapOfResponseOrRefValues {

					if content, okay := response.Response.Content["application/json"]; okay {

						if content.Schema.SchemaReference != nil {
							fmt.Println(5, resType, content.Schema.SchemaReference.Ref)
						} else if schema := content.Schema.Schema; schema != nil {

							for _, ref := range schema.AllOf {

								if ref.SchemaReference != nil {
									// fmt.Println(71, ref.SchemaReference.Ref)
								} else if ref.Schema != nil {
									action.Out.Dto = OpenApiPropertiesToFirebackDto(ref.Schema)

								}
							}

						}
					}

				}
			}

			actions = append(actions, &action)
		}
	}

	dtos := OpenApiSchemasToFirebackDtos(s.Components.Schemas.MapOfSchemaOrRefValues)

	return &Module2{
		Dto:     dtos,
		Name:    s.Info.Title,
		Version: s.Info.Version,
		// Description: *s.Info.Description,
		Actions: actions,
	}
}

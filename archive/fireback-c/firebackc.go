package workspaces

import fibreackc "github.com/torabian/fireback/modules/workspaces/codegen/fireback-c"

func CComputedField(field *Module3Field, isWorkspace bool) string {
	prefix := ""
	if !isWorkspace {
		prefix = "workspaces."
	}
	switch field.Type {
	case "string", "text", "html":
		return "*string"
	case "enum":
		return "*string"
	case "one":
		if field.Module != "" {
			return field.Module + "." + field.Target
		}
		return field.Target
	case "array":
		return field.PublicName()
	case "daterange":
		return prefix + "XDate"
	case "any":
		return "interface{}"
	case "many2many":
		if field.Module != "" {
			return field.Module + "." + field.Target
		}
		return field.Target
	case "arrayP":
		return "[]" + field.Primitive
	case "int64", "int32", "int", "float64", "float32", "float", "bool":
		return "*" + field.Type
	case "Timestamp":
		return "*string"
	case "datenano":
		return "int64"
	case "boolean":
		return "*bool"
	case "double":
		return "*float64"
	case "object":
		return field.PublicName()
	case "json":
		return "JSON"
	case "date":
		return prefix + "XDate"
	default:
		return "*" + field.Type
	}
}

func CEntityDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + "Entity.dyno.c"
}
func CEntityExtensionDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + "Entity.c"
}
func CEntityHeaderDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + "Entity.dyno.h"
}

var FirebackCGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName: "FirebackC",
	ComputeField: CComputedField,
	Templates:    fibreackc.FbCtpl,

	EntityGeneratorTemplate: "CEntity.tpl",
	EntityDiskName:          CEntityDiskName,
	EntityExtensionDiskName: CEntityExtensionDiskName,
	EntityHeaderDiskName:    CEntityHeaderDiskName,

	DtoDiskName: GoDtoDiskName,
}

// This was in cli code gen
// {
// 	Flags: commonFlags,
// 	Name:  "cem",
// 	Usage: "Generates the C embedded tools for microcontrollers",
// 	Action: func(c *cli.Context) error {

// 		RunCodeGen(xapp, GenContextFromCli(c, FirebackCGenCatalog))

// 		return nil
// 	},
// },

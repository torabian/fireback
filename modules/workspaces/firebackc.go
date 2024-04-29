package workspaces

import fibreackc "github.com/torabian/fireback/modules/workspaces/codegen/fireback-c"

func CComputedField(field *Module2Field, isWorkspace bool) string {
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

func CEntityDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Entity.dyno.c"
}
func CEntityExtensionDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Entity.c"
}
func CEntityHeaderDiskName(x *Module2Entity) string {
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

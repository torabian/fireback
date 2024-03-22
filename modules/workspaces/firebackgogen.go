package workspaces

import firebackgo "github.com/torabian/fireback/modules/workspaces/codegen/firebackgo"

func GolangComputedField(field *Module2Field, isWorkspace bool) string {
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

func GoEntityDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Entity.dyno.go"
}
func GoEntityExtensionDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Entity.go"
}
func GoActionDiskName(moduleName string) string {
	return ToUpper(moduleName) + "CustomActions.dyno.go"
}

func GoDtoDiskName(x *Module2DtoBase) string {
	return ToUpper(x.Name) + "Dto.dyno.go"
}

var FirebackGoGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName:                     "FirebackGo",
	ComputeField:                     GolangComputedField,
	Templates:                        firebackgo.FbGoTpl,
	EntityGeneratorTemplate:          "GoEntity.tpl",
	DtoGeneratorTemplate:             "GoDto.tpl",
	CteSqlTemplate:                   "SqlCteQuery.tpl",
	ActionGeneratorTemplate:          "GoAction.tpl",
	ActionDiskName:                   GoActionDiskName,
	EntityDiskName:                   GoEntityDiskName,
	EntityExtensionGeneratorTemplate: "GoEntityExtension.tpl",
	EntityExtensionDiskName:          GoEntityExtensionDiskName,

	DtoDiskName: GoDtoDiskName,
}

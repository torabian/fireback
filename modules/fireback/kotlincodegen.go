package fireback

import (
	kotlintpl "github.com/torabian/fireback/modules/fireback/codegen/andkot"
	kotlininclude "github.com/torabian/fireback/modules/fireback/codegen/andkot/include"
)

func KotlinPrimitive(primitive string) string {
	switch primitive {
	case "string", "text":
		return "String"
	case "int64", "int32", "int":
		return "Int"
	case "float64", "float32", "float":
		return "Float"
	case "bool":
		return "Boolean"
	case "double":
		return "Double"
	default:
		return "String"
	}
}

func KotlinComputedField(field *Module3Field, isWorkspace bool) string {
	switch field.Type {
	case "string", "text":
		return "String"
	case "one":
		if field.Module != "" {
			return field.Module + "." + field.Target
		}
		return field.Target
	case "int64", "int32", "int":
		return "Int"
	case "float64", "float32", "float":
		return "Float"
	case "html", "json":
		return "String"
	case "array":
		return "List<" + field.PublicName() + ">"
	case "many2many":
		target := field.Target
		if field.Module != "" {
			target = field.Module + "." + field.Target
		}
		return "List<" + target + ">"
	case "object":
		if field.Module != "" {
			return field.Module + "." + field.PublicName()
		}
		return field.PublicName()
	case "arrayP":
		return "List<" + KotlinPrimitive(field.Primitive) + ">"
	case "bool":
		return "Boolean"
	case "enum":
		return "String"
	case "date":
		return "java.util.Date"
	case "daterange":
		return "com.fireback.DateRange"
	case "Timestamp", "datenano":
		return "String"
	case "double":
		return "Double"
	case "any":
		return "Any"
	default:
		return field.Type
	}
}
func KotlinEntityDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + ".kt"
}

func KotlinDtoDiskName(x *Module3Dto) string {
	return ToUpper(x.Name) + "Dto.kt"
}

func KotlinFormDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + "Form.kt"
}

func KotlinRpcDiskName(x *Module3Action) string {
	return ToUpper(x.GetFuncName()) + ".kt"
}

func KotlinActionDiskName(action *Module3Action, moduleName string) string {
	return ToUpper(action.Name) + "Action.kt"
}

var KotlinGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName:            "andkot",
	ComputeField:            KotlinComputedField,
	RpcPostDiskName:         KotlinRpcDiskName,
	RpcPost:                 "KotlinRpcPost.tpl",
	Templates:               kotlintpl.KotlinTpls,
	IncludeDirectory:        &kotlininclude.JavaInclude,
	RpcQueryDiskName:        KotlinRpcDiskName,
	RpcQuery:                "KotlinRpcQuery.tpl",
	EntityGeneratorTemplate: "KotlinEntity.tpl",
	DtoGeneratorTemplate:    "KotlinDto.tpl",
	ActionGeneratorTemplate: "KotlinActionDto.tpl",
	SingleActionDiskName:    KotlinActionDiskName,
	Prettier: func(name string) {
		// Java formatter is super slow. We skip it for now.
		// But make sure you have the google java format jar
		// from: https://github.com/google/google-java-format/releases, and it would format:
		// if strings.HasSuffix(name, ".java") {
		// 	RunExecCmd("java", []string{"-jar", "./google-java-format.jar", "-i", name})
		// }
	},

	EntityDiskName: KotlinEntityDiskName,
	DtoDiskName:    KotlinDtoDiskName,
}

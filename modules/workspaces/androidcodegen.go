package workspaces

import (
	javatpl "github.com/torabian/fireback/modules/workspaces/codegen/java"
	javainclude "github.com/torabian/fireback/modules/workspaces/codegen/java/include"
)

func JavaPrimitve(primitive string) string {
	switch primitive {
	case "string", "text":
		return "String"
	case "int64", "int32", "int":
		return "Integer"
	case "float64", "float32", "float":
		return "Float"
	case "bool":
		return "Boolean"
	case "double":
		return "Double"
	default:
		// Sometimes dto's which are the primitive in golang, actually
		// are not compiled via fireback, because they are internal.
		// For now, they are not accessible, and we consider them as String
		// Which is wrong. The use case of such classes is very limited,
		// should not be a problem on major cases.
		return "String"
	}
}

func JavaComputedField(field *Module3Field, isWorkspace bool) string {
	switch field.Type {
	case "string", "text":
		return "String"
	case "one":
		if field.Module != "" {
			return field.Module + "." + field.Target
		}
		return field.Target
	case "int64", "int32", "int":
		return "int"
	case "float64", "float32", "float":
		return "float"
	case "html":
		// Think about this, make an object called Html maybe?
		return "String"
	case "json":
		return "String"
		// return TsCalcJsonField(field)
	case "array":
		return field.PublicName() + "[]"
	case "many2many":
		if field.Module != "" {
			return field.Module + "." + field.Target + "[]"
		}
		return field.Target + "[]"
	case "object":
		if field.Module != "" {
			return field.Module + "." + field.PublicName()
		}
		return field.PublicName()
	case "arrayP":
		return JavaPrimitve(field.Primitive) + "[]"
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
		return "double"
	case "any":
		return "String"

	default:
		return field.Type
	}
}

func JavaEntityDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + "Entity.java"
}

func JavaDtoDiskName(x *Module3DtoBase) string {
	return ToUpper(x.Name) + "Dto.java"
}

func JavaFormDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + "Form.java"
}

func JavaRpcCommonDiskName(x *Module3Action) string {
	return ToUpper(x.GetFuncName()) + ".java"
}

func JavaActionDiskName(action *Module3Action, moduleName string) string {
	return ToUpper(action.Name) + "Action.java"
}

var JavaGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName:            "android",
	ComputeField:            JavaComputedField,
	RpcPostDiskName:         JavaRpcCommonDiskName,
	RpcPost:                 "JavaRpcPost.tpl",
	Templates:               javatpl.JavaTpls,
	IncludeDirectory:        &javainclude.JavaInclude,
	RpcQueryDiskName:        JavaRpcCommonDiskName,
	RpcQuery:                "JavaRpcQuery.tpl",
	EntityGeneratorTemplate: "JavaEntity.tpl",
	DtoGeneratorTemplate:    "JavaDto.tpl",
	ActionGeneratorTemplate: "JavaActionDto.tpl",
	SingleActionDiskName:    JavaActionDiskName,
	Prettier: func(name string) {
		// Java formatter is super slow. We skip it for now.
		// But make sure you have the google java format jar
		// from: https://github.com/google/google-java-format/releases, and it would format:
		// if strings.HasSuffix(name, ".java") {
		// 	RunExecCmd("java", []string{"-jar", "./google-java-format.jar", "-i", name})
		// }
	},

	EntityDiskName: JavaEntityDiskName,
	DtoDiskName:    JavaDtoDiskName,
}

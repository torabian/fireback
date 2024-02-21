package workspaces

import (
	swiftpl "github.com/torabian/fireback/modules/workspaces/codegen/swift"
	swiftinclude "github.com/torabian/fireback/modules/workspaces/codegen/swift/include"
)

func SwiftPrimitve(primitive string) string {
	switch primitive {
	case "string", "text":
		return "String"
	case "int64", "int32", "int":
		return "Int"
	case "float64", "float32", "float":
		return "Float64"
	case "bool":
		return "Bool"
	case "double":
		return "Double"
	default:
		return "Any"
	}
}

func SwiftComputedField(field *Module2Field) string {
	switch field.Type {
	case "string", "text":
		return "String?"
	case "one":
		return field.Target + "?"
	case "array":
		return field.PublicName()
	case "many2many":
		return field.Target
	case "object":
		return field.PublicName() + "?"
	case "json", "any":
		return "String?"
	case "arrayP":
		return "[" + SwiftPrimitve(field.Primitive) + "]?"
	case "int64", "int32", "int":
		return "Int?"
	case "float64", "float32", "float":
		return "Float64?"
	case "bool":
		return "Bool?"
	case "date":
		return "Date?"
	case "Timestamp", "datenano":
		return "String?"
	case "double":
		return "Double?"
	default:
		return field.Type + "?"
	}
}

func SwiftDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Entity.swift"
}

func SwiftDtoDiskName(x *Module2DtoBase) string {
	return ToUpper(x.Name) + "Dto.swift"
}

func SwiftFormDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Form.swift"
}

func SwiftRpcCommonDiskName(x *Module2Action) string {
	return ToUpper(x.GetFuncName()) + ".swift"
}

func SwiftActionDiskName(moduleName string) string {
	return ToUpper(moduleName) + "ActionsDto.swift"
}

var SwiftGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName:            "swift",
	ComputeField:            SwiftComputedField,
	IncludeDirectory:        &swiftinclude.SwiftInclude,
	ActionDiskName:          SwiftActionDiskName,
	ActionGeneratorTemplate: "SwiftActionDto.tpl",
	Templates:               swiftpl.SwiftTpl,
	EntityGeneratorTemplate: "SwiftEntity.tpl",
	DtoGeneratorTemplate:    "SwiftDto.tpl",
	// FormGeneratorTemplate:   "SwiftUIForm.tpl",
	EntityDiskName:   SwiftDiskName,
	FormDiskName:     SwiftFormDiskName,
	RpcQueryDiskName: SwiftRpcCommonDiskName,

	RpcPost:         "SwiftRpcPost.tpl",
	RpcQuery:        "SwiftRpcQuery.tpl",
	RpcPostDiskName: SwiftRpcCommonDiskName,
	DtoDiskName:     SwiftDtoDiskName,
}

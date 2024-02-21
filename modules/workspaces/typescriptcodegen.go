package workspaces

import (
	"strings"

	typescripttpl "github.com/torabian/fireback/modules/workspaces/codegen/typescript"
	tsinclude "github.com/torabian/fireback/modules/workspaces/codegen/typescript/include"
)

func TsPrimitve(primitive string) string {
	switch primitive {
	case "string", "text":
		return "string"
	case "int64", "int32", "int":
		return "number"
	case "float64", "float32", "float":
		return "number"
	case "bool":
		return "boolean"
	case "double":
		return "number"
	default:
		return "unknown"
	}
}

func TsCalcJsonField(field *Module2Field) string {
	t := []string{}

	if len(field.Matches) > 0 {

		for _, match := range field.Matches {
			if match.Dto != nil {
				t = append(t, match.PublicName())
			}
		}

	} else {
		t = append(t, "any")
	}

	return strings.Join(t, "|")
}

func TsComputedField(field *Module2Field) string {
	switch field.Type {
	case "string", "text":
		return "string"
	case "one":
		return field.Target
	case "enum":
		items := []string{}
		for _, item := range field.OfType {
			items = append(items, "\""+item.Key+"\"")
		}
		return strings.Join(items, " | ")
	case "json":
		return TsCalcJsonField(field)
	case "many2many":
		return field.Target + "[]"
	case "array":
		return field.PublicName() + "[]"
	case "arrayP":
		return TsPrimitve(field.Primitive) + "[]"
	case "html":
		return "string"
	case "int64", "int32", "int":
		return "number"
	case "float64", "float32", "float":
		return "number"
	case "bool":
		return "boolean"
	case "Timestamp", "datenano":
		return "string"
	case "date":
		return "Date"
	case "double":
		return "number"
	case "object":
		return field.PublicName()
	default:
		return field.Type
	}
}

func TypeScriptEntityDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Entity.ts"
}

func TypeScriptDtoDiskName(x *Module2DtoBase) string {
	return ToUpper(x.Name) + "Dto.ts"
}

func TypeScriptFormDiskName(x *Module2Entity) string {
	return ToUpper(x.Name) + "Form.ts"
}

func TypeScriptRpcQueryDiskName(x *Module2Action) string {
	return "use" + ToUpper(x.GetFuncName()) + ".ts"
}

func TsActionDiskName(moduleName string) string {
	return ToUpper(moduleName) + "ActionsDto.ts"
}

var TypeScriptGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName:            "TypeScript",
	ComputeField:            TsComputedField,
	Templates:               typescripttpl.TypeScriptTpl,
	IncludeDirectory:        &tsinclude.TypeScriptInclude,
	EntityGeneratorTemplate: "TypescriptEntity.tpl",
	DtoGeneratorTemplate:    "TypescriptDto.tpl",
	ActionDiskName:          TsActionDiskName,
	ActionGeneratorTemplate: "TsActionDto.tpl",
	RpcQueryDiskName:        TypeScriptRpcQueryDiskName,
	RpcDeleteDiskName:       TypeScriptRpcQueryDiskName,
	RpcPatchDiskName:        TypeScriptRpcQueryDiskName,
	RpcPostDiskName:         TypeScriptRpcQueryDiskName,
	RpcGetOneDiskName:       TypeScriptRpcQueryDiskName,
	RpcPatchBulkDiskName:    TypeScriptRpcQueryDiskName,
	RpcReactiveDiskName:     TypeScriptRpcQueryDiskName,
	RpcQuery:                "RpcQuery.tpl",
	RpcDelete:               "RpcDelete.tpl",
	RpcPatchBulk:            "RpcPatchBulk.tpl",
	RpcPatch:                "RpcPatch.tpl",
	RpcGetOne:               "RpcGetOne.tpl",
	RpcPost:                 "RpcPost.tpl",
	RpcReactive:             "RpcReactive.tpl",
	EntityDiskName:          TypeScriptEntityDiskName,
	DtoDiskName:             TypeScriptDtoDiskName,
}

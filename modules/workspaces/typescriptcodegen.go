package workspaces

import (
	"strings"

	typescripttpl "github.com/torabian/fireback/modules/workspaces/codegen/typescript"
	angularinclude "github.com/torabian/fireback/modules/workspaces/codegen/typescript/angular-include"
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

func TsComputedField(field *Module2Field, isWorkspace bool) string {
	switch field.Type {
	case "string", "text":
		return "string"
	case "one":
		return field.Target
	case "daterange":
		return "any"
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
		return "string"
		// return field.Type
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

// Angular is class based, makes sense to put the actions of specific entity into a class
func AngularRouteGroupFileName(x *Module2Action) string {
	return ToUpper(x.Group) + "Rpc.ts"
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
	RpcQuery:                "ReactRpcQuery.tpl",
	RpcDelete:               "ReactRpcDelete.tpl",
	RpcPatchBulk:            "ReactRpcPatchBulk.tpl",
	RpcPatch:                "ReactRpcPatch.tpl",
	RpcGetOne:               "ReactRpcGetOne.tpl",
	RpcPost:                 "ReactRpcPost.tpl",
	RpcReactive:             "ReactRpcReactive.tpl",
	EntityDiskName:          TypeScriptEntityDiskName,
	DtoDiskName:             TypeScriptDtoDiskName,
}

var AngularGenCatalog CodeGenCatalog = CodeGenCatalog{
	LanguageName:            "TypeScript",
	ComputeField:            TsComputedField,
	Templates:               typescripttpl.TypeScriptTpl,
	IncludeDirectory:        &angularinclude.AngularInclude,
	EntityGeneratorTemplate: "TypescriptEntity.tpl",
	DtoGeneratorTemplate:    "TypescriptDto.tpl",
	ActionDiskName:          TsActionDiskName,
	ActionGeneratorTemplate: "TsActionDto.tpl",

	EntityClassTemplate:  "AngularEntity.tpl",
	EntityClassDiskName:  AngularRouteGroupFileName,
	RpcQueryDiskName:     TypeScriptRpcQueryDiskName,
	RpcDeleteDiskName:    TypeScriptRpcQueryDiskName,
	RpcPatchDiskName:     TypeScriptRpcQueryDiskName,
	RpcPostDiskName:      TypeScriptRpcQueryDiskName,
	RpcGetOneDiskName:    TypeScriptRpcQueryDiskName,
	RpcPatchBulkDiskName: TypeScriptRpcQueryDiskName,
	RpcReactiveDiskName:  TypeScriptRpcQueryDiskName,
	RpcQuery:             "AngularRpcQuery.tpl",
	RpcDelete:            "AngularRpcDelete.tpl",
	RpcPatchBulk:         "AngularRpcPatchBulk.tpl",
	RpcPatch:             "AngularRpcPatch.tpl",
	RpcGetOne:            "AngularRpcGetOne.tpl",
	RpcPost:              "AngularRpcPost.tpl",
	RpcReactive:          "AngularRpcReactive.tpl",
	EntityDiskName:       TypeScriptEntityDiskName,
	DtoDiskName:          TypeScriptDtoDiskName,
}

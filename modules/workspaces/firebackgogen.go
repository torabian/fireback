package workspaces

import (
	"os/exec"
	"strings"

	firebackgo "github.com/torabian/fireback/modules/workspaces/codegen/firebackgo"
)

func GolangComputedField(field *Module3Field, isWorkspace bool) string {
	prefix := ""
	if !isWorkspace {
		prefix = "workspaces."
	}
	switch field.Type {

	case "string", "text", "html", "enum":
		return "string"
	case "string?", "text?", "html?", "enum?":
		return prefix + "String"
	case "duration?":
		return prefix + "Duration"

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
	case "int64", "int32", "int", "float64", "float32", "bool":
		return field.Type
	case "int64?", "int32?", "int?", "float64?", "float32?", "bool?":
		return prefix + strings.ReplaceAll(ToUpper(field.Type), "?", "")
	case "Timestamp":
		return "*string"
	case "datenano":
		return "int64"
	case "boolean":
		return "*bool"
	case "double":
		return "*float64"
	case "object", "embed":
		return field.PublicName()
	case "json":
		return "JSON"
	case "date":
		return prefix + "XDate"
	case "datetime":
		return "*" + prefix + "XDateTime"
	default:
		// Let's return string anyway for unknown types, because it's gonna
		// prevent the code generate to break to some extend
		return "*string"
		// return "*" + field.Type
	}
}

func GoEntityDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + "Entity.dyno.go"
}
func GoEntityExtensionDiskName(x *Module3Entity) string {
	return ToUpper(x.Name) + "Entity.go"
}
func GoActionDiskName(moduleName string) string {
	return ToUpper(moduleName) + "CustomActions.dyno.go"
}

func GoDtoDiskName(x *Module3Dto) string {
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
	Prettier: func(name string) {
		if strings.HasSuffix(name, ".go") {
			RunExecCmd(GO_BIN, []string{"fmt", name})
		}
	},
	DtoDiskName: GoDtoDiskName,
}

var GO_BIN string = ""

func init() {
	goBinary, err := exec.LookPath("go")
	if err != nil {
		return
	}
	GO_BIN = goBinary
}

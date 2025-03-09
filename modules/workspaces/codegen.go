/**
* Code generation for fireback projects.
* This is |NOT| golang code generation, rather generates code after compile time for different
* Platforms
 */
package workspaces

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	reflect "reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/gin-gonic/gin"
	"github.com/swaggest/openapi-go/openapi3"
	firebackgo "github.com/torabian/fireback/modules/workspaces/codegen/firebackgo"
	"gopkg.in/yaml.v2"
)

var FIELD_TYPE_ARRAY string = "array"
var FIELD_TYPE_ARRAYP string = "arrayP"
var FIELD_TYPE_JSON string = "json"
var FIELD_TYPE_ONE string = "one"
var FIELD_TYPE_DATE string = "date"
var FIELD_TYPE_MANY2MANY string = "many2many"
var FIELD_TYPE_OBJECT string = "object"
var FIELD_TYPE_EMBED string = "embed"
var FIELD_TYPE_ENUM string = "enum"
var FIELD_TYPE_COMPUTED string = "computed"
var FIELD_TYPE_TEXT string = "text"
var FIELD_TYPE_STRING string = "string"
var FIELD_TYPE_ANY string = "any"
var ROUTE_FORMAT_DELETE string = "DELETE_DSL"
var ROUTE_FORMAT_QUERY string = "QUERY"
var ROUTE_FORMAT_POST string = "POST_ONE"
var ROUTE_FORMAT_GET_ONE string = "GET_ONE"
var ROUTE_FORMAT_REACTIVE string = "REACTIVE"
var ROUTE_FORMAT_PATCH string = "PATCH_ONE"
var ROUTE_FORMAT_PATCH_BULK string = "PATCH_BULK"

// We extend things in the front-end from a base entity, this is not necessary anymore
// But let it be for reference, maybe later we want add something to it (?)
var DefaultFields []*Module3Field = []*Module3Field{
	// {Name: "visibility", Type: "string"},
	// {Name: "parentId", Type: "string"},
	// {Name: "linkerId", Type: "string"},
	// {Name: "workspaceId", Type: "string"},
	// {Name: "linkedId", Type: "string"},
	// {Name: "uniqueId", Type: "string"},
	// {Name: "userId", Type: "string"},
	// // {Name: "rank", Type: "float64"},
	// {Name: "updated", Type: "float64"},
	// {Name: "created", Type: "float64"},
	// {Name: "createdFormatted", Type: "string"},
	// {Name: "updatedFormatted", Type: "string"},
}

func (x *Module3FieldMatch) PublicName() string {
	if x.Dto == nil {
		return ""
	}

	return ToUpper(*x.Dto) + "Dto"
}

func (x *Module3ActionBody) EntityPure() string {
	if x.Entity != "" {
		return strings.ReplaceAll(x.Entity, "Entity", "")
	}

	return ""
}

func (x *Module3Action) DashedName() string {
	return ToSnakeCase(x.Name)
}

func (x *Module3) Upper() string {
	if x.Name == "" {
		return ToUpper(x.Name)
	}
	return ToUpper(x.Name)
}

func (x *Module3Field) ComputedCliName() string {
	return strings.ReplaceAll(ToSnakeCase((x.Name)), "_", "-")
}

func (x *Module3Field) DistinctBy() string {
	return ""
}

func (x *Module3Action) ComputedCliName() string {
	if x.CliName != "" {
		return x.CliName
	}
	return strings.ReplaceAll(ToSnakeCase((x.Name)), "_", "-")
}

func (x *Module3ActionBody) Template() string {
	return "-"
}

func (x *Module3Action) Template() string {
	return "-"
	// return x.DashedName()
}

func (x *Module3Action) ComputeRequestEntity() string {
	if x.In == nil {
		return ""
	}
	if x.In.Entity != "" {
		return "&" + x.In.Entity + "{}"
	}
	if x.In.Dto != "" {
		return "&" + x.In.Dto + "{}"
	}

	if len(x.In.Fields) > 0 {
		return "&" + x.Upper() + "ActionReqDto{}"
	}

	return ""
}

func (x *Module3Action) ComputeRequestEntityS() string {
	if x.In.Entity != "" {
		return x.In.Entity
	}
	if x.In.Dto != "" {
		return x.In.Dto
	}

	if len(x.In.Fields) > 0 {
		return x.Upper() + "ActionReqDto"
	}

	return ""
}

func (x *Module3Action) FormatComputed() string {
	if x.Method == "REACTIVE" || x.Method == "reactive" {
		return "REACTIVE"
	}

	if strings.ToLower(x.Method) == "query" {
		return "QUERY"
	}

	if x.Format != "" {
		return strings.ToUpper(x.Format)
	}

	if x.Method == "get" {
		return "GET_ONE"
	}

	return "POST_ONE"
}

func (x *Module3Action) DashedPluralName() string {

	pluralize2 := pluralize.NewClient()
	return strings.ReplaceAll(ToSnakeCase(pluralize2.Plural(x.Name)), "_", "-")
}
func (x *Module3Action) ComputedUrl() string {

	if x.Url != "" {
		return x.Url
	}

	return "/" + x.DashedPluralName()
}

func (x *Module3Action) MethodAllUpper() string {

	return strings.ToUpper(x.Method)
}

func (x *Module3Enum) KeyUpper() string {
	return ToUpper(x.Key)
}

func (x *Module3Action) ComputeResponseEntity() string {
	if x.Out == nil {
		return `string("")`
	}
	if x.Out.Entity != "" {
		return "&" + x.Out.Entity + "{}"
	}
	if x.Out.Dto != "" {
		return "&" + x.Out.Dto + "{}"
	}

	if len(x.Out.Fields) > 0 {
		return "&" + x.Upper() + "ActionResDto{}"
	}

	return "&OkayResponseDto{}"
}
func (x *Module3Action) ComputeResponseEntityS() string {
	if x.Out == nil {
		return ``
	}

	if x.Out.Entity != "" {
		return x.Out.Entity
	}
	if x.Out.Dto != "" {
		return x.Out.Dto
	}

	if len(x.Out.Fields) > 0 {
		return x.Upper() + "ActionResDto"
	}

	return "OkayResponseDto"
}

type TypeScriptGenContext struct {
	IncludeStaticField      bool
	IncludeFirebackDef      bool
	IncludeStaticNavigation bool
}

type CodeGenContext struct {
	// Where the content will be exported to
	Path string

	// Used in golang which indicates the relative path
	RelativePath    string
	RelativePathDot string
	EntityPath      string

	// Location of the sdk which entities will be there
	// such as @/fireback/sdk, ...
	UiSdkDir string

	// Type of the generation, (swift, etc)
	Type string

	// Generation
	OpenApiFile string

	// Generation
	GofModuleName string

	FirebackUIDir string

	// Only build specific modules
	Modules []string

	// Only build specific modules
	ModulesOnDisk []string

	// Set of functions, and meta data to generate for that specific target
	Catalog CodeGenCatalog

	NoCache bool

	Ts TypeScriptGenContext
}

// Used in go structs to configurate the gorm tag
func (x *Module3Field) ComputedExcerptSize() int {
	if x.ExcerptSize == 0 {
		return 100
	}

	return x.ExcerptSize
}

func (x *Module3Field) ComputedGormTag() string {
	if x.Gorm != "" {
		return x.Gorm
	}

	if x.Type == FIELD_TYPE_TEXT {
		return "text"
	}

	if x.Type == FIELD_TYPE_MANY2MANY {
		return "many2many:" + x.BelongingEntityName + "_" + x.PrivateName() + ";foreignKey:UniqueId;references:UniqueId"
	}

	if x.Type == FIELD_TYPE_EMBED {
		return "embedded"
	}

	if x.Type == FIELD_TYPE_ARRAY || x.Type == FIELD_TYPE_OBJECT {
		return "foreignKey:LinkerId;references:UniqueId;constraint:OnDelete:CASCADE"
	}

	if x.Type == FIELD_TYPE_ONE {
		return "foreignKey:" + x.PublicName() + "Id;references:UniqueId"
	}

	if x.Type == FIELD_TYPE_ANY {
		return "-"
	}

	return ""
}

func (x *Module3Field) ComputedSqlTag() string {
	if x.Sql != "" {
		return x.Sql
	}

	if x.Type == FIELD_TYPE_COMPUTED || x.Type == FIELD_TYPE_ANY {
		return "-"
	}

	return ""
}
func (x *Module3Field) PrivateNameUnderscore() string {
	return "_" + x.Name
}
func (x *Module3Field) UpperPlural() string {
	pluralize2 := pluralize.NewClient()
	return ToUpper(pluralize2.Plural(x.Name))
}

func CalcAllPolyglotEntities(m []*Module3Field) []string {
	items := []string{}
	for _, item := range m {
		if item.Translate {
			items = append(items, item.Name)
		}

		if item.Type == FIELD_TYPE_OBJECT || item.Type == FIELD_TYPE_ARRAY {
			items = append(items, CalcAllPolyglotEntities(item.Fields)...)
		}
	}

	return items
}

func (x *Module3Entity) CompletePolyglotFields() []string {
	return CalcAllPolyglotEntities(x.Fields)
}

func (x *Module3Entity) DistinctByAllUpper() string {
	return strings.ToUpper(x.DistinctBy)
}

func (x *Module3Entity) DistinctByAllLower() string {
	return strings.ToLower(x.DistinctBy)
}

func (x *Module3Entity) ComputedCliName() string {
	if x.CliName != "" {
		return x.CliName
	}
	return strings.ToLower(x.Name)
}

func (x *Module3Entity) ComputedCliDescription() string {
	if x.Description != "" {
		return x.Description
	}

	return ""

}

func (x *Module3Field) ComputedCliDescription() string {

	if x.Type == FIELD_TYPE_ENUM {
		items := []string{}
		for _, item := range x.OfType {
			items = append(items, "'"+item.Key+"'")
		}
		return "One of: " + strings.Join(items, ", ")
	}

	if x.Description != "" {
		return x.Description
	}

	return x.Name
}

func (x *Module3Field) CompleteFields() []*Module3Field {
	return x.Fields
}

func (x *Module3Field) IsRequired() bool {
	return strings.Contains(x.Validate, "required")
}

// On cli level for interactive access
func (x *Module3Field) IsRecommended() bool {
	return x.Recommended
}

func (x *Module3Field) PrivateName() string {
	return x.Name
}

func (x *Module3) PublicName() string {
	return ToUpper(x.Name)
}
func (x *Module3Field) PublicName() string {
	return ToUpper(x.Name)
}
func (x *Module3Field) AllUpper() string {
	return strings.ToUpper(CamelCaseToWordsUnderlined(x.Name))
}
func (x *Module3Field) UnderscoreName() string {
	return strings.ToLower(CamelCaseToWordsUnderlined(x.Name))
}
func (x *Module3Entity) UnderscoreName() string {
	return strings.ToLower(CamelCaseToWordsUnderlined(x.Name))
}
func (x *Module3Field) TargetWithModule() string {
	if x.Module != "" {
		return x.Module + "." + ToUpper(x.Target)
	}
	return ToUpper(x.Target)
}
func (x *Module3Field) TargetWithoutEntity() string {
	return strings.ReplaceAll(x.Target, "Entity", "")
}
func (x *Module3Field) TargetWithoutEntityPlural() string {
	pluralize2 := pluralize.NewClient()
	return ToUpper(pluralize2.Plural(x.TargetWithoutEntity()))
}
func (x *Module3Field) TargetWithModuleWithoutEntity() string {
	return strings.ReplaceAll(x.TargetWithModule(), "Entity", "")
}
func (x *Module3Field) TargetWithModuleWithoutEntityPluralize() string {
	pluralize2 := pluralize.NewClient()
	return pluralize2.Plural(strings.ReplaceAll(x.TargetWithModule(), "Entity", ""))
}
func (x *Module3Action) Upper() string {
	return ToUpper(x.Name)
}

func (x *Module3Action) ActionReqDto() string {
	if x.In == nil {
		return ""
	}

	if x.In.Entity != "" {
		return "*" + x.In.Entity
	}
	if x.In.Dto != "" {
		return "*" + x.In.Dto
	}

	if len(x.In.Fields) > 0 {
		return "*" + x.Upper() + "ActionReqDto"
	}

	return "nil"
}

func (x *Module3Action) ActionResDto() string {
	if x.Out == nil {
		return "string"
	}
	prefix := ""
	if strings.ToLower(x.Format) == "query" || strings.ToLower(x.Method) == "query" {
		prefix = "[]"
	}

	if x.Out.Primitive != "" {
		return x.Out.Primitive
	}

	if x.Out.Entity != "" {
		return prefix + "*" + x.Out.Entity
	}
	if x.Out.Dto != "" {
		return prefix + "*" + x.Out.Dto
	}

	if len(x.Out.Fields) > 0 {
		return prefix + "*" + x.Upper() + "ActionResDto"
	}

	return "nil"
}

func (x *Module3Field) DashedName() string {
	return ToSnakeCase(x.Name)
}
func (x *Module3Field) DefaultEmptySymbol() string {
	switch x.Type {
	case "string", "text":
		return `""`
	case "one":
		return x.Target + "?"
	case "array", "many2many":
		return "nil"
	case "int64", "int32", "int":
		return "0"
	case "float64", "float32", "float":
		return "0.0"
	case "bool":
		return "FALSE"
	case "Timestamp", "datenano":
		return `""`
	case "double":
		return "0.0"
	default:
		return "nil"
	}

}
func (x *Module3Entity) CompleteFields() []*Module3Field {
	var all []*Module3Field = []*Module3Field{}
	all = append(all,
		x.Fields...,
	)
	all = append(all,
		DefaultFields...,
	)

	return all
}
func (x *Module3Remote) CompleteFields() []*Module3Field {
	var all []*Module3Field = []*Module3Field{}
	all = append(all,
		x.Query...,
	)

	return all
}

func (x *Module3Dto) CompleteFields() []*Module3Field {
	var all []*Module3Field = []*Module3Field{}
	all = append(all,
		x.Fields...,
	)
	all = append(all,
		DefaultFields...,
	)

	return all
}

// Represents things which are needed to be imported
type ImportDependencyStrategy struct {
	Items []string
	Path  string
}

type ImportMapRow struct {
	Items []string
}
type ImportMap map[string]*ImportMapRow

func (x ImportMap) IsValid() bool {
	return true
	// Now if this exists, it means the import is wrong :D
	return x["..//"] == nil
}

func mergeImportMaps(maps ...ImportMap) ImportMap {
	result := make(ImportMap)

	for _, m := range maps {
		for key, row := range m {
			if _, ok := result[key]; !ok {
				result[key] = &ImportMapRow{}
			}
			// Merge items and make them unique
			itemSet := make(map[string]struct{})
			for _, item := range append(result[key].Items, row.Items...) {
				itemSet[item] = struct{}{}
			}
			result[key].Items = make([]string, 0, len(itemSet))
			for item := range itemSet {
				result[key].Items = append(result[key].Items, item)
			}
		}
	}

	return result
}

var generatorHash map[string]string = map[string]string{}

func GetMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}

func WriteFileGen(ctx *CodeGenContext, name string, data []byte, perm os.FileMode) error {

	gen, okay := generatorHash[name]

	newGen := GetMD5Hash(data)
	if okay && gen == newGen {
		return nil
	}

	generatorHash[name] = newGen

	err := os.WriteFile(name, data, perm)

	if err != nil {
		return err
	}

	if ctx.Catalog.Prettier != nil {
		ctx.Catalog.Prettier(name)
	}

	return nil
}

type ReconfigDto struct {
	ProjectSource  string
	NewProjectName string
	BinaryName     string
	Description    string
	Languages      []string
}

/**
* Renames the project into another project, useful for fast bootstrap
**/
func Reconfig(scheme ReconfigDto) error {

	// Change the make file
	{
		data, _ := ioutil.ReadFile("Makefile")
		d := string(data)
		d = strings.ReplaceAll(d, scheme.ProjectSource, scheme.NewProjectName)
		os.WriteFile("Makefile", []byte(d), 0644)
	}
	{
		data, _ := ioutil.ReadFile(filepath.Join(".vscode", "tasks.json"))
		d := string(data)
		d = strings.ReplaceAll(d, ` f `, scheme.BinaryName+` `)
		os.WriteFile(filepath.Join(".vscode", "tasks.json"), []byte(d), 0644)
	}
	{
		data, _ := ioutil.ReadFile(filepath.Join("cmd", "fireback", "main.go"))
		d := string(data)
		d = strings.ReplaceAll(d, `var PRODUCT_NAMESPACENAME = "fireback"`, `var PRODUCT_NAMESPACENAME = "`+scheme.NewProjectName+`"`)
		d = strings.ReplaceAll(d, `var PRODUCT_DESCRIPTION = "Fireback core microservice"`, `var PRODUCT_DESCRIPTION = "`+scheme.Description+`"`)
		os.WriteFile(filepath.Join("cmd", "fireback", "main.go"), []byte(d), 0644)
	}

	{
		data, _ := ioutil.ReadFile(filepath.Join("cmd", "fireback", "Makefile"))
		d := string(data)
		d = strings.ReplaceAll(d, "project = fireback", "project = "+scheme.NewProjectName)
		d = strings.ReplaceAll(d, "projectBinary = f", "projectBinary = "+scheme.BinaryName)
		os.WriteFile(filepath.Join("cmd", "fireback", "Makefile"), []byte(d), 0644)
	}

	{
		err := os.Rename(
			filepath.Join("cmd", "fireback"),
			filepath.Join("cmd", scheme.NewProjectName+"-server"),
		)
		if err != nil {
			return err
		}
	}

	return nil

}

func GenerateRpcCodeOnDisk(ctx *CodeGenContext, route *Module3Action, exportDir string) {

	content, exportPath := GenerateRpcCodeString(ctx, route, exportDir)

	if exportPath != "" {

		err3 := WriteFileGen(ctx, exportPath, content, 0644)
		if err3 != nil {
			fmt.Println("Error on writing content on RPC:", exportPath, err3)
		}
	}
}

func GenerateRpcCodeString(ctx *CodeGenContext, route *Module3Action, exportDir string) ([]byte, string) {
	method := strings.ToUpper(route.Method)
	if (route.Format == ROUTE_FORMAT_POST || method == "POST") && ctx.Catalog.RpcPost != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcPost)
		if err != nil {
			log.Fatalln("Generating post call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcPostDiskName(route))
		return EscapeLines(data), exportPath
	}

	if route.Format == ROUTE_FORMAT_QUERY && ctx.Catalog.RpcQuery != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcQuery)
		if err != nil {
			log.Fatalln("Generating rpc query call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcQueryDiskName(route))
		return EscapeLines(data), exportPath
	}
	if (route.Format == ROUTE_FORMAT_DELETE || method == "DELETE") && ctx.Catalog.RpcDelete != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcDelete)
		if err != nil {
			log.Fatalln("Generating delete rpc call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcDeleteDiskName(route))
		return EscapeLines(data), exportPath
	}
	if (route.Format == ROUTE_FORMAT_PATCH || method == "PATCH") && ctx.Catalog.RpcPatch != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcPatch)
		if err != nil {
			log.Fatalln("Generating rpc patch call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcPatchDiskName(route))
		return EscapeLines(data), exportPath
	}
	if route.Format == ROUTE_FORMAT_PATCH_BULK && ctx.Catalog.RpcPatchBulk != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcPatchBulk)
		if err != nil {
			log.Fatalln("Generating rpc patch call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcPatchBulkDiskName(route))
		return EscapeLines(data), exportPath
	}
	if (route.Format == ROUTE_FORMAT_REACTIVE || method == ROUTE_FORMAT_REACTIVE) && ctx.Catalog.RpcReactive != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcReactive)
		if err != nil {
			log.Fatalln("Generating rpc reactive call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcReactiveDiskName(route))
		return EscapeLines(data), exportPath
	}
	if route.Format == ROUTE_FORMAT_GET_ONE && ctx.Catalog.RpcGetOne != "" {
		data, err := route.RenderTemplate(ctx, ctx.Catalog.Templates, ctx.Catalog.RpcGetOne)
		if err != nil {
			log.Fatalln("Generating rpc get one call error", err)
			return []byte(""), ""
		}
		exportPath := filepath.Join(exportDir, ctx.Catalog.RpcGetOneDiskName(route))
		return EscapeLines(data), exportPath
	}
	return []byte(""), ""
}

func GenMoveIncludeDir(ctx *CodeGenContext) {

	// Move the include directory
	if ctx.Catalog.IncludeDirectory != nil {
		files, err := GetAllFilenames(ctx.Catalog.IncludeDirectory, "")
		if err != nil {
			log.Fatalln(err)
		}

		for _, file := range files {
			exportFile := filepath.Join(ctx.Path, file)
			content, err4 := ReadEmbedFileContent(ctx.Catalog.IncludeDirectory, file)
			if err4 != nil {
				log.Fatalln(err)
			}
			dir := filepath.Dir(exportFile)
			os.MkdirAll(dir, os.ModePerm)
			err3 := WriteFileGen(ctx, exportFile, []byte(content), 0644)
			if err3 != nil {
				log.Fatalln(err)
			}
		}
	}

}

func GenGetModules(xapp *FirebackApp, ctx *CodeGenContext) []*Module3 {

	j := []*Module3{}
	if len(ctx.ModulesOnDisk) > 0 && ctx.ModulesOnDisk[0] != "" {
		j = append(j, ListModule3FilesFromDisk(ctx.ModulesOnDisk)...)
	} else {
		j = append(j, ListModule3Files(xapp)...)
	}

	return j
}

func writeGenCache(ctx *CodeGenContext) {
	cacheMapFile := filepath.Join(ctx.Path, "cachemap.json")

	cache, _ := json.MarshalIndent(generatorHash, "", "  ")
	os.WriteFile(cacheMapFile, cache, 0644)
}

func ReadGenCache(ctx *CodeGenContext) {
	cacheMapFile := filepath.Join(ctx.Path, "cachemap.json")
	if !ctx.NoCache {
		ReadJsonFile(cacheMapFile, &generatorHash)

	}
}

func GenRpcCode(ctx *CodeGenContext, modules []*ModuleProvider, mode string) {

	for _, item := range modules {
		m := item.ToModule3()
		exportDir := filepath.Join(ctx.Path, "modules", item.Name)
		if m.Namespace != "" {
			exportDir = filepath.Join(ctx.Path, "modules", item.Namespace, item.Name)
		}
		perr := os.MkdirAll(exportDir, os.ModePerm)
		if perr != nil {
			log.Fatalln(perr)
		}

		actions := item.Actions

		if item.ActionsBundle != nil {
			actions = append(actions, item.ActionsBundle.Actions)
		}

		for _, bundle := range item.EntityBundles {
			actions = append(actions, bundle.Actions)
		}

		// Reading the custom actions is missing

		for _, actions := range actions {
			if mode == "disk" {
				for _, action := range actions {
					action.RootModule = &m
					GenerateRpcCodeOnDisk(ctx, &action, exportDir)
				}
			}

			if mode == "class" {
				if len(actions) > 0 {

					importMap := ImportMap{}
					fileToWrite := filepath.Join(exportDir, ctx.Catalog.EntityClassDiskName(&actions[0]))
					content := []byte("")

					for _, action := range actions {

						action.RootModule = &m
						maps := actions[0].ImportDependecies()
						importMap = mergeImportMaps(importMap, maps)
						partial, _ := GenerateRpcCodeString(ctx, &action, exportDir)

						content = append(content, []byte("\r\n")...)
						content = append(content, EscapeLines(partial)...)
						content = append(content, []byte("\r\n")...)
					}

					render, err2 := RenderRpcGroupClassBody(
						ctx,
						ctx.Catalog.Templates,
						ctx.Catalog.EntityClassTemplate,
						content,
						importMap,
					)
					if err2 != nil {
						log.Fatalln("Error on rendering the content", err2)
					}

					err3 := WriteFileGen(ctx, fileToWrite, (render), 0644)
					if err3 != nil {
						log.Fatalln("Error on writing content for Actions class file:", fileToWrite, err3)
					}
				}

			}
		}
	}
}

func GenRpcCodeExternal(ctx *CodeGenContext, modules []*Module3, mode string) {

	for _, item := range modules {

		exportDir := filepath.Join(ctx.Path, item.Name)
		// perr := os.MkdirAll(exportDir, os.ModePerm)
		// if perr != nil {
		// 	log.Fatalln(perr)
		// }

		if mode == "disk" {
			for _, action := range item.Actions {
				action.RootModule = item
				GenerateRpcCodeOnDisk(ctx, action, exportDir)
			}
		}

		if mode == "class" {
			if len(item.Actions) > 0 {

				importMap := ImportMap{}
				fileToWrite := filepath.Join(exportDir, ctx.Catalog.EntityClassDiskName(item.Actions[0]))
				content := []byte("")

				for _, action := range item.Actions {
					action.RootModule = item
					maps := action.ImportDependecies()
					importMap = mergeImportMaps(importMap, maps)
					partial, _ := GenerateRpcCodeString(ctx, action, exportDir)

					content = append(content, []byte("\r\n")...)
					content = append(content, EscapeLines(partial)...)
					content = append(content, []byte("\r\n")...)
				}

				render, err2 := RenderRpcGroupClassBody(
					ctx,
					ctx.Catalog.Templates,
					ctx.Catalog.EntityClassTemplate,
					content,
					importMap,
				)
				if err2 != nil {
					log.Fatalln("Error on rendering the content", err2)
				}

				err3 := WriteFileGen(ctx, fileToWrite, (render), 0644)
				if err3 != nil {
					log.Fatalln("Error on writing content for Actions class file:", fileToWrite, err3)
				}
			}

		}
	}
}

// For openapi3, we create FirebackApp not from internal, rather an external json
func GetOpenAPiXServer(ctx *CodeGenContext) (*FirebackApp, []*Module3) {
	data, _ := ioutil.ReadFile(ctx.OpenApiFile)
	s := openapi3.Spec{}

	if err := s.UnmarshalJSON(data); err != nil {
		log.Fatal("Converting json content:", err)
	}

	virtualModule := OpenApiToFireback(s)
	modules := []*Module3{
		virtualModule,
	}
	app := &FirebackApp{
		Modules: []*ModuleProvider{
			{
				Actions: [][]Module3Action{
					// virtualModule.Actions,
				},
			},
		},
	}
	return app, modules
}

func extractModuleName(goModFilePath string) (string, error) {
	file, err := os.Open(goModFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("module name not found in %s", goModFilePath)
}

/*
* Creates a yml, and also a golang module file in modules directory
 */
func NewGoNativeModule(name string, dist string, autoImport string) error {

	folderName := strings.ToLower(dist)
	args := gin.H{
		"path": folderName,
		"Name": ToUpper(name),
		"name": name,
	}

	goModule, err := CompileString(&firebackgo.FbGoTpl, "GoModule.tpl", args)
	if err != nil {
		return err
	}

	goModuleDef, err := CompileString(&firebackgo.FbGoTpl, "GoModuleDef.tpl", args)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(folderName), os.ModePerm); err != nil {
		return err
	}

	moduleName := filepath.Join(folderName, ToUpper(name)+"Module.go")

	if err := os.WriteFile(moduleName, []byte(goModule), 0644); err != nil {
		return err
	}

	yamlName := filepath.Join(folderName, ToUpper(name)+"Module3.yml")
	if err := os.WriteFile(yamlName, []byte(goModuleDef), 0644); err != nil {
		return err
	}

	MAGIC_LINE := "// do not remove this comment line - it's used by fireback to append new modules"

	if autoImport != "" {

		if !Exists("go.mod") {
			return nil
		}

		moduleName, err0 := extractModuleName("go.mod")
		if err0 != nil {
			return nil
		}

		if data, err := os.ReadFile(autoImport); err != nil {
			return err
		} else {
			j := string(data)
			m := strings.ReplaceAll(
				j,
				MAGIC_LINE,
				MAGIC_LINE+"\r\n\t\t"+ToLower(name)+"."+ToUpper(name)+"ModuleSetup(nil),",
			)

			m = strings.ReplaceAll(
				m,
				"import (",
				"import ("+"\r\n\t\""+moduleName+"/"+folderName+"\"\r\n",
			)
			os.WriteFile(autoImport, []byte(m), 0644)
		}
	}

	return nil
}

func CompileString(fs *embed.FS, fname string, params gin.H) (string, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer

	err = t.ExecuteTemplate(&tpl, fname, params)

	return tpl.String(), err
}

func RunCodeGen(xapp *FirebackApp, ctx *CodeGenContext) error {
	// For now, I have separate work flow for binary and external definitions
	// Which might need to be fixed.

	if len(ctx.ModulesOnDisk) > 0 || ctx.OpenApiFile != "" {
		return RunCodeGenExternal(ctx)
	}

	os.MkdirAll(ctx.Path, os.ModePerm)
	ReadGenCache(ctx)
	GenMoveIncludeDir(ctx)

	app := xapp
	modules := GenGetModules(xapp, ctx)

	// Generate the classes, definitions, structs
	for _, item := range modules {
		item.Generate(ctx)
	}

	mode := "disk"
	if ctx.Catalog.EntityClassTemplate != "" {
		mode = "class"
	}
	GenRpcCode(ctx, app.Modules, mode)

	writeGenCache(ctx)

	return nil
}

// This is used for pure items from definition, not internal binary
func RunCodeGenExternal(ctx *CodeGenContext) error {
	os.MkdirAll(ctx.Path, os.ModePerm)
	ReadGenCache(ctx)
	GenMoveIncludeDir(ctx)

	modules := []*Module3{}

	for _, def := range ctx.ModulesOnDisk {
		var m Module3
		ReadYamlFile[Module3](def, &m)
		modules = append(modules, &m)
	}

	if ctx.OpenApiFile != "" {
		_, modules = GetOpenAPiXServer(ctx)
	}

	// Generate the classes, definitions, structs
	for _, item := range modules {
		// if len(ctx.Modules) > 0 && !Contains(ctx.Modules, item.Name) {
		// 	continue
		// }
		item.Generate(ctx)
	}

	mode := "disk"
	if ctx.Catalog.EntityClassTemplate != "" {
		mode = "class"
	}

	GenRpcCodeExternal(ctx, modules, mode)

	writeGenCache(ctx)

	return nil
}

func FindModuleByPath(xapp *FirebackApp, modulePath string) *Module3 {
	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if defFile, err := GetSeederFilenames(item.Definitions, ""); err != nil {
			fmt.Println(err.Error())
		} else {

			for _, path := range defFile {
				var mod2 Module3
				ReadYamlFileEmbed(item.Definitions, path, &mod2)
				ComputeMacros(&mod2)
				if mod2.Name == modulePath {
					return &mod2
				}
			}
		}

	}

	return nil
}

func ListModule3WithEntities(xapp *FirebackApp) []string {
	items := []string{}
	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if defFile, err := GetSeederFilenames(item.Definitions, ""); err != nil {
			fmt.Println(err.Error())
		} else {

			for _, path := range defFile {
				var mod2 Module3
				ReadYamlFileEmbed(item.Definitions, path, &mod2)
				ComputeMacros(&mod2)
				for _, entity := range mod2.Entities {
					pf := mod2.Name
					if mod2.Namespace != "" {
						pf = mod2.Namespace
					}
					items = append(items, pf+"."+entity.Name)
				}
			}
		}
	}

	return items
}

func FindModule3Entity(xapp *FirebackApp, address string) *Module3Entity {
	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if defFile, err := GetSeederFilenames(item.Definitions, ""); err != nil {
			fmt.Println(err.Error())
		} else {

			for _, path := range defFile {
				var mod2 Module3
				ReadYamlFileEmbed(item.Definitions, path, &mod2)
				ComputeMacros(&mod2)
				for _, entity := range mod2.Entities {
					if mod2.Name+"."+entity.Name == address {
						return &entity
					}
				}
			}
		}
	}

	return nil
}

func ListModule3Files(xapp *FirebackApp) []*Module3 {
	items := []*Module3{}
	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if defFile, err := GetSeederFilenames(item.Definitions, ""); err != nil {
			fmt.Println(err.Error())
		} else {

			for _, path := range defFile {
				var mod2 Module3
				ReadYamlFileEmbed(item.Definitions, path, &mod2)
				ComputeMacros(&mod2)
				items = append(items, &mod2)
			}
		}

	}

	return items
}

func ListModule3Entities(xapp *FirebackApp, modulePath string) []Module3Entity {
	module := FindModuleByPath(xapp, modulePath)
	if module == nil {
		return []Module3Entity{}
	}
	return module.Entities
}

func ListModule3FilesFromDisk(files []string) []*Module3 {
	items := []*Module3{}

	for _, item := range files {
		var mod2 Module3
		ReadYamlFile(item, &mod2)
		ComputeMacros(&mod2)
		items = append(items, &mod2)
	}

	return items
}

func ComputeComplexGormField(entity *Module3Entity, fields []*Module3Field) {
	if len(fields) == 0 {
		return
	}

	for _, field := range fields {
		field.BelongingEntityName = entity.Name

		if field.Type == FIELD_TYPE_OBJECT || field.Type == FIELD_TYPE_EMBED || field.Type == FIELD_TYPE_ARRAY {
			ComputeComplexGormField(entity, field.Fields)
		}
	}
}

func ComputeFieldTypes(fields []*Module3Field, isWorkspace bool, fn func(field *Module3Field, isWorkspace bool) string,
) {
	if len(fields) == 0 {
		return
	}

	for _, field := range fields {
		field.ComputedType = fn(field, isWorkspace)

		if field.Type == FIELD_TYPE_OBJECT || field.Type == FIELD_TYPE_EMBED || field.Type == FIELD_TYPE_ARRAY {
			ComputeFieldTypes(field.Fields, isWorkspace, fn)
		}
	}
}

// Removes pointers
func ComputeFieldTypesAbsolute(fields []*Module3Field, isWorkspace bool, fn func(field *Module3Field, isWorkspace bool) string,
) {
	if len(fields) == 0 {
		return
	}

	for _, field := range fields {
		field.ComputedType = strings.ReplaceAll(fn(field, isWorkspace), "*", "")

		if field.Type == FIELD_TYPE_OBJECT || field.Type == FIELD_TYPE_EMBED || field.Type == FIELD_TYPE_ARRAY {
			ComputeFieldTypesAbsolute(field.Fields, isWorkspace, fn)
		}
	}
}

type CodeGenCatalog struct {
	LanguageName            string
	ComputeField            func(field *Module3Field, isWorkspace bool) string
	EntityDiskName          func(x *Module3Entity) string
	EntityExtensionDiskName func(x *Module3Entity) string
	EntityClassDiskName     func(x *Module3Action) string

	// A function that does a post formatting on the file, when it's saved
	Prettier func(string)

	// Maybe only useful for C/C++
	EntityHeaderDiskName          func(x *Module3Entity) string
	EntityHeaderExtensionDiskName func(x *Module3Entity) string

	ActionDiskName func(modulename string) string

	// When you want each action to be written in separate file
	SingleActionDiskName             func(action *Module3Action, modulename string) string
	DtoDiskName                      func(x *Module3Dto) string
	FormDiskName                     func(x *Module3Entity) string
	RpcQueryDiskName                 func(x *Module3Action) string
	RpcDeleteDiskName                func(x *Module3Action) string
	RpcGetOneDiskName                func(x *Module3Action) string
	RpcPatchBulkDiskName             func(x *Module3Action) string
	RpcReactiveDiskName              func(x *Module3Action) string
	RpcPatchDiskName                 func(x *Module3Action) string
	RpcPostDiskName                  func(x *Module3Action) string
	Templates                        embed.FS
	IncludeDirectory                 *embed.FS
	Partials                         *embed.FS
	EntityClassTemplate              string
	EntityGeneratorTemplate          string
	EntityExtensionGeneratorTemplate string
	DtoGeneratorTemplate             string
	CteSqlTemplate                   string
	ActionGeneratorTemplate          string
	FormGeneratorTemplate            string
	RpcQuery                         string
	RpcPatch                         string
	RpcPatchBulk                     string
	RpcGetOne                        string
	RpcReactive                      string
	RpcPost                          string
	RpcDelete                        string
}

func EscapeLines(data []byte) []byte {
	d := string(data)

	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)

	return []byte(re.ReplaceAllString(d, ""))

}

// This function would enable or disable functionality in a module based on
// some predefined condition in Fireback itself, then on the module global config
// before making anycomputes
func FeatureSetMacro(x *Module3) {
	// Implement such logic here
}

func ComputeMacros(x *Module3) {
	for _, item := range x.Macros {
		if item.Using == "eav" {
			EavMacro(item, x)
		}
	}
}

func GofModuleGenerationFlow(x *Module3, ctx *CodeGenContext, exportDir string, isWorkspace bool) {

	for _, remote := range x.Remotes {

		if remote.Query != nil {
			ComputeFieldTypesAbsolute(remote.Query, isWorkspace, ctx.Catalog.ComputeField)
		}
		if remote.Out != nil {
			ComputeFieldTypesAbsolute(remote.Out.Fields, isWorkspace, ctx.Catalog.ComputeField)
		}

		if remote.In != nil {
			ComputeFieldTypesAbsolute(remote.In.Fields, isWorkspace, ctx.Catalog.ComputeField)
		}

	}

	for _, task := range x.Tasks {
		if task.In != nil {
			ComputeFieldTypesAbsolute(task.In.Fields, isWorkspace, ctx.Catalog.ComputeField)
		}

	}

	exportPath := filepath.Join(exportDir, ToUpper(x.Name)+"Module.dyno.go")
	data, err := x.RenderTemplate(ctx, ctx.Catalog.Templates, "GoModuleDyno.tpl")
	if err != nil {
		fmt.Println("Error on module dyno file:", exportPath, err)
	}

	if !HasQueries(ctx.Path) {
		err7 := CreateQueriesDirectory(ctx.Path)
		if err7 != nil {
			fmt.Println("Error on entity mocks directory generation:", err7)
		}
	}

	if !HasMetasFolder(ctx.Path) {
		if err9 := CreateMetaDirectory(ctx.Path); err9 != nil {
			fmt.Errorf("Error on writing content: err: %v", err9)
		}
	}

	err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
	if err3 != nil {
		fmt.Println("Error on writing content:", exportPath, err3)
	}
}

/**
*	Common code generator
**/
func (x *Module3) Generate(ctx *CodeGenContext) {
	isWorkspace := x.MetaWorkspace
	os.MkdirAll(ctx.Path, os.ModePerm)
	exportDir := filepath.Join(ctx.Path)

	if x.Namespace != "" {
		exportDir = filepath.Join(ctx.Path, x.Namespace)
	}

	// This is nasty. Let's change the entry point of the compiler soon for different
	// languages. Initial idea was to add as many as targets, now it's only go/react
	if ctx.Catalog.LanguageName != "FirebackGo" {
		exportDir = filepath.Join(ctx.Path, "modules", x.Name)
		if x.Namespace != "" {
			exportDir = filepath.Join(ctx.Path, "modules", x.Namespace, x.Name)
		}
	}

	perr := os.MkdirAll(exportDir, os.ModePerm)
	if perr != nil {
		log.Fatalln(perr)
	}

	ComputeMacros(x)

	if ctx.Catalog.LanguageName == "FirebackGo" {
		GofModuleGenerationFlow(x, ctx, exportDir, isWorkspace)
	}

	for _, query := range x.Queries {

		exportPath := filepath.Join(exportDir, "queries", ToUpper(query.Name)+".vsql")

		data := []byte("--- Generated VSQL file. Do not modify directly, check yaml definition instead \r\n" + query.Query)

		err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
		if err3 != nil {
			fmt.Println("Error on writing vsql:", exportPath, err3)
		}

	}

	for _, dto := range x.Dto {

		// Computing field types is important for target writter.
		ComputeFieldTypes(dto.CompleteFields(), isWorkspace, ctx.Catalog.ComputeField)

		// Step 0: Generate the Entity
		if ctx.Catalog.DtoGeneratorTemplate != "" {
			exportPath := filepath.Join(exportDir, ctx.Catalog.DtoDiskName(&dto))

			data, err := dto.RenderTemplate(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.DtoGeneratorTemplate,
				x,
			)
			if err != nil {
				fmt.Println("Error on dto generation:", err)
			} else {
				err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content:", exportPath, err3)
				}
			}
		}

	}

	// Render actions specific dtos if they have their own
	if ctx.Catalog.ActionDiskName != nil {

		exportPath := filepath.Join(exportDir, ctx.Catalog.ActionDiskName(x.Name))

		if len(x.Actions) > 0 || ctx.Catalog.LanguageName == "FirebackGo" {

			data, err := x.RenderActions(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.ActionGeneratorTemplate,
			)

			if err != nil {
				fmt.Println("Error on action generation:", err)
			} else {
				err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content:", exportPath, err3)
				}
			}

			if ctx.Catalog.LanguageName == "FirebackGo" {
				// Let's also check if the actions files are there, if not skip them.
				for _, action := range x.Actions {
					actionImplementationFile := filepath.Join(exportDir, action.Upper()+"Action.go")
					hasFile := Exists(actionImplementationFile)

					if !hasFile {

						wsPrefix := "workspaces."
						if x.MetaWorkspace {
							wsPrefix = ""
							isWorkspace = true
						}

						params := gin.H{
							"m":        x,
							"a":        action,
							"wsprefix": wsPrefix,
						}
						data, err5 := getActionTemplate(params)
						if err5 != nil {
							fmt.Println("Error creating action default template:", exportPath, err5)
						}
						err4 := WriteFileGen(ctx, actionImplementationFile, EscapeLines(data), 0644)
						if err4 != nil {
							fmt.Println("Error creating action default template:", exportPath, err4)
						}
					}
				}
			}
		}
	}

	if ctx.Catalog.SingleActionDiskName != nil {
		if len(x.Actions) > 0 {

			for _, action := range x.Actions {
				exportPath := filepath.Join(exportDir, ctx.Catalog.SingleActionDiskName(action, x.Name))

				data, err := action.Render(
					x,
					ctx,
					ctx.Catalog.Templates,
					ctx.Catalog.ActionGeneratorTemplate,
				)
				if err != nil {
					fmt.Println("Error on action generation:", err)
				} else {
					err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
					if err3 != nil {
						fmt.Println("Error on writing content:", exportPath, err3)
					}
				}
			}
		}
	}

	for _, entity := range x.Entities {

		// Computing field types is important for target writter.
		ComputeFieldTypes(entity.CompleteFields(), isWorkspace, ctx.Catalog.ComputeField)
		ComputeComplexGormField(&entity, entity.CompleteFields())

		entityAddress := filepath.Join(exportDir, ctx.Catalog.EntityDiskName(&entity))

		if ctx.Catalog.LanguageName == "FirebackGo" {
			if !HasSeeders(ctx.Path, &entity) {
				err7 := CreateSeederDirectory(ctx.Path, &entity)
				if err7 != nil {
					fmt.Println("Error on entity seeders directory generation:", err7)
				}
			}

			if !HasMocks(ctx.Path, &entity) {
				err7 := CreateMockDirectory(ctx.Path, &entity)
				if err7 != nil {
					fmt.Println("Error on entity mocks directory generation:", err7)
				}
			}
			if !HasMocks(ctx.Path, &entity) {
				err7 := CreateMockDirectory(ctx.Path, &entity)
				if err7 != nil {
					fmt.Println("Error on entity mocks directory generation:", err7)
				}
			}
		}

		if ctx.Catalog.EntityHeaderDiskName != nil {
			exportPath := filepath.Join(exportDir, ctx.Catalog.EntityHeaderDiskName(&entity))

			params := gin.H{
				"implementation": "skip",
			}

			data, err := entity.RenderTemplate(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.EntityGeneratorTemplate,
				x,
				params,
			)

			if err != nil {
				fmt.Println("Error on entity extension generation:", err)
			} else {
				err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
				if err3 != nil {
					fmt.Println("Error on writing content:", exportPath, err3)
				}
			}
		}

		if ctx.Catalog.EntityExtensionGeneratorTemplate != "" {
			exportPath := filepath.Join(exportDir, ctx.Catalog.EntityExtensionDiskName(&entity))

			// We only render the extension, if this entity is first time being created
			if !Exists(exportPath) && !Exists(entityAddress) {
				data, err := entity.RenderTemplate(
					ctx,
					ctx.Catalog.Templates,
					ctx.Catalog.EntityExtensionGeneratorTemplate,
					x,
					nil,
				)
				if err != nil {
					fmt.Println("Error on entity extension generation:", err)
				} else {
					err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
					if err3 != nil {
						fmt.Println("Error on writing content:", exportPath, err3)
					}
				}

			}

		}

		// Step 0: Generate the Entity
		if ctx.Catalog.EntityGeneratorTemplate != "" {

			data, err := entity.RenderTemplate(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.EntityGeneratorTemplate,
				x,
				nil,
			)
			if err != nil {
				log.Fatalln("Error on entity generation:", err)
			} else {
				err3 := WriteFileGen(ctx, entityAddress, EscapeLines(data), 0644)
				if err3 != nil {
					log.Fatalln("Error on writing content:", entityAddress, err3)
				}
			}
		}

		// Step 0: Cte SQL Render
		if entity.Cte && ctx.Catalog.CteSqlTemplate != "" {

			{
				os.MkdirAll(filepath.Join(exportDir, "queries"), os.ModePerm)
				CreateQueryIndex(ctx.Path)
				exportPath := filepath.Join(exportDir, "queries", entity.Upper()+"Cte.vsql")
				data, err := entity.RenderCteSqlTemplate(
					ctx,
					ctx.Catalog.Templates,
					ctx.Catalog.CteSqlTemplate,
					x,
					"sql",
				)
				if err != nil {
					log.Fatalln("Error on cte sql generation:", err)
				} else {
					err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
					if err3 != nil {
						log.Fatalln("Error on writing content:", exportPath, err3)
					}
				}
			}
		}

		// Step 1: Generate the form, (fields, sections, inputs, etc...)
		if ctx.Catalog.FormGeneratorTemplate != "" {
			exportPath := filepath.Join(exportDir, ctx.Catalog.FormDiskName(&entity))

			data, err := entity.RenderTemplate(
				ctx,
				ctx.Catalog.Templates,
				ctx.Catalog.FormGeneratorTemplate,
				x,
				nil,
			)
			if err != nil {
				log.Fatalln("Error on UI generation:", err)
			} else {
				err3 := WriteFileGen(ctx, exportPath, EscapeLines(data), 0644)
				if err3 != nil {
					log.Fatalln("Error on writing content:", exportPath, err3)
				}
			}
		}

	}
}

func ToUpper(t string) string {
	if t == "" {
		return ""
	}
	return strings.ToUpper(t[0:1]) + t[1:]
}
func ToLower(t string) string {
	if t == "" {
		return ""
	}
	return strings.ToLower(t[0:1]) + t[1:]
}

func (x *Module3Entity) EntityName() string {
	return ToUpper(x.Name) + "Entity"
}

func (x *Module3Entity) ObjectName() string {
	return x.EntityName()
}

func getActionTemplate(data interface{}) ([]byte, error) {
	tmplStr := `package {{ .m.Name }}

	func init() {
		// Override the implementation with our actual code.
		{{ .a.Upper }}ActionImp = {{ .a.Upper }}Action
	}
	
	func {{ .a.Upper }}Action(
      {{ if .a.ComputeRequestEntity }}{{ if ne .a.ActionReqDto "nil" }}req {{ .a.ActionReqDto }}, {{ end}}{{end}}
      q {{ .wsprefix }}QueryDSL) ({{ .a.ActionResDto }},
      {{ if (eq .a.FormatComputed "QUERY") }} *{{ .wsprefix }}QueryResultMeta, {{ end }}
      *{{ .wsprefix }}IError,
    ) {
		// Implement the logic here.
		
		return {{ if (eq .a.ActionResDto "string")}} "" {{ else }} nil {{ end }}, {{ if (eq .a.FormatComputed "QUERY") }} nil, {{ end }} nil
	}
`

	tmpl, err := template.New("greeting").Parse(tmplStr)
	if err != nil {
		return []byte{}, err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		return []byte{}, err
	}

	return result.Bytes(), nil
}

func (x *Module3Dto) ObjectName() string {
	return x.DtoName()
}

func (x *Module3Entity) HasExtendedQuer() bool {
	return len(x.Queries) > 0 && Contains(x.Queries, "extended")
}

func (x *Module3Entity) EventCreated() string {
	return x.AllUpper() + "_EVENT_CREATED"
}

func (x *Module3Entity) EventUpdated() string {
	return x.AllUpper() + "_EVENT_UPDATED"
}

func (x *Module3Entity) AllUpper() string {
	return strings.ToUpper(CamelCaseToWordsUnderlined(x.Name))
}

func (x *Module3Entity) AllLower() string {
	return strings.ToLower(CamelCaseToWordsDashed(x.Name))
}

func (x *Module3Entity) HumanReadable() string {
	return strings.ToLower(CamelCaseToWords(x.Name))
}

func (x *Module3) AllUpper() string {
	return strings.ToUpper(CamelCaseToWordsUnderlined(x.Name))
}

func (x *Module3) AllLower() string {
	return strings.ToLower(CamelCaseToWordsDashed(x.Name))
}

func (x *Module3Permission) AllUpper() string {
	return strings.ToUpper(CamelCaseToWordsUnderlined(x.Key))
}

func (x *Module3Permission) AllLower() string {
	return strings.ToLower(CamelCaseToWordsDashed(x.Key))
}

func (x *Module3Entity) PolyglotName() string {
	return x.EntityName() + "Polyglot"
}
func (x *Module3Entity) HasTranslations() bool {
	for _, field := range x.Fields {
		if field.Translate {
			return true
		}
	}
	return false
}
func (x *Module3Entity) DefinitionJson() string {
	data, _ := json.MarshalIndent(x, "", "  ")
	return string(data)
}
func (x *Module3Action) Json() string {
	data, _ := json.MarshalIndent(x, "", "  ")
	return string(data)
}
func (x *Module3Dto) DtoName() string {
	return ToUpper(x.Name) + "Dto"
}
func (x *Module3Dto) DefinitionJson() string {
	data, _ := json.MarshalIndent(x, "", "  ")
	return string(data)
}
func (x *Module3Entity) Upper() string {
	return ToUpper(x.Name)
}

func (x *Module3Dto) Upper() string {
	return ToUpper(x.Name)
}

/**
*	In module2 definitions we do have array and object fields,
*	which need to be stored in database in their own table
*	so we need to create those classes, etc...
**/
func GetArrayOrObjectFieldsFlatten(depth int, parentType string, depthName string, fields []*Module3Field, ctx *CodeGenContext, isWorkspace bool) []*Module3Field {
	items := []*Module3Field{}
	if len(fields) == 0 {
		return items
	}

	for _, item := range fields {
		if item.Type == FIELD_TYPE_EMBED {
			item.IsVirtualObject = true
		}

		if item.Type != FIELD_TYPE_OBJECT && item.Type != FIELD_TYPE_ARRAY && item.Type != FIELD_TYPE_EMBED {
			item.ComputedType = ctx.Catalog.ComputeField(item, isWorkspace)
			continue
		} else {
			item.LinkedTo = depthName
			if depth == 0 {
				item.LinkedTo += parentType
			}
			item.ComputedType = depthName + ctx.Catalog.ComputeField(item, isWorkspace)

		}

		item.FullName = depthName + item.PublicName()
		items = append(items, item)
		items = append(items, GetArrayOrObjectFieldsFlatten(
			depth+1,
			parentType,
			item.FullName,
			item.Fields, ctx, isWorkspace)...,
		)
	}

	return items
}

func ChildItems(x *Module3Entity, ctx *CodeGenContext, isWorkspace bool) []*Module3Field {

	return GetArrayOrObjectFieldsFlatten(0, "Entity", x.Upper(), x.Fields, ctx, isWorkspace)

}

func ChildItemsActionIn(x *Module3Action, ctx *CodeGenContext, isWorkspace bool) []*Module3Field {
	if x.In == nil {
		return []*Module3Field{}
	}
	return GetArrayOrObjectFieldsFlatten(0, "Entity", x.Upper()+"ReqDto", x.In.Fields, ctx, isWorkspace)
}

func ChildItemsActionOut(x *Module3Action, ctx *CodeGenContext, isWorkspace bool) []*Module3Field {
	if x.Out == nil {
		return []*Module3Field{}
	}
	return GetArrayOrObjectFieldsFlatten(0, "Entity", x.Upper()+"ResDto", x.Out.Fields, ctx, isWorkspace)
}

func ChildItemsCommon(prefix string, x []*Module3Field, ctx *CodeGenContext, isWorkspace bool) []*Module3Field {

	return GetArrayOrObjectFieldsFlatten(0, "Entity", prefix, x, ctx, isWorkspace)

}

func ChildItemsX(x *Module3Dto, ctx *CodeGenContext, isWorkspace bool) []*Module3Field {

	return GetArrayOrObjectFieldsFlatten(0, "Dto", x.Upper(), x.Fields, ctx, isWorkspace)

}

func FieldsChildren(
	fields []*Module3Field,
	ctx *CodeGenContext,
	isWorkspace bool,
	name string,
	affix string,
) []*Module3Field {
	return GetArrayOrObjectFieldsFlatten(0, affix, name, fields, ctx, isWorkspace)
}

func (x *Module3Field) PluralName() string {

	pluralize2 := pluralize.NewClient()
	return pluralize2.Plural(x.Name)
}

func (x *Module3Entity) PluralNameUpper() string {

	pluralize2 := pluralize.NewClient()
	return ToUpper(pluralize2.Plural(x.Name))
}

func (x *Module3Entity) PluralName() string {

	pluralize2 := pluralize.NewClient()
	return pluralize2.Plural(x.Name)
}

func (x *Module3Dto) Template() string {
	return x.DashedName()
}

func (x *Module3Dto) Templates() string {
	pluralize2 := pluralize.NewClient()
	return strings.ToLower(pluralize2.Plural(x.Name))
}

func (x *Module3Dto) TemplatesLower() string {
	return x.PluralNameLower()
}

func (x *Module3Entity) Template() string {
	return x.DashedName()
}

func (x *Module3Entity) Templates() string {
	pluralize2 := pluralize.NewClient()
	return strings.ToLower(pluralize2.Plural(x.Name))
}

func (x *Module3Entity) TemplatesLower() string {
	return x.PluralNameLower()
}

func (x *Module3Entity) PluralNameLower() string {

	pluralize2 := pluralize.NewClient()
	return strings.ToLower(pluralize2.Plural(x.Name))
}

func (x *Module3Dto) PluralNameLower() string {

	pluralize2 := pluralize.NewClient()
	return strings.ToLower(pluralize2.Plural(x.Name))
}

func (x *Module3Entity) DashedPluralName() string {

	pluralize2 := pluralize.NewClient()
	return strings.ReplaceAll(ToSnakeCase(pluralize2.Plural(x.Name)), "_", "-")
}

func (x *Module3Entity) TableName() string {

	return ToSnakeCase((x.Name))
}

func (x *Module3Dto) DashedName() string {
	return strings.ReplaceAll(ToSnakeCase(x.Name), "_", "-")
}
func (x *Module3Entity) DashedName() string {
	return strings.ReplaceAll(ToSnakeCase(x.Name), "_", "-")
}

func (x *Module3Entity) FormName() string {
	return ToUpper(x.Name) + "Form"
}

func ImportDependecies(fields []*Module3Field) []ImportDependencyStrategy {
	items := []ImportDependencyStrategy{
		{
			Path:  "../../core/definitions",
			Items: []string{"BaseDto", "BaseEntity"},
		},
	}

	for _, field := range fields {

		if field.Type == FIELD_TYPE_JSON {

			if len(field.Matches) > 0 {
				for _, item := range field.Matches {
					if item.Dto != nil {
						items = append(items, ImportDependencyStrategy{
							Items: []string{item.PublicName()},

							// This loads only the same directory items
							Path: "./" + item.PublicName(),
						})
					}
				}
			}
		}

		if field.Type == FIELD_TYPE_ARRAY || field.Type == FIELD_TYPE_OBJECT || field.Type == FIELD_TYPE_EMBED {
			items = append(items, ImportDependecies(field.Fields)...)
		}

		if field.Type != FIELD_TYPE_ONE && field.Type != FIELD_TYPE_MANY2MANY {
			continue
		}

		// Computed path is based on the idea that every thing is only one folder
		// inside. This might required revise if we want to build for C++, etc

		computedPath := ""
		if field.Module != "" {
			computedPath = "../" + field.Module + "/" + field.Target
		} else {
			if field.RootClass != "" {
				computedPath = "./" + field.RootClass

			} else {
				computedPath = "./" + field.Target

			}
		}

		items = append(items, ImportDependencyStrategy{
			Items: []string{field.Target},
			Path:  computedPath,
		})
	}

	return items
}

func (x *Module3) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

func (x *Module3) Yaml() string {
	if x != nil {
		str, _ := yaml.Marshal(x)
		return (string(str))
	}
	return ""
}

/*
Give me golang function, which gets and string, and convert it into camelCase,
and output should only consist of numbers and english alphabet lowercase and uppercase (only).
*/

func ToCamelCaseClean(input string) string {
	splitBySpecial := regexp.MustCompile("[^A-Za-z0-9]+")
	words := splitBySpecial.Split(input, -1)

	var result string
	for _, word := range words {
		// Convert each word to camel case
		word = strings.ToLower(word)
		if word == "" {
			continue
		}
		word = strings.ToUpper(word[0:1]) + word[1:]
		result += word
	}

	// Remove non-alphanumeric characters
	nonAlphaNumeric := regexp.MustCompile("[^A-Za-z0-9]")
	result = nonAlphaNumeric.ReplaceAllString(result, "")

	return ToLower(result)
}

func ImportGoDependencies(fields []*Module3Field, importGroupPrefix string) []ImportDependencyStrategy {
	items := []ImportDependencyStrategy{}

	for _, field := range fields {
		actualPrefix := importGroupPrefix

		if field.Provider != "" {
			actualPrefix = field.Provider

			// this is magical keyword resolves to the fireback on github
			if field.Provider == "fireback" {
				actualPrefix = "github.com/torabian/fireback/modules/"
			}

			if !strings.HasSuffix(actualPrefix, "/") {
				actualPrefix = actualPrefix + "/"
			}
		}
		// if field.Type == FIELD_TYPE_JSON {
		// 	items = append(items, ImportDependencyStrategy{
		// 		Items: []string{field.Target},
		// 	})
		// }

		if field.Type == FIELD_TYPE_ARRAY || field.Type == FIELD_TYPE_OBJECT || field.Type == FIELD_TYPE_EMBED {

			items = append(items, ImportGoDependencies(field.Fields, actualPrefix)...)
		}

		if field.Type != FIELD_TYPE_ONE && field.Type != FIELD_TYPE_MANY2MANY {
			continue
		}

		if field.Type == FIELD_TYPE_DATE {
			items = append(items, ImportDependencyStrategy{
				Path: "time",
			})
		}

		if field.Module != "" && field.Module != "workspaces" {
			items = append(items, ImportDependencyStrategy{
				Items: []string{field.Target},
				Path:  actualPrefix + field.Module,
			})

		}
	}

	return items
}

func (x Module3Action) RequestRootObjectName() string {
	reqValue := reflect.ValueOf(x.RequestEntity)
	if reqValue.MethodByName("RootObjectName").IsValid() {
		res := reqValue.MethodByName("RootObjectName").Call(nil)

		if len(res) > 0 {
			return res[0].String()
		}
	}
	return ""
}

func (x Module3Action) ResponseRootObjectName() string {
	reqValue := reflect.ValueOf(x.ResponseEntity)
	if reqValue.MethodByName("RootObjectName").IsValid() {
		res := reqValue.MethodByName("RootObjectName").Call(nil)

		if len(res) > 0 {
			return res[0].String()
		}
	}
	return ""
}

func (x Module3Action) RequestExample() string {
	if x.RequestEntity == nil {
		return ""
	}

	reqValue := reflect.ValueOf(x.RequestEntity)
	if reqValue.MethodByName("Seeder").IsValid() {
		res := reqValue.MethodByName("Seeder").Call(nil)

		if len(res) > 0 {
			return res[0].String()
		}
	}
	return ""
}

// In Typescript, we put all req/res dtos related to custom actions
// into a single file, this is why we need to get the correct name
func TsObjectName(objectName string, rootObjectName string) string {

	if strings.Contains(objectName, "ActionReq") || strings.Contains(objectName, "ActionRes") {
		return ToUpper(rootObjectName) + "ActionsDto"
	}

	return rootObjectName
}

func (x ImportMap) AppendImportMapRow(key string, row *ImportMapRow) {
	if x[key] == nil {
		x[key] = row
	} else {
		// Complete this logic over time, if necessary
		for _, item := range row.Items {
			if !Contains(x[key].Items, item) {
				x[key].Items = append(x[key].Items, item)
			}
		}
	}
}

func (x Module3Action) ImportDependecies() ImportMap {
	m := ImportMap{}

	if x.RequestEntity != nil {
		meta := x.RequestEntityMeta()
		fileName := meta.ClassName
		if name := x.RequestRootObjectName(); name != "" {
			fileName = TsObjectName(meta.ClassName, name)
		}

		m.AppendImportMapRow("../"+meta.Module+"/"+fileName, &ImportMapRow{
			Items: []string{meta.ClassName},
		})
	} else if x.RequestEntityComputed() != "" {
		// This is more for external defined files
		m.AppendImportMapRow("./"+x.RequestEntityComputed(), &ImportMapRow{
			Items: []string{x.RequestEntityComputed()},
		})
	}

	if x.ResponseEntity != nil {
		meta := x.ResponseEntityMeta()
		fileName := meta.ClassName
		if name := x.ResponseRootObjectName(); name != "" {
			fileName = TsObjectName(meta.ClassName, name)
		}

		m.AppendImportMapRow("../"+meta.Module+"/"+fileName, &ImportMapRow{
			Items: []string{meta.ClassName},
		})

	} else if x.ResponseEntityComputed() != "" {
		// This is more for external defined files
		m.AppendImportMapRow("./"+x.ResponseEntityComputed(), &ImportMapRow{
			Items: []string{x.ResponseEntityComputed()},
		})
	}

	return m
}

// Converts import strategy into unique map to be ported into the template.
// ImportDependencies might generate duplicate elements, here we make them unique
// or any other last moment changes
func (x *Module3Entity) ImportGroupResolver(prefix string) ImportMap {

	deps := ImportGoDependencies(x.Fields, prefix)

	m := ImportMap{}
	for _, dep := range deps {
		if m[dep.Path] == nil {
			m[dep.Path] = &ImportMapRow{}
		}

		for _, klass := range dep.Items {
			if !Contains(m[dep.Path].Items, klass) && klass != x.EntityName() {
				m[dep.Path].Items = append(m[dep.Path].Items, klass)
			}
		}

	}

	return m

}
func (x *Module3) TsActionsImport() ImportMap {
	m := ImportMap{}

	for _, action := range x.Actions {

		if action.In != nil {

			deps := ImportDependecies(action.In.Fields)

			for _, dep := range deps {
				if m[dep.Path] == nil {
					m[dep.Path] = &ImportMapRow{}
				}

				for _, klass := range dep.Items {
					if !Contains(m[dep.Path].Items, klass) {
						m[dep.Path].Items = append(m[dep.Path].Items, klass)
					}
				}

			}
		}

		if action.Out != nil {
			deps := ImportDependecies(action.Out.Fields)

			for _, dep := range deps {
				if m[dep.Path] == nil {
					m[dep.Path] = &ImportMapRow{}
				}

				for _, klass := range dep.Items {
					if !Contains(m[dep.Path].Items, klass) {
						m[dep.Path].Items = append(m[dep.Path].Items, klass)
					}
				}

			}
		}
	}
	return m

}
func (x *Module3Entity) ImportDependecies() ImportMap {

	deps := ImportDependecies(x.Fields)

	m := ImportMap{}
	for _, dep := range deps {
		if m[dep.Path] == nil {
			m[dep.Path] = &ImportMapRow{}
		}

		for _, klass := range dep.Items {
			if !Contains(m[dep.Path].Items, klass) && klass != x.EntityName() {
				m[dep.Path].Items = append(m[dep.Path].Items, klass)
			}
		}

	}

	return m

}
func (x *Module3Dto) ImportDependecies() ImportMap {

	deps := ImportDependecies(x.Fields)

	m := ImportMap{}
	for _, dep := range deps {
		if m[dep.Path] == nil {
			m[dep.Path] = &ImportMapRow{}
		}

		for _, klass := range dep.Items {
			if !Contains(m[dep.Path].Items, klass) && klass != x.DtoName() {
				m[dep.Path].Items = append(m[dep.Path].Items, klass)
			}
		}

	}

	return m

}

func HasSeeders(dir string, entity *Module3Entity) bool {
	checkee := filepath.Join(dir, "seeders", entity.Upper())
	if _, err := os.Stat(checkee); !os.IsNotExist(err) {
		return true
	}

	return false
}
func CreateSeederDirectory(dir string, entity *Module3Entity) error {
	basePath := filepath.Join(dir, "seeders", entity.Upper())
	indexPath := filepath.Join(basePath, "index.go")

	// Create the directory, and add index.go into it
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return err
	}
	var indexContent = `package seeders

import "embed"

//go:embed *
var ViewsFs embed.FS
`

	return os.WriteFile(indexPath, []byte(indexContent), 0644)
}
func CreateMetaDirectory(dir string) error {
	basePath := filepath.Join(dir, "metas")
	indexPath := filepath.Join(basePath, "index.go")

	fmt.Println("Creating:", indexPath, basePath)

	// Create the directory, and add index.go into it
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return err
	}
	var indexContent = `package metas

import "embed"

//go:embed *
var MetaFs embed.FS

`

	return os.WriteFile(indexPath, []byte(indexContent), 0644)
}

func CreateQueryIndex(dir string) error {

	basePath := filepath.Join(dir, "queries")
	indexPath := filepath.Join(basePath, "index.go")

	if Exists(indexPath) {
		return nil
	}

	var indexContent = `package queries

import "embed"

//go:embed *
var QueriesFs embed.FS
`

	return os.WriteFile(indexPath, []byte(indexContent), 0644)

}

func HasMetasFolder(dir string) bool {
	checkee := filepath.Join(dir, "metas")

	if _, err := os.Stat(checkee); !os.IsNotExist(err) {

		return true
	}

	return false
}

func HasMetas(dir string, entity *Module3Entity) bool {
	checkee := filepath.Join(dir, "seeders", entity.Upper())

	if _, err := os.Stat(checkee); !os.IsNotExist(err) {

		return true
	}

	return false
}

func HasMocks(dir string, entity *Module3Entity) bool {
	mocks := filepath.Join(dir, "mocks", entity.Upper())

	if _, err := os.Stat(mocks); !os.IsNotExist(err) {
		return true
	}
	return false
}

func HasQueries(dir string) bool {
	mocks := filepath.Join(dir, "queries")

	if _, err := os.Stat(mocks); !os.IsNotExist(err) {
		return true
	}
	return false
}

func CreateMockDirectory(dir string, entity *Module3Entity) error {
	basePath := filepath.Join(dir, "mocks", entity.Upper())
	indexPath := filepath.Join(basePath, "index.go")

	// Create the directory, and add index.go into it
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return err
	}
	var indexContent = `package mocks

import "embed"

//go:embed *
var ViewsFs embed.FS
`

	return os.WriteFile(indexPath, []byte(indexContent), 0644)
}
func CreateQueriesDirectory(dir string) error {
	basePath := filepath.Join(dir, "queries")
	indexPath := filepath.Join(basePath, "index.go")

	// Create the directory, and add index.go into it
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return err
	}
	var indexContent = `package queries

import "embed"

//go:embed *
var QueriesFs embed.FS

`

	return os.WriteFile(indexPath, []byte(indexContent), 0644)
}

func generateRange(start, end int) []int {
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = i + start
	}
	return result
}

func SafeIndex(slice []interface{}, index int) bool {
	if index < 0 || index >= len(slice) {
		return false
	}
	return true
}

func EscapeDoubleQuotes(input string) string {
	return strings.ReplaceAll(input, `"`, `\"`)
}

func goComment(comment string) string {
	// Escape problematic characters and split into lines
	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		lines[i] = "// " + strings.ReplaceAll(line, "*/", "* /") // Escape `*/`
	}
	return strings.Join(lines, "\n")
}

func typescriptComment(comment string) string {
	// Escape problematic characters and split into lines
	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "*/", "* /") // Escape `*/`
	}
	return strings.Join(lines, "\n")
}

var CommonMap = template.FuncMap{
	"endsWithDto": func(s string) bool {
		return strings.HasSuffix(s, "Dto")
	},
	"goComment":         goComment,
	"until":             generateRange,
	"typescriptComment": typescriptComment,
	"join":              strings.Join,
	"trim":              strings.TrimSpace,
	"upper":             ToUpper,
	"lower":             ToLower,
	"snakeUpper":        ToSnakeUpper,
	"escape":            EscapeDoubleQuotes,
	"safeIndex":         SafeIndex,
	"hasSuffix":         strings.HasSuffix,
	"arr":               func(els ...any) []any { return els },
	"inc": func(i int) int {
		return i + 1
	},
	"fx": func(fieldName string, depth int) string {
		return fieldName + "[index" + fmt.Sprintf("%v", depth) + "]."
	},
}

func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})

	// Copy values from map1 to merged
	for key, value := range map1 {
		merged[key] = value
	}

	// Copy values from map2 to merged, overwriting existing keys
	for key, value := range map2 {
		merged[key] = value
	}

	return merged
}
func (x *Module3Entity) RenderTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	module *Module3,
	map2 map[string]interface{},
) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	isWorkspace := false
	if module.MetaWorkspace {
		wsPrefix = ""
		isWorkspace = true
	}

	params := gin.H{
		"e":           x,
		"m":           module,
		"fv":          FIREBACK_VERSION,
		"gofModule":   ctx.GofModuleName,
		"ctx":         ctx,
		"children":    ChildItems(x, ctx, isWorkspace),
		"imports":     x.ImportDependecies(),
		"goimports":   x.ImportGroupResolver(ctx.GofModuleName + "/"),
		"javaimports": x.ImportGroupResolver("com.fireback.modules."),
		"wsprefix":    wsPrefix,
		"hasMetas":    HasMetas(ctx.Path, x),
	}

	err = t.ExecuteTemplate(&tpl, fname, mergeMaps(params, map2))

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func (x *Module3Entity) RenderCteSqlTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	module *Module3,
	tp string,
) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	if module.MetaWorkspace {
		wsPrefix = ""
	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"m":        module,
		"wsprefix": wsPrefix,
		"type":     tp,
		"e":        x,
		"fv":       FIREBACK_VERSION,
		"ctx":      ctx,
	})
	if err != nil {
		return []byte{}, err
	}

	y := string(tpl.Bytes())
	y = strings.ReplaceAll(y, "template", x.TableName())
	return []byte(y), nil
}

func (action *Module3Action) Render(
	x *Module3,
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	isWorkspace := false

	wsPrefix := "workspaces."
	if x.MetaWorkspace {
		wsPrefix = ""
		isWorkspace = true
	}

	ComputeFieldTypesAbsolute(action.Query, isWorkspace, ctx.Catalog.ComputeField)

	if action.In != nil && len(action.In.Fields) > 0 {
		ComputeFieldTypes(action.In.Fields, isWorkspace, ctx.Catalog.ComputeField)
	}
	if action.Out != nil && len(action.Out.Fields) > 0 {

		ComputeFieldTypes(action.Out.Fields, isWorkspace, ctx.Catalog.ComputeField)
	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"m":                   x,
		"ctx":                 ctx,
		"woo":                 33,
		"childrenIn":          ChildItemsActionIn(action, ctx, isWorkspace),
		"remoteQueryChildren": RemoteQueryAppend(ctx, x.Remotes, isWorkspace),
		"taskChildren":        RemoteTaskAppend(ctx, x.Tasks, isWorkspace),
		"gofModule":           ctx.GofModuleName,
		"queriesChildren":     QueryAppend(ctx, x.Queries, isWorkspace),
		"childrenOut":         ChildItemsActionOut(action, ctx, isWorkspace),
		"fv":                  FIREBACK_VERSION,
		"wsprefix":            wsPrefix,
		"action":              action,
		"tsactionimports":     x.TsActionsImport(),
	})
	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func (x *Module3) RenderActions(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
) ([]byte, error) {
	isWorkspace := false

	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	if x.MetaWorkspace {
		wsPrefix = ""
		isWorkspace = true
	}

	itemsIn := [][]*Module3Field{}
	itemsOut := [][]*Module3Field{}

	for _, task := range x.Tasks {
		if task.In != nil {
			ComputeFieldTypesAbsolute(task.In.Fields, isWorkspace, ctx.Catalog.ComputeField)
		}

	}

	for _, action := range x.Actions {

		ComputeFieldTypesAbsolute(action.Query, isWorkspace, ctx.Catalog.ComputeField)
		itemsIn = append(itemsIn, ChildItemsActionIn(action, ctx, isWorkspace))
		itemsOut = append(itemsOut, ChildItemsActionOut(action, ctx, isWorkspace))
		// if len(action.In.Fields) > 0 {
		// 	ComputeFieldTypes(action.In.Fields, isWorkspace, ctx.Catalog.ComputeField)
		// }
		// if len(action.Out.Fields) > 0 {
		// 	ComputeFieldTypes(action.Out.Fields, isWorkspace, ctx.Catalog.ComputeField)
		// }

	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"m":                   x,
		"fv":                  FIREBACK_VERSION,
		"wsprefix":            wsPrefix,
		"woo":                 31,
		"childrenIn":          itemsIn,
		"childrenOut":         itemsOut,
		"remoteQueryChildren": RemoteActionsAppend(ctx, x.Actions, isWorkspace),
		"taskChildren":        RemoteTaskAppend(ctx, x.Tasks, isWorkspace),
		"queriesChildren":     QueryAppend(ctx, x.Queries, isWorkspace),
		"gofModule":           ctx.GofModuleName,
		"tsactionimports":     x.TsActionsImport(),
		"ctx":                 ctx,
	})
	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func (x *Module3Entity) GetSqlFields() []string {
	items := []string{
		"template_entities.parent_id",
		"template_entities.visibility",
		"template_entities.updated",
		"template_entities.created",
	}
	for _, field := range x.Fields {
		if field.Type == "object" {
			continue
		}
		if field.Type == "one" {
			items = append(items, "template_entities."+ToSnakeCase(field.Name)+"_id")
		} else {
			items = append(items, "template_entities."+ToSnakeCase(field.Name))
		}

	}

	return items
}

func (x *Module3Entity) GetSqlFieldNames() []string {
	items := []string{"parent_id", "visibility", "updated", "created"}
	for _, field := range x.Fields {
		if field.Type == "object" {
			continue
		}

		if field.Type == "one" {
			items = append(items, ToSnakeCase(field.Name)+"_id")
		} else {
			items = append(items, ToSnakeCase(field.Name))
		}

	}

	return items
}

func (x *Module3Entity) GetSqlFieldNamesAfter() []string {
	items := []string{
		"template_entities_cte.parent_id",
		"template_entities_cte.visibility",
		"template_entities_cte.updated",
		"template_entities_cte.created",
	}
	for _, field := range x.Fields {
		if field.Type == "object" {
			continue
		}

		if field.Type == "one" {
			items = append(items, "template_entities_cte."+ToSnakeCase(field.Name)+"_id\n")
		} else {

			if field.Translate {
				items = append(items, "template_entity_polyglots."+ToSnakeCase(field.Name)+"\n")
			} else {
				items = append(items, "template_entities_cte."+ToSnakeCase(field.Name)+"\n")
			}
		}

	}

	return items
}

func (x *Module3Dto) RenderTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	module *Module3,
) ([]byte, error) {

	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	isWorkspace := false
	if module.MetaWorkspace {
		wsPrefix = ""
		isWorkspace = true
	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"e":        x,
		"children": ChildItemsX(x, ctx, isWorkspace),
		"imports":  x.ImportDependecies(),
		"m":        module,
		"ctx":      ctx,
		"fv":       FIREBACK_VERSION,
		"wsprefix": wsPrefix,
	})

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func RemoteChildrenMapResponse(ctx *CodeGenContext, remotes []*Module3Remote, isWorkspace bool) [][]*Module3Field {
	res := [][]*Module3Field{}

	for _, item := range remotes {
		if item.Out == nil || len(item.Out.Fields) == 0 {
			continue
		}
		res = append(res, ChildItemsCommon(ToUpper(item.Name)+"Res", item.Out.Fields, ctx, isWorkspace))
	}

	return res
}

func RemoteTaskAppend(ctx *CodeGenContext, remotes []*Module3Task, isWorkspace bool) [][]*Module3Field {
	res := [][]*Module3Field{}

	for _, item := range remotes {
		if item.In == nil || len(item.In.Fields) == 0 {
			continue
		}
		res = append(res, ChildItemsCommon(ToUpper(item.Name)+"Task", item.In.Fields, ctx, isWorkspace))
	}

	return res
}

func QueryAppend(ctx *CodeGenContext, queries []*Module3Query, isWorkspace bool) [][]*Module3Field {
	res := [][]*Module3Field{}

	for _, item := range queries {
		if item.Columns == nil || len(item.Columns.Fields) == 0 {
			continue
		}
		res = append(res, ChildItemsCommon(ToUpper(item.Name)+"Query", item.Columns.Fields, ctx, isWorkspace))
	}

	return res
}

func RemoteQueryAppend(ctx *CodeGenContext, remotes []*Module3Remote, isWorkspace bool) [][]*Module3Field {
	res := [][]*Module3Field{}

	for _, item := range remotes {
		if len(item.Query) == 0 {
			continue
		}
		res = append(res, ChildItemsCommon(ToUpper(item.Name)+"Query", item.Query, ctx, isWorkspace))
	}

	return res
}

func extractRouteParams(route string) []*Module3Field {
	// Split the route by '/'
	parts := strings.Split(route, "/")

	var params []*Module3Field

	// Iterate over the parts and extract variables
	for _, part := range parts {
		// Variables start with ':' or '*'
		if strings.HasPrefix(part, ":") || strings.HasPrefix(part, "*") {
			// Remove ':' or '*' and append to params
			param := strings.TrimPrefix(part, ":")
			param = strings.TrimPrefix(param, "*")
			params = append(params, &Module3Field{
				Name: param,
				Type: "string",
			})
		}
	}

	if len(params) > 0 {
		return []*Module3Field{
			{
				Name:   "pathParams",
				Type:   "object",
				Fields: params,
			},
		}
	}

	return params
}

func RemoteActionsAppend(ctx *CodeGenContext, remotes []*Module3Action, isWorkspace bool) [][]*Module3Field {
	res := [][]*Module3Field{}

	for _, item := range remotes {
		item.Query = append(item.Query, extractRouteParams(item.Url)...)
		if len(item.Query) == 0 {
			continue
		}
		res = append(res, ChildItemsCommon(ToUpper(item.Name)+"Query", item.Query, ctx, isWorkspace))
	}

	return res
}

func RemoteChildrenMapRequest(ctx *CodeGenContext, remotes []*Module3Remote, isWorkspace bool) [][]*Module3Field {
	res := [][]*Module3Field{}

	for _, item := range remotes {
		if item.In == nil || len(item.In.Fields) == 0 {
			continue
		}
		res = append(res, ChildItemsCommon(ToUpper(item.Name)+"Req", item.In.Fields, ctx, isWorkspace))
	}

	return res
}

func ActionChildrenMapResponse(ctx *CodeGenContext, actions []*Module3Action, isWorkspace bool) [][]*Module3Field {
	res := [][]*Module3Field{}

	for _, item := range actions {
		if item.Out == nil {
			continue
		}
		if len(item.Out.Fields) == 0 {
			continue
		}
		res = append(res, ChildItemsCommon(ToUpper(item.Name), item.Out.Fields, ctx, isWorkspace))
	}

	return res
}
func ActionChildrenMapRequest(ctx *CodeGenContext, actions []*Module3Action, isWorkspace bool) [][]*Module3Field {
	res := [][]*Module3Field{}
	for _, item := range actions {
		if item.In == nil {
			continue
		}
		if len(item.In.Fields) == 0 {
			continue
		}
		res = append(res, ChildItemsCommon(ToUpper(item.Name), item.In.Fields, ctx, isWorkspace))
	}

	return res
}

func (x *Module3) RenderTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
) ([]byte, error) {

	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	wsPrefix := "workspaces."
	isWorkspace := false
	if x.MetaWorkspace {
		isWorkspace = true
		wsPrefix = ""
	}

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"m":                    x,
		"fv":                   FIREBACK_VERSION,
		"remoteQueryChildren":  RemoteQueryAppend(ctx, x.Remotes, isWorkspace),
		"taskChildren":         RemoteTaskAppend(ctx, x.Tasks, isWorkspace),
		"queriesChildren":      QueryAppend(ctx, x.Queries, isWorkspace),
		"gofModule":            ctx.GofModuleName,
		"remoteResChildrenMap": RemoteChildrenMapResponse(ctx, x.Remotes, isWorkspace),
		"remoteReqChildrenMap": RemoteChildrenMapRequest(ctx, x.Remotes, isWorkspace),
		"actionResChildrenMap": ActionChildrenMapResponse(ctx, x.Actions, isWorkspace),
		"actionReqChildrenMap": ActionChildrenMapRequest(ctx, x.Actions, isWorkspace),
		"wsprefix":             wsPrefix,
		"ctx":                  ctx,
	})

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func (x Module3Action) RenderTemplate(ctx *CodeGenContext, fs embed.FS, fname string) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}

	wsPrefix := "workspaces."

	var tpl bytes.Buffer
	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"r":        x,
		"m":        x.RootModule,
		"ctx":      ctx,
		"fv":       FIREBACK_VERSION,
		"imports":  x.ImportDependecies(),
		"woo":      11,
		"wsprefix": wsPrefix,
	})
	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func RenderRpcGroupClassBody(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	content []byte,
	importMap ImportMap,
) ([]byte, error) {
	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}

	wsPrefix := "workspaces."

	var tpl bytes.Buffer
	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"content":  content,
		"ctx":      ctx,
		"woo":      15,
		"imports":  importMap,
		"fv":       FIREBACK_VERSION,
		"wsprefix": wsPrefix,
	})
	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

// Detect undefined and null in golang
// I think there is a huge problem to use this
// How to pass it to gorm?
// type SampleEntity or Dto struct {
//     Field1 Optional[string] `json:"field1"`
//     Field2 Optional[bool]   `json:"field2"`
//     Field3 Optional[int32]  `json:"field3"`
// }

type Optional[T any] struct {
	Defined bool
	Value   *T
}

// UnmarshalJSON is implemented by deferring to the wrapped type (T).
// It will be called only if the value is defined in the JSON payload.
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return json.Unmarshal(data, &o.Value)
}

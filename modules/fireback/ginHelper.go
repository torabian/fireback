package fireback

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/stoewer/go-strcase"
)

type LangQ struct {
	Lang   string
	Region string
	Q      float64
}

func ParseAcceptLanguage(acptLang string) []LangQ {

	var lqs []LangQ

	if acptLang == "" {
		return lqs
	}

	langQStrs := strings.Split(acptLang, ",")
	if len(langQStrs) == 0 {
		return lqs
	}

	for _, langQStr := range langQStrs {
		trimedLangQStr := strings.Trim(langQStr, " ")

		langQ := strings.Split(trimedLangQStr, ";")
		if len(langQ) == 1 {
			langRegion := strings.Split(langQ[0], "-")
			lq := LangQ{langRegion[0], langRegion[1], 1}
			lqs = append(lqs, lq)
		} else {
			qp := strings.Split(langQ[1], "=")
			q, err := strconv.ParseFloat(qp[1], 64)
			if err != nil {
				panic(err)
			}
			langRegion := strings.Split(langQ[0], "-")
			lq := LangQ{langRegion[0], langRegion[1], q}
			lqs = append(lqs, lq)
		}
	}
	return lqs
}

func GetMainLanguageFromAcceptLanguage(acceptLanguage string) string {
	query := ParseAcceptLanguage(acceptLanguage)

	if len(query) > 0 {
		return query[0].Lang
	}

	// default language
	return "en"
}

func GinMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		headers := []string{
			"accept",
			"authorization",
			"content-type",
			"content-length",
			"x-csrf-token",
			"token",
			"session",
			"origin",
			"host",
			"connection",
			"accept-encoding",
			"accept-language",
			"x-requested-with",
			"workspace",
			"workspace-id",
			"role-id",
			"deep",
			"query",
			"x-request-id",
			"x-http-method-override",
			"upload-length",
			"upload-offset",
			"tus-resumable",
			"upload-metadata",
			"upload-defer-length",
			"upload-concat",
			"user-agent",
			"referrer",
		}
		c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))

		c.Writer.Header().Set("Access-Control-Expose-Headers", "Upload-Offset, Location, Tus-Resumable")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

var ResolveStrategyPublic = "public"
var ResolveStrategyUser = "user"
var ResolveStrategyWorkspace = "workspace"

// Used for actions generally
type SecurityModel struct {
	// Only users which belong to root and actively selected the root workspace can
	// write to this entity from Fireback default functionality
	AllowOnRoot bool `json:"allowOnRoot,omitempty" yaml:"allowOnRoot,omitempty"`

	// Set of permissions which are required for this service.
	ActionRequires []PermissionInfo `json:"requires,omitempty" yaml:"requires,omitempty"`

	// Resolve strategy is by default on the workspace, you can change it by user
	// also. Be sure of the consequences
	ResolveStrategy string `json:"resolveStrategy,omitempty" yaml:"resolveStrategy,omitempty"`
}

// Used for defining the entity overall action permissions
type EntitySecurityModel struct {
	// Only users which belong to root and actively selected the root workspace can write to this entity from Fireback default functionality. Read mechanism won't be affected.
	WriteOnRoot *bool `json:"writeOnRoot,omitempty" yaml:"writeOnRoot,omitempty" jsonschema:"description=Only users which belong to root and actively selected the root workspace can write to this entity from Fireback default functionality. Read mechanism won't be affected."`

	// Only users which belong to root and actively selected the root workspace can read from entity from Fireback default functionality. Write mechanism is not affected.
	ReadOnRoot *bool `json:"readOnRoot,omitempty" yaml:"readOnRoot,omitempty" jsonschema:"description=Only users which belong to root and actively selected the root workspace can read from entity from Fireback default functionality. Write mechanism is not affected."`

	// Resolve strategy means that the content belongs either to workspace or user. It affects the query.
	ResolveStrategy *string `json:"resolveStrategy,omitempty" yaml:"resolveStrategy,omitempty" jsonschema:"enum=workspace,enum=user, description=Resolve strategy means that the content belongs either to workspace or user. It affects the query."`
}

func WithAuthorization(securityModel *SecurityModel) gin.HandlerFunc {
	return WithAuthorizationFn(securityModel)
}

// Converts the security policy and action into the gin
func CastRouteToHandler(r Module3Action) []gin.HandlerFunc {

	items := []gin.HandlerFunc{}

	// Handle security model - to this moment only WithAuth... is used,
	// Seems other models are not required
	if r.SecurityModel != nil && r.SecurityModel.ResolveStrategy != ResolveStrategyPublic {
		if r.Method == "REACTIVE" {
			items = append([]gin.HandlerFunc{WithSocketAuthorization(r.SecurityModel)}, items...)

		} else {

			items = append([]gin.HandlerFunc{WithAuthorization(r.SecurityModel)}, items...)
		}
	}

	items = append(items, r.Handlers...)

	return items
}

func CastRoutes(routes []Module3Action, r *gin.Engine) {
	for _, route := range routes {

		if route.Url == "" {
			continue
		}
		if route.Method == "GET" {
			r.GET(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "REACTIVE" {
			r.GET(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "DELETE" {
			r.DELETE(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "POST" {
			r.POST(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "PATCH" {
			r.PATCH(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "OPTIONS" {
			r.OPTIONS(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "HEAD" {
			r.HEAD(route.Url, CastRouteToHandler(route)...)
		}
	}
}

func CastRoutes2(routes []Module3Action, r *gin.RouterGroup) {
	for _, route := range routes {

		if route.Url == "" {
			continue
		}
		if route.Method == "GET" {
			r.GET(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "REACTIVE" {
			r.GET(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "DELETE" {
			r.DELETE(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "POST" {
			r.POST(route.Url, CastRouteToHandler(route)...)
		}
		// WebRtc is also a post request in it's nature
		if route.Method == "WEBRTC" {
			r.POST(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "PATCH" {
			r.PATCH(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "OPTIONS" {
			r.OPTIONS(route.Url, CastRouteToHandler(route)...)
		}
		if route.Method == "HEAD" {
			r.HEAD(route.Url, CastRouteToHandler(route)...)
		}
	}
}

type HttpRouteInformation struct {
	Method         string
	Url            string
	RequestEntity  string
	TargetEntity   string
	ResponseEntity string
	Action         string
	Params         []string
}
type HttpRouteInformationFile struct {
	ModuleName    string
	SubModuleName string
	Schema        []EntityJsonField
	Routes        []*HttpRouteInformation
}

func GetTypeString(myvar interface{}) string {
	pathRemover := regexp.MustCompile("(\\[).*/")
	t := reflect.TypeOf(myvar)

	if t != nil {
		full := (fmt.Sprintf("%s", t))
		full = pathRemover.ReplaceAllString(full, "$1")
		return full
	}

	return ""
}

func SplitFnToModuleAndFunc(input string) (string, string, string) {

	items := strings.Split(input, "/")

	fullName := items[len(items)-1]
	moduleName := strings.Split(fullName, ".")[0]
	modulePath := strings.Join(items[0:len(items)-1], "/")

	return fullName, modulePath + "/" + moduleName, moduleName
}

func UniqueString(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func GetInterfaceName(temp interface{}) string {
	return reflect.ValueOf(temp).Elem().String()
}

func GetFunctionNameFull(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func WriteMockDataToFile(lang string, region string, entityName string, data2 any) {

	fmt.Println("Writing:", lang, region, entityName, data2)
	body, err := json.MarshalIndent(data2, "", "  ")
	if err != nil {
		// log.Fatal(err)
	}

	os.Mkdir("./artifacts/md/"+lang, 0777)
	os.WriteFile("./artifacts/md/"+lang+"/"+entityName+".json", body, 0644)
}

func WriteEntitySchema(name string, data interface{}, mod string) {

	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./artifacts/entity-schema", 0777)
	os.WriteFile("./artifacts/entity-schema/"+name+".json", body, 0644)
}

type EntityResolvedInformation struct {
	Module        string
	ClassName     string
	GenericGroups []string
}

func EntityFromString(str string) EntityResolvedInformation {
	dot := strings.Index(str, ".")

	vModule := ""
	vClassName := ""
	genericGroups := []string{}

	if dot == -1 {
		return EntityResolvedInformation{
			Module:        vModule,
			ClassName:     vClassName,
			GenericGroups: genericGroups,
		}
	}

	// Match this pattern *[]module.entity

	if strings.Contains(str, "*[]") {
		str = strings.ReplaceAll(str, "*[]", "")
		dot = strings.Index(str, ".")
		vModule = strings.ReplaceAll(str[0:dot], "*", "")
		vClassName = str[dot+1:]
	} else if strings.Contains(str, "[") && strings.Contains(str, "]") {
		// Match the generic patten
		startBracket := strings.Index(str, "[")
		endBracket := strings.Index(str, "]")

		between := str[startBracket+1 : endBracket]
		dot = strings.Index(between, ".")
		vModule = strings.ReplaceAll(between[0:dot], "*", "")
		vClassName = between[dot+1:]
	} else {
		// Simple workspace.entity type

		vModule = strings.ReplaceAll(str[0:dot], "*", "")
		vClassName = str[dot+1:]
	}

	return EntityResolvedInformation{
		Module:        vModule,
		ClassName:     vClassName,
		GenericGroups: genericGroups,
	}
}

func (route Module3Action) ResponseEntityMeta() EntityResolvedInformation {

	return EntityFromString(GetTypeString(route.ResponseEntity))
}

func (x Module3Action) ResponseEntityComputed() string {
	if x.Out == nil {
		return "any"
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

	return "any"
	// j := EntityFromString(GetTypeString(route.ResponseEntity))
	// return j.ClassName
}

func (route Module3Action) ResponseEntityComputedSplit() string {
	return strings.ReplaceAll(route.ResponseEntityComputed(), "ResDto", ".Res")
}

func (route Module3Action) RequestEntityComputedSplit() string {
	return strings.ReplaceAll(route.RequestEntityComputed(), "ReqDto", ".Req")
}

func (route Module3Action) RequestEntityMeta() EntityResolvedInformation {
	return EntityFromString(GetTypeString(route.RequestEntity))
}

func (x Module3Action) RequestEntityComputed() string {
	if x.In == nil {
		return ""
	}

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

	// j := EntityFromString(GetTypeString(route.RequestEntity))
	// return j.ClassName
}

func (route Module3Action) EntityKey() string {
	t := ""
	if route.Method == "DELETE" {
		if route.TargetEntity != nil {
			t = GetTypeString(route.TargetEntity)
		}
	} else {
		if route.ResponseEntity != nil {
			t = GetTypeString(route.ResponseEntity)
		} else if route.ResponseEntityComputed() != "" {
			t = route.ResponseEntityComputed()
		}
	}

	t = strings.ReplaceAll(t, "[]", "")

	return t
}

func (route Module3Action) UrlParams() []string {
	params := []string{}
	parts := strings.Split(route.Url, "/")

	for _, part := range parts {
		if part != "" && part[0:1] == ":" {
			params = append(params, part)
		}
	}

	return params
}

func (route Module3Action) GetFuncName() string {
	fnName := route.ExternFuncName

	if fnName == "" {
		fnName = route.Method + strings.ReplaceAll(route.Url, "/", "_")
		fnName = strings.Replace(fnName, ":", "_by_", 1)
		fnName = strings.ReplaceAll(fnName, ":", "_and_")
		fnName = strcase.LowerCamelCase(fnName)
	}

	return fnName
}

func (route Module3Action) GetFuncNameUpper() string {
	return ToUpper(route.GetFuncName())
}

type IResponseDelete struct {
	Data *struct {
		RowsAffected int64 `json:"rowsAffected"`
	} `json:"data"`
	Error *struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Errors  []struct {
			Location string `json:"location"`
			Message  string `json:"message"`
		} `json:"errors"`
	} `json:"error"`
}

type IResponse[T any] struct {
	Data  *T      `json:"data"`
	Error *IError `json:"error"`
	// Error *struct {
	// 	Message string `json:"message"`
	// 	Code    string `json:"code"`
	// 	Errors  []struct {
	// 		Location string `json:"location"`
	// 		Message  string `json:"message"`
	// 	} `json:"errors"`
	// } `json:"error"`
}

type IResponseList[T any] struct {
	Data *struct {
		Items []T `json:"items"`
	} `json:"data"`
	Error *struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Errors  []struct {
			Location string `json:"location"`
			Message  string `json:"message"`
		} `json:"errors"`
	} `json:"error"`
}

type TestResult struct {
	Service      string `json:"service"`
	RequestBody  any    `json:"requestBody"`
	ResponseBody any    `json:"responseBody"`
	Name         string `json:"name"`
	Method       string `json:"method"`
}

func DocumentTestResult(testResult TestResult) {
	str, _ := json.MarshalIndent(testResult, "", "  ")
	os.Mkdir("TestResults", 0777)
	os.WriteFile("TestResults/"+testResult.Name+".json", str, 0777)
}

func RenderTemplateToGin(ctx *gin.Context, path string, ui fs.FS, data any) {

	filename := filepath.Base(path)
	tmpl, err := template.New("").Funcs(CommonMap).ParseFS(ui, path, "SharedParticles.tpl")
	if err != nil {
		ctx.JSON(500, gin.H{
			"error1": err.Error(),
		})
		return
	}

	err = tmpl.ExecuteTemplate(ctx.Writer, filename, data)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error2": err.Error(),
		})
		return
	}
}

package workspaces

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
	"gorm.io/gorm"
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

// @meta(include)
type QueryDSL struct {
	Query        string   `json:"query"`
	StartIndex   int      `json:"startIndex"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Deep         bool     `json:"deep"`
	Sort         string   `json:"sort"`
	UniqueId     string   `json:"uniqueId"`
	UniqueIdList string   `json:"uniqueListId"`
	WithPreloads []string `json:"withPreloads"`
	JsonQuery    string   `json:"jsonQuery"`

	Tx *gorm.DB
	// This event will be trigged in the system, if that action is done
	TriggerEventName string `json:"-"`

	Authorization string `json:"authorization"`

	// Parsed languages
	AcceptLanguage []LangQ `json:"-"`

	// This is the person who is requesting, regardless of the workspace
	UserId string `json:"-"`

	LinkerId string `json:"-"`

	// This is the person who is requesting, regardless of the workspace
	SearchPhrase string `json:"searchPhrase"`

	// This is the workspace which user is working inside, usually data belongs there
	WorkspaceId string `json:"-"`

	// Those capabilities which user has
	ActionRequires []string `json:"-"`

	// List of permissions that this request is affecting
	RequestAffectingScopes []string `json:"-"`

	// This is the capabilities that user has
	UserHas []string `json:"-"`

	InternalQuery string   `json:"-"`
	Language      string   `json:"-"`
	Region        string   `json:"-"`
	Preloads      []string `json:"-"`
}

func (x QueryDSL) Json() string {
	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))

}

func GinMiddleware() gin.HandlerFunc {
	config := GetAppConfig()

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Headers.AccessControlAllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", config.Headers.AccessControlAllowHeaders)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

type SecurityModel struct {
	ActionRequires           []string
	OnlyInSpecificWorkspaces []string
	Model                    string
}

// Converts the security policy and action into the gin
func CastRouteToHandler(r Module2Action) []gin.HandlerFunc {

	items := []gin.HandlerFunc{}

	// Handle security model - to this moment only WithAuth... is used,
	// Seems other models are not required
	if len(r.SecurityModel.ActionRequires) > 0 {
		items = append(items, WithAuthorization(r.SecurityModel.ActionRequires))
	}

	// If there are no handlers, we need to automatically add them

	// I failed here
	// if len(r.Handlers) == 0 {

	// 	if r.Format == "POST_ONE" {

	// 		items = append(items, func(c *gin.Context) {
	// 			HttpPostEntity[any](c, r.Action)
	// 		})
	// 	}
	// }

	items = append(items, r.Handlers...)

	return items
}

func CastRoutes(routes []Module2Action, r *gin.Engine) {
	for _, route := range routes {

		if route.Virtual {
			continue
		}
		if route.Method == "GET" || route.Method == "REACTIVE" {
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

type HttpRouteInformation struct {
	Method         string
	Url            string
	ExternFuncName string // the function which will be called by third party
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

func (route Module2Action) ResponseEntityMeta() EntityResolvedInformation {

	return EntityFromString(GetTypeString(route.ResponseEntity))
}

func (route Module2Action) ResponseEntityComputed() string {

	j := EntityFromString(GetTypeString(route.ResponseEntity))
	return j.ClassName
}

func (route Module2Action) ResponseEntityComputedSplit() string {
	return strings.ReplaceAll(route.ResponseEntityComputed(), "ResDto", ".Res")
}

func (route Module2Action) RequestEntityComputedSplit() string {
	return strings.ReplaceAll(route.RequestEntityComputed(), "ReqDto", ".Req")
}

func (route Module2Action) RequestEntityMeta() EntityResolvedInformation {
	return EntityFromString(GetTypeString(route.RequestEntity))
}

func (route Module2Action) RequestEntityComputed() string {
	j := EntityFromString(GetTypeString(route.RequestEntity))
	return j.ClassName
}

func (route Module2Action) EntityKey() string {
	t := ""
	if route.Method == "DELETE" {
		t = GetTypeString(route.TargetEntity)
	} else {
		t = GetTypeString(route.ResponseEntity)
	}

	t = strings.ReplaceAll(t, "[]", "")

	return t
}

func (route Module2Action) UrlParams() []string {
	params := []string{}
	parts := strings.Split(route.Url, "/")

	for _, part := range parts {
		if part != "" && part[0:1] == ":" {
			params = append(params, part)
		}
	}

	return params
}

func (route Module2Action) GetFuncName() string {
	fnName := route.ExternFuncName

	if fnName == "" {
		fnName = route.Method + strings.ReplaceAll(route.Url, "/", "_")
		fnName = strings.Replace(fnName, ":", "_by_", 1)
		fnName = strings.ReplaceAll(fnName, ":", "_and_")
		fnName = strcase.LowerCamelCase(fnName)
	}

	return fnName
}

func (route Module2Action) GetFuncNameUpper() string {
	return ToUpper(route.GetFuncName())
}

func WriteHttpInformationToFile(routes *[]Module2Action, schema []EntityJsonField,
	subModuleName string, mod string) {

	data := []*HttpRouteInformation{}

	for _, route := range *routes {

		action := ""

		if route.Action != nil {
			action = GetFunctionName(route.Action)
		}

		entity := &HttpRouteInformation{
			Method:         route.Method,
			Url:            strings.TrimPrefix(route.Url, "/"),
			ExternFuncName: route.GetFuncName(),
			RequestEntity:  GetTypeString(route.RequestEntity),
			ResponseEntity: GetTypeString(route.ResponseEntity),
			TargetEntity:   GetTypeString(route.TargetEntity),
			Action:         action,
			Params:         []string{},
		}

		parts := strings.Split(route.Url, "/")

		for _, part := range parts {
			if part != "" && part[0:1] == ":" {
				entity.Params = append(entity.Params, part)
			}
		}

		data = append(data, entity)
	}

	data2 := HttpRouteInformationFile{
		SubModuleName: subModuleName,
		ModuleName:    mod,
		Routes:        data,
		Schema:        schema,
	}

	body, err := json.MarshalIndent(data2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./artifacts/intermediate-http", 0777)
	os.WriteFile("./artifacts/intermediate-http/"+subModuleName+".json", body, 0644)

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

package fireback

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

type LangQ struct {
	Lang   string
	Region string
	Q      float64
}

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.Writer.Header()

		h.Set("Access-Control-Allow-Origin", "*")
		h.Set("Access-Control-Allow-Methods", "*")
		h.Set("Access-Control-Allow-Headers", "*")
		h.Set("Access-Control-Expose-Headers", "*")
		h.Set("Access-Control-Max-Age", "86400")

		// Only enable this if you're NOT using "*"
		// h.Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

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

package fireback

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/urfave/cli"
)

var Upgrader = websocket.Upgrader{
	//Solve "request origin not allowed by Upgrader.CheckOrigin"
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type QueryableAction[T any] func(query QueryDSL) ([]*T, *QueryResultMeta, error)

func BindCli(c *cli.Context, entity any) (any, error) {
	reqValue := reflect.Indirect(reflect.ValueOf(entity))
	if reqValue.MethodByName("FromCli").IsValid() {
		args := []reflect.Value{reflect.ValueOf(c)}

		res := reqValue.MethodByName("FromCli").Call(args)

		if len(res) > 0 {
			return res[0].Interface(), nil
		}

		return nil, nil
	}

	return nil, errors.New("cannot bind the cli")
}

func zeroValueT[T any]() T {
	var zeroVal T
	return zeroVal
}
func CastAnyToT[T any](val interface{}) T {
	t, ok := val.(T)
	if !ok {
		// Handle the case where the type assertion fails
		return zeroValueT[T]()
	}
	return t
}

func CliPostEntity[T any, V any](c *cli.Context, fn func(T, QueryDSL) (*V, *IError), security *SecurityModel) (*V, *IError) {
	f := CommonCliQueryDSLBuilderAuthorize(c, security)
	var body T

	if result, err := BindCli(c, &body); err != nil {
		fmt.Println("CORRECT_BODY_SIGNATURE_IS_NEEDED", err)
		return nil, GormErrorToIError(err)
	} else {
		return fn(CastAnyToT[T](result), f)
	}

}

func ginBodyToBytes(c *gin.Context) ([]byte, *IError) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, Create401Error(&FirebackMessages.BodyIsEmptyEof, []string{})
		} else if errors.Is(err, io.ErrUnexpectedEOF) {
			return nil, Create401Error(&FirebackMessages.BodyUnexpectedEof, []string{})
		} else if errors.Is(err, http.ErrBodyReadAfterClose) {
			return nil, Create401Error(&FirebackMessages.BodyReadAfterClose, []string{})
		} else {
			return nil, Create401Error(&FirebackMessages.UnknownErrorReadingBody, []string{})
		}
	}

	// Reset the body so it can be read again later
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

func isMap(m interface{}) bool {
	if m == nil {
		return false
	}

	rt := reflect.TypeOf(m)
	return rt.Kind() == reflect.Map
}

func isSlice(m interface{}) bool {
	if m == nil {
		return false
	}

	rt := reflect.TypeOf(m)
	return rt.Kind() == reflect.Slice
}

/**
*	This is an specific type of translation. It gets a json recursivly,
*   and if finds 'translations' field, it would replace the parent with that value
**/
func RecursiveJsonTranslate(content map[string]interface{}, lang string) map[string]interface{} {

	val, hasTranslations := content["translations"]
	if hasTranslations {
		translations := val.([]interface{})
		for _, translate := range translations {
			dictionary, okay := translate.(map[string]interface{})

			if !okay {
				continue
			}

			if dictionary["languageId"] != lang {
				continue
			}

			for key, value := range dictionary {
				if key == "languageId" {
					continue
				}
				content[key] = value.(string)
			}
		}
	}

	for k, v := range content {

		if isMap(v) {
			sub, ok := v.(map[string]interface{})
			if ok {
				content[k] = RecursiveJsonTranslate(sub, lang)
			}
		}
		if isSlice(v) {
			if k != "translations" {
				var data []map[string]interface{} = []map[string]interface{}{}

				for _, n := range v.([]interface{}) {
					mapped, okay := n.(map[string]interface{})
					if okay {
						data = append(data, n.(map[string]interface{}))
						continue
					} else {
						data = append(data, RecursiveJsonTranslate(mapped, lang))
					}
				}
				content[k] = data

			}

		}

	}

	return content

}

func PolyglotQueryHandler(entity any, query *QueryDSL) map[string]interface{} {

	str, _ := json.MarshalIndent(entity, "", "  ")
	var content map[string]interface{}
	json.Unmarshal(str, &content)

	// @todo: Huge bug here. It touches also the json content, which it should not ever touch.
	// perhaps querying the content from database level should be fixed
	// RecursiveJsonTranslate(content, query.Language)

	val, ok := content["translations"]

	if !ok {
		return content
	}

	translations := val.([]interface{})

	for _, translate := range translations {
		dictionary, okay := translate.(map[string]interface{})
		if !okay {
			continue
		}

		if dictionary["languageId"] != query.Language {
			continue
		}

		for key, value := range dictionary {

			if key == "languageId" {
				continue
			}
			content[key] = value.(string)
		}
	}

	return content

}

func QueryEntitySuccessResult[T any](f QueryDSL, items []T, meta *QueryResultMeta) gin.H {
	mappedItems := []map[string]interface{}{}
	for _, item := range items {
		content := PolyglotQueryHandler(item, &f)
		mappedItems = append(mappedItems, content)
	}

	data := gin.H{
		"startIndex":   f.StartIndex,
		"itemsPerPage": f.ItemsPerPage,
		"items":        mappedItems,
	}

	if meta != nil {
		data["next"] = gin.H{
			"cursor": meta.Cursor,
		}
		data["totalItems"] = meta.TotalItems
		data["totalAvailableItems"] = meta.TotalAvailableItems
	}

	return gin.H{
		"data": data,
	}
}

// Use it for requests which are kinda having body, such as post, put, patch, etc.
// It would read the body (either if it's json, form-data, yaml, etc, based on headers)
// and cast it to the 'body'. Make sure calling this with &body, not body
// Extend this function if you want to support different formats.
func ReadGinRequestBodyAndCastToGoStruct(c *gin.Context, body any, f QueryDSL) (aborted bool) {

	bodyBytes, err := ginBodyToBytes(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.ToPublicEndUser(&f)})
		return true
	}

	switch DetectGinContentType(c) {

	case ContentTypeYAML:
		if err := BindYamlStringWithDetails(bodyBytes, body); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.ToPublicEndUser(&f)})
			return true
		}
	case ContentTypeFormData:
		if err := BindMultiPartFormDataWithDetails(c, body); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.ToPublicEndUser(&f)})
			return true
		}
	case ContentTypeURLEncoded:
		if err := BindFormUrlEncodedWithDetails(c, body); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.ToPublicEndUser(&f)})
			return true
		}
	case ContentTypeXML:
		if err := BindXmlStringWithDetails(bodyBytes, body); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.ToPublicEndUser(&f)})
			return true
		}

	default:
		if err := BindJsonStringWithDetails(bodyBytes, body); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.ToPublicEndUser(&f)})
			return true
		}
	}

	return false
}

type BulkRecordRequest[T any] struct {
	Records []*T `json:"records"`
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func UniversalJsonString(okayResult interface{}, errorResult *IError) string {

	if errorResult != nil {
		data, _ := json.MarshalIndent(errorResult, "", "  ")
		return string(data)
	}

	data, _ := json.MarshalIndent(gin.H{
		"data": okayResult,
	}, "", "  ")
	return string(data)
}

func UniversalJsonStringFormatQuery(okayResult interface{}, count int64, errorResult error) string {

	if okayResult == nil {
		eedata, _ := json.MarshalIndent(errorResult, "", "  ")
		return string(eedata)
	}

	data, _ := json.MarshalIndent(gin.H{
		"data": gin.H{
			"items":      okayResult,
			"totalItems": count,
		},
	}, "", "  ")
	return string(data)
}

func DtoFromString[T any](input string) T {
	var body T
	json.Unmarshal([]byte(input), &body)
	return body
}

type BindingActionItem struct {
	FunctionName     string
	FunctionFullName string
	GenericName      string
	Format           string
	ModuleName       string
	IPCSecurity      string
	Url              string
}

type BindingTemplateData struct {
	Imports []string
	Items   []BindingActionItem
}

func GetType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

var IpcTemplate string = `
package main

import (
	{{ with .Imports }}
		{{ range . }}
			"{{ .}}"
		{{ end }}
	{{ end }}
)

{{ with .Items }}
	{{ range .}}
		// URL: {{ .Url }}
		{{ if eq .Format "QUERY" }}
		// dto is not used, for compatibilty we have it
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonStringFormatQuery(
				{{ .FunctionFullName}}(
					fireback.ActionArgumentFormatQuery(query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
		{{ if eq .Format "GET_ONE" }}
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonString(
				{{ .FunctionFullName}}(
					fireback.ActionArgumentFormatQuery(query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
		{{ if eq .Format "POST_ONE" }}
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonString(
				{{ .FunctionFullName}}(
					fireback.ActionArgumentsFormatPostOne[{{ .GenericName }}](dto, query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
		{{ if eq .Format "PATCH_ONE" }}
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonString(
				{{ .FunctionFullName}}(
					fireback.ActionArgumentsFormatUpdateOne[{{ .GenericName }}](dto, query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
		{{ if eq .Format "DELETE_DSL" }}
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonString(
				{{ .FunctionFullName}}(
					fireback.ActionArgumentsFormatDeleteDSL(query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
	{{ end }}
{{ end }}

`

func GenerateBindings(items []Module3Action) cli.Command {
	return cli.Command{

		Name:    "bindings",
		Aliases: []string{"bi"},
		Usage:   "Generates the bindings",
		Action: func(c *cli.Context) error {

			data := BindingTemplateData{
				Imports: []string{
					"github.com/torabian/fireback/modules/fireback",
				},
			}
			for _, item := range items {
				if item.Action != nil {
					fn := GetFunctionName(item.Action)
					ffn := GetFunctionNameFull(item.Action)
					imp, importPath, moduleName := SplitFnToModuleAndFunc(ffn)

					genericName := ""
					if item.Format == "POST_ONE" || item.Format == "PATCH_ONE" {
						genericName = strings.Replace(moduleName+"."+GetType(item.RequestEntity), "*", "", 1)
					}

					data.Imports = UniqueString(append(data.Imports, importPath))
					data.Items = append(data.Items, BindingActionItem{
						FunctionName:     fn,
						FunctionFullName: imp,
						Format:           item.Format,
						GenericName:      genericName,
						ModuleName:       moduleName,
						Url:              item.Url,
					})

				}
			}
			var tpl bytes.Buffer

			t := template.Must(template.New("html-tmpl").Parse(IpcTemplate))
			err := t.Execute(&tpl, data)
			if err != nil {
				panic(err)
			}
			result := tpl.String()

			target := "./cmd/academy-desktop/ipc-binding.go"
			err2 := os.WriteFile(target, []byte(result), 0644)
			if err2 == nil {
				fmt.Println("Bindings wrote to:", target)
			}

			return nil
		},
	}
}

type PivotResult struct {
	UniqueId0 string `json:"uniqueId_0" gorm:"unique_id0"`
	Name0     string `json:"name_0" gorm:"name0"`

	UniqueId1 string `json:"uniqueId_1" gorm:"unique_id1"`
	Name1     string `json:"name_1" gorm:"name1"`

	UniqueId2 string `json:"uniqueId_2" gorm:"unique_id2"`
	Name2     string `json:"name_2" gorm:"name2"`

	UniqueId3 string `json:"uniqueId_3" gorm:"unique_id3"`
	Name3     string `json:"name_3" gorm:"name3"`

	UniqueId4 string `json:"uniqueId_4" gorm:"unique_id4"`
	Name4     string `json:"name_4" gorm:"name4"`

	UniqueId5 string `json:"uniqueId_5" gorm:"unique_id5"`
	Name5     string `json:"name_5" gorm:"name5"`

	UniqueId6 string `json:"uniqueId_6" gorm:"unique_id6"`
	Name6     string `json:"name_6" gorm:"name6"`
}

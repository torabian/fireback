package workspaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	Any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/gorilla/websocket"
	"github.com/urfave/cli"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var upgrader = websocket.Upgrader{
	//Solve "request origin not allowed by Upgrader.CheckOrigin"
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HttpReactiveQuery[T any](ctx *gin.Context, fn func(QueryDSL, chan bool, chan map[string]interface{}) chan *T) {
	f := ExtractQueryDslFromGinContext(ctx)

	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		c.WriteJSON(GormErrorToIError(err))

		c.Close()
		return
	}

	done := make(chan bool)
	read := make(chan map[string]interface{})

	// Add done channel here to be passed later on
	act := fn(f, done, read)

	defer c.Close()

	go func() {

		for {
			// defer close(done)

			var k interface{} = nil
			err := c.ReadJSON(&k)

			if err != nil {
				fmt.Println("Closed")
				close(done)
				// done <- true
				// close(done)
				close(act)
				return
			}
			read <- k.(map[string]interface{})
		}
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done222")
			// defer close(act)

			return
		case row, more := <-act:
			err := c.WriteJSON(row)
			if err != nil || !more {
				return
			}

		}
	}

}

func HttpSocketRequest(ctx *gin.Context, fn func(QueryDSL, func(string)), onRead func(QueryDSL, interface{})) {
	f := ExtractQueryDslFromGinContext(ctx)

	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		c.WriteJSON(GormErrorToIError(err))

		c.Close()
		return
	}

	go func() {
		for {
			var k interface{} = nil
			err := c.ReadJSON(&k)

			if err != nil {

				return
			}
			onRead(f, k)
		}
	}()

	fn(f, func(data string) {

		c.WriteJSON(data)
	})
}

func BindCli(c *cli.Context, entity any) (any, error) {
	reqValue := reflect.ValueOf(entity)
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

func CliPostEntity[T any, V any](c *cli.Context, fn func(T, QueryDSL) (V, *IError)) {
	f := CommonCliQueryDSLBuilder(c)
	var body T

	if result, err := BindCli(c, &body); err != nil {
		fmt.Println("CORRECT_BODY_SIGNATURE_IS_NEEDED")
	} else {
		fmt.Println(result)
	}
	if entity, err := fn(body, f); err != nil {
		fmt.Println(err.Error())
	} else {
		f, _ := json.MarshalIndent(entity, "", "  ")
		fmt.Println(string(f))
	}

}

func HttpPostEntity[T any, V any](c *gin.Context, fn func(T, QueryDSL) (V, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	id := c.Param("uniqueId")
	if id != "" {
		f.UniqueId = id
	}

	var body T

	if err := c.BindJSON(&body); err != nil {

		c.AbortWithStatusJSON(500, gin.H{
			"error": gin.H{
				"code": "CORRECT_BODY_SIGNATURE_IS_NEEDED",
			},
		})
		return
	}

	if entity, err := fn(body, f); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err, "data": entity})
	} else {
		c.JSON(200, gin.H{
			"data": entity,
		})
	}
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

	RecursiveJsonTranslate(content, query.Language)

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

	return gin.H{
		"data": gin.H{
			"startIndex":          f.StartIndex,
			"itemsPerPage":        f.ItemsPerPage,
			"items":               mappedItems,
			"totalItems":          meta.TotalItems,
			"totalAvailableItems": meta.TotalAvailableItems,
		},
	}
}

func HttpStreamFileChannel(
	c *gin.Context,
	fn func(query QueryDSL) (chan []byte, *IError),
) {
	f := ExtractQueryDslFromGinContext(c)
	chanStream, err := fn(f)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	GinStreamFromChannel(c, chanStream)

}
func HttpQueryEntity[T any](
	c *gin.Context,
	fn func(query QueryDSL) ([]T, *QueryResultMeta, error),
) {

	f := ExtractQueryDslFromGinContext(c)

	if items, count, err := fn(f); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	} else {
		result := QueryEntitySuccessResult(f, items, count)
		result["jsonQuery"] = c.Query("jsonQuery")
		c.JSON(200, result)
	}
}
func HttpQueryEntity2[T any](
	c *gin.Context,
	fn func(query QueryDSL) ([]T, *QueryResultMeta, *IError),
) {

	f := ExtractQueryDslFromGinContext(c)

	if items, count, err := fn(f); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	} else {
		result := QueryEntitySuccessResult(f, items, count)
		result["jsonQuery"] = c.Query("jsonQuery")
		c.JSON(200, result)
	}
}

func HttpRawQuery[T any](
	c *gin.Context,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, *IError),
) {

	f := ExtractQueryDslFromGinContext(c)

	if items, count, err := fn(f); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	} else {

		mappedItems := []map[string]interface{}{}
		for _, item := range items {
			content := PolyglotQueryHandler(item, &f)
			mappedItems = append(mappedItems, content)
		}

		c.JSON(200, gin.H{
			"data": gin.H{
				"startIndex":          f.StartIndex,
				"itemsPerPage":        f.ItemsPerPage,
				"items":               mappedItems,
				"totalItems":          count.TotalItems,
				"totalAvailableItems": count.TotalAvailableItems,
			},
		})
	}
}

func HttpGetEntity[T any](
	c *gin.Context,
	fn func(QueryDSL) (T, *IError),
) {
	id := c.Param("uniqueId")
	f := ExtractQueryDslFromGinContext(c)
	f.UniqueId = id

	if item, err := fn(f); err != nil {

		code := http.StatusBadRequest

		if err.HttpCode > 0 {
			code = int(err.HttpCode)
		}
		c.AbortWithStatusJSON(code, gin.H{
			"error": err,
		})

	} else {
		c.JSON(200, gin.H{
			"data": PolyglotQueryHandler(item, &f),
		})
	}
}

func HttpRemoveEntity[T any](c *gin.Context, fn func(QueryDSL) (T, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	var body DeleteRequest
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	f.Query = body.Query

	if item, err := fn(f); err != nil {
		c.JSON(400, gin.H{
			"error": gin.H{
				"message": err.Error(),
			},
		})
	} else {
		c.JSON(200, gin.H{
			"data": gin.H{
				"rowsAffected": item,
			},
		})
	}
}

func HttpUpdateEntity[T any, V any](c *gin.Context, fn func(QueryDSL, T) (V, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	var body T
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	if entity, err := fn(f, body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(200, gin.H{
			"data": entity,
		})
	}
}

type BulkRecordRequest[T any] struct {
	Records []*T `json:"records"`
}

func HttpUpdateEntities[T any](c *gin.Context, fn func(QueryDSL, *BulkRecordRequest[T]) (*BulkRecordRequest[T], *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	var body BulkRecordRequest[T]
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	if entity, err := fn(f, &body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(200, gin.H{
			"data": entity,
		})
	}
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
					workspaces.ActionArgumentFormatQuery(query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
		{{ if eq .Format "GET_ONE" }}
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonString(
				{{ .FunctionFullName}}(
					workspaces.ActionArgumentFormatQuery(query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
		{{ if eq .Format "POST_ONE" }}
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonString(
				{{ .FunctionFullName}}(
					workspaces.ActionArgumentsFormatPostOne[{{ .GenericName }}](dto, query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
		{{ if eq .Format "PATCH_ONE" }}
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonString(
				{{ .FunctionFullName}}(
					workspaces.ActionArgumentsFormatUpdateOne[{{ .GenericName }}](dto, query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
		{{ if eq .Format "DELETE_DSL" }}
		func (a *AppIPCBridge) {{ .FunctionName}}(dto string, query string) string {
			return UniversalJsonString(
				{{ .FunctionFullName}}(
					workspaces.ActionArgumentsFormatDeleteDSL(query, "{{ .Url }}", "{{ .IPCSecurity }}"),
				),
			)
		}
		{{ end }}
	{{ end }}
{{ end }}

`

func GenerateBindings(items []Module2Action) cli.Command {
	return cli.Command{

		Name:    "bindings",
		Aliases: []string{"bi"},
		Usage:   "Generates the bindings",
		Action: func(c *cli.Context) error {

			data := BindingTemplateData{
				Imports: []string{
					"github.com/torabian/fireback/modules/workspaces",
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

func ConvertAnyToInterface(anyValue *Any.Any) (interface{}, error) {
	var value interface{}
	bytesValue := &wrappers.BytesValue{}
	err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
	if err != nil {
		return value, err
	}
	uErr := json.Unmarshal(bytesValue.Value, &value)
	if err != nil {
		return value, uErr
	}
	return value, nil
}

func ConvertInterfaceToAny(v interface{}) (*Any.Any, error) {
	anyValue := &Any.Any{}
	bytes, _ := json.Marshal(v)
	bytesValue := &wrappers.BytesValue{
		Value: bytes,
	}
	err := anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})
	return anyValue, err
}

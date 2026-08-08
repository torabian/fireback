package fireback

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v2"
)

var Upgrader = websocket.Upgrader{
	//Solve "request origin not allowed by Upgrader.CheckOrigin"
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type QueryableAction[T any] func(query QueryDSL) ([]*T, *QueryResultMeta, error)

func BindCli(c *cli.Command, entity any) (any, error) {
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

func QueryEntitySuccessResult[T any](f QueryDSL, items []T, meta *QueryResultMeta) gin.H {

	var formatted = []json.RawMessage{}
	for _, item := range items {
		if len([]string(f.SelectableColumn)) > 0 {
			data, _ := MarshalWithWhitelist(item, []string(f.SelectableColumn))
			formatted = append(formatted, json.RawMessage(data))
		} else {

			data, _ := json.Marshal(item)
			formatted = append(formatted, json.RawMessage(data))
		}
	}

	data := gin.H{
		"startIndex":   f.StartIndex,
		"itemsPerPage": f.ItemsPerPage,
		"items":        formatted,
	}

	if meta != nil {
		data["next"] = gin.H{
			"cursor": meta.Cursor,
		}
		data["totalItems"] = meta.TotalItems
		data["totalAvailableItems"] = meta.TotalAvailableItems
	}

	if f.G != nil && IsYaml(f.G) || f.C != nil && (f.C.String("x-accept") == "application/x-yaml" || f.C.String("x-accept") == "application/yaml" || f.C.String("x-accept") == "text/yaml" || f.C.String("x-accept") == "yaml" || f.C.String("x-accept") == "yml") {
		var yamlFormatted []any
		for _, item := range formatted {
			var m any
			_ = json.Unmarshal(item, &m)
			yamlFormatted = append(yamlFormatted, m)
		}
		data["items"] = yamlFormatted
	}

	return gin.H{
		"data": data,
	}
}

func abortWithError(c *gin.Context, err error, f QueryDSL) bool {
	accept := c.GetHeader("Accept")
	isYAML := accept == "application/x-yaml" || accept == "application/yaml" || accept == "text/yaml"

	if isYAML {
		c.Header("Content-Type", "application/x-yaml")
		yamlData, marshalErr := yaml.Marshal(err)
		if marshalErr != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to marshal yaml"})
			return true
		}
		c.Writer.WriteHeader(500)
		c.Writer.Write(yamlData)
	} else {
		c.AbortWithStatusJSON(500, err)
	}
	return true
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

func GetType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
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

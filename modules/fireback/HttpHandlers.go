package fireback

import (
	"encoding/json"
	"io"
	"net/http"
	reflect "reflect"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func HttpUpdateEntity[T any, V any](c *gin.Context, fn func(QueryDSL, T) (V, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	var body T

	if ReadGinRequestBodyAndCastToGoStruct(c, &body, f) {
		return
	}

	if entity, err := fn(f, body); err != nil {
		GinWriteContent(c, int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&f)})
	} else {
		GinWriteContent(c, 200, gin.H{
			"data": entity,
		})
	}
}

func HttpUpdateEntities[T any](c *gin.Context, fn func(QueryDSL, *BulkRecordRequest[T]) (*BulkRecordRequest[T], *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	var body BulkRecordRequest[T]

	if ReadGinRequestBodyAndCastToGoStruct(c, &body, f) {
		return
	}

	if entity, err := fn(f, &body); err != nil {
		GinWriteContent(c, int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&f)})
	} else {
		GinWriteContent(c, 200, gin.H{
			"data": entity,
		})
	}
}

func HttpRemoveEntity[T any](c *gin.Context, fn func(QueryDSL) (T, *IError)) {
	f := ExtractQueryDslFromGinContext(c)
	var body DeleteRequest
	if ReadGinRequestBodyAndCastToGoStruct(c, &body, f) {
		return
	}
	f.Query = body.Query
	if item, err := fn(f); err != nil {
		GinWriteContent(c, int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&f)})
	} else {
		GinWriteContent(c, 200, gin.H{
			"data": gin.H{
				"rowsAffected": item,
			},
		})
	}
}

func HttpReactiveQuery[T any](ctx *gin.Context, fn func(QueryDSL, chan bool, chan map[string]interface{}) chan *T) {
	f := ExtractQueryDslFromGinContext(ctx)

	c, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
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
			var k interface{} = nil
			err := c.ReadJSON(&k)

			if err != nil {
				close(done)
				close(act)
				return
			}
			read <- k.(map[string]interface{})
		}
	}()

	for {
		select {
		case <-done:
			return
		case row, more := <-act:
			err := c.WriteJSON(row)
			if err != nil || !more {
				return
			}

		}
	}

}

func HttpSocketRequest(ctx *gin.Context, fn func(QueryDSL, func([]byte)), onRead func(QueryDSL, SocketReadChan)) {
	f := ExtractQueryDslFromGinContext(ctx)

	c, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		c.WriteJSON(GormErrorToIError(err))

		c.Close()
		return
	}

	f.RawSocketConnection = c

	go func() {
		for {
			_, k, err := c.ReadMessage()

			onRead(f, SocketReadChan{
				Data:  k,
				Error: err,
			})

			if err != nil {
				return
			}
		}
	}()

	fn(f, func(data []byte) {
		c.WriteMessage(websocket.TextMessage, data)
	})
}

func HttpSocketRequest2(ctx *gin.Context, fn func(QueryDSL, func([]byte)), onRead func(QueryDSL, SocketReadChan)) {
	f := ExtractQueryDslFromGinContext(ctx)

	c, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		c.WriteJSON(GormErrorToIError(err))

		c.Close()
		return
	}

	f.RawSocketConnection = c

	go func() {
		for {
			_, k, err := c.ReadMessage()

			onRead(f, SocketReadChan{
				Data:  k,
				Error: err,
			})

			if err != nil {
				return
			}
		}
	}()

	fn(f, func(data []byte) {
		c.WriteMessage(websocket.TextMessage, data)
	})
}

func HttpPost[V any](c *gin.Context, fn func(QueryDSL) (V, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	if result, err := fn(f); err != nil {
		GinWriteContent(c, int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&f), "data": result})
	} else {
		GinWriteContent(c, 200, gin.H{
			"data": result,
		})
	}
}

func HttpPostXhtml(c *gin.Context, fn func(QueryDSL) (*XHtml, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	if result, err := fn(f); err != nil {
		GinWriteContent(c, int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&f), "data": result})
	} else {
		if result != nil {
			RenderPage(result.ScreensFs, c, result.TemplateName, result.Params)
		} else {
			c.AbortWithStatus(404)
		}
	}
}

func HttpPostEntity[T any, V any](c *gin.Context, fn func(T, QueryDSL) (V, *IError)) {
	f := ExtractQueryDslFromGinContext(c)
	id := c.Param("uniqueId")
	if id != "" {
		f.UniqueId = id
	}

	var body T
	if ReadGinRequestBodyAndCastToGoStruct(c, &body, f) {
		return
	}

	// If type of entity is string, or bool, or number, then just send the result as string
	entity, err := fn(body, f)

	switch v := any(entity).(type) {
	case string:
		c.String(http.StatusOK, v)
		return
	}

	if err != nil {
		code := int(err.HttpCode)
		if code == 0 {
			code = 403
		}

		GinWriteContent(c, code, gin.H{"error": err.ToPublicEndUser(&f), "data": entity})
		return
	}

	GinWriteContent(c, 200, gin.H{"data": entity})
}

func HttpPostEntityXhtml[T any](c *gin.Context, fn func(T, QueryDSL) (*XHtml, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	id := c.Param("uniqueId")
	if id != "" {
		f.UniqueId = id
	}

	var body T

	if ReadGinRequestBodyAndCastToGoStruct(c, &body, f) {
		return
	}

	if result, err := fn(body, f); err != nil {
		GinWriteContent(c, int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&f), "data": result})
	} else {
		if result != nil {
			RenderPage(result.ScreensFs, c, result.TemplateName, result.Params)
		}
	}
}

func HttpPostWebrtc[T any, V any](c *gin.Context, fn func(T, QueryDSL) (V, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	var body T

	if ReadGinRequestBodyAndCastToGoStruct(c, &body, f) {
		return
	}

	if entity, err := fn(body, f); err != nil {
		GinWriteContent(c, int(err.HttpCode), gin.H{"error": err.ToPublicEndUser(&f), "data": entity})
	} else {
		GinWriteContent(c, 200, gin.H{
			"data": entity,
		})
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
			"error": err.ToPublicEndUser(&f),
		})
	}

	GinStreamFromChannel(c, chanStream)

}

func getFilterQuery(qs interface{}, f *QueryDSL) {

	method := reflect.ValueOf(qs).MethodByName("GetQuery")
	if method.IsValid() {
		results := method.Call(nil) // Call the method with no arguments
		if len(results) > 0 {
			f.FilterQuery = results[0].Interface().(string)
		}
	}
}
func getJqQuery(qs interface{}, f *QueryDSL) {

	method := reflect.ValueOf(qs).MethodByName("GetJq")
	if method.IsValid() {
		results := method.Call(nil) // Call the method with no arguments
		if len(results) > 0 {
			f.JqQuery = results[0].Interface().(string)
		}
	}
}

func isYaml(headerValue string) bool {

	return slices.Contains([]string{
		"application/x-yaml",
		"application/yaml",
		"text/yaml",
		"yaml",
		"yml",
	}, headerValue)

}
func isCsv(headerValue string) bool {

	return slices.Contains([]string{
		"text/csv",
		"csv",
	}, headerValue)

}

func IsYaml(c *gin.Context) bool {
	return isYaml(c.GetHeader("Accept"))
}

func IsYamlCli(c *cli.Context) bool {
	return isYaml(c.String("x-accept"))
}

func IsCsvCli(c *cli.Context) bool {
	return isCsv(c.String("x-accept"))
}

// When done with a http handler, you can use this to write the content
// Use it for successful operations
func GinWriteContent(c *gin.Context, code int, content gin.H) {

	isYAML := IsYaml(c)

	if isYAML {
		c.Header("Content-Type", "application/x-yaml")
		c.Status(code)
		yamlData, err := yaml.Marshal(content)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "failed to marshal yaml"})
			return
		}
		c.Writer.Write(yamlData)

		return
	}

	c.JSON(code, content)
}

func getHttpF(qs interface{}, c *gin.Context) QueryDSL {
	f := ExtractQueryDslFromGinContext(c)

	if qs != nil {
		InitQueryStruct(qs)

		QueriableFieldFromGinContext(reflect.ValueOf(qs), "", c)
		DetectSelectFieldsInSQL(qs, &f)
		getFilterQuery(qs, &f)
		getJqQuery(qs, &f)
	}

	return f
}

func HttpQueryEntity[T any](
	c *gin.Context,
	fn func(query QueryDSL) ([]T, *QueryResultMeta, *IError),
	qs interface{},
) {
	f := getHttpF(qs, c)
	if items, count, err := fn(f); err != nil {
		GinWriteContent(c, 500, gin.H{"error": err})
	} else {
		result := QueryEntitySuccessResult(f, items, count)
		result["jsonQuery"] = c.Query("jsonQuery")
		GinWriteContent(c, 200, result)
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

		GinWriteContent(c, code, gin.H{
			"error": err.ToPublicEndUser(&f),
		})
	} else {

		data := PolyglotQueryHandler(item, &f)
		GinWriteContent(c, 200, gin.H{
			"data": data,
		})
	}
}

func HttpGetXHtml(
	c *gin.Context,
	fn func(QueryDSL) (*XHtml, *IError),
) {
	id := c.Param("uniqueId")
	f := ExtractQueryDslFromGinContext(c)
	f.UniqueId = id
	f.G = c

	if item, err := fn(f); err != nil {

		code := http.StatusBadRequest

		if err.HttpCode > 0 {
			code = int(err.HttpCode)
		}

		GinWriteContent(c, code, gin.H{
			"error": err.ToPublicEndUser(&f),
		})
	} else {
		if item != nil {
			RenderPage(item.ScreensFs, c, item.TemplateName, item.Params)
		} else {
			c.AbortWithStatus(404)
		}
	}
}

func toBytes(v any) []byte {
	switch t := v.(type) {
	case []byte:
		return t
	case string:
		return []byte(t)
	default:
		b, _ := json.Marshal(t)
		return b
	}
}

func WriteResponse(c *gin.Context, status int, resp emigo.EmiActionResult) {
	payload := resp.GetPayload()

	if payload == nil {
		c.Status(status)
		return
	}

	headers := resp.GetRespHeaders()
	if headers == nil {
		c.JSON(status, payload)

		return
	}
	switch p := payload.(type) {
	case func(io.Writer) error:
		// template streaming
		if err := p(c.Writer); err != nil {
			c.Error(err)
			return
		}

	case io.Reader:
		// file or image streaming
		_, err := io.Copy(c.Writer, p)
		if err != nil {
			c.Error(err)
			return
		}

	default:
		// fallback based on content-type
		ct := headers["Content-Type"]
		c.Data(status, ct, toBytes(payload))

	}
}

type GinLanguageTmp struct {
	g *gin.Context
}

func (x *GinLanguageTmp) GetLanguage() string {
	return GetAcceptFromGinHeaders(x.g)
}

func WriteActionResponseToGin(m *gin.Context, resp emigo.EmiActionResult, err error) {
	j := GinLanguageTmp{g: m}

	if err != nil && !reflect.ValueOf(err).IsNil() {

		if ierr, ok := err.(*IError); ok {
			code := int(ierr.HttpCode)
			if code == 0 {
				code = 403
			}

			m.JSON(http.StatusInternalServerError, gin.H{
				"error": ierr.ToPublicEndUser(&j),
			})
			return
		}

		m.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If the handler returned nil (and no error), it means the response was handled manually.
	if resp == nil || reflect.ValueOf(resp).IsNil() {
		return
	}

	// Apply headers
	for k, v := range resp.GetRespHeaders() {
		m.Header(k, v)
	}

	// Apply status and payload
	status := resp.GetStatusCode()
	if status == 0 {
		status = http.StatusOK
	}

	// Now here it's a bit weird. Because GetPayload returns an interface, we need to check the content type,
	// and call appropriate options. How can we extend this function, to easily add text, stream, json, yaml,
	// based on reading resp.Headers
	WriteResponse(m, status, resp)
}

package fireback

import (
	"net/http"
	reflect "reflect"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func HttpUpdateEntity[T any, V any](c *gin.Context, fn func(QueryDSL, T) (V, *IError)) {
	f := ExtractQueryDslFromGinContext(c)

	var body T

	if ReadGinRequestBodyAndCastToGoStruct(c, &body, f) {
		return
	}

	if entity, err := fn(f, body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.ToPublicEndUser(&f)})
	} else {
		c.JSON(200, gin.H{
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.ToPublicEndUser(&f)})
	} else {
		c.JSON(200, gin.H{
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
		c.JSON(400, gin.H{
			"error": err.ToPublicEndUser(&f),
		})
	} else {
		c.JSON(200, gin.H{
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

func HttpSocketRequest(ctx *gin.Context, fn func(QueryDSL, func(string)), onRead func(QueryDSL, interface{})) {
	f := ExtractQueryDslFromGinContext(ctx)

	c, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		c.WriteJSON(GormErrorToIError(err))

		c.Close()
		return
	}

	go func() {
		for {
			_, k, err := c.ReadMessage()
			if err != nil {
				return
			}
			onRead(f, string(k))
		}
	}()

	fn(f, func(data string) {

		c.WriteMessage(websocket.TextMessage, []byte(data))
	})
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

	if entity, err := fn(body, f); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.ToPublicEndUser(&f), "data": entity})
	} else {
		c.JSON(200, gin.H{
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
func HttpQueryEntity[T any](
	c *gin.Context,
	fn func(query QueryDSL) ([]T, *QueryResultMeta, error),
	qs interface{},
) {

	f := ExtractQueryDslFromGinContext(c)

	QueriableFieldFromGinContext(reflect.ValueOf(qs), "", c)

	method := reflect.ValueOf(qs).MethodByName("GetQuery")
	if method.IsValid() {
		results := method.Call(nil) // Call the method with no arguments

		// Check if it returns at least one result
		if len(results) > 0 {
			f.Query = results[0].Interface().(string)
		}
	}

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
			"error": err.ToPublicEndUser(&f),
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
			"error": err.ToPublicEndUser(&f),
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
			"error": err.ToPublicEndUser(&f),
		})

	} else {

		data := PolyglotQueryHandler(item, &f)

		c.JSON(200, gin.H{
			"data": data,
		})
	}
}

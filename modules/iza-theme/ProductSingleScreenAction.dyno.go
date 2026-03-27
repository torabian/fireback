package izaTheme

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/**
* Action to communicate with the action ProductSingleScreenAction
 */
func ProductSingleScreenActionMeta() struct {
	Name        string
	CliName     string
	URL         string
	Method      string
	Description string
} {
	return struct {
		Name        string
		CliName     string
		URL         string
		Method      string
		Description string
	}{
		Name:        "ProductSingleScreenAction",
		CliName:     "product-single-screen-action",
		URL:         "/product/:id",
		Method:      "GET",
		Description: `When a user opens a product single screen`,
	}
}

type ProductSingleScreenActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *ProductSingleScreenActionResponse) SetContentType(contentType string) *ProductSingleScreenActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ProductSingleScreenActionResponse) AsStream(r io.Reader, contentType string) *ProductSingleScreenActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ProductSingleScreenActionResponse) AsJSON(payload any) *ProductSingleScreenActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *ProductSingleScreenActionResponse) AsHTML(payload string) *ProductSingleScreenActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ProductSingleScreenActionResponse) AsBytes(payload []byte) *ProductSingleScreenActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ProductSingleScreenActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ProductSingleScreenActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ProductSingleScreenActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ProductSingleScreenActionRaw registers a raw Gin route for the ProductSingleScreenAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ProductSingleScreenActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ProductSingleScreenActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ProductSingleScreenActionRequestSig = func(c ProductSingleScreenActionRequest) (*ProductSingleScreenActionResponse, error)

// ProductSingleScreenActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ProductSingleScreenAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ProductSingleScreenActionHandler(
	handler ProductSingleScreenActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ProductSingleScreenActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := ProductSingleScreenActionRequest{
			Params:      ProductSingleScreenActionPathParameterFromGin(m),
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
		resp, err := handler(req)
		if err != nil {
			m.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), it means the response was handled manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			m.Header(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			m.JSON(status, resp.Payload)
		} else {
			m.Status(status)
		}
	}
}

// ProductSingleScreenAction is a high-level convenience wrapper around ProductSingleScreenActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ProductSingleScreenActionGin(r gin.IRoutes, handler ProductSingleScreenActionRequestSig) {
	method, url, h := ProductSingleScreenActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Path parameters for ProductSingleScreenAction
 */
type ProductSingleScreenActionPathParameter struct {
	Id string
}

// Converts a placeholder url, and applies the parameters to it.
func ProductSingleScreenActionPathParameterApply(params ProductSingleScreenActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, "id", fmt.Sprintf("%v", params.Id))
	return templateUrl
}

// Creates the parameters from the gin
// Creates the parameters from the gin
func ProductSingleScreenActionPathParameterFromGin(g *gin.Context) ProductSingleScreenActionPathParameter {
	res := ProductSingleScreenActionPathParameter{}
	res.Id = g.Param("id")
	return res
}

/**
 * Query parameters for ProductSingleScreenAction
 */
// Query wrapper with private fields
type ProductSingleScreenActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ProductSingleScreenActionQueryFromString(rawQuery string) ProductSingleScreenActionQuery {
	v := ProductSingleScreenActionQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	v.values = values
	v.mapped = mapped
	return v
}
func ProductSingleScreenActionQueryFromGin(c *gin.Context) ProductSingleScreenActionQuery {
	return ProductSingleScreenActionQueryFromString(c.Request.URL.RawQuery)
}
func ProductSingleScreenActionQueryFromHttp(r *http.Request) ProductSingleScreenActionQuery {
	return ProductSingleScreenActionQueryFromString(r.URL.RawQuery)
}
func (q ProductSingleScreenActionQuery) Values() url.Values {
	return q.values
}
func (q ProductSingleScreenActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ProductSingleScreenActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ProductSingleScreenActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ProductSingleScreenActionRequest struct {
	Params      ProductSingleScreenActionPathParameter
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type ProductSingleScreenActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ProductSingleScreenActionCall(
	req ProductSingleScreenActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ProductSingleScreenActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ProductSingleScreenActionMeta()
		baseURL := meta.URL
		// Build final URL with query string
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		// if UrlValues present, encode and append
		if len(req.QueryParams) > 0 {
			u.RawQuery = req.QueryParams.Encode()
		}
		req0, err := http.NewRequest(meta.Method, u.String(), nil)
		if err != nil {
			return nil, err
		}
		httpReq = req0
	} else {
		httpReq = config.Httpr
	}
	httpReq.Header = req.Headers
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var result ProductSingleScreenActionResult
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if resp.StatusCode >= 400 {
		return &result, fmt.Errorf("request failed: %s", respBody)
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}

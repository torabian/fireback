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
)

/**
* Action to communicate with the action HomePageAction
 */
func HomePageActionMeta() struct {
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
		Name:        "HomePageAction",
		CliName:     "home-page-action",
		URL:         "/",
		Method:      "GET",
		Description: `Homepage of the shop`,
	}
}

type HomePageActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *HomePageActionResponse) SetContentType(contentType string) *HomePageActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *HomePageActionResponse) AsStream(r io.Reader, contentType string) *HomePageActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *HomePageActionResponse) AsJSON(payload any) *HomePageActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *HomePageActionResponse) AsHTML(payload string) *HomePageActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *HomePageActionResponse) AsBytes(payload []byte) *HomePageActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x HomePageActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x HomePageActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x HomePageActionResponse) GetPayload() interface{} {
	return x.Payload
}

// HomePageActionRaw registers a raw Gin route for the HomePageAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func HomePageActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := HomePageActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type HomePageActionRequestSig = func(c HomePageActionRequest) (*HomePageActionResponse, error)

// HomePageActionHandler returns the HTTP method, route URL, and a typed Gin handler for the HomePageAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func HomePageActionHandler(
	handler HomePageActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := HomePageActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := HomePageActionRequest{
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

// HomePageAction is a high-level convenience wrapper around HomePageActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func HomePageActionGin(r gin.IRoutes, handler HomePageActionRequestSig) {
	method, url, h := HomePageActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for HomePageAction
 */
// Query wrapper with private fields
type HomePageActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func HomePageActionQueryFromString(rawQuery string) HomePageActionQuery {
	v := HomePageActionQuery{}
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
func HomePageActionQueryFromGin(c *gin.Context) HomePageActionQuery {
	return HomePageActionQueryFromString(c.Request.URL.RawQuery)
}
func HomePageActionQueryFromHttp(r *http.Request) HomePageActionQuery {
	return HomePageActionQueryFromString(r.URL.RawQuery)
}
func (q HomePageActionQuery) Values() url.Values {
	return q.values
}
func (q HomePageActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *HomePageActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *HomePageActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type HomePageActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type HomePageActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func HomePageActionCall(
	req HomePageActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*HomePageActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := HomePageActionMeta()
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
	var result HomePageActionResult
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

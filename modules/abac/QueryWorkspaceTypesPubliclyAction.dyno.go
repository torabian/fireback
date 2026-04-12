package abac

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
* Action to communicate with the action QueryWorkspaceTypesPubliclyAction
 */
func QueryWorkspaceTypesPubliclyActionMeta() struct {
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
		Name:        "QueryWorkspaceTypesPubliclyAction",
		CliName:     "public-types",
		URL:         "/workspace/public/types",
		Method:      "GET",
		Description: `Returns the workspaces types available in the project publicly without authentication, and the value could be used upon signup to go different route.`,
	}
}
func GetQueryWorkspaceTypesPubliclyActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "title",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "slug",
			Type: "string",
		},
	}
}
func CastQueryWorkspaceTypesPubliclyActionResFromCli(c emigo.CliCastable) QueryWorkspaceTypesPubliclyActionRes {
	data := QueryWorkspaceTypesPubliclyActionRes{}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("slug") {
		data.Slug = c.String("slug")
	}
	return data
}

// The base class definition for queryWorkspaceTypesPubliclyActionRes
type QueryWorkspaceTypesPubliclyActionRes struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	UniqueId    string `json:"uniqueId" yaml:"uniqueId"`
	Slug        string `json:"slug" yaml:"slug"`
}

func (x *QueryWorkspaceTypesPubliclyActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type QueryWorkspaceTypesPubliclyActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *QueryWorkspaceTypesPubliclyActionResponse) SetContentType(contentType string) *QueryWorkspaceTypesPubliclyActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *QueryWorkspaceTypesPubliclyActionResponse) AsStream(r io.Reader, contentType string) *QueryWorkspaceTypesPubliclyActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *QueryWorkspaceTypesPubliclyActionResponse) AsJSON(payload any) *QueryWorkspaceTypesPubliclyActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *QueryWorkspaceTypesPubliclyActionResponse) AsHTML(payload string) *QueryWorkspaceTypesPubliclyActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *QueryWorkspaceTypesPubliclyActionResponse) AsBytes(payload []byte) *QueryWorkspaceTypesPubliclyActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x QueryWorkspaceTypesPubliclyActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x QueryWorkspaceTypesPubliclyActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x QueryWorkspaceTypesPubliclyActionResponse) GetPayload() interface{} {
	return x.Payload
}

// QueryWorkspaceTypesPubliclyActionRaw registers a raw Gin route for the QueryWorkspaceTypesPubliclyAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func QueryWorkspaceTypesPubliclyActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := QueryWorkspaceTypesPubliclyActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type QueryWorkspaceTypesPubliclyActionRequestSig = func(c QueryWorkspaceTypesPubliclyActionRequest) (*QueryWorkspaceTypesPubliclyActionResponse, error)

// QueryWorkspaceTypesPubliclyActionHandler returns the HTTP method, route URL, and a typed Gin handler for the QueryWorkspaceTypesPubliclyAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func QueryWorkspaceTypesPubliclyActionHandler(
	handler QueryWorkspaceTypesPubliclyActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := QueryWorkspaceTypesPubliclyActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := QueryWorkspaceTypesPubliclyActionRequest{
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

// QueryWorkspaceTypesPubliclyAction is a high-level convenience wrapper around QueryWorkspaceTypesPubliclyActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func QueryWorkspaceTypesPubliclyActionGin(r gin.IRoutes, handler QueryWorkspaceTypesPubliclyActionRequestSig) {
	method, url, h := QueryWorkspaceTypesPubliclyActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for QueryWorkspaceTypesPubliclyAction
 */
// Query wrapper with private fields
type QueryWorkspaceTypesPubliclyActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func QueryWorkspaceTypesPubliclyActionQueryFromString(rawQuery string) QueryWorkspaceTypesPubliclyActionQuery {
	v := QueryWorkspaceTypesPubliclyActionQuery{}
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
func QueryWorkspaceTypesPubliclyActionQueryFromGin(c *gin.Context) QueryWorkspaceTypesPubliclyActionQuery {
	return QueryWorkspaceTypesPubliclyActionQueryFromString(c.Request.URL.RawQuery)
}
func QueryWorkspaceTypesPubliclyActionQueryFromHttp(r *http.Request) QueryWorkspaceTypesPubliclyActionQuery {
	return QueryWorkspaceTypesPubliclyActionQueryFromString(r.URL.RawQuery)
}
func (q QueryWorkspaceTypesPubliclyActionQuery) Values() url.Values {
	return q.values
}
func (q QueryWorkspaceTypesPubliclyActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *QueryWorkspaceTypesPubliclyActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *QueryWorkspaceTypesPubliclyActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type QueryWorkspaceTypesPubliclyActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type QueryWorkspaceTypesPubliclyActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func QueryWorkspaceTypesPubliclyActionCall(
	req QueryWorkspaceTypesPubliclyActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*QueryWorkspaceTypesPubliclyActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := QueryWorkspaceTypesPubliclyActionMeta()
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
	var result QueryWorkspaceTypesPubliclyActionResult
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

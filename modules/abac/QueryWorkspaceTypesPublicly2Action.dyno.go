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
* Action to communicate with the action QueryWorkspaceTypesPublicly2Action
 */
func QueryWorkspaceTypesPublicly2ActionMeta() struct {
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
		Name:        "QueryWorkspaceTypesPublicly2Action",
		CliName:     "public-types",
		URL:         "/workspace/public/types2",
		Method:      "GET",
		Description: `Returns the workspaces types available in the project publicly without authentication, and the value could be used upon signup to go different route.`,
	}
}
func GetQueryWorkspaceTypesPublicly2ActionResCliFlags(prefix string) []emigo.CliFlag {
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
func CastQueryWorkspaceTypesPublicly2ActionResFromCli(c emigo.CliCastable) QueryWorkspaceTypesPublicly2ActionRes {
	data := QueryWorkspaceTypesPublicly2ActionRes{}
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

// The base class definition for queryWorkspaceTypesPublicly2ActionRes
type QueryWorkspaceTypesPublicly2ActionRes struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	UniqueId    string `json:"uniqueId" yaml:"uniqueId"`
	Slug        string `json:"slug" yaml:"slug"`
}

func (x *QueryWorkspaceTypesPublicly2ActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type QueryWorkspaceTypesPublicly2ActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *QueryWorkspaceTypesPublicly2ActionResponse) SetContentType(contentType string) *QueryWorkspaceTypesPublicly2ActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *QueryWorkspaceTypesPublicly2ActionResponse) AsStream(r io.Reader, contentType string) *QueryWorkspaceTypesPublicly2ActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *QueryWorkspaceTypesPublicly2ActionResponse) AsJSON(payload any) *QueryWorkspaceTypesPublicly2ActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *QueryWorkspaceTypesPublicly2ActionResponse) AsHTML(payload string) *QueryWorkspaceTypesPublicly2ActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *QueryWorkspaceTypesPublicly2ActionResponse) AsBytes(payload []byte) *QueryWorkspaceTypesPublicly2ActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x QueryWorkspaceTypesPublicly2ActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x QueryWorkspaceTypesPublicly2ActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x QueryWorkspaceTypesPublicly2ActionResponse) GetPayload() interface{} {
	return x.Payload
}

// QueryWorkspaceTypesPublicly2ActionRaw registers a raw Gin route for the QueryWorkspaceTypesPublicly2Action action.
// This gives the developer full control over middleware, handlers, and response handling.
func QueryWorkspaceTypesPublicly2ActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := QueryWorkspaceTypesPublicly2ActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type QueryWorkspaceTypesPublicly2ActionRequestSig = func(c QueryWorkspaceTypesPublicly2ActionRequest) (*QueryWorkspaceTypesPublicly2ActionResponse, error)

// QueryWorkspaceTypesPublicly2ActionHandler returns the HTTP method, route URL, and a typed Gin handler for the QueryWorkspaceTypesPublicly2Action action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func QueryWorkspaceTypesPublicly2ActionHandler(
	handler QueryWorkspaceTypesPublicly2ActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := QueryWorkspaceTypesPublicly2ActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := QueryWorkspaceTypesPublicly2ActionRequest{
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

// QueryWorkspaceTypesPublicly2Action is a high-level convenience wrapper around QueryWorkspaceTypesPublicly2ActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func QueryWorkspaceTypesPublicly2ActionGin(r gin.IRoutes, handler QueryWorkspaceTypesPublicly2ActionRequestSig) {
	method, url, h := QueryWorkspaceTypesPublicly2ActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for QueryWorkspaceTypesPublicly2Action
 */
// Query wrapper with private fields
type QueryWorkspaceTypesPublicly2ActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func QueryWorkspaceTypesPublicly2ActionQueryFromString(rawQuery string) QueryWorkspaceTypesPublicly2ActionQuery {
	v := QueryWorkspaceTypesPublicly2ActionQuery{}
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
func QueryWorkspaceTypesPublicly2ActionQueryFromGin(c *gin.Context) QueryWorkspaceTypesPublicly2ActionQuery {
	return QueryWorkspaceTypesPublicly2ActionQueryFromString(c.Request.URL.RawQuery)
}
func QueryWorkspaceTypesPublicly2ActionQueryFromHttp(r *http.Request) QueryWorkspaceTypesPublicly2ActionQuery {
	return QueryWorkspaceTypesPublicly2ActionQueryFromString(r.URL.RawQuery)
}
func (q QueryWorkspaceTypesPublicly2ActionQuery) Values() url.Values {
	return q.values
}
func (q QueryWorkspaceTypesPublicly2ActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *QueryWorkspaceTypesPublicly2ActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *QueryWorkspaceTypesPublicly2ActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type QueryWorkspaceTypesPublicly2ActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type QueryWorkspaceTypesPublicly2ActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func QueryWorkspaceTypesPublicly2ActionCall(
	req QueryWorkspaceTypesPublicly2ActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*QueryWorkspaceTypesPublicly2ActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := QueryWorkspaceTypesPublicly2ActionMeta()
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
	var result QueryWorkspaceTypesPublicly2ActionResult
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

package abac

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action QueryWorkspaceTypesPubliclyAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of QueryWorkspaceTypesPubliclyAction
func QueryWorkspaceTypesPubliclyAction(c QueryWorkspaceTypesPubliclyActionRequest) (*QueryWorkspaceTypesPubliclyActionResponse, error) {
	return &QueryWorkspaceTypesPubliclyActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
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
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
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

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *QueryWorkspaceTypesPubliclyActionResponse) WithIdeal(payload QueryWorkspaceTypesPubliclyActionRes) *QueryWorkspaceTypesPubliclyActionResponse {
	x.Payload = payload
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
			Body:        nil,
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
	Body        interface{}
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	GinCtx *gin.Context
	// Urfave context, per each request
	CliCtx *cli.Command
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func (x QueryWorkspaceTypesPubliclyActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x QueryWorkspaceTypesPubliclyActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type QueryWorkspaceTypesPubliclyActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func QueryWorkspaceTypesPubliclyActionClientCreateUrl(
	req QueryWorkspaceTypesPubliclyActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := QueryWorkspaceTypesPubliclyActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// Build final URL with query string
	u, err := url.Parse(urlAddr)
	if err != nil {
		return nil, err
	}
	// if UrlValues present, encode and append
	if len(req.QueryParams) > 0 {
		u.RawQuery = req.QueryParams.Encode()
	}
	return u, nil
}
func QueryWorkspaceTypesPubliclyActionClientExecuteTyped(httpReq *http.Request) (*QueryWorkspaceTypesPubliclyActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result QueryWorkspaceTypesPubliclyActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &QueryWorkspaceTypesPubliclyActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &QueryWorkspaceTypesPubliclyActionResponse{Payload: result}, err
	}
	return &QueryWorkspaceTypesPubliclyActionResponse{Payload: result}, nil
}
func QueryWorkspaceTypesPubliclyActionClientBuildRequest(req QueryWorkspaceTypesPubliclyActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := QueryWorkspaceTypesPubliclyActionMeta()
	httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header = make(http.Header)
	// copy defaults
	for k, v := range config.Headers {
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
	}
	// override with request-specific headers
	for k, v := range req.Headers {
		httpReq.Header.Del(k) // ensure override, not duplicate
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
	}
	return httpReq, nil
}
func QueryWorkspaceTypesPubliclyActionCall(
	req QueryWorkspaceTypesPubliclyActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*QueryWorkspaceTypesPubliclyActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := QueryWorkspaceTypesPubliclyActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := QueryWorkspaceTypesPubliclyActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return QueryWorkspaceTypesPubliclyActionClientExecuteTyped(r)
}

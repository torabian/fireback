package abac

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action QueryUserRoleWorkspacesAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of QueryUserRoleWorkspacesAction
func QueryUserRoleWorkspacesAction(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error) {
	return &QueryUserRoleWorkspacesActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func QueryUserRoleWorkspacesActionMeta() struct {
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
		Name:        "QueryUserRoleWorkspacesAction",
		CliName:     "urw",
		URL:         "/urw/query",
		Method:      "GET",
		Description: `Returns the workspaces that user belongs to, as well as his role in there, and the permissions for each role`,
	}
}
func GetQueryUserRoleWorkspacesActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name:        prefix + "capabilities",
			Type:        "slice",
			Description: "Workspace level capabilities which are available",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "roles",
			Type: "array",
		},
	}
}
func CastQueryUserRoleWorkspacesActionResFromCli(c emigo.CliCastable) QueryUserRoleWorkspacesActionRes {
	data := QueryUserRoleWorkspacesActionRes{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("capabilities") {
		emigo.InflatePossibleSlice(c.String("capabilities"), &data.Capabilities)
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("roles") {
		data.Roles = emigo.CapturePossibleArray(CastQueryUserRoleWorkspacesActionResRolesFromCli, "roles", c)
	}
	return data
}

// The base class definition for queryUserRoleWorkspacesActionRes
type QueryUserRoleWorkspacesActionRes struct {
	Name string `json:"name" yaml:"name"`
	// Workspace level capabilities which are available
	Capabilities []string                                           `json:"capabilities" yaml:"capabilities"`
	UniqueId     string                                             `json:"uniqueId" yaml:"uniqueId"`
	Roles        emigo.Array[QueryUserRoleWorkspacesActionResRoles] `json:"roles" yaml:"roles"`
}

func GetQueryUserRoleWorkspacesActionResRolesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name:        prefix + "capabilities",
			Type:        "slice",
			Description: "Capabilities related to this role which are available",
		},
	}
}
func CastQueryUserRoleWorkspacesActionResRolesFromCli(c emigo.CliCastable) QueryUserRoleWorkspacesActionResRoles {
	data := QueryUserRoleWorkspacesActionResRoles{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("capabilities") {
		emigo.InflatePossibleSlice(c.String("capabilities"), &data.Capabilities)
	}
	return data
}

// The base class definition for roles
type QueryUserRoleWorkspacesActionResRoles struct {
	Name     string `json:"name" yaml:"name"`
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// Capabilities related to this role which are available
	Capabilities []string `json:"capabilities" yaml:"capabilities"`
}

func (x *QueryUserRoleWorkspacesActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type QueryUserRoleWorkspacesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *QueryUserRoleWorkspacesActionResponse) SetContentType(contentType string) *QueryUserRoleWorkspacesActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *QueryUserRoleWorkspacesActionResponse) AsStream(r io.Reader, contentType string) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *QueryUserRoleWorkspacesActionResponse) AsJSON(payload any) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *QueryUserRoleWorkspacesActionResponse) WithIdeal(payload QueryUserRoleWorkspacesActionRes) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = payload
	return x
}
func (x *QueryUserRoleWorkspacesActionResponse) AsHTML(payload string) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *QueryUserRoleWorkspacesActionResponse) AsBytes(payload []byte) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x QueryUserRoleWorkspacesActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x QueryUserRoleWorkspacesActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x QueryUserRoleWorkspacesActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type QueryUserRoleWorkspacesActionRequestSig = func(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error)

/**
 * Query parameters for QueryUserRoleWorkspacesAction
 */
// Query wrapper with private fields
type QueryUserRoleWorkspacesActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func QueryUserRoleWorkspacesActionQueryFromString(rawQuery string) QueryUserRoleWorkspacesActionQuery {
	v := QueryUserRoleWorkspacesActionQuery{}
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
func QueryUserRoleWorkspacesActionQueryFromHttp(r *http.Request) QueryUserRoleWorkspacesActionQuery {
	return QueryUserRoleWorkspacesActionQueryFromString(r.URL.RawQuery)
}
func (q QueryUserRoleWorkspacesActionQuery) Values() url.Values {
	return q.values
}
func (q QueryUserRoleWorkspacesActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *QueryUserRoleWorkspacesActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *QueryUserRoleWorkspacesActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type QueryUserRoleWorkspacesActionRequest struct {
	Body        interface{}
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	// Now it's interface, so the code gen doesn't depend on the instance
	// or gin package. Make sure you cast is later into *gin.Context, or whatever
	// your framework is passing when creating a request.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	GinCtx interface{}
	// Cli library helper (urfave) by default. The instance is interface{}, and you
	// need to manually cast it to the *cli.Command, so gives you freedom and independence
	// of external library.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	CliCtx interface{}
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func QueryUserRoleWorkspacesActionClientCreateUrl(
	req QueryUserRoleWorkspacesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := QueryUserRoleWorkspacesActionMeta()
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
func QueryUserRoleWorkspacesActionClientExecuteTyped(httpReq *http.Request) (*QueryUserRoleWorkspacesActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result QueryUserRoleWorkspacesActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &QueryUserRoleWorkspacesActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &QueryUserRoleWorkspacesActionResponse{Payload: result}, err
	}
	return &QueryUserRoleWorkspacesActionResponse{Payload: result}, nil
}
func QueryUserRoleWorkspacesActionClientBuildRequest(req QueryUserRoleWorkspacesActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := QueryUserRoleWorkspacesActionMeta()
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
func QueryUserRoleWorkspacesActionCall(
	req QueryUserRoleWorkspacesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*QueryUserRoleWorkspacesActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := QueryUserRoleWorkspacesActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := QueryUserRoleWorkspacesActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return QueryUserRoleWorkspacesActionClientExecuteTyped(r)
}

// QueryUserRoleWorkspacesActionRaw registers a raw Gin route for the QueryUserRoleWorkspacesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func QueryUserRoleWorkspacesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := QueryUserRoleWorkspacesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// QueryUserRoleWorkspacesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the QueryUserRoleWorkspacesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func QueryUserRoleWorkspacesActionHandler(
	handler func(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := QueryUserRoleWorkspacesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := QueryUserRoleWorkspacesActionRequest{
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

// QueryUserRoleWorkspacesActionGin is a high-level convenience wrapper around QueryUserRoleWorkspacesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func QueryUserRoleWorkspacesActionGin(r gin.IRoutes, handler func(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error)) {
	method, url, h := QueryUserRoleWorkspacesActionHandler(handler)
	r.Handle(method, url, h)
}
func (x QueryUserRoleWorkspacesActionRequest) IsGin() bool {
	if x.GinCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.GinCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}
func QueryUserRoleWorkspacesActionQueryFromGin(c *gin.Context) QueryUserRoleWorkspacesActionQuery {
	return QueryUserRoleWorkspacesActionQueryFromString(c.Request.URL.RawQuery)
}

// QueryUserRoleWorkspacesActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the QueryUserRoleWorkspacesAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *QueryUserRoleWorkspacesActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func QueryUserRoleWorkspacesActionHttpHandler(
	handler func(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := QueryUserRoleWorkspacesActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := QueryUserRoleWorkspacesActionRequest{
			Body:        nil,
			QueryParams: r.URL.Query(),
			Headers:     r.Header,
		}
		resp, err := handler(req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), the response was handled
		// manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			if w.Header().Get("Content-Type") == "" {
				w.Header().Set("Content-Type", "application/json")
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(resp.Payload)
		} else {
			w.WriteHeader(status)
		}
	}
}

// QueryUserRoleWorkspacesActionHttp is a high-level convenience wrapper around
// QueryUserRoleWorkspacesActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func QueryUserRoleWorkspacesActionHttp(
	mux *http.ServeMux,
	handler func(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error),
) {
	method, pattern, h := QueryUserRoleWorkspacesActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

package suggestion

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action QueryAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of QueryAction
func QueryAction(c QueryActionRequest) (*QueryActionResponse, error) {
	return &QueryActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func QueryActionMeta() struct {
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
		Name:        "QueryAction",
		CliName:     "query-action",
		URL:         "/suggestion/query",
		Method:      "POST",
		Description: `The final result of the query, it is a list of content entities based on the search. It is a list of content entities based on the search.`,
	}
}

// The base class definition for queryActionReq
type QueryActionReq struct {
	// The number of items to return.
	ItemsPerPage int `json:"itemsPerPage" yaml:"itemsPerPage"`
	// The index of the first item to return.
	StartIndex int `json:"startIndex" yaml:"startIndex"`
	// The query to search for. It is a string that will be used to search for the content.
	Phrase string `json:"phrase" yaml:"phrase"`
}

func (x *QueryActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// The base class definition for queryActionRes
type QueryActionRes struct {
	Items emigo.Array[QueryActionResItems] `json:"items" yaml:"items"`
}

// The base class definition for items
type QueryActionResItems struct {
	// The title of the content.
	Title string `json:"title" yaml:"title"`
	// The excerpt
	Excerpt string `json:"excerpt" yaml:"excerpt"`
	// The content type group
	ContentType string `json:"contentType" yaml:"contentType"`
}

func (x *QueryActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type QueryActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *QueryActionResponse) SetContentType(contentType string) *QueryActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *QueryActionResponse) AsStream(r io.Reader, contentType string) *QueryActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *QueryActionResponse) AsJSON(payload any) *QueryActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *QueryActionResponse) WithIdeal(payload QueryActionRes) *QueryActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *QueryActionResponse) AsIdeal() (*QueryActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res QueryActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *QueryActionResponse) AsHTML(payload string) *QueryActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *QueryActionResponse) AsBytes(payload []byte) *QueryActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x QueryActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x QueryActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x QueryActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type QueryActionRequestSig = func(c QueryActionRequest) (*QueryActionResponse, error)

/**
 * Query parameters for QueryAction
 */
// Query wrapper with private fields
type QueryActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func QueryActionQueryFromString(rawQuery string) QueryActionQuery {
	v := QueryActionQuery{}
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
func QueryActionQueryFromHttp(r *http.Request) QueryActionQuery {
	return QueryActionQueryFromString(r.URL.RawQuery)
}
func (q QueryActionQuery) Values() url.Values {
	return q.values
}
func (q QueryActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *QueryActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *QueryActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type QueryActionRequest struct {
	Body        QueryActionReq
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

// Returns the gin ctx. You need to manually cast this to .(*gin.Context)
func (x QueryActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x QueryActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func QueryActionClientCreateUrl(
	req QueryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := QueryActionMeta()
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
func QueryActionClientExecuteTyped(httpReq *http.Request) (*QueryActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result QueryActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}
func QueryActionClientBuildRequest(req QueryActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := QueryActionMeta()
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), bytes.NewReader(bodyBytes))
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
func QueryActionCall(
	req QueryActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*QueryActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := QueryActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := QueryActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return QueryActionClientExecuteTyped(r)
}

// QueryActionRaw registers a raw Gin route for the QueryAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func QueryActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := QueryActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// QueryActionHandler returns the HTTP method, route URL, and a typed Gin handler for the QueryAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func QueryActionHandler(
	handler func(c QueryActionRequest) (*QueryActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := QueryActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body QueryActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := QueryActionRequest{
			Body:        body,
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

// QueryActionGin is a high-level convenience wrapper around QueryActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func QueryActionGin(r gin.IRoutes, handler func(c QueryActionRequest) (*QueryActionResponse, error)) {
	method, url, h := QueryActionHandler(handler)
	r.Handle(method, url, h)
}
func (x QueryActionRequest) IsGin() bool {
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
func QueryActionQueryFromGin(c *gin.Context) QueryActionQuery {
	return QueryActionQueryFromString(c.Request.URL.RawQuery)
}

// QueryActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the QueryAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *QueryActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func QueryActionHttpHandler(
	handler func(c QueryActionRequest) (*QueryActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := QueryActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body QueryActionReq
		if r.Body != nil {
			defer r.Body.Close()
			if data, _ := io.ReadAll(r.Body); len(data) > 0 {
				if err := json.Unmarshal(data, &body); err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON: " + err.Error()})
					return
				}
			}
		}
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := QueryActionRequest{
			Body:        body,
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

// QueryActionHttp is a high-level convenience wrapper around
// QueryActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func QueryActionHttp(
	mux *http.ServeMux,
	handler func(c QueryActionRequest) (*QueryActionResponse, error),
) {
	method, pattern, h := QueryActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

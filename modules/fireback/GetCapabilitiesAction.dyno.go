package fireback

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
* Action to communicate with the action GetCapabilitiesAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetCapabilitiesAction
func GetCapabilitiesAction(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error) {
	return &GetCapabilitiesActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetCapabilitiesActionMeta() struct {
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
		Name:        "GetCapabilitiesAction",
		CliName:     "get-capabilities-action",
		URL:         "/capabilities2",
		Method:      "GET",
		Description: `Get the capabilities inside a system`,
	}
}

// The base class definition for getCapabilitiesActionRes
type GetCapabilitiesActionRes struct {
	Capabilities emigo.Collection[CapabilityInfoDto] `json:"capabilities" yaml:"capabilities"`
	Nested       emigo.Collection[CapabilityInfoDto] `json:"nested" yaml:"nested"`
}

func (x *GetCapabilitiesActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetCapabilitiesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetCapabilitiesActionResponse) SetContentType(contentType string) *GetCapabilitiesActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetCapabilitiesActionResponse) AsStream(r io.Reader, contentType string) *GetCapabilitiesActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetCapabilitiesActionResponse) AsJSON(payload any) *GetCapabilitiesActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetCapabilitiesActionResponse) WithIdeal(payload GetCapabilitiesActionRes) *GetCapabilitiesActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *GetCapabilitiesActionResponse) AsIdeal() (*GetCapabilitiesActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res GetCapabilitiesActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *GetCapabilitiesActionResponse) AsHTML(payload string) *GetCapabilitiesActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetCapabilitiesActionResponse) AsBytes(payload []byte) *GetCapabilitiesActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetCapabilitiesActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetCapabilitiesActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetCapabilitiesActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type GetCapabilitiesActionRequestSig = func(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error)

/**
 * Query parameters for GetCapabilitiesAction
 */
// Query wrapper with private fields
type GetCapabilitiesActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetCapabilitiesActionQueryFromString(rawQuery string) GetCapabilitiesActionQuery {
	v := GetCapabilitiesActionQuery{}
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
func GetCapabilitiesActionQueryFromHttp(r *http.Request) GetCapabilitiesActionQuery {
	return GetCapabilitiesActionQueryFromString(r.URL.RawQuery)
}
func (q GetCapabilitiesActionQuery) Values() url.Values {
	return q.values
}
func (q GetCapabilitiesActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetCapabilitiesActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetCapabilitiesActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetCapabilitiesActionRequest struct {
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

// Returns the gin ctx. You need to manually cast this to .(*gin.Context)
func (x GetCapabilitiesActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x GetCapabilitiesActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func GetCapabilitiesActionClientCreateUrl(
	req GetCapabilitiesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetCapabilitiesActionMeta()
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
func GetCapabilitiesActionClientExecuteTyped(httpReq *http.Request) (*GetCapabilitiesActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetCapabilitiesActionResponse
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
func GetCapabilitiesActionClientBuildRequest(req GetCapabilitiesActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetCapabilitiesActionMeta()
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
func GetCapabilitiesActionCall(
	req GetCapabilitiesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetCapabilitiesActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetCapabilitiesActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetCapabilitiesActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetCapabilitiesActionClientExecuteTyped(r)
}

// GetCapabilitiesActionRaw registers a raw Gin route for the GetCapabilitiesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetCapabilitiesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetCapabilitiesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// GetCapabilitiesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetCapabilitiesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetCapabilitiesActionHandler(
	handler func(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GetCapabilitiesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetCapabilitiesActionRequest{
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

// GetCapabilitiesActionGin is a high-level convenience wrapper around GetCapabilitiesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetCapabilitiesActionGin(r gin.IRoutes, handler func(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error)) {
	method, url, h := GetCapabilitiesActionHandler(handler)
	r.Handle(method, url, h)
}
func (x GetCapabilitiesActionRequest) IsGin() bool {
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
func GetCapabilitiesActionQueryFromGin(c *gin.Context) GetCapabilitiesActionQuery {
	return GetCapabilitiesActionQueryFromString(c.Request.URL.RawQuery)
}

// GetCapabilitiesActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the GetCapabilitiesAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *GetCapabilitiesActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func GetCapabilitiesActionHttpHandler(
	handler func(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := GetCapabilitiesActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := GetCapabilitiesActionRequest{
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

// GetCapabilitiesActionHttp is a high-level convenience wrapper around
// GetCapabilitiesActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func GetCapabilitiesActionHttp(
	mux *http.ServeMux,
	handler func(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error),
) {
	method, pattern, h := GetCapabilitiesActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

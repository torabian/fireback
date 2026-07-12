package abac

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
* Action to communicate with the action ClassicSigninAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ClassicSigninAction
func ClassicSigninAction(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error) {
	return &ClassicSigninActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ClassicSigninActionMeta() struct {
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
		Name:        "ClassicSigninAction",
		CliName:     "in",
		URL:         "/passports/signin/classic",
		Method:      "POST",
		Description: `Signin publicly to and account using class passports (email, password)`,
	}
}

// The base class definition for classicSigninActionReq
type ClassicSigninActionReq struct {
	Value    string `json:"value" validate:"required" yaml:"value"`
	Password string `json:"password" validate:"required" yaml:"password"`
	// Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.
	TotpCode string `json:"totpCode" yaml:"totpCode"`
	// Session secret when logging in to the application requires more steps to complete.
	SessionSecret string `json:"sessionSecret" yaml:"sessionSecret"`
}

func (x *ClassicSigninActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// The base class definition for classicSigninActionRes
type ClassicSigninActionRes struct {
	Session emigo.One[UserSessionDto] `json:"session" yaml:"session"`
	// The next possible action which is suggested.
	Next []string `json:"next" yaml:"next"`
	// In case the account doesn't have totp, but enforced by installation, this value will contain the link
	TotpUrl string `json:"totpUrl" yaml:"totpUrl"`
	// Returns a secret session if the authentication requires more steps.
	SessionSecret string `json:"sessionSecret" yaml:"sessionSecret"`
}

func (x *ClassicSigninActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ClassicSigninActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ClassicSigninActionResponse) SetContentType(contentType string) *ClassicSigninActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ClassicSigninActionResponse) AsStream(r io.Reader, contentType string) *ClassicSigninActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ClassicSigninActionResponse) AsJSON(payload any) *ClassicSigninActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ClassicSigninActionResponse) WithIdeal(payload ClassicSigninActionRes) *ClassicSigninActionResponse {
	x.Payload = payload
	return x
}
func (x *ClassicSigninActionResponse) AsHTML(payload string) *ClassicSigninActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ClassicSigninActionResponse) AsBytes(payload []byte) *ClassicSigninActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ClassicSigninActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ClassicSigninActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ClassicSigninActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type ClassicSigninActionRequestSig = func(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error)

/**
 * Query parameters for ClassicSigninAction
 */
// Query wrapper with private fields
type ClassicSigninActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ClassicSigninActionQueryFromString(rawQuery string) ClassicSigninActionQuery {
	v := ClassicSigninActionQuery{}
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
func ClassicSigninActionQueryFromHttp(r *http.Request) ClassicSigninActionQuery {
	return ClassicSigninActionQueryFromString(r.URL.RawQuery)
}
func (q ClassicSigninActionQuery) Values() url.Values {
	return q.values
}
func (q ClassicSigninActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ClassicSigninActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ClassicSigninActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ClassicSigninActionRequest struct {
	Body        ClassicSigninActionReq
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

func ClassicSigninActionClientCreateUrl(
	req ClassicSigninActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ClassicSigninActionMeta()
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
func ClassicSigninActionClientExecuteTyped(httpReq *http.Request) (*ClassicSigninActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ClassicSigninActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ClassicSigninActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ClassicSigninActionResponse{Payload: result}, err
	}
	return &ClassicSigninActionResponse{Payload: result}, nil
}
func ClassicSigninActionClientBuildRequest(req ClassicSigninActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ClassicSigninActionMeta()
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
func ClassicSigninActionCall(
	req ClassicSigninActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ClassicSigninActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ClassicSigninActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ClassicSigninActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ClassicSigninActionClientExecuteTyped(r)
}

// ClassicSigninActionRaw registers a raw Gin route for the ClassicSigninAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ClassicSigninActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ClassicSigninActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// ClassicSigninActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ClassicSigninAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ClassicSigninActionHandler(
	handler func(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ClassicSigninActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ClassicSigninActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ClassicSigninActionRequest{
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

// ClassicSigninActionGin is a high-level convenience wrapper around ClassicSigninActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ClassicSigninActionGin(r gin.IRoutes, handler func(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error)) {
	method, url, h := ClassicSigninActionHandler(handler)
	r.Handle(method, url, h)
}
func (x ClassicSigninActionRequest) IsGin() bool {
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
func ClassicSigninActionQueryFromGin(c *gin.Context) ClassicSigninActionQuery {
	return ClassicSigninActionQueryFromString(c.Request.URL.RawQuery)
}

// ClassicSigninActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the ClassicSigninAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *ClassicSigninActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func ClassicSigninActionHttpHandler(
	handler func(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := ClassicSigninActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body ClassicSigninActionReq
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
		req := ClassicSigninActionRequest{
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

// ClassicSigninActionHttp is a high-level convenience wrapper around
// ClassicSigninActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func ClassicSigninActionHttp(
	mux *http.ServeMux,
	handler func(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error),
) {
	method, pattern, h := ClassicSigninActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

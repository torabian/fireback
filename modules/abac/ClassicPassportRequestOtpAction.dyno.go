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
* Action to communicate with the action ClassicPassportRequestOtpAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ClassicPassportRequestOtpAction
func ClassicPassportRequestOtpAction(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error) {
	return &ClassicPassportRequestOtpActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ClassicPassportRequestOtpActionMeta() struct {
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
		Name:        "ClassicPassportRequestOtpAction",
		CliName:     "otp-request",
		URL:         "/workspace/passport/request-otp",
		Method:      "POST",
		Description: `Triggers an otp request, and will send an sms or email to the passport. This endpoint is not used for login, but rather makes a request at initial step. Later you can call classicPassportOtp to get in.`,
	}
}
func GetClassicPassportRequestOtpActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "Passport value (email, phone number) which would be receiving the otp code.",
		},
	}
}
func CastClassicPassportRequestOtpActionReqFromCli(c emigo.CliCastable) ClassicPassportRequestOtpActionReq {
	data := ClassicPassportRequestOtpActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	return data
}

// The base class definition for classicPassportRequestOtpActionReq
type ClassicPassportRequestOtpActionReq struct {
	// Passport value (email, phone number) which would be receiving the otp code.
	Value string `json:"value" validate:"required" yaml:"value"`
}

func (x *ClassicPassportRequestOtpActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetClassicPassportRequestOtpActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "suspend-until",
			Type: "int64",
		},
		{
			Name: prefix + "valid-until",
			Type: "int64",
		},
		{
			Name: prefix + "blocked-until",
			Type: "int64",
		},
		{
			Name:        prefix + "seconds-to-unblock",
			Type:        "int64",
			Description: "The amount of time left to unblock for next request",
		},
	}
}
func CastClassicPassportRequestOtpActionResFromCli(c emigo.CliCastable) ClassicPassportRequestOtpActionRes {
	data := ClassicPassportRequestOtpActionRes{}
	if c.IsSet("suspend-until") {
		data.SuspendUntil = int64(c.Int64("suspend-until"))
	}
	if c.IsSet("valid-until") {
		data.ValidUntil = int64(c.Int64("valid-until"))
	}
	if c.IsSet("blocked-until") {
		data.BlockedUntil = int64(c.Int64("blocked-until"))
	}
	if c.IsSet("seconds-to-unblock") {
		data.SecondsToUnblock = int64(c.Int64("seconds-to-unblock"))
	}
	return data
}

// The base class definition for classicPassportRequestOtpActionRes
type ClassicPassportRequestOtpActionRes struct {
	SuspendUntil int64 `json:"suspendUntil" yaml:"suspendUntil"`
	ValidUntil   int64 `json:"validUntil" yaml:"validUntil"`
	BlockedUntil int64 `json:"blockedUntil" yaml:"blockedUntil"`
	// The amount of time left to unblock for next request
	SecondsToUnblock int64 `json:"secondsToUnblock" yaml:"secondsToUnblock"`
}

func (x *ClassicPassportRequestOtpActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ClassicPassportRequestOtpActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ClassicPassportRequestOtpActionResponse) SetContentType(contentType string) *ClassicPassportRequestOtpActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ClassicPassportRequestOtpActionResponse) AsStream(r io.Reader, contentType string) *ClassicPassportRequestOtpActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ClassicPassportRequestOtpActionResponse) AsJSON(payload any) *ClassicPassportRequestOtpActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ClassicPassportRequestOtpActionResponse) WithIdeal(payload ClassicPassportRequestOtpActionRes) *ClassicPassportRequestOtpActionResponse {
	x.Payload = payload
	return x
}
func (x *ClassicPassportRequestOtpActionResponse) AsHTML(payload string) *ClassicPassportRequestOtpActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ClassicPassportRequestOtpActionResponse) AsBytes(payload []byte) *ClassicPassportRequestOtpActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ClassicPassportRequestOtpActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ClassicPassportRequestOtpActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ClassicPassportRequestOtpActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type ClassicPassportRequestOtpActionRequestSig = func(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error)

/**
 * Query parameters for ClassicPassportRequestOtpAction
 */
// Query wrapper with private fields
type ClassicPassportRequestOtpActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ClassicPassportRequestOtpActionQueryFromString(rawQuery string) ClassicPassportRequestOtpActionQuery {
	v := ClassicPassportRequestOtpActionQuery{}
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
func ClassicPassportRequestOtpActionQueryFromHttp(r *http.Request) ClassicPassportRequestOtpActionQuery {
	return ClassicPassportRequestOtpActionQueryFromString(r.URL.RawQuery)
}
func (q ClassicPassportRequestOtpActionQuery) Values() url.Values {
	return q.values
}
func (q ClassicPassportRequestOtpActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ClassicPassportRequestOtpActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ClassicPassportRequestOtpActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ClassicPassportRequestOtpActionRequest struct {
	Body        ClassicPassportRequestOtpActionReq
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

func ClassicPassportRequestOtpActionClientCreateUrl(
	req ClassicPassportRequestOtpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ClassicPassportRequestOtpActionMeta()
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
func ClassicPassportRequestOtpActionClientExecuteTyped(httpReq *http.Request) (*ClassicPassportRequestOtpActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ClassicPassportRequestOtpActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ClassicPassportRequestOtpActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ClassicPassportRequestOtpActionResponse{Payload: result}, err
	}
	return &ClassicPassportRequestOtpActionResponse{Payload: result}, nil
}
func ClassicPassportRequestOtpActionClientBuildRequest(req ClassicPassportRequestOtpActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ClassicPassportRequestOtpActionMeta()
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
func ClassicPassportRequestOtpActionCall(
	req ClassicPassportRequestOtpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ClassicPassportRequestOtpActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ClassicPassportRequestOtpActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ClassicPassportRequestOtpActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ClassicPassportRequestOtpActionClientExecuteTyped(r)
}

// ClassicPassportRequestOtpActionRaw registers a raw Gin route for the ClassicPassportRequestOtpAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ClassicPassportRequestOtpActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ClassicPassportRequestOtpActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// ClassicPassportRequestOtpActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ClassicPassportRequestOtpAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ClassicPassportRequestOtpActionHandler(
	handler func(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ClassicPassportRequestOtpActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ClassicPassportRequestOtpActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ClassicPassportRequestOtpActionRequest{
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

// ClassicPassportRequestOtpActionGin is a high-level convenience wrapper around ClassicPassportRequestOtpActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ClassicPassportRequestOtpActionGin(r gin.IRoutes, handler func(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error)) {
	method, url, h := ClassicPassportRequestOtpActionHandler(handler)
	r.Handle(method, url, h)
}
func (x ClassicPassportRequestOtpActionRequest) IsGin() bool {
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
func ClassicPassportRequestOtpActionQueryFromGin(c *gin.Context) ClassicPassportRequestOtpActionQuery {
	return ClassicPassportRequestOtpActionQueryFromString(c.Request.URL.RawQuery)
}
func (x ClassicPassportRequestOtpActionRequest) IsCli() bool {
	if x.CliCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.CliCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}

// ClassicPassportRequestOtpActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the ClassicPassportRequestOtpAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *ClassicPassportRequestOtpActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func ClassicPassportRequestOtpActionHttpHandler(
	handler func(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := ClassicPassportRequestOtpActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body ClassicPassportRequestOtpActionReq
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
		req := ClassicPassportRequestOtpActionRequest{
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

// ClassicPassportRequestOtpActionHttp is a high-level convenience wrapper around
// ClassicPassportRequestOtpActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func ClassicPassportRequestOtpActionHttp(
	mux *http.ServeMux,
	handler func(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error),
) {
	method, pattern, h := ClassicPassportRequestOtpActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

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
* Action to communicate with the action ClassicPassportOtpAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ClassicPassportOtpAction
func ClassicPassportOtpAction(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error) {
	return &ClassicPassportOtpActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func ClassicPassportOtpActionMeta() struct {
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
		Name:        "ClassicPassportOtpAction",
		CliName:     "otp",
		URL:         "/workspace/passport/otp",
		Method:      "POST",
		Description: `Authenticate the user publicly for classic methods using communication service, such as sms, call, or email. You need to call classicPassportRequestOtp beforehand to send a otp code, and then validate it with this API. Also checkClassicPassport action might already sent the otp, so make sure you don't send it twice.`,
	}
}

// The base class definition for classicPassportOtpActionReq
type ClassicPassportOtpActionReq struct {
	Value string `json:"value" validate:"required" yaml:"value"`
	Otp   string `json:"otp" validate:"required" yaml:"otp"`
}

func (x *ClassicPassportOtpActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// The base class definition for classicPassportOtpActionRes
type ClassicPassportOtpActionRes struct {
	// Upon successful authentication, there will be a session dto generated, which is a ground information of authorized user and can be stored in front-end.
	Session emigo.OneNullable[UserSessionDto] `json:"session" yaml:"session"`
	// If time based otp is available, we add it response to make it easier for ui.
	TotpUrl string `json:"totpUrl" yaml:"totpUrl"`
	// The session secret will be used to call complete user registration api.
	SessionSecret string `json:"sessionSecret" yaml:"sessionSecret"`
	// If return true, means the OTP is correct and user needs to be created before continue the authentication process.
	ContinueWithCreation bool `json:"continueWithCreation" yaml:"continueWithCreation"`
}

func (x *ClassicPassportOtpActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ClassicPassportOtpActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *ClassicPassportOtpActionResponse) SetContentType(contentType string) *ClassicPassportOtpActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ClassicPassportOtpActionResponse) AsStream(r io.Reader, contentType string) *ClassicPassportOtpActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ClassicPassportOtpActionResponse) AsJSON(payload any) *ClassicPassportOtpActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ClassicPassportOtpActionResponse) WithIdeal(payload ClassicPassportOtpActionRes) *ClassicPassportOtpActionResponse {
	x.Payload = payload
	return x
}
func (x *ClassicPassportOtpActionResponse) AsHTML(payload string) *ClassicPassportOtpActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ClassicPassportOtpActionResponse) AsBytes(payload []byte) *ClassicPassportOtpActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ClassicPassportOtpActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ClassicPassportOtpActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ClassicPassportOtpActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type ClassicPassportOtpActionRequestSig = func(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error)

/**
 * Query parameters for ClassicPassportOtpAction
 */
// Query wrapper with private fields
type ClassicPassportOtpActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ClassicPassportOtpActionQueryFromString(rawQuery string) ClassicPassportOtpActionQuery {
	v := ClassicPassportOtpActionQuery{}
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
func ClassicPassportOtpActionQueryFromHttp(r *http.Request) ClassicPassportOtpActionQuery {
	return ClassicPassportOtpActionQueryFromString(r.URL.RawQuery)
}
func (q ClassicPassportOtpActionQuery) Values() url.Values {
	return q.values
}
func (q ClassicPassportOtpActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ClassicPassportOtpActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ClassicPassportOtpActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ClassicPassportOtpActionRequest struct {
	Body        ClassicPassportOtpActionReq
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

func ClassicPassportOtpActionClientCreateUrl(
	req ClassicPassportOtpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ClassicPassportOtpActionMeta()
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
func ClassicPassportOtpActionClientExecuteTyped(httpReq *http.Request) (*ClassicPassportOtpActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ClassicPassportOtpActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ClassicPassportOtpActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ClassicPassportOtpActionResponse{Payload: result}, err
	}
	return &ClassicPassportOtpActionResponse{Payload: result}, nil
}
func ClassicPassportOtpActionClientBuildRequest(req ClassicPassportOtpActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ClassicPassportOtpActionMeta()
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
func ClassicPassportOtpActionCall(
	req ClassicPassportOtpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ClassicPassportOtpActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ClassicPassportOtpActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ClassicPassportOtpActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ClassicPassportOtpActionClientExecuteTyped(r)
}

// ClassicPassportOtpActionRaw registers a raw Gin route for the ClassicPassportOtpAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ClassicPassportOtpActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ClassicPassportOtpActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// ClassicPassportOtpActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ClassicPassportOtpAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ClassicPassportOtpActionHandler(
	handler func(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := ClassicPassportOtpActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ClassicPassportOtpActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ClassicPassportOtpActionRequest{
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

// ClassicPassportOtpActionGin is a high-level convenience wrapper around ClassicPassportOtpActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ClassicPassportOtpActionGin(r gin.IRoutes, handler func(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error)) {
	method, url, h := ClassicPassportOtpActionHandler(handler)
	r.Handle(method, url, h)
}
func (x ClassicPassportOtpActionRequest) IsGin() bool {
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
func ClassicPassportOtpActionQueryFromGin(c *gin.Context) ClassicPassportOtpActionQuery {
	return ClassicPassportOtpActionQueryFromString(c.Request.URL.RawQuery)
}

// ClassicPassportOtpActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the ClassicPassportOtpAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *ClassicPassportOtpActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func ClassicPassportOtpActionHttpHandler(
	handler func(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := ClassicPassportOtpActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body ClassicPassportOtpActionReq
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
		req := ClassicPassportOtpActionRequest{
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

// ClassicPassportOtpActionHttp is a high-level convenience wrapper around
// ClassicPassportOtpActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func ClassicPassportOtpActionHttp(
	mux *http.ServeMux,
	handler func(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error),
) {
	method, pattern, h := ClassicPassportOtpActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

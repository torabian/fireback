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
* Action to communicate with the action OauthAuthenticateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of OauthAuthenticateAction
func OauthAuthenticateAction(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error) {
	return &OauthAuthenticateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func OauthAuthenticateActionMeta() struct {
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
		Name:        "OauthAuthenticateAction",
		CliName:     "oauth-authenticate-action",
		URL:         "/passport/via-oauth",
		Method:      "POST",
		Description: `When a token is got from a oauth service such as google, we send the token here to authenticate the user. To me seems this doesn't need to have 2FA or anything, so we return the session directly, or maybe there needs to be next step.`,
	}
}
func GetOauthAuthenticateActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "token",
			Type:        "string",
			Description: "The token that Auth2 provider returned to the front-end, which will be used to validate the backend",
		},
		{
			Name:        prefix + "service",
			Type:        "string",
			Description: "The service name, such as 'google' which later backend will use to authorize the token and create the user.",
		},
	}
}
func CastOauthAuthenticateActionReqFromCli(c emigo.CliCastable) OauthAuthenticateActionReq {
	data := OauthAuthenticateActionReq{}
	if c.IsSet("token") {
		data.Token = c.String("token")
	}
	if c.IsSet("service") {
		data.Service = c.String("service")
	}
	return data
}

// The base class definition for oauthAuthenticateActionReq
type OauthAuthenticateActionReq struct {
	// The token that Auth2 provider returned to the front-end, which will be used to validate the backend
	Token string `json:"token" yaml:"token"`
	// The service name, such as 'google' which later backend will use to authorize the token and create the user.
	Service string `json:"service" yaml:"service"`
}

func (x *OauthAuthenticateActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetOauthAuthenticateActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "session",
			Type: "one",
		},
		{
			Name:        prefix + "next",
			Type:        "slice",
			Description: "The next possible action which is suggested.",
		},
	}
}
func CastOauthAuthenticateActionResFromCli(c emigo.CliCastable) OauthAuthenticateActionRes {
	data := OauthAuthenticateActionRes{}
	if c.IsSet("session") {
		data.Session = emigo.CapturePossibleOne(CastUserSessionDtoFromCli, "session", c)
	}
	if c.IsSet("next") {
		emigo.InflatePossibleSlice(c.String("next"), &data.Next)
	}
	return data
}

// The base class definition for oauthAuthenticateActionRes
type OauthAuthenticateActionRes struct {
	Session emigo.One[UserSessionDto] `json:"session" yaml:"session"`
	// The next possible action which is suggested.
	Next []string `json:"next" yaml:"next"`
}

func (x *OauthAuthenticateActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type OauthAuthenticateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *OauthAuthenticateActionResponse) SetContentType(contentType string) *OauthAuthenticateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *OauthAuthenticateActionResponse) AsStream(r io.Reader, contentType string) *OauthAuthenticateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *OauthAuthenticateActionResponse) AsJSON(payload any) *OauthAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *OauthAuthenticateActionResponse) WithIdeal(payload OauthAuthenticateActionRes) *OauthAuthenticateActionResponse {
	x.Payload = payload
	return x
}
func (x *OauthAuthenticateActionResponse) AsHTML(payload string) *OauthAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *OauthAuthenticateActionResponse) AsBytes(payload []byte) *OauthAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x OauthAuthenticateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x OauthAuthenticateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x OauthAuthenticateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type OauthAuthenticateActionRequestSig = func(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error)

/**
 * Query parameters for OauthAuthenticateAction
 */
// Query wrapper with private fields
type OauthAuthenticateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func OauthAuthenticateActionQueryFromString(rawQuery string) OauthAuthenticateActionQuery {
	v := OauthAuthenticateActionQuery{}
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
func OauthAuthenticateActionQueryFromHttp(r *http.Request) OauthAuthenticateActionQuery {
	return OauthAuthenticateActionQueryFromString(r.URL.RawQuery)
}
func (q OauthAuthenticateActionQuery) Values() url.Values {
	return q.values
}
func (q OauthAuthenticateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *OauthAuthenticateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *OauthAuthenticateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type OauthAuthenticateActionRequest struct {
	Body        OauthAuthenticateActionReq
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

func OauthAuthenticateActionClientCreateUrl(
	req OauthAuthenticateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := OauthAuthenticateActionMeta()
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
func OauthAuthenticateActionClientExecuteTyped(httpReq *http.Request) (*OauthAuthenticateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result OauthAuthenticateActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &OauthAuthenticateActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &OauthAuthenticateActionResponse{Payload: result}, err
	}
	return &OauthAuthenticateActionResponse{Payload: result}, nil
}
func OauthAuthenticateActionClientBuildRequest(req OauthAuthenticateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := OauthAuthenticateActionMeta()
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
func OauthAuthenticateActionCall(
	req OauthAuthenticateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*OauthAuthenticateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := OauthAuthenticateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := OauthAuthenticateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return OauthAuthenticateActionClientExecuteTyped(r)
}

// OauthAuthenticateActionRaw registers a raw Gin route for the OauthAuthenticateAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func OauthAuthenticateActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := OauthAuthenticateActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// OauthAuthenticateActionHandler returns the HTTP method, route URL, and a typed Gin handler for the OauthAuthenticateAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func OauthAuthenticateActionHandler(
	handler func(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := OauthAuthenticateActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body OauthAuthenticateActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := OauthAuthenticateActionRequest{
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

// OauthAuthenticateActionGin is a high-level convenience wrapper around OauthAuthenticateActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func OauthAuthenticateActionGin(r gin.IRoutes, handler func(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error)) {
	method, url, h := OauthAuthenticateActionHandler(handler)
	r.Handle(method, url, h)
}
func (x OauthAuthenticateActionRequest) IsGin() bool {
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
func OauthAuthenticateActionQueryFromGin(c *gin.Context) OauthAuthenticateActionQuery {
	return OauthAuthenticateActionQueryFromString(c.Request.URL.RawQuery)
}
func (x OauthAuthenticateActionRequest) IsCli() bool {
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

// OauthAuthenticateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the OauthAuthenticateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *OauthAuthenticateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func OauthAuthenticateActionHttpHandler(
	handler func(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := OauthAuthenticateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body OauthAuthenticateActionReq
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
		req := OauthAuthenticateActionRequest{
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

// OauthAuthenticateActionHttp is a high-level convenience wrapper around
// OauthAuthenticateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func OauthAuthenticateActionHttp(
	mux *http.ServeMux,
	handler func(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error),
) {
	method, pattern, h := OauthAuthenticateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

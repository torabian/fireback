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
* Action to communicate with the action UserPassportsAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of UserPassportsAction
func UserPassportsAction(c UserPassportsActionRequest) (*UserPassportsActionResponse, error) {
	return &UserPassportsActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func UserPassportsActionMeta() struct {
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
		Name:        "UserPassportsAction",
		CliName:     "user-passports-action",
		URL:         "/user/passports",
		Method:      "GET",
		Description: `Returns list of passports belongs to an specific user.`,
	}
}
func GetUserPassportsActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "The passport value, such as email address or phone number",
		},
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "Unique identifier of the passport to operate some action on top of it",
		},
		{
			Name:        prefix + "type",
			Type:        "string",
			Description: "The type of the passport, such as email, phone number",
		},
		{
			Name:        prefix + "totp-confirmed",
			Type:        "bool",
			Description: "Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.",
		},
	}
}
func CastUserPassportsActionResFromCli(c emigo.CliCastable) UserPassportsActionRes {
	data := UserPassportsActionRes{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("totp-confirmed") {
		data.TotpConfirmed = bool(c.Bool("totp-confirmed"))
	}
	return data
}

// The base class definition for userPassportsActionRes
type UserPassportsActionRes struct {
	// The passport value, such as email address or phone number
	Value string `json:"value" yaml:"value"`
	// Unique identifier of the passport to operate some action on top of it
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// The type of the passport, such as email, phone number
	Type string `json:"type" yaml:"type"`
	// Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
	TotpConfirmed bool `json:"totpConfirmed" yaml:"totpConfirmed"`
}

func (x *UserPassportsActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type UserPassportsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *UserPassportsActionResponse) SetContentType(contentType string) *UserPassportsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *UserPassportsActionResponse) AsStream(r io.Reader, contentType string) *UserPassportsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *UserPassportsActionResponse) AsJSON(payload any) *UserPassportsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *UserPassportsActionResponse) WithIdeal(payload UserPassportsActionRes) *UserPassportsActionResponse {
	x.Payload = payload
	return x
}
func (x *UserPassportsActionResponse) AsHTML(payload string) *UserPassportsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *UserPassportsActionResponse) AsBytes(payload []byte) *UserPassportsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x UserPassportsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x UserPassportsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x UserPassportsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type UserPassportsActionRequestSig = func(c UserPassportsActionRequest) (*UserPassportsActionResponse, error)

/**
 * Query parameters for UserPassportsAction
 */
// Query wrapper with private fields
type UserPassportsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func UserPassportsActionQueryFromString(rawQuery string) UserPassportsActionQuery {
	v := UserPassportsActionQuery{}
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
func UserPassportsActionQueryFromHttp(r *http.Request) UserPassportsActionQuery {
	return UserPassportsActionQueryFromString(r.URL.RawQuery)
}
func (q UserPassportsActionQuery) Values() url.Values {
	return q.values
}
func (q UserPassportsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *UserPassportsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *UserPassportsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type UserPassportsActionRequest struct {
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

func UserPassportsActionClientCreateUrl(
	req UserPassportsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := UserPassportsActionMeta()
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
func UserPassportsActionClientExecuteTyped(httpReq *http.Request) (*UserPassportsActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result UserPassportsActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &UserPassportsActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &UserPassportsActionResponse{Payload: result}, err
	}
	return &UserPassportsActionResponse{Payload: result}, nil
}
func UserPassportsActionClientBuildRequest(req UserPassportsActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := UserPassportsActionMeta()
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
func UserPassportsActionCall(
	req UserPassportsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UserPassportsActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := UserPassportsActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := UserPassportsActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return UserPassportsActionClientExecuteTyped(r)
}

// UserPassportsActionRaw registers a raw Gin route for the UserPassportsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func UserPassportsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := UserPassportsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// UserPassportsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the UserPassportsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func UserPassportsActionHandler(
	handler func(c UserPassportsActionRequest) (*UserPassportsActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := UserPassportsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := UserPassportsActionRequest{
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

// UserPassportsActionGin is a high-level convenience wrapper around UserPassportsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func UserPassportsActionGin(r gin.IRoutes, handler func(c UserPassportsActionRequest) (*UserPassportsActionResponse, error)) {
	method, url, h := UserPassportsActionHandler(handler)
	r.Handle(method, url, h)
}
func (x UserPassportsActionRequest) IsGin() bool {
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
func UserPassportsActionQueryFromGin(c *gin.Context) UserPassportsActionQuery {
	return UserPassportsActionQueryFromString(c.Request.URL.RawQuery)
}
func (x UserPassportsActionRequest) IsCli() bool {
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

// UserPassportsActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the UserPassportsAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *UserPassportsActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func UserPassportsActionHttpHandler(
	handler func(c UserPassportsActionRequest) (*UserPassportsActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := UserPassportsActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := UserPassportsActionRequest{
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

// UserPassportsActionHttp is a high-level convenience wrapper around
// UserPassportsActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func UserPassportsActionHttp(
	mux *http.ServeMux,
	handler func(c UserPassportsActionRequest) (*UserPassportsActionResponse, error),
) {
	method, pattern, h := UserPassportsActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

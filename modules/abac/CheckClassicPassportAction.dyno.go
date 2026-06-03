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
* Action to communicate with the action CheckClassicPassportAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CheckClassicPassportAction
func CheckClassicPassportAction(c CheckClassicPassportActionRequest) (*CheckClassicPassportActionResponse, error) {
	return &CheckClassicPassportActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CheckClassicPassportActionMeta() struct {
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
		Name:        "CheckClassicPassportAction",
		CliName:     "ccp",
		URL:         "/workspace/passport/check",
		Method:      "POST",
		Description: `Checks if a classic passport (email, phone) exists or not, used in multi step authentication`,
	}
}
func GetCheckClassicPassportActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name:        prefix + "security-token",
			Type:        "string",
			Description: "This can be the value of ReCaptcha2, ReCaptcha3, or generate security image or voice for verification. Will be used based on the configuration.",
		},
	}
}
func CastCheckClassicPassportActionReqFromCli(c emigo.CliCastable) CheckClassicPassportActionReq {
	data := CheckClassicPassportActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("security-token") {
		data.SecurityToken = c.String("security-token")
	}
	return data
}

// The base class definition for checkClassicPassportActionReq
type CheckClassicPassportActionReq struct {
	Value string `json:"value" validate:"required" yaml:"value"`
	// This can be the value of ReCaptcha2, ReCaptcha3, or generate security image or voice for verification. Will be used based on the configuration.
	SecurityToken string `json:"securityToken" yaml:"securityToken"`
}

func (x *CheckClassicPassportActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCheckClassicPassportActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "next",
			Type:        "slice",
			Description: "The next possible action which is suggested.",
		},
		{
			Name:        prefix + "flags",
			Type:        "slice",
			Description: "Extra information that can be useful actually when doing onboarding. Make sure sensitive information doesn't go out.",
		},
		{
			Name:        prefix + "otp-info",
			Type:        "object?",
			Description: "If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.",
		},
	}
}
func CastCheckClassicPassportActionResFromCli(c emigo.CliCastable) CheckClassicPassportActionRes {
	data := CheckClassicPassportActionRes{}
	if c.IsSet("next") {
		emigo.InflatePossibleSlice(c.String("next"), &data.Next)
	}
	if c.IsSet("flags") {
		emigo.InflatePossibleSlice(c.String("flags"), &data.Flags)
	}
	if c.IsSet("otp-info") {
		emigo.ParseNullable(c.String("otp-info"), &data.OtpInfo)
	}
	return data
}

// The base class definition for checkClassicPassportActionRes
type CheckClassicPassportActionRes struct {
	// The next possible action which is suggested.
	Next []string `json:"next" yaml:"next"`
	// Extra information that can be useful actually when doing onboarding. Make sure sensitive information doesn't go out.
	Flags []string `json:"flags" yaml:"flags"`
	// If the endpoint automatically triggers a send otp, then it would be holding that information, Also the otp information can become available.
	OtpInfo emigo.Nullable[CheckClassicPassportActionResOtpInfo] `json:"otpInfo" yaml:"otpInfo"`
}

func GetCheckClassicPassportActionResOtpInfoCliFlags(prefix string) []emigo.CliFlag {
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
func CastCheckClassicPassportActionResOtpInfoFromCli(c emigo.CliCastable) CheckClassicPassportActionResOtpInfo {
	data := CheckClassicPassportActionResOtpInfo{}
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

// The base class definition for otpInfo
type CheckClassicPassportActionResOtpInfo struct {
	SuspendUntil int64 `json:"suspendUntil" yaml:"suspendUntil"`
	ValidUntil   int64 `json:"validUntil" yaml:"validUntil"`
	BlockedUntil int64 `json:"blockedUntil" yaml:"blockedUntil"`
	// The amount of time left to unblock for next request
	SecondsToUnblock int64 `json:"secondsToUnblock" yaml:"secondsToUnblock"`
}

func (x *CheckClassicPassportActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CheckClassicPassportActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CheckClassicPassportActionResponse) SetContentType(contentType string) *CheckClassicPassportActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CheckClassicPassportActionResponse) AsStream(r io.Reader, contentType string) *CheckClassicPassportActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CheckClassicPassportActionResponse) AsJSON(payload any) *CheckClassicPassportActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CheckClassicPassportActionResponse) WithIdeal(payload CheckClassicPassportActionRes) *CheckClassicPassportActionResponse {
	x.Payload = payload
	return x
}
func (x *CheckClassicPassportActionResponse) AsHTML(payload string) *CheckClassicPassportActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CheckClassicPassportActionResponse) AsBytes(payload []byte) *CheckClassicPassportActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CheckClassicPassportActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CheckClassicPassportActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CheckClassicPassportActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type CheckClassicPassportActionRequestSig = func(c CheckClassicPassportActionRequest) (*CheckClassicPassportActionResponse, error)

/**
 * Query parameters for CheckClassicPassportAction
 */
// Query wrapper with private fields
type CheckClassicPassportActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CheckClassicPassportActionQueryFromString(rawQuery string) CheckClassicPassportActionQuery {
	v := CheckClassicPassportActionQuery{}
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
func CheckClassicPassportActionQueryFromHttp(r *http.Request) CheckClassicPassportActionQuery {
	return CheckClassicPassportActionQueryFromString(r.URL.RawQuery)
}
func (q CheckClassicPassportActionQuery) Values() url.Values {
	return q.values
}
func (q CheckClassicPassportActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CheckClassicPassportActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CheckClassicPassportActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CheckClassicPassportActionRequest struct {
	Body        CheckClassicPassportActionReq
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

func CheckClassicPassportActionClientCreateUrl(
	req CheckClassicPassportActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CheckClassicPassportActionMeta()
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
func CheckClassicPassportActionClientExecuteTyped(httpReq *http.Request) (*CheckClassicPassportActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CheckClassicPassportActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CheckClassicPassportActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &CheckClassicPassportActionResponse{Payload: result}, err
	}
	return &CheckClassicPassportActionResponse{Payload: result}, nil
}
func CheckClassicPassportActionClientBuildRequest(req CheckClassicPassportActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CheckClassicPassportActionMeta()
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
func CheckClassicPassportActionCall(
	req CheckClassicPassportActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CheckClassicPassportActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CheckClassicPassportActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CheckClassicPassportActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CheckClassicPassportActionClientExecuteTyped(r)
}

// CheckClassicPassportActionRaw registers a raw Gin route for the CheckClassicPassportAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CheckClassicPassportActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CheckClassicPassportActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// CheckClassicPassportActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CheckClassicPassportAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CheckClassicPassportActionHandler(
	handler func(c CheckClassicPassportActionRequest) (*CheckClassicPassportActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := CheckClassicPassportActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body CheckClassicPassportActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := CheckClassicPassportActionRequest{
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

// CheckClassicPassportActionGin is a high-level convenience wrapper around CheckClassicPassportActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CheckClassicPassportActionGin(r gin.IRoutes, handler func(c CheckClassicPassportActionRequest) (*CheckClassicPassportActionResponse, error)) {
	method, url, h := CheckClassicPassportActionHandler(handler)
	r.Handle(method, url, h)
}
func (x CheckClassicPassportActionRequest) IsGin() bool {
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
func CheckClassicPassportActionQueryFromGin(c *gin.Context) CheckClassicPassportActionQuery {
	return CheckClassicPassportActionQueryFromString(c.Request.URL.RawQuery)
}

// CheckClassicPassportActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the CheckClassicPassportAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *CheckClassicPassportActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func CheckClassicPassportActionHttpHandler(
	handler func(c CheckClassicPassportActionRequest) (*CheckClassicPassportActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := CheckClassicPassportActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body CheckClassicPassportActionReq
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
		req := CheckClassicPassportActionRequest{
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

// CheckClassicPassportActionHttp is a high-level convenience wrapper around
// CheckClassicPassportActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func CheckClassicPassportActionHttp(
	mux *http.ServeMux,
	handler func(c CheckClassicPassportActionRequest) (*CheckClassicPassportActionResponse, error),
) {
	method, pattern, h := CheckClassicPassportActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

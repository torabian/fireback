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
* Action to communicate with the action CheckPassportMethodsAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CheckPassportMethodsAction
func CheckPassportMethodsAction(c CheckPassportMethodsActionRequest) (*CheckPassportMethodsActionResponse, error) {
	return &CheckPassportMethodsActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CheckPassportMethodsActionMeta() struct {
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
		Name:        "CheckPassportMethodsAction",
		CliName:     "check-passport-methods",
		URL:         "/passports/available-methods",
		Method:      "GET",
		Description: `Publicly available information to create the authentication form, and show users how they can signin or signup to the system. Based on the PassportMethod entities, it will compute the available methods for the user, considering their region (IP for example)`,
	}
}
func GetCheckPassportMethodsActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "email",
			Type: "bool",
		},
		{
			Name: prefix + "phone",
			Type: "bool",
		},
		{
			Name: prefix + "google",
			Type: "bool",
		},
		{
			Name: prefix + "facebook",
			Type: "bool",
		},
		{
			Name: prefix + "google-o-auth-client-key",
			Type: "string",
		},
		{
			Name: prefix + "facebook-app-id",
			Type: "string",
		},
		{
			Name: prefix + "enabled-recaptcha2",
			Type: "bool",
		},
		{
			Name: prefix + "recaptcha2-client-key",
			Type: "string",
		},
	}
}
func CastCheckPassportMethodsActionResFromCli(c emigo.CliCastable) CheckPassportMethodsActionRes {
	data := CheckPassportMethodsActionRes{}
	if c.IsSet("email") {
		data.Email = bool(c.Bool("email"))
	}
	if c.IsSet("phone") {
		data.Phone = bool(c.Bool("phone"))
	}
	if c.IsSet("google") {
		data.Google = bool(c.Bool("google"))
	}
	if c.IsSet("facebook") {
		data.Facebook = bool(c.Bool("facebook"))
	}
	if c.IsSet("google-o-auth-client-key") {
		data.GoogleOAuthClientKey = c.String("google-o-auth-client-key")
	}
	if c.IsSet("facebook-app-id") {
		data.FacebookAppId = c.String("facebook-app-id")
	}
	if c.IsSet("enabled-recaptcha2") {
		data.EnabledRecaptcha2 = bool(c.Bool("enabled-recaptcha2"))
	}
	if c.IsSet("recaptcha2-client-key") {
		data.Recaptcha2ClientKey = c.String("recaptcha2-client-key")
	}
	return data
}

// The base class definition for checkPassportMethodsActionRes
type CheckPassportMethodsActionRes struct {
	Email                bool   `json:"email" yaml:"email"`
	Phone                bool   `json:"phone" yaml:"phone"`
	Google               bool   `json:"google" yaml:"google"`
	Facebook             bool   `json:"facebook" yaml:"facebook"`
	GoogleOAuthClientKey string `json:"googleOAuthClientKey" yaml:"googleOAuthClientKey"`
	FacebookAppId        string `json:"facebookAppId" yaml:"facebookAppId"`
	EnabledRecaptcha2    bool   `json:"enabledRecaptcha2" yaml:"enabledRecaptcha2"`
	Recaptcha2ClientKey  string `json:"recaptcha2ClientKey" yaml:"recaptcha2ClientKey"`
}

func (x *CheckPassportMethodsActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CheckPassportMethodsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CheckPassportMethodsActionResponse) SetContentType(contentType string) *CheckPassportMethodsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CheckPassportMethodsActionResponse) AsStream(r io.Reader, contentType string) *CheckPassportMethodsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CheckPassportMethodsActionResponse) AsJSON(payload any) *CheckPassportMethodsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CheckPassportMethodsActionResponse) WithIdeal(payload CheckPassportMethodsActionRes) *CheckPassportMethodsActionResponse {
	x.Payload = payload
	return x
}
func (x *CheckPassportMethodsActionResponse) AsHTML(payload string) *CheckPassportMethodsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CheckPassportMethodsActionResponse) AsBytes(payload []byte) *CheckPassportMethodsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CheckPassportMethodsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CheckPassportMethodsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CheckPassportMethodsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// CheckPassportMethodsActionRaw registers a raw Gin route for the CheckPassportMethodsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CheckPassportMethodsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CheckPassportMethodsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type CheckPassportMethodsActionRequestSig = func(c CheckPassportMethodsActionRequest) (*CheckPassportMethodsActionResponse, error)

// CheckPassportMethodsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CheckPassportMethodsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CheckPassportMethodsActionHandler(
	handler CheckPassportMethodsActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := CheckPassportMethodsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := CheckPassportMethodsActionRequest{
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

// CheckPassportMethodsAction is a high-level convenience wrapper around CheckPassportMethodsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CheckPassportMethodsActionGin(r gin.IRoutes, handler CheckPassportMethodsActionRequestSig) {
	method, url, h := CheckPassportMethodsActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for CheckPassportMethodsAction
 */
// Query wrapper with private fields
type CheckPassportMethodsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CheckPassportMethodsActionQueryFromString(rawQuery string) CheckPassportMethodsActionQuery {
	v := CheckPassportMethodsActionQuery{}
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
func CheckPassportMethodsActionQueryFromGin(c *gin.Context) CheckPassportMethodsActionQuery {
	return CheckPassportMethodsActionQueryFromString(c.Request.URL.RawQuery)
}
func CheckPassportMethodsActionQueryFromHttp(r *http.Request) CheckPassportMethodsActionQuery {
	return CheckPassportMethodsActionQueryFromString(r.URL.RawQuery)
}
func (q CheckPassportMethodsActionQuery) Values() url.Values {
	return q.values
}
func (q CheckPassportMethodsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CheckPassportMethodsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CheckPassportMethodsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CheckPassportMethodsActionRequest struct {
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

func (x CheckPassportMethodsActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x CheckPassportMethodsActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type CheckPassportMethodsActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func CheckPassportMethodsActionClientCreateUrl(
	req CheckPassportMethodsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CheckPassportMethodsActionMeta()
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
func CheckPassportMethodsActionClientExecuteTyped(httpReq *http.Request) (*CheckPassportMethodsActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CheckPassportMethodsActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CheckPassportMethodsActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &CheckPassportMethodsActionResponse{Payload: result}, err
	}
	return &CheckPassportMethodsActionResponse{Payload: result}, nil
}
func CheckPassportMethodsActionClientBuildRequest(req CheckPassportMethodsActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CheckPassportMethodsActionMeta()
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
func CheckPassportMethodsActionCall(
	req CheckPassportMethodsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CheckPassportMethodsActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CheckPassportMethodsActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CheckPassportMethodsActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CheckPassportMethodsActionClientExecuteTyped(r)
}

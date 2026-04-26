package abac

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ClassicSigninAction
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
func GetClassicSigninActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
		{
			Name: prefix + "totp-code",
			Type: "string",
		},
		{
			Name: prefix + "session-secret",
			Type: "string",
		},
	}
}
func CastClassicSigninActionReqFromCli(c emigo.CliCastable) ClassicSigninActionReq {
	data := ClassicSigninActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	if c.IsSet("totp-code") {
		data.TotpCode = c.String("totp-code")
	}
	if c.IsSet("session-secret") {
		data.SessionSecret = c.String("session-secret")
	}
	return data
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
func GetClassicSigninActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "session",
			Type: "one",
		},
		{
			Name: prefix + "next",
			Type: "slice",
		},
		{
			Name: prefix + "totp-url",
			Type: "string",
		},
		{
			Name: prefix + "session-secret",
			Type: "string",
		},
	}
}
func CastClassicSigninActionResFromCli(c emigo.CliCastable) ClassicSigninActionRes {
	data := ClassicSigninActionRes{}
	if c.IsSet("next") {
		emigo.InflatePossibleSlice(c.String("next"), &data.Next)
	}
	if c.IsSet("totp-url") {
		data.TotpUrl = c.String("totp-url")
	}
	if c.IsSet("session-secret") {
		data.SessionSecret = c.String("session-secret")
	}
	return data
}

// The base class definition for classicSigninActionRes
type ClassicSigninActionRes struct {
	Session UserSessionDto `json:"session" yaml:"session"`
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

// ClassicSigninActionRaw registers a raw Gin route for the ClassicSigninAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ClassicSigninActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ClassicSigninActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ClassicSigninActionRequestSig = func(c ClassicSigninActionRequest) (*ClassicSigninActionResponse, error)

// ClassicSigninActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ClassicSigninAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ClassicSigninActionHandler(
	handler ClassicSigninActionRequestSig,
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

// ClassicSigninAction is a high-level convenience wrapper around ClassicSigninActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ClassicSigninActionGin(r gin.IRoutes, handler ClassicSigninActionRequestSig) {
	method, url, h := ClassicSigninActionHandler(handler)
	r.Handle(method, url, h)
}

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
func ClassicSigninActionQueryFromGin(c *gin.Context) ClassicSigninActionQuery {
	return ClassicSigninActionQueryFromString(c.Request.URL.RawQuery)
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
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type ClassicSigninActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ClassicSigninActionCall(
	req ClassicSigninActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ClassicSigninActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ClassicSigninActionMeta()
		baseURL := meta.URL
		// Build final URL with query string
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		// if UrlValues present, encode and append
		if len(req.QueryParams) > 0 {
			u.RawQuery = req.QueryParams.Encode()
		}
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		req0, err := http.NewRequest(meta.Method, u.String(), bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}
		httpReq = req0
	} else {
		httpReq = config.Httpr
	}
	httpReq.Header = req.Headers
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var result ClassicSigninActionResult
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if resp.StatusCode >= 400 {
		return &result, fmt.Errorf("request failed: %s", respBody)
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}

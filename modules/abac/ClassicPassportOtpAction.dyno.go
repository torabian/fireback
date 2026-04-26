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
* Action to communicate with the action ClassicPassportOtpAction
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
func GetClassicPassportOtpActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "otp",
			Type: "string",
		},
	}
}
func CastClassicPassportOtpActionReqFromCli(c emigo.CliCastable) ClassicPassportOtpActionReq {
	data := ClassicPassportOtpActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("otp") {
		data.Otp = c.String("otp")
	}
	return data
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
func GetClassicPassportOtpActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "session",
			Type: "one?",
		},
		{
			Name: prefix + "totp-url",
			Type: "string",
		},
		{
			Name: prefix + "session-secret",
			Type: "string",
		},
		{
			Name: prefix + "continue-with-creation",
			Type: "bool",
		},
	}
}
func CastClassicPassportOtpActionResFromCli(c emigo.CliCastable) ClassicPassportOtpActionRes {
	data := ClassicPassportOtpActionRes{}
	if c.IsSet("session") {
		emigo.ParseNullable(c.String("session"), &data.Session)
	}
	if c.IsSet("totp-url") {
		data.TotpUrl = c.String("totp-url")
	}
	if c.IsSet("session-secret") {
		data.SessionSecret = c.String("session-secret")
	}
	if c.IsSet("continue-with-creation") {
		data.ContinueWithCreation = bool(c.Bool("continue-with-creation"))
	}
	return data
}

// The base class definition for classicPassportOtpActionRes
type ClassicPassportOtpActionRes struct {
	// Upon successful authentication, there will be a session dto generated, which is a ground information of authorized user and can be stored in front-end.
	Session emigo.Nullable[UserSessionDto] `json:"session" yaml:"session"`
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

// ClassicPassportOtpActionRaw registers a raw Gin route for the ClassicPassportOtpAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ClassicPassportOtpActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ClassicPassportOtpActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ClassicPassportOtpActionRequestSig = func(c ClassicPassportOtpActionRequest) (*ClassicPassportOtpActionResponse, error)

// ClassicPassportOtpActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ClassicPassportOtpAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ClassicPassportOtpActionHandler(
	handler ClassicPassportOtpActionRequestSig,
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

// ClassicPassportOtpAction is a high-level convenience wrapper around ClassicPassportOtpActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ClassicPassportOtpActionGin(r gin.IRoutes, handler ClassicPassportOtpActionRequestSig) {
	method, url, h := ClassicPassportOtpActionHandler(handler)
	r.Handle(method, url, h)
}

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
func ClassicPassportOtpActionQueryFromGin(c *gin.Context) ClassicPassportOtpActionQuery {
	return ClassicPassportOtpActionQueryFromString(c.Request.URL.RawQuery)
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
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type ClassicPassportOtpActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ClassicPassportOtpActionCall(
	req ClassicPassportOtpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ClassicPassportOtpActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ClassicPassportOtpActionMeta()
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
	var result ClassicPassportOtpActionResult
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

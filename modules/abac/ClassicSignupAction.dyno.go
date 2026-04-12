package abac

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli"
)

/**
* Action to communicate with the action ClassicSignupAction
 */
func ClassicSignupActionMeta() struct {
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
		Name:        "ClassicSignupAction",
		CliName:     "up",
		URL:         "/passports/signup/classic",
		Method:      "POST",
		Description: `Signup a user into system via public access (aka website visitors) using either email or phone number.`,
	}
}
func GetClassicSignupActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "session-secret",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "enum",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
		{
			Name: prefix + "first-name",
			Type: "string",
		},
		{
			Name: prefix + "last-name",
			Type: "string",
		},
		{
			Name: prefix + "invite-id",
			Type: "string?",
		},
		{
			Name: prefix + "public-join-key-id",
			Type: "string?",
		},
		{
			Name: prefix + "workspace-type-id",
			Type: "string?",
		},
	}
}
func CastClassicSignupActionReqFromCli(c emigo.CliCastable) ClassicSignupActionReq {
	data := ClassicSignupActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("session-secret") {
		data.SessionSecret = c.String("session-secret")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("invite-id") {
		emigo.ParseNullable(c.String("invite-id"), &data.InviteId)
	}
	if c.IsSet("public-join-key-id") {
		emigo.ParseNullable(c.String("public-join-key-id"), &data.PublicJoinKeyId)
	}
	if c.IsSet("workspace-type-id") {
		emigo.ParseNullable(c.String("workspace-type-id"), &data.WorkspaceTypeId)
	}
	return data
}

// The base class definition for classicSignupActionReq
type ClassicSignupActionReq struct {
	Value string `json:"value" validate:"required" yaml:"value"`
	// Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.
	SessionSecret   string                 `json:"sessionSecret" yaml:"sessionSecret"`
	Type            string                 `json:"type" validate:"required" yaml:"type"`
	Password        string                 `json:"password" validate:"required" yaml:"password"`
	FirstName       string                 `json:"firstName" validate:"required" yaml:"firstName"`
	LastName        string                 `json:"lastName" validate:"required" yaml:"lastName"`
	InviteId        emigo.Nullable[string] `json:"inviteId" yaml:"inviteId"`
	PublicJoinKeyId emigo.Nullable[string] `json:"publicJoinKeyId" yaml:"publicJoinKeyId"`
	WorkspaceTypeId emigo.Nullable[string] `json:"workspaceTypeId" validate:"required" yaml:"workspaceTypeId"`
}

func (x *ClassicSignupActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetClassicSignupActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "session",
			Type: "one",
		},
		{
			Name: prefix + "totp-url",
			Type: "string",
		},
		{
			Name: prefix + "continue-to-totp",
			Type: "bool",
		},
		{
			Name: prefix + "forced-totp",
			Type: "bool",
		},
	}
}
func CastClassicSignupActionResFromCli(c emigo.CliCastable) ClassicSignupActionRes {
	data := ClassicSignupActionRes{}
	if c.IsSet("totp-url") {
		data.TotpUrl = c.String("totp-url")
	}
	if c.IsSet("continue-to-totp") {
		data.ContinueToTotp = bool(c.Bool("continue-to-totp"))
	}
	if c.IsSet("forced-totp") {
		data.ForcedTotp = bool(c.Bool("forced-totp"))
	}
	return data
}

// The base class definition for classicSignupActionRes
type ClassicSignupActionRes struct {
	// Returns the user session in case that signup is completely successful.
	Session UserSessionDto `json:"session" yaml:"session"`
	// If time based otp is available, we add it response to make it easier for ui.
	TotpUrl string `json:"totpUrl" yaml:"totpUrl"`
	// Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.
	ContinueToTotp bool `json:"continueToTotp" yaml:"continueToTotp"`
	// Determines if user must complete totp in order to continue based on workspace or installation
	ForcedTotp bool `json:"forcedTotp" yaml:"forcedTotp"`
}

func (x *ClassicSignupActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ClassicSignupActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *ClassicSignupActionResponse) SetContentType(contentType string) *ClassicSignupActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ClassicSignupActionResponse) AsStream(r io.Reader, contentType string) *ClassicSignupActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ClassicSignupActionResponse) AsJSON(payload any) *ClassicSignupActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *ClassicSignupActionResponse) AsHTML(payload string) *ClassicSignupActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ClassicSignupActionResponse) AsBytes(payload []byte) *ClassicSignupActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ClassicSignupActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ClassicSignupActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ClassicSignupActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ClassicSignupActionRaw registers a raw Gin route for the ClassicSignupAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ClassicSignupActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ClassicSignupActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ClassicSignupActionRequestSig = func(c ClassicSignupActionRequest) (*ClassicSignupActionResponse, error)

// ClassicSignupActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ClassicSignupAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ClassicSignupActionHandler(
	handler ClassicSignupActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ClassicSignupActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ClassicSignupActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ClassicSignupActionRequest{
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

// ClassicSignupAction is a high-level convenience wrapper around ClassicSignupActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ClassicSignupActionGin(r gin.IRoutes, handler ClassicSignupActionRequestSig) {
	method, url, h := ClassicSignupActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for ClassicSignupAction
 */
// Query wrapper with private fields
type ClassicSignupActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ClassicSignupActionQueryFromString(rawQuery string) ClassicSignupActionQuery {
	v := ClassicSignupActionQuery{}
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
func ClassicSignupActionQueryFromGin(c *gin.Context) ClassicSignupActionQuery {
	return ClassicSignupActionQueryFromString(c.Request.URL.RawQuery)
}
func ClassicSignupActionQueryFromHttp(r *http.Request) ClassicSignupActionQuery {
	return ClassicSignupActionQueryFromString(r.URL.RawQuery)
}
func (q ClassicSignupActionQuery) Values() url.Values {
	return q.values
}
func (q ClassicSignupActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ClassicSignupActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ClassicSignupActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ClassicSignupActionRequest struct {
	Body        ClassicSignupActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type ClassicSignupActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ClassicSignupActionCall(
	req ClassicSignupActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ClassicSignupActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ClassicSignupActionMeta()
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
	var result ClassicSignupActionResult
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

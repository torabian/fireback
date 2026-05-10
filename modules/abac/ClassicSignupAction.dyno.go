package abac

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action ClassicSignupAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of ClassicSignupAction
func ClassicSignupAction(c ClassicSignupActionRequest) (*ClassicSignupActionResponse, error) {
	return &ClassicSignupActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
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
			Name:        prefix + "session-secret",
			Type:        "string",
			Description: "Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.",
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
			Name:        prefix + "session",
			Type:        "one",
			Description: "Returns the user session in case that signup is completely successful.",
		},
		{
			Name:        prefix + "totp-url",
			Type:        "string",
			Description: "If time based otp is available, we add it response to make it easier for ui.",
		},
		{
			Name:        prefix + "continue-to-totp",
			Type:        "bool",
			Description: "Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.",
		},
		{
			Name:        prefix + "forced-totp",
			Type:        "bool",
			Description: "Determines if user must complete totp in order to continue based on workspace or installation",
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
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
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

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ClassicSignupActionResponse) WithIdeal(payload ClassicSignupActionRes) *ClassicSignupActionResponse {
	x.Payload = payload
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

func (x ClassicSignupActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x ClassicSignupActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type ClassicSignupActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func ClassicSignupActionClientCreateUrl(
	req ClassicSignupActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := ClassicSignupActionMeta()
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
func ClassicSignupActionClientExecuteTyped(httpReq *http.Request) (*ClassicSignupActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result ClassicSignupActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ClassicSignupActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &ClassicSignupActionResponse{Payload: result}, err
	}
	return &ClassicSignupActionResponse{Payload: result}, nil
}
func ClassicSignupActionClientBuildRequest(req ClassicSignupActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := ClassicSignupActionMeta()
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
func ClassicSignupActionCall(
	req ClassicSignupActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ClassicSignupActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := ClassicSignupActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := ClassicSignupActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return ClassicSignupActionClientExecuteTyped(r)
}

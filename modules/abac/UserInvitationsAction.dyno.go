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
* Action to communicate with the action UserInvitationsAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of UserInvitationsAction
func UserInvitationsAction(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error) {
	return &UserInvitationsActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func UserInvitationsActionMeta() struct {
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
		Name:        "UserInvitationsAction",
		CliName:     "user-invitations-action",
		URL:         "/users/invitations",
		Method:      "GET",
		Description: `Shows the invitations for an specific user, if the invited member already has a account. It's based on the passports, so if the passport is authenticated we will show them.`,
	}
}
func GetUserInvitationsActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "user-id",
			Type:        "string",
			Description: "UserUniqueId",
		},
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "Invitation unique id",
		},
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "The value of the passport (email/phone)",
		},
		{
			Name:        prefix + "role-name",
			Type:        "string",
			Description: "Name of the role that user will get",
		},
		{
			Name:        prefix + "workspace-name",
			Type:        "string",
			Description: "Name of the workspace which user is invited to.",
		},
		{
			Name:        prefix + "type",
			Type:        "string",
			Description: "The method of the invitation, such as email.",
		},
		{
			Name:        prefix + "cover-letter",
			Type:        "string",
			Description: "The content that user will receive to understand the reason of the letter.",
		},
	}
}
func CastUserInvitationsActionResFromCli(c emigo.CliCastable) UserInvitationsActionRes {
	data := UserInvitationsActionRes{}
	if c.IsSet("user-id") {
		data.UserId = c.String("user-id")
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("role-name") {
		data.RoleName = c.String("role-name")
	}
	if c.IsSet("workspace-name") {
		data.WorkspaceName = c.String("workspace-name")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("cover-letter") {
		data.CoverLetter = c.String("cover-letter")
	}
	return data
}

// The base class definition for userInvitationsActionRes
type UserInvitationsActionRes struct {
	// UserUniqueId
	UserId string `json:"userId" yaml:"userId"`
	// Invitation unique id
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// The value of the passport (email/phone)
	Value string `json:"value" yaml:"value"`
	// Name of the role that user will get
	RoleName string `json:"roleName" yaml:"roleName"`
	// Name of the workspace which user is invited to.
	WorkspaceName string `json:"workspaceName" yaml:"workspaceName"`
	// The method of the invitation, such as email.
	Type string `json:"type" yaml:"type"`
	// The content that user will receive to understand the reason of the letter.
	CoverLetter string `json:"coverLetter" yaml:"coverLetter"`
}

func (x *UserInvitationsActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type UserInvitationsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *UserInvitationsActionResponse) SetContentType(contentType string) *UserInvitationsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *UserInvitationsActionResponse) AsStream(r io.Reader, contentType string) *UserInvitationsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *UserInvitationsActionResponse) AsJSON(payload any) *UserInvitationsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *UserInvitationsActionResponse) WithIdeal(payload UserInvitationsActionRes) *UserInvitationsActionResponse {
	x.Payload = payload
	return x
}
func (x *UserInvitationsActionResponse) AsHTML(payload string) *UserInvitationsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *UserInvitationsActionResponse) AsBytes(payload []byte) *UserInvitationsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x UserInvitationsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x UserInvitationsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x UserInvitationsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// UserInvitationsActionRaw registers a raw Gin route for the UserInvitationsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func UserInvitationsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := UserInvitationsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type UserInvitationsActionRequestSig = func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error)

// UserInvitationsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the UserInvitationsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func UserInvitationsActionHandler(
	handler UserInvitationsActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := UserInvitationsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := UserInvitationsActionRequest{
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

// UserInvitationsAction is a high-level convenience wrapper around UserInvitationsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func UserInvitationsActionGin(r gin.IRoutes, handler UserInvitationsActionRequestSig) {
	method, url, h := UserInvitationsActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for UserInvitationsAction
 */
// Query wrapper with private fields
type UserInvitationsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func UserInvitationsActionQueryFromString(rawQuery string) UserInvitationsActionQuery {
	v := UserInvitationsActionQuery{}
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
func UserInvitationsActionQueryFromGin(c *gin.Context) UserInvitationsActionQuery {
	return UserInvitationsActionQueryFromString(c.Request.URL.RawQuery)
}
func UserInvitationsActionQueryFromHttp(r *http.Request) UserInvitationsActionQuery {
	return UserInvitationsActionQueryFromString(r.URL.RawQuery)
}
func (q UserInvitationsActionQuery) Values() url.Values {
	return q.values
}
func (q UserInvitationsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *UserInvitationsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *UserInvitationsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type UserInvitationsActionRequest struct {
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

func (x UserInvitationsActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x UserInvitationsActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type UserInvitationsActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func UserInvitationsActionClientCreateUrl(
	req UserInvitationsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := UserInvitationsActionMeta()
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
func UserInvitationsActionClientExecuteTyped(httpReq *http.Request) (*UserInvitationsActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result UserInvitationsActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &UserInvitationsActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &UserInvitationsActionResponse{Payload: result}, err
	}
	return &UserInvitationsActionResponse{Payload: result}, nil
}
func UserInvitationsActionClientBuildRequest(req UserInvitationsActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := UserInvitationsActionMeta()
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
func UserInvitationsActionCall(
	req UserInvitationsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UserInvitationsActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := UserInvitationsActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := UserInvitationsActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return UserInvitationsActionClientExecuteTyped(r)
}

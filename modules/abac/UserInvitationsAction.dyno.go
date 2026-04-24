package abac

import (
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
* Action to communicate with the action UserInvitationsAction
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
			Name: prefix + "user-id",
			Type: "string",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "role-name",
			Type: "string",
		},
		{
			Name: prefix + "workspace-name",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "cover-letter",
			Type: "string",
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
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type UserInvitationsActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func UserInvitationsActionCall(
	req UserInvitationsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UserInvitationsActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := UserInvitationsActionMeta()
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
		req0, err := http.NewRequest(meta.Method, u.String(), nil)
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
	var result UserInvitationsActionResult
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

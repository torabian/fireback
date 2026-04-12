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
* Action to communicate with the action InviteToWorkspaceAction
 */
func InviteToWorkspaceActionMeta() struct {
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
		Name:        "InviteToWorkspaceAction",
		CliName:     "invite",
		URL:         "/invite/to/workspace",
		Method:      "POST",
		Description: `Invite a new person (either a user, with passport or without passport)`,
	}
}
func GetInviteToWorkspaceActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "public-key",
			Type: "string",
		},
		{
			Name: prefix + "cover-letter",
			Type: "string",
		},
		{
			Name: prefix + "target-user-locale",
			Type: "string",
		},
		{
			Name: prefix + "email",
			Type: "string",
		},
		{
			Name: prefix + "phonenumber",
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
			Name: prefix + "force-email-address",
			Type: "bool?",
		},
		{
			Name: prefix + "force-phone-number",
			Type: "bool?",
		},
		{
			Name: prefix + "role-id",
			Type: "string",
		},
	}
}
func CastInviteToWorkspaceActionReqFromCli(c emigo.CliCastable) InviteToWorkspaceActionReq {
	data := InviteToWorkspaceActionReq{}
	if c.IsSet("public-key") {
		data.PublicKey = c.String("public-key")
	}
	if c.IsSet("cover-letter") {
		data.CoverLetter = c.String("cover-letter")
	}
	if c.IsSet("target-user-locale") {
		data.TargetUserLocale = c.String("target-user-locale")
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("phonenumber") {
		data.Phonenumber = c.String("phonenumber")
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("force-email-address") {
		emigo.ParseNullable(c.String("force-email-address"), &data.ForceEmailAddress)
	}
	if c.IsSet("force-phone-number") {
		emigo.ParseNullable(c.String("force-phone-number"), &data.ForcePhoneNumber)
	}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
	}
	return data
}

// The base class definition for inviteToWorkspaceActionReq
type InviteToWorkspaceActionReq struct {
	// A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.
	PublicKey string `json:"publicKey" yaml:"publicKey"`
	// The content that user will receive to understand the reason of the letter.
	CoverLetter string `json:"coverLetter" yaml:"coverLetter"`
	// If the invited person has a different language, then you can define that so the interface for him will be automatically translated.
	TargetUserLocale string `json:"targetUserLocale" yaml:"targetUserLocale"`
	// The email address of the person which is invited.
	Email string `json:"email" yaml:"email"`
	// The phone number of the person which is invited.
	Phonenumber string `json:"phonenumber" yaml:"phonenumber"`
	// First name of the person which is invited
	FirstName string `json:"firstName" validate:"required" yaml:"firstName"`
	// Last name of the person which is invited.
	LastName string `json:"lastName" validate:"required" yaml:"lastName"`
	// If forced, the email address cannot be changed by the user which has been invited.
	ForceEmailAddress emigo.Nullable[bool] `json:"forceEmailAddress" yaml:"forceEmailAddress"`
	// If forced, user cannot change the phone number and needs to complete signup.
	ForcePhoneNumber emigo.Nullable[bool] `json:"forcePhoneNumber" yaml:"forcePhoneNumber"`
	// The role which invitee get if they accept the request.
	RoleId string `json:"roleId" validate:"required" yaml:"roleId"`
}

func (x *InviteToWorkspaceActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type InviteToWorkspaceActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *InviteToWorkspaceActionResponse) SetContentType(contentType string) *InviteToWorkspaceActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *InviteToWorkspaceActionResponse) AsStream(r io.Reader, contentType string) *InviteToWorkspaceActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *InviteToWorkspaceActionResponse) AsJSON(payload any) *InviteToWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *InviteToWorkspaceActionResponse) AsHTML(payload string) *InviteToWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *InviteToWorkspaceActionResponse) AsBytes(payload []byte) *InviteToWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x InviteToWorkspaceActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x InviteToWorkspaceActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x InviteToWorkspaceActionResponse) GetPayload() interface{} {
	return x.Payload
}

// InviteToWorkspaceActionRaw registers a raw Gin route for the InviteToWorkspaceAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func InviteToWorkspaceActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := InviteToWorkspaceActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type InviteToWorkspaceActionRequestSig = func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error)

// InviteToWorkspaceActionHandler returns the HTTP method, route URL, and a typed Gin handler for the InviteToWorkspaceAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func InviteToWorkspaceActionHandler(
	handler InviteToWorkspaceActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := InviteToWorkspaceActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body InviteToWorkspaceActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := InviteToWorkspaceActionRequest{
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

// InviteToWorkspaceAction is a high-level convenience wrapper around InviteToWorkspaceActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func InviteToWorkspaceActionGin(r gin.IRoutes, handler InviteToWorkspaceActionRequestSig) {
	method, url, h := InviteToWorkspaceActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for InviteToWorkspaceAction
 */
// Query wrapper with private fields
type InviteToWorkspaceActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func InviteToWorkspaceActionQueryFromString(rawQuery string) InviteToWorkspaceActionQuery {
	v := InviteToWorkspaceActionQuery{}
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
func InviteToWorkspaceActionQueryFromGin(c *gin.Context) InviteToWorkspaceActionQuery {
	return InviteToWorkspaceActionQueryFromString(c.Request.URL.RawQuery)
}
func InviteToWorkspaceActionQueryFromHttp(r *http.Request) InviteToWorkspaceActionQuery {
	return InviteToWorkspaceActionQueryFromString(r.URL.RawQuery)
}
func (q InviteToWorkspaceActionQuery) Values() url.Values {
	return q.values
}
func (q InviteToWorkspaceActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *InviteToWorkspaceActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *InviteToWorkspaceActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type InviteToWorkspaceActionRequest struct {
	Body        InviteToWorkspaceActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type InviteToWorkspaceActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func InviteToWorkspaceActionCall(
	req InviteToWorkspaceActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*InviteToWorkspaceActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := InviteToWorkspaceActionMeta()
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
	var result InviteToWorkspaceActionResult
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

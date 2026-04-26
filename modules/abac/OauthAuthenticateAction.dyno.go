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
* Action to communicate with the action OauthAuthenticateAction
 */
func OauthAuthenticateActionMeta() struct {
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
		Name:        "OauthAuthenticateAction",
		CliName:     "oauth-authenticate-action",
		URL:         "/passport/via-oauth",
		Method:      "POST",
		Description: `When a token is got from a oauth service such as google, we send the token here to authenticate the user. To me seems this doesn't need to have 2FA or anything, so we return the session directly, or maybe there needs to be next step.`,
	}
}
func GetOauthAuthenticateActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "token",
			Type: "string",
		},
		{
			Name: prefix + "service",
			Type: "string",
		},
	}
}
func CastOauthAuthenticateActionReqFromCli(c emigo.CliCastable) OauthAuthenticateActionReq {
	data := OauthAuthenticateActionReq{}
	if c.IsSet("token") {
		data.Token = c.String("token")
	}
	if c.IsSet("service") {
		data.Service = c.String("service")
	}
	return data
}

// The base class definition for oauthAuthenticateActionReq
type OauthAuthenticateActionReq struct {
	// The token that Auth2 provider returned to the front-end, which will be used to validate the backend
	Token string `json:"token" yaml:"token"`
	// The service name, such as 'google' which later backend will use to authorize the token and create the user.
	Service string `json:"service" yaml:"service"`
}

func (x *OauthAuthenticateActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetOauthAuthenticateActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "session",
			Type: "one",
		},
		{
			Name: prefix + "next",
			Type: "slice",
		},
	}
}
func CastOauthAuthenticateActionResFromCli(c emigo.CliCastable) OauthAuthenticateActionRes {
	data := OauthAuthenticateActionRes{}
	if c.IsSet("next") {
		emigo.InflatePossibleSlice(c.String("next"), &data.Next)
	}
	return data
}

// The base class definition for oauthAuthenticateActionRes
type OauthAuthenticateActionRes struct {
	Session UserSessionDto `json:"session" yaml:"session"`
	// The next possible action which is suggested.
	Next []string `json:"next" yaml:"next"`
}

func (x *OauthAuthenticateActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type OauthAuthenticateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *OauthAuthenticateActionResponse) SetContentType(contentType string) *OauthAuthenticateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *OauthAuthenticateActionResponse) AsStream(r io.Reader, contentType string) *OauthAuthenticateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *OauthAuthenticateActionResponse) AsJSON(payload any) *OauthAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *OauthAuthenticateActionResponse) WithIdeal(payload OauthAuthenticateActionRes) *OauthAuthenticateActionResponse {
	x.Payload = payload
	return x
}
func (x *OauthAuthenticateActionResponse) AsHTML(payload string) *OauthAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *OauthAuthenticateActionResponse) AsBytes(payload []byte) *OauthAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x OauthAuthenticateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x OauthAuthenticateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x OauthAuthenticateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// OauthAuthenticateActionRaw registers a raw Gin route for the OauthAuthenticateAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func OauthAuthenticateActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := OauthAuthenticateActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type OauthAuthenticateActionRequestSig = func(c OauthAuthenticateActionRequest) (*OauthAuthenticateActionResponse, error)

// OauthAuthenticateActionHandler returns the HTTP method, route URL, and a typed Gin handler for the OauthAuthenticateAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func OauthAuthenticateActionHandler(
	handler OauthAuthenticateActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := OauthAuthenticateActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body OauthAuthenticateActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := OauthAuthenticateActionRequest{
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

// OauthAuthenticateAction is a high-level convenience wrapper around OauthAuthenticateActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func OauthAuthenticateActionGin(r gin.IRoutes, handler OauthAuthenticateActionRequestSig) {
	method, url, h := OauthAuthenticateActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for OauthAuthenticateAction
 */
// Query wrapper with private fields
type OauthAuthenticateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func OauthAuthenticateActionQueryFromString(rawQuery string) OauthAuthenticateActionQuery {
	v := OauthAuthenticateActionQuery{}
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
func OauthAuthenticateActionQueryFromGin(c *gin.Context) OauthAuthenticateActionQuery {
	return OauthAuthenticateActionQueryFromString(c.Request.URL.RawQuery)
}
func OauthAuthenticateActionQueryFromHttp(r *http.Request) OauthAuthenticateActionQuery {
	return OauthAuthenticateActionQueryFromString(r.URL.RawQuery)
}
func (q OauthAuthenticateActionQuery) Values() url.Values {
	return q.values
}
func (q OauthAuthenticateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *OauthAuthenticateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *OauthAuthenticateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type OauthAuthenticateActionRequest struct {
	Body        OauthAuthenticateActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type OauthAuthenticateActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func OauthAuthenticateActionCall(
	req OauthAuthenticateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*OauthAuthenticateActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := OauthAuthenticateActionMeta()
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
	var result OauthAuthenticateActionResult
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

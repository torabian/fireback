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
* Action to communicate with the action UserPassportsAction
 */
func UserPassportsActionMeta() struct {
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
		Name:        "UserPassportsAction",
		CliName:     "user-passports-action",
		URL:         "/user/passports",
		Method:      "GET",
		Description: `Returns list of passports belongs to an specific user.`,
	}
}
func GetUserPassportsActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "totp-confirmed",
			Type: "bool",
		},
	}
}
func CastUserPassportsActionResFromCli(c emigo.CliCastable) UserPassportsActionRes {
	data := UserPassportsActionRes{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("totp-confirmed") {
		data.TotpConfirmed = bool(c.Bool("totp-confirmed"))
	}
	return data
}

// The base class definition for userPassportsActionRes
type UserPassportsActionRes struct {
	// The passport value, such as email address or phone number
	Value string `json:"value" yaml:"value"`
	// Unique identifier of the passport to operate some action on top of it
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// The type of the passport, such as email, phone number
	Type string `json:"type" yaml:"type"`
	// Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
	TotpConfirmed bool `json:"totpConfirmed" yaml:"totpConfirmed"`
}

func (x *UserPassportsActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type UserPassportsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *UserPassportsActionResponse) SetContentType(contentType string) *UserPassportsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *UserPassportsActionResponse) AsStream(r io.Reader, contentType string) *UserPassportsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *UserPassportsActionResponse) AsJSON(payload any) *UserPassportsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *UserPassportsActionResponse) WithIdeal(payload UserPassportsActionRes) *UserPassportsActionResponse {
	x.Payload = payload
	return x
}
func (x *UserPassportsActionResponse) AsHTML(payload string) *UserPassportsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *UserPassportsActionResponse) AsBytes(payload []byte) *UserPassportsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x UserPassportsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x UserPassportsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x UserPassportsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// UserPassportsActionRaw registers a raw Gin route for the UserPassportsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func UserPassportsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := UserPassportsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type UserPassportsActionRequestSig = func(c UserPassportsActionRequest) (*UserPassportsActionResponse, error)

// UserPassportsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the UserPassportsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func UserPassportsActionHandler(
	handler UserPassportsActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := UserPassportsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := UserPassportsActionRequest{
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

// UserPassportsAction is a high-level convenience wrapper around UserPassportsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func UserPassportsActionGin(r gin.IRoutes, handler UserPassportsActionRequestSig) {
	method, url, h := UserPassportsActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for UserPassportsAction
 */
// Query wrapper with private fields
type UserPassportsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func UserPassportsActionQueryFromString(rawQuery string) UserPassportsActionQuery {
	v := UserPassportsActionQuery{}
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
func UserPassportsActionQueryFromGin(c *gin.Context) UserPassportsActionQuery {
	return UserPassportsActionQueryFromString(c.Request.URL.RawQuery)
}
func UserPassportsActionQueryFromHttp(r *http.Request) UserPassportsActionQuery {
	return UserPassportsActionQueryFromString(r.URL.RawQuery)
}
func (q UserPassportsActionQuery) Values() url.Values {
	return q.values
}
func (q UserPassportsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *UserPassportsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *UserPassportsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type UserPassportsActionRequest struct {
	Body        interface{}
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type UserPassportsActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func UserPassportsActionCall(
	req UserPassportsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UserPassportsActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := UserPassportsActionMeta()
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
	var result UserPassportsActionResult
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

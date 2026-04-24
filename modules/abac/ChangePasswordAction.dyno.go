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
* Action to communicate with the action ChangePasswordAction
 */
func ChangePasswordActionMeta() struct {
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
		Name:        "ChangePasswordAction",
		CliName:     "cp",
		URL:         "/passport/change-password",
		Method:      "POST",
		Description: `Change the password for a given passport of the user. User needs to be authenticated in order to be able to change the password for a given account.`,
	}
}
func GetChangePasswordActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "password",
			Type: "string",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
	}
}
func CastChangePasswordActionReqFromCli(c emigo.CliCastable) ChangePasswordActionReq {
	data := ChangePasswordActionReq{}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	return data
}

// The base class definition for changePasswordActionReq
type ChangePasswordActionReq struct {
	// New password meeting the security requirements.
	Password string `json:"password" validate:"required" yaml:"password"`
	// The passport uniqueId (not the email or phone number) which password would be applied to. Don't confuse with value.
	UniqueId string `json:"uniqueId" validate:"required" yaml:"uniqueId"`
}

func (x *ChangePasswordActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetChangePasswordActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "changed",
			Type: "bool",
		},
	}
}
func CastChangePasswordActionResFromCli(c emigo.CliCastable) ChangePasswordActionRes {
	data := ChangePasswordActionRes{}
	if c.IsSet("changed") {
		data.Changed = bool(c.Bool("changed"))
	}
	return data
}

// The base class definition for changePasswordActionRes
type ChangePasswordActionRes struct {
	Changed bool `json:"changed" yaml:"changed"`
}

func (x *ChangePasswordActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ChangePasswordActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *ChangePasswordActionResponse) SetContentType(contentType string) *ChangePasswordActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ChangePasswordActionResponse) AsStream(r io.Reader, contentType string) *ChangePasswordActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ChangePasswordActionResponse) AsJSON(payload any) *ChangePasswordActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *ChangePasswordActionResponse) WithIdeal(payload ChangePasswordActionRes) *ChangePasswordActionResponse {
	x.Payload = payload
	return x
}
func (x *ChangePasswordActionResponse) AsHTML(payload string) *ChangePasswordActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ChangePasswordActionResponse) AsBytes(payload []byte) *ChangePasswordActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ChangePasswordActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ChangePasswordActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ChangePasswordActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ChangePasswordActionRaw registers a raw Gin route for the ChangePasswordAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ChangePasswordActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ChangePasswordActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ChangePasswordActionRequestSig = func(c ChangePasswordActionRequest) (*ChangePasswordActionResponse, error)

// ChangePasswordActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ChangePasswordAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ChangePasswordActionHandler(
	handler ChangePasswordActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ChangePasswordActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ChangePasswordActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ChangePasswordActionRequest{
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

// ChangePasswordAction is a high-level convenience wrapper around ChangePasswordActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ChangePasswordActionGin(r gin.IRoutes, handler ChangePasswordActionRequestSig) {
	method, url, h := ChangePasswordActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for ChangePasswordAction
 */
// Query wrapper with private fields
type ChangePasswordActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ChangePasswordActionQueryFromString(rawQuery string) ChangePasswordActionQuery {
	v := ChangePasswordActionQuery{}
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
func ChangePasswordActionQueryFromGin(c *gin.Context) ChangePasswordActionQuery {
	return ChangePasswordActionQueryFromString(c.Request.URL.RawQuery)
}
func ChangePasswordActionQueryFromHttp(r *http.Request) ChangePasswordActionQuery {
	return ChangePasswordActionQueryFromString(r.URL.RawQuery)
}
func (q ChangePasswordActionQuery) Values() url.Values {
	return q.values
}
func (q ChangePasswordActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ChangePasswordActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ChangePasswordActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ChangePasswordActionRequest struct {
	Body        ChangePasswordActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type ChangePasswordActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ChangePasswordActionCall(
	req ChangePasswordActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ChangePasswordActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ChangePasswordActionMeta()
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
	var result ChangePasswordActionResult
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

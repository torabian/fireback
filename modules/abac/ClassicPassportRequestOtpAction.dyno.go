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
* Action to communicate with the action ClassicPassportRequestOtpAction
 */
func ClassicPassportRequestOtpActionMeta() struct {
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
		Name:        "ClassicPassportRequestOtpAction",
		CliName:     "otp-request",
		URL:         "/workspace/passport/request-otp",
		Method:      "POST",
		Description: `Triggers an otp request, and will send an sms or email to the passport. This endpoint is not used for login, but rather makes a request at initial step. Later you can call classicPassportOtp to get in.`,
	}
}
func GetClassicPassportRequestOtpActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
	}
}
func CastClassicPassportRequestOtpActionReqFromCli(c emigo.CliCastable) ClassicPassportRequestOtpActionReq {
	data := ClassicPassportRequestOtpActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	return data
}

// The base class definition for classicPassportRequestOtpActionReq
type ClassicPassportRequestOtpActionReq struct {
	// Passport value (email, phone number) which would be receiving the otp code.
	Value string `json:"value" validate:"required" yaml:"value"`
}

func (x *ClassicPassportRequestOtpActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetClassicPassportRequestOtpActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "suspend-until",
			Type: "int64",
		},
		{
			Name: prefix + "valid-until",
			Type: "int64",
		},
		{
			Name: prefix + "blocked-until",
			Type: "int64",
		},
		{
			Name: prefix + "seconds-to-unblock",
			Type: "int64",
		},
	}
}
func CastClassicPassportRequestOtpActionResFromCli(c emigo.CliCastable) ClassicPassportRequestOtpActionRes {
	data := ClassicPassportRequestOtpActionRes{}
	if c.IsSet("suspend-until") {
		data.SuspendUntil = int64(c.Int64("suspend-until"))
	}
	if c.IsSet("valid-until") {
		data.ValidUntil = int64(c.Int64("valid-until"))
	}
	if c.IsSet("blocked-until") {
		data.BlockedUntil = int64(c.Int64("blocked-until"))
	}
	if c.IsSet("seconds-to-unblock") {
		data.SecondsToUnblock = int64(c.Int64("seconds-to-unblock"))
	}
	return data
}

// The base class definition for classicPassportRequestOtpActionRes
type ClassicPassportRequestOtpActionRes struct {
	SuspendUntil int64 `json:"suspendUntil" yaml:"suspendUntil"`
	ValidUntil   int64 `json:"validUntil" yaml:"validUntil"`
	BlockedUntil int64 `json:"blockedUntil" yaml:"blockedUntil"`
	// The amount of time left to unblock for next request
	SecondsToUnblock int64 `json:"secondsToUnblock" yaml:"secondsToUnblock"`
}

func (x *ClassicPassportRequestOtpActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type ClassicPassportRequestOtpActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *ClassicPassportRequestOtpActionResponse) SetContentType(contentType string) *ClassicPassportRequestOtpActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *ClassicPassportRequestOtpActionResponse) AsStream(r io.Reader, contentType string) *ClassicPassportRequestOtpActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *ClassicPassportRequestOtpActionResponse) AsJSON(payload any) *ClassicPassportRequestOtpActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *ClassicPassportRequestOtpActionResponse) AsHTML(payload string) *ClassicPassportRequestOtpActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *ClassicPassportRequestOtpActionResponse) AsBytes(payload []byte) *ClassicPassportRequestOtpActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x ClassicPassportRequestOtpActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x ClassicPassportRequestOtpActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x ClassicPassportRequestOtpActionResponse) GetPayload() interface{} {
	return x.Payload
}

// ClassicPassportRequestOtpActionRaw registers a raw Gin route for the ClassicPassportRequestOtpAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func ClassicPassportRequestOtpActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := ClassicPassportRequestOtpActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type ClassicPassportRequestOtpActionRequestSig = func(c ClassicPassportRequestOtpActionRequest) (*ClassicPassportRequestOtpActionResponse, error)

// ClassicPassportRequestOtpActionHandler returns the HTTP method, route URL, and a typed Gin handler for the ClassicPassportRequestOtpAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func ClassicPassportRequestOtpActionHandler(
	handler ClassicPassportRequestOtpActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := ClassicPassportRequestOtpActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body ClassicPassportRequestOtpActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := ClassicPassportRequestOtpActionRequest{
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

// ClassicPassportRequestOtpAction is a high-level convenience wrapper around ClassicPassportRequestOtpActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func ClassicPassportRequestOtpActionGin(r gin.IRoutes, handler ClassicPassportRequestOtpActionRequestSig) {
	method, url, h := ClassicPassportRequestOtpActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for ClassicPassportRequestOtpAction
 */
// Query wrapper with private fields
type ClassicPassportRequestOtpActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func ClassicPassportRequestOtpActionQueryFromString(rawQuery string) ClassicPassportRequestOtpActionQuery {
	v := ClassicPassportRequestOtpActionQuery{}
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
func ClassicPassportRequestOtpActionQueryFromGin(c *gin.Context) ClassicPassportRequestOtpActionQuery {
	return ClassicPassportRequestOtpActionQueryFromString(c.Request.URL.RawQuery)
}
func ClassicPassportRequestOtpActionQueryFromHttp(r *http.Request) ClassicPassportRequestOtpActionQuery {
	return ClassicPassportRequestOtpActionQueryFromString(r.URL.RawQuery)
}
func (q ClassicPassportRequestOtpActionQuery) Values() url.Values {
	return q.values
}
func (q ClassicPassportRequestOtpActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *ClassicPassportRequestOtpActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *ClassicPassportRequestOtpActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type ClassicPassportRequestOtpActionRequest struct {
	Body        ClassicPassportRequestOtpActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type ClassicPassportRequestOtpActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func ClassicPassportRequestOtpActionCall(
	req ClassicPassportRequestOtpActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*ClassicPassportRequestOtpActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := ClassicPassportRequestOtpActionMeta()
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
	var result ClassicPassportRequestOtpActionResult
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

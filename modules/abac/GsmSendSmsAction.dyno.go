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
* Action to communicate with the action GsmSendSmsAction
 */
func GsmSendSmsActionMeta() struct {
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
		Name:        "GsmSendSmsAction",
		CliName:     "sms",
		URL:         "/gsm/send/sms",
		Method:      "POST",
		Description: `Send a text message using default root notification configuration`,
	}
}
func GetGsmSendSmsActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "to-number",
			Type: "string",
		},
		{
			Name: prefix + "body",
			Type: "string",
		},
	}
}
func CastGsmSendSmsActionReqFromCli(c emigo.CliCastable) GsmSendSmsActionReq {
	data := GsmSendSmsActionReq{}
	if c.IsSet("to-number") {
		data.ToNumber = c.String("to-number")
	}
	if c.IsSet("body") {
		data.Body = c.String("body")
	}
	return data
}

// The base class definition for gsmSendSmsActionReq
type GsmSendSmsActionReq struct {
	ToNumber string `json:"toNumber" validate:"required" yaml:"toNumber"`
	Body     string `json:"body" validate:"required" yaml:"body"`
}

func (x *GsmSendSmsActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetGsmSendSmsActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "queue-id",
			Type: "string",
		},
	}
}
func CastGsmSendSmsActionResFromCli(c emigo.CliCastable) GsmSendSmsActionRes {
	data := GsmSendSmsActionRes{}
	if c.IsSet("queue-id") {
		data.QueueId = c.String("queue-id")
	}
	return data
}

// The base class definition for gsmSendSmsActionRes
type GsmSendSmsActionRes struct {
	QueueId string `json:"queueId" yaml:"queueId"`
}

func (x *GsmSendSmsActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GsmSendSmsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *GsmSendSmsActionResponse) SetContentType(contentType string) *GsmSendSmsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GsmSendSmsActionResponse) AsStream(r io.Reader, contentType string) *GsmSendSmsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GsmSendSmsActionResponse) AsJSON(payload any) *GsmSendSmsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GsmSendSmsActionResponse) WithIdeal(payload GsmSendSmsActionRes) *GsmSendSmsActionResponse {
	x.Payload = payload
	return x
}
func (x *GsmSendSmsActionResponse) AsHTML(payload string) *GsmSendSmsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GsmSendSmsActionResponse) AsBytes(payload []byte) *GsmSendSmsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GsmSendSmsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GsmSendSmsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GsmSendSmsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GsmSendSmsActionRaw registers a raw Gin route for the GsmSendSmsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GsmSendSmsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GsmSendSmsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GsmSendSmsActionRequestSig = func(c GsmSendSmsActionRequest) (*GsmSendSmsActionResponse, error)

// GsmSendSmsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GsmSendSmsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GsmSendSmsActionHandler(
	handler GsmSendSmsActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GsmSendSmsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body GsmSendSmsActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := GsmSendSmsActionRequest{
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

// GsmSendSmsAction is a high-level convenience wrapper around GsmSendSmsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GsmSendSmsActionGin(r gin.IRoutes, handler GsmSendSmsActionRequestSig) {
	method, url, h := GsmSendSmsActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for GsmSendSmsAction
 */
// Query wrapper with private fields
type GsmSendSmsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GsmSendSmsActionQueryFromString(rawQuery string) GsmSendSmsActionQuery {
	v := GsmSendSmsActionQuery{}
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
func GsmSendSmsActionQueryFromGin(c *gin.Context) GsmSendSmsActionQuery {
	return GsmSendSmsActionQueryFromString(c.Request.URL.RawQuery)
}
func GsmSendSmsActionQueryFromHttp(r *http.Request) GsmSendSmsActionQuery {
	return GsmSendSmsActionQueryFromString(r.URL.RawQuery)
}
func (q GsmSendSmsActionQuery) Values() url.Values {
	return q.values
}
func (q GsmSendSmsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GsmSendSmsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GsmSendSmsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GsmSendSmsActionRequest struct {
	Body        GsmSendSmsActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type GsmSendSmsActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func GsmSendSmsActionCall(
	req GsmSendSmsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GsmSendSmsActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GsmSendSmsActionMeta()
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
	var result GsmSendSmsActionResult
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

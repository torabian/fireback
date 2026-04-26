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
* Action to communicate with the action GsmSendSmsWithProviderAction
 */
func GsmSendSmsWithProviderActionMeta() struct {
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
		Name:        "GsmSendSmsWithProviderAction",
		CliName:     "smsp",
		URL:         "/gsmProvider/send/sms",
		Method:      "POST",
		Description: `Send a text message using an specific gsm provider`,
	}
}
func GetGsmSendSmsWithProviderActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "gsm-provider-id",
			Type: "string",
		},
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
func CastGsmSendSmsWithProviderActionReqFromCli(c emigo.CliCastable) GsmSendSmsWithProviderActionReq {
	data := GsmSendSmsWithProviderActionReq{}
	if c.IsSet("gsm-provider-id") {
		data.GsmProviderId = c.String("gsm-provider-id")
	}
	if c.IsSet("to-number") {
		data.ToNumber = c.String("to-number")
	}
	if c.IsSet("body") {
		data.Body = c.String("body")
	}
	return data
}

// The base class definition for gsmSendSmsWithProviderActionReq
type GsmSendSmsWithProviderActionReq struct {
	GsmProviderId string `json:"gsmProviderId" yaml:"gsmProviderId"`
	ToNumber      string `json:"toNumber" validate:"required" yaml:"toNumber"`
	Body          string `json:"body" validate:"required" yaml:"body"`
}

func (x *GsmSendSmsWithProviderActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetGsmSendSmsWithProviderActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "queue-id",
			Type: "string",
		},
	}
}
func CastGsmSendSmsWithProviderActionResFromCli(c emigo.CliCastable) GsmSendSmsWithProviderActionRes {
	data := GsmSendSmsWithProviderActionRes{}
	if c.IsSet("queue-id") {
		data.QueueId = c.String("queue-id")
	}
	return data
}

// The base class definition for gsmSendSmsWithProviderActionRes
type GsmSendSmsWithProviderActionRes struct {
	QueueId string `json:"queueId" yaml:"queueId"`
}

func (x *GsmSendSmsWithProviderActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GsmSendSmsWithProviderActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *GsmSendSmsWithProviderActionResponse) SetContentType(contentType string) *GsmSendSmsWithProviderActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GsmSendSmsWithProviderActionResponse) AsStream(r io.Reader, contentType string) *GsmSendSmsWithProviderActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GsmSendSmsWithProviderActionResponse) AsJSON(payload any) *GsmSendSmsWithProviderActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GsmSendSmsWithProviderActionResponse) WithIdeal(payload GsmSendSmsWithProviderActionRes) *GsmSendSmsWithProviderActionResponse {
	x.Payload = payload
	return x
}
func (x *GsmSendSmsWithProviderActionResponse) AsHTML(payload string) *GsmSendSmsWithProviderActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GsmSendSmsWithProviderActionResponse) AsBytes(payload []byte) *GsmSendSmsWithProviderActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GsmSendSmsWithProviderActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GsmSendSmsWithProviderActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GsmSendSmsWithProviderActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GsmSendSmsWithProviderActionRaw registers a raw Gin route for the GsmSendSmsWithProviderAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GsmSendSmsWithProviderActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GsmSendSmsWithProviderActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GsmSendSmsWithProviderActionRequestSig = func(c GsmSendSmsWithProviderActionRequest) (*GsmSendSmsWithProviderActionResponse, error)

// GsmSendSmsWithProviderActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GsmSendSmsWithProviderAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GsmSendSmsWithProviderActionHandler(
	handler GsmSendSmsWithProviderActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GsmSendSmsWithProviderActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body GsmSendSmsWithProviderActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := GsmSendSmsWithProviderActionRequest{
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

// GsmSendSmsWithProviderAction is a high-level convenience wrapper around GsmSendSmsWithProviderActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GsmSendSmsWithProviderActionGin(r gin.IRoutes, handler GsmSendSmsWithProviderActionRequestSig) {
	method, url, h := GsmSendSmsWithProviderActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for GsmSendSmsWithProviderAction
 */
// Query wrapper with private fields
type GsmSendSmsWithProviderActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GsmSendSmsWithProviderActionQueryFromString(rawQuery string) GsmSendSmsWithProviderActionQuery {
	v := GsmSendSmsWithProviderActionQuery{}
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
func GsmSendSmsWithProviderActionQueryFromGin(c *gin.Context) GsmSendSmsWithProviderActionQuery {
	return GsmSendSmsWithProviderActionQueryFromString(c.Request.URL.RawQuery)
}
func GsmSendSmsWithProviderActionQueryFromHttp(r *http.Request) GsmSendSmsWithProviderActionQuery {
	return GsmSendSmsWithProviderActionQueryFromString(r.URL.RawQuery)
}
func (q GsmSendSmsWithProviderActionQuery) Values() url.Values {
	return q.values
}
func (q GsmSendSmsWithProviderActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GsmSendSmsWithProviderActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GsmSendSmsWithProviderActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GsmSendSmsWithProviderActionRequest struct {
	Body        GsmSendSmsWithProviderActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type GsmSendSmsWithProviderActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func GsmSendSmsWithProviderActionCall(
	req GsmSendSmsWithProviderActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GsmSendSmsWithProviderActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GsmSendSmsWithProviderActionMeta()
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
	var result GsmSendSmsWithProviderActionResult
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

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
* Action to communicate with the action SendEmailWithProviderAction
 */
func SendEmailWithProviderActionMeta() struct {
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
		Name:        "SendEmailWithProviderAction",
		CliName:     "emailp",
		URL:         "/emailProvider/send",
		Method:      "POST",
		Description: `Send a text message using an specific gsm provider`,
	}
}
func GetSendEmailWithProviderActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "email-provider",
			Type: "one",
		},
		{
			Name: prefix + "to-address",
			Type: "string",
		},
		{
			Name: prefix + "body",
			Type: "string",
		},
	}
}
func CastSendEmailWithProviderActionReqFromCli(c emigo.CliCastable) SendEmailWithProviderActionReq {
	data := SendEmailWithProviderActionReq{}
	if c.IsSet("to-address") {
		data.ToAddress = c.String("to-address")
	}
	if c.IsSet("body") {
		data.Body = c.String("body")
	}
	return data
}

// The base class definition for sendEmailWithProviderActionReq
type SendEmailWithProviderActionReq struct {
	EmailProvider EmailProviderEntity `json:"emailProvider" yaml:"emailProvider"`
	ToAddress     string              `json:"toAddress" validate:"required" yaml:"toAddress"`
	Body          string              `json:"body" validate:"required" yaml:"body"`
}

func (x *SendEmailWithProviderActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSendEmailWithProviderActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "queue-id",
			Type: "string",
		},
	}
}
func CastSendEmailWithProviderActionResFromCli(c emigo.CliCastable) SendEmailWithProviderActionRes {
	data := SendEmailWithProviderActionRes{}
	if c.IsSet("queue-id") {
		data.QueueId = c.String("queue-id")
	}
	return data
}

// The base class definition for sendEmailWithProviderActionRes
type SendEmailWithProviderActionRes struct {
	QueueId string `json:"queueId" yaml:"queueId"`
}

func (x *SendEmailWithProviderActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type SendEmailWithProviderActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *SendEmailWithProviderActionResponse) SetContentType(contentType string) *SendEmailWithProviderActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *SendEmailWithProviderActionResponse) AsStream(r io.Reader, contentType string) *SendEmailWithProviderActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *SendEmailWithProviderActionResponse) AsJSON(payload any) *SendEmailWithProviderActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *SendEmailWithProviderActionResponse) AsHTML(payload string) *SendEmailWithProviderActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *SendEmailWithProviderActionResponse) AsBytes(payload []byte) *SendEmailWithProviderActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x SendEmailWithProviderActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x SendEmailWithProviderActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x SendEmailWithProviderActionResponse) GetPayload() interface{} {
	return x.Payload
}

// SendEmailWithProviderActionRaw registers a raw Gin route for the SendEmailWithProviderAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func SendEmailWithProviderActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := SendEmailWithProviderActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type SendEmailWithProviderActionRequestSig = func(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error)

// SendEmailWithProviderActionHandler returns the HTTP method, route URL, and a typed Gin handler for the SendEmailWithProviderAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func SendEmailWithProviderActionHandler(
	handler SendEmailWithProviderActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := SendEmailWithProviderActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body SendEmailWithProviderActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := SendEmailWithProviderActionRequest{
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

// SendEmailWithProviderAction is a high-level convenience wrapper around SendEmailWithProviderActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func SendEmailWithProviderActionGin(r gin.IRoutes, handler SendEmailWithProviderActionRequestSig) {
	method, url, h := SendEmailWithProviderActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for SendEmailWithProviderAction
 */
// Query wrapper with private fields
type SendEmailWithProviderActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func SendEmailWithProviderActionQueryFromString(rawQuery string) SendEmailWithProviderActionQuery {
	v := SendEmailWithProviderActionQuery{}
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
func SendEmailWithProviderActionQueryFromGin(c *gin.Context) SendEmailWithProviderActionQuery {
	return SendEmailWithProviderActionQueryFromString(c.Request.URL.RawQuery)
}
func SendEmailWithProviderActionQueryFromHttp(r *http.Request) SendEmailWithProviderActionQuery {
	return SendEmailWithProviderActionQueryFromString(r.URL.RawQuery)
}
func (q SendEmailWithProviderActionQuery) Values() url.Values {
	return q.values
}
func (q SendEmailWithProviderActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *SendEmailWithProviderActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *SendEmailWithProviderActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type SendEmailWithProviderActionRequest struct {
	Body        SendEmailWithProviderActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type SendEmailWithProviderActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func SendEmailWithProviderActionCall(
	req SendEmailWithProviderActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*SendEmailWithProviderActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := SendEmailWithProviderActionMeta()
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
	var result SendEmailWithProviderActionResult
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

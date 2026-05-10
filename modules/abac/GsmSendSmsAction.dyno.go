package abac

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GsmSendSmsAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GsmSendSmsAction
func GsmSendSmsAction(c GsmSendSmsActionRequest) (*GsmSendSmsActionResponse, error) {
	return &GsmSendSmsActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
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
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
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

func (x GsmSendSmsActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GsmSendSmsActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GsmSendSmsActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GsmSendSmsActionClientCreateUrl(
	req GsmSendSmsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GsmSendSmsActionMeta()
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
func GsmSendSmsActionClientExecuteTyped(httpReq *http.Request) (*GsmSendSmsActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GsmSendSmsActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GsmSendSmsActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GsmSendSmsActionResponse{Payload: result}, err
	}
	return &GsmSendSmsActionResponse{Payload: result}, nil
}
func GsmSendSmsActionClientBuildRequest(req GsmSendSmsActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GsmSendSmsActionMeta()
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), bytes.NewReader(bodyBytes))
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
func GsmSendSmsActionCall(
	req GsmSendSmsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GsmSendSmsActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GsmSendSmsActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GsmSendSmsActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GsmSendSmsActionClientExecuteTyped(r)
}

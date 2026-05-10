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
* Action to communicate with the action SendEmailAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of SendEmailAction
func SendEmailAction(c SendEmailActionRequest) (*SendEmailActionResponse, error) {
	return &SendEmailActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func SendEmailActionMeta() struct {
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
		Name:        "SendEmailAction",
		CliName:     "email",
		URL:         "/email/send",
		Method:      "POST",
		Description: `Send a email using default root notification configuration`,
	}
}
func GetSendEmailActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "provider-id",
			Type:        "string",
			Description: "Sending a test email requires to be sent through an specific email provider.",
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
func CastSendEmailActionReqFromCli(c emigo.CliCastable) SendEmailActionReq {
	data := SendEmailActionReq{}
	if c.IsSet("provider-id") {
		data.ProviderId = c.String("provider-id")
	}
	if c.IsSet("to-address") {
		data.ToAddress = c.String("to-address")
	}
	if c.IsSet("body") {
		data.Body = c.String("body")
	}
	return data
}

// The base class definition for sendEmailActionReq
type SendEmailActionReq struct {
	// Sending a test email requires to be sent through an specific email provider.
	ProviderId string `json:"providerId" yaml:"providerId"`
	ToAddress  string `json:"toAddress" validate:"required" yaml:"toAddress"`
	Body       string `json:"body" validate:"required" yaml:"body"`
}

func (x *SendEmailActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSendEmailActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "queue-id",
			Type: "string",
		},
	}
}
func CastSendEmailActionResFromCli(c emigo.CliCastable) SendEmailActionRes {
	data := SendEmailActionRes{}
	if c.IsSet("queue-id") {
		data.QueueId = c.String("queue-id")
	}
	return data
}

// The base class definition for sendEmailActionRes
type SendEmailActionRes struct {
	QueueId string `json:"queueId" yaml:"queueId"`
}

func (x *SendEmailActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type SendEmailActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *SendEmailActionResponse) SetContentType(contentType string) *SendEmailActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *SendEmailActionResponse) AsStream(r io.Reader, contentType string) *SendEmailActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *SendEmailActionResponse) AsJSON(payload any) *SendEmailActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *SendEmailActionResponse) WithIdeal(payload SendEmailActionRes) *SendEmailActionResponse {
	x.Payload = payload
	return x
}
func (x *SendEmailActionResponse) AsHTML(payload string) *SendEmailActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *SendEmailActionResponse) AsBytes(payload []byte) *SendEmailActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x SendEmailActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x SendEmailActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x SendEmailActionResponse) GetPayload() interface{} {
	return x.Payload
}

// SendEmailActionRaw registers a raw Gin route for the SendEmailAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func SendEmailActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := SendEmailActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type SendEmailActionRequestSig = func(c SendEmailActionRequest) (*SendEmailActionResponse, error)

// SendEmailActionHandler returns the HTTP method, route URL, and a typed Gin handler for the SendEmailAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func SendEmailActionHandler(
	handler SendEmailActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := SendEmailActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body SendEmailActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := SendEmailActionRequest{
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

// SendEmailAction is a high-level convenience wrapper around SendEmailActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func SendEmailActionGin(r gin.IRoutes, handler SendEmailActionRequestSig) {
	method, url, h := SendEmailActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for SendEmailAction
 */
// Query wrapper with private fields
type SendEmailActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func SendEmailActionQueryFromString(rawQuery string) SendEmailActionQuery {
	v := SendEmailActionQuery{}
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
func SendEmailActionQueryFromGin(c *gin.Context) SendEmailActionQuery {
	return SendEmailActionQueryFromString(c.Request.URL.RawQuery)
}
func SendEmailActionQueryFromHttp(r *http.Request) SendEmailActionQuery {
	return SendEmailActionQueryFromString(r.URL.RawQuery)
}
func (q SendEmailActionQuery) Values() url.Values {
	return q.values
}
func (q SendEmailActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *SendEmailActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *SendEmailActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type SendEmailActionRequest struct {
	Body        SendEmailActionReq
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

func (x SendEmailActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x SendEmailActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type SendEmailActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func SendEmailActionClientCreateUrl(
	req SendEmailActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := SendEmailActionMeta()
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
func SendEmailActionClientExecuteTyped(httpReq *http.Request) (*SendEmailActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result SendEmailActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &SendEmailActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &SendEmailActionResponse{Payload: result}, err
	}
	return &SendEmailActionResponse{Payload: result}, nil
}
func SendEmailActionClientBuildRequest(req SendEmailActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := SendEmailActionMeta()
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
func SendEmailActionCall(
	req SendEmailActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*SendEmailActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := SendEmailActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := SendEmailActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return SendEmailActionClientExecuteTyped(r)
}

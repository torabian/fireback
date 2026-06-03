package abac

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action SendEmailWithProviderAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of SendEmailWithProviderAction
func SendEmailWithProviderAction(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error) {
	return &SendEmailWithProviderActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
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
	if c.IsSet("email-provider") {
		data.EmailProvider = emigo.CapturePossibleOne(CastEmailProviderEntityFromCli, "email-provider", c)
	}
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
	EmailProvider emigo.One[EmailProviderEntity] `json:"emailProvider" yaml:"emailProvider"`
	ToAddress     string                         `json:"toAddress" validate:"required" yaml:"toAddress"`
	Body          string                         `json:"body" validate:"required" yaml:"body"`
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
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
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

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *SendEmailWithProviderActionResponse) WithIdeal(payload SendEmailWithProviderActionRes) *SendEmailWithProviderActionResponse {
	x.Payload = payload
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

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type SendEmailWithProviderActionRequestSig = func(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error)

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
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	// Now it's interface, so the code gen doesn't depend on the instance
	// or gin package. Make sure you cast is later into *gin.Context, or whatever
	// your framework is passing when creating a request.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	GinCtx interface{}
	// Cli library helper (urfave) by default. The instance is interface{}, and you
	// need to manually cast it to the *cli.Command, so gives you freedom and independence
	// of external library.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	CliCtx interface{}
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func SendEmailWithProviderActionClientCreateUrl(
	req SendEmailWithProviderActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := SendEmailWithProviderActionMeta()
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
func SendEmailWithProviderActionClientExecuteTyped(httpReq *http.Request) (*SendEmailWithProviderActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result SendEmailWithProviderActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &SendEmailWithProviderActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &SendEmailWithProviderActionResponse{Payload: result}, err
	}
	return &SendEmailWithProviderActionResponse{Payload: result}, nil
}
func SendEmailWithProviderActionClientBuildRequest(req SendEmailWithProviderActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := SendEmailWithProviderActionMeta()
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
func SendEmailWithProviderActionCall(
	req SendEmailWithProviderActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*SendEmailWithProviderActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := SendEmailWithProviderActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := SendEmailWithProviderActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return SendEmailWithProviderActionClientExecuteTyped(r)
}

// SendEmailWithProviderActionRaw registers a raw Gin route for the SendEmailWithProviderAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func SendEmailWithProviderActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := SendEmailWithProviderActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// SendEmailWithProviderActionHandler returns the HTTP method, route URL, and a typed Gin handler for the SendEmailWithProviderAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func SendEmailWithProviderActionHandler(
	handler func(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error),
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

// SendEmailWithProviderActionGin is a high-level convenience wrapper around SendEmailWithProviderActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func SendEmailWithProviderActionGin(r gin.IRoutes, handler func(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error)) {
	method, url, h := SendEmailWithProviderActionHandler(handler)
	r.Handle(method, url, h)
}
func (x SendEmailWithProviderActionRequest) IsGin() bool {
	if x.GinCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.GinCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}
func SendEmailWithProviderActionQueryFromGin(c *gin.Context) SendEmailWithProviderActionQuery {
	return SendEmailWithProviderActionQueryFromString(c.Request.URL.RawQuery)
}

// SendEmailWithProviderActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the SendEmailWithProviderAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *SendEmailWithProviderActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func SendEmailWithProviderActionHttpHandler(
	handler func(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := SendEmailWithProviderActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body SendEmailWithProviderActionReq
		if r.Body != nil {
			defer r.Body.Close()
			if data, _ := io.ReadAll(r.Body); len(data) > 0 {
				if err := json.Unmarshal(data, &body); err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON: " + err.Error()})
					return
				}
			}
		}
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := SendEmailWithProviderActionRequest{
			Body:        body,
			QueryParams: r.URL.Query(),
			Headers:     r.Header,
		}
		resp, err := handler(req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), the response was handled
		// manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			if w.Header().Get("Content-Type") == "" {
				w.Header().Set("Content-Type", "application/json")
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(resp.Payload)
		} else {
			w.WriteHeader(status)
		}
	}
}

// SendEmailWithProviderActionHttp is a high-level convenience wrapper around
// SendEmailWithProviderActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func SendEmailWithProviderActionHttp(
	mux *http.ServeMux,
	handler func(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error),
) {
	method, pattern, h := SendEmailWithProviderActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

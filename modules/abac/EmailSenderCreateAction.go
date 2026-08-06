package abac

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action EmailSenderCreateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of EmailSenderCreateAction
func EmailSenderCreateAction(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error) {
	return &EmailSenderCreateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func EmailSenderCreateActionMeta() struct {
	Name        string
	CliName     string
	CliShort    string
	URL         string
	Method      string
	Description string
} {
	return struct {
		Name        string
		CliName     string
		CliShort    string
		URL         string
		Method      string
		Description string
	}{
		Name:        "EmailSenderCreateAction",
		CliName:     "email-sender-create-action",
		CliShort:    "emailSender-c",
		URL:         "/emailSender",
		Method:      "POST",
		Description: `Creates a new "emailSender" row.`,
	}
}

type EmailSenderCreateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *EmailSenderCreateActionResponse) SetContentType(contentType string) *EmailSenderCreateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *EmailSenderCreateActionResponse) AsStream(r io.Reader, contentType string) *EmailSenderCreateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *EmailSenderCreateActionResponse) AsJSON(payload any) *EmailSenderCreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *EmailSenderCreateActionResponse) WithIdeal(payload EmailSenderDto) *EmailSenderCreateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *EmailSenderCreateActionResponse) AsIdeal() (*EmailSenderDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res EmailSenderDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *EmailSenderCreateActionResponse) AsHTML(payload string) *EmailSenderCreateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *EmailSenderCreateActionResponse) AsBytes(payload []byte) *EmailSenderCreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x EmailSenderCreateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x EmailSenderCreateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x EmailSenderCreateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type EmailSenderCreateActionRequestSig = func(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error)

/**
 * Query parameters for EmailSenderCreateAction
 */
// Query wrapper with private fields
type EmailSenderCreateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func EmailSenderCreateActionQueryFromString(rawQuery string) EmailSenderCreateActionQuery {
	v := EmailSenderCreateActionQuery{}
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
func EmailSenderCreateActionQueryFromHttp(r *http.Request) EmailSenderCreateActionQuery {
	return EmailSenderCreateActionQueryFromString(r.URL.RawQuery)
}
func (q EmailSenderCreateActionQuery) Values() url.Values {
	return q.values
}
func (q EmailSenderCreateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *EmailSenderCreateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *EmailSenderCreateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type EmailSenderCreateActionRequest struct {
	Body        EmailSenderDto
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

// Returns the gin ctx. You need to manually cast this to .(*gin.Context)
func (x EmailSenderCreateActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x EmailSenderCreateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func EmailSenderCreateActionClientCreateUrl(
	req EmailSenderCreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := EmailSenderCreateActionMeta()
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
func EmailSenderCreateActionClientExecuteTyped(httpReq *http.Request) (*EmailSenderCreateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result EmailSenderCreateActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}
func EmailSenderCreateActionClientBuildRequest(req EmailSenderCreateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := EmailSenderCreateActionMeta()
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
func EmailSenderCreateActionCall(
	req EmailSenderCreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*EmailSenderCreateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := EmailSenderCreateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := EmailSenderCreateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return EmailSenderCreateActionClientExecuteTyped(r)
}

// EmailSenderCreateActionRaw registers a raw Gin route for the EmailSenderCreateAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func EmailSenderCreateActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := EmailSenderCreateActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// EmailSenderCreateActionHandler returns the HTTP method, route URL, and a typed Gin handler for the EmailSenderCreateAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func EmailSenderCreateActionHandler(
	handler func(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := EmailSenderCreateActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body EmailSenderDto
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := EmailSenderCreateActionRequest{
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

// EmailSenderCreateActionGin is a high-level convenience wrapper around EmailSenderCreateActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func EmailSenderCreateActionGin(r gin.IRoutes, handler func(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error)) {
	method, url, h := EmailSenderCreateActionHandler(handler)
	r.Handle(method, url, h)
}
func (x EmailSenderCreateActionRequest) IsGin() bool {
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
func EmailSenderCreateActionQueryFromGin(c *gin.Context) EmailSenderCreateActionQuery {
	return EmailSenderCreateActionQueryFromString(c.Request.URL.RawQuery)
}
func (x EmailSenderCreateActionRequest) IsCli() bool {
	if x.CliCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.CliCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}

// EmailSenderCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the EmailSenderCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func EmailSenderCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// EmailSenderCreateActionCliHandler builds a full *cli.Command for the
// EmailSenderCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a EmailSenderCreateActionRequest the same way
// EmailSenderCreateActionHandler (Gin) and EmailSenderCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func EmailSenderCreateActionCliHandler(
	handler func(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error),
) *cli.Command {
	meta := EmailSenderCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: EmailSenderCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := EmailSenderCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// EmailSenderCreateActionCli is a high-level convenience wrapper around
// EmailSenderCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way EmailSenderCreateActionGin
// registers a route on a Gin engine.
func EmailSenderCreateActionCli(
	app *cli.Command,
	handler func(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, EmailSenderCreateActionCliHandler(handler))
}

// EmailSenderCreateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the EmailSenderCreateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *EmailSenderCreateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func EmailSenderCreateActionHttpHandler(
	handler func(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := EmailSenderCreateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body EmailSenderDto
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
		req := EmailSenderCreateActionRequest{
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

// EmailSenderCreateActionHttp is a high-level convenience wrapper around
// EmailSenderCreateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func EmailSenderCreateActionHttp(
	mux *http.ServeMux,
	handler func(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error),
) {
	method, pattern, h := EmailSenderCreateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

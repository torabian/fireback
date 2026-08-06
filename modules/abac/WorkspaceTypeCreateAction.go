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
* Action to communicate with the action WorkspaceTypeCreateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of WorkspaceTypeCreateAction
func WorkspaceTypeCreateAction(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error) {
	return &WorkspaceTypeCreateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func WorkspaceTypeCreateActionMeta() struct {
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
		Name:        "WorkspaceTypeCreateAction",
		CliName:     "workspace-type-create-action",
		CliShort:    "workspaceType-c",
		URL:         "/workspaceType",
		Method:      "POST",
		Description: `Creates a new "workspaceType" row.`,
	}
}

type WorkspaceTypeCreateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *WorkspaceTypeCreateActionResponse) SetContentType(contentType string) *WorkspaceTypeCreateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *WorkspaceTypeCreateActionResponse) AsStream(r io.Reader, contentType string) *WorkspaceTypeCreateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *WorkspaceTypeCreateActionResponse) AsJSON(payload any) *WorkspaceTypeCreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *WorkspaceTypeCreateActionResponse) WithIdeal(payload WorkspaceTypeDto) *WorkspaceTypeCreateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *WorkspaceTypeCreateActionResponse) AsIdeal() (*WorkspaceTypeDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res WorkspaceTypeDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *WorkspaceTypeCreateActionResponse) AsHTML(payload string) *WorkspaceTypeCreateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *WorkspaceTypeCreateActionResponse) AsBytes(payload []byte) *WorkspaceTypeCreateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x WorkspaceTypeCreateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x WorkspaceTypeCreateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x WorkspaceTypeCreateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type WorkspaceTypeCreateActionRequestSig = func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error)

/**
 * Query parameters for WorkspaceTypeCreateAction
 */
// Query wrapper with private fields
type WorkspaceTypeCreateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func WorkspaceTypeCreateActionQueryFromString(rawQuery string) WorkspaceTypeCreateActionQuery {
	v := WorkspaceTypeCreateActionQuery{}
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
func WorkspaceTypeCreateActionQueryFromHttp(r *http.Request) WorkspaceTypeCreateActionQuery {
	return WorkspaceTypeCreateActionQueryFromString(r.URL.RawQuery)
}
func (q WorkspaceTypeCreateActionQuery) Values() url.Values {
	return q.values
}
func (q WorkspaceTypeCreateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *WorkspaceTypeCreateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *WorkspaceTypeCreateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type WorkspaceTypeCreateActionRequest struct {
	Body        WorkspaceTypeDto
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
func (x WorkspaceTypeCreateActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x WorkspaceTypeCreateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func WorkspaceTypeCreateActionClientCreateUrl(
	req WorkspaceTypeCreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := WorkspaceTypeCreateActionMeta()
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
func WorkspaceTypeCreateActionClientExecuteTyped(httpReq *http.Request) (*WorkspaceTypeCreateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result WorkspaceTypeCreateActionResponse
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
func WorkspaceTypeCreateActionClientBuildRequest(req WorkspaceTypeCreateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := WorkspaceTypeCreateActionMeta()
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
func WorkspaceTypeCreateActionCall(
	req WorkspaceTypeCreateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*WorkspaceTypeCreateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := WorkspaceTypeCreateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := WorkspaceTypeCreateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return WorkspaceTypeCreateActionClientExecuteTyped(r)
}

// WorkspaceTypeCreateActionRaw registers a raw Gin route for the WorkspaceTypeCreateAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func WorkspaceTypeCreateActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := WorkspaceTypeCreateActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// WorkspaceTypeCreateActionHandler returns the HTTP method, route URL, and a typed Gin handler for the WorkspaceTypeCreateAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func WorkspaceTypeCreateActionHandler(
	handler func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := WorkspaceTypeCreateActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body WorkspaceTypeDto
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := WorkspaceTypeCreateActionRequest{
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

// WorkspaceTypeCreateActionGin is a high-level convenience wrapper around WorkspaceTypeCreateActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func WorkspaceTypeCreateActionGin(r gin.IRoutes, handler func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error)) {
	method, url, h := WorkspaceTypeCreateActionHandler(handler)
	r.Handle(method, url, h)
}
func (x WorkspaceTypeCreateActionRequest) IsGin() bool {
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
func WorkspaceTypeCreateActionQueryFromGin(c *gin.Context) WorkspaceTypeCreateActionQuery {
	return WorkspaceTypeCreateActionQueryFromString(c.Request.URL.RawQuery)
}
func (x WorkspaceTypeCreateActionRequest) IsCli() bool {
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

// WorkspaceTypeCreateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceTypeCreateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceTypeCreateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// WorkspaceTypeCreateActionCliHandler builds a full *cli.Command for the
// WorkspaceTypeCreateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceTypeCreateActionRequest the same way
// WorkspaceTypeCreateActionHandler (Gin) and WorkspaceTypeCreateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceTypeCreateActionCliHandler(
	handler func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error),
) *cli.Command {
	meta := WorkspaceTypeCreateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceTypeCreateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceTypeCreateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceTypeCreateActionCli is a high-level convenience wrapper around
// WorkspaceTypeCreateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceTypeCreateActionGin
// registers a route on a Gin engine.
func WorkspaceTypeCreateActionCli(
	app *cli.Command,
	handler func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceTypeCreateActionCliHandler(handler))
}

// WorkspaceTypeCreateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the WorkspaceTypeCreateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *WorkspaceTypeCreateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func WorkspaceTypeCreateActionHttpHandler(
	handler func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := WorkspaceTypeCreateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body WorkspaceTypeDto
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
		req := WorkspaceTypeCreateActionRequest{
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

// WorkspaceTypeCreateActionHttp is a high-level convenience wrapper around
// WorkspaceTypeCreateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func WorkspaceTypeCreateActionHttp(
	mux *http.ServeMux,
	handler func(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error),
) {
	method, pattern, h := WorkspaceTypeCreateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

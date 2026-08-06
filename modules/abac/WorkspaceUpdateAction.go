package abac

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

/**
* Action to communicate with the action WorkspaceUpdateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of WorkspaceUpdateAction
func WorkspaceUpdateAction(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error) {
	return &WorkspaceUpdateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func WorkspaceUpdateActionMeta() struct {
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
		Name:        "WorkspaceUpdateAction",
		CliName:     "workspace-update-action",
		CliShort:    "workspace-u",
		URL:         "/workspace/:uniqueId",
		Method:      "PATCH",
		Description: `Applies a partial update to a "workspace" row by uniqueId.`,
	}
}

type WorkspaceUpdateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *WorkspaceUpdateActionResponse) SetContentType(contentType string) *WorkspaceUpdateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *WorkspaceUpdateActionResponse) AsStream(r io.Reader, contentType string) *WorkspaceUpdateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *WorkspaceUpdateActionResponse) AsJSON(payload any) *WorkspaceUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *WorkspaceUpdateActionResponse) WithIdeal(payload WorkspaceDto) *WorkspaceUpdateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *WorkspaceUpdateActionResponse) AsIdeal() (*WorkspaceDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res WorkspaceDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *WorkspaceUpdateActionResponse) AsHTML(payload string) *WorkspaceUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *WorkspaceUpdateActionResponse) AsBytes(payload []byte) *WorkspaceUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x WorkspaceUpdateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x WorkspaceUpdateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x WorkspaceUpdateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type WorkspaceUpdateActionRequestSig = func(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error)

/**
 * Path parameters for WorkspaceUpdateAction
 */
type WorkspaceUpdateActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func WorkspaceUpdateActionPathParameterApply(params WorkspaceUpdateActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func WorkspaceUpdateActionPathParameterFromFn(fn func(key string) string) WorkspaceUpdateActionPathParameter {
	res := WorkspaceUpdateActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for WorkspaceUpdateAction
 */
// Query wrapper with private fields
type WorkspaceUpdateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func WorkspaceUpdateActionQueryFromString(rawQuery string) WorkspaceUpdateActionQuery {
	v := WorkspaceUpdateActionQuery{}
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
func WorkspaceUpdateActionQueryFromHttp(r *http.Request) WorkspaceUpdateActionQuery {
	return WorkspaceUpdateActionQueryFromString(r.URL.RawQuery)
}
func (q WorkspaceUpdateActionQuery) Values() url.Values {
	return q.values
}
func (q WorkspaceUpdateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *WorkspaceUpdateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *WorkspaceUpdateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type WorkspaceUpdateActionRequest struct {
	Body        WorkspaceOptionalDto
	Params      WorkspaceUpdateActionPathParameter
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
func (x WorkspaceUpdateActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x WorkspaceUpdateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func WorkspaceUpdateActionClientCreateUrl(
	req WorkspaceUpdateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := WorkspaceUpdateActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = WorkspaceUpdateActionPathParameterApply(req.Params, urlAddr)
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
func WorkspaceUpdateActionClientExecuteTyped(httpReq *http.Request) (*WorkspaceUpdateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result WorkspaceUpdateActionResponse
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
func WorkspaceUpdateActionClientBuildRequest(req WorkspaceUpdateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := WorkspaceUpdateActionMeta()
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
func WorkspaceUpdateActionCall(
	req WorkspaceUpdateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*WorkspaceUpdateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := WorkspaceUpdateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := WorkspaceUpdateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return WorkspaceUpdateActionClientExecuteTyped(r)
}
func WorkspaceUpdateActionPathParameterFromGin(g *gin.Context) WorkspaceUpdateActionPathParameter {
	return WorkspaceUpdateActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// WorkspaceUpdateActionRaw registers a raw Gin route for the WorkspaceUpdateAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func WorkspaceUpdateActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := WorkspaceUpdateActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// WorkspaceUpdateActionHandler returns the HTTP method, route URL, and a typed Gin handler for the WorkspaceUpdateAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func WorkspaceUpdateActionHandler(
	handler func(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := WorkspaceUpdateActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body WorkspaceOptionalDto
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := WorkspaceUpdateActionRequest{
			Body:        body,
			Params:      WorkspaceUpdateActionPathParameterFromGin(m),
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

// WorkspaceUpdateActionGin is a high-level convenience wrapper around WorkspaceUpdateActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func WorkspaceUpdateActionGin(r gin.IRoutes, handler func(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error)) {
	method, url, h := WorkspaceUpdateActionHandler(handler)
	r.Handle(method, url, h)
}
func (x WorkspaceUpdateActionRequest) IsGin() bool {
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
func WorkspaceUpdateActionQueryFromGin(c *gin.Context) WorkspaceUpdateActionQuery {
	return WorkspaceUpdateActionQueryFromString(c.Request.URL.RawQuery)
}
func GetWorkspaceUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func WorkspaceUpdateActionPathParameterFromCli(c *cli.Command) WorkspaceUpdateActionPathParameter {
	return WorkspaceUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x WorkspaceUpdateActionRequest) IsCli() bool {
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

// WorkspaceUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// WorkspaceUpdateActionCliHandler builds a full *cli.Command for the
// WorkspaceUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceUpdateActionRequest the same way
// WorkspaceUpdateActionHandler (Gin) and WorkspaceUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceUpdateActionCliHandler(
	handler func(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error),
) *cli.Command {
	meta := WorkspaceUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      WorkspaceUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceUpdateActionCli is a high-level convenience wrapper around
// WorkspaceUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceUpdateActionGin
// registers a route on a Gin engine.
func WorkspaceUpdateActionCli(
	app *cli.Command,
	handler func(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceUpdateActionCliHandler(handler))
}

// WorkspaceUpdateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the WorkspaceUpdateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *WorkspaceUpdateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func WorkspaceUpdateActionHttpHandler(
	handler func(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := WorkspaceUpdateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body WorkspaceOptionalDto
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
		req := WorkspaceUpdateActionRequest{
			Body: body,
			Params: WorkspaceUpdateActionPathParameterFromFn(func(key string) string {
				return r.PathValue(key)
			}),
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

// WorkspaceUpdateActionHttp is a high-level convenience wrapper around
// WorkspaceUpdateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func WorkspaceUpdateActionHttp(
	mux *http.ServeMux,
	handler func(c WorkspaceUpdateActionRequest) (*WorkspaceUpdateActionResponse, error),
) {
	method, pattern, h := WorkspaceUpdateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

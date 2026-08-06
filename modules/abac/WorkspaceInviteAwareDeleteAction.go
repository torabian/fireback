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
* Action to communicate with the action WorkspaceInviteAwareDeleteAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of WorkspaceInviteAwareDeleteAction
func WorkspaceInviteAwareDeleteAction(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error) {
	return &WorkspaceInviteAwareDeleteActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func WorkspaceInviteAwareDeleteActionMeta() struct {
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
		Name:        "WorkspaceInviteAwareDeleteAction",
		CliName:     "workspace-invite-aware-delete-action",
		CliShort:    "workspaceInvite-d",
		URL:         "/workspaceInvite/delete",
		Method:      "POST",
		Description: `Deletes the given "workspaceInvite" uniqueIds, along with everything workspaceInviteAwareDeletePreview reports.`,
	}
}

// The base class definition for workspaceInviteAwareDeleteActionReq
type WorkspaceInviteAwareDeleteActionReq struct {
	UniqueIds []string `json:"uniqueIds" yaml:"uniqueIds"`
}

func (x *WorkspaceInviteAwareDeleteActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWorkspaceInviteAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastWorkspaceInviteAwareDeleteActionReqFromCli(c emigo.CliCastable) WorkspaceInviteAwareDeleteActionReq {
	data := WorkspaceInviteAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}

type WorkspaceInviteAwareDeleteActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *WorkspaceInviteAwareDeleteActionResponse) SetContentType(contentType string) *WorkspaceInviteAwareDeleteActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *WorkspaceInviteAwareDeleteActionResponse) AsStream(r io.Reader, contentType string) *WorkspaceInviteAwareDeleteActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *WorkspaceInviteAwareDeleteActionResponse) AsJSON(payload any) *WorkspaceInviteAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *WorkspaceInviteAwareDeleteActionResponse) AsHTML(payload string) *WorkspaceInviteAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *WorkspaceInviteAwareDeleteActionResponse) AsBytes(payload []byte) *WorkspaceInviteAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x WorkspaceInviteAwareDeleteActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x WorkspaceInviteAwareDeleteActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x WorkspaceInviteAwareDeleteActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type WorkspaceInviteAwareDeleteActionRequestSig = func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error)

/**
 * Query parameters for WorkspaceInviteAwareDeleteAction
 */
// Query wrapper with private fields
type WorkspaceInviteAwareDeleteActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func WorkspaceInviteAwareDeleteActionQueryFromString(rawQuery string) WorkspaceInviteAwareDeleteActionQuery {
	v := WorkspaceInviteAwareDeleteActionQuery{}
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
func WorkspaceInviteAwareDeleteActionQueryFromHttp(r *http.Request) WorkspaceInviteAwareDeleteActionQuery {
	return WorkspaceInviteAwareDeleteActionQueryFromString(r.URL.RawQuery)
}
func (q WorkspaceInviteAwareDeleteActionQuery) Values() url.Values {
	return q.values
}
func (q WorkspaceInviteAwareDeleteActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *WorkspaceInviteAwareDeleteActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *WorkspaceInviteAwareDeleteActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type WorkspaceInviteAwareDeleteActionRequest struct {
	Body        WorkspaceInviteAwareDeleteActionReq
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
func (x WorkspaceInviteAwareDeleteActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x WorkspaceInviteAwareDeleteActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func WorkspaceInviteAwareDeleteActionClientCreateUrl(
	req WorkspaceInviteAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := WorkspaceInviteAwareDeleteActionMeta()
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
func WorkspaceInviteAwareDeleteActionClientExecuteTyped(httpReq *http.Request) (*WorkspaceInviteAwareDeleteActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result WorkspaceInviteAwareDeleteActionResponse
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
func WorkspaceInviteAwareDeleteActionClientBuildRequest(req WorkspaceInviteAwareDeleteActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := WorkspaceInviteAwareDeleteActionMeta()
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
func WorkspaceInviteAwareDeleteActionCall(
	req WorkspaceInviteAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*WorkspaceInviteAwareDeleteActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := WorkspaceInviteAwareDeleteActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := WorkspaceInviteAwareDeleteActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return WorkspaceInviteAwareDeleteActionClientExecuteTyped(r)
}

// WorkspaceInviteAwareDeleteActionRaw registers a raw Gin route for the WorkspaceInviteAwareDeleteAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func WorkspaceInviteAwareDeleteActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := WorkspaceInviteAwareDeleteActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// WorkspaceInviteAwareDeleteActionHandler returns the HTTP method, route URL, and a typed Gin handler for the WorkspaceInviteAwareDeleteAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func WorkspaceInviteAwareDeleteActionHandler(
	handler func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := WorkspaceInviteAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body WorkspaceInviteAwareDeleteActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := WorkspaceInviteAwareDeleteActionRequest{
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

// WorkspaceInviteAwareDeleteActionGin is a high-level convenience wrapper around WorkspaceInviteAwareDeleteActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func WorkspaceInviteAwareDeleteActionGin(r gin.IRoutes, handler func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error)) {
	method, url, h := WorkspaceInviteAwareDeleteActionHandler(handler)
	r.Handle(method, url, h)
}
func (x WorkspaceInviteAwareDeleteActionRequest) IsGin() bool {
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
func WorkspaceInviteAwareDeleteActionQueryFromGin(c *gin.Context) WorkspaceInviteAwareDeleteActionQuery {
	return WorkspaceInviteAwareDeleteActionQueryFromString(c.Request.URL.RawQuery)
}
func (x WorkspaceInviteAwareDeleteActionRequest) IsCli() bool {
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

// WorkspaceInviteAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WorkspaceInviteAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WorkspaceInviteAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWorkspaceInviteAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// WorkspaceInviteAwareDeleteActionCliHandler builds a full *cli.Command for the
// WorkspaceInviteAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WorkspaceInviteAwareDeleteActionRequest the same way
// WorkspaceInviteAwareDeleteActionHandler (Gin) and WorkspaceInviteAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WorkspaceInviteAwareDeleteActionCliHandler(
	handler func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error),
) *cli.Command {
	meta := WorkspaceInviteAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WorkspaceInviteAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WorkspaceInviteAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastWorkspaceInviteAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WorkspaceInviteAwareDeleteActionCli is a high-level convenience wrapper around
// WorkspaceInviteAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WorkspaceInviteAwareDeleteActionGin
// registers a route on a Gin engine.
func WorkspaceInviteAwareDeleteActionCli(
	app *cli.Command,
	handler func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, WorkspaceInviteAwareDeleteActionCliHandler(handler))
}

// WorkspaceInviteAwareDeleteActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the WorkspaceInviteAwareDeleteAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *WorkspaceInviteAwareDeleteActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func WorkspaceInviteAwareDeleteActionHttpHandler(
	handler func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := WorkspaceInviteAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body WorkspaceInviteAwareDeleteActionReq
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
		req := WorkspaceInviteAwareDeleteActionRequest{
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

// WorkspaceInviteAwareDeleteActionHttp is a high-level convenience wrapper around
// WorkspaceInviteAwareDeleteActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func WorkspaceInviteAwareDeleteActionHttp(
	mux *http.ServeMux,
	handler func(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error),
) {
	method, pattern, h := WorkspaceInviteAwareDeleteActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

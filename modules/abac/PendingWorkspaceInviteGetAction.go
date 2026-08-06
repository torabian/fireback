package abac

import (
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
* Action to communicate with the action PendingWorkspaceInviteGetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PendingWorkspaceInviteGetAction
func PendingWorkspaceInviteGetAction(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error) {
	return &PendingWorkspaceInviteGetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PendingWorkspaceInviteGetActionMeta() struct {
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
		Name:        "PendingWorkspaceInviteGetAction",
		CliName:     "pending-workspace-invite-get-action",
		CliShort:    "pendingWorkspaceInvite-g",
		URL:         "/pendingWorkspaceInvite/:uniqueId",
		Method:      "GET",
		Description: `Looks up a single "pendingWorkspaceInvite" row by uniqueId.`,
	}
}

type PendingWorkspaceInviteGetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PendingWorkspaceInviteGetActionResponse) SetContentType(contentType string) *PendingWorkspaceInviteGetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PendingWorkspaceInviteGetActionResponse) AsStream(r io.Reader, contentType string) *PendingWorkspaceInviteGetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PendingWorkspaceInviteGetActionResponse) AsJSON(payload any) *PendingWorkspaceInviteGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *PendingWorkspaceInviteGetActionResponse) WithIdeal(payload PendingWorkspaceInviteDto) *PendingWorkspaceInviteGetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *PendingWorkspaceInviteGetActionResponse) AsIdeal() (*PendingWorkspaceInviteDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res PendingWorkspaceInviteDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *PendingWorkspaceInviteGetActionResponse) AsHTML(payload string) *PendingWorkspaceInviteGetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PendingWorkspaceInviteGetActionResponse) AsBytes(payload []byte) *PendingWorkspaceInviteGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PendingWorkspaceInviteGetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PendingWorkspaceInviteGetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PendingWorkspaceInviteGetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type PendingWorkspaceInviteGetActionRequestSig = func(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error)

/**
 * Path parameters for PendingWorkspaceInviteGetAction
 */
type PendingWorkspaceInviteGetActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func PendingWorkspaceInviteGetActionPathParameterApply(params PendingWorkspaceInviteGetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func PendingWorkspaceInviteGetActionPathParameterFromFn(fn func(key string) string) PendingWorkspaceInviteGetActionPathParameter {
	res := PendingWorkspaceInviteGetActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for PendingWorkspaceInviteGetAction
 */
// Query wrapper with private fields
type PendingWorkspaceInviteGetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PendingWorkspaceInviteGetActionQueryFromString(rawQuery string) PendingWorkspaceInviteGetActionQuery {
	v := PendingWorkspaceInviteGetActionQuery{}
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
func PendingWorkspaceInviteGetActionQueryFromHttp(r *http.Request) PendingWorkspaceInviteGetActionQuery {
	return PendingWorkspaceInviteGetActionQueryFromString(r.URL.RawQuery)
}
func (q PendingWorkspaceInviteGetActionQuery) Values() url.Values {
	return q.values
}
func (q PendingWorkspaceInviteGetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PendingWorkspaceInviteGetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PendingWorkspaceInviteGetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PendingWorkspaceInviteGetActionRequest struct {
	Body        interface{}
	Params      PendingWorkspaceInviteGetActionPathParameter
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
func (x PendingWorkspaceInviteGetActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x PendingWorkspaceInviteGetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func PendingWorkspaceInviteGetActionClientCreateUrl(
	req PendingWorkspaceInviteGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PendingWorkspaceInviteGetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = PendingWorkspaceInviteGetActionPathParameterApply(req.Params, urlAddr)
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
func PendingWorkspaceInviteGetActionClientExecuteTyped(httpReq *http.Request) (*PendingWorkspaceInviteGetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PendingWorkspaceInviteGetActionResponse
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
func PendingWorkspaceInviteGetActionClientBuildRequest(req PendingWorkspaceInviteGetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PendingWorkspaceInviteGetActionMeta()
	httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), nil)
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
func PendingWorkspaceInviteGetActionCall(
	req PendingWorkspaceInviteGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PendingWorkspaceInviteGetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PendingWorkspaceInviteGetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PendingWorkspaceInviteGetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PendingWorkspaceInviteGetActionClientExecuteTyped(r)
}
func PendingWorkspaceInviteGetActionPathParameterFromGin(g *gin.Context) PendingWorkspaceInviteGetActionPathParameter {
	return PendingWorkspaceInviteGetActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// PendingWorkspaceInviteGetActionRaw registers a raw Gin route for the PendingWorkspaceInviteGetAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PendingWorkspaceInviteGetActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PendingWorkspaceInviteGetActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// PendingWorkspaceInviteGetActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PendingWorkspaceInviteGetAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PendingWorkspaceInviteGetActionHandler(
	handler func(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PendingWorkspaceInviteGetActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PendingWorkspaceInviteGetActionRequest{
			Body:        nil,
			Params:      PendingWorkspaceInviteGetActionPathParameterFromGin(m),
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

// PendingWorkspaceInviteGetActionGin is a high-level convenience wrapper around PendingWorkspaceInviteGetActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PendingWorkspaceInviteGetActionGin(r gin.IRoutes, handler func(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error)) {
	method, url, h := PendingWorkspaceInviteGetActionHandler(handler)
	r.Handle(method, url, h)
}
func (x PendingWorkspaceInviteGetActionRequest) IsGin() bool {
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
func PendingWorkspaceInviteGetActionQueryFromGin(c *gin.Context) PendingWorkspaceInviteGetActionQuery {
	return PendingWorkspaceInviteGetActionQueryFromString(c.Request.URL.RawQuery)
}
func GetPendingWorkspaceInviteGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func PendingWorkspaceInviteGetActionPathParameterFromCli(c *cli.Command) PendingWorkspaceInviteGetActionPathParameter {
	return PendingWorkspaceInviteGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x PendingWorkspaceInviteGetActionRequest) IsCli() bool {
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

// PendingWorkspaceInviteGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PendingWorkspaceInviteGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PendingWorkspaceInviteGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPendingWorkspaceInviteGetActionPathParameterCliFlags(""))...)
	return flags
}

// PendingWorkspaceInviteGetActionCliHandler builds a full *cli.Command for the
// PendingWorkspaceInviteGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PendingWorkspaceInviteGetActionRequest the same way
// PendingWorkspaceInviteGetActionHandler (Gin) and PendingWorkspaceInviteGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PendingWorkspaceInviteGetActionCliHandler(
	handler func(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error),
) *cli.Command {
	meta := PendingWorkspaceInviteGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PendingWorkspaceInviteGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PendingWorkspaceInviteGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      PendingWorkspaceInviteGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PendingWorkspaceInviteGetActionCli is a high-level convenience wrapper around
// PendingWorkspaceInviteGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PendingWorkspaceInviteGetActionGin
// registers a route on a Gin engine.
func PendingWorkspaceInviteGetActionCli(
	app *cli.Command,
	handler func(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error),
) {
	app.Commands = append(app.Commands, PendingWorkspaceInviteGetActionCliHandler(handler))
}

// PendingWorkspaceInviteGetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the PendingWorkspaceInviteGetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *PendingWorkspaceInviteGetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func PendingWorkspaceInviteGetActionHttpHandler(
	handler func(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := PendingWorkspaceInviteGetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := PendingWorkspaceInviteGetActionRequest{
			Body: nil,
			Params: PendingWorkspaceInviteGetActionPathParameterFromFn(func(key string) string {
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

// PendingWorkspaceInviteGetActionHttp is a high-level convenience wrapper around
// PendingWorkspaceInviteGetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func PendingWorkspaceInviteGetActionHttp(
	mux *http.ServeMux,
	handler func(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error),
) {
	method, pattern, h := PendingWorkspaceInviteGetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

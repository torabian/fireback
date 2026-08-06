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
* Action to communicate with the action CapabilityGetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CapabilityGetAction
func CapabilityGetAction(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error) {
	return &CapabilityGetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CapabilityGetActionMeta() struct {
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
		Name:        "CapabilityGetAction",
		CliName:     "capability-get-action",
		CliShort:    "capability-g",
		URL:         "/capability/:uniqueId",
		Method:      "GET",
		Description: `Looks up a single "capability" row by uniqueId.`,
	}
}

type CapabilityGetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CapabilityGetActionResponse) SetContentType(contentType string) *CapabilityGetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CapabilityGetActionResponse) AsStream(r io.Reader, contentType string) *CapabilityGetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CapabilityGetActionResponse) AsJSON(payload any) *CapabilityGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CapabilityGetActionResponse) WithIdeal(payload CapabilityDto) *CapabilityGetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *CapabilityGetActionResponse) AsIdeal() (*CapabilityDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res CapabilityDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *CapabilityGetActionResponse) AsHTML(payload string) *CapabilityGetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CapabilityGetActionResponse) AsBytes(payload []byte) *CapabilityGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CapabilityGetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CapabilityGetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CapabilityGetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type CapabilityGetActionRequestSig = func(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error)

/**
 * Path parameters for CapabilityGetAction
 */
type CapabilityGetActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func CapabilityGetActionPathParameterApply(params CapabilityGetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func CapabilityGetActionPathParameterFromFn(fn func(key string) string) CapabilityGetActionPathParameter {
	res := CapabilityGetActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for CapabilityGetAction
 */
// Query wrapper with private fields
type CapabilityGetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CapabilityGetActionQueryFromString(rawQuery string) CapabilityGetActionQuery {
	v := CapabilityGetActionQuery{}
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
func CapabilityGetActionQueryFromHttp(r *http.Request) CapabilityGetActionQuery {
	return CapabilityGetActionQueryFromString(r.URL.RawQuery)
}
func (q CapabilityGetActionQuery) Values() url.Values {
	return q.values
}
func (q CapabilityGetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CapabilityGetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CapabilityGetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CapabilityGetActionRequest struct {
	Body        interface{}
	Params      CapabilityGetActionPathParameter
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
func (x CapabilityGetActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x CapabilityGetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func CapabilityGetActionClientCreateUrl(
	req CapabilityGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CapabilityGetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = CapabilityGetActionPathParameterApply(req.Params, urlAddr)
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
func CapabilityGetActionClientExecuteTyped(httpReq *http.Request) (*CapabilityGetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CapabilityGetActionResponse
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
func CapabilityGetActionClientBuildRequest(req CapabilityGetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CapabilityGetActionMeta()
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
func CapabilityGetActionCall(
	req CapabilityGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CapabilityGetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CapabilityGetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CapabilityGetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CapabilityGetActionClientExecuteTyped(r)
}
func CapabilityGetActionPathParameterFromGin(g *gin.Context) CapabilityGetActionPathParameter {
	return CapabilityGetActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// CapabilityGetActionRaw registers a raw Gin route for the CapabilityGetAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CapabilityGetActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CapabilityGetActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// CapabilityGetActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CapabilityGetAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CapabilityGetActionHandler(
	handler func(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := CapabilityGetActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := CapabilityGetActionRequest{
			Body:        nil,
			Params:      CapabilityGetActionPathParameterFromGin(m),
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

// CapabilityGetActionGin is a high-level convenience wrapper around CapabilityGetActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CapabilityGetActionGin(r gin.IRoutes, handler func(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error)) {
	method, url, h := CapabilityGetActionHandler(handler)
	r.Handle(method, url, h)
}
func (x CapabilityGetActionRequest) IsGin() bool {
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
func CapabilityGetActionQueryFromGin(c *gin.Context) CapabilityGetActionQuery {
	return CapabilityGetActionQueryFromString(c.Request.URL.RawQuery)
}
func GetCapabilityGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func CapabilityGetActionPathParameterFromCli(c *cli.Command) CapabilityGetActionPathParameter {
	return CapabilityGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x CapabilityGetActionRequest) IsCli() bool {
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

// CapabilityGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the CapabilityGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func CapabilityGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetCapabilityGetActionPathParameterCliFlags(""))...)
	return flags
}

// CapabilityGetActionCliHandler builds a full *cli.Command for the
// CapabilityGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a CapabilityGetActionRequest the same way
// CapabilityGetActionHandler (Gin) and CapabilityGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func CapabilityGetActionCliHandler(
	handler func(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error),
) *cli.Command {
	meta := CapabilityGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: CapabilityGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := CapabilityGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      CapabilityGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// CapabilityGetActionCli is a high-level convenience wrapper around
// CapabilityGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way CapabilityGetActionGin
// registers a route on a Gin engine.
func CapabilityGetActionCli(
	app *cli.Command,
	handler func(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error),
) {
	app.Commands = append(app.Commands, CapabilityGetActionCliHandler(handler))
}

// CapabilityGetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the CapabilityGetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *CapabilityGetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func CapabilityGetActionHttpHandler(
	handler func(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := CapabilityGetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := CapabilityGetActionRequest{
			Body: nil,
			Params: CapabilityGetActionPathParameterFromFn(func(key string) string {
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

// CapabilityGetActionHttp is a high-level convenience wrapper around
// CapabilityGetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func CapabilityGetActionHttp(
	mux *http.ServeMux,
	handler func(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error),
) {
	method, pattern, h := CapabilityGetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

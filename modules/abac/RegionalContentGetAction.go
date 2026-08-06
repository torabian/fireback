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
* Action to communicate with the action RegionalContentGetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of RegionalContentGetAction
func RegionalContentGetAction(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error) {
	return &RegionalContentGetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func RegionalContentGetActionMeta() struct {
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
		Name:        "RegionalContentGetAction",
		CliName:     "regional-content-get-action",
		CliShort:    "regionalContent-g",
		URL:         "/regionalContent/:uniqueId",
		Method:      "GET",
		Description: `Looks up a single "regionalContent" row by uniqueId.`,
	}
}

type RegionalContentGetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *RegionalContentGetActionResponse) SetContentType(contentType string) *RegionalContentGetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *RegionalContentGetActionResponse) AsStream(r io.Reader, contentType string) *RegionalContentGetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *RegionalContentGetActionResponse) AsJSON(payload any) *RegionalContentGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *RegionalContentGetActionResponse) WithIdeal(payload RegionalContentDto) *RegionalContentGetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *RegionalContentGetActionResponse) AsIdeal() (*RegionalContentDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res RegionalContentDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *RegionalContentGetActionResponse) AsHTML(payload string) *RegionalContentGetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *RegionalContentGetActionResponse) AsBytes(payload []byte) *RegionalContentGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x RegionalContentGetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x RegionalContentGetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x RegionalContentGetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type RegionalContentGetActionRequestSig = func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error)

/**
 * Path parameters for RegionalContentGetAction
 */
type RegionalContentGetActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func RegionalContentGetActionPathParameterApply(params RegionalContentGetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func RegionalContentGetActionPathParameterFromFn(fn func(key string) string) RegionalContentGetActionPathParameter {
	res := RegionalContentGetActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for RegionalContentGetAction
 */
// Query wrapper with private fields
type RegionalContentGetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func RegionalContentGetActionQueryFromString(rawQuery string) RegionalContentGetActionQuery {
	v := RegionalContentGetActionQuery{}
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
func RegionalContentGetActionQueryFromHttp(r *http.Request) RegionalContentGetActionQuery {
	return RegionalContentGetActionQueryFromString(r.URL.RawQuery)
}
func (q RegionalContentGetActionQuery) Values() url.Values {
	return q.values
}
func (q RegionalContentGetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *RegionalContentGetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *RegionalContentGetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type RegionalContentGetActionRequest struct {
	Body        interface{}
	Params      RegionalContentGetActionPathParameter
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
func (x RegionalContentGetActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x RegionalContentGetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func RegionalContentGetActionClientCreateUrl(
	req RegionalContentGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := RegionalContentGetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = RegionalContentGetActionPathParameterApply(req.Params, urlAddr)
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
func RegionalContentGetActionClientExecuteTyped(httpReq *http.Request) (*RegionalContentGetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result RegionalContentGetActionResponse
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
func RegionalContentGetActionClientBuildRequest(req RegionalContentGetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := RegionalContentGetActionMeta()
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
func RegionalContentGetActionCall(
	req RegionalContentGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*RegionalContentGetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := RegionalContentGetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := RegionalContentGetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return RegionalContentGetActionClientExecuteTyped(r)
}
func RegionalContentGetActionPathParameterFromGin(g *gin.Context) RegionalContentGetActionPathParameter {
	return RegionalContentGetActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// RegionalContentGetActionRaw registers a raw Gin route for the RegionalContentGetAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func RegionalContentGetActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := RegionalContentGetActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// RegionalContentGetActionHandler returns the HTTP method, route URL, and a typed Gin handler for the RegionalContentGetAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func RegionalContentGetActionHandler(
	handler func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := RegionalContentGetActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := RegionalContentGetActionRequest{
			Body:        nil,
			Params:      RegionalContentGetActionPathParameterFromGin(m),
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

// RegionalContentGetActionGin is a high-level convenience wrapper around RegionalContentGetActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func RegionalContentGetActionGin(r gin.IRoutes, handler func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error)) {
	method, url, h := RegionalContentGetActionHandler(handler)
	r.Handle(method, url, h)
}
func (x RegionalContentGetActionRequest) IsGin() bool {
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
func RegionalContentGetActionQueryFromGin(c *gin.Context) RegionalContentGetActionQuery {
	return RegionalContentGetActionQueryFromString(c.Request.URL.RawQuery)
}
func GetRegionalContentGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func RegionalContentGetActionPathParameterFromCli(c *cli.Command) RegionalContentGetActionPathParameter {
	return RegionalContentGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x RegionalContentGetActionRequest) IsCli() bool {
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

// RegionalContentGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RegionalContentGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RegionalContentGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRegionalContentGetActionPathParameterCliFlags(""))...)
	return flags
}

// RegionalContentGetActionCliHandler builds a full *cli.Command for the
// RegionalContentGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RegionalContentGetActionRequest the same way
// RegionalContentGetActionHandler (Gin) and RegionalContentGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RegionalContentGetActionCliHandler(
	handler func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error),
) *cli.Command {
	meta := RegionalContentGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RegionalContentGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RegionalContentGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      RegionalContentGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RegionalContentGetActionCli is a high-level convenience wrapper around
// RegionalContentGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RegionalContentGetActionGin
// registers a route on a Gin engine.
func RegionalContentGetActionCli(
	app *cli.Command,
	handler func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error),
) {
	app.Commands = append(app.Commands, RegionalContentGetActionCliHandler(handler))
}

// RegionalContentGetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the RegionalContentGetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *RegionalContentGetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func RegionalContentGetActionHttpHandler(
	handler func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := RegionalContentGetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := RegionalContentGetActionRequest{
			Body: nil,
			Params: RegionalContentGetActionPathParameterFromFn(func(key string) string {
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

// RegionalContentGetActionHttp is a high-level convenience wrapper around
// RegionalContentGetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func RegionalContentGetActionHttp(
	mux *http.ServeMux,
	handler func(c RegionalContentGetActionRequest) (*RegionalContentGetActionResponse, error),
) {
	method, pattern, h := RegionalContentGetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

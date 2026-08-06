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
* Action to communicate with the action TableViewSizingGetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of TableViewSizingGetAction
func TableViewSizingGetAction(c TableViewSizingGetActionRequest) (*TableViewSizingGetActionResponse, error) {
	return &TableViewSizingGetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func TableViewSizingGetActionMeta() struct {
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
		Name:        "TableViewSizingGetAction",
		CliName:     "table-view-sizing-get-action",
		CliShort:    "tableViewSizing-g",
		URL:         "/tableViewSizing/:uniqueId",
		Method:      "GET",
		Description: `Looks up a single "tableViewSizing" row by uniqueId.`,
	}
}

type TableViewSizingGetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *TableViewSizingGetActionResponse) SetContentType(contentType string) *TableViewSizingGetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *TableViewSizingGetActionResponse) AsStream(r io.Reader, contentType string) *TableViewSizingGetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *TableViewSizingGetActionResponse) AsJSON(payload any) *TableViewSizingGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *TableViewSizingGetActionResponse) WithIdeal(payload TableViewSizingDto) *TableViewSizingGetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *TableViewSizingGetActionResponse) AsIdeal() (*TableViewSizingDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res TableViewSizingDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *TableViewSizingGetActionResponse) AsHTML(payload string) *TableViewSizingGetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *TableViewSizingGetActionResponse) AsBytes(payload []byte) *TableViewSizingGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x TableViewSizingGetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x TableViewSizingGetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x TableViewSizingGetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type TableViewSizingGetActionRequestSig = func(c TableViewSizingGetActionRequest) (*TableViewSizingGetActionResponse, error)

/**
 * Path parameters for TableViewSizingGetAction
 */
type TableViewSizingGetActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func TableViewSizingGetActionPathParameterApply(params TableViewSizingGetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func TableViewSizingGetActionPathParameterFromFn(fn func(key string) string) TableViewSizingGetActionPathParameter {
	res := TableViewSizingGetActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for TableViewSizingGetAction
 */
// Query wrapper with private fields
type TableViewSizingGetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func TableViewSizingGetActionQueryFromString(rawQuery string) TableViewSizingGetActionQuery {
	v := TableViewSizingGetActionQuery{}
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
func TableViewSizingGetActionQueryFromHttp(r *http.Request) TableViewSizingGetActionQuery {
	return TableViewSizingGetActionQueryFromString(r.URL.RawQuery)
}
func (q TableViewSizingGetActionQuery) Values() url.Values {
	return q.values
}
func (q TableViewSizingGetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *TableViewSizingGetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *TableViewSizingGetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type TableViewSizingGetActionRequest struct {
	Body        interface{}
	Params      TableViewSizingGetActionPathParameter
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
func (x TableViewSizingGetActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x TableViewSizingGetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func TableViewSizingGetActionClientCreateUrl(
	req TableViewSizingGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := TableViewSizingGetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = TableViewSizingGetActionPathParameterApply(req.Params, urlAddr)
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
func TableViewSizingGetActionClientExecuteTyped(httpReq *http.Request) (*TableViewSizingGetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result TableViewSizingGetActionResponse
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
func TableViewSizingGetActionClientBuildRequest(req TableViewSizingGetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := TableViewSizingGetActionMeta()
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
func TableViewSizingGetActionCall(
	req TableViewSizingGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*TableViewSizingGetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := TableViewSizingGetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := TableViewSizingGetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return TableViewSizingGetActionClientExecuteTyped(r)
}
func TableViewSizingGetActionPathParameterFromGin(g *gin.Context) TableViewSizingGetActionPathParameter {
	return TableViewSizingGetActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// TableViewSizingGetActionRaw registers a raw Gin route for the TableViewSizingGetAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func TableViewSizingGetActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := TableViewSizingGetActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// TableViewSizingGetActionHandler returns the HTTP method, route URL, and a typed Gin handler for the TableViewSizingGetAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func TableViewSizingGetActionHandler(
	handler func(c TableViewSizingGetActionRequest) (*TableViewSizingGetActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := TableViewSizingGetActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := TableViewSizingGetActionRequest{
			Body:        nil,
			Params:      TableViewSizingGetActionPathParameterFromGin(m),
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

// TableViewSizingGetActionGin is a high-level convenience wrapper around TableViewSizingGetActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func TableViewSizingGetActionGin(r gin.IRoutes, handler func(c TableViewSizingGetActionRequest) (*TableViewSizingGetActionResponse, error)) {
	method, url, h := TableViewSizingGetActionHandler(handler)
	r.Handle(method, url, h)
}
func (x TableViewSizingGetActionRequest) IsGin() bool {
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
func TableViewSizingGetActionQueryFromGin(c *gin.Context) TableViewSizingGetActionQuery {
	return TableViewSizingGetActionQueryFromString(c.Request.URL.RawQuery)
}
func GetTableViewSizingGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func TableViewSizingGetActionPathParameterFromCli(c *cli.Command) TableViewSizingGetActionPathParameter {
	return TableViewSizingGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x TableViewSizingGetActionRequest) IsCli() bool {
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

// TableViewSizingGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TableViewSizingGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TableViewSizingGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTableViewSizingGetActionPathParameterCliFlags(""))...)
	return flags
}

// TableViewSizingGetActionCliHandler builds a full *cli.Command for the
// TableViewSizingGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TableViewSizingGetActionRequest the same way
// TableViewSizingGetActionHandler (Gin) and TableViewSizingGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TableViewSizingGetActionCliHandler(
	handler func(c TableViewSizingGetActionRequest) (*TableViewSizingGetActionResponse, error),
) *cli.Command {
	meta := TableViewSizingGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TableViewSizingGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TableViewSizingGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      TableViewSizingGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TableViewSizingGetActionCli is a high-level convenience wrapper around
// TableViewSizingGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TableViewSizingGetActionGin
// registers a route on a Gin engine.
func TableViewSizingGetActionCli(
	app *cli.Command,
	handler func(c TableViewSizingGetActionRequest) (*TableViewSizingGetActionResponse, error),
) {
	app.Commands = append(app.Commands, TableViewSizingGetActionCliHandler(handler))
}

// TableViewSizingGetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the TableViewSizingGetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *TableViewSizingGetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func TableViewSizingGetActionHttpHandler(
	handler func(c TableViewSizingGetActionRequest) (*TableViewSizingGetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := TableViewSizingGetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := TableViewSizingGetActionRequest{
			Body: nil,
			Params: TableViewSizingGetActionPathParameterFromFn(func(key string) string {
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

// TableViewSizingGetActionHttp is a high-level convenience wrapper around
// TableViewSizingGetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func TableViewSizingGetActionHttp(
	mux *http.ServeMux,
	handler func(c TableViewSizingGetActionRequest) (*TableViewSizingGetActionResponse, error),
) {
	method, pattern, h := TableViewSizingGetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

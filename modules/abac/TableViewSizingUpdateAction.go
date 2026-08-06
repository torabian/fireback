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
* Action to communicate with the action TableViewSizingUpdateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of TableViewSizingUpdateAction
func TableViewSizingUpdateAction(c TableViewSizingUpdateActionRequest) (*TableViewSizingUpdateActionResponse, error) {
	return &TableViewSizingUpdateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func TableViewSizingUpdateActionMeta() struct {
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
		Name:        "TableViewSizingUpdateAction",
		CliName:     "table-view-sizing-update-action",
		CliShort:    "tableViewSizing-u",
		URL:         "/tableViewSizing/:uniqueId",
		Method:      "PATCH",
		Description: `Applies a partial update to a "tableViewSizing" row by uniqueId.`,
	}
}

type TableViewSizingUpdateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *TableViewSizingUpdateActionResponse) SetContentType(contentType string) *TableViewSizingUpdateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *TableViewSizingUpdateActionResponse) AsStream(r io.Reader, contentType string) *TableViewSizingUpdateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *TableViewSizingUpdateActionResponse) AsJSON(payload any) *TableViewSizingUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *TableViewSizingUpdateActionResponse) WithIdeal(payload TableViewSizingDto) *TableViewSizingUpdateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *TableViewSizingUpdateActionResponse) AsIdeal() (*TableViewSizingDto, error) {
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
func (x *TableViewSizingUpdateActionResponse) AsHTML(payload string) *TableViewSizingUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *TableViewSizingUpdateActionResponse) AsBytes(payload []byte) *TableViewSizingUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x TableViewSizingUpdateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x TableViewSizingUpdateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x TableViewSizingUpdateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type TableViewSizingUpdateActionRequestSig = func(c TableViewSizingUpdateActionRequest) (*TableViewSizingUpdateActionResponse, error)

/**
 * Path parameters for TableViewSizingUpdateAction
 */
type TableViewSizingUpdateActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func TableViewSizingUpdateActionPathParameterApply(params TableViewSizingUpdateActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func TableViewSizingUpdateActionPathParameterFromFn(fn func(key string) string) TableViewSizingUpdateActionPathParameter {
	res := TableViewSizingUpdateActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for TableViewSizingUpdateAction
 */
// Query wrapper with private fields
type TableViewSizingUpdateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func TableViewSizingUpdateActionQueryFromString(rawQuery string) TableViewSizingUpdateActionQuery {
	v := TableViewSizingUpdateActionQuery{}
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
func TableViewSizingUpdateActionQueryFromHttp(r *http.Request) TableViewSizingUpdateActionQuery {
	return TableViewSizingUpdateActionQueryFromString(r.URL.RawQuery)
}
func (q TableViewSizingUpdateActionQuery) Values() url.Values {
	return q.values
}
func (q TableViewSizingUpdateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *TableViewSizingUpdateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *TableViewSizingUpdateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type TableViewSizingUpdateActionRequest struct {
	Body        TableViewSizingOptionalDto
	Params      TableViewSizingUpdateActionPathParameter
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
func (x TableViewSizingUpdateActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x TableViewSizingUpdateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func TableViewSizingUpdateActionClientCreateUrl(
	req TableViewSizingUpdateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := TableViewSizingUpdateActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = TableViewSizingUpdateActionPathParameterApply(req.Params, urlAddr)
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
func TableViewSizingUpdateActionClientExecuteTyped(httpReq *http.Request) (*TableViewSizingUpdateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result TableViewSizingUpdateActionResponse
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
func TableViewSizingUpdateActionClientBuildRequest(req TableViewSizingUpdateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := TableViewSizingUpdateActionMeta()
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
func TableViewSizingUpdateActionCall(
	req TableViewSizingUpdateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*TableViewSizingUpdateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := TableViewSizingUpdateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := TableViewSizingUpdateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return TableViewSizingUpdateActionClientExecuteTyped(r)
}
func TableViewSizingUpdateActionPathParameterFromGin(g *gin.Context) TableViewSizingUpdateActionPathParameter {
	return TableViewSizingUpdateActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// TableViewSizingUpdateActionRaw registers a raw Gin route for the TableViewSizingUpdateAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func TableViewSizingUpdateActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := TableViewSizingUpdateActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// TableViewSizingUpdateActionHandler returns the HTTP method, route URL, and a typed Gin handler for the TableViewSizingUpdateAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func TableViewSizingUpdateActionHandler(
	handler func(c TableViewSizingUpdateActionRequest) (*TableViewSizingUpdateActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := TableViewSizingUpdateActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body TableViewSizingOptionalDto
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := TableViewSizingUpdateActionRequest{
			Body:        body,
			Params:      TableViewSizingUpdateActionPathParameterFromGin(m),
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

// TableViewSizingUpdateActionGin is a high-level convenience wrapper around TableViewSizingUpdateActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func TableViewSizingUpdateActionGin(r gin.IRoutes, handler func(c TableViewSizingUpdateActionRequest) (*TableViewSizingUpdateActionResponse, error)) {
	method, url, h := TableViewSizingUpdateActionHandler(handler)
	r.Handle(method, url, h)
}
func (x TableViewSizingUpdateActionRequest) IsGin() bool {
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
func TableViewSizingUpdateActionQueryFromGin(c *gin.Context) TableViewSizingUpdateActionQuery {
	return TableViewSizingUpdateActionQueryFromString(c.Request.URL.RawQuery)
}
func GetTableViewSizingUpdateActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func TableViewSizingUpdateActionPathParameterFromCli(c *cli.Command) TableViewSizingUpdateActionPathParameter {
	return TableViewSizingUpdateActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x TableViewSizingUpdateActionRequest) IsCli() bool {
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

// TableViewSizingUpdateActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TableViewSizingUpdateAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TableViewSizingUpdateActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTableViewSizingUpdateActionPathParameterCliFlags(""))...)
	return flags
}

// TableViewSizingUpdateActionCliHandler builds a full *cli.Command for the
// TableViewSizingUpdateAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TableViewSizingUpdateActionRequest the same way
// TableViewSizingUpdateActionHandler (Gin) and TableViewSizingUpdateActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TableViewSizingUpdateActionCliHandler(
	handler func(c TableViewSizingUpdateActionRequest) (*TableViewSizingUpdateActionResponse, error),
) *cli.Command {
	meta := TableViewSizingUpdateActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TableViewSizingUpdateActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TableViewSizingUpdateActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      TableViewSizingUpdateActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TableViewSizingUpdateActionCli is a high-level convenience wrapper around
// TableViewSizingUpdateActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TableViewSizingUpdateActionGin
// registers a route on a Gin engine.
func TableViewSizingUpdateActionCli(
	app *cli.Command,
	handler func(c TableViewSizingUpdateActionRequest) (*TableViewSizingUpdateActionResponse, error),
) {
	app.Commands = append(app.Commands, TableViewSizingUpdateActionCliHandler(handler))
}

// TableViewSizingUpdateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the TableViewSizingUpdateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *TableViewSizingUpdateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func TableViewSizingUpdateActionHttpHandler(
	handler func(c TableViewSizingUpdateActionRequest) (*TableViewSizingUpdateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := TableViewSizingUpdateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body TableViewSizingOptionalDto
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
		req := TableViewSizingUpdateActionRequest{
			Body: body,
			Params: TableViewSizingUpdateActionPathParameterFromFn(func(key string) string {
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

// TableViewSizingUpdateActionHttp is a high-level convenience wrapper around
// TableViewSizingUpdateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func TableViewSizingUpdateActionHttp(
	mux *http.ServeMux,
	handler func(c TableViewSizingUpdateActionRequest) (*TableViewSizingUpdateActionResponse, error),
) {
	method, pattern, h := TableViewSizingUpdateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

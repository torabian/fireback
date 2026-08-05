package fireback

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

/**
* Action to communicate with the action WebPushConfigBrowseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of WebPushConfigBrowseAction
func WebPushConfigBrowseAction(c WebPushConfigBrowseActionRequest) (*WebPushConfigBrowseActionResponse, error) {
	return &WebPushConfigBrowseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func WebPushConfigBrowseActionMeta() struct {
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
		Name:        "WebPushConfigBrowseAction",
		CliName:     "web-push-config-browse-action",
		CliShort:    "webPushConfig-b",
		URL:         "/webPushConfig/browse",
		Method:      "GET",
		Description: `Returns "webPushConfig" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).`,
	}
}

type WebPushConfigBrowseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *WebPushConfigBrowseActionResponse) SetContentType(contentType string) *WebPushConfigBrowseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *WebPushConfigBrowseActionResponse) AsStream(r io.Reader, contentType string) *WebPushConfigBrowseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *WebPushConfigBrowseActionResponse) AsJSON(payload any) *WebPushConfigBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *WebPushConfigBrowseActionResponse) WithIdeal(payload WebPushConfigOptionalDto) *WebPushConfigBrowseActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *WebPushConfigBrowseActionResponse) AsIdeal() (*WebPushConfigOptionalDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res WebPushConfigOptionalDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *WebPushConfigBrowseActionResponse) AsHTML(payload string) *WebPushConfigBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *WebPushConfigBrowseActionResponse) AsBytes(payload []byte) *WebPushConfigBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x WebPushConfigBrowseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x WebPushConfigBrowseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x WebPushConfigBrowseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type WebPushConfigBrowseActionRequestSig = func(c WebPushConfigBrowseActionRequest) (*WebPushConfigBrowseActionResponse, error)

/**
 * Query parameters for WebPushConfigBrowseAction
 */
// Query wrapper with private fields
type WebPushConfigBrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func WebPushConfigBrowseActionQueryFromString(rawQuery string) WebPushConfigBrowseActionQuery {
	v := WebPushConfigBrowseActionQuery{}
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
func WebPushConfigBrowseActionQueryFromHttp(r *http.Request) WebPushConfigBrowseActionQuery {
	return WebPushConfigBrowseActionQueryFromString(r.URL.RawQuery)
}
func (q WebPushConfigBrowseActionQuery) Values() url.Values {
	return q.values
}
func (q WebPushConfigBrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *WebPushConfigBrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *WebPushConfigBrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type WebPushConfigBrowseActionRequest struct {
	Body        interface{}
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
func (x WebPushConfigBrowseActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x WebPushConfigBrowseActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func WebPushConfigBrowseActionClientCreateUrl(
	req WebPushConfigBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := WebPushConfigBrowseActionMeta()
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
func WebPushConfigBrowseActionClientExecuteTyped(httpReq *http.Request) (*WebPushConfigBrowseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result WebPushConfigBrowseActionResponse
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
func WebPushConfigBrowseActionClientBuildRequest(req WebPushConfigBrowseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := WebPushConfigBrowseActionMeta()
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
func WebPushConfigBrowseActionCall(
	req WebPushConfigBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*WebPushConfigBrowseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := WebPushConfigBrowseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := WebPushConfigBrowseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return WebPushConfigBrowseActionClientExecuteTyped(r)
}

// WebPushConfigBrowseActionRaw registers a raw Gin route for the WebPushConfigBrowseAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func WebPushConfigBrowseActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := WebPushConfigBrowseActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// WebPushConfigBrowseActionHandler returns the HTTP method, route URL, and a typed Gin handler for the WebPushConfigBrowseAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func WebPushConfigBrowseActionHandler(
	handler func(c WebPushConfigBrowseActionRequest) (*WebPushConfigBrowseActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := WebPushConfigBrowseActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := WebPushConfigBrowseActionRequest{
			Body:        nil,
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

// WebPushConfigBrowseActionGin is a high-level convenience wrapper around WebPushConfigBrowseActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func WebPushConfigBrowseActionGin(r gin.IRoutes, handler func(c WebPushConfigBrowseActionRequest) (*WebPushConfigBrowseActionResponse, error)) {
	method, url, h := WebPushConfigBrowseActionHandler(handler)
	r.Handle(method, url, h)
}
func (x WebPushConfigBrowseActionRequest) IsGin() bool {
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
func WebPushConfigBrowseActionQueryFromGin(c *gin.Context) WebPushConfigBrowseActionQuery {
	return WebPushConfigBrowseActionQueryFromString(c.Request.URL.RawQuery)
}
func GetWebPushConfigBrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-filter",
			Type: "string",
		},
		{
			Name: prefix + "qs-sort",
			Type: "string",
		},
		{
			Name: prefix + "qs-start-index",
			Type: "int",
		},
		{
			Name: prefix + "qs-items-per-page",
			Type: "int",
		},
		{
			Name: prefix + "qs-cursor",
			Type: "string",
		},
	}
}

// WebPushConfigBrowseActionQueryFromCli extracts and casts query parameters the same way
// WebPushConfigBrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func WebPushConfigBrowseActionQueryFromCli(c *cli.Command) WebPushConfigBrowseActionQuery {
	data := WebPushConfigBrowseActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-filter") {
		data.Filter = c.String("qs-filter")
		values.Set("filter", data.Filter)
	}
	if c.IsSet("qs-sort") {
		data.Sort = c.String("qs-sort")
		values.Set("sort", data.Sort)
	}
	if c.IsSet("qs-start-index") {
		data.StartIndex = int(c.Int64("qs-start-index"))
		values.Set("startIndex", strconv.FormatInt(int64(data.StartIndex), 10))
	}
	if c.IsSet("qs-items-per-page") {
		data.ItemsPerPage = int(c.Int64("qs-items-per-page"))
		values.Set("itemsPerPage", strconv.FormatInt(int64(data.ItemsPerPage), 10))
	}
	if c.IsSet("qs-cursor") {
		data.Cursor = c.String("qs-cursor")
		values.Set("cursor", data.Cursor)
	}
	data.SetValues(values)
	return data
}
func (x WebPushConfigBrowseActionRequest) IsCli() bool {
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

// WebPushConfigBrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the WebPushConfigBrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func WebPushConfigBrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetWebPushConfigBrowseActionQueryCliFlags(""))...)
	return flags
}

// WebPushConfigBrowseActionCliHandler builds a full *cli.Command for the
// WebPushConfigBrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a WebPushConfigBrowseActionRequest the same way
// WebPushConfigBrowseActionHandler (Gin) and WebPushConfigBrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func WebPushConfigBrowseActionCliHandler(
	handler func(c WebPushConfigBrowseActionRequest) (*WebPushConfigBrowseActionResponse, error),
) *cli.Command {
	meta := WebPushConfigBrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: WebPushConfigBrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := WebPushConfigBrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = WebPushConfigBrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// WebPushConfigBrowseActionCli is a high-level convenience wrapper around
// WebPushConfigBrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way WebPushConfigBrowseActionGin
// registers a route on a Gin engine.
func WebPushConfigBrowseActionCli(
	app *cli.Command,
	handler func(c WebPushConfigBrowseActionRequest) (*WebPushConfigBrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, WebPushConfigBrowseActionCliHandler(handler))
}

// WebPushConfigBrowseActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the WebPushConfigBrowseAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *WebPushConfigBrowseActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func WebPushConfigBrowseActionHttpHandler(
	handler func(c WebPushConfigBrowseActionRequest) (*WebPushConfigBrowseActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := WebPushConfigBrowseActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := WebPushConfigBrowseActionRequest{
			Body:        nil,
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

// WebPushConfigBrowseActionHttp is a high-level convenience wrapper around
// WebPushConfigBrowseActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func WebPushConfigBrowseActionHttp(
	mux *http.ServeMux,
	handler func(c WebPushConfigBrowseActionRequest) (*WebPushConfigBrowseActionResponse, error),
) {
	method, pattern, h := WebPushConfigBrowseActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

package interfacetools

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
* Action to communicate with the action AppMenuBrowseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of AppMenuBrowseAction
func AppMenuBrowseAction(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error) {
	return &AppMenuBrowseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func AppMenuBrowseActionMeta() struct {
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
		Name:        "AppMenuBrowseAction",
		CliName:     "app-menu-browse-action",
		CliShort:    "appMenu-b",
		URL:         "/appMenu/browse",
		Method:      "GET",
		Description: `Returns "appMenu" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).`,
	}
}

type AppMenuBrowseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *AppMenuBrowseActionResponse) SetContentType(contentType string) *AppMenuBrowseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *AppMenuBrowseActionResponse) AsStream(r io.Reader, contentType string) *AppMenuBrowseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *AppMenuBrowseActionResponse) AsJSON(payload any) *AppMenuBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *AppMenuBrowseActionResponse) WithIdeal(payload AppMenuOptionalDto) *AppMenuBrowseActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *AppMenuBrowseActionResponse) AsIdeal() (*AppMenuOptionalDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res AppMenuOptionalDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *AppMenuBrowseActionResponse) AsHTML(payload string) *AppMenuBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *AppMenuBrowseActionResponse) AsBytes(payload []byte) *AppMenuBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x AppMenuBrowseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x AppMenuBrowseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x AppMenuBrowseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type AppMenuBrowseActionRequestSig = func(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error)

/**
 * Query parameters for AppMenuBrowseAction
 */
// Query wrapper with private fields
type AppMenuBrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func AppMenuBrowseActionQueryFromString(rawQuery string) AppMenuBrowseActionQuery {
	v := AppMenuBrowseActionQuery{}
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
func AppMenuBrowseActionQueryFromHttp(r *http.Request) AppMenuBrowseActionQuery {
	return AppMenuBrowseActionQueryFromString(r.URL.RawQuery)
}
func (q AppMenuBrowseActionQuery) Values() url.Values {
	return q.values
}
func (q AppMenuBrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *AppMenuBrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *AppMenuBrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type AppMenuBrowseActionRequest struct {
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
func (x AppMenuBrowseActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x AppMenuBrowseActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func AppMenuBrowseActionClientCreateUrl(
	req AppMenuBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := AppMenuBrowseActionMeta()
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
func AppMenuBrowseActionClientExecuteTyped(httpReq *http.Request) (*AppMenuBrowseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result AppMenuBrowseActionResponse
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
func AppMenuBrowseActionClientBuildRequest(req AppMenuBrowseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := AppMenuBrowseActionMeta()
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
func AppMenuBrowseActionCall(
	req AppMenuBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*AppMenuBrowseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := AppMenuBrowseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := AppMenuBrowseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return AppMenuBrowseActionClientExecuteTyped(r)
}

// AppMenuBrowseActionRaw registers a raw Gin route for the AppMenuBrowseAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func AppMenuBrowseActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := AppMenuBrowseActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// AppMenuBrowseActionHandler returns the HTTP method, route URL, and a typed Gin handler for the AppMenuBrowseAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func AppMenuBrowseActionHandler(
	handler func(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := AppMenuBrowseActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := AppMenuBrowseActionRequest{
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

// AppMenuBrowseActionGin is a high-level convenience wrapper around AppMenuBrowseActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func AppMenuBrowseActionGin(r gin.IRoutes, handler func(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error)) {
	method, url, h := AppMenuBrowseActionHandler(handler)
	r.Handle(method, url, h)
}
func (x AppMenuBrowseActionRequest) IsGin() bool {
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
func AppMenuBrowseActionQueryFromGin(c *gin.Context) AppMenuBrowseActionQuery {
	return AppMenuBrowseActionQueryFromString(c.Request.URL.RawQuery)
}
func GetAppMenuBrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
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

// AppMenuBrowseActionQueryFromCli extracts and casts query parameters the same way
// AppMenuBrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func AppMenuBrowseActionQueryFromCli(c *cli.Command) AppMenuBrowseActionQuery {
	data := AppMenuBrowseActionQuery{}
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
func (x AppMenuBrowseActionRequest) IsCli() bool {
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

// AppMenuBrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AppMenuBrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AppMenuBrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAppMenuBrowseActionQueryCliFlags(""))...)
	return flags
}

// AppMenuBrowseActionCliHandler builds a full *cli.Command for the
// AppMenuBrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AppMenuBrowseActionRequest the same way
// AppMenuBrowseActionHandler (Gin) and AppMenuBrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AppMenuBrowseActionCliHandler(
	handler func(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error),
) *cli.Command {
	meta := AppMenuBrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AppMenuBrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AppMenuBrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = AppMenuBrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AppMenuBrowseActionCli is a high-level convenience wrapper around
// AppMenuBrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AppMenuBrowseActionGin
// registers a route on a Gin engine.
func AppMenuBrowseActionCli(
	app *cli.Command,
	handler func(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, AppMenuBrowseActionCliHandler(handler))
}

// AppMenuBrowseActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the AppMenuBrowseAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *AppMenuBrowseActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func AppMenuBrowseActionHttpHandler(
	handler func(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := AppMenuBrowseActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := AppMenuBrowseActionRequest{
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

// AppMenuBrowseActionHttp is a high-level convenience wrapper around
// AppMenuBrowseActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func AppMenuBrowseActionHttp(
	mux *http.ServeMux,
	handler func(c AppMenuBrowseActionRequest) (*AppMenuBrowseActionResponse, error),
) {
	method, pattern, h := AppMenuBrowseActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

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
* Action to communicate with the action CapabilityBrowseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CapabilityBrowseAction
func CapabilityBrowseAction(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error) {
	return &CapabilityBrowseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CapabilityBrowseActionMeta() struct {
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
		Name:        "CapabilityBrowseAction",
		CliName:     "capability-browse-action",
		CliShort:    "capability-b",
		URL:         "/capability/browse",
		Method:      "GET",
		Description: `Returns "capability" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).`,
	}
}

type CapabilityBrowseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CapabilityBrowseActionResponse) SetContentType(contentType string) *CapabilityBrowseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CapabilityBrowseActionResponse) AsStream(r io.Reader, contentType string) *CapabilityBrowseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CapabilityBrowseActionResponse) AsJSON(payload any) *CapabilityBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CapabilityBrowseActionResponse) WithIdeal(payload CapabilityOptionalDto) *CapabilityBrowseActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *CapabilityBrowseActionResponse) AsIdeal() (*CapabilityOptionalDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res CapabilityOptionalDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *CapabilityBrowseActionResponse) AsHTML(payload string) *CapabilityBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CapabilityBrowseActionResponse) AsBytes(payload []byte) *CapabilityBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CapabilityBrowseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CapabilityBrowseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CapabilityBrowseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type CapabilityBrowseActionRequestSig = func(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error)

/**
 * Query parameters for CapabilityBrowseAction
 */
// Query wrapper with private fields
type CapabilityBrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func CapabilityBrowseActionQueryFromString(rawQuery string) CapabilityBrowseActionQuery {
	v := CapabilityBrowseActionQuery{}
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
func CapabilityBrowseActionQueryFromHttp(r *http.Request) CapabilityBrowseActionQuery {
	return CapabilityBrowseActionQueryFromString(r.URL.RawQuery)
}
func (q CapabilityBrowseActionQuery) Values() url.Values {
	return q.values
}
func (q CapabilityBrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CapabilityBrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CapabilityBrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CapabilityBrowseActionRequest struct {
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
func (x CapabilityBrowseActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x CapabilityBrowseActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func CapabilityBrowseActionClientCreateUrl(
	req CapabilityBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CapabilityBrowseActionMeta()
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
func CapabilityBrowseActionClientExecuteTyped(httpReq *http.Request) (*CapabilityBrowseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CapabilityBrowseActionResponse
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
func CapabilityBrowseActionClientBuildRequest(req CapabilityBrowseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CapabilityBrowseActionMeta()
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
func CapabilityBrowseActionCall(
	req CapabilityBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CapabilityBrowseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CapabilityBrowseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CapabilityBrowseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CapabilityBrowseActionClientExecuteTyped(r)
}

// CapabilityBrowseActionRaw registers a raw Gin route for the CapabilityBrowseAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CapabilityBrowseActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CapabilityBrowseActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// CapabilityBrowseActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CapabilityBrowseAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CapabilityBrowseActionHandler(
	handler func(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := CapabilityBrowseActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := CapabilityBrowseActionRequest{
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

// CapabilityBrowseActionGin is a high-level convenience wrapper around CapabilityBrowseActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CapabilityBrowseActionGin(r gin.IRoutes, handler func(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error)) {
	method, url, h := CapabilityBrowseActionHandler(handler)
	r.Handle(method, url, h)
}
func (x CapabilityBrowseActionRequest) IsGin() bool {
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
func CapabilityBrowseActionQueryFromGin(c *gin.Context) CapabilityBrowseActionQuery {
	return CapabilityBrowseActionQueryFromString(c.Request.URL.RawQuery)
}
func GetCapabilityBrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
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

// CapabilityBrowseActionQueryFromCli extracts and casts query parameters the same way
// CapabilityBrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func CapabilityBrowseActionQueryFromCli(c *cli.Command) CapabilityBrowseActionQuery {
	data := CapabilityBrowseActionQuery{}
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
func (x CapabilityBrowseActionRequest) IsCli() bool {
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

// CapabilityBrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the CapabilityBrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func CapabilityBrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetCapabilityBrowseActionQueryCliFlags(""))...)
	return flags
}

// CapabilityBrowseActionCliHandler builds a full *cli.Command for the
// CapabilityBrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a CapabilityBrowseActionRequest the same way
// CapabilityBrowseActionHandler (Gin) and CapabilityBrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func CapabilityBrowseActionCliHandler(
	handler func(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error),
) *cli.Command {
	meta := CapabilityBrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: CapabilityBrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := CapabilityBrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = CapabilityBrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// CapabilityBrowseActionCli is a high-level convenience wrapper around
// CapabilityBrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way CapabilityBrowseActionGin
// registers a route on a Gin engine.
func CapabilityBrowseActionCli(
	app *cli.Command,
	handler func(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, CapabilityBrowseActionCliHandler(handler))
}

// CapabilityBrowseActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the CapabilityBrowseAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *CapabilityBrowseActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func CapabilityBrowseActionHttpHandler(
	handler func(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := CapabilityBrowseActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := CapabilityBrowseActionRequest{
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

// CapabilityBrowseActionHttp is a high-level convenience wrapper around
// CapabilityBrowseActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func CapabilityBrowseActionHttp(
	mux *http.ServeMux,
	handler func(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error),
) {
	method, pattern, h := CapabilityBrowseActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

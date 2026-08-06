package abac

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
* Action to communicate with the action PublicJoinKeyBrowseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PublicJoinKeyBrowseAction
func PublicJoinKeyBrowseAction(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error) {
	return &PublicJoinKeyBrowseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PublicJoinKeyBrowseActionMeta() struct {
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
		Name:        "PublicJoinKeyBrowseAction",
		CliName:     "public-join-key-browse-action",
		CliShort:    "publicJoinKey-b",
		URL:         "/publicJoinKey/browse",
		Method:      "GET",
		Description: `Returns "publicJoinKey" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).`,
	}
}

type PublicJoinKeyBrowseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PublicJoinKeyBrowseActionResponse) SetContentType(contentType string) *PublicJoinKeyBrowseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PublicJoinKeyBrowseActionResponse) AsStream(r io.Reader, contentType string) *PublicJoinKeyBrowseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PublicJoinKeyBrowseActionResponse) AsJSON(payload any) *PublicJoinKeyBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *PublicJoinKeyBrowseActionResponse) WithIdeal(payload PublicJoinKeyOptionalDto) *PublicJoinKeyBrowseActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *PublicJoinKeyBrowseActionResponse) AsIdeal() (*PublicJoinKeyOptionalDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res PublicJoinKeyOptionalDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *PublicJoinKeyBrowseActionResponse) AsHTML(payload string) *PublicJoinKeyBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PublicJoinKeyBrowseActionResponse) AsBytes(payload []byte) *PublicJoinKeyBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PublicJoinKeyBrowseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PublicJoinKeyBrowseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PublicJoinKeyBrowseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type PublicJoinKeyBrowseActionRequestSig = func(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error)

/**
 * Query parameters for PublicJoinKeyBrowseAction
 */
// Query wrapper with private fields
type PublicJoinKeyBrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func PublicJoinKeyBrowseActionQueryFromString(rawQuery string) PublicJoinKeyBrowseActionQuery {
	v := PublicJoinKeyBrowseActionQuery{}
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
func PublicJoinKeyBrowseActionQueryFromHttp(r *http.Request) PublicJoinKeyBrowseActionQuery {
	return PublicJoinKeyBrowseActionQueryFromString(r.URL.RawQuery)
}
func (q PublicJoinKeyBrowseActionQuery) Values() url.Values {
	return q.values
}
func (q PublicJoinKeyBrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PublicJoinKeyBrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PublicJoinKeyBrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PublicJoinKeyBrowseActionRequest struct {
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
func (x PublicJoinKeyBrowseActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x PublicJoinKeyBrowseActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func PublicJoinKeyBrowseActionClientCreateUrl(
	req PublicJoinKeyBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PublicJoinKeyBrowseActionMeta()
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
func PublicJoinKeyBrowseActionClientExecuteTyped(httpReq *http.Request) (*PublicJoinKeyBrowseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PublicJoinKeyBrowseActionResponse
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
func PublicJoinKeyBrowseActionClientBuildRequest(req PublicJoinKeyBrowseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PublicJoinKeyBrowseActionMeta()
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
func PublicJoinKeyBrowseActionCall(
	req PublicJoinKeyBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PublicJoinKeyBrowseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PublicJoinKeyBrowseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PublicJoinKeyBrowseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PublicJoinKeyBrowseActionClientExecuteTyped(r)
}

// PublicJoinKeyBrowseActionRaw registers a raw Gin route for the PublicJoinKeyBrowseAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PublicJoinKeyBrowseActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PublicJoinKeyBrowseActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// PublicJoinKeyBrowseActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PublicJoinKeyBrowseAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PublicJoinKeyBrowseActionHandler(
	handler func(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PublicJoinKeyBrowseActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PublicJoinKeyBrowseActionRequest{
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

// PublicJoinKeyBrowseActionGin is a high-level convenience wrapper around PublicJoinKeyBrowseActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PublicJoinKeyBrowseActionGin(r gin.IRoutes, handler func(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error)) {
	method, url, h := PublicJoinKeyBrowseActionHandler(handler)
	r.Handle(method, url, h)
}
func (x PublicJoinKeyBrowseActionRequest) IsGin() bool {
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
func PublicJoinKeyBrowseActionQueryFromGin(c *gin.Context) PublicJoinKeyBrowseActionQuery {
	return PublicJoinKeyBrowseActionQueryFromString(c.Request.URL.RawQuery)
}
func GetPublicJoinKeyBrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
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

// PublicJoinKeyBrowseActionQueryFromCli extracts and casts query parameters the same way
// PublicJoinKeyBrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func PublicJoinKeyBrowseActionQueryFromCli(c *cli.Command) PublicJoinKeyBrowseActionQuery {
	data := PublicJoinKeyBrowseActionQuery{}
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
func (x PublicJoinKeyBrowseActionRequest) IsCli() bool {
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

// PublicJoinKeyBrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PublicJoinKeyBrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PublicJoinKeyBrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPublicJoinKeyBrowseActionQueryCliFlags(""))...)
	return flags
}

// PublicJoinKeyBrowseActionCliHandler builds a full *cli.Command for the
// PublicJoinKeyBrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PublicJoinKeyBrowseActionRequest the same way
// PublicJoinKeyBrowseActionHandler (Gin) and PublicJoinKeyBrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PublicJoinKeyBrowseActionCliHandler(
	handler func(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error),
) *cli.Command {
	meta := PublicJoinKeyBrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PublicJoinKeyBrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PublicJoinKeyBrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = PublicJoinKeyBrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PublicJoinKeyBrowseActionCli is a high-level convenience wrapper around
// PublicJoinKeyBrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PublicJoinKeyBrowseActionGin
// registers a route on a Gin engine.
func PublicJoinKeyBrowseActionCli(
	app *cli.Command,
	handler func(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, PublicJoinKeyBrowseActionCliHandler(handler))
}

// PublicJoinKeyBrowseActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the PublicJoinKeyBrowseAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *PublicJoinKeyBrowseActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func PublicJoinKeyBrowseActionHttpHandler(
	handler func(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := PublicJoinKeyBrowseActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := PublicJoinKeyBrowseActionRequest{
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

// PublicJoinKeyBrowseActionHttp is a high-level convenience wrapper around
// PublicJoinKeyBrowseActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func PublicJoinKeyBrowseActionHttp(
	mux *http.ServeMux,
	handler func(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error),
) {
	method, pattern, h := PublicJoinKeyBrowseActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

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
)

/**
* Action to communicate with the action GsmProviderAwareDeletePreviewAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GsmProviderAwareDeletePreviewAction
func GsmProviderAwareDeletePreviewAction(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error) {
	return &GsmProviderAwareDeletePreviewActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GsmProviderAwareDeletePreviewActionMeta() struct {
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
		Name:        "GsmProviderAwareDeletePreviewAction",
		CliName:     "gsm-provider-aware-delete-preview-action",
		CliShort:    "gsmProvider-dp",
		URL:         "/gsmProvider/delete-preview",
		Method:      "GET",
		Description: `Reports what deleting the given "gsmProvider" uniqueIds would affect, without deleting anything.`,
	}
}

// The base class definition for gsmProviderAwareDeletePreviewActionRes
type GsmProviderAwareDeletePreviewActionRes struct {
	Message  string                                                      `json:"message" yaml:"message"`
	Affected emigo.Array[GsmProviderAwareDeletePreviewActionResAffected] `json:"affected" yaml:"affected"`
}

// The base class definition for affected
type GsmProviderAwareDeletePreviewActionResAffected struct {
	Relation string `json:"relation" yaml:"relation"`
	Count    int64  `json:"count" yaml:"count"`
}

func (x *GsmProviderAwareDeletePreviewActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetGsmProviderAwareDeletePreviewActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "message",
			Type: "string",
		},
		{
			Name: prefix + "affected",
			Type: "array",
		},
	}
}
func CastGsmProviderAwareDeletePreviewActionResFromCli(c emigo.CliCastable) GsmProviderAwareDeletePreviewActionRes {
	data := GsmProviderAwareDeletePreviewActionRes{}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("affected") {
		data.Affected = emigo.CapturePossibleArray(CastGsmProviderAwareDeletePreviewActionResAffectedFromCli, "affected", c)
	}
	return data
}
func GetGsmProviderAwareDeletePreviewActionResAffectedCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "relation",
			Type: "string",
		},
		{
			Name: prefix + "count",
			Type: "int64",
		},
	}
}
func CastGsmProviderAwareDeletePreviewActionResAffectedFromCli(c emigo.CliCastable) GsmProviderAwareDeletePreviewActionResAffected {
	data := GsmProviderAwareDeletePreviewActionResAffected{}
	if c.IsSet("relation") {
		data.Relation = c.String("relation")
	}
	if c.IsSet("count") {
		data.Count = int64(c.Int64("count"))
	}
	return data
}

type GsmProviderAwareDeletePreviewActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GsmProviderAwareDeletePreviewActionResponse) SetContentType(contentType string) *GsmProviderAwareDeletePreviewActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GsmProviderAwareDeletePreviewActionResponse) AsStream(r io.Reader, contentType string) *GsmProviderAwareDeletePreviewActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GsmProviderAwareDeletePreviewActionResponse) AsJSON(payload any) *GsmProviderAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GsmProviderAwareDeletePreviewActionResponse) WithIdeal(payload GsmProviderAwareDeletePreviewActionRes) *GsmProviderAwareDeletePreviewActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *GsmProviderAwareDeletePreviewActionResponse) AsIdeal() (*GsmProviderAwareDeletePreviewActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res GsmProviderAwareDeletePreviewActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *GsmProviderAwareDeletePreviewActionResponse) AsHTML(payload string) *GsmProviderAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GsmProviderAwareDeletePreviewActionResponse) AsBytes(payload []byte) *GsmProviderAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GsmProviderAwareDeletePreviewActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GsmProviderAwareDeletePreviewActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GsmProviderAwareDeletePreviewActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type GsmProviderAwareDeletePreviewActionRequestSig = func(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error)

/**
 * Query parameters for GsmProviderAwareDeletePreviewAction
 */
// Query wrapper with private fields
type GsmProviderAwareDeletePreviewActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	UniqueIds []string `json:"uniqueIds"`
}

func GsmProviderAwareDeletePreviewActionQueryFromString(rawQuery string) GsmProviderAwareDeletePreviewActionQuery {
	v := GsmProviderAwareDeletePreviewActionQuery{}
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
func GsmProviderAwareDeletePreviewActionQueryFromHttp(r *http.Request) GsmProviderAwareDeletePreviewActionQuery {
	return GsmProviderAwareDeletePreviewActionQueryFromString(r.URL.RawQuery)
}
func (q GsmProviderAwareDeletePreviewActionQuery) Values() url.Values {
	return q.values
}
func (q GsmProviderAwareDeletePreviewActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GsmProviderAwareDeletePreviewActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GsmProviderAwareDeletePreviewActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GsmProviderAwareDeletePreviewActionRequest struct {
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
func (x GsmProviderAwareDeletePreviewActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x GsmProviderAwareDeletePreviewActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func GsmProviderAwareDeletePreviewActionClientCreateUrl(
	req GsmProviderAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GsmProviderAwareDeletePreviewActionMeta()
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
func GsmProviderAwareDeletePreviewActionClientExecuteTyped(httpReq *http.Request) (*GsmProviderAwareDeletePreviewActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GsmProviderAwareDeletePreviewActionResponse
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
func GsmProviderAwareDeletePreviewActionClientBuildRequest(req GsmProviderAwareDeletePreviewActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GsmProviderAwareDeletePreviewActionMeta()
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
func GsmProviderAwareDeletePreviewActionCall(
	req GsmProviderAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GsmProviderAwareDeletePreviewActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GsmProviderAwareDeletePreviewActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GsmProviderAwareDeletePreviewActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GsmProviderAwareDeletePreviewActionClientExecuteTyped(r)
}

// GsmProviderAwareDeletePreviewActionRaw registers a raw Gin route for the GsmProviderAwareDeletePreviewAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GsmProviderAwareDeletePreviewActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GsmProviderAwareDeletePreviewActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// GsmProviderAwareDeletePreviewActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GsmProviderAwareDeletePreviewAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GsmProviderAwareDeletePreviewActionHandler(
	handler func(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GsmProviderAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GsmProviderAwareDeletePreviewActionRequest{
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

// GsmProviderAwareDeletePreviewActionGin is a high-level convenience wrapper around GsmProviderAwareDeletePreviewActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GsmProviderAwareDeletePreviewActionGin(r gin.IRoutes, handler func(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error)) {
	method, url, h := GsmProviderAwareDeletePreviewActionHandler(handler)
	r.Handle(method, url, h)
}
func (x GsmProviderAwareDeletePreviewActionRequest) IsGin() bool {
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
func GsmProviderAwareDeletePreviewActionQueryFromGin(c *gin.Context) GsmProviderAwareDeletePreviewActionQuery {
	return GsmProviderAwareDeletePreviewActionQueryFromString(c.Request.URL.RawQuery)
}
func GetGsmProviderAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// GsmProviderAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// GsmProviderAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func GsmProviderAwareDeletePreviewActionQueryFromCli(c *cli.Command) GsmProviderAwareDeletePreviewActionQuery {
	data := GsmProviderAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x GsmProviderAwareDeletePreviewActionRequest) IsCli() bool {
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

// GsmProviderAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the GsmProviderAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func GsmProviderAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetGsmProviderAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// GsmProviderAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// GsmProviderAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a GsmProviderAwareDeletePreviewActionRequest the same way
// GsmProviderAwareDeletePreviewActionHandler (Gin) and GsmProviderAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func GsmProviderAwareDeletePreviewActionCliHandler(
	handler func(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := GsmProviderAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: GsmProviderAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := GsmProviderAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = GsmProviderAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// GsmProviderAwareDeletePreviewActionCli is a high-level convenience wrapper around
// GsmProviderAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way GsmProviderAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func GsmProviderAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, GsmProviderAwareDeletePreviewActionCliHandler(handler))
}

// GsmProviderAwareDeletePreviewActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the GsmProviderAwareDeletePreviewAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *GsmProviderAwareDeletePreviewActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func GsmProviderAwareDeletePreviewActionHttpHandler(
	handler func(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := GsmProviderAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := GsmProviderAwareDeletePreviewActionRequest{
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

// GsmProviderAwareDeletePreviewActionHttp is a high-level convenience wrapper around
// GsmProviderAwareDeletePreviewActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func GsmProviderAwareDeletePreviewActionHttp(
	mux *http.ServeMux,
	handler func(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error),
) {
	method, pattern, h := GsmProviderAwareDeletePreviewActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

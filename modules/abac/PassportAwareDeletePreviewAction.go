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
* Action to communicate with the action PassportAwareDeletePreviewAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PassportAwareDeletePreviewAction
func PassportAwareDeletePreviewAction(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error) {
	return &PassportAwareDeletePreviewActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PassportAwareDeletePreviewActionMeta() struct {
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
		Name:        "PassportAwareDeletePreviewAction",
		CliName:     "passport-aware-delete-preview-action",
		CliShort:    "passport-dp",
		URL:         "/passport/delete-preview",
		Method:      "GET",
		Description: `Reports what deleting the given "passport" uniqueIds would affect, without deleting anything.`,
	}
}

// The base class definition for passportAwareDeletePreviewActionRes
type PassportAwareDeletePreviewActionRes struct {
	Message  string                                                   `json:"message" yaml:"message"`
	Affected emigo.Array[PassportAwareDeletePreviewActionResAffected] `json:"affected" yaml:"affected"`
}

// The base class definition for affected
type PassportAwareDeletePreviewActionResAffected struct {
	Relation string `json:"relation" yaml:"relation"`
	Count    int64  `json:"count" yaml:"count"`
}

func (x *PassportAwareDeletePreviewActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPassportAwareDeletePreviewActionResCliFlags(prefix string) []emigo.CliFlag {
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
func CastPassportAwareDeletePreviewActionResFromCli(c emigo.CliCastable) PassportAwareDeletePreviewActionRes {
	data := PassportAwareDeletePreviewActionRes{}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("affected") {
		data.Affected = emigo.CapturePossibleArray(CastPassportAwareDeletePreviewActionResAffectedFromCli, "affected", c)
	}
	return data
}
func GetPassportAwareDeletePreviewActionResAffectedCliFlags(prefix string) []emigo.CliFlag {
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
func CastPassportAwareDeletePreviewActionResAffectedFromCli(c emigo.CliCastable) PassportAwareDeletePreviewActionResAffected {
	data := PassportAwareDeletePreviewActionResAffected{}
	if c.IsSet("relation") {
		data.Relation = c.String("relation")
	}
	if c.IsSet("count") {
		data.Count = int64(c.Int64("count"))
	}
	return data
}

type PassportAwareDeletePreviewActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PassportAwareDeletePreviewActionResponse) SetContentType(contentType string) *PassportAwareDeletePreviewActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PassportAwareDeletePreviewActionResponse) AsStream(r io.Reader, contentType string) *PassportAwareDeletePreviewActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PassportAwareDeletePreviewActionResponse) AsJSON(payload any) *PassportAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *PassportAwareDeletePreviewActionResponse) WithIdeal(payload PassportAwareDeletePreviewActionRes) *PassportAwareDeletePreviewActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *PassportAwareDeletePreviewActionResponse) AsIdeal() (*PassportAwareDeletePreviewActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res PassportAwareDeletePreviewActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *PassportAwareDeletePreviewActionResponse) AsHTML(payload string) *PassportAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PassportAwareDeletePreviewActionResponse) AsBytes(payload []byte) *PassportAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PassportAwareDeletePreviewActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PassportAwareDeletePreviewActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PassportAwareDeletePreviewActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type PassportAwareDeletePreviewActionRequestSig = func(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error)

/**
 * Query parameters for PassportAwareDeletePreviewAction
 */
// Query wrapper with private fields
type PassportAwareDeletePreviewActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	UniqueIds []string `json:"uniqueIds"`
}

func PassportAwareDeletePreviewActionQueryFromString(rawQuery string) PassportAwareDeletePreviewActionQuery {
	v := PassportAwareDeletePreviewActionQuery{}
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
func PassportAwareDeletePreviewActionQueryFromHttp(r *http.Request) PassportAwareDeletePreviewActionQuery {
	return PassportAwareDeletePreviewActionQueryFromString(r.URL.RawQuery)
}
func (q PassportAwareDeletePreviewActionQuery) Values() url.Values {
	return q.values
}
func (q PassportAwareDeletePreviewActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PassportAwareDeletePreviewActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PassportAwareDeletePreviewActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PassportAwareDeletePreviewActionRequest struct {
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
func (x PassportAwareDeletePreviewActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x PassportAwareDeletePreviewActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func PassportAwareDeletePreviewActionClientCreateUrl(
	req PassportAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PassportAwareDeletePreviewActionMeta()
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
func PassportAwareDeletePreviewActionClientExecuteTyped(httpReq *http.Request) (*PassportAwareDeletePreviewActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PassportAwareDeletePreviewActionResponse
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
func PassportAwareDeletePreviewActionClientBuildRequest(req PassportAwareDeletePreviewActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PassportAwareDeletePreviewActionMeta()
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
func PassportAwareDeletePreviewActionCall(
	req PassportAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PassportAwareDeletePreviewActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PassportAwareDeletePreviewActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PassportAwareDeletePreviewActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PassportAwareDeletePreviewActionClientExecuteTyped(r)
}

// PassportAwareDeletePreviewActionRaw registers a raw Gin route for the PassportAwareDeletePreviewAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PassportAwareDeletePreviewActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PassportAwareDeletePreviewActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// PassportAwareDeletePreviewActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PassportAwareDeletePreviewAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PassportAwareDeletePreviewActionHandler(
	handler func(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PassportAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PassportAwareDeletePreviewActionRequest{
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

// PassportAwareDeletePreviewActionGin is a high-level convenience wrapper around PassportAwareDeletePreviewActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PassportAwareDeletePreviewActionGin(r gin.IRoutes, handler func(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error)) {
	method, url, h := PassportAwareDeletePreviewActionHandler(handler)
	r.Handle(method, url, h)
}
func (x PassportAwareDeletePreviewActionRequest) IsGin() bool {
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
func PassportAwareDeletePreviewActionQueryFromGin(c *gin.Context) PassportAwareDeletePreviewActionQuery {
	return PassportAwareDeletePreviewActionQueryFromString(c.Request.URL.RawQuery)
}
func GetPassportAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// PassportAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// PassportAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func PassportAwareDeletePreviewActionQueryFromCli(c *cli.Command) PassportAwareDeletePreviewActionQuery {
	data := PassportAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x PassportAwareDeletePreviewActionRequest) IsCli() bool {
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

// PassportAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PassportAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PassportAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPassportAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// PassportAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// PassportAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PassportAwareDeletePreviewActionRequest the same way
// PassportAwareDeletePreviewActionHandler (Gin) and PassportAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PassportAwareDeletePreviewActionCliHandler(
	handler func(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := PassportAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PassportAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PassportAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = PassportAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PassportAwareDeletePreviewActionCli is a high-level convenience wrapper around
// PassportAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PassportAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func PassportAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, PassportAwareDeletePreviewActionCliHandler(handler))
}

// PassportAwareDeletePreviewActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the PassportAwareDeletePreviewAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *PassportAwareDeletePreviewActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func PassportAwareDeletePreviewActionHttpHandler(
	handler func(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := PassportAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := PassportAwareDeletePreviewActionRequest{
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

// PassportAwareDeletePreviewActionHttp is a high-level convenience wrapper around
// PassportAwareDeletePreviewActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func PassportAwareDeletePreviewActionHttp(
	mux *http.ServeMux,
	handler func(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error),
) {
	method, pattern, h := PassportAwareDeletePreviewActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

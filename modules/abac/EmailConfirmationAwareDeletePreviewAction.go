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
* Action to communicate with the action EmailConfirmationAwareDeletePreviewAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of EmailConfirmationAwareDeletePreviewAction
func EmailConfirmationAwareDeletePreviewAction(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error) {
	return &EmailConfirmationAwareDeletePreviewActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func EmailConfirmationAwareDeletePreviewActionMeta() struct {
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
		Name:        "EmailConfirmationAwareDeletePreviewAction",
		CliName:     "email-confirmation-aware-delete-preview-action",
		CliShort:    "emailConfirmation-dp",
		URL:         "/emailConfirmation/delete-preview",
		Method:      "GET",
		Description: `Reports what deleting the given "emailConfirmation" uniqueIds would affect, without deleting anything.`,
	}
}

// The base class definition for emailConfirmationAwareDeletePreviewActionRes
type EmailConfirmationAwareDeletePreviewActionRes struct {
	Message  string                                                            `json:"message" yaml:"message"`
	Affected emigo.Array[EmailConfirmationAwareDeletePreviewActionResAffected] `json:"affected" yaml:"affected"`
}

// The base class definition for affected
type EmailConfirmationAwareDeletePreviewActionResAffected struct {
	Relation string `json:"relation" yaml:"relation"`
	Count    int64  `json:"count" yaml:"count"`
}

func (x *EmailConfirmationAwareDeletePreviewActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEmailConfirmationAwareDeletePreviewActionResCliFlags(prefix string) []emigo.CliFlag {
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
func CastEmailConfirmationAwareDeletePreviewActionResFromCli(c emigo.CliCastable) EmailConfirmationAwareDeletePreviewActionRes {
	data := EmailConfirmationAwareDeletePreviewActionRes{}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("affected") {
		data.Affected = emigo.CapturePossibleArray(CastEmailConfirmationAwareDeletePreviewActionResAffectedFromCli, "affected", c)
	}
	return data
}
func GetEmailConfirmationAwareDeletePreviewActionResAffectedCliFlags(prefix string) []emigo.CliFlag {
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
func CastEmailConfirmationAwareDeletePreviewActionResAffectedFromCli(c emigo.CliCastable) EmailConfirmationAwareDeletePreviewActionResAffected {
	data := EmailConfirmationAwareDeletePreviewActionResAffected{}
	if c.IsSet("relation") {
		data.Relation = c.String("relation")
	}
	if c.IsSet("count") {
		data.Count = int64(c.Int64("count"))
	}
	return data
}

type EmailConfirmationAwareDeletePreviewActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *EmailConfirmationAwareDeletePreviewActionResponse) SetContentType(contentType string) *EmailConfirmationAwareDeletePreviewActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *EmailConfirmationAwareDeletePreviewActionResponse) AsStream(r io.Reader, contentType string) *EmailConfirmationAwareDeletePreviewActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *EmailConfirmationAwareDeletePreviewActionResponse) AsJSON(payload any) *EmailConfirmationAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *EmailConfirmationAwareDeletePreviewActionResponse) WithIdeal(payload EmailConfirmationAwareDeletePreviewActionRes) *EmailConfirmationAwareDeletePreviewActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *EmailConfirmationAwareDeletePreviewActionResponse) AsIdeal() (*EmailConfirmationAwareDeletePreviewActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res EmailConfirmationAwareDeletePreviewActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *EmailConfirmationAwareDeletePreviewActionResponse) AsHTML(payload string) *EmailConfirmationAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *EmailConfirmationAwareDeletePreviewActionResponse) AsBytes(payload []byte) *EmailConfirmationAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x EmailConfirmationAwareDeletePreviewActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x EmailConfirmationAwareDeletePreviewActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x EmailConfirmationAwareDeletePreviewActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type EmailConfirmationAwareDeletePreviewActionRequestSig = func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error)

/**
 * Query parameters for EmailConfirmationAwareDeletePreviewAction
 */
// Query wrapper with private fields
type EmailConfirmationAwareDeletePreviewActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	UniqueIds []string `json:"uniqueIds"`
}

func EmailConfirmationAwareDeletePreviewActionQueryFromString(rawQuery string) EmailConfirmationAwareDeletePreviewActionQuery {
	v := EmailConfirmationAwareDeletePreviewActionQuery{}
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
func EmailConfirmationAwareDeletePreviewActionQueryFromHttp(r *http.Request) EmailConfirmationAwareDeletePreviewActionQuery {
	return EmailConfirmationAwareDeletePreviewActionQueryFromString(r.URL.RawQuery)
}
func (q EmailConfirmationAwareDeletePreviewActionQuery) Values() url.Values {
	return q.values
}
func (q EmailConfirmationAwareDeletePreviewActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *EmailConfirmationAwareDeletePreviewActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *EmailConfirmationAwareDeletePreviewActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type EmailConfirmationAwareDeletePreviewActionRequest struct {
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
func (x EmailConfirmationAwareDeletePreviewActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x EmailConfirmationAwareDeletePreviewActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func EmailConfirmationAwareDeletePreviewActionClientCreateUrl(
	req EmailConfirmationAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := EmailConfirmationAwareDeletePreviewActionMeta()
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
func EmailConfirmationAwareDeletePreviewActionClientExecuteTyped(httpReq *http.Request) (*EmailConfirmationAwareDeletePreviewActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result EmailConfirmationAwareDeletePreviewActionResponse
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
func EmailConfirmationAwareDeletePreviewActionClientBuildRequest(req EmailConfirmationAwareDeletePreviewActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := EmailConfirmationAwareDeletePreviewActionMeta()
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
func EmailConfirmationAwareDeletePreviewActionCall(
	req EmailConfirmationAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*EmailConfirmationAwareDeletePreviewActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := EmailConfirmationAwareDeletePreviewActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := EmailConfirmationAwareDeletePreviewActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return EmailConfirmationAwareDeletePreviewActionClientExecuteTyped(r)
}

// EmailConfirmationAwareDeletePreviewActionRaw registers a raw Gin route for the EmailConfirmationAwareDeletePreviewAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func EmailConfirmationAwareDeletePreviewActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := EmailConfirmationAwareDeletePreviewActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// EmailConfirmationAwareDeletePreviewActionHandler returns the HTTP method, route URL, and a typed Gin handler for the EmailConfirmationAwareDeletePreviewAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func EmailConfirmationAwareDeletePreviewActionHandler(
	handler func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := EmailConfirmationAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := EmailConfirmationAwareDeletePreviewActionRequest{
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

// EmailConfirmationAwareDeletePreviewActionGin is a high-level convenience wrapper around EmailConfirmationAwareDeletePreviewActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func EmailConfirmationAwareDeletePreviewActionGin(r gin.IRoutes, handler func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error)) {
	method, url, h := EmailConfirmationAwareDeletePreviewActionHandler(handler)
	r.Handle(method, url, h)
}
func (x EmailConfirmationAwareDeletePreviewActionRequest) IsGin() bool {
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
func EmailConfirmationAwareDeletePreviewActionQueryFromGin(c *gin.Context) EmailConfirmationAwareDeletePreviewActionQuery {
	return EmailConfirmationAwareDeletePreviewActionQueryFromString(c.Request.URL.RawQuery)
}
func GetEmailConfirmationAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// EmailConfirmationAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// EmailConfirmationAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func EmailConfirmationAwareDeletePreviewActionQueryFromCli(c *cli.Command) EmailConfirmationAwareDeletePreviewActionQuery {
	data := EmailConfirmationAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x EmailConfirmationAwareDeletePreviewActionRequest) IsCli() bool {
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

// EmailConfirmationAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the EmailConfirmationAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func EmailConfirmationAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEmailConfirmationAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// EmailConfirmationAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// EmailConfirmationAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a EmailConfirmationAwareDeletePreviewActionRequest the same way
// EmailConfirmationAwareDeletePreviewActionHandler (Gin) and EmailConfirmationAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func EmailConfirmationAwareDeletePreviewActionCliHandler(
	handler func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := EmailConfirmationAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: EmailConfirmationAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := EmailConfirmationAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = EmailConfirmationAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// EmailConfirmationAwareDeletePreviewActionCli is a high-level convenience wrapper around
// EmailConfirmationAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way EmailConfirmationAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func EmailConfirmationAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, EmailConfirmationAwareDeletePreviewActionCliHandler(handler))
}

// EmailConfirmationAwareDeletePreviewActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the EmailConfirmationAwareDeletePreviewAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *EmailConfirmationAwareDeletePreviewActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func EmailConfirmationAwareDeletePreviewActionHttpHandler(
	handler func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := EmailConfirmationAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := EmailConfirmationAwareDeletePreviewActionRequest{
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

// EmailConfirmationAwareDeletePreviewActionHttp is a high-level convenience wrapper around
// EmailConfirmationAwareDeletePreviewActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func EmailConfirmationAwareDeletePreviewActionHttp(
	mux *http.ServeMux,
	handler func(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error),
) {
	method, pattern, h := EmailConfirmationAwareDeletePreviewActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

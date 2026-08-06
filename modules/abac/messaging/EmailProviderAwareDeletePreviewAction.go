package messaging

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
* Action to communicate with the action EmailProviderAwareDeletePreviewAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of EmailProviderAwareDeletePreviewAction
func EmailProviderAwareDeletePreviewAction(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error) {
	return &EmailProviderAwareDeletePreviewActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func EmailProviderAwareDeletePreviewActionMeta() struct {
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
		Name:        "EmailProviderAwareDeletePreviewAction",
		CliName:     "email-provider-aware-delete-preview-action",
		CliShort:    "emailProvider-dp",
		URL:         "/emailProvider/delete-preview",
		Method:      "GET",
		Description: `Reports what deleting the given "emailProvider" uniqueIds would affect, without deleting anything.`,
	}
}

// The base class definition for emailProviderAwareDeletePreviewActionRes
type EmailProviderAwareDeletePreviewActionRes struct {
	Message  string                                                        `json:"message" yaml:"message"`
	Affected emigo.Array[EmailProviderAwareDeletePreviewActionResAffected] `json:"affected" yaml:"affected"`
}

// The base class definition for affected
type EmailProviderAwareDeletePreviewActionResAffected struct {
	Relation string `json:"relation" yaml:"relation"`
	Count    int64  `json:"count" yaml:"count"`
}

func (x *EmailProviderAwareDeletePreviewActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEmailProviderAwareDeletePreviewActionResCliFlags(prefix string) []emigo.CliFlag {
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
func CastEmailProviderAwareDeletePreviewActionResFromCli(c emigo.CliCastable) EmailProviderAwareDeletePreviewActionRes {
	data := EmailProviderAwareDeletePreviewActionRes{}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("affected") {
		data.Affected = emigo.CapturePossibleArray(CastEmailProviderAwareDeletePreviewActionResAffectedFromCli, "affected", c)
	}
	return data
}
func GetEmailProviderAwareDeletePreviewActionResAffectedCliFlags(prefix string) []emigo.CliFlag {
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
func CastEmailProviderAwareDeletePreviewActionResAffectedFromCli(c emigo.CliCastable) EmailProviderAwareDeletePreviewActionResAffected {
	data := EmailProviderAwareDeletePreviewActionResAffected{}
	if c.IsSet("relation") {
		data.Relation = c.String("relation")
	}
	if c.IsSet("count") {
		data.Count = int64(c.Int64("count"))
	}
	return data
}

type EmailProviderAwareDeletePreviewActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *EmailProviderAwareDeletePreviewActionResponse) SetContentType(contentType string) *EmailProviderAwareDeletePreviewActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *EmailProviderAwareDeletePreviewActionResponse) AsStream(r io.Reader, contentType string) *EmailProviderAwareDeletePreviewActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *EmailProviderAwareDeletePreviewActionResponse) AsJSON(payload any) *EmailProviderAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *EmailProviderAwareDeletePreviewActionResponse) WithIdeal(payload EmailProviderAwareDeletePreviewActionRes) *EmailProviderAwareDeletePreviewActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *EmailProviderAwareDeletePreviewActionResponse) AsIdeal() (*EmailProviderAwareDeletePreviewActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res EmailProviderAwareDeletePreviewActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *EmailProviderAwareDeletePreviewActionResponse) AsHTML(payload string) *EmailProviderAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *EmailProviderAwareDeletePreviewActionResponse) AsBytes(payload []byte) *EmailProviderAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x EmailProviderAwareDeletePreviewActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x EmailProviderAwareDeletePreviewActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x EmailProviderAwareDeletePreviewActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type EmailProviderAwareDeletePreviewActionRequestSig = func(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error)

/**
 * Query parameters for EmailProviderAwareDeletePreviewAction
 */
// Query wrapper with private fields
type EmailProviderAwareDeletePreviewActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	UniqueIds []string `json:"uniqueIds"`
}

func EmailProviderAwareDeletePreviewActionQueryFromString(rawQuery string) EmailProviderAwareDeletePreviewActionQuery {
	v := EmailProviderAwareDeletePreviewActionQuery{}
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
func EmailProviderAwareDeletePreviewActionQueryFromHttp(r *http.Request) EmailProviderAwareDeletePreviewActionQuery {
	return EmailProviderAwareDeletePreviewActionQueryFromString(r.URL.RawQuery)
}
func (q EmailProviderAwareDeletePreviewActionQuery) Values() url.Values {
	return q.values
}
func (q EmailProviderAwareDeletePreviewActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *EmailProviderAwareDeletePreviewActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *EmailProviderAwareDeletePreviewActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type EmailProviderAwareDeletePreviewActionRequest struct {
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
func (x EmailProviderAwareDeletePreviewActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x EmailProviderAwareDeletePreviewActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func EmailProviderAwareDeletePreviewActionClientCreateUrl(
	req EmailProviderAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := EmailProviderAwareDeletePreviewActionMeta()
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
func EmailProviderAwareDeletePreviewActionClientExecuteTyped(httpReq *http.Request) (*EmailProviderAwareDeletePreviewActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result EmailProviderAwareDeletePreviewActionResponse
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
func EmailProviderAwareDeletePreviewActionClientBuildRequest(req EmailProviderAwareDeletePreviewActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := EmailProviderAwareDeletePreviewActionMeta()
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
func EmailProviderAwareDeletePreviewActionCall(
	req EmailProviderAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*EmailProviderAwareDeletePreviewActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := EmailProviderAwareDeletePreviewActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := EmailProviderAwareDeletePreviewActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return EmailProviderAwareDeletePreviewActionClientExecuteTyped(r)
}

// EmailProviderAwareDeletePreviewActionRaw registers a raw Gin route for the EmailProviderAwareDeletePreviewAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func EmailProviderAwareDeletePreviewActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := EmailProviderAwareDeletePreviewActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// EmailProviderAwareDeletePreviewActionHandler returns the HTTP method, route URL, and a typed Gin handler for the EmailProviderAwareDeletePreviewAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func EmailProviderAwareDeletePreviewActionHandler(
	handler func(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := EmailProviderAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := EmailProviderAwareDeletePreviewActionRequest{
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

// EmailProviderAwareDeletePreviewActionGin is a high-level convenience wrapper around EmailProviderAwareDeletePreviewActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func EmailProviderAwareDeletePreviewActionGin(r gin.IRoutes, handler func(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error)) {
	method, url, h := EmailProviderAwareDeletePreviewActionHandler(handler)
	r.Handle(method, url, h)
}
func (x EmailProviderAwareDeletePreviewActionRequest) IsGin() bool {
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
func EmailProviderAwareDeletePreviewActionQueryFromGin(c *gin.Context) EmailProviderAwareDeletePreviewActionQuery {
	return EmailProviderAwareDeletePreviewActionQueryFromString(c.Request.URL.RawQuery)
}
func GetEmailProviderAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// EmailProviderAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// EmailProviderAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func EmailProviderAwareDeletePreviewActionQueryFromCli(c *cli.Command) EmailProviderAwareDeletePreviewActionQuery {
	data := EmailProviderAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x EmailProviderAwareDeletePreviewActionRequest) IsCli() bool {
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

// EmailProviderAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the EmailProviderAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func EmailProviderAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetEmailProviderAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// EmailProviderAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// EmailProviderAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a EmailProviderAwareDeletePreviewActionRequest the same way
// EmailProviderAwareDeletePreviewActionHandler (Gin) and EmailProviderAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func EmailProviderAwareDeletePreviewActionCliHandler(
	handler func(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := EmailProviderAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: EmailProviderAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := EmailProviderAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = EmailProviderAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// EmailProviderAwareDeletePreviewActionCli is a high-level convenience wrapper around
// EmailProviderAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way EmailProviderAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func EmailProviderAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, EmailProviderAwareDeletePreviewActionCliHandler(handler))
}

// EmailProviderAwareDeletePreviewActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the EmailProviderAwareDeletePreviewAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *EmailProviderAwareDeletePreviewActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func EmailProviderAwareDeletePreviewActionHttpHandler(
	handler func(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := EmailProviderAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := EmailProviderAwareDeletePreviewActionRequest{
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

// EmailProviderAwareDeletePreviewActionHttp is a high-level convenience wrapper around
// EmailProviderAwareDeletePreviewActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func EmailProviderAwareDeletePreviewActionHttp(
	mux *http.ServeMux,
	handler func(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error),
) {
	method, pattern, h := EmailProviderAwareDeletePreviewActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

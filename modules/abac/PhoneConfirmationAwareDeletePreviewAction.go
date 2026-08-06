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
* Action to communicate with the action PhoneConfirmationAwareDeletePreviewAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PhoneConfirmationAwareDeletePreviewAction
func PhoneConfirmationAwareDeletePreviewAction(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error) {
	return &PhoneConfirmationAwareDeletePreviewActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PhoneConfirmationAwareDeletePreviewActionMeta() struct {
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
		Name:        "PhoneConfirmationAwareDeletePreviewAction",
		CliName:     "phone-confirmation-aware-delete-preview-action",
		CliShort:    "phoneConfirmation-dp",
		URL:         "/phoneConfirmation/delete-preview",
		Method:      "GET",
		Description: `Reports what deleting the given "phoneConfirmation" uniqueIds would affect, without deleting anything.`,
	}
}

// The base class definition for phoneConfirmationAwareDeletePreviewActionRes
type PhoneConfirmationAwareDeletePreviewActionRes struct {
	Message  string                                                            `json:"message" yaml:"message"`
	Affected emigo.Array[PhoneConfirmationAwareDeletePreviewActionResAffected] `json:"affected" yaml:"affected"`
}

// The base class definition for affected
type PhoneConfirmationAwareDeletePreviewActionResAffected struct {
	Relation string `json:"relation" yaml:"relation"`
	Count    int64  `json:"count" yaml:"count"`
}

func (x *PhoneConfirmationAwareDeletePreviewActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPhoneConfirmationAwareDeletePreviewActionResCliFlags(prefix string) []emigo.CliFlag {
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
func CastPhoneConfirmationAwareDeletePreviewActionResFromCli(c emigo.CliCastable) PhoneConfirmationAwareDeletePreviewActionRes {
	data := PhoneConfirmationAwareDeletePreviewActionRes{}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("affected") {
		data.Affected = emigo.CapturePossibleArray(CastPhoneConfirmationAwareDeletePreviewActionResAffectedFromCli, "affected", c)
	}
	return data
}
func GetPhoneConfirmationAwareDeletePreviewActionResAffectedCliFlags(prefix string) []emigo.CliFlag {
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
func CastPhoneConfirmationAwareDeletePreviewActionResAffectedFromCli(c emigo.CliCastable) PhoneConfirmationAwareDeletePreviewActionResAffected {
	data := PhoneConfirmationAwareDeletePreviewActionResAffected{}
	if c.IsSet("relation") {
		data.Relation = c.String("relation")
	}
	if c.IsSet("count") {
		data.Count = int64(c.Int64("count"))
	}
	return data
}

type PhoneConfirmationAwareDeletePreviewActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PhoneConfirmationAwareDeletePreviewActionResponse) SetContentType(contentType string) *PhoneConfirmationAwareDeletePreviewActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PhoneConfirmationAwareDeletePreviewActionResponse) AsStream(r io.Reader, contentType string) *PhoneConfirmationAwareDeletePreviewActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PhoneConfirmationAwareDeletePreviewActionResponse) AsJSON(payload any) *PhoneConfirmationAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *PhoneConfirmationAwareDeletePreviewActionResponse) WithIdeal(payload PhoneConfirmationAwareDeletePreviewActionRes) *PhoneConfirmationAwareDeletePreviewActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *PhoneConfirmationAwareDeletePreviewActionResponse) AsIdeal() (*PhoneConfirmationAwareDeletePreviewActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res PhoneConfirmationAwareDeletePreviewActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *PhoneConfirmationAwareDeletePreviewActionResponse) AsHTML(payload string) *PhoneConfirmationAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PhoneConfirmationAwareDeletePreviewActionResponse) AsBytes(payload []byte) *PhoneConfirmationAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PhoneConfirmationAwareDeletePreviewActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PhoneConfirmationAwareDeletePreviewActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PhoneConfirmationAwareDeletePreviewActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type PhoneConfirmationAwareDeletePreviewActionRequestSig = func(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error)

/**
 * Query parameters for PhoneConfirmationAwareDeletePreviewAction
 */
// Query wrapper with private fields
type PhoneConfirmationAwareDeletePreviewActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	UniqueIds []string `json:"uniqueIds"`
}

func PhoneConfirmationAwareDeletePreviewActionQueryFromString(rawQuery string) PhoneConfirmationAwareDeletePreviewActionQuery {
	v := PhoneConfirmationAwareDeletePreviewActionQuery{}
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
func PhoneConfirmationAwareDeletePreviewActionQueryFromHttp(r *http.Request) PhoneConfirmationAwareDeletePreviewActionQuery {
	return PhoneConfirmationAwareDeletePreviewActionQueryFromString(r.URL.RawQuery)
}
func (q PhoneConfirmationAwareDeletePreviewActionQuery) Values() url.Values {
	return q.values
}
func (q PhoneConfirmationAwareDeletePreviewActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PhoneConfirmationAwareDeletePreviewActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PhoneConfirmationAwareDeletePreviewActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PhoneConfirmationAwareDeletePreviewActionRequest struct {
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
func (x PhoneConfirmationAwareDeletePreviewActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x PhoneConfirmationAwareDeletePreviewActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func PhoneConfirmationAwareDeletePreviewActionClientCreateUrl(
	req PhoneConfirmationAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PhoneConfirmationAwareDeletePreviewActionMeta()
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
func PhoneConfirmationAwareDeletePreviewActionClientExecuteTyped(httpReq *http.Request) (*PhoneConfirmationAwareDeletePreviewActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PhoneConfirmationAwareDeletePreviewActionResponse
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
func PhoneConfirmationAwareDeletePreviewActionClientBuildRequest(req PhoneConfirmationAwareDeletePreviewActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PhoneConfirmationAwareDeletePreviewActionMeta()
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
func PhoneConfirmationAwareDeletePreviewActionCall(
	req PhoneConfirmationAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PhoneConfirmationAwareDeletePreviewActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PhoneConfirmationAwareDeletePreviewActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PhoneConfirmationAwareDeletePreviewActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PhoneConfirmationAwareDeletePreviewActionClientExecuteTyped(r)
}

// PhoneConfirmationAwareDeletePreviewActionRaw registers a raw Gin route for the PhoneConfirmationAwareDeletePreviewAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PhoneConfirmationAwareDeletePreviewActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PhoneConfirmationAwareDeletePreviewActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// PhoneConfirmationAwareDeletePreviewActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PhoneConfirmationAwareDeletePreviewAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PhoneConfirmationAwareDeletePreviewActionHandler(
	handler func(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PhoneConfirmationAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PhoneConfirmationAwareDeletePreviewActionRequest{
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

// PhoneConfirmationAwareDeletePreviewActionGin is a high-level convenience wrapper around PhoneConfirmationAwareDeletePreviewActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PhoneConfirmationAwareDeletePreviewActionGin(r gin.IRoutes, handler func(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error)) {
	method, url, h := PhoneConfirmationAwareDeletePreviewActionHandler(handler)
	r.Handle(method, url, h)
}
func (x PhoneConfirmationAwareDeletePreviewActionRequest) IsGin() bool {
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
func PhoneConfirmationAwareDeletePreviewActionQueryFromGin(c *gin.Context) PhoneConfirmationAwareDeletePreviewActionQuery {
	return PhoneConfirmationAwareDeletePreviewActionQueryFromString(c.Request.URL.RawQuery)
}
func GetPhoneConfirmationAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// PhoneConfirmationAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// PhoneConfirmationAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func PhoneConfirmationAwareDeletePreviewActionQueryFromCli(c *cli.Command) PhoneConfirmationAwareDeletePreviewActionQuery {
	data := PhoneConfirmationAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x PhoneConfirmationAwareDeletePreviewActionRequest) IsCli() bool {
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

// PhoneConfirmationAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PhoneConfirmationAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PhoneConfirmationAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPhoneConfirmationAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// PhoneConfirmationAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// PhoneConfirmationAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PhoneConfirmationAwareDeletePreviewActionRequest the same way
// PhoneConfirmationAwareDeletePreviewActionHandler (Gin) and PhoneConfirmationAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PhoneConfirmationAwareDeletePreviewActionCliHandler(
	handler func(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := PhoneConfirmationAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PhoneConfirmationAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PhoneConfirmationAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = PhoneConfirmationAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PhoneConfirmationAwareDeletePreviewActionCli is a high-level convenience wrapper around
// PhoneConfirmationAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PhoneConfirmationAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func PhoneConfirmationAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, PhoneConfirmationAwareDeletePreviewActionCliHandler(handler))
}

// PhoneConfirmationAwareDeletePreviewActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the PhoneConfirmationAwareDeletePreviewAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *PhoneConfirmationAwareDeletePreviewActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func PhoneConfirmationAwareDeletePreviewActionHttpHandler(
	handler func(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := PhoneConfirmationAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := PhoneConfirmationAwareDeletePreviewActionRequest{
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

// PhoneConfirmationAwareDeletePreviewActionHttp is a high-level convenience wrapper around
// PhoneConfirmationAwareDeletePreviewActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func PhoneConfirmationAwareDeletePreviewActionHttp(
	mux *http.ServeMux,
	handler func(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error),
) {
	method, pattern, h := PhoneConfirmationAwareDeletePreviewActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

package messaging

import (
	"bytes"
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
* Action to communicate with the action GsmProviderAwareDeleteAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GsmProviderAwareDeleteAction
func GsmProviderAwareDeleteAction(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error) {
	return &GsmProviderAwareDeleteActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GsmProviderAwareDeleteActionMeta() struct {
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
		Name:        "GsmProviderAwareDeleteAction",
		CliName:     "gsm-provider-aware-delete-action",
		CliShort:    "gsmProvider-d",
		URL:         "/gsmProvider/delete",
		Method:      "POST",
		Description: `Deletes the given "gsmProvider" uniqueIds, along with everything gsmProviderAwareDeletePreview reports.`,
	}
}

// The base class definition for gsmProviderAwareDeleteActionReq
type GsmProviderAwareDeleteActionReq struct {
	UniqueIds []string `json:"uniqueIds" yaml:"uniqueIds"`
}

func (x *GsmProviderAwareDeleteActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetGsmProviderAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastGsmProviderAwareDeleteActionReqFromCli(c emigo.CliCastable) GsmProviderAwareDeleteActionReq {
	data := GsmProviderAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}

type GsmProviderAwareDeleteActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GsmProviderAwareDeleteActionResponse) SetContentType(contentType string) *GsmProviderAwareDeleteActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GsmProviderAwareDeleteActionResponse) AsStream(r io.Reader, contentType string) *GsmProviderAwareDeleteActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GsmProviderAwareDeleteActionResponse) AsJSON(payload any) *GsmProviderAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *GsmProviderAwareDeleteActionResponse) AsHTML(payload string) *GsmProviderAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GsmProviderAwareDeleteActionResponse) AsBytes(payload []byte) *GsmProviderAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GsmProviderAwareDeleteActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GsmProviderAwareDeleteActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GsmProviderAwareDeleteActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type GsmProviderAwareDeleteActionRequestSig = func(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error)

/**
 * Query parameters for GsmProviderAwareDeleteAction
 */
// Query wrapper with private fields
type GsmProviderAwareDeleteActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GsmProviderAwareDeleteActionQueryFromString(rawQuery string) GsmProviderAwareDeleteActionQuery {
	v := GsmProviderAwareDeleteActionQuery{}
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
func GsmProviderAwareDeleteActionQueryFromHttp(r *http.Request) GsmProviderAwareDeleteActionQuery {
	return GsmProviderAwareDeleteActionQueryFromString(r.URL.RawQuery)
}
func (q GsmProviderAwareDeleteActionQuery) Values() url.Values {
	return q.values
}
func (q GsmProviderAwareDeleteActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GsmProviderAwareDeleteActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GsmProviderAwareDeleteActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GsmProviderAwareDeleteActionRequest struct {
	Body        GsmProviderAwareDeleteActionReq
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
func (x GsmProviderAwareDeleteActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x GsmProviderAwareDeleteActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func GsmProviderAwareDeleteActionClientCreateUrl(
	req GsmProviderAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GsmProviderAwareDeleteActionMeta()
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
func GsmProviderAwareDeleteActionClientExecuteTyped(httpReq *http.Request) (*GsmProviderAwareDeleteActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GsmProviderAwareDeleteActionResponse
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
func GsmProviderAwareDeleteActionClientBuildRequest(req GsmProviderAwareDeleteActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GsmProviderAwareDeleteActionMeta()
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
func GsmProviderAwareDeleteActionCall(
	req GsmProviderAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GsmProviderAwareDeleteActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GsmProviderAwareDeleteActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GsmProviderAwareDeleteActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GsmProviderAwareDeleteActionClientExecuteTyped(r)
}

// GsmProviderAwareDeleteActionRaw registers a raw Gin route for the GsmProviderAwareDeleteAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GsmProviderAwareDeleteActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GsmProviderAwareDeleteActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// GsmProviderAwareDeleteActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GsmProviderAwareDeleteAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GsmProviderAwareDeleteActionHandler(
	handler func(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := GsmProviderAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body GsmProviderAwareDeleteActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := GsmProviderAwareDeleteActionRequest{
			Body:        body,
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

// GsmProviderAwareDeleteActionGin is a high-level convenience wrapper around GsmProviderAwareDeleteActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GsmProviderAwareDeleteActionGin(r gin.IRoutes, handler func(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error)) {
	method, url, h := GsmProviderAwareDeleteActionHandler(handler)
	r.Handle(method, url, h)
}
func (x GsmProviderAwareDeleteActionRequest) IsGin() bool {
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
func GsmProviderAwareDeleteActionQueryFromGin(c *gin.Context) GsmProviderAwareDeleteActionQuery {
	return GsmProviderAwareDeleteActionQueryFromString(c.Request.URL.RawQuery)
}
func (x GsmProviderAwareDeleteActionRequest) IsCli() bool {
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

// GsmProviderAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the GsmProviderAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func GsmProviderAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetGsmProviderAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// GsmProviderAwareDeleteActionCliHandler builds a full *cli.Command for the
// GsmProviderAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a GsmProviderAwareDeleteActionRequest the same way
// GsmProviderAwareDeleteActionHandler (Gin) and GsmProviderAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func GsmProviderAwareDeleteActionCliHandler(
	handler func(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error),
) *cli.Command {
	meta := GsmProviderAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: GsmProviderAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := GsmProviderAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastGsmProviderAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// GsmProviderAwareDeleteActionCli is a high-level convenience wrapper around
// GsmProviderAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way GsmProviderAwareDeleteActionGin
// registers a route on a Gin engine.
func GsmProviderAwareDeleteActionCli(
	app *cli.Command,
	handler func(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, GsmProviderAwareDeleteActionCliHandler(handler))
}

// GsmProviderAwareDeleteActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the GsmProviderAwareDeleteAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *GsmProviderAwareDeleteActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func GsmProviderAwareDeleteActionHttpHandler(
	handler func(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := GsmProviderAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body GsmProviderAwareDeleteActionReq
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
		req := GsmProviderAwareDeleteActionRequest{
			Body:        body,
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

// GsmProviderAwareDeleteActionHttp is a high-level convenience wrapper around
// GsmProviderAwareDeleteActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func GsmProviderAwareDeleteActionHttp(
	mux *http.ServeMux,
	handler func(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error),
) {
	method, pattern, h := GsmProviderAwareDeleteActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

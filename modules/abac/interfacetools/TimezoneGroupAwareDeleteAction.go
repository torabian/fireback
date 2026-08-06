package interfacetools

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
* Action to communicate with the action TimezoneGroupAwareDeleteAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of TimezoneGroupAwareDeleteAction
func TimezoneGroupAwareDeleteAction(c TimezoneGroupAwareDeleteActionRequest) (*TimezoneGroupAwareDeleteActionResponse, error) {
	return &TimezoneGroupAwareDeleteActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func TimezoneGroupAwareDeleteActionMeta() struct {
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
		Name:        "TimezoneGroupAwareDeleteAction",
		CliName:     "timezone-group-aware-delete-action",
		CliShort:    "timezoneGroup-d",
		URL:         "/timezoneGroup/delete",
		Method:      "POST",
		Description: `Deletes the given "timezoneGroup" uniqueIds, along with everything timezoneGroupAwareDeletePreview reports.`,
	}
}

// The base class definition for timezoneGroupAwareDeleteActionReq
type TimezoneGroupAwareDeleteActionReq struct {
	UniqueIds []string `json:"uniqueIds" yaml:"uniqueIds"`
}

func (x *TimezoneGroupAwareDeleteActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetTimezoneGroupAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastTimezoneGroupAwareDeleteActionReqFromCli(c emigo.CliCastable) TimezoneGroupAwareDeleteActionReq {
	data := TimezoneGroupAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}

type TimezoneGroupAwareDeleteActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *TimezoneGroupAwareDeleteActionResponse) SetContentType(contentType string) *TimezoneGroupAwareDeleteActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *TimezoneGroupAwareDeleteActionResponse) AsStream(r io.Reader, contentType string) *TimezoneGroupAwareDeleteActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *TimezoneGroupAwareDeleteActionResponse) AsJSON(payload any) *TimezoneGroupAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *TimezoneGroupAwareDeleteActionResponse) AsHTML(payload string) *TimezoneGroupAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *TimezoneGroupAwareDeleteActionResponse) AsBytes(payload []byte) *TimezoneGroupAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x TimezoneGroupAwareDeleteActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x TimezoneGroupAwareDeleteActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x TimezoneGroupAwareDeleteActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type TimezoneGroupAwareDeleteActionRequestSig = func(c TimezoneGroupAwareDeleteActionRequest) (*TimezoneGroupAwareDeleteActionResponse, error)

/**
 * Query parameters for TimezoneGroupAwareDeleteAction
 */
// Query wrapper with private fields
type TimezoneGroupAwareDeleteActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func TimezoneGroupAwareDeleteActionQueryFromString(rawQuery string) TimezoneGroupAwareDeleteActionQuery {
	v := TimezoneGroupAwareDeleteActionQuery{}
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
func TimezoneGroupAwareDeleteActionQueryFromHttp(r *http.Request) TimezoneGroupAwareDeleteActionQuery {
	return TimezoneGroupAwareDeleteActionQueryFromString(r.URL.RawQuery)
}
func (q TimezoneGroupAwareDeleteActionQuery) Values() url.Values {
	return q.values
}
func (q TimezoneGroupAwareDeleteActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *TimezoneGroupAwareDeleteActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *TimezoneGroupAwareDeleteActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type TimezoneGroupAwareDeleteActionRequest struct {
	Body        TimezoneGroupAwareDeleteActionReq
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
func (x TimezoneGroupAwareDeleteActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x TimezoneGroupAwareDeleteActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func TimezoneGroupAwareDeleteActionClientCreateUrl(
	req TimezoneGroupAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := TimezoneGroupAwareDeleteActionMeta()
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
func TimezoneGroupAwareDeleteActionClientExecuteTyped(httpReq *http.Request) (*TimezoneGroupAwareDeleteActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result TimezoneGroupAwareDeleteActionResponse
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
func TimezoneGroupAwareDeleteActionClientBuildRequest(req TimezoneGroupAwareDeleteActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := TimezoneGroupAwareDeleteActionMeta()
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
func TimezoneGroupAwareDeleteActionCall(
	req TimezoneGroupAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*TimezoneGroupAwareDeleteActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := TimezoneGroupAwareDeleteActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := TimezoneGroupAwareDeleteActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return TimezoneGroupAwareDeleteActionClientExecuteTyped(r)
}

// TimezoneGroupAwareDeleteActionRaw registers a raw Gin route for the TimezoneGroupAwareDeleteAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func TimezoneGroupAwareDeleteActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := TimezoneGroupAwareDeleteActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// TimezoneGroupAwareDeleteActionHandler returns the HTTP method, route URL, and a typed Gin handler for the TimezoneGroupAwareDeleteAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func TimezoneGroupAwareDeleteActionHandler(
	handler func(c TimezoneGroupAwareDeleteActionRequest) (*TimezoneGroupAwareDeleteActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := TimezoneGroupAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body TimezoneGroupAwareDeleteActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := TimezoneGroupAwareDeleteActionRequest{
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

// TimezoneGroupAwareDeleteActionGin is a high-level convenience wrapper around TimezoneGroupAwareDeleteActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func TimezoneGroupAwareDeleteActionGin(r gin.IRoutes, handler func(c TimezoneGroupAwareDeleteActionRequest) (*TimezoneGroupAwareDeleteActionResponse, error)) {
	method, url, h := TimezoneGroupAwareDeleteActionHandler(handler)
	r.Handle(method, url, h)
}
func (x TimezoneGroupAwareDeleteActionRequest) IsGin() bool {
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
func TimezoneGroupAwareDeleteActionQueryFromGin(c *gin.Context) TimezoneGroupAwareDeleteActionQuery {
	return TimezoneGroupAwareDeleteActionQueryFromString(c.Request.URL.RawQuery)
}
func (x TimezoneGroupAwareDeleteActionRequest) IsCli() bool {
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

// TimezoneGroupAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TimezoneGroupAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TimezoneGroupAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTimezoneGroupAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// TimezoneGroupAwareDeleteActionCliHandler builds a full *cli.Command for the
// TimezoneGroupAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TimezoneGroupAwareDeleteActionRequest the same way
// TimezoneGroupAwareDeleteActionHandler (Gin) and TimezoneGroupAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TimezoneGroupAwareDeleteActionCliHandler(
	handler func(c TimezoneGroupAwareDeleteActionRequest) (*TimezoneGroupAwareDeleteActionResponse, error),
) *cli.Command {
	meta := TimezoneGroupAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TimezoneGroupAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TimezoneGroupAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastTimezoneGroupAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TimezoneGroupAwareDeleteActionCli is a high-level convenience wrapper around
// TimezoneGroupAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TimezoneGroupAwareDeleteActionGin
// registers a route on a Gin engine.
func TimezoneGroupAwareDeleteActionCli(
	app *cli.Command,
	handler func(c TimezoneGroupAwareDeleteActionRequest) (*TimezoneGroupAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, TimezoneGroupAwareDeleteActionCliHandler(handler))
}

// TimezoneGroupAwareDeleteActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the TimezoneGroupAwareDeleteAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *TimezoneGroupAwareDeleteActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func TimezoneGroupAwareDeleteActionHttpHandler(
	handler func(c TimezoneGroupAwareDeleteActionRequest) (*TimezoneGroupAwareDeleteActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := TimezoneGroupAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body TimezoneGroupAwareDeleteActionReq
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
		req := TimezoneGroupAwareDeleteActionRequest{
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

// TimezoneGroupAwareDeleteActionHttp is a high-level convenience wrapper around
// TimezoneGroupAwareDeleteActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func TimezoneGroupAwareDeleteActionHttp(
	mux *http.ServeMux,
	handler func(c TimezoneGroupAwareDeleteActionRequest) (*TimezoneGroupAwareDeleteActionResponse, error),
) {
	method, pattern, h := TimezoneGroupAwareDeleteActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

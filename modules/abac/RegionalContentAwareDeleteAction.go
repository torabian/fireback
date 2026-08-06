package abac

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
* Action to communicate with the action RegionalContentAwareDeleteAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of RegionalContentAwareDeleteAction
func RegionalContentAwareDeleteAction(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error) {
	return &RegionalContentAwareDeleteActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func RegionalContentAwareDeleteActionMeta() struct {
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
		Name:        "RegionalContentAwareDeleteAction",
		CliName:     "regional-content-aware-delete-action",
		CliShort:    "regionalContent-d",
		URL:         "/regionalContent/delete",
		Method:      "POST",
		Description: `Deletes the given "regionalContent" uniqueIds, along with everything regionalContentAwareDeletePreview reports.`,
	}
}

// The base class definition for regionalContentAwareDeleteActionReq
type RegionalContentAwareDeleteActionReq struct {
	UniqueIds []string `json:"uniqueIds" yaml:"uniqueIds"`
}

func (x *RegionalContentAwareDeleteActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetRegionalContentAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastRegionalContentAwareDeleteActionReqFromCli(c emigo.CliCastable) RegionalContentAwareDeleteActionReq {
	data := RegionalContentAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}

type RegionalContentAwareDeleteActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *RegionalContentAwareDeleteActionResponse) SetContentType(contentType string) *RegionalContentAwareDeleteActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *RegionalContentAwareDeleteActionResponse) AsStream(r io.Reader, contentType string) *RegionalContentAwareDeleteActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *RegionalContentAwareDeleteActionResponse) AsJSON(payload any) *RegionalContentAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *RegionalContentAwareDeleteActionResponse) AsHTML(payload string) *RegionalContentAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *RegionalContentAwareDeleteActionResponse) AsBytes(payload []byte) *RegionalContentAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x RegionalContentAwareDeleteActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x RegionalContentAwareDeleteActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x RegionalContentAwareDeleteActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type RegionalContentAwareDeleteActionRequestSig = func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error)

/**
 * Query parameters for RegionalContentAwareDeleteAction
 */
// Query wrapper with private fields
type RegionalContentAwareDeleteActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func RegionalContentAwareDeleteActionQueryFromString(rawQuery string) RegionalContentAwareDeleteActionQuery {
	v := RegionalContentAwareDeleteActionQuery{}
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
func RegionalContentAwareDeleteActionQueryFromHttp(r *http.Request) RegionalContentAwareDeleteActionQuery {
	return RegionalContentAwareDeleteActionQueryFromString(r.URL.RawQuery)
}
func (q RegionalContentAwareDeleteActionQuery) Values() url.Values {
	return q.values
}
func (q RegionalContentAwareDeleteActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *RegionalContentAwareDeleteActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *RegionalContentAwareDeleteActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type RegionalContentAwareDeleteActionRequest struct {
	Body        RegionalContentAwareDeleteActionReq
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
func (x RegionalContentAwareDeleteActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x RegionalContentAwareDeleteActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func RegionalContentAwareDeleteActionClientCreateUrl(
	req RegionalContentAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := RegionalContentAwareDeleteActionMeta()
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
func RegionalContentAwareDeleteActionClientExecuteTyped(httpReq *http.Request) (*RegionalContentAwareDeleteActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result RegionalContentAwareDeleteActionResponse
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
func RegionalContentAwareDeleteActionClientBuildRequest(req RegionalContentAwareDeleteActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := RegionalContentAwareDeleteActionMeta()
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
func RegionalContentAwareDeleteActionCall(
	req RegionalContentAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*RegionalContentAwareDeleteActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := RegionalContentAwareDeleteActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := RegionalContentAwareDeleteActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return RegionalContentAwareDeleteActionClientExecuteTyped(r)
}

// RegionalContentAwareDeleteActionRaw registers a raw Gin route for the RegionalContentAwareDeleteAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func RegionalContentAwareDeleteActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := RegionalContentAwareDeleteActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// RegionalContentAwareDeleteActionHandler returns the HTTP method, route URL, and a typed Gin handler for the RegionalContentAwareDeleteAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func RegionalContentAwareDeleteActionHandler(
	handler func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := RegionalContentAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body RegionalContentAwareDeleteActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := RegionalContentAwareDeleteActionRequest{
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

// RegionalContentAwareDeleteActionGin is a high-level convenience wrapper around RegionalContentAwareDeleteActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func RegionalContentAwareDeleteActionGin(r gin.IRoutes, handler func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error)) {
	method, url, h := RegionalContentAwareDeleteActionHandler(handler)
	r.Handle(method, url, h)
}
func (x RegionalContentAwareDeleteActionRequest) IsGin() bool {
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
func RegionalContentAwareDeleteActionQueryFromGin(c *gin.Context) RegionalContentAwareDeleteActionQuery {
	return RegionalContentAwareDeleteActionQueryFromString(c.Request.URL.RawQuery)
}
func (x RegionalContentAwareDeleteActionRequest) IsCli() bool {
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

// RegionalContentAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the RegionalContentAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func RegionalContentAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetRegionalContentAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// RegionalContentAwareDeleteActionCliHandler builds a full *cli.Command for the
// RegionalContentAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a RegionalContentAwareDeleteActionRequest the same way
// RegionalContentAwareDeleteActionHandler (Gin) and RegionalContentAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func RegionalContentAwareDeleteActionCliHandler(
	handler func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error),
) *cli.Command {
	meta := RegionalContentAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: RegionalContentAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := RegionalContentAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastRegionalContentAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// RegionalContentAwareDeleteActionCli is a high-level convenience wrapper around
// RegionalContentAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way RegionalContentAwareDeleteActionGin
// registers a route on a Gin engine.
func RegionalContentAwareDeleteActionCli(
	app *cli.Command,
	handler func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, RegionalContentAwareDeleteActionCliHandler(handler))
}

// RegionalContentAwareDeleteActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the RegionalContentAwareDeleteAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *RegionalContentAwareDeleteActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func RegionalContentAwareDeleteActionHttpHandler(
	handler func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := RegionalContentAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body RegionalContentAwareDeleteActionReq
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
		req := RegionalContentAwareDeleteActionRequest{
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

// RegionalContentAwareDeleteActionHttp is a high-level convenience wrapper around
// RegionalContentAwareDeleteActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func RegionalContentAwareDeleteActionHttp(
	mux *http.ServeMux,
	handler func(c RegionalContentAwareDeleteActionRequest) (*RegionalContentAwareDeleteActionResponse, error),
) {
	method, pattern, h := RegionalContentAwareDeleteActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

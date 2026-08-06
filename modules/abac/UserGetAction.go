package abac

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

/**
* Action to communicate with the action UserGetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of UserGetAction
func UserGetAction(c UserGetActionRequest) (*UserGetActionResponse, error) {
	return &UserGetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func UserGetActionMeta() struct {
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
		Name:        "UserGetAction",
		CliName:     "user-get-action",
		CliShort:    "user-g",
		URL:         "/user/:uniqueId",
		Method:      "GET",
		Description: `Looks up a single "user" row by uniqueId.`,
	}
}

type UserGetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *UserGetActionResponse) SetContentType(contentType string) *UserGetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *UserGetActionResponse) AsStream(r io.Reader, contentType string) *UserGetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *UserGetActionResponse) AsJSON(payload any) *UserGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *UserGetActionResponse) WithIdeal(payload UserDto) *UserGetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *UserGetActionResponse) AsIdeal() (*UserDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res UserDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *UserGetActionResponse) AsHTML(payload string) *UserGetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *UserGetActionResponse) AsBytes(payload []byte) *UserGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x UserGetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x UserGetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x UserGetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type UserGetActionRequestSig = func(c UserGetActionRequest) (*UserGetActionResponse, error)

/**
 * Path parameters for UserGetAction
 */
type UserGetActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func UserGetActionPathParameterApply(params UserGetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func UserGetActionPathParameterFromFn(fn func(key string) string) UserGetActionPathParameter {
	res := UserGetActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for UserGetAction
 */
// Query wrapper with private fields
type UserGetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func UserGetActionQueryFromString(rawQuery string) UserGetActionQuery {
	v := UserGetActionQuery{}
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
func UserGetActionQueryFromHttp(r *http.Request) UserGetActionQuery {
	return UserGetActionQueryFromString(r.URL.RawQuery)
}
func (q UserGetActionQuery) Values() url.Values {
	return q.values
}
func (q UserGetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *UserGetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *UserGetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type UserGetActionRequest struct {
	Body        interface{}
	Params      UserGetActionPathParameter
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
func (x UserGetActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x UserGetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func UserGetActionClientCreateUrl(
	req UserGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := UserGetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = UserGetActionPathParameterApply(req.Params, urlAddr)
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
func UserGetActionClientExecuteTyped(httpReq *http.Request) (*UserGetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result UserGetActionResponse
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
func UserGetActionClientBuildRequest(req UserGetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := UserGetActionMeta()
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
func UserGetActionCall(
	req UserGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UserGetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := UserGetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := UserGetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return UserGetActionClientExecuteTyped(r)
}
func UserGetActionPathParameterFromGin(g *gin.Context) UserGetActionPathParameter {
	return UserGetActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// UserGetActionRaw registers a raw Gin route for the UserGetAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func UserGetActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := UserGetActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// UserGetActionHandler returns the HTTP method, route URL, and a typed Gin handler for the UserGetAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func UserGetActionHandler(
	handler func(c UserGetActionRequest) (*UserGetActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := UserGetActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := UserGetActionRequest{
			Body:        nil,
			Params:      UserGetActionPathParameterFromGin(m),
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

// UserGetActionGin is a high-level convenience wrapper around UserGetActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func UserGetActionGin(r gin.IRoutes, handler func(c UserGetActionRequest) (*UserGetActionResponse, error)) {
	method, url, h := UserGetActionHandler(handler)
	r.Handle(method, url, h)
}
func (x UserGetActionRequest) IsGin() bool {
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
func UserGetActionQueryFromGin(c *gin.Context) UserGetActionQuery {
	return UserGetActionQueryFromString(c.Request.URL.RawQuery)
}
func GetUserGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func UserGetActionPathParameterFromCli(c *cli.Command) UserGetActionPathParameter {
	return UserGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x UserGetActionRequest) IsCli() bool {
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

// UserGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserGetActionPathParameterCliFlags(""))...)
	return flags
}

// UserGetActionCliHandler builds a full *cli.Command for the
// UserGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserGetActionRequest the same way
// UserGetActionHandler (Gin) and UserGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserGetActionCliHandler(
	handler func(c UserGetActionRequest) (*UserGetActionResponse, error),
) *cli.Command {
	meta := UserGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      UserGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserGetActionCli is a high-level convenience wrapper around
// UserGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserGetActionGin
// registers a route on a Gin engine.
func UserGetActionCli(
	app *cli.Command,
	handler func(c UserGetActionRequest) (*UserGetActionResponse, error),
) {
	app.Commands = append(app.Commands, UserGetActionCliHandler(handler))
}

// UserGetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the UserGetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *UserGetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func UserGetActionHttpHandler(
	handler func(c UserGetActionRequest) (*UserGetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := UserGetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := UserGetActionRequest{
			Body: nil,
			Params: UserGetActionPathParameterFromFn(func(key string) string {
				return r.PathValue(key)
			}),
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

// UserGetActionHttp is a high-level convenience wrapper around
// UserGetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func UserGetActionHttp(
	mux *http.ServeMux,
	handler func(c UserGetActionRequest) (*UserGetActionResponse, error),
) {
	method, pattern, h := UserGetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

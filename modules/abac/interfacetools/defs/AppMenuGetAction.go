package interfacetoolsdefs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
)

/**
* Action to communicate with t he action AppMenuGetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of AppMenuGetAction
func AppMenuGetAction(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error) {
	return &AppMenuGetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func AppMenuGetActionMeta() struct {
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
		Name:        "AppMenuGetAction",
		CliName:     "get",
		CliShort:    "appMenu-g",
		URL:         "/appMenu/:uniqueId",
		Method:      "GET",
		Description: `Looks up a single "appMenu" row by uniqueId.`,
	}
}

type AppMenuGetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *AppMenuGetActionResponse) SetContentType(contentType string) *AppMenuGetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *AppMenuGetActionResponse) AsStream(r io.Reader, contentType string) *AppMenuGetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *AppMenuGetActionResponse) AsJSON(payload any) *AppMenuGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *AppMenuGetActionResponse) WithIdeal(payload AppMenuDto) *AppMenuGetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *AppMenuGetActionResponse) AsIdeal() (*AppMenuDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res AppMenuDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *AppMenuGetActionResponse) AsHTML(payload string) *AppMenuGetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *AppMenuGetActionResponse) AsBytes(payload []byte) *AppMenuGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x AppMenuGetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x AppMenuGetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x AppMenuGetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type AppMenuGetActionRequestSig = func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error)

/**
 * Path parameters for AppMenuGetAction
 */
type AppMenuGetActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func AppMenuGetActionPathParameterApply(params AppMenuGetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func AppMenuGetActionPathParameterFromFn(fn func(key string) string) AppMenuGetActionPathParameter {
	res := AppMenuGetActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for AppMenuGetAction
 */
// Query wrapper with private fields
type AppMenuGetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func AppMenuGetActionQueryFromString(rawQuery string) AppMenuGetActionQuery {
	v := AppMenuGetActionQuery{}
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
func AppMenuGetActionQueryFromHttp(r *http.Request) AppMenuGetActionQuery {
	return AppMenuGetActionQueryFromString(r.URL.RawQuery)
}
func (q AppMenuGetActionQuery) Values() url.Values {
	return q.values
}
func (q AppMenuGetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *AppMenuGetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *AppMenuGetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type AppMenuGetActionRequest struct {
	Body        interface{}
	Params      AppMenuGetActionPathParameter
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
func (x AppMenuGetActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x AppMenuGetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func AppMenuGetActionClientCreateUrl(
	req AppMenuGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := AppMenuGetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = AppMenuGetActionPathParameterApply(req.Params, urlAddr)
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
func AppMenuGetActionClientExecuteTyped(httpReq *http.Request) (*AppMenuGetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result AppMenuGetActionResponse
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
func AppMenuGetActionClientBuildRequest(req AppMenuGetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := AppMenuGetActionMeta()
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
func AppMenuGetActionCall(
	req AppMenuGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*AppMenuGetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := AppMenuGetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := AppMenuGetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return AppMenuGetActionClientExecuteTyped(r)
}
func AppMenuGetActionPathParameterFromGin(g *gin.Context) AppMenuGetActionPathParameter {
	return AppMenuGetActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// AppMenuGetActionRaw registers a raw Gin route for the AppMenuGetAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func AppMenuGetActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := AppMenuGetActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// AppMenuGetActionHandler returns the HTTP method, route URL, and a typed Gin handler for the AppMenuGetAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func AppMenuGetActionHandler(
	handler func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := AppMenuGetActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := AppMenuGetActionRequest{
			Body:        nil,
			Params:      AppMenuGetActionPathParameterFromGin(m),
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
		resp, err := handler(req)
		if err != nil {
			// Some deeper call inside handler (e.g. a security/authorization check
			// that rejects the request before the handler's own business logic ever
			// runs) may have already written and aborted the response itself - gin
			// tracks that on the ResponseWriter regardless of who did the writing.
			// Rendering the bubbled-up error on top of that would append a second,
			// invalid JSON body after the first.
			if m.Writer.Written() {
				return
			}
			status := http.StatusInternalServerError
			// If the error knows how to render itself for a given language (e.g.
			// fireback.IError, whose ferror.Error.ToPublicJSON resolves its
			// {"$": ..., "en": ..., "fa": ...} message map down to one string), let it -
			// picking the language the same way the rest of the app resolves it: the
			// "acceptLanguage" query param first, else the Accept-Language header, else
			// "en".
			if converter, ok := err.(interface {
				ToPublicJSON(lang string) ([]byte, int32)
			}); ok {
				lang := m.Query("acceptLanguage")
				if lang == "" {
					lang = m.GetHeader("Accept-Language")
					if i := strings.IndexAny(lang, ",;-"); i >= 0 {
						lang = lang[:i]
					}
					lang = strings.ToLower(strings.TrimSpace(lang))
				}
				if lang == "" {
					lang = "en"
				}
				body, code := converter.ToPublicJSON(lang)
				if code != 0 {
					status = int(code)
				}
				// Nest the resolved object under "error" (rather than writing it as the
				// bare response body) so every error shape - this one, the generic
				// forwarded-JSON one below, and the plain-string one - answers with the
				// same {"error": ...} envelope. json.RawMessage keeps body embedded as
				// real JSON instead of being re-escaped into a string.
				m.JSON(status, gin.H{"error": json.RawMessage(body)})
				return
			}
			// Otherwise, other action errors may still stringify themselves as an
			// indented JSON object via their Error() method. If that's what we got,
			// forward it nested under "error" as real JSON (optionally honoring its own
			// "httpCode" field for the response status) instead of re-escaping it into a
			// string, which is what plain errors still get.
			msg := err.Error()
			trimmed := strings.TrimSpace(msg)
			if strings.HasPrefix(trimmed, "{") && json.Valid([]byte(trimmed)) {
				var probe struct {
					HttpCode int32 `json:"httpCode"`
				}
				if uErr := json.Unmarshal([]byte(trimmed), &probe); uErr == nil && probe.HttpCode != 0 {
					status = int(probe.HttpCode)
				}
				m.JSON(status, gin.H{"error": json.RawMessage(trimmed)})
				return
			}
			m.JSON(status, gin.H{"error": msg})
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

// AppMenuGetActionGin is a high-level convenience wrapper around AppMenuGetActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func AppMenuGetActionGin(r gin.IRoutes, handler func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error)) {
	method, url, h := AppMenuGetActionHandler(handler)
	r.Handle(method, url, h)
}
func (x AppMenuGetActionRequest) IsGin() bool {
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
func AppMenuGetActionQueryFromGin(c *gin.Context) AppMenuGetActionQuery {
	return AppMenuGetActionQueryFromString(c.Request.URL.RawQuery)
}
func GetAppMenuGetActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-uniqueId",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func AppMenuGetActionPathParameterFromCli(c *cli.Command) AppMenuGetActionPathParameter {
	return AppMenuGetActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x AppMenuGetActionRequest) IsCli() bool {
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

// AppMenuGetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the AppMenuGetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func AppMenuGetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetAppMenuGetActionPathParameterCliFlags(""))...)
	return flags
}

// AppMenuGetActionCliHandler builds a full *cli.Command for the
// AppMenuGetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a AppMenuGetActionRequest the same way
// AppMenuGetActionHandler (Gin) and AppMenuGetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func AppMenuGetActionCliHandler(
	handler func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error),
) *cli.Command {
	meta := AppMenuGetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: AppMenuGetActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := AppMenuGetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Params:      AppMenuGetActionPathParameterFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// AppMenuGetActionCli is a high-level convenience wrapper around
// AppMenuGetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way AppMenuGetActionGin
// registers a route on a Gin engine.
func AppMenuGetActionCli(
	app *cli.Command,
	handler func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error),
) {
	app.Commands = append(app.Commands, AppMenuGetActionCliHandler(handler))
}

// AppMenuGetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the AppMenuGetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *AppMenuGetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func AppMenuGetActionHttpHandler(
	handler func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := AppMenuGetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := AppMenuGetActionRequest{
			Body: nil,
			Params: AppMenuGetActionPathParameterFromFn(func(key string) string {
				return r.PathValue(key)
			}),
			QueryParams: r.URL.Query(),
			Headers:     r.Header,
		}
		resp, err := handler(req)
		if err != nil {
			status := http.StatusInternalServerError
			w.Header().Set("Content-Type", "application/json")
			// If the error knows how to render itself for a given language (e.g.
			// fireback.IError, whose ferror.Error.ToPublicJSON resolves its
			// {"$": ..., "en": ..., "fa": ...} message map down to one string), let it -
			// picking the language the same way the rest of the app resolves it: the
			// "acceptLanguage" query param first, else the Accept-Language header, else
			// "en".
			if converter, ok := err.(interface {
				ToPublicJSON(lang string) ([]byte, int32)
			}); ok {
				lang := r.URL.Query().Get("acceptLanguage")
				if lang == "" {
					lang = r.Header.Get("Accept-Language")
					if i := strings.IndexAny(lang, ",;-"); i >= 0 {
						lang = lang[:i]
					}
					lang = strings.ToLower(strings.TrimSpace(lang))
				}
				if lang == "" {
					lang = "en"
				}
				body, code := converter.ToPublicJSON(lang)
				if code != 0 {
					status = int(code)
				}
				// Nest the resolved object under "error" (rather than writing it as the
				// bare response body) so every error shape - this one, the generic
				// forwarded-JSON one below, and the plain-string one - answers with the
				// same {"error": ...} envelope. json.RawMessage keeps body embedded as
				// real JSON instead of being re-escaped into a string.
				wrapped, wErr := json.Marshal(map[string]json.RawMessage{"error": json.RawMessage(body)})
				w.WriteHeader(status)
				if wErr == nil {
					w.Write(wrapped)
				} else {
					w.Write(body)
				}
				return
			}
			// Otherwise, other action errors may still stringify themselves as an
			// indented JSON object via their Error() method. If that's what we got,
			// forward it nested under "error" as real JSON (optionally honoring its own
			// "httpCode" field for the response status) instead of re-escaping it into a
			// string, which is what plain errors still get.
			msg := err.Error()
			trimmed := strings.TrimSpace(msg)
			if strings.HasPrefix(trimmed, "{") && json.Valid([]byte(trimmed)) {
				var probe struct {
					HttpCode int32 `json:"httpCode"`
				}
				if uErr := json.Unmarshal([]byte(trimmed), &probe); uErr == nil && probe.HttpCode != 0 {
					status = int(probe.HttpCode)
				}
				wrapped, wErr := json.Marshal(map[string]json.RawMessage{"error": json.RawMessage(trimmed)})
				w.WriteHeader(status)
				if wErr == nil {
					w.Write(wrapped)
				} else {
					w.Write([]byte(trimmed))
				}
				return
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(map[string]string{"error": msg})
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

// AppMenuGetActionHttp is a high-level convenience wrapper around
// AppMenuGetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func AppMenuGetActionHttp(
	mux *http.ServeMux,
	handler func(c AppMenuGetActionRequest) (*AppMenuGetActionResponse, error),
) {
	method, pattern, h := AppMenuGetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

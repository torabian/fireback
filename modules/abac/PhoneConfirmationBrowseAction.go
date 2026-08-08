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
	"strconv"
	"strings"
)

/**
* Action to communicate with the action PhoneConfirmationBrowseAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PhoneConfirmationBrowseAction
func PhoneConfirmationBrowseAction(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error) {
	return &PhoneConfirmationBrowseActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PhoneConfirmationBrowseActionMeta() struct {
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
		Name:        "PhoneConfirmationBrowseAction",
		CliName:     "phone-confirmation-browse-action",
		CliShort:    "phoneConfirmation-b",
		URL:         "/phoneConfirmation/browse",
		Method:      "GET",
		Description: `Returns "phoneConfirmation" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).`,
	}
}

type PhoneConfirmationBrowseActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PhoneConfirmationBrowseActionResponse) SetContentType(contentType string) *PhoneConfirmationBrowseActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PhoneConfirmationBrowseActionResponse) AsStream(r io.Reader, contentType string) *PhoneConfirmationBrowseActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PhoneConfirmationBrowseActionResponse) AsJSON(payload any) *PhoneConfirmationBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *PhoneConfirmationBrowseActionResponse) WithIdeal(payload PhoneConfirmationOptionalDto) *PhoneConfirmationBrowseActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *PhoneConfirmationBrowseActionResponse) AsIdeal() (*PhoneConfirmationOptionalDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res PhoneConfirmationOptionalDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *PhoneConfirmationBrowseActionResponse) AsHTML(payload string) *PhoneConfirmationBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PhoneConfirmationBrowseActionResponse) AsBytes(payload []byte) *PhoneConfirmationBrowseActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PhoneConfirmationBrowseActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PhoneConfirmationBrowseActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PhoneConfirmationBrowseActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type PhoneConfirmationBrowseActionRequestSig = func(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error)

/**
 * Query parameters for PhoneConfirmationBrowseAction
 */
// Query wrapper with private fields
type PhoneConfirmationBrowseActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	Filter       string `json:"filter"`
	Sort         string `json:"sort"`
	StartIndex   int    `json:"startIndex"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Cursor       string `json:"cursor"`
}

func PhoneConfirmationBrowseActionQueryFromString(rawQuery string) PhoneConfirmationBrowseActionQuery {
	v := PhoneConfirmationBrowseActionQuery{}
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
func PhoneConfirmationBrowseActionQueryFromHttp(r *http.Request) PhoneConfirmationBrowseActionQuery {
	return PhoneConfirmationBrowseActionQueryFromString(r.URL.RawQuery)
}
func (q PhoneConfirmationBrowseActionQuery) Values() url.Values {
	return q.values
}
func (q PhoneConfirmationBrowseActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PhoneConfirmationBrowseActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PhoneConfirmationBrowseActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PhoneConfirmationBrowseActionRequest struct {
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
func (x PhoneConfirmationBrowseActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x PhoneConfirmationBrowseActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func PhoneConfirmationBrowseActionClientCreateUrl(
	req PhoneConfirmationBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PhoneConfirmationBrowseActionMeta()
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
func PhoneConfirmationBrowseActionClientExecuteTyped(httpReq *http.Request) (*PhoneConfirmationBrowseActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PhoneConfirmationBrowseActionResponse
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
func PhoneConfirmationBrowseActionClientBuildRequest(req PhoneConfirmationBrowseActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PhoneConfirmationBrowseActionMeta()
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
func PhoneConfirmationBrowseActionCall(
	req PhoneConfirmationBrowseActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PhoneConfirmationBrowseActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PhoneConfirmationBrowseActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PhoneConfirmationBrowseActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PhoneConfirmationBrowseActionClientExecuteTyped(r)
}

// PhoneConfirmationBrowseActionRaw registers a raw Gin route for the PhoneConfirmationBrowseAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PhoneConfirmationBrowseActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PhoneConfirmationBrowseActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// PhoneConfirmationBrowseActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PhoneConfirmationBrowseAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PhoneConfirmationBrowseActionHandler(
	handler func(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PhoneConfirmationBrowseActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := PhoneConfirmationBrowseActionRequest{
			Body:        nil,
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

// PhoneConfirmationBrowseActionGin is a high-level convenience wrapper around PhoneConfirmationBrowseActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PhoneConfirmationBrowseActionGin(r gin.IRoutes, handler func(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error)) {
	method, url, h := PhoneConfirmationBrowseActionHandler(handler)
	r.Handle(method, url, h)
}
func (x PhoneConfirmationBrowseActionRequest) IsGin() bool {
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
func PhoneConfirmationBrowseActionQueryFromGin(c *gin.Context) PhoneConfirmationBrowseActionQuery {
	return PhoneConfirmationBrowseActionQueryFromString(c.Request.URL.RawQuery)
}
func GetPhoneConfirmationBrowseActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-filter",
			Type: "string",
		},
		{
			Name: prefix + "qs-sort",
			Type: "string",
		},
		{
			Name: prefix + "qs-start-index",
			Type: "int",
		},
		{
			Name: prefix + "qs-items-per-page",
			Type: "int",
		},
		{
			Name: prefix + "qs-cursor",
			Type: "string",
		},
	}
}

// PhoneConfirmationBrowseActionQueryFromCli extracts and casts query parameters the same way
// PhoneConfirmationBrowseActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func PhoneConfirmationBrowseActionQueryFromCli(c *cli.Command) PhoneConfirmationBrowseActionQuery {
	data := PhoneConfirmationBrowseActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-filter") {
		data.Filter = c.String("qs-filter")
		values.Set("filter", data.Filter)
	}
	if c.IsSet("qs-sort") {
		data.Sort = c.String("qs-sort")
		values.Set("sort", data.Sort)
	}
	if c.IsSet("qs-start-index") {
		data.StartIndex = int(c.Int64("qs-start-index"))
		values.Set("startIndex", strconv.FormatInt(int64(data.StartIndex), 10))
	}
	if c.IsSet("qs-items-per-page") {
		data.ItemsPerPage = int(c.Int64("qs-items-per-page"))
		values.Set("itemsPerPage", strconv.FormatInt(int64(data.ItemsPerPage), 10))
	}
	if c.IsSet("qs-cursor") {
		data.Cursor = c.String("qs-cursor")
		values.Set("cursor", data.Cursor)
	}
	data.SetValues(values)
	return data
}
func (x PhoneConfirmationBrowseActionRequest) IsCli() bool {
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

// PhoneConfirmationBrowseActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PhoneConfirmationBrowseAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PhoneConfirmationBrowseActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPhoneConfirmationBrowseActionQueryCliFlags(""))...)
	return flags
}

// PhoneConfirmationBrowseActionCliHandler builds a full *cli.Command for the
// PhoneConfirmationBrowseAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PhoneConfirmationBrowseActionRequest the same way
// PhoneConfirmationBrowseActionHandler (Gin) and PhoneConfirmationBrowseActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PhoneConfirmationBrowseActionCliHandler(
	handler func(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error),
) *cli.Command {
	meta := PhoneConfirmationBrowseActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PhoneConfirmationBrowseActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PhoneConfirmationBrowseActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = PhoneConfirmationBrowseActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PhoneConfirmationBrowseActionCli is a high-level convenience wrapper around
// PhoneConfirmationBrowseActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PhoneConfirmationBrowseActionGin
// registers a route on a Gin engine.
func PhoneConfirmationBrowseActionCli(
	app *cli.Command,
	handler func(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error),
) {
	app.Commands = append(app.Commands, PhoneConfirmationBrowseActionCliHandler(handler))
}

// PhoneConfirmationBrowseActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the PhoneConfirmationBrowseAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *PhoneConfirmationBrowseActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func PhoneConfirmationBrowseActionHttpHandler(
	handler func(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := PhoneConfirmationBrowseActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := PhoneConfirmationBrowseActionRequest{
			Body:        nil,
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

// PhoneConfirmationBrowseActionHttp is a high-level convenience wrapper around
// PhoneConfirmationBrowseActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func PhoneConfirmationBrowseActionHttp(
	mux *http.ServeMux,
	handler func(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error),
) {
	method, pattern, h := PhoneConfirmationBrowseActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

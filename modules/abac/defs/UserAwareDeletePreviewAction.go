package abacdefs

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
	"strings"
)

/**
* Action to communicate with the action UserAwareDeletePreviewAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of UserAwareDeletePreviewAction
func UserAwareDeletePreviewAction(c UserAwareDeletePreviewActionRequest) (*UserAwareDeletePreviewActionResponse, error) {
	return &UserAwareDeletePreviewActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func UserAwareDeletePreviewActionMeta() struct {
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
		Name:        "UserAwareDeletePreviewAction",
		CliName:     "delete-preview",
		CliShort:    "user-dp",
		URL:         "/user/delete-preview",
		Method:      "GET",
		Description: `Reports what deleting the given "user" uniqueIds would affect, without deleting anything.`,
	}
}

// The base class definition for userAwareDeletePreviewActionRes
type UserAwareDeletePreviewActionRes struct {
	Message  string                                               `json:"message" yaml:"message"`
	Affected emigo.Array[UserAwareDeletePreviewActionResAffected] `json:"affected" yaml:"affected"`
}

// The base class definition for affected
type UserAwareDeletePreviewActionResAffected struct {
	Relation string `json:"relation" yaml:"relation"`
	Count    int64  `json:"count" yaml:"count"`
}

func (x *UserAwareDeletePreviewActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserAwareDeletePreviewActionResCliFlags(prefix string) []emigo.CliFlag {
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
func CastUserAwareDeletePreviewActionResFromCli(c emigo.CliCastable) UserAwareDeletePreviewActionRes {
	data := UserAwareDeletePreviewActionRes{}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("affected") {
		data.Affected = emigo.CapturePossibleArray(CastUserAwareDeletePreviewActionResAffectedFromCli, "affected", c)
	}
	return data
}
func GetUserAwareDeletePreviewActionResAffectedCliFlags(prefix string) []emigo.CliFlag {
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
func CastUserAwareDeletePreviewActionResAffectedFromCli(c emigo.CliCastable) UserAwareDeletePreviewActionResAffected {
	data := UserAwareDeletePreviewActionResAffected{}
	if c.IsSet("relation") {
		data.Relation = c.String("relation")
	}
	if c.IsSet("count") {
		data.Count = int64(c.Int64("count"))
	}
	return data
}

type UserAwareDeletePreviewActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *UserAwareDeletePreviewActionResponse) SetContentType(contentType string) *UserAwareDeletePreviewActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *UserAwareDeletePreviewActionResponse) AsStream(r io.Reader, contentType string) *UserAwareDeletePreviewActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *UserAwareDeletePreviewActionResponse) AsJSON(payload any) *UserAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *UserAwareDeletePreviewActionResponse) WithIdeal(payload UserAwareDeletePreviewActionRes) *UserAwareDeletePreviewActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *UserAwareDeletePreviewActionResponse) AsIdeal() (*UserAwareDeletePreviewActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res UserAwareDeletePreviewActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *UserAwareDeletePreviewActionResponse) AsHTML(payload string) *UserAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *UserAwareDeletePreviewActionResponse) AsBytes(payload []byte) *UserAwareDeletePreviewActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x UserAwareDeletePreviewActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x UserAwareDeletePreviewActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x UserAwareDeletePreviewActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type UserAwareDeletePreviewActionRequestSig = func(c UserAwareDeletePreviewActionRequest) (*UserAwareDeletePreviewActionResponse, error)

/**
 * Query parameters for UserAwareDeletePreviewAction
 */
// Query wrapper with private fields
type UserAwareDeletePreviewActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
	UniqueIds []string `json:"uniqueIds"`
}

func UserAwareDeletePreviewActionQueryFromString(rawQuery string) UserAwareDeletePreviewActionQuery {
	v := UserAwareDeletePreviewActionQuery{}
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
func UserAwareDeletePreviewActionQueryFromHttp(r *http.Request) UserAwareDeletePreviewActionQuery {
	return UserAwareDeletePreviewActionQueryFromString(r.URL.RawQuery)
}
func (q UserAwareDeletePreviewActionQuery) Values() url.Values {
	return q.values
}
func (q UserAwareDeletePreviewActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *UserAwareDeletePreviewActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *UserAwareDeletePreviewActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type UserAwareDeletePreviewActionRequest struct {
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
func (x UserAwareDeletePreviewActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x UserAwareDeletePreviewActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func UserAwareDeletePreviewActionClientCreateUrl(
	req UserAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := UserAwareDeletePreviewActionMeta()
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
func UserAwareDeletePreviewActionClientExecuteTyped(httpReq *http.Request) (*UserAwareDeletePreviewActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result UserAwareDeletePreviewActionResponse
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
func UserAwareDeletePreviewActionClientBuildRequest(req UserAwareDeletePreviewActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := UserAwareDeletePreviewActionMeta()
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
func UserAwareDeletePreviewActionCall(
	req UserAwareDeletePreviewActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UserAwareDeletePreviewActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := UserAwareDeletePreviewActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := UserAwareDeletePreviewActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return UserAwareDeletePreviewActionClientExecuteTyped(r)
}

// UserAwareDeletePreviewActionRaw registers a raw Gin route for the UserAwareDeletePreviewAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func UserAwareDeletePreviewActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := UserAwareDeletePreviewActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// UserAwareDeletePreviewActionHandler returns the HTTP method, route URL, and a typed Gin handler for the UserAwareDeletePreviewAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func UserAwareDeletePreviewActionHandler(
	handler func(c UserAwareDeletePreviewActionRequest) (*UserAwareDeletePreviewActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := UserAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := UserAwareDeletePreviewActionRequest{
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

// UserAwareDeletePreviewActionGin is a high-level convenience wrapper around UserAwareDeletePreviewActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func UserAwareDeletePreviewActionGin(r gin.IRoutes, handler func(c UserAwareDeletePreviewActionRequest) (*UserAwareDeletePreviewActionResponse, error)) {
	method, url, h := UserAwareDeletePreviewActionHandler(handler)
	r.Handle(method, url, h)
}
func (x UserAwareDeletePreviewActionRequest) IsGin() bool {
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
func UserAwareDeletePreviewActionQueryFromGin(c *gin.Context) UserAwareDeletePreviewActionQuery {
	return UserAwareDeletePreviewActionQueryFromString(c.Request.URL.RawQuery)
}
func GetUserAwareDeletePreviewActionQueryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "qs-unique-ids",
			Type: "slice",
		},
	}
}

// UserAwareDeletePreviewActionQueryFromCli extracts and casts query parameters the same way
// UserAwareDeletePreviewActionQueryFromString does, but reads them off urfave v3 CLI flags instead
// of a raw query string. The underlying url.Values (as returned by .Values()) is filled
// in using each field's real name, so code consuming req.QueryParams behaves the same
// whether the request came from HTTP or from the CLI.
func UserAwareDeletePreviewActionQueryFromCli(c *cli.Command) UserAwareDeletePreviewActionQuery {
	data := UserAwareDeletePreviewActionQuery{}
	values := url.Values{}
	if c.IsSet("qs-unique-ids") {
		raw := c.String("qs-unique-ids")
		emigo.InflatePossibleSlice(raw, &data.UniqueIds)
		values.Set("uniqueIds", raw)
	}
	data.SetValues(values)
	return data
}
func (x UserAwareDeletePreviewActionRequest) IsCli() bool {
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

// UserAwareDeletePreviewActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserAwareDeletePreviewAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserAwareDeletePreviewActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetUserAwareDeletePreviewActionQueryCliFlags(""))...)
	return flags
}

// UserAwareDeletePreviewActionCliHandler builds a full *cli.Command for the
// UserAwareDeletePreviewAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserAwareDeletePreviewActionRequest the same way
// UserAwareDeletePreviewActionHandler (Gin) and UserAwareDeletePreviewActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserAwareDeletePreviewActionCliHandler(
	handler func(c UserAwareDeletePreviewActionRequest) (*UserAwareDeletePreviewActionResponse, error),
) *cli.Command {
	meta := UserAwareDeletePreviewActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserAwareDeletePreviewActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserAwareDeletePreviewActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		req.QueryParams = UserAwareDeletePreviewActionQueryFromCli(c).Values()
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserAwareDeletePreviewActionCli is a high-level convenience wrapper around
// UserAwareDeletePreviewActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserAwareDeletePreviewActionGin
// registers a route on a Gin engine.
func UserAwareDeletePreviewActionCli(
	app *cli.Command,
	handler func(c UserAwareDeletePreviewActionRequest) (*UserAwareDeletePreviewActionResponse, error),
) {
	app.Commands = append(app.Commands, UserAwareDeletePreviewActionCliHandler(handler))
}

// UserAwareDeletePreviewActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the UserAwareDeletePreviewAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *UserAwareDeletePreviewActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func UserAwareDeletePreviewActionHttpHandler(
	handler func(c UserAwareDeletePreviewActionRequest) (*UserAwareDeletePreviewActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := UserAwareDeletePreviewActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := UserAwareDeletePreviewActionRequest{
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

// UserAwareDeletePreviewActionHttp is a high-level convenience wrapper around
// UserAwareDeletePreviewActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func UserAwareDeletePreviewActionHttp(
	mux *http.ServeMux,
	handler func(c UserAwareDeletePreviewActionRequest) (*UserAwareDeletePreviewActionResponse, error),
) {
	method, pattern, h := UserAwareDeletePreviewActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

package abacdefs

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
	"strings"
)

/**
* Action to communicate with the action CompletePassportPasswordResetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CompletePassportPasswordResetAction
func CompletePassportPasswordResetAction(c CompletePassportPasswordResetActionRequest) (*CompletePassportPasswordResetActionResponse, error) {
	return &CompletePassportPasswordResetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CompletePassportPasswordResetActionMeta() struct {
	Name        string
	CliName     string
	URL         string
	Method      string
	Description string
} {
	return struct {
		Name        string
		CliName     string
		URL         string
		Method      string
		Description string
	}{
		Name:        "CompletePassportPasswordResetAction",
		CliName:     "reset-password",
		URL:         "/passport/reset-password",
		Method:      "POST",
		Description: `Public. Finishes the self-service reset flow started by ClassicPassportRequestOtp or SendPassportResetEmail - given the passport value and the OTP code that was emailed/texted to it, sets a new password and signs the user in, the same way a successful ClassicPassportOtp login would.`,
	}
}

// The base class definition for completePassportPasswordResetActionReq
type CompletePassportPasswordResetActionReq struct {
	// Passport value (email, phone number) the OTP was sent to.
	Value string `json:"value" validate:"required" yaml:"value"`
	// The OTP code emailed/texted to the passport value.
	Otp string `json:"otp" validate:"required" yaml:"otp"`
	// New password meeting the security requirements.
	Password string `json:"password" validate:"required" yaml:"password"`
}

func (x *CompletePassportPasswordResetActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCompletePassportPasswordResetActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "Passport value (email, phone number) the OTP was sent to.",
		},
		{
			Name:        prefix + "otp",
			Type:        "string",
			Description: "The OTP code emailed/texted to the passport value.",
		},
		{
			Name:        prefix + "password",
			Type:        "string",
			Description: "New password meeting the security requirements.",
		},
	}
}
func CastCompletePassportPasswordResetActionReqFromCli(c emigo.CliCastable) CompletePassportPasswordResetActionReq {
	data := CompletePassportPasswordResetActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("otp") {
		data.Otp = c.String("otp")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

// The base class definition for completePassportPasswordResetActionRes
type CompletePassportPasswordResetActionRes struct {
	Reset   bool                              `json:"reset" yaml:"reset"`
	Session emigo.OneNullable[UserSessionDto] `json:"session" yaml:"session"`
}

func (x *CompletePassportPasswordResetActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCompletePassportPasswordResetActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "reset",
			Type: "bool",
		},
		{
			Name: prefix + "session",
			Type: "one?",
		},
	}
}
func CastCompletePassportPasswordResetActionResFromCli(c emigo.CliCastable) CompletePassportPasswordResetActionRes {
	data := CompletePassportPasswordResetActionRes{}
	if c.IsSet("reset") {
		data.Reset = bool(c.Bool("reset"))
	}
	if c.IsSet("session") {
		data.Session = emigo.CapturePossibleOneNullable(CastUserSessionDtoFromCli, "session", c)
	}
	return data
}

type CompletePassportPasswordResetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CompletePassportPasswordResetActionResponse) SetContentType(contentType string) *CompletePassportPasswordResetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CompletePassportPasswordResetActionResponse) AsStream(r io.Reader, contentType string) *CompletePassportPasswordResetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CompletePassportPasswordResetActionResponse) AsJSON(payload any) *CompletePassportPasswordResetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CompletePassportPasswordResetActionResponse) WithIdeal(payload CompletePassportPasswordResetActionRes) *CompletePassportPasswordResetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *CompletePassportPasswordResetActionResponse) AsIdeal() (*CompletePassportPasswordResetActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res CompletePassportPasswordResetActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *CompletePassportPasswordResetActionResponse) AsHTML(payload string) *CompletePassportPasswordResetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CompletePassportPasswordResetActionResponse) AsBytes(payload []byte) *CompletePassportPasswordResetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CompletePassportPasswordResetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CompletePassportPasswordResetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CompletePassportPasswordResetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type CompletePassportPasswordResetActionRequestSig = func(c CompletePassportPasswordResetActionRequest) (*CompletePassportPasswordResetActionResponse, error)

/**
 * Query parameters for CompletePassportPasswordResetAction
 */
// Query wrapper with private fields
type CompletePassportPasswordResetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CompletePassportPasswordResetActionQueryFromString(rawQuery string) CompletePassportPasswordResetActionQuery {
	v := CompletePassportPasswordResetActionQuery{}
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
func CompletePassportPasswordResetActionQueryFromHttp(r *http.Request) CompletePassportPasswordResetActionQuery {
	return CompletePassportPasswordResetActionQueryFromString(r.URL.RawQuery)
}
func (q CompletePassportPasswordResetActionQuery) Values() url.Values {
	return q.values
}
func (q CompletePassportPasswordResetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CompletePassportPasswordResetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CompletePassportPasswordResetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CompletePassportPasswordResetActionRequest struct {
	Body        CompletePassportPasswordResetActionReq
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
func (x CompletePassportPasswordResetActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x CompletePassportPasswordResetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func CompletePassportPasswordResetActionClientCreateUrl(
	req CompletePassportPasswordResetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CompletePassportPasswordResetActionMeta()
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
func CompletePassportPasswordResetActionClientExecuteTyped(httpReq *http.Request) (*CompletePassportPasswordResetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CompletePassportPasswordResetActionResponse
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
func CompletePassportPasswordResetActionClientBuildRequest(req CompletePassportPasswordResetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CompletePassportPasswordResetActionMeta()
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
func CompletePassportPasswordResetActionCall(
	req CompletePassportPasswordResetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CompletePassportPasswordResetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CompletePassportPasswordResetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CompletePassportPasswordResetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CompletePassportPasswordResetActionClientExecuteTyped(r)
}

// CompletePassportPasswordResetActionRaw registers a raw Gin route for the CompletePassportPasswordResetAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CompletePassportPasswordResetActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CompletePassportPasswordResetActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// CompletePassportPasswordResetActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CompletePassportPasswordResetAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CompletePassportPasswordResetActionHandler(
	handler func(c CompletePassportPasswordResetActionRequest) (*CompletePassportPasswordResetActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := CompletePassportPasswordResetActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body CompletePassportPasswordResetActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := CompletePassportPasswordResetActionRequest{
			Body:        body,
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

// CompletePassportPasswordResetActionGin is a high-level convenience wrapper around CompletePassportPasswordResetActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CompletePassportPasswordResetActionGin(r gin.IRoutes, handler func(c CompletePassportPasswordResetActionRequest) (*CompletePassportPasswordResetActionResponse, error)) {
	method, url, h := CompletePassportPasswordResetActionHandler(handler)
	r.Handle(method, url, h)
}
func (x CompletePassportPasswordResetActionRequest) IsGin() bool {
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
func CompletePassportPasswordResetActionQueryFromGin(c *gin.Context) CompletePassportPasswordResetActionQuery {
	return CompletePassportPasswordResetActionQueryFromString(c.Request.URL.RawQuery)
}
func (x CompletePassportPasswordResetActionRequest) IsCli() bool {
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

// CompletePassportPasswordResetActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the CompletePassportPasswordResetAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func CompletePassportPasswordResetActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetCompletePassportPasswordResetActionReqCliFlags(""))...)
	return flags
}

// CompletePassportPasswordResetActionCliHandler builds a full *cli.Command for the
// CompletePassportPasswordResetAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a CompletePassportPasswordResetActionRequest the same way
// CompletePassportPasswordResetActionHandler (Gin) and CompletePassportPasswordResetActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func CompletePassportPasswordResetActionCliHandler(
	handler func(c CompletePassportPasswordResetActionRequest) (*CompletePassportPasswordResetActionResponse, error),
) *cli.Command {
	meta := CompletePassportPasswordResetActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: CompletePassportPasswordResetActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := CompletePassportPasswordResetActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastCompletePassportPasswordResetActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// CompletePassportPasswordResetActionCli is a high-level convenience wrapper around
// CompletePassportPasswordResetActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way CompletePassportPasswordResetActionGin
// registers a route on a Gin engine.
func CompletePassportPasswordResetActionCli(
	app *cli.Command,
	handler func(c CompletePassportPasswordResetActionRequest) (*CompletePassportPasswordResetActionResponse, error),
) {
	app.Commands = append(app.Commands, CompletePassportPasswordResetActionCliHandler(handler))
}

// CompletePassportPasswordResetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the CompletePassportPasswordResetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *CompletePassportPasswordResetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func CompletePassportPasswordResetActionHttpHandler(
	handler func(c CompletePassportPasswordResetActionRequest) (*CompletePassportPasswordResetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := CompletePassportPasswordResetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body CompletePassportPasswordResetActionReq
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
		req := CompletePassportPasswordResetActionRequest{
			Body:        body,
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

// CompletePassportPasswordResetActionHttp is a high-level convenience wrapper around
// CompletePassportPasswordResetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func CompletePassportPasswordResetActionHttp(
	mux *http.ServeMux,
	handler func(c CompletePassportPasswordResetActionRequest) (*CompletePassportPasswordResetActionResponse, error),
) {
	method, pattern, h := CompletePassportPasswordResetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

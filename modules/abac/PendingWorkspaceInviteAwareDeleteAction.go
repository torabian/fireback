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
	"strings"
)

/**
* Action to communicate with the action PendingWorkspaceInviteAwareDeleteAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of PendingWorkspaceInviteAwareDeleteAction
func PendingWorkspaceInviteAwareDeleteAction(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error) {
	return &PendingWorkspaceInviteAwareDeleteActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func PendingWorkspaceInviteAwareDeleteActionMeta() struct {
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
		Name:        "PendingWorkspaceInviteAwareDeleteAction",
		CliName:     "pending-workspace-invite-aware-delete-action",
		CliShort:    "pendingWorkspaceInvite-d",
		URL:         "/pendingWorkspaceInvite/delete",
		Method:      "POST",
		Description: `Deletes the given "pendingWorkspaceInvite" uniqueIds, along with everything pendingWorkspaceInviteAwareDeletePreview reports.`,
	}
}

// The base class definition for pendingWorkspaceInviteAwareDeleteActionReq
type PendingWorkspaceInviteAwareDeleteActionReq struct {
	UniqueIds []string `json:"uniqueIds" yaml:"uniqueIds"`
}

func (x *PendingWorkspaceInviteAwareDeleteActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPendingWorkspaceInviteAwareDeleteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-ids",
			Type: "slice",
		},
	}
}
func CastPendingWorkspaceInviteAwareDeleteActionReqFromCli(c emigo.CliCastable) PendingWorkspaceInviteAwareDeleteActionReq {
	data := PendingWorkspaceInviteAwareDeleteActionReq{}
	if c.IsSet("unique-ids") {
		emigo.InflatePossibleSlice(c.String("unique-ids"), &data.UniqueIds)
	}
	return data
}

type PendingWorkspaceInviteAwareDeleteActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *PendingWorkspaceInviteAwareDeleteActionResponse) SetContentType(contentType string) *PendingWorkspaceInviteAwareDeleteActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *PendingWorkspaceInviteAwareDeleteActionResponse) AsStream(r io.Reader, contentType string) *PendingWorkspaceInviteAwareDeleteActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *PendingWorkspaceInviteAwareDeleteActionResponse) AsJSON(payload any) *PendingWorkspaceInviteAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *PendingWorkspaceInviteAwareDeleteActionResponse) AsHTML(payload string) *PendingWorkspaceInviteAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *PendingWorkspaceInviteAwareDeleteActionResponse) AsBytes(payload []byte) *PendingWorkspaceInviteAwareDeleteActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x PendingWorkspaceInviteAwareDeleteActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x PendingWorkspaceInviteAwareDeleteActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x PendingWorkspaceInviteAwareDeleteActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type PendingWorkspaceInviteAwareDeleteActionRequestSig = func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error)

/**
 * Query parameters for PendingWorkspaceInviteAwareDeleteAction
 */
// Query wrapper with private fields
type PendingWorkspaceInviteAwareDeleteActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func PendingWorkspaceInviteAwareDeleteActionQueryFromString(rawQuery string) PendingWorkspaceInviteAwareDeleteActionQuery {
	v := PendingWorkspaceInviteAwareDeleteActionQuery{}
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
func PendingWorkspaceInviteAwareDeleteActionQueryFromHttp(r *http.Request) PendingWorkspaceInviteAwareDeleteActionQuery {
	return PendingWorkspaceInviteAwareDeleteActionQueryFromString(r.URL.RawQuery)
}
func (q PendingWorkspaceInviteAwareDeleteActionQuery) Values() url.Values {
	return q.values
}
func (q PendingWorkspaceInviteAwareDeleteActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *PendingWorkspaceInviteAwareDeleteActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *PendingWorkspaceInviteAwareDeleteActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type PendingWorkspaceInviteAwareDeleteActionRequest struct {
	Body        PendingWorkspaceInviteAwareDeleteActionReq
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
func (x PendingWorkspaceInviteAwareDeleteActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x PendingWorkspaceInviteAwareDeleteActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func PendingWorkspaceInviteAwareDeleteActionClientCreateUrl(
	req PendingWorkspaceInviteAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := PendingWorkspaceInviteAwareDeleteActionMeta()
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
func PendingWorkspaceInviteAwareDeleteActionClientExecuteTyped(httpReq *http.Request) (*PendingWorkspaceInviteAwareDeleteActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result PendingWorkspaceInviteAwareDeleteActionResponse
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
func PendingWorkspaceInviteAwareDeleteActionClientBuildRequest(req PendingWorkspaceInviteAwareDeleteActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := PendingWorkspaceInviteAwareDeleteActionMeta()
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
func PendingWorkspaceInviteAwareDeleteActionCall(
	req PendingWorkspaceInviteAwareDeleteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*PendingWorkspaceInviteAwareDeleteActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := PendingWorkspaceInviteAwareDeleteActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := PendingWorkspaceInviteAwareDeleteActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return PendingWorkspaceInviteAwareDeleteActionClientExecuteTyped(r)
}

// PendingWorkspaceInviteAwareDeleteActionRaw registers a raw Gin route for the PendingWorkspaceInviteAwareDeleteAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func PendingWorkspaceInviteAwareDeleteActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := PendingWorkspaceInviteAwareDeleteActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// PendingWorkspaceInviteAwareDeleteActionHandler returns the HTTP method, route URL, and a typed Gin handler for the PendingWorkspaceInviteAwareDeleteAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func PendingWorkspaceInviteAwareDeleteActionHandler(
	handler func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := PendingWorkspaceInviteAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body PendingWorkspaceInviteAwareDeleteActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := PendingWorkspaceInviteAwareDeleteActionRequest{
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

// PendingWorkspaceInviteAwareDeleteActionGin is a high-level convenience wrapper around PendingWorkspaceInviteAwareDeleteActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func PendingWorkspaceInviteAwareDeleteActionGin(r gin.IRoutes, handler func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error)) {
	method, url, h := PendingWorkspaceInviteAwareDeleteActionHandler(handler)
	r.Handle(method, url, h)
}
func (x PendingWorkspaceInviteAwareDeleteActionRequest) IsGin() bool {
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
func PendingWorkspaceInviteAwareDeleteActionQueryFromGin(c *gin.Context) PendingWorkspaceInviteAwareDeleteActionQuery {
	return PendingWorkspaceInviteAwareDeleteActionQueryFromString(c.Request.URL.RawQuery)
}
func (x PendingWorkspaceInviteAwareDeleteActionRequest) IsCli() bool {
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

// PendingWorkspaceInviteAwareDeleteActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the PendingWorkspaceInviteAwareDeleteAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func PendingWorkspaceInviteAwareDeleteActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetPendingWorkspaceInviteAwareDeleteActionReqCliFlags(""))...)
	return flags
}

// PendingWorkspaceInviteAwareDeleteActionCliHandler builds a full *cli.Command for the
// PendingWorkspaceInviteAwareDeleteAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a PendingWorkspaceInviteAwareDeleteActionRequest the same way
// PendingWorkspaceInviteAwareDeleteActionHandler (Gin) and PendingWorkspaceInviteAwareDeleteActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func PendingWorkspaceInviteAwareDeleteActionCliHandler(
	handler func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error),
) *cli.Command {
	meta := PendingWorkspaceInviteAwareDeleteActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: PendingWorkspaceInviteAwareDeleteActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := PendingWorkspaceInviteAwareDeleteActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastPendingWorkspaceInviteAwareDeleteActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// PendingWorkspaceInviteAwareDeleteActionCli is a high-level convenience wrapper around
// PendingWorkspaceInviteAwareDeleteActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way PendingWorkspaceInviteAwareDeleteActionGin
// registers a route on a Gin engine.
func PendingWorkspaceInviteAwareDeleteActionCli(
	app *cli.Command,
	handler func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error),
) {
	app.Commands = append(app.Commands, PendingWorkspaceInviteAwareDeleteActionCliHandler(handler))
}

// PendingWorkspaceInviteAwareDeleteActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the PendingWorkspaceInviteAwareDeleteAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *PendingWorkspaceInviteAwareDeleteActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func PendingWorkspaceInviteAwareDeleteActionHttpHandler(
	handler func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := PendingWorkspaceInviteAwareDeleteActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body PendingWorkspaceInviteAwareDeleteActionReq
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
		req := PendingWorkspaceInviteAwareDeleteActionRequest{
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

// PendingWorkspaceInviteAwareDeleteActionHttp is a high-level convenience wrapper around
// PendingWorkspaceInviteAwareDeleteActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func PendingWorkspaceInviteAwareDeleteActionHttp(
	mux *http.ServeMux,
	handler func(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error),
) {
	method, pattern, h := PendingWorkspaceInviteAwareDeleteActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

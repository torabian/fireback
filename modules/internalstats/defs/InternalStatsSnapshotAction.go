package internalstatsdefs

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
* Action to communicate with the action InternalStatsSnapshotAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of InternalStatsSnapshotAction
func InternalStatsSnapshotAction(c InternalStatsSnapshotActionRequest) (*InternalStatsSnapshotActionResponse, error) {
	return &InternalStatsSnapshotActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func InternalStatsSnapshotActionMeta() struct {
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
		Name:        "InternalStatsSnapshotAction",
		CliName:     "snapshot",
		URL:         "/internal-stats/snapshot",
		Method:      "GET",
		Description: `One point-in-time snapshot of every measured stat, as JSON. Root workspace token required by default - see InternalStatsModuleConfig.Authorize.`,
	}
}

// The base class definition for internalStatsSnapshotActionRes
type InternalStatsSnapshotActionRes struct {
	// RFC3339 timestamp this snapshot was collected at.
	GeneratedAt string `json:"generatedAt" yaml:"generatedAt"`
	Hostname    string `json:"hostname" yaml:"hostname"`
	// Every measured stat, in a stable display order (grouped by category).
	Items emigo.Array[InternalStatsSnapshotActionResItems] `json:"items" yaml:"items"`
}

// The base class definition for items
type InternalStatsSnapshotActionResItems struct {
	// Stable machine-readable identifier, e.g. cpu.usedPercent.
	Key string `json:"key" yaml:"key"`
	// Human-readable label for display, e.g. CPU Usage.
	Label string `json:"label" yaml:"label"`
	// Display grouping, e.g. CPU, Memory, Disk, Network, Runtime, Host.
	Category string `json:"category" yaml:"category"`
	// Pre-formatted display value, e.g. 62.3% or 8.0 GB.
	Value string `json:"value" yaml:"value"`
	// The same value as a plain number, for programmatic use (0 when not numeric).
	RawValue float64 `json:"rawValue" yaml:"rawValue"`
	// Unit of rawValue, e.g. percent, bytes, seconds. Empty for non-numeric stats.
	Unit string `json:"unit" yaml:"unit"`
	// One of ok, warn, critical, info - a coarse threshold-based read on this stat.
	Severity string `json:"severity" yaml:"severity"`
}

func (x *InternalStatsSnapshotActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInternalStatsSnapshotActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "generated-at",
			Type:        "string",
			Description: "RFC3339 timestamp this snapshot was collected at.",
		},
		{
			Name: prefix + "hostname",
			Type: "string",
		},
		{
			Name:        prefix + "items",
			Type:        "array",
			Description: "Every measured stat, in a stable display order (grouped by category).",
		},
	}
}
func CastInternalStatsSnapshotActionResFromCli(c emigo.CliCastable) InternalStatsSnapshotActionRes {
	data := InternalStatsSnapshotActionRes{}
	if c.IsSet("generated-at") {
		data.GeneratedAt = c.String("generated-at")
	}
	if c.IsSet("hostname") {
		data.Hostname = c.String("hostname")
	}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastInternalStatsSnapshotActionResItemsFromCli, "items", c)
	}
	return data
}
func GetInternalStatsSnapshotActionResItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "key",
			Type:        "string",
			Description: "Stable machine-readable identifier, e.g. cpu.usedPercent.",
		},
		{
			Name:        prefix + "label",
			Type:        "string",
			Description: "Human-readable label for display, e.g. CPU Usage.",
		},
		{
			Name:        prefix + "category",
			Type:        "string",
			Description: "Display grouping, e.g. CPU, Memory, Disk, Network, Runtime, Host.",
		},
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "Pre-formatted display value, e.g. 62.3% or 8.0 GB.",
		},
		{
			Name:        prefix + "raw-value",
			Type:        "float64",
			Description: "The same value as a plain number, for programmatic use (0 when not numeric).",
		},
		{
			Name:        prefix + "unit",
			Type:        "string",
			Description: "Unit of rawValue, e.g. percent, bytes, seconds. Empty for non-numeric stats.",
		},
		{
			Name:        prefix + "severity",
			Type:        "string",
			Description: "One of ok, warn, critical, info - a coarse threshold-based read on this stat.",
		},
	}
}
func CastInternalStatsSnapshotActionResItemsFromCli(c emigo.CliCastable) InternalStatsSnapshotActionResItems {
	data := InternalStatsSnapshotActionResItems{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	if c.IsSet("category") {
		data.Category = c.String("category")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("raw-value") {
		data.RawValue = float64(c.Float64("raw-value"))
	}
	if c.IsSet("unit") {
		data.Unit = c.String("unit")
	}
	if c.IsSet("severity") {
		data.Severity = c.String("severity")
	}
	return data
}

type InternalStatsSnapshotActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *InternalStatsSnapshotActionResponse) SetContentType(contentType string) *InternalStatsSnapshotActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *InternalStatsSnapshotActionResponse) AsStream(r io.Reader, contentType string) *InternalStatsSnapshotActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *InternalStatsSnapshotActionResponse) AsJSON(payload any) *InternalStatsSnapshotActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *InternalStatsSnapshotActionResponse) WithIdeal(payload InternalStatsSnapshotActionRes) *InternalStatsSnapshotActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *InternalStatsSnapshotActionResponse) AsIdeal() (*InternalStatsSnapshotActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res InternalStatsSnapshotActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *InternalStatsSnapshotActionResponse) AsHTML(payload string) *InternalStatsSnapshotActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *InternalStatsSnapshotActionResponse) AsBytes(payload []byte) *InternalStatsSnapshotActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x InternalStatsSnapshotActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x InternalStatsSnapshotActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x InternalStatsSnapshotActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type InternalStatsSnapshotActionRequestSig = func(c InternalStatsSnapshotActionRequest) (*InternalStatsSnapshotActionResponse, error)

/**
 * Query parameters for InternalStatsSnapshotAction
 */
// Query wrapper with private fields
type InternalStatsSnapshotActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func InternalStatsSnapshotActionQueryFromString(rawQuery string) InternalStatsSnapshotActionQuery {
	v := InternalStatsSnapshotActionQuery{}
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
func InternalStatsSnapshotActionQueryFromHttp(r *http.Request) InternalStatsSnapshotActionQuery {
	return InternalStatsSnapshotActionQueryFromString(r.URL.RawQuery)
}
func (q InternalStatsSnapshotActionQuery) Values() url.Values {
	return q.values
}
func (q InternalStatsSnapshotActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *InternalStatsSnapshotActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *InternalStatsSnapshotActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type InternalStatsSnapshotActionRequest struct {
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
func (x InternalStatsSnapshotActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x InternalStatsSnapshotActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func InternalStatsSnapshotActionClientCreateUrl(
	req InternalStatsSnapshotActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := InternalStatsSnapshotActionMeta()
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
func InternalStatsSnapshotActionClientExecuteTyped(httpReq *http.Request) (*InternalStatsSnapshotActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result InternalStatsSnapshotActionResponse
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
func InternalStatsSnapshotActionClientBuildRequest(req InternalStatsSnapshotActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := InternalStatsSnapshotActionMeta()
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
func InternalStatsSnapshotActionCall(
	req InternalStatsSnapshotActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*InternalStatsSnapshotActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := InternalStatsSnapshotActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := InternalStatsSnapshotActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return InternalStatsSnapshotActionClientExecuteTyped(r)
}

// InternalStatsSnapshotActionRaw registers a raw Gin route for the InternalStatsSnapshotAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func InternalStatsSnapshotActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := InternalStatsSnapshotActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// InternalStatsSnapshotActionHandler returns the HTTP method, route URL, and a typed Gin handler for the InternalStatsSnapshotAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func InternalStatsSnapshotActionHandler(
	handler func(c InternalStatsSnapshotActionRequest) (*InternalStatsSnapshotActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := InternalStatsSnapshotActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := InternalStatsSnapshotActionRequest{
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

// InternalStatsSnapshotActionGin is a high-level convenience wrapper around InternalStatsSnapshotActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func InternalStatsSnapshotActionGin(r gin.IRoutes, handler func(c InternalStatsSnapshotActionRequest) (*InternalStatsSnapshotActionResponse, error)) {
	method, url, h := InternalStatsSnapshotActionHandler(handler)
	r.Handle(method, url, h)
}
func (x InternalStatsSnapshotActionRequest) IsGin() bool {
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
func InternalStatsSnapshotActionQueryFromGin(c *gin.Context) InternalStatsSnapshotActionQuery {
	return InternalStatsSnapshotActionQueryFromString(c.Request.URL.RawQuery)
}
func (x InternalStatsSnapshotActionRequest) IsCli() bool {
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

// InternalStatsSnapshotActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the InternalStatsSnapshotAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func InternalStatsSnapshotActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// InternalStatsSnapshotActionCliHandler builds a full *cli.Command for the
// InternalStatsSnapshotAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a InternalStatsSnapshotActionRequest the same way
// InternalStatsSnapshotActionHandler (Gin) and InternalStatsSnapshotActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func InternalStatsSnapshotActionCliHandler(
	handler func(c InternalStatsSnapshotActionRequest) (*InternalStatsSnapshotActionResponse, error),
) *cli.Command {
	meta := InternalStatsSnapshotActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: InternalStatsSnapshotActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := InternalStatsSnapshotActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// InternalStatsSnapshotActionCli is a high-level convenience wrapper around
// InternalStatsSnapshotActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way InternalStatsSnapshotActionGin
// registers a route on a Gin engine.
func InternalStatsSnapshotActionCli(
	app *cli.Command,
	handler func(c InternalStatsSnapshotActionRequest) (*InternalStatsSnapshotActionResponse, error),
) {
	app.Commands = append(app.Commands, InternalStatsSnapshotActionCliHandler(handler))
}

// InternalStatsSnapshotActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the InternalStatsSnapshotAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *InternalStatsSnapshotActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func InternalStatsSnapshotActionHttpHandler(
	handler func(c InternalStatsSnapshotActionRequest) (*InternalStatsSnapshotActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := InternalStatsSnapshotActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := InternalStatsSnapshotActionRequest{
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

// InternalStatsSnapshotActionHttp is a high-level convenience wrapper around
// InternalStatsSnapshotActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func InternalStatsSnapshotActionHttp(
	mux *http.ServeMux,
	handler func(c InternalStatsSnapshotActionRequest) (*InternalStatsSnapshotActionResponse, error),
) {
	method, pattern, h := InternalStatsSnapshotActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

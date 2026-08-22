package walletpublicdefs

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
* Action to communicate with the action TopupAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of TopupAction
func TopupAction(c TopupActionRequest) (*TopupActionResponse, error) {
	return &TopupActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func TopupActionMeta() struct {
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
		Name:        "TopupAction",
		CliName:     "topup-action",
		CliShort:    "topup",
		URL:         "/wallet/topup",
		Method:      "POST",
		Description: `Starts a topup of walletId through gatewayCode: creates a pending walletPaymentAttempt and asks the gateway adapter to initiate payment, returning whatever the caller needs to complete it (redirect URL and/or a client secret, gateway-dependent). idempotencyKey makes retrying a timed-out call safe - it will not create a second attempt at the gateway.`,
	}
}

// The base class definition for topupActionReq
type TopupActionReq struct {
	// Unique id of the wallet to top up.
	WalletId string `json:"walletId" validate:"required" yaml:"walletId"`
	// Code of the walletGateway to pay through.
	GatewayCode string `json:"gatewayCode" validate:"required" yaml:"gatewayCode"`
	// Amount to top up, as a positive minor-units string.
	Amount string `json:"amount" validate:"required" yaml:"amount"`
	// Makes this topup-initiation safe to retry.
	IdempotencyKey string `json:"idempotencyKey" validate:"required" yaml:"idempotencyKey"`
	// Where to send the caller's browser back to once a redirect-based gateway (Przelewy24, ZarinPal, BLIK) completes the payment. Not needed for gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow) - effectively required for the others.
	ReturnUrl emigo.Nullable[string] `json:"returnUrl" yaml:"returnUrl"`
}

func (x *TopupActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetTopupActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "wallet-id",
			Type:        "string",
			Description: "Unique id of the wallet to top up.",
		},
		{
			Name:        prefix + "gateway-code",
			Type:        "string",
			Description: "Code of the walletGateway to pay through.",
		},
		{
			Name:        prefix + "amount",
			Type:        "string",
			Description: "Amount to top up, as a positive minor-units string.",
		},
		{
			Name:        prefix + "idempotency-key",
			Type:        "string",
			Description: "Makes this topup-initiation safe to retry.",
		},
		{
			Name:        prefix + "return-url",
			Type:        "string?",
			Description: "Where to send the caller's browser back to once a redirect-based gateway (Przelewy24, ZarinPal, BLIK) completes the payment. Not needed for gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow) - effectively required for the others.",
		},
	}
}
func CastTopupActionReqFromCli(c emigo.CliCastable) TopupActionReq {
	data := TopupActionReq{}
	if c.IsSet("wallet-id") {
		data.WalletId = c.String("wallet-id")
	}
	if c.IsSet("gateway-code") {
		data.GatewayCode = c.String("gateway-code")
	}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("idempotency-key") {
		data.IdempotencyKey = c.String("idempotency-key")
	}
	if c.IsSet("return-url") {
		emigo.ParseNullable(c.String("return-url"), &data.ReturnUrl)
	}
	return data
}

// The base class definition for topupActionRes
type TopupActionRes struct {
	// The created payment attempt.
	Attempt TopupActionResAttempt `json:"attempt" yaml:"attempt"`
	// URL to send the owner to, for gateways that need one.
	RedirectUrl emigo.Nullable[string] `json:"redirectUrl" yaml:"redirectUrl"`
	// Client-side secret/token, for gateways that need one instead.
	ClientSecret emigo.Nullable[string] `json:"clientSecret" yaml:"clientSecret"`
}

// The base class definition for attempt
type TopupActionResAttempt struct {
	UniqueId    string `json:"uniqueId" yaml:"uniqueId"`
	Status      string `json:"status" yaml:"status"`
	GatewayCode string `json:"gatewayCode" yaml:"gatewayCode"`
}

func (x *TopupActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetTopupActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "attempt",
			Type:        "object",
			Children:    GetTopupActionResAttemptCliFlags("attempt-"),
			Description: "The created payment attempt.",
		},
		{
			Name:        prefix + "redirect-url",
			Type:        "string?",
			Description: "URL to send the owner to, for gateways that need one.",
		},
		{
			Name:        prefix + "client-secret",
			Type:        "string?",
			Description: "Client-side secret/token, for gateways that need one instead.",
		},
	}
}
func CastTopupActionResFromCli(c emigo.CliCastable) TopupActionRes {
	data := TopupActionRes{}
	if c.IsSet("attempt") {
		data.Attempt = CastTopupActionResAttemptFromCli(c)
	}
	if c.IsSet("redirect-url") {
		emigo.ParseNullable(c.String("redirect-url"), &data.RedirectUrl)
	}
	if c.IsSet("client-secret") {
		emigo.ParseNullable(c.String("client-secret"), &data.ClientSecret)
	}
	return data
}
func GetTopupActionResAttemptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "status",
			Type: "string",
		},
		{
			Name: prefix + "gateway-code",
			Type: "string",
		},
	}
}
func CastTopupActionResAttemptFromCli(c emigo.CliCastable) TopupActionResAttempt {
	data := TopupActionResAttempt{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("gateway-code") {
		data.GatewayCode = c.String("gateway-code")
	}
	return data
}

type TopupActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *TopupActionResponse) SetContentType(contentType string) *TopupActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *TopupActionResponse) AsStream(r io.Reader, contentType string) *TopupActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *TopupActionResponse) AsJSON(payload any) *TopupActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *TopupActionResponse) WithIdeal(payload TopupActionRes) *TopupActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *TopupActionResponse) AsIdeal() (*TopupActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res TopupActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *TopupActionResponse) AsHTML(payload string) *TopupActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *TopupActionResponse) AsBytes(payload []byte) *TopupActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x TopupActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x TopupActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x TopupActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type TopupActionRequestSig = func(c TopupActionRequest) (*TopupActionResponse, error)

/**
 * Query parameters for TopupAction
 */
// Query wrapper with private fields
type TopupActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func TopupActionQueryFromString(rawQuery string) TopupActionQuery {
	v := TopupActionQuery{}
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
func TopupActionQueryFromHttp(r *http.Request) TopupActionQuery {
	return TopupActionQueryFromString(r.URL.RawQuery)
}
func (q TopupActionQuery) Values() url.Values {
	return q.values
}
func (q TopupActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *TopupActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *TopupActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type TopupActionRequest struct {
	Body        TopupActionReq
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
func (x TopupActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x TopupActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func TopupActionClientCreateUrl(
	req TopupActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := TopupActionMeta()
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
func TopupActionClientExecuteTyped(httpReq *http.Request) (*TopupActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result TopupActionResponse
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
func TopupActionClientBuildRequest(req TopupActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := TopupActionMeta()
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
func TopupActionCall(
	req TopupActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*TopupActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := TopupActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := TopupActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return TopupActionClientExecuteTyped(r)
}

// TopupActionRaw registers a raw Gin route for the TopupAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func TopupActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := TopupActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// TopupActionHandler returns the HTTP method, route URL, and a typed Gin handler for the TopupAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func TopupActionHandler(
	handler func(c TopupActionRequest) (*TopupActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := TopupActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body TopupActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := TopupActionRequest{
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

// TopupActionGin is a high-level convenience wrapper around TopupActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func TopupActionGin(r gin.IRoutes, handler func(c TopupActionRequest) (*TopupActionResponse, error)) {
	method, url, h := TopupActionHandler(handler)
	r.Handle(method, url, h)
}
func (x TopupActionRequest) IsGin() bool {
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
func TopupActionQueryFromGin(c *gin.Context) TopupActionQuery {
	return TopupActionQueryFromString(c.Request.URL.RawQuery)
}
func (x TopupActionRequest) IsCli() bool {
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

// TopupActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the TopupAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func TopupActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetTopupActionReqCliFlags(""))...)
	return flags
}

// TopupActionCliHandler builds a full *cli.Command for the
// TopupAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a TopupActionRequest the same way
// TopupActionHandler (Gin) and TopupActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func TopupActionCliHandler(
	handler func(c TopupActionRequest) (*TopupActionResponse, error),
) *cli.Command {
	meta := TopupActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: TopupActionCliFlags(),
	}
	cmd.Aliases = []string{meta.CliShort}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := TopupActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastTopupActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// TopupActionCli is a high-level convenience wrapper around
// TopupActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way TopupActionGin
// registers a route on a Gin engine.
func TopupActionCli(
	app *cli.Command,
	handler func(c TopupActionRequest) (*TopupActionResponse, error),
) {
	app.Commands = append(app.Commands, TopupActionCliHandler(handler))
}

// TopupActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the TopupAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *TopupActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func TopupActionHttpHandler(
	handler func(c TopupActionRequest) (*TopupActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := TopupActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body TopupActionReq
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
		req := TopupActionRequest{
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

// TopupActionHttp is a high-level convenience wrapper around
// TopupActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func TopupActionHttp(
	mux *http.ServeMux,
	handler func(c TopupActionRequest) (*TopupActionResponse, error),
) {
	method, pattern, h := TopupActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

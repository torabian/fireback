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
* Action to communicate with the action UserInvitationsAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of UserInvitationsAction
func UserInvitationsAction(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error) {
	return &UserInvitationsActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func UserInvitationsActionMeta() struct {
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
		Name:        "UserInvitationsAction",
		CliName:     "user-invitations-action",
		URL:         "/users/invitations",
		Method:      "GET",
		Description: `Shows the invitations for an specific user, if the invited member already has a account. It's based on the passports, so if the passport is authenticated we will show them.`,
	}
}

// The base class definition for userInvitationsActionRes
type UserInvitationsActionRes struct {
	// UserUniqueId
	UserId string `json:"userId" yaml:"userId"`
	// Invitation unique id
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// The value of the passport (email/phone)
	Value string `json:"value" yaml:"value"`
	// Name of the role that user will get
	RoleName string `json:"roleName" yaml:"roleName"`
	// Name of the workspace which user is invited to.
	WorkspaceName string `json:"workspaceName" yaml:"workspaceName"`
	// The method of the invitation, such as email.
	Type string `json:"type" yaml:"type"`
	// The content that user will receive to understand the reason of the letter.
	CoverLetter string `json:"coverLetter" yaml:"coverLetter"`
}

func (x *UserInvitationsActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserInvitationsActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "user-id",
			Type:        "string",
			Description: "UserUniqueId",
		},
		{
			Name:        prefix + "unique-id",
			Type:        "string",
			Description: "Invitation unique id",
		},
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "The value of the passport (email/phone)",
		},
		{
			Name:        prefix + "role-name",
			Type:        "string",
			Description: "Name of the role that user will get",
		},
		{
			Name:        prefix + "workspace-name",
			Type:        "string",
			Description: "Name of the workspace which user is invited to.",
		},
		{
			Name:        prefix + "type",
			Type:        "string",
			Description: "The method of the invitation, such as email.",
		},
		{
			Name:        prefix + "cover-letter",
			Type:        "string",
			Description: "The content that user will receive to understand the reason of the letter.",
		},
	}
}
func CastUserInvitationsActionResFromCli(c emigo.CliCastable) UserInvitationsActionRes {
	data := UserInvitationsActionRes{}
	if c.IsSet("user-id") {
		data.UserId = c.String("user-id")
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("role-name") {
		data.RoleName = c.String("role-name")
	}
	if c.IsSet("workspace-name") {
		data.WorkspaceName = c.String("workspace-name")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("cover-letter") {
		data.CoverLetter = c.String("cover-letter")
	}
	return data
}

type UserInvitationsActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *UserInvitationsActionResponse) SetContentType(contentType string) *UserInvitationsActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *UserInvitationsActionResponse) AsStream(r io.Reader, contentType string) *UserInvitationsActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *UserInvitationsActionResponse) AsJSON(payload any) *UserInvitationsActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *UserInvitationsActionResponse) WithIdeal(payload UserInvitationsActionRes) *UserInvitationsActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *UserInvitationsActionResponse) AsIdeal() (*UserInvitationsActionRes, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res UserInvitationsActionRes
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *UserInvitationsActionResponse) AsHTML(payload string) *UserInvitationsActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *UserInvitationsActionResponse) AsBytes(payload []byte) *UserInvitationsActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x UserInvitationsActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x UserInvitationsActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x UserInvitationsActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type UserInvitationsActionRequestSig = func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error)

/**
 * Query parameters for UserInvitationsAction
 */
// Query wrapper with private fields
type UserInvitationsActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func UserInvitationsActionQueryFromString(rawQuery string) UserInvitationsActionQuery {
	v := UserInvitationsActionQuery{}
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
func UserInvitationsActionQueryFromHttp(r *http.Request) UserInvitationsActionQuery {
	return UserInvitationsActionQueryFromString(r.URL.RawQuery)
}
func (q UserInvitationsActionQuery) Values() url.Values {
	return q.values
}
func (q UserInvitationsActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *UserInvitationsActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *UserInvitationsActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type UserInvitationsActionRequest struct {
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
func (x UserInvitationsActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x UserInvitationsActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func UserInvitationsActionClientCreateUrl(
	req UserInvitationsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := UserInvitationsActionMeta()
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
func UserInvitationsActionClientExecuteTyped(httpReq *http.Request) (*UserInvitationsActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result UserInvitationsActionResponse
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
func UserInvitationsActionClientBuildRequest(req UserInvitationsActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := UserInvitationsActionMeta()
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
func UserInvitationsActionCall(
	req UserInvitationsActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*UserInvitationsActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := UserInvitationsActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := UserInvitationsActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return UserInvitationsActionClientExecuteTyped(r)
}

// UserInvitationsActionRaw registers a raw Gin route for the UserInvitationsAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func UserInvitationsActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := UserInvitationsActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// UserInvitationsActionHandler returns the HTTP method, route URL, and a typed Gin handler for the UserInvitationsAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func UserInvitationsActionHandler(
	handler func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := UserInvitationsActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := UserInvitationsActionRequest{
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

// UserInvitationsActionGin is a high-level convenience wrapper around UserInvitationsActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func UserInvitationsActionGin(r gin.IRoutes, handler func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error)) {
	method, url, h := UserInvitationsActionHandler(handler)
	r.Handle(method, url, h)
}
func (x UserInvitationsActionRequest) IsGin() bool {
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
func UserInvitationsActionQueryFromGin(c *gin.Context) UserInvitationsActionQuery {
	return UserInvitationsActionQueryFromString(c.Request.URL.RawQuery)
}
func (x UserInvitationsActionRequest) IsCli() bool {
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

// UserInvitationsActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the UserInvitationsAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func UserInvitationsActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// UserInvitationsActionCliHandler builds a full *cli.Command for the
// UserInvitationsAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a UserInvitationsActionRequest the same way
// UserInvitationsActionHandler (Gin) and UserInvitationsActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func UserInvitationsActionCliHandler(
	handler func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error),
) *cli.Command {
	meta := UserInvitationsActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: UserInvitationsActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := UserInvitationsActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// UserInvitationsActionCli is a high-level convenience wrapper around
// UserInvitationsActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way UserInvitationsActionGin
// registers a route on a Gin engine.
func UserInvitationsActionCli(
	app *cli.Command,
	handler func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error),
) {
	app.Commands = append(app.Commands, UserInvitationsActionCliHandler(handler))
}

// UserInvitationsActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the UserInvitationsAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *UserInvitationsActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func UserInvitationsActionHttpHandler(
	handler func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := UserInvitationsActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := UserInvitationsActionRequest{
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

// UserInvitationsActionHttp is a high-level convenience wrapper around
// UserInvitationsActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func UserInvitationsActionHttp(
	mux *http.ServeMux,
	handler func(c UserInvitationsActionRequest) (*UserInvitationsActionResponse, error),
) {
	method, pattern, h := UserInvitationsActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

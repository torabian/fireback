package abacdefs

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/**
* Action to communicate with the action NotificationConfigGetAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of NotificationConfigGetAction
func NotificationConfigGetAction(c NotificationConfigGetActionRequest) (*NotificationConfigGetActionResponse, error) {
	return &NotificationConfigGetActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func NotificationConfigGetActionMeta() struct {
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
		Name:        "NotificationConfigGetAction",
		CliName:     "get",
		CliShort:    "g",
		URL:         "/notificationConfig/:uniqueId",
		Method:      "GET",
		Description: `Looks up a single "notificationConfig" row by uniqueId.`,
	}
}

type NotificationConfigGetActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *NotificationConfigGetActionResponse) SetContentType(contentType string) *NotificationConfigGetActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *NotificationConfigGetActionResponse) AsStream(r io.Reader, contentType string) *NotificationConfigGetActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *NotificationConfigGetActionResponse) AsJSON(payload any) *NotificationConfigGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *NotificationConfigGetActionResponse) WithIdeal(payload NotificationConfigDto) *NotificationConfigGetActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *NotificationConfigGetActionResponse) AsIdeal() (*NotificationConfigDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res NotificationConfigDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *NotificationConfigGetActionResponse) AsHTML(payload string) *NotificationConfigGetActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *NotificationConfigGetActionResponse) AsBytes(payload []byte) *NotificationConfigGetActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x NotificationConfigGetActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x NotificationConfigGetActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x NotificationConfigGetActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type NotificationConfigGetActionRequestSig = func(c NotificationConfigGetActionRequest) (*NotificationConfigGetActionResponse, error)

/**
 * Path parameters for NotificationConfigGetAction
 */
type NotificationConfigGetActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func NotificationConfigGetActionPathParameterApply(params NotificationConfigGetActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func NotificationConfigGetActionPathParameterFromFn(fn func(key string) string) NotificationConfigGetActionPathParameter {
	res := NotificationConfigGetActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for NotificationConfigGetAction
 */
// Query wrapper with private fields
type NotificationConfigGetActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func NotificationConfigGetActionQueryFromString(rawQuery string) NotificationConfigGetActionQuery {
	v := NotificationConfigGetActionQuery{}
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
func NotificationConfigGetActionQueryFromHttp(r *http.Request) NotificationConfigGetActionQuery {
	return NotificationConfigGetActionQueryFromString(r.URL.RawQuery)
}
func (q NotificationConfigGetActionQuery) Values() url.Values {
	return q.values
}
func (q NotificationConfigGetActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *NotificationConfigGetActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *NotificationConfigGetActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type NotificationConfigGetActionRequest struct {
	Body        interface{}
	Params      NotificationConfigGetActionPathParameter
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
func (x NotificationConfigGetActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x NotificationConfigGetActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func NotificationConfigGetActionClientCreateUrl(
	req NotificationConfigGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := NotificationConfigGetActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = NotificationConfigGetActionPathParameterApply(req.Params, urlAddr)
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
func NotificationConfigGetActionClientExecuteTyped(httpReq *http.Request) (*NotificationConfigGetActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result NotificationConfigGetActionResponse
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
func NotificationConfigGetActionClientBuildRequest(req NotificationConfigGetActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := NotificationConfigGetActionMeta()
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
func NotificationConfigGetActionCall(
	req NotificationConfigGetActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*NotificationConfigGetActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := NotificationConfigGetActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := NotificationConfigGetActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return NotificationConfigGetActionClientExecuteTyped(r)
}

// NotificationConfigGetActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the NotificationConfigGetAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *NotificationConfigGetActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func NotificationConfigGetActionHttpHandler(
	handler func(c NotificationConfigGetActionRequest) (*NotificationConfigGetActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := NotificationConfigGetActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := NotificationConfigGetActionRequest{
			Body: nil,
			Params: NotificationConfigGetActionPathParameterFromFn(func(key string) string {
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

// NotificationConfigGetActionHttp is a high-level convenience wrapper around
// NotificationConfigGetActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func NotificationConfigGetActionHttp(
	mux *http.ServeMux,
	handler func(c NotificationConfigGetActionRequest) (*NotificationConfigGetActionResponse, error),
) {
	method, pattern, h := NotificationConfigGetActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

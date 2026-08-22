package abacdefs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/**
* Action to communicate with the action CapabilityUpdateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CapabilityUpdateAction
func CapabilityUpdateAction(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error) {
	return &CapabilityUpdateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CapabilityUpdateActionMeta() struct {
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
		Name:        "CapabilityUpdateAction",
		CliName:     "update",
		CliShort:    "u",
		URL:         "/capability/:uniqueId",
		Method:      "PATCH",
		Description: `Applies a partial update to a "capability" row by uniqueId.`,
	}
}

type CapabilityUpdateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CapabilityUpdateActionResponse) SetContentType(contentType string) *CapabilityUpdateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CapabilityUpdateActionResponse) AsStream(r io.Reader, contentType string) *CapabilityUpdateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CapabilityUpdateActionResponse) AsJSON(payload any) *CapabilityUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CapabilityUpdateActionResponse) WithIdeal(payload CapabilityDto) *CapabilityUpdateActionResponse {
	x.Payload = payload
	return x
}

// Use this for client calls, so the payload is being casted
func (x *CapabilityUpdateActionResponse) AsIdeal() (*CapabilityDto, error) {
	b, err := json.Marshal(x.GetPayload())
	if err != nil {
		return nil, err
	}
	var res CapabilityDto
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
func (x *CapabilityUpdateActionResponse) AsHTML(payload string) *CapabilityUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CapabilityUpdateActionResponse) AsBytes(payload []byte) *CapabilityUpdateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CapabilityUpdateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CapabilityUpdateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CapabilityUpdateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type CapabilityUpdateActionRequestSig = func(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error)

/**
 * Path parameters for CapabilityUpdateAction
 */
type CapabilityUpdateActionPathParameter struct {
	UniqueId string
}

// Converts a placeholder url, and applies the parameters to it.
func CapabilityUpdateActionPathParameterApply(params CapabilityUpdateActionPathParameter, templateUrl string) string {
	templateUrl = strings.ReplaceAll(templateUrl, ":uniqueId", fmt.Sprintf("%v", params.UniqueId))
	return templateUrl
}

// General purpose to extract the value and cast based on type.
func CapabilityUpdateActionPathParameterFromFn(fn func(key string) string) CapabilityUpdateActionPathParameter {
	res := CapabilityUpdateActionPathParameter{}
	res.UniqueId = fn("uniqueId")
	return res
}

/**
 * Query parameters for CapabilityUpdateAction
 */
// Query wrapper with private fields
type CapabilityUpdateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CapabilityUpdateActionQueryFromString(rawQuery string) CapabilityUpdateActionQuery {
	v := CapabilityUpdateActionQuery{}
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
func CapabilityUpdateActionQueryFromHttp(r *http.Request) CapabilityUpdateActionQuery {
	return CapabilityUpdateActionQueryFromString(r.URL.RawQuery)
}
func (q CapabilityUpdateActionQuery) Values() url.Values {
	return q.values
}
func (q CapabilityUpdateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CapabilityUpdateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CapabilityUpdateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CapabilityUpdateActionRequest struct {
	Body        CapabilityOptionalDto
	Params      CapabilityUpdateActionPathParameter
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
func (x CapabilityUpdateActionRequest) GetGinCtx() interface{} {
	return x.GinCtx
}

// Returns the urfave 3 cli context. You need to manullay cast to .(*cli.Command)
func (x CapabilityUpdateActionRequest) GetCliCtx() interface{} {
	return x.CliCtx
}
func CapabilityUpdateActionClientCreateUrl(
	req CapabilityUpdateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CapabilityUpdateActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// In case there is a path parameter, we need to apply that.
	urlAddr = CapabilityUpdateActionPathParameterApply(req.Params, urlAddr)
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
func CapabilityUpdateActionClientExecuteTyped(httpReq *http.Request) (*CapabilityUpdateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CapabilityUpdateActionResponse
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
func CapabilityUpdateActionClientBuildRequest(req CapabilityUpdateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CapabilityUpdateActionMeta()
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
func CapabilityUpdateActionCall(
	req CapabilityUpdateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CapabilityUpdateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CapabilityUpdateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CapabilityUpdateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CapabilityUpdateActionClientExecuteTyped(r)
}

// CapabilityUpdateActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the CapabilityUpdateAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *CapabilityUpdateActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func CapabilityUpdateActionHttpHandler(
	handler func(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := CapabilityUpdateActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body CapabilityOptionalDto
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
		req := CapabilityUpdateActionRequest{
			Body: body,
			Params: CapabilityUpdateActionPathParameterFromFn(func(key string) string {
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

// CapabilityUpdateActionHttp is a high-level convenience wrapper around
// CapabilityUpdateActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func CapabilityUpdateActionHttp(
	mux *http.ServeMux,
	handler func(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error),
) {
	method, pattern, h := CapabilityUpdateActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

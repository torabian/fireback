package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
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
		CliName:     "browse",
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

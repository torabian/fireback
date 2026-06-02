package abac

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

/**
* Action to communicate with the action InviteToWorkspaceAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of InviteToWorkspaceAction
func InviteToWorkspaceAction(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error) {
	return &InviteToWorkspaceActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func InviteToWorkspaceActionMeta() struct {
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
		Name:        "InviteToWorkspaceAction",
		CliName:     "invite",
		URL:         "/workspace/invite",
		Method:      "POST",
		Description: `Invite a new person (either a user, with passport or without passport)`,
	}
}

type InviteToWorkspaceActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *InviteToWorkspaceActionResponse) SetContentType(contentType string) *InviteToWorkspaceActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *InviteToWorkspaceActionResponse) AsStream(r io.Reader, contentType string) *InviteToWorkspaceActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *InviteToWorkspaceActionResponse) AsJSON(payload any) *InviteToWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *InviteToWorkspaceActionResponse) AsHTML(payload string) *InviteToWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *InviteToWorkspaceActionResponse) AsBytes(payload []byte) *InviteToWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x InviteToWorkspaceActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x InviteToWorkspaceActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x InviteToWorkspaceActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type InviteToWorkspaceActionRequestSig = func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error)

/**
 * Query parameters for InviteToWorkspaceAction
 */
// Query wrapper with private fields
type InviteToWorkspaceActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func InviteToWorkspaceActionQueryFromString(rawQuery string) InviteToWorkspaceActionQuery {
	v := InviteToWorkspaceActionQuery{}
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
func InviteToWorkspaceActionQueryFromHttp(r *http.Request) InviteToWorkspaceActionQuery {
	return InviteToWorkspaceActionQueryFromString(r.URL.RawQuery)
}
func (q InviteToWorkspaceActionQuery) Values() url.Values {
	return q.values
}
func (q InviteToWorkspaceActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *InviteToWorkspaceActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *InviteToWorkspaceActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type InviteToWorkspaceActionRequest struct {
	Body        WorkspaceInvitationDto
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

func InviteToWorkspaceActionClientCreateUrl(
	req InviteToWorkspaceActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := InviteToWorkspaceActionMeta()
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
func InviteToWorkspaceActionClientExecuteTyped(httpReq *http.Request) (*InviteToWorkspaceActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result InviteToWorkspaceActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &InviteToWorkspaceActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &InviteToWorkspaceActionResponse{Payload: result}, err
	}
	return &InviteToWorkspaceActionResponse{Payload: result}, nil
}
func InviteToWorkspaceActionClientBuildRequest(req InviteToWorkspaceActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := InviteToWorkspaceActionMeta()
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
func InviteToWorkspaceActionCall(
	req InviteToWorkspaceActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*InviteToWorkspaceActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := InviteToWorkspaceActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := InviteToWorkspaceActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return InviteToWorkspaceActionClientExecuteTyped(r)
}

// InviteToWorkspaceActionRaw registers a raw Gin route for the InviteToWorkspaceAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func InviteToWorkspaceActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := InviteToWorkspaceActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// InviteToWorkspaceActionHandler returns the HTTP method, route URL, and a typed Gin handler for the InviteToWorkspaceAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func InviteToWorkspaceActionHandler(
	handler func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := InviteToWorkspaceActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body WorkspaceInvitationDto
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := InviteToWorkspaceActionRequest{
			Body:        body,
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
		resp, err := handler(req)
		if err != nil {
			m.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

// InviteToWorkspaceActionGin is a high-level convenience wrapper around InviteToWorkspaceActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func InviteToWorkspaceActionGin(r gin.IRoutes, handler func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error)) {
	method, url, h := InviteToWorkspaceActionHandler(handler)
	r.Handle(method, url, h)
}
func (x InviteToWorkspaceActionRequest) IsGin() bool {
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
func InviteToWorkspaceActionQueryFromGin(c *gin.Context) InviteToWorkspaceActionQuery {
	return InviteToWorkspaceActionQueryFromString(c.Request.URL.RawQuery)
}
func (x InviteToWorkspaceActionRequest) IsCli() bool {
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

// InviteToWorkspaceActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the InviteToWorkspaceAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *InviteToWorkspaceActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func InviteToWorkspaceActionHttpHandler(
	handler func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := InviteToWorkspaceActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body WorkspaceInvitationDto
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
		req := InviteToWorkspaceActionRequest{
			Body:        body,
			QueryParams: r.URL.Query(),
			Headers:     r.Header,
		}
		resp, err := handler(req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
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

// InviteToWorkspaceActionHttp is a high-level convenience wrapper around
// InviteToWorkspaceActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func InviteToWorkspaceActionHttp(
	mux *http.ServeMux,
	handler func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error),
) {
	method, pattern, h := InviteToWorkspaceActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}

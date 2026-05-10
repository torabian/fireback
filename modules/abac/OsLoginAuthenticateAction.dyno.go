package abac

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action OsLoginAuthenticateAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of OsLoginAuthenticateAction
func OsLoginAuthenticateAction(c OsLoginAuthenticateActionRequest) (*OsLoginAuthenticateActionResponse, error) {
	return &OsLoginAuthenticateActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func OsLoginAuthenticateActionMeta() struct {
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
		Name:        "OsLoginAuthenticateAction",
		CliName:     "oslogin",
		URL:         "/passports/os/login",
		Method:      "GET",
		Description: `Logins into the system using operating system (current) user, and store the information for them. Useful for desktop applications.`,
	}
}

type OsLoginAuthenticateActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *OsLoginAuthenticateActionResponse) SetContentType(contentType string) *OsLoginAuthenticateActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *OsLoginAuthenticateActionResponse) AsStream(r io.Reader, contentType string) *OsLoginAuthenticateActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *OsLoginAuthenticateActionResponse) AsJSON(payload any) *OsLoginAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *OsLoginAuthenticateActionResponse) AsHTML(payload string) *OsLoginAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *OsLoginAuthenticateActionResponse) AsBytes(payload []byte) *OsLoginAuthenticateActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x OsLoginAuthenticateActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x OsLoginAuthenticateActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x OsLoginAuthenticateActionResponse) GetPayload() interface{} {
	return x.Payload
}

// OsLoginAuthenticateActionRaw registers a raw Gin route for the OsLoginAuthenticateAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func OsLoginAuthenticateActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := OsLoginAuthenticateActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type OsLoginAuthenticateActionRequestSig = func(c OsLoginAuthenticateActionRequest) (*OsLoginAuthenticateActionResponse, error)

// OsLoginAuthenticateActionHandler returns the HTTP method, route URL, and a typed Gin handler for the OsLoginAuthenticateAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func OsLoginAuthenticateActionHandler(
	handler OsLoginAuthenticateActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := OsLoginAuthenticateActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := OsLoginAuthenticateActionRequest{
			Body:        nil,
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

// OsLoginAuthenticateAction is a high-level convenience wrapper around OsLoginAuthenticateActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func OsLoginAuthenticateActionGin(r gin.IRoutes, handler OsLoginAuthenticateActionRequestSig) {
	method, url, h := OsLoginAuthenticateActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for OsLoginAuthenticateAction
 */
// Query wrapper with private fields
type OsLoginAuthenticateActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func OsLoginAuthenticateActionQueryFromString(rawQuery string) OsLoginAuthenticateActionQuery {
	v := OsLoginAuthenticateActionQuery{}
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
func OsLoginAuthenticateActionQueryFromGin(c *gin.Context) OsLoginAuthenticateActionQuery {
	return OsLoginAuthenticateActionQueryFromString(c.Request.URL.RawQuery)
}
func OsLoginAuthenticateActionQueryFromHttp(r *http.Request) OsLoginAuthenticateActionQuery {
	return OsLoginAuthenticateActionQueryFromString(r.URL.RawQuery)
}
func (q OsLoginAuthenticateActionQuery) Values() url.Values {
	return q.values
}
func (q OsLoginAuthenticateActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *OsLoginAuthenticateActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *OsLoginAuthenticateActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type OsLoginAuthenticateActionRequest struct {
	Body        interface{}
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	GinCtx *gin.Context
	// Urfave context, per each request
	CliCtx *cli.Command
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func (x OsLoginAuthenticateActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x OsLoginAuthenticateActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type OsLoginAuthenticateActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func OsLoginAuthenticateActionClientCreateUrl(
	req OsLoginAuthenticateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := OsLoginAuthenticateActionMeta()
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
func OsLoginAuthenticateActionClientExecuteTyped(httpReq *http.Request) (*OsLoginAuthenticateActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result OsLoginAuthenticateActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &OsLoginAuthenticateActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &OsLoginAuthenticateActionResponse{Payload: result}, err
	}
	return &OsLoginAuthenticateActionResponse{Payload: result}, nil
}
func OsLoginAuthenticateActionClientBuildRequest(req OsLoginAuthenticateActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := OsLoginAuthenticateActionMeta()
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
func OsLoginAuthenticateActionCall(
	req OsLoginAuthenticateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*OsLoginAuthenticateActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := OsLoginAuthenticateActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := OsLoginAuthenticateActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return OsLoginAuthenticateActionClientExecuteTyped(r)
}

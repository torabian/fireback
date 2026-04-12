package abac

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action OsLoginAuthenticateAction
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
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type OsLoginAuthenticateActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func OsLoginAuthenticateActionCall(
	req OsLoginAuthenticateActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*OsLoginAuthenticateActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := OsLoginAuthenticateActionMeta()
		baseURL := meta.URL
		// Build final URL with query string
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		// if UrlValues present, encode and append
		if len(req.QueryParams) > 0 {
			u.RawQuery = req.QueryParams.Encode()
		}
		req0, err := http.NewRequest(meta.Method, u.String(), nil)
		if err != nil {
			return nil, err
		}
		httpReq = req0
	} else {
		httpReq = config.Httpr
	}
	httpReq.Header = req.Headers
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var result OsLoginAuthenticateActionResult
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if resp.StatusCode >= 400 {
		return &result, fmt.Errorf("request failed: %s", respBody)
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}

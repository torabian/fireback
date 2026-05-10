package abac

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
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

// InviteToWorkspaceActionRaw registers a raw Gin route for the InviteToWorkspaceAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func InviteToWorkspaceActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := InviteToWorkspaceActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type InviteToWorkspaceActionRequestSig = func(c InviteToWorkspaceActionRequest) (*InviteToWorkspaceActionResponse, error)

// InviteToWorkspaceActionHandler returns the HTTP method, route URL, and a typed Gin handler for the InviteToWorkspaceAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func InviteToWorkspaceActionHandler(
	handler InviteToWorkspaceActionRequestSig,
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

// InviteToWorkspaceAction is a high-level convenience wrapper around InviteToWorkspaceActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func InviteToWorkspaceActionGin(r gin.IRoutes, handler InviteToWorkspaceActionRequestSig) {
	method, url, h := InviteToWorkspaceActionHandler(handler)
	r.Handle(method, url, h)
}

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
func InviteToWorkspaceActionQueryFromGin(c *gin.Context) InviteToWorkspaceActionQuery {
	return InviteToWorkspaceActionQueryFromString(c.Request.URL.RawQuery)
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
	GinCtx *gin.Context
	// Urfave context, per each request
	CliCtx *cli.Command
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func (x InviteToWorkspaceActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x InviteToWorkspaceActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type InviteToWorkspaceActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
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

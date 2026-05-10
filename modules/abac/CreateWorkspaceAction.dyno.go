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
* Action to communicate with the action CreateWorkspaceAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CreateWorkspaceAction
func CreateWorkspaceAction(c CreateWorkspaceActionRequest) (*CreateWorkspaceActionResponse, error) {
	return &CreateWorkspaceActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CreateWorkspaceActionMeta() struct {
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
		Name:        "CreateWorkspaceAction",
		CliName:     "create-workspace-action",
		URL:         "/workspaces/create",
		Method:      "POST",
		Description: ``,
	}
}
func GetCreateWorkspaceActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateWorkspaceActionReqFromCli(c emigo.CliCastable) CreateWorkspaceActionReq {
	data := CreateWorkspaceActionReq{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for createWorkspaceActionReq
type CreateWorkspaceActionReq struct {
	Name string `json:"name" yaml:"name"`
}

func (x *CreateWorkspaceActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCreateWorkspaceActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "workspace-id",
			Type: "string",
		},
	}
}
func CastCreateWorkspaceActionResFromCli(c emigo.CliCastable) CreateWorkspaceActionRes {
	data := CreateWorkspaceActionRes{}
	if c.IsSet("workspace-id") {
		data.WorkspaceId = c.String("workspace-id")
	}
	return data
}

// The base class definition for createWorkspaceActionRes
type CreateWorkspaceActionRes struct {
	WorkspaceId string `json:"workspaceId" yaml:"workspaceId"`
}

func (x *CreateWorkspaceActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CreateWorkspaceActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CreateWorkspaceActionResponse) SetContentType(contentType string) *CreateWorkspaceActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CreateWorkspaceActionResponse) AsStream(r io.Reader, contentType string) *CreateWorkspaceActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CreateWorkspaceActionResponse) AsJSON(payload any) *CreateWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CreateWorkspaceActionResponse) WithIdeal(payload CreateWorkspaceActionRes) *CreateWorkspaceActionResponse {
	x.Payload = payload
	return x
}
func (x *CreateWorkspaceActionResponse) AsHTML(payload string) *CreateWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CreateWorkspaceActionResponse) AsBytes(payload []byte) *CreateWorkspaceActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CreateWorkspaceActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CreateWorkspaceActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CreateWorkspaceActionResponse) GetPayload() interface{} {
	return x.Payload
}

// CreateWorkspaceActionRaw registers a raw Gin route for the CreateWorkspaceAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CreateWorkspaceActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CreateWorkspaceActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type CreateWorkspaceActionRequestSig = func(c CreateWorkspaceActionRequest) (*CreateWorkspaceActionResponse, error)

// CreateWorkspaceActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CreateWorkspaceAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CreateWorkspaceActionHandler(
	handler CreateWorkspaceActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := CreateWorkspaceActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body CreateWorkspaceActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := CreateWorkspaceActionRequest{
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

// CreateWorkspaceAction is a high-level convenience wrapper around CreateWorkspaceActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CreateWorkspaceActionGin(r gin.IRoutes, handler CreateWorkspaceActionRequestSig) {
	method, url, h := CreateWorkspaceActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for CreateWorkspaceAction
 */
// Query wrapper with private fields
type CreateWorkspaceActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CreateWorkspaceActionQueryFromString(rawQuery string) CreateWorkspaceActionQuery {
	v := CreateWorkspaceActionQuery{}
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
func CreateWorkspaceActionQueryFromGin(c *gin.Context) CreateWorkspaceActionQuery {
	return CreateWorkspaceActionQueryFromString(c.Request.URL.RawQuery)
}
func CreateWorkspaceActionQueryFromHttp(r *http.Request) CreateWorkspaceActionQuery {
	return CreateWorkspaceActionQueryFromString(r.URL.RawQuery)
}
func (q CreateWorkspaceActionQuery) Values() url.Values {
	return q.values
}
func (q CreateWorkspaceActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CreateWorkspaceActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CreateWorkspaceActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CreateWorkspaceActionRequest struct {
	Body        CreateWorkspaceActionReq
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

func (x CreateWorkspaceActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x CreateWorkspaceActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type CreateWorkspaceActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func CreateWorkspaceActionClientCreateUrl(
	req CreateWorkspaceActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CreateWorkspaceActionMeta()
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
func CreateWorkspaceActionClientExecuteTyped(httpReq *http.Request) (*CreateWorkspaceActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CreateWorkspaceActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CreateWorkspaceActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &CreateWorkspaceActionResponse{Payload: result}, err
	}
	return &CreateWorkspaceActionResponse{Payload: result}, nil
}
func CreateWorkspaceActionClientBuildRequest(req CreateWorkspaceActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CreateWorkspaceActionMeta()
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
func CreateWorkspaceActionCall(
	req CreateWorkspaceActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CreateWorkspaceActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CreateWorkspaceActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CreateWorkspaceActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CreateWorkspaceActionClientExecuteTyped(r)
}

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
* Action to communicate with the action QueryUserRoleWorkspacesAction
 */
func QueryUserRoleWorkspacesActionMeta() struct {
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
		Name:        "QueryUserRoleWorkspacesAction",
		CliName:     "urw",
		URL:         "/urw/query",
		Method:      "GET",
		Description: `Returns the workspaces that user belongs to, as well as his role in there, and the permissions for each role`,
	}
}
func GetQueryUserRoleWorkspacesActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "capabilities",
			Type: "slice",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "roles",
			Type: "array",
		},
	}
}
func CastQueryUserRoleWorkspacesActionResFromCli(c emigo.CliCastable) QueryUserRoleWorkspacesActionRes {
	data := QueryUserRoleWorkspacesActionRes{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("capabilities") {
		emigo.InflatePossibleSlice(c.String("capabilities"), &data.Capabilities)
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("roles") {
		data.Roles = emigo.CapturePossibleArray(CastQueryUserRoleWorkspacesActionResRolesFromCli, "roles", c)
	}
	return data
}

// The base class definition for queryUserRoleWorkspacesActionRes
type QueryUserRoleWorkspacesActionRes struct {
	Name string `json:"name" yaml:"name"`
	// Workspace level capabilities which are available
	Capabilities []string                                `json:"capabilities" yaml:"capabilities"`
	UniqueId     string                                  `json:"uniqueId" yaml:"uniqueId"`
	Roles        []QueryUserRoleWorkspacesActionResRoles `json:"roles" yaml:"roles"`
}

func GetQueryUserRoleWorkspacesActionResRolesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "capabilities",
			Type: "slice",
		},
	}
}
func CastQueryUserRoleWorkspacesActionResRolesFromCli(c emigo.CliCastable) QueryUserRoleWorkspacesActionResRoles {
	data := QueryUserRoleWorkspacesActionResRoles{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("capabilities") {
		emigo.InflatePossibleSlice(c.String("capabilities"), &data.Capabilities)
	}
	return data
}

// The base class definition for roles
type QueryUserRoleWorkspacesActionResRoles struct {
	Name     string `json:"name" yaml:"name"`
	UniqueId string `json:"uniqueId" yaml:"uniqueId"`
	// Capabilities related to this role which are available
	Capabilities []string `json:"capabilities" yaml:"capabilities"`
}

func (x *QueryUserRoleWorkspacesActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type QueryUserRoleWorkspacesActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *QueryUserRoleWorkspacesActionResponse) SetContentType(contentType string) *QueryUserRoleWorkspacesActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *QueryUserRoleWorkspacesActionResponse) AsStream(r io.Reader, contentType string) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *QueryUserRoleWorkspacesActionResponse) AsJSON(payload any) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *QueryUserRoleWorkspacesActionResponse) WithIdeal(payload QueryUserRoleWorkspacesActionRes) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = payload
	return x
}
func (x *QueryUserRoleWorkspacesActionResponse) AsHTML(payload string) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *QueryUserRoleWorkspacesActionResponse) AsBytes(payload []byte) *QueryUserRoleWorkspacesActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x QueryUserRoleWorkspacesActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x QueryUserRoleWorkspacesActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x QueryUserRoleWorkspacesActionResponse) GetPayload() interface{} {
	return x.Payload
}

// QueryUserRoleWorkspacesActionRaw registers a raw Gin route for the QueryUserRoleWorkspacesAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func QueryUserRoleWorkspacesActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := QueryUserRoleWorkspacesActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type QueryUserRoleWorkspacesActionRequestSig = func(c QueryUserRoleWorkspacesActionRequest) (*QueryUserRoleWorkspacesActionResponse, error)

// QueryUserRoleWorkspacesActionHandler returns the HTTP method, route URL, and a typed Gin handler for the QueryUserRoleWorkspacesAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func QueryUserRoleWorkspacesActionHandler(
	handler QueryUserRoleWorkspacesActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := QueryUserRoleWorkspacesActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := QueryUserRoleWorkspacesActionRequest{
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

// QueryUserRoleWorkspacesAction is a high-level convenience wrapper around QueryUserRoleWorkspacesActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func QueryUserRoleWorkspacesActionGin(r gin.IRoutes, handler QueryUserRoleWorkspacesActionRequestSig) {
	method, url, h := QueryUserRoleWorkspacesActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for QueryUserRoleWorkspacesAction
 */
// Query wrapper with private fields
type QueryUserRoleWorkspacesActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func QueryUserRoleWorkspacesActionQueryFromString(rawQuery string) QueryUserRoleWorkspacesActionQuery {
	v := QueryUserRoleWorkspacesActionQuery{}
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
func QueryUserRoleWorkspacesActionQueryFromGin(c *gin.Context) QueryUserRoleWorkspacesActionQuery {
	return QueryUserRoleWorkspacesActionQueryFromString(c.Request.URL.RawQuery)
}
func QueryUserRoleWorkspacesActionQueryFromHttp(r *http.Request) QueryUserRoleWorkspacesActionQuery {
	return QueryUserRoleWorkspacesActionQueryFromString(r.URL.RawQuery)
}
func (q QueryUserRoleWorkspacesActionQuery) Values() url.Values {
	return q.values
}
func (q QueryUserRoleWorkspacesActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *QueryUserRoleWorkspacesActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *QueryUserRoleWorkspacesActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type QueryUserRoleWorkspacesActionRequest struct {
	Body        interface{}
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type QueryUserRoleWorkspacesActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func QueryUserRoleWorkspacesActionCall(
	req QueryUserRoleWorkspacesActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*QueryUserRoleWorkspacesActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := QueryUserRoleWorkspacesActionMeta()
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
	var result QueryUserRoleWorkspacesActionResult
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

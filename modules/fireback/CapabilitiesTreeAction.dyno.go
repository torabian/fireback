package fireback

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
* Action to communicate with the action CapabilitiesTreeAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CapabilitiesTreeAction
func CapabilitiesTreeAction(c CapabilitiesTreeActionRequest) (*CapabilitiesTreeActionResponse, error) {
	return &CapabilitiesTreeActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CapabilitiesTreeActionMeta() struct {
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
		Name:        "CapabilitiesTreeAction",
		CliName:     "treex",
		URL:         "/capabilitiesTree",
		Method:      "GET",
		Description: `dLists all of the capabilities in database as a array of string as root access`,
	}
}
func GetCapabilitiesTreeActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "capabilities",
			Type: "collection",
		},
		{
			Name: prefix + "nested",
			Type: "collection",
		},
	}
}
func CastCapabilitiesTreeActionResFromCli(c emigo.CliCastable) CapabilitiesTreeActionRes {
	data := CapabilitiesTreeActionRes{}
	return data
}

// The base class definition for capabilitiesTreeActionRes
type CapabilitiesTreeActionRes struct {
	Capabilities []CapabilityInfoDto `json:"capabilities" yaml:"capabilities"`
	Nested       []CapabilityInfoDto `json:"nested" yaml:"nested"`
}

func (x *CapabilitiesTreeActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CapabilitiesTreeActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CapabilitiesTreeActionResponse) SetContentType(contentType string) *CapabilitiesTreeActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CapabilitiesTreeActionResponse) AsStream(r io.Reader, contentType string) *CapabilitiesTreeActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CapabilitiesTreeActionResponse) AsJSON(payload any) *CapabilitiesTreeActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *CapabilitiesTreeActionResponse) WithIdeal(payload CapabilitiesTreeActionRes) *CapabilitiesTreeActionResponse {
	x.Payload = payload
	return x
}
func (x *CapabilitiesTreeActionResponse) AsHTML(payload string) *CapabilitiesTreeActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CapabilitiesTreeActionResponse) AsBytes(payload []byte) *CapabilitiesTreeActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CapabilitiesTreeActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CapabilitiesTreeActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CapabilitiesTreeActionResponse) GetPayload() interface{} {
	return x.Payload
}

// CapabilitiesTreeActionRaw registers a raw Gin route for the CapabilitiesTreeAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CapabilitiesTreeActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CapabilitiesTreeActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type CapabilitiesTreeActionRequestSig = func(c CapabilitiesTreeActionRequest) (*CapabilitiesTreeActionResponse, error)

// CapabilitiesTreeActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CapabilitiesTreeAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CapabilitiesTreeActionHandler(
	handler CapabilitiesTreeActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := CapabilitiesTreeActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := CapabilitiesTreeActionRequest{
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

// CapabilitiesTreeAction is a high-level convenience wrapper around CapabilitiesTreeActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CapabilitiesTreeActionGin(r gin.IRoutes, handler CapabilitiesTreeActionRequestSig) {
	method, url, h := CapabilitiesTreeActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for CapabilitiesTreeAction
 */
// Query wrapper with private fields
type CapabilitiesTreeActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CapabilitiesTreeActionQueryFromString(rawQuery string) CapabilitiesTreeActionQuery {
	v := CapabilitiesTreeActionQuery{}
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
func CapabilitiesTreeActionQueryFromGin(c *gin.Context) CapabilitiesTreeActionQuery {
	return CapabilitiesTreeActionQueryFromString(c.Request.URL.RawQuery)
}
func CapabilitiesTreeActionQueryFromHttp(r *http.Request) CapabilitiesTreeActionQuery {
	return CapabilitiesTreeActionQueryFromString(r.URL.RawQuery)
}
func (q CapabilitiesTreeActionQuery) Values() url.Values {
	return q.values
}
func (q CapabilitiesTreeActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CapabilitiesTreeActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CapabilitiesTreeActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CapabilitiesTreeActionRequest struct {
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

func (x CapabilitiesTreeActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x CapabilitiesTreeActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type CapabilitiesTreeActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func CapabilitiesTreeActionClientCreateUrl(
	req CapabilitiesTreeActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CapabilitiesTreeActionMeta()
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
func CapabilitiesTreeActionClientExecuteTyped(httpReq *http.Request) (*CapabilitiesTreeActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CapabilitiesTreeActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CapabilitiesTreeActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &CapabilitiesTreeActionResponse{Payload: result}, err
	}
	return &CapabilitiesTreeActionResponse{Payload: result}, nil
}
func CapabilitiesTreeActionClientBuildRequest(req CapabilitiesTreeActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CapabilitiesTreeActionMeta()
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
func CapabilitiesTreeActionCall(
	req CapabilitiesTreeActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CapabilitiesTreeActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CapabilitiesTreeActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CapabilitiesTreeActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CapabilitiesTreeActionClientExecuteTyped(r)
}

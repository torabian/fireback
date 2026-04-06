package fireback

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
* Action to communicate with the action CapabilitiesTreeAction
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
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type CapabilitiesTreeActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func CapabilitiesTreeActionCall(
	req CapabilitiesTreeActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CapabilitiesTreeActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := CapabilitiesTreeActionMeta()
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
	var result CapabilitiesTreeActionResult
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

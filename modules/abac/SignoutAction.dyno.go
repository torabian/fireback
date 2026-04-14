package abac

import (
	"bytes"
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
* Action to communicate with the action SignoutAction
 */
func SignoutActionMeta() struct {
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
		Name:        "SignoutAction",
		CliName:     "signout-action",
		URL:         "/passport/signout",
		Method:      "POST",
		Description: `Signout the user, clears cookies or does anything else if needed.`,
	}
}
func GetSignoutActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "clear",
			Type: "bool?",
		},
	}
}
func CastSignoutActionReqFromCli(c emigo.CliCastable) SignoutActionReq {
	data := SignoutActionReq{}
	if c.IsSet("clear") {
		emigo.ParseNullable(c.String("clear"), &data.Clear)
	}
	return data
}

// The base class definition for signoutActionReq
type SignoutActionReq struct {
	Clear emigo.Nullable[bool] `json:"clear" yaml:"clear"`
}

func (x *SignoutActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSignoutActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "okay",
			Type: "bool",
		},
	}
}
func CastSignoutActionResFromCli(c emigo.CliCastable) SignoutActionRes {
	data := SignoutActionRes{}
	if c.IsSet("okay") {
		data.Okay = bool(c.Bool("okay"))
	}
	return data
}

// The base class definition for signoutActionRes
type SignoutActionRes struct {
	Okay bool `json:"okay" yaml:"okay"`
}

func (x *SignoutActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type SignoutActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *SignoutActionResponse) SetContentType(contentType string) *SignoutActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *SignoutActionResponse) AsStream(r io.Reader, contentType string) *SignoutActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *SignoutActionResponse) AsJSON(payload any) *SignoutActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *SignoutActionResponse) AsHTML(payload string) *SignoutActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *SignoutActionResponse) AsBytes(payload []byte) *SignoutActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x SignoutActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x SignoutActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x SignoutActionResponse) GetPayload() interface{} {
	return x.Payload
}

// SignoutActionRaw registers a raw Gin route for the SignoutAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func SignoutActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := SignoutActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type SignoutActionRequestSig = func(c SignoutActionRequest) (*SignoutActionResponse, error)

// SignoutActionHandler returns the HTTP method, route URL, and a typed Gin handler for the SignoutAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func SignoutActionHandler(
	handler SignoutActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := SignoutActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body SignoutActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := SignoutActionRequest{
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

// SignoutAction is a high-level convenience wrapper around SignoutActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func SignoutActionGin(r gin.IRoutes, handler SignoutActionRequestSig) {
	method, url, h := SignoutActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for SignoutAction
 */
// Query wrapper with private fields
type SignoutActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func SignoutActionQueryFromString(rawQuery string) SignoutActionQuery {
	v := SignoutActionQuery{}
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
func SignoutActionQueryFromGin(c *gin.Context) SignoutActionQuery {
	return SignoutActionQueryFromString(c.Request.URL.RawQuery)
}
func SignoutActionQueryFromHttp(r *http.Request) SignoutActionQuery {
	return SignoutActionQueryFromString(r.URL.RawQuery)
}
func (q SignoutActionQuery) Values() url.Values {
	return q.values
}
func (q SignoutActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *SignoutActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *SignoutActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type SignoutActionRequest struct {
	Body        SignoutActionReq
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type SignoutActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func SignoutActionCall(
	req SignoutActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*SignoutActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := SignoutActionMeta()
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
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		req0, err := http.NewRequest(meta.Method, u.String(), bytes.NewReader(bodyBytes))
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
	var result SignoutActionResult
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

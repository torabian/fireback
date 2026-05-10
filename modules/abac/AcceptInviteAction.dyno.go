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
* Action to communicate with the action AcceptInviteAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of AcceptInviteAction
func AcceptInviteAction(c AcceptInviteActionRequest) (*AcceptInviteActionResponse, error) {
	return &AcceptInviteActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func AcceptInviteActionMeta() struct {
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
		Name:        "AcceptInviteAction",
		CliName:     "accept-invite-action",
		URL:         "/user/invitation/accept",
		Method:      "POST",
		Description: `Use it when user accepts an invitation, and it will complete the joining process`,
	}
}
func GetAcceptInviteActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "invitation-unique-id",
			Type:        "string",
			Description: "The invitation id which will be used to process",
		},
	}
}
func CastAcceptInviteActionReqFromCli(c emigo.CliCastable) AcceptInviteActionReq {
	data := AcceptInviteActionReq{}
	if c.IsSet("invitation-unique-id") {
		data.InvitationUniqueId = c.String("invitation-unique-id")
	}
	return data
}

// The base class definition for acceptInviteActionReq
type AcceptInviteActionReq struct {
	// The invitation id which will be used to process
	InvitationUniqueId string `json:"invitationUniqueId" validate:"required" yaml:"invitationUniqueId"`
}

func (x *AcceptInviteActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetAcceptInviteActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "accepted",
			Type: "bool",
		},
	}
}
func CastAcceptInviteActionResFromCli(c emigo.CliCastable) AcceptInviteActionRes {
	data := AcceptInviteActionRes{}
	if c.IsSet("accepted") {
		data.Accepted = bool(c.Bool("accepted"))
	}
	return data
}

// The base class definition for acceptInviteActionRes
type AcceptInviteActionRes struct {
	Accepted bool `json:"accepted" yaml:"accepted"`
}

func (x *AcceptInviteActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type AcceptInviteActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *AcceptInviteActionResponse) SetContentType(contentType string) *AcceptInviteActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *AcceptInviteActionResponse) AsStream(r io.Reader, contentType string) *AcceptInviteActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *AcceptInviteActionResponse) AsJSON(payload any) *AcceptInviteActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *AcceptInviteActionResponse) WithIdeal(payload AcceptInviteActionRes) *AcceptInviteActionResponse {
	x.Payload = payload
	return x
}
func (x *AcceptInviteActionResponse) AsHTML(payload string) *AcceptInviteActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *AcceptInviteActionResponse) AsBytes(payload []byte) *AcceptInviteActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x AcceptInviteActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x AcceptInviteActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x AcceptInviteActionResponse) GetPayload() interface{} {
	return x.Payload
}

// AcceptInviteActionRaw registers a raw Gin route for the AcceptInviteAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func AcceptInviteActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := AcceptInviteActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type AcceptInviteActionRequestSig = func(c AcceptInviteActionRequest) (*AcceptInviteActionResponse, error)

// AcceptInviteActionHandler returns the HTTP method, route URL, and a typed Gin handler for the AcceptInviteAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func AcceptInviteActionHandler(
	handler AcceptInviteActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := AcceptInviteActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body AcceptInviteActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := AcceptInviteActionRequest{
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

// AcceptInviteAction is a high-level convenience wrapper around AcceptInviteActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func AcceptInviteActionGin(r gin.IRoutes, handler AcceptInviteActionRequestSig) {
	method, url, h := AcceptInviteActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for AcceptInviteAction
 */
// Query wrapper with private fields
type AcceptInviteActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func AcceptInviteActionQueryFromString(rawQuery string) AcceptInviteActionQuery {
	v := AcceptInviteActionQuery{}
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
func AcceptInviteActionQueryFromGin(c *gin.Context) AcceptInviteActionQuery {
	return AcceptInviteActionQueryFromString(c.Request.URL.RawQuery)
}
func AcceptInviteActionQueryFromHttp(r *http.Request) AcceptInviteActionQuery {
	return AcceptInviteActionQueryFromString(r.URL.RawQuery)
}
func (q AcceptInviteActionQuery) Values() url.Values {
	return q.values
}
func (q AcceptInviteActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *AcceptInviteActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *AcceptInviteActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type AcceptInviteActionRequest struct {
	Body        AcceptInviteActionReq
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

func (x AcceptInviteActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x AcceptInviteActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type AcceptInviteActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func AcceptInviteActionClientCreateUrl(
	req AcceptInviteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := AcceptInviteActionMeta()
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
func AcceptInviteActionClientExecuteTyped(httpReq *http.Request) (*AcceptInviteActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result AcceptInviteActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &AcceptInviteActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &AcceptInviteActionResponse{Payload: result}, err
	}
	return &AcceptInviteActionResponse{Payload: result}, nil
}
func AcceptInviteActionClientBuildRequest(req AcceptInviteActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := AcceptInviteActionMeta()
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
func AcceptInviteActionCall(
	req AcceptInviteActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*AcceptInviteActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := AcceptInviteActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := AcceptInviteActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return AcceptInviteActionClientExecuteTyped(r)
}

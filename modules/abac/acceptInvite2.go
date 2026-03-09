package abac

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
)

/**
* Action to communicate with the action AcceptInvite2Action
 */
func AcceptInvite2ActionMeta() struct {
	Name    string
	CliName string
	URL     string
	Method  string
} {
	return struct {
		Name    string
		CliName string
		URL     string
		Method  string
	}{
		Name:    "AcceptInvite2Action",
		CliName: "accept-invite2-action",
		URL:     "/v2/user/invitation/accept",
		Method:  "POST",
	}
}
func GetAcceptInvite2ActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "invitation-unique-id",
			Type: "string",
		},
	}
}
func CastAcceptInvite2ActionReqFromCli(c emigo.CliCastable) AcceptInvite2ActionReq {
	data := AcceptInvite2ActionReq{}
	if c.IsSet("invitation-unique-id") {
		data.InvitationUniqueId = c.String("invitation-unique-id")
	}
	return data
}

// The base class definition for acceptInvite2ActionReq
type AcceptInvite2ActionReq struct {
	// The invitation id which will be used to process
	InvitationUniqueId string `validate:"required" json:"invitationUniqueId" yaml:"invitationUniqueId"`
}

func (x *AcceptInvite2ActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type AcceptInvite2ActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// AcceptInvite2ActionRaw registers a raw Gin route for the AcceptInvite2Action action.
// This gives the developer full control over middleware, handlers, and response handling.
func AcceptInvite2ActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := AcceptInvite2ActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // AcceptInvite2ActionHandler returns the HTTP method, route URL, and a typed Gin handler for the AcceptInvite2Action action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func AcceptInvite2ActionHandler(
	handler func(c AcceptInvite2ActionRequest, gin *gin.Context) (*AcceptInvite2ActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := AcceptInvite2ActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body AcceptInvite2ActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := AcceptInvite2ActionRequest{
			Body:        body,
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
		}
		resp, err := handler(req, m)
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

// AcceptInvite2Action is a high-level convenience wrapper around AcceptInvite2ActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func AcceptInvite2Action(r gin.IRoutes, handler func(c AcceptInvite2ActionRequest, gin *gin.Context) (*AcceptInvite2ActionResponse, error)) {
	method, url, h := AcceptInvite2ActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for AcceptInvite2Action
 */
// Query wrapper with private fields
type AcceptInvite2ActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func AcceptInvite2ActionQueryFromString(rawQuery string) AcceptInvite2ActionQuery {
	v := AcceptInvite2ActionQuery{}
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
func AcceptInvite2ActionQueryFromGin(c *gin.Context) AcceptInvite2ActionQuery {
	return AcceptInvite2ActionQueryFromString(c.Request.URL.RawQuery)
}
func AcceptInvite2ActionQueryFromHttp(r *http.Request) AcceptInvite2ActionQuery {
	return AcceptInvite2ActionQueryFromString(r.URL.RawQuery)
}
func (q AcceptInvite2ActionQuery) Values() url.Values {
	return q.values
}
func (q AcceptInvite2ActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *AcceptInvite2ActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *AcceptInvite2ActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type AcceptInvite2ActionRequest struct {
	Body        AcceptInvite2ActionReq
	QueryParams url.Values
	Headers     http.Header
}
type AcceptInvite2ActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func AcceptInvite2ActionCall(
	req AcceptInvite2ActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*AcceptInvite2ActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := AcceptInvite2ActionMeta()
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
	var result AcceptInvite2ActionResult
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

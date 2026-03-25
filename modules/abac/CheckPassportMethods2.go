package abac

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action CheckPassportMethods2Action
 */
func CheckPassportMethods2ActionMeta() struct {
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
		Name:    "CheckPassportMethods2Action",
		CliName: "check-passport-methods2-action",
		URL:     "/passports/available-methods2",
		Method:  "GET",
	}
}
func GetCheckPassportMethods2ActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "email",
			Type: "bool",
		},
		{
			Name: prefix + "phone",
			Type: "bool",
		},
		{
			Name: prefix + "google",
			Type: "bool",
		},
		{
			Name: prefix + "facebook",
			Type: "bool",
		},
		{
			Name: prefix + "google-o-auth-client-key",
			Type: "string",
		},
		{
			Name: prefix + "facebook-app-id",
			Type: "string",
		},
		{
			Name: prefix + "enabled-recaptcha2",
			Type: "bool",
		},
		{
			Name: prefix + "recaptcha2-client-key",
			Type: "string",
		},
	}
}
func CastCheckPassportMethods2ActionResFromCli(c emigo.CliCastable) CheckPassportMethods2ActionRes {
	data := CheckPassportMethods2ActionRes{}
	if c.IsSet("email") {
		data.Email = bool(c.Bool("email"))
	}
	if c.IsSet("phone") {
		data.Phone = bool(c.Bool("phone"))
	}
	if c.IsSet("google") {
		data.Google = bool(c.Bool("google"))
	}
	if c.IsSet("facebook") {
		data.Facebook = bool(c.Bool("facebook"))
	}
	if c.IsSet("google-o-auth-client-key") {
		data.GoogleOAuthClientKey = c.String("google-o-auth-client-key")
	}
	if c.IsSet("facebook-app-id") {
		data.FacebookAppId = c.String("facebook-app-id")
	}
	if c.IsSet("enabled-recaptcha2") {
		data.EnabledRecaptcha2 = bool(c.Bool("enabled-recaptcha2"))
	}
	if c.IsSet("recaptcha2-client-key") {
		data.Recaptcha2ClientKey = c.String("recaptcha2-client-key")
	}
	return data
}

// The base class definition for checkPassportMethods2ActionRes
type CheckPassportMethods2ActionRes struct {
	Email                bool   `json:"email" yaml:"email"`
	Phone                bool   `json:"phone" yaml:"phone"`
	Google               bool   `json:"google" yaml:"google"`
	Facebook             bool   `json:"facebook" yaml:"facebook"`
	GoogleOAuthClientKey string `json:"googleOAuthClientKey" yaml:"googleOAuthClientKey"`
	FacebookAppId        string `json:"facebookAppId" yaml:"facebookAppId"`
	EnabledRecaptcha2    bool   `json:"enabledRecaptcha2" yaml:"enabledRecaptcha2"`
	Recaptcha2ClientKey  string `json:"recaptcha2ClientKey" yaml:"recaptcha2ClientKey"`
}

func (x *CheckPassportMethods2ActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CheckPassportMethods2ActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// CheckPassportMethods2ActionRaw registers a raw Gin route for the CheckPassportMethods2Action action.
// This gives the developer full control over middleware, handlers, and response handling.
func CheckPassportMethods2ActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CheckPassportMethods2ActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type CheckPassportMethods2ActionRequestSig = func(c CheckPassportMethods2ActionRequest, gin *gin.Context) (*CheckPassportMethods2ActionResponse, error)

// CheckPassportMethods2ActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CheckPassportMethods2Action action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CheckPassportMethods2ActionHandler(
	handler CheckPassportMethods2ActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := CheckPassportMethods2ActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := CheckPassportMethods2ActionRequest{
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

// CheckPassportMethods2Action is a high-level convenience wrapper around CheckPassportMethods2ActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CheckPassportMethods2Action(r gin.IRoutes, handler CheckPassportMethods2ActionRequestSig) {
	method, url, h := CheckPassportMethods2ActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for CheckPassportMethods2Action
 */
// Query wrapper with private fields
type CheckPassportMethods2ActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CheckPassportMethods2ActionQueryFromString(rawQuery string) CheckPassportMethods2ActionQuery {
	v := CheckPassportMethods2ActionQuery{}
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
func CheckPassportMethods2ActionQueryFromGin(c *gin.Context) CheckPassportMethods2ActionQuery {
	return CheckPassportMethods2ActionQueryFromString(c.Request.URL.RawQuery)
}
func CheckPassportMethods2ActionQueryFromHttp(r *http.Request) CheckPassportMethods2ActionQuery {
	return CheckPassportMethods2ActionQueryFromString(r.URL.RawQuery)
}
func (q CheckPassportMethods2ActionQuery) Values() url.Values {
	return q.values
}
func (q CheckPassportMethods2ActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CheckPassportMethods2ActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CheckPassportMethods2ActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CheckPassportMethods2ActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
}
type CheckPassportMethods2ActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func CheckPassportMethods2ActionCall(
	req CheckPassportMethods2ActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CheckPassportMethods2ActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := CheckPassportMethods2ActionMeta()
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
	var result CheckPassportMethods2ActionResult
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

package abacdefs

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

func WorkspaceTypeGetActionPathParameterFromGin(g *gin.Context) WorkspaceTypeGetActionPathParameter {
	return WorkspaceTypeGetActionPathParameterFromFn(func(key string) string {
		return g.Param(key)
	})
}

// WorkspaceTypeGetActionRaw registers a raw Gin route for the WorkspaceTypeGetAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func WorkspaceTypeGetActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := WorkspaceTypeGetActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

// WorkspaceTypeGetActionHandler returns the HTTP method, route URL, and a typed Gin handler for the WorkspaceTypeGetAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func WorkspaceTypeGetActionHandler(
	handler func(c WorkspaceTypeGetActionRequest) (*WorkspaceTypeGetActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := WorkspaceTypeGetActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := WorkspaceTypeGetActionRequest{
			Body:        nil,
			Params:      WorkspaceTypeGetActionPathParameterFromGin(m),
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
		resp, err := handler(req)
		if err != nil {
			// Some deeper call inside handler (e.g. a security/authorization check
			// that rejects the request before the handler's own business logic ever
			// runs) may have already written and aborted the response itself - gin
			// tracks that on the ResponseWriter regardless of who did the writing.
			// Rendering the bubbled-up error on top of that would append a second,
			// invalid JSON body after the first.
			if m.Writer.Written() {
				return
			}
			status := http.StatusInternalServerError
			// If the error knows how to render itself for a given language (e.g.
			// fireback.IError, whose ferror.Error.ToPublicJSON resolves its
			// {"$": ..., "en": ..., "fa": ...} message map down to one string), let it -
			// picking the language the same way the rest of the app resolves it: the
			// "acceptLanguage" query param first, else the Accept-Language header, else
			// "en".
			if converter, ok := err.(interface {
				ToPublicJSON(lang string) ([]byte, int32)
			}); ok {
				lang := m.Query("acceptLanguage")
				if lang == "" {
					lang = m.GetHeader("Accept-Language")
					if i := strings.IndexAny(lang, ",;-"); i >= 0 {
						lang = lang[:i]
					}
					lang = strings.ToLower(strings.TrimSpace(lang))
				}
				if lang == "" {
					lang = "en"
				}
				body, code := converter.ToPublicJSON(lang)
				if code != 0 {
					status = int(code)
				}
				// Nest the resolved object under "error" (rather than writing it as the
				// bare response body) so every error shape - this one, the generic
				// forwarded-JSON one below, and the plain-string one - answers with the
				// same {"error": ...} envelope. json.RawMessage keeps body embedded as
				// real JSON instead of being re-escaped into a string.
				m.JSON(status, gin.H{"error": json.RawMessage(body)})
				return
			}
			// Otherwise, other action errors may still stringify themselves as an
			// indented JSON object via their Error() method. If that's what we got,
			// forward it nested under "error" as real JSON (optionally honoring its own
			// "httpCode" field for the response status) instead of re-escaping it into a
			// string, which is what plain errors still get.
			msg := err.Error()
			trimmed := strings.TrimSpace(msg)
			if strings.HasPrefix(trimmed, "{") && json.Valid([]byte(trimmed)) {
				var probe struct {
					HttpCode int32 `json:"httpCode"`
				}
				if uErr := json.Unmarshal([]byte(trimmed), &probe); uErr == nil && probe.HttpCode != 0 {
					status = int(probe.HttpCode)
				}
				m.JSON(status, gin.H{"error": json.RawMessage(trimmed)})
				return
			}
			m.JSON(status, gin.H{"error": msg})
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

// WorkspaceTypeGetActionGin is a high-level convenience wrapper around WorkspaceTypeGetActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func WorkspaceTypeGetActionGin(r gin.IRoutes, handler func(c WorkspaceTypeGetActionRequest) (*WorkspaceTypeGetActionResponse, error)) {
	method, url, h := WorkspaceTypeGetActionHandler(handler)
	r.Handle(method, url, h)
}
func (x WorkspaceTypeGetActionRequest) IsGin() bool {
	if x.GinCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.GinCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}
func WorkspaceTypeGetActionQueryFromGin(c *gin.Context) WorkspaceTypeGetActionQuery {
	return WorkspaceTypeGetActionQueryFromString(c.Request.URL.RawQuery)
}

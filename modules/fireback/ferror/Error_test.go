package ferror

import (
	"encoding/json"
	"testing"
)

// TestError_ToPublicJSON covers the language resolution the Emi-generated Gin/http
// action handlers rely on (see go-action-gin-render.go / go-action-http-render.go in the
// emi compiler): an Error's {"$": ..., "en": ..., "fa": ...} message map must collapse
// down to a single message string in the requested language, both at the top level and
// for each nested field error - not leak the whole language map to the client.
func TestError_ToPublicJSON(t *testing.T) {
	err := &Error{
		Message: ErrorItem{
			"$": "CannotCreateWorkspaceType", "en": "You cannot create workspace type due to some validation errors.", "fa": "خطای اعتبارسنجی.",
		},
		Errors: []*FieldError{
			{
				Location: "roleId",
				Message: &ErrorItem{
					"$": "RoleIsNecessary", "en": "Role needs to be defined and exist.", "fa": "نقش باید تعریف شود.",
				},
			},
		},
		HttpCode: 400,
	}

	t.Run("resolves the requested language", func(t *testing.T) {
		body, code := err.ToPublicJSON("fa")
		if code != 400 {
			t.Errorf("expected httpCode 400, got %d", code)
		}

		var pub PublicError
		if uErr := json.Unmarshal(body, &pub); uErr != nil {
			t.Fatalf("ToPublicJSON did not return valid JSON: %v\nbody: %s", uErr, body)
		}

		if pub.Message != "CannotCreateWorkspaceType" {
			t.Errorf("expected the machine-readable code to survive as Message, got %q", pub.Message)
		}
		if pub.MessageTranslated != "خطای اعتبارسنجی." {
			t.Errorf("expected the fa translation, got %q", pub.MessageTranslated)
		}
		if len(pub.Errors) != 1 || pub.Errors[0].MessageTranslated != "نقش باید تعریف شود." {
			t.Errorf("expected the nested field error translated to fa, got %+v", pub.Errors)
		}

		// The full language map must not leak into the public payload.
		if string(body) == err.Error() {
			t.Errorf("ToPublicJSON must not just re-emit the raw multi-language Error()")
		}
	})

	t.Run("falls back to en for an unknown language", func(t *testing.T) {
		body, _ := err.ToPublicJSON("xx")
		var pub PublicError
		if uErr := json.Unmarshal(body, &pub); uErr != nil {
			t.Fatalf("ToPublicJSON did not return valid JSON: %v\nbody: %s", uErr, body)
		}
		// ToPublicEndUser looks up r.Message[lang] directly with no fallback chain -
		// an unrecognized language code resolves to "", which is the existing,
		// unchanged ToPublicEndUser behavior this method is a thin wrapper over.
		if pub.MessageTranslated != "" {
			t.Errorf("expected an empty translation for an unrecognized language code, got %q", pub.MessageTranslated)
		}
	})

	t.Run("satisfies the duck-typed interface Emi's generated handlers assert against", func(t *testing.T) {
		var asError error = err
		if _, ok := asError.(interface {
			ToPublicJSON(lang string) ([]byte, int32)
		}); !ok {
			t.Fatal("expected *Error to satisfy the ToPublicJSON(lang string) ([]byte, int32) interface")
		}
	})

	t.Run("defaults the http code to 500 when unset", func(t *testing.T) {
		unset := &Error{Message: ErrorItem{"$": "X", "en": "x"}}
		_, code := unset.ToPublicJSON("en")
		if code != 500 {
			t.Errorf("expected a default httpCode of 500, got %d", code)
		}
	})
}

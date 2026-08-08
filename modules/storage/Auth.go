package storage

import "github.com/gin-gonic/gin"

// AuthContext is the identity resolved from an incoming request by
// StorageModuleConfig.Authenticate: who is making this request, which
// workspace they're acting in, and (best-effort, framework-specific) what
// access level they hold. It's persisted per upload (see FileRecord's
// UserId/WorkspaceId/AccessLevel) and fed to StorageModuleConfig.Quota to
// decide how much storage this identity is allowed to use.
type AuthContext struct {
	UserId      string
	WorkspaceId string
	AccessLevel string
}

// Anonymous is the zero-value AuthContext: what every request is treated as
// when StorageModuleConfig.Authenticate is nil. No owner is recorded on
// uploads made this way, Quota is never consulted, and (since there's no
// UserId to compare against) ownership checks let any request through -
// this is this module's original behavior, from before this mechanism
// existed.
var Anonymous = AuthContext{}

func (a AuthContext) isAnonymous() bool {
	return a.UserId == ""
}

// resolveAuth calls cfg.Authenticate, or returns Anonymous with no error if
// cfg.Authenticate is nil (authentication disabled).
func resolveAuth(c *gin.Context, cfg *StorageModuleConfig) (AuthContext, error) {
	if cfg == nil || cfg.Authenticate == nil {
		return Anonymous, nil
	}
	return cfg.Authenticate(c)
}

// authorizeOwner reports whether auth may act on an upload whose recorded
// owner is ownerId. A nil ownerId - no owner was recorded, e.g. the upload
// predates this mechanism or was made while Authenticate was nil - is open
// to anyone, matching this module's pre-auth behavior. Otherwise, only the
// matching, non-anonymous owner may act on it.
func authorizeOwner(auth AuthContext, ownerId *string) bool {
	if ownerId == nil {
		return true
	}
	return !auth.isAnonymous() && auth.UserId == *ownerId
}

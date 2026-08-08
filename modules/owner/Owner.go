// Package owner defines the minimal identity a request/event belongs to - the
// workspace and user it's scoped to - so packages that only need that identity (not
// the rest of fireback.QueryDSL, or fireback at all) can depend on this instead of the
// concrete QueryDSL type. See modules/fireback/QueryDSL.go for QueryDSL's own
// GetWorkspaceId/GetUserId methods, which satisfy this interface.
package owner

// Owner identifies who something belongs to: a workspace and, within it, a user.
type Owner interface {
	// GetWorkspaceId returns the unique-id of the workspace this belongs to.
	GetWorkspaceId() string

	// GetUserId returns the unique-id of the user this belongs to.
	GetUserId() string
}

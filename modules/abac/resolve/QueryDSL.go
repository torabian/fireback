package resolve

import (
	"encoding/json"

	"github.com/torabian/fireback/modules/fireback/application"
)

type MsgContext struct {

	// Automatically assigned UserId to the request after analising the token
	// This will be used to save each entity and determine the owner of the record
	UserId string `json:"-"`

	ResolveStrategy string `json:"-"`

	// This is the workspace which user is working inside, usually data belongs there
	WorkspaceId string `json:"-"`

	// Those capabilities which user has
	ActionRequires []application.PermissionInfo `json:"-"`

	// This is the capabilities that user has
	UserHas []string `json:"-"`

	UserAccessPerWorkspace *UserAccessPerWorkspaceDto `json:"-" yaml:"-"`

	// This is limitation of that workspace
	WorkspaceHas []string `json:"-"`

	InternalQuery string `json:"-"`
	Language      string `json:"-"`
}

func (x MsgContext) Json() string {
	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))

}

func (x MsgContext) GetLanguage() string {
	return x.Language
}

// GetWorkspaceId implements owner.Owner.
func (x MsgContext) GetWorkspaceId() string {
	return x.WorkspaceId
}

// GetUserId implements owner.Owner.
func (x MsgContext) GetUserId() string {
	return x.UserId
}

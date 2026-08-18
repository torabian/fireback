package fireback

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/application"
)

var ResolveStrategyPublic = "public"
var ResolveStrategyUser = "user"
var ResolveStrategyWorkspace = "workspace"

// Used for actions generally
type SecurityModel struct {
	// Only users which belong to root and actively selected the root workspace can
	// write to this entity from Fireback default functionality
	AllowOnRoot bool `json:"allowOnRoot,omitempty" yaml:"allowOnRoot,omitempty"`

	// Set of permissions which are required for this service.
	ActionRequires []application.PermissionInfo `json:"requires,omitempty" yaml:"requires,omitempty"`

	// Resolve strategy is by default on the workspace, you can change it by user
	// also. Be sure of the consequences
	ResolveStrategy string `json:"resolveStrategy,omitempty" yaml:"resolveStrategy,omitempty"`
}

// Used for defining the entity overall action permissions
type EntitySecurityModel struct {
	// Only users which belong to root and actively selected the root workspace can write to this entity from Fireback default functionality. Read mechanism won't be affected.
	WriteOnRoot *bool `json:"writeOnRoot,omitempty" yaml:"writeOnRoot,omitempty" jsonschema:"description=Only users which belong to root and actively selected the root workspace can write to this entity from Fireback default functionality. Read mechanism won't be affected."`

	// Only users which belong to root and actively selected the root workspace can read from entity from Fireback default functionality. Write mechanism is not affected.
	ReadOnRoot *bool `json:"readOnRoot,omitempty" yaml:"readOnRoot,omitempty" jsonschema:"description=Only users which belong to root and actively selected the root workspace can read from entity from Fireback default functionality. Write mechanism is not affected."`

	// Resolve strategy means that the content belongs either to workspace or user. It affects the query.
	ResolveStrategy *string `json:"resolveStrategy,omitempty" yaml:"resolveStrategy,omitempty" jsonschema:"enum=workspace,enum=user, description=Resolve strategy means that the content belongs either to workspace or user. It affects the query."`
}

func WithAuthorization(securityModel *SecurityModel) gin.HandlerFunc {
	return WithAuthorizationFn(securityModel)
}

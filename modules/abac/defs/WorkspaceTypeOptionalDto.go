package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for workspaceTypeOptionalDto
type WorkspaceTypeOptionalDto struct {
	UniqueId    emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Title       emigo.Nullable[string] `json:"title" validate:"required,omitempty,min=1,max=250" yaml:"title"`
	Description emigo.Nullable[string] `json:"description" yaml:"description"`
	// Unique, URL-safe identifier for this workspace type. Must start with "/", contain only lowercase a-z and dashes after that, and be at most 50 characters long. Format is enforced in ValidateTheWorkspaceTypeEntity (WorkspaceTypeActions.go); gorm:"unique" here is the DB-level backstop for the uniqueness half of that.
	Slug emigo.Nullable[string] `json:"slug" validate:"required,omitempty,min=2,max=50" yaml:"slug"`
	// The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.
	RoleId      emigo.Nullable[string]  `json:"roleId" validate:"required" yaml:"roleId"`
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId      emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceTypeOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for roleDto
type RoleDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Name     string                 `json:"name" validate:"required,omitempty,min=1,max=200" yaml:"name"`
	// The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).
	CapabilitiesListId complexes.JSON          `json:"capabilitiesListId" yaml:"capabilitiesListId"`
	IsDeletable        emigo.Nullable[bool]    `json:"isDeletable" yaml:"isDeletable"`
	IsUpdatable        emigo.Nullable[bool]    `json:"isUpdatable" yaml:"isUpdatable"`
	WorkspaceId        emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId             emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt          abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt          abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *RoleDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

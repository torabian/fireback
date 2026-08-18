package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for passportMethodDto
type PassportMethodDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Type     string                 `json:"type" validate:"oneof=email phone google facebook,required" yaml:"type"`
	// The region which would be using this method of passports for authentication. In Fireback open-source, only 'global' is available.
	Region string `json:"region" validate:"required,oneof=global" yaml:"region"`
	// Client key for those methods such as 'google' which require oauth client key
	ClientKey string `json:"clientKey" yaml:"clientKey"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PassportMethodDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

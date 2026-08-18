package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for emailConfirmationOptionalDto
type EmailConfirmationOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user this confirmation belongs to.
	UserId    emigo.Nullable[string] `json:"userId" yaml:"userId"`
	Status    emigo.Nullable[string] `json:"status" yaml:"status"`
	Email     emigo.Nullable[string] `json:"email" yaml:"email"`
	Key       emigo.Nullable[string] `json:"key" yaml:"key"`
	ExpiresAt emigo.Nullable[string] `json:"expiresAt" yaml:"expiresAt"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *EmailConfirmationOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

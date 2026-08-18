package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for emailConfirmationDto
type EmailConfirmationDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user this confirmation belongs to.
	UserId    emigo.Nullable[string] `json:"userId" yaml:"userId"`
	Status    string                 `json:"status" yaml:"status"`
	Email     string                 `json:"email" yaml:"email"`
	Key       string                 `json:"key" yaml:"key"`
	ExpiresAt string                 `json:"expiresAt" yaml:"expiresAt"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *EmailConfirmationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

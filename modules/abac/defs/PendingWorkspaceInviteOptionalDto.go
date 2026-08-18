package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for pendingWorkspaceInviteOptionalDto
type PendingWorkspaceInviteOptionalDto struct {
	UniqueId      emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Value         emigo.Nullable[string] `json:"value" yaml:"value"`
	Type          emigo.Nullable[string] `json:"type" yaml:"type"`
	CoverLetter   emigo.Nullable[string] `json:"coverLetter" yaml:"coverLetter"`
	WorkspaceName emigo.Nullable[string] `json:"workspaceName" yaml:"workspaceName"`
	// The unique-id of the role which invitee will get if they accept the request.
	RoleId emigo.Nullable[string] `json:"roleId" yaml:"roleId"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PendingWorkspaceInviteOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

package abac

import (
	"encoding"
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
func GetPendingWorkspaceInviteOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "value",
			Type: "string?",
		},
		{
			Name: prefix + "type",
			Type: "string?",
		},
		{
			Name: prefix + "cover-letter",
			Type: "string?",
		},
		{
			Name: prefix + "workspace-name",
			Type: "string?",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string?",
			Description: "The unique-id of the role which invitee will get if they accept the request.",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user which created/owns the record.",
		},
		{
			Name: prefix + "created-at",
			Type: "complex",
		},
		{
			Name: prefix + "updated-at",
			Type: "complex",
		},
	}
}
func CastPendingWorkspaceInviteOptionalDtoFromCli(c emigo.CliCastable) PendingWorkspaceInviteOptionalDto {
	data := PendingWorkspaceInviteOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("value") {
		emigo.ParseNullable(c.String("value"), &data.Value)
	}
	if c.IsSet("type") {
		emigo.ParseNullable(c.String("type"), &data.Type)
	}
	if c.IsSet("cover-letter") {
		emigo.ParseNullable(c.String("cover-letter"), &data.CoverLetter)
	}
	if c.IsSet("workspace-name") {
		emigo.ParseNullable(c.String("workspace-name"), &data.WorkspaceName)
	}
	if c.IsSet("role-id") {
		emigo.ParseNullable(c.String("role-id"), &data.RoleId)
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	if c.IsSet("updated-at") {
		if u, ok := any(&data.UpdatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("updated-at")))
		}
	}
	return data
}

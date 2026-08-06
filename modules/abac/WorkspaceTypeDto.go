package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for workspaceTypeDto
type WorkspaceTypeDto struct {
	UniqueId    emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Title       string                 `json:"title" yaml:"title"`
	Description string                 `json:"description" yaml:"description"`
	Slug        string                 `json:"slug" yaml:"slug"`
	// The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.
	RoleId      string                  `json:"roleId" yaml:"roleId"`
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId      emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceTypeDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWorkspaceTypeDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "title",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
		{
			Name: prefix + "slug",
			Type: "string",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string",
			Description: "The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.",
		},
		{
			Name: prefix + "workspace-id",
			Type: "string?",
		},
		{
			Name: prefix + "user-id",
			Type: "string?",
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
func CastWorkspaceTypeDtoFromCli(c emigo.CliCastable) WorkspaceTypeDto {
	data := WorkspaceTypeDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	if c.IsSet("slug") {
		data.Slug = c.String("slug")
	}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
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

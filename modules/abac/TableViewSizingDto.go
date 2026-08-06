package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for tableViewSizingDto
type TableViewSizingDto struct {
	UniqueId  emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	TableName string                 `json:"tableName" yaml:"tableName"`
	Sizes     string                 `json:"sizes" yaml:"sizes"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// The time the record was created. Populated automatically by gorm.
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	// The time the record was last updated. Populated automatically by gorm.
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *TableViewSizingDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetTableViewSizingDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "table-name",
			Type: "string",
		},
		{
			Name: prefix + "sizes",
			Type: "string",
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
			Name:        prefix + "created-at",
			Type:        "complex",
			Description: "The time the record was created. Populated automatically by gorm.",
		},
		{
			Name:        prefix + "updated-at",
			Type:        "complex",
			Description: "The time the record was last updated. Populated automatically by gorm.",
		},
	}
}
func CastTableViewSizingDtoFromCli(c emigo.CliCastable) TableViewSizingDto {
	data := TableViewSizingDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("table-name") {
		data.TableName = c.String("table-name")
	}
	if c.IsSet("sizes") {
		data.Sizes = c.String("sizes")
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

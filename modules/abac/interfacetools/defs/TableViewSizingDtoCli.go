//go:build !wasm

package interfacetoolsdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

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

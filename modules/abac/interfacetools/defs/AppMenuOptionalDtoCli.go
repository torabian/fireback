//go:build !wasm

package interfacetoolsdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetAppMenuOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "label",
			Type:        "complex",
			Description: "Label that will be visible to user, as a locale -> text map (e.g. {\"en\": \"Home\", \"fa\": \"خانه\"}) - see complexes.TString.",
		},
		{
			Name:        prefix + "href",
			Type:        "string?",
			Description: "Location that will be navigated in case of click or selection on ui",
		},
		{
			Name:        prefix + "icon",
			Type:        "string?",
			Description: "Icon string address which matches the resources on the front-end apps.",
		},
		{
			Name:        prefix + "active-matcher",
			Type:        "string?",
			Description: "Custom window location url matchers, for inner screens.",
		},
		{
			Name:        prefix + "capability-id",
			Type:        "string?",
			Description: "The unique-id of the capability which is required for the menu to be visible.",
		},
		{
			Name:        prefix + "parent-id",
			Type:        "string?",
			Description: "The unique-id of the parent menu item, for nested/tree menus.",
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
func CastAppMenuOptionalDtoFromCli(c emigo.CliCastable) AppMenuOptionalDto {
	data := AppMenuOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("label") {
		if u, ok := any(&data.Label).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("label")))
		}
	}
	if c.IsSet("href") {
		emigo.ParseNullable(c.String("href"), &data.Href)
	}
	if c.IsSet("icon") {
		emigo.ParseNullable(c.String("icon"), &data.Icon)
	}
	if c.IsSet("active-matcher") {
		emigo.ParseNullable(c.String("active-matcher"), &data.ActiveMatcher)
	}
	if c.IsSet("capability-id") {
		emigo.ParseNullable(c.String("capability-id"), &data.CapabilityId)
	}
	if c.IsSet("parent-id") {
		emigo.ParseNullable(c.String("parent-id"), &data.ParentId)
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

package interfacetoolsdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// The base class definition for appMenuOptionalDto
type AppMenuOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// Label that will be visible to user, as a locale -> text map (e.g. {"en": "Home", "fa": "خانه"}) - see complexes.TString.
	Label complexes.TString `json:"label" yaml:"label"`
	// Location that will be navigated in case of click or selection on ui
	Href emigo.Nullable[string] `json:"href" yaml:"href"`
	// Icon string address which matches the resources on the front-end apps.
	Icon emigo.Nullable[string] `json:"icon" yaml:"icon"`
	// Custom window location url matchers, for inner screens.
	ActiveMatcher emigo.Nullable[string] `json:"activeMatcher" yaml:"activeMatcher"`
	// The unique-id of the capability which is required for the menu to be visible.
	CapabilityId emigo.Nullable[string] `json:"capabilityId" yaml:"capabilityId"`
	// The unique-id of the parent menu item, for nested/tree menus.
	ParentId    emigo.Nullable[string]  `json:"parentId" yaml:"parentId"`
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId      emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *AppMenuOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

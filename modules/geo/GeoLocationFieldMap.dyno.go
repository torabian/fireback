package geo

import "github.com/torabian/fireback/modules/workspaces"

type GeoLocationFieldMap struct {
	Name workspaces.TranslatedString `yaml:"name"`

	Code workspaces.TranslatedString `yaml:"code"`

	Type workspaces.TranslatedString `yaml:"type"`

	Status workspaces.TranslatedString `yaml:"status"`

	Flag workspaces.TranslatedString `yaml:"flag"`

	OfficialName workspaces.TranslatedString `yaml:"officialName"`
}

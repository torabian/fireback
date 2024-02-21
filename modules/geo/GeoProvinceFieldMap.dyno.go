package geo

import "github.com/torabian/fireback/modules/workspaces"

type GeoProvinceFieldMap struct {
	Name workspaces.TranslatedString `yaml:"name"`

	Country workspaces.TranslatedString `yaml:"country"`
}

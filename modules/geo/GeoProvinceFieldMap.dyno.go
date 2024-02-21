package geo

import "pixelplux.com/fireback/modules/workspaces"

type GeoProvinceFieldMap struct {
	Name workspaces.TranslatedString `yaml:"name"`

	Country workspaces.TranslatedString `yaml:"country"`
}

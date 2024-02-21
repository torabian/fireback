package geo

import "github.com/torabian/fireback/modules/workspaces"

type LocationDataFieldMap struct {
	Lat workspaces.TranslatedString `yaml:"lat"`

	Lng workspaces.TranslatedString `yaml:"lng"`

	PhysicalAddress workspaces.TranslatedString `yaml:"physicalAddress"`
}

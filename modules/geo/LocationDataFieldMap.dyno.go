package geo

import "pixelplux.com/fireback/modules/workspaces"

type LocationDataFieldMap struct {
	Lat workspaces.TranslatedString `yaml:"lat"`

	Lng workspaces.TranslatedString `yaml:"lng"`

	PhysicalAddress workspaces.TranslatedString `yaml:"physicalAddress"`
}

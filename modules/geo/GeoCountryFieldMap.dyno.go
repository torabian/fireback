package geo

import "pixelplux.com/fireback/modules/workspaces"

type GeoCountryFieldMap struct {
	Status workspaces.TranslatedString `yaml:"status"`

	Flag workspaces.TranslatedString `yaml:"flag"`

	CommonName workspaces.TranslatedString `yaml:"commonName"`

	OfficialName workspaces.TranslatedString `yaml:"officialName"`
}

package worldtimezone

import "github.com/torabian/fireback/modules/workspaces"

type TimezoneGroupFieldMap struct {
	Value workspaces.TranslatedString `yaml:"value"`

	Abbr workspaces.TranslatedString `yaml:"abbr"`

	Offset workspaces.TranslatedString `yaml:"offset"`

	Isdst workspaces.TranslatedString `yaml:"isdst"`

	Text workspaces.TranslatedString `yaml:"text"`

	UtcItems workspaces.TranslatedString `yaml:"utcItems"`
}

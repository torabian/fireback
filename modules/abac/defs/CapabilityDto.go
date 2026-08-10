package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for capabilityDto
type CapabilityDto struct {
	UniqueId    emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
}

func (x *CapabilityDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCapabilityDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
	}
}
func CastCapabilityDtoFromCli(c emigo.CliCastable) CapabilityDto {
	data := CapabilityDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	return data
}

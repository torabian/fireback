package abac

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for capabilityOptionalDto
type CapabilityOptionalDto struct {
	UniqueId    emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Name        emigo.Nullable[string] `json:"name" yaml:"name"`
	Description emigo.Nullable[string] `json:"description" yaml:"description"`
}

func (x *CapabilityOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCapabilityOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "name",
			Type: "string?",
		},
		{
			Name: prefix + "description",
			Type: "string?",
		},
	}
}
func CastCapabilityOptionalDtoFromCli(c emigo.CliCastable) CapabilityOptionalDto {
	data := CapabilityOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("name") {
		emigo.ParseNullable(c.String("name"), &data.Name)
	}
	if c.IsSet("description") {
		emigo.ParseNullable(c.String("description"), &data.Description)
	}
	return data
}

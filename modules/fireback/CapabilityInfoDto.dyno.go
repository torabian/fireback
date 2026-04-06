package fireback

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetCapabilityInfoDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "children",
			Type: "collection",
		},
	}
}
func CastCapabilityInfoDtoFromCli(c emigo.CliCastable) CapabilityInfoDto {
	data := CapabilityInfoDto{}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for capabilityInfoDto
type CapabilityInfoDto struct {
	UniqueId string              `json:"uniqueId" yaml:"uniqueId"`
	Name     string              `json:"name" yaml:"name"`
	Children []CapabilityInfoDto `json:"children" yaml:"children"`
}

func (x *CapabilityInfoDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

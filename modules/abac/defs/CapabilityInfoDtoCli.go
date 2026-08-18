//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

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
	if c.IsSet("children") {
		data.Children = emigo.CapturePossibleCollection(CastCapabilityInfoDtoFromCli, "children", c)
	}
	return data
}

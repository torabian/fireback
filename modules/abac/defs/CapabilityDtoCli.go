//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

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

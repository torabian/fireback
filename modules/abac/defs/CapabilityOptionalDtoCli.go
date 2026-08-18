//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

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

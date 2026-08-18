//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetImportRequestDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "file",
			Type: "string",
		},
	}
}
func CastImportRequestDtoFromCli(c emigo.CliCastable) ImportRequestDto {
	data := ImportRequestDto{}
	if c.IsSet("file") {
		data.File = c.String("file")
	}
	return data
}

//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetExchangeKeyInformationDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "key",
			Type: "string",
		},
		{
			Name: prefix + "visibility",
			Type: "string",
		},
	}
}
func CastExchangeKeyInformationDtoFromCli(c emigo.CliCastable) ExchangeKeyInformationDto {
	data := ExchangeKeyInformationDto{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("visibility") {
		data.Visibility = c.String("visibility")
	}
	return data
}

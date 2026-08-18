//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetPhoneNumberAccountCreationDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "phone-number",
			Type: "string",
		},
	}
}
func CastPhoneNumberAccountCreationDtoFromCli(c emigo.CliCastable) PhoneNumberAccountCreationDto {
	data := PhoneNumberAccountCreationDto{}
	if c.IsSet("phone-number") {
		data.PhoneNumber = c.String("phone-number")
	}
	return data
}

//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetEmailAccountSigninDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "email",
			Type: "string",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
	}
}
func CastEmailAccountSigninDtoFromCli(c emigo.CliCastable) EmailAccountSigninDto {
	data := EmailAccountSigninDto{}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

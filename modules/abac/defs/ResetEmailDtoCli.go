//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetResetEmailDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "password",
			Type: "string",
		},
	}
}
func CastResetEmailDtoFromCli(c emigo.CliCastable) ResetEmailDto {
	data := ResetEmailDto{}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

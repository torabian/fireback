//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetOtpAuthenticateDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name: prefix + "otp",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
	}
}
func CastOtpAuthenticateDtoFromCli(c emigo.CliCastable) OtpAuthenticateDto {
	data := OtpAuthenticateDto{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("otp") {
		data.Otp = c.String("otp")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

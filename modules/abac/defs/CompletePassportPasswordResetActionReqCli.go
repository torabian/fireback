//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetCompletePassportPasswordResetActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "Passport value (email, phone number) the OTP was sent to.",
		},
		{
			Name:        prefix + "otp",
			Type:        "string",
			Description: "The OTP code emailed/texted to the passport value.",
		},
		{
			Name:        prefix + "password",
			Type:        "string",
			Description: "New password meeting the security requirements.",
		},
	}
}
func CastCompletePassportPasswordResetActionReqFromCli(c emigo.CliCastable) CompletePassportPasswordResetActionReq {
	data := CompletePassportPasswordResetActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("otp") {
		data.Otp = c.String("otp")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

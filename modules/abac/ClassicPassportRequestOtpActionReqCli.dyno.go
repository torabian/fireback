//go:build !wasm

package abac

func GetClassicPassportRequestOtpActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "Passport value (email, phone number) which would be receiving the otp code.",
		},
	}
}
func CastClassicPassportRequestOtpActionReqFromCli(c emigo.CliCastable) ClassicPassportRequestOtpActionReq {
	data := ClassicPassportRequestOtpActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	return data
}

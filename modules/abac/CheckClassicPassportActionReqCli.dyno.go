//go:build !wasm

package abac

import "github.com/torabian/emi/emigo"

func GetCheckClassicPassportActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name:        prefix + "security-token",
			Type:        "string",
			Description: "This can be the value of ReCaptcha2, ReCaptcha3, or generate security image or voice for verification. Will be used based on the configuration.",
		},
	}
}
func CastCheckClassicPassportActionReqFromCli(c emigo.CliCastable) CheckClassicPassportActionReq {
	data := CheckClassicPassportActionReq{}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("security-token") {
		data.SecurityToken = c.String("security-token")
	}
	return data
}

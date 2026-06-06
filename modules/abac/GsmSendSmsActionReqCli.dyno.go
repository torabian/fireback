//go:build !wasm

package abac

import "github.com/torabian/emi/emigo"

func GetGsmSendSmsActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "to-number",
			Type: "string",
		},
		{
			Name: prefix + "body",
			Type: "string",
		},
	}
}
func CastGsmSendSmsActionReqFromCli(c emigo.CliCastable) GsmSendSmsActionReq {
	data := GsmSendSmsActionReq{}
	if c.IsSet("to-number") {
		data.ToNumber = c.String("to-number")
	}
	if c.IsSet("body") {
		data.Body = c.String("body")
	}
	return data
}

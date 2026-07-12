//go:build !wasm

package abac

import "github.com/torabian/emi/emigo"

func GetOauthAuthenticateActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "token",
			Type:        "string",
			Description: "The token that Auth2 provider returned to the front-end, which will be used to validate the backend",
		},
		{
			Name:        prefix + "service",
			Type:        "string",
			Description: "The service name, such as 'google' which later backend will use to authorize the token and create the user.",
		},
	}
}
func CastOauthAuthenticateActionReqFromCli(c emigo.CliCastable) OauthAuthenticateActionReq {
	data := OauthAuthenticateActionReq{}
	if c.IsSet("token") {
		data.Token = c.String("token")
	}
	if c.IsSet("service") {
		data.Service = c.String("service")
	}
	return data
}

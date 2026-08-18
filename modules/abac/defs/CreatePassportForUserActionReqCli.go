//go:build !wasm

package abacdefs

import "github.com/torabian/emi/emigo"

func GetCreatePassportForUserActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "user-id",
			Type:        "string",
			Description: "UniqueId of the existing user to create the passport for.",
		},
		{
			Name:        prefix + "type",
			Type:        "string",
			Description: "The passport method: 'email' or 'phone'.",
		},
		{
			Name:        prefix + "value",
			Type:        "string",
			Description: "The email address or phone number for this passport. Must be globally unique across every passport.",
		},
		{
			Name:        prefix + "password",
			Type:        "string",
			Description: "Plaintext password to set on the new passport (minimum 6 characters) - hashed before it's stored, never returned.",
		},
	}
}
func CastCreatePassportForUserActionReqFromCli(c emigo.CliCastable) CreatePassportForUserActionReq {
	data := CreatePassportForUserActionReq{}
	if c.IsSet("user-id") {
		data.UserId = c.String("user-id")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	return data
}

//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetPassportEntityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name:        prefix + "third-party-verifier",
			Type:        "string",
			Description: "When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "user-id",
			Type: "string?",
		},
		{
			Name: prefix + "value",
			Type: "string",
		},
		{
			Name:        prefix + "totp-secret",
			Type:        "string",
			Description: "Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.",
		},
		{
			Name:        prefix + "totp-confirmed",
			Type:        "bool?",
			Description: "Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.",
		},
		{
			Name: prefix + "password",
			Type: "string",
		},
		{
			Name: prefix + "confirmed",
			Type: "bool?",
		},
		{
			Name: prefix + "access-token",
			Type: "string",
		},
		{
			Name: prefix + "workspace-id",
			Type: "string?",
		},
		{
			Name: prefix + "created-at",
			Type: "complex",
		},
		{
			Name: prefix + "updated-at",
			Type: "complex",
		},
	}
}
func CastPassportEntityFromCli(c emigo.CliCastable) PassportEntity {
	data := PassportEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("third-party-verifier") {
		data.ThirdPartyVerifier = c.String("third-party-verifier")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	if c.IsSet("totp-secret") {
		data.TotpSecret = c.String("totp-secret")
	}
	if c.IsSet("totp-confirmed") {
		emigo.ParseNullable(c.String("totp-confirmed"), &data.TotpConfirmed)
	}
	if c.IsSet("password") {
		data.Password = c.String("password")
	}
	if c.IsSet("confirmed") {
		emigo.ParseNullable(c.String("confirmed"), &data.Confirmed)
	}
	if c.IsSet("access-token") {
		data.AccessToken = c.String("access-token")
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	if c.IsSet("updated-at") {
		if u, ok := any(&data.UpdatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("updated-at")))
		}
	}
	return data
}

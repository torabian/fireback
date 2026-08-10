package abacdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for passportDto
type PassportDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.
	ThirdPartyVerifier string                 `json:"thirdPartyVerifier" yaml:"thirdPartyVerifier"`
	Type               string                 `json:"type" validate:"required" yaml:"type"`
	UserId             emigo.Nullable[string] `json:"userId" yaml:"userId"`
	Value              string                 `json:"value" validate:"required" yaml:"value"`
	// Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
	TotpSecret string `json:"totpSecret" yaml:"totpSecret"`
	// Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
	TotpConfirmed emigo.Nullable[bool]    `json:"totpConfirmed" yaml:"totpConfirmed"`
	Password      string                  `json:"-" yaml:"-"`
	Confirmed     emigo.Nullable[bool]    `json:"confirmed" yaml:"confirmed"`
	AccessToken   string                  `json:"accessToken" yaml:"accessToken"`
	WorkspaceId   emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt     abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt     abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PassportDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPassportDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
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
func CastPassportDtoFromCli(c emigo.CliCastable) PassportDto {
	data := PassportDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
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

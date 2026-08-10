package abacdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for publicAuthenticationOptionalDto
type PublicAuthenticationOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// The unique-id of the user which this record belongs to.
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// If the application requires totp dual factor upon account creation, we create a secret here and pass the link
	TotpSecret emigo.Nullable[string] `json:"totpSecret" yaml:"totpSecret"`
	// The url which will be converted into QR code on client side to scan
	TotpLink emigo.Nullable[string] `json:"totpLink" yaml:"totpLink"`
	// The unique-id of the passport this record belongs to.
	PassportId emigo.Nullable[string] `json:"passportId" yaml:"passportId"`
	// This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account
	SessionSecret       emigo.Nullable[string] `json:"sessionSecret" yaml:"sessionSecret"`
	PassportValue       emigo.Nullable[string] `json:"passportValue" yaml:"passportValue"`
	IsInCreationProcess emigo.Nullable[bool]   `json:"isInCreationProcess" yaml:"isInCreationProcess"`
	Status              emigo.Nullable[string] `json:"status" yaml:"status"`
	BlockedUntil        emigo.Nullable[int64]  `json:"blockedUntil" yaml:"blockedUntil"`
	Otp                 emigo.Nullable[string] `json:"otp" yaml:"otp"`
	RecoveryAbsoluteUrl emigo.Nullable[string] `json:"recoveryAbsoluteUrl" sql:"-" yaml:"recoveryAbsoluteUrl"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt   abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt   abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PublicAuthenticationOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetPublicAuthenticationOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user which this record belongs to.",
		},
		{
			Name:        prefix + "totp-secret",
			Type:        "string?",
			Description: "If the application requires totp dual factor upon account creation, we create a secret here and pass the link",
		},
		{
			Name:        prefix + "totp-link",
			Type:        "string?",
			Description: "The url which will be converted into QR code on client side to scan",
		},
		{
			Name:        prefix + "passport-id",
			Type:        "string?",
			Description: "The unique-id of the passport this record belongs to.",
		},
		{
			Name:        prefix + "session-secret",
			Type:        "string?",
			Description: "This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account",
		},
		{
			Name: prefix + "passport-value",
			Type: "string?",
		},
		{
			Name: prefix + "is-in-creation-process",
			Type: "bool?",
		},
		{
			Name: prefix + "status",
			Type: "string?",
		},
		{
			Name: prefix + "blocked-until",
			Type: "int64?",
		},
		{
			Name: prefix + "otp",
			Type: "string?",
		},
		{
			Name: prefix + "recovery-absolute-url",
			Type: "string?",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which content belongs to.",
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
func CastPublicAuthenticationOptionalDtoFromCli(c emigo.CliCastable) PublicAuthenticationOptionalDto {
	data := PublicAuthenticationOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
	}
	if c.IsSet("totp-secret") {
		emigo.ParseNullable(c.String("totp-secret"), &data.TotpSecret)
	}
	if c.IsSet("totp-link") {
		emigo.ParseNullable(c.String("totp-link"), &data.TotpLink)
	}
	if c.IsSet("passport-id") {
		emigo.ParseNullable(c.String("passport-id"), &data.PassportId)
	}
	if c.IsSet("session-secret") {
		emigo.ParseNullable(c.String("session-secret"), &data.SessionSecret)
	}
	if c.IsSet("passport-value") {
		emigo.ParseNullable(c.String("passport-value"), &data.PassportValue)
	}
	if c.IsSet("is-in-creation-process") {
		emigo.ParseNullable(c.String("is-in-creation-process"), &data.IsInCreationProcess)
	}
	if c.IsSet("status") {
		emigo.ParseNullable(c.String("status"), &data.Status)
	}
	if c.IsSet("blocked-until") {
		emigo.ParseNullable(c.String("blocked-until"), &data.BlockedUntil)
	}
	if c.IsSet("otp") {
		emigo.ParseNullable(c.String("otp"), &data.Otp)
	}
	if c.IsSet("recovery-absolute-url") {
		emigo.ParseNullable(c.String("recovery-absolute-url"), &data.RecoveryAbsoluteUrl)
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

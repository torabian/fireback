package abacdefs

import (
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

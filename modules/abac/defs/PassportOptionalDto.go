package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for passportOptionalDto
type PassportOptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.
	ThirdPartyVerifier emigo.Nullable[string] `json:"thirdPartyVerifier" yaml:"thirdPartyVerifier"`
	Type               emigo.Nullable[string] `json:"type" validate:"required" yaml:"type"`
	UserId             emigo.Nullable[string] `json:"userId" yaml:"userId"`
	Value              emigo.Nullable[string] `json:"value" validate:"required" yaml:"value"`
	// Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
	TotpSecret emigo.Nullable[string] `json:"totpSecret" yaml:"totpSecret"`
	// Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
	TotpConfirmed emigo.Nullable[bool]    `json:"totpConfirmed" yaml:"totpConfirmed"`
	Password      emigo.Nullable[string]  `json:"-" yaml:"-"`
	Confirmed     emigo.Nullable[bool]    `json:"confirmed" yaml:"confirmed"`
	AccessToken   emigo.Nullable[string]  `json:"accessToken" yaml:"accessToken"`
	WorkspaceId   emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	CreatedAt     abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt     abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *PassportOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

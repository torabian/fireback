package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for workspaceInviteDto
type WorkspaceInviteDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.
	PublicKey string `json:"publicKey" yaml:"publicKey"`
	// The content that user will receive to understand the reason of the letter.
	CoverLetter string `json:"coverLetter" yaml:"coverLetter"`
	// If the invited person has a different language, then you can define that so the interface for him will be automatically translated.
	TargetUserLocale string `json:"targetUserLocale" yaml:"targetUserLocale"`
	// The email address of the person which is invited.
	Email string `json:"email" yaml:"email"`
	// The phone number of the person which is invited.
	Phonenumber string `json:"phonenumber" yaml:"phonenumber"`
	// The unique-id of the workspace which user is being invited to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// First name of the person which is invited
	FirstName string `json:"firstName" validate:"required" yaml:"firstName"`
	// Last name of the person which is invited.
	LastName string `json:"lastName" validate:"required" yaml:"lastName"`
	// If forced, the email address cannot be changed by the user which has been invited.
	ForceEmailAddress emigo.Nullable[bool] `json:"forceEmailAddress" yaml:"forceEmailAddress"`
	// If forced, user cannot change the phone number and needs to complete signup.
	ForcePhoneNumber emigo.Nullable[bool] `json:"forcePhoneNumber" yaml:"forcePhoneNumber"`
	// The role which invitee get if they accept the request.
	RoleId emigo.Nullable[string] `json:"roleId" validate:"required" yaml:"roleId"`
	// The unique-id of the user which created/owns the record.
	UserId    emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceInviteDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

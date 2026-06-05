package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

// The base class definition for workspaceInvitationDto
type WorkspaceInvitationDto struct {
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
	// Workspace which user is being invite to.
	Workspace emigo.One[WorkspaceEntity] `json:"workspace" yaml:"workspace"`
	// First name of the person which is invited
	FirstName string `json:"firstName" validate:"required" yaml:"firstName"`
	// Last name of the person which is invited.
	LastName string `json:"lastName" validate:"required" yaml:"lastName"`
	// If forced, the email address cannot be changed by the user which has been invited.
	ForceEmailAddress emigo.Nullable[bool] `json:"forceEmailAddress" yaml:"forceEmailAddress"`
	// If forced, user cannot change the phone number and needs to complete signup.
	ForcePhoneNumber emigo.Nullable[bool] `json:"forcePhoneNumber" yaml:"forcePhoneNumber"`
	// The role which invitee get if they accept the request.
	RoleId string `json:"roleId" validate:"required" yaml:"roleId"`
}

func (x *WorkspaceInvitationDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

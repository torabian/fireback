package abac

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetWorkspaceInvitationDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "public-key",
			Type: "string",
		},
		{
			Name: prefix + "cover-letter",
			Type: "string",
		},
		{
			Name: prefix + "target-user-locale",
			Type: "string",
		},
		{
			Name: prefix + "email",
			Type: "string",
		},
		{
			Name: prefix + "phonenumber",
			Type: "string",
		},
		{
			Name: prefix + "workspace",
			Type: "one",
		},
		{
			Name: prefix + "first-name",
			Type: "string",
		},
		{
			Name: prefix + "last-name",
			Type: "string",
		},
		{
			Name: prefix + "force-email-address",
			Type: "bool?",
		},
		{
			Name: prefix + "force-phone-number",
			Type: "bool?",
		},
		{
			Name: prefix + "role-id",
			Type: "string",
		},
	}
}
func CastWorkspaceInvitationDtoFromCli(c emigo.CliCastable) WorkspaceInvitationDto {
	data := WorkspaceInvitationDto{}
	if c.IsSet("public-key") {
		data.PublicKey = c.String("public-key")
	}
	if c.IsSet("cover-letter") {
		data.CoverLetter = c.String("cover-letter")
	}
	if c.IsSet("target-user-locale") {
		data.TargetUserLocale = c.String("target-user-locale")
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("phonenumber") {
		data.Phonenumber = c.String("phonenumber")
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("force-email-address") {
		emigo.ParseNullable(c.String("force-email-address"), &data.ForceEmailAddress)
	}
	if c.IsSet("force-phone-number") {
		emigo.ParseNullable(c.String("force-phone-number"), &data.ForcePhoneNumber)
	}
	if c.IsSet("role-id") {
		data.RoleId = c.String("role-id")
	}
	return data
}

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
	Workspace WorkspaceEntity `json:"workspace" yaml:"workspace"`
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

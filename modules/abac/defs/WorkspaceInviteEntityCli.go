//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetWorkspaceInviteEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "public-key",
			Type:        "string",
			Description: "A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.",
		},
		{
			Name:        prefix + "cover-letter",
			Type:        "string",
			Description: "The content that user will receive to understand the reason of the letter.",
		},
		{
			Name:        prefix + "target-user-locale",
			Type:        "string",
			Description: "If the invited person has a different language, then you can define that so the interface for him will be automatically translated.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "The email address of the person which is invited.",
		},
		{
			Name:        prefix + "phonenumber",
			Type:        "string",
			Description: "The phone number of the person which is invited.",
		},
		{
			Name:        prefix + "workspace-id",
			Type:        "string?",
			Description: "The unique-id of the workspace which user is being invited to.",
		},
		{
			Name:        prefix + "first-name",
			Type:        "string",
			Description: "First name of the person which is invited",
		},
		{
			Name:        prefix + "last-name",
			Type:        "string",
			Description: "Last name of the person which is invited.",
		},
		{
			Name:        prefix + "force-email-address",
			Type:        "bool?",
			Description: "If forced, the email address cannot be changed by the user which has been invited.",
		},
		{
			Name:        prefix + "force-phone-number",
			Type:        "bool?",
			Description: "If forced, user cannot change the phone number and needs to complete signup.",
		},
		{
			Name:        prefix + "role-id",
			Type:        "string?",
			Description: "The role which invitee get if they accept the request.",
		},
		{
			Name:        prefix + "user-id",
			Type:        "string?",
			Description: "The unique-id of the user which created/owns the record.",
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
func CastWorkspaceInviteEntityFromCli(c emigo.CliCastable) WorkspaceInviteEntity {
	data := WorkspaceInviteEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
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
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
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
		emigo.ParseNullable(c.String("role-id"), &data.RoleId)
	}
	if c.IsSet("user-id") {
		emigo.ParseNullable(c.String("user-id"), &data.UserId)
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

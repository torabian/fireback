//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetWorkspaceConfigOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name:        prefix + "enable-recaptcha2",
			Type:        "bool?",
			Description: "Enables the recaptcha2 for authentication flow.",
		},
		{
			Name:        prefix + "enable-otp",
			Type:        "bool?",
			Description: "Enables the otp option. It's not forcing it, so user can choose if they want otp or password.",
		},
		{
			Name:        prefix + "require-otp-on-signup",
			Type:        "bool?",
			Description: "Forces the user to have otp verification before can create an account. They can define their password still.",
		},
		{
			Name:        prefix + "require-otp-on-signin",
			Type:        "bool?",
			Description: "Forces the user to use otp when signing in. Even if they have password set, they won't use it and only will be able to signin using that otp.",
		},
		{
			Name:        prefix + "recaptcha2-server-key",
			Type:        "string?",
			Description: "Secret which would be used to decrypt if the recaptcha is correct. Should not be available publicly.",
		},
		{
			Name:        prefix + "recaptcha2-client-key",
			Type:        "string?",
			Description: "Secret which would be used for recaptcha2 on the client side. Can be publicly visible, and upon authenticating users it would be sent to front-end.",
		},
		{
			Name:        prefix + "enable-totp",
			Type:        "bool?",
			Description: "Enables user to make 2FA using apps such as google authenticator or microsoft authenticator.",
		},
		{
			Name:        prefix + "force-totp",
			Type:        "bool?",
			Description: "Forces the user to setup a 2FA in order to access their account. Users which did not setup this won't be affected.",
		},
		{
			Name:        prefix + "force-password-on-phone",
			Type:        "bool?",
			Description: "Forces users who want to create account using phone number to also set a password on their account",
		},
		{
			Name:        prefix + "force-person-name-on-phone",
			Type:        "bool?",
			Description: "Forces the creation of account using phone number to ask for user first name and last name",
		},
		{
			Name:        prefix + "general-email-provider-id",
			Type:        "string?",
			Description: "The unique-id of the email provider service, which will be used to send the messages using it's service.",
		},
		{
			Name:        prefix + "general-gsm-provider-id",
			Type:        "string?",
			Description: "The unique-id of the general service which would be used to send text messages (sms).",
		},
		{
			Name:        prefix + "invite-to-workspace-content-id",
			Type:        "string?",
			Description: "The unique-id of the template used as default when a user is inviting a third-party into their own workspace.",
		},
		{
			Name:        prefix + "email-otp-content-id",
			Type:        "string?",
			Description: "The unique-id of the template used to fill the message for email one-time-password requests.",
		},
		{
			Name:        prefix + "sms-otp-content-id",
			Type:        "string?",
			Description: "The unique-id of the template used for OTP text messages, including the one time password code.",
		},
		{
			Name: prefix + "workspace-id",
			Type: "string?",
		},
		{
			Name: prefix + "user-id",
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
func CastWorkspaceConfigOptionalDtoFromCli(c emigo.CliCastable) WorkspaceConfigOptionalDto {
	data := WorkspaceConfigOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("enable-recaptcha2") {
		emigo.ParseNullable(c.String("enable-recaptcha2"), &data.EnableRecaptcha2)
	}
	if c.IsSet("enable-otp") {
		emigo.ParseNullable(c.String("enable-otp"), &data.EnableOtp)
	}
	if c.IsSet("require-otp-on-signup") {
		emigo.ParseNullable(c.String("require-otp-on-signup"), &data.RequireOtpOnSignup)
	}
	if c.IsSet("require-otp-on-signin") {
		emigo.ParseNullable(c.String("require-otp-on-signin"), &data.RequireOtpOnSignin)
	}
	if c.IsSet("recaptcha2-server-key") {
		emigo.ParseNullable(c.String("recaptcha2-server-key"), &data.Recaptcha2ServerKey)
	}
	if c.IsSet("recaptcha2-client-key") {
		emigo.ParseNullable(c.String("recaptcha2-client-key"), &data.Recaptcha2ClientKey)
	}
	if c.IsSet("enable-totp") {
		emigo.ParseNullable(c.String("enable-totp"), &data.EnableTotp)
	}
	if c.IsSet("force-totp") {
		emigo.ParseNullable(c.String("force-totp"), &data.ForceTotp)
	}
	if c.IsSet("force-password-on-phone") {
		emigo.ParseNullable(c.String("force-password-on-phone"), &data.ForcePasswordOnPhone)
	}
	if c.IsSet("force-person-name-on-phone") {
		emigo.ParseNullable(c.String("force-person-name-on-phone"), &data.ForcePersonNameOnPhone)
	}
	if c.IsSet("general-email-provider-id") {
		emigo.ParseNullable(c.String("general-email-provider-id"), &data.GeneralEmailProviderId)
	}
	if c.IsSet("general-gsm-provider-id") {
		emigo.ParseNullable(c.String("general-gsm-provider-id"), &data.GeneralGsmProviderId)
	}
	if c.IsSet("invite-to-workspace-content-id") {
		emigo.ParseNullable(c.String("invite-to-workspace-content-id"), &data.InviteToWorkspaceContentId)
	}
	if c.IsSet("email-otp-content-id") {
		emigo.ParseNullable(c.String("email-otp-content-id"), &data.EmailOtpContentId)
	}
	if c.IsSet("sms-otp-content-id") {
		emigo.ParseNullable(c.String("sms-otp-content-id"), &data.SmsOtpContentId)
	}
	if c.IsSet("workspace-id") {
		emigo.ParseNullable(c.String("workspace-id"), &data.WorkspaceId)
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

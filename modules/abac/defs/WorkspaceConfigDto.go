package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for workspaceConfigDto
type WorkspaceConfigDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	// Enables the recaptcha2 for authentication flow.
	EnableRecaptcha2 emigo.Nullable[bool] `json:"enableRecaptcha2" yaml:"enableRecaptcha2"`
	// Enables the otp option. It's not forcing it, so user can choose if they want otp or password.
	EnableOtp emigo.Nullable[bool] `json:"enableOtp" yaml:"enableOtp"`
	// Forces the user to have otp verification before can create an account. They can define their password still.
	RequireOtpOnSignup emigo.Nullable[bool] `json:"requireOtpOnSignup" yaml:"requireOtpOnSignup"`
	// Forces the user to use otp when signing in. Even if they have password set, they won't use it and only will be able to signin using that otp.
	RequireOtpOnSignin emigo.Nullable[bool] `json:"requireOtpOnSignin" yaml:"requireOtpOnSignin"`
	// Secret which would be used to decrypt if the recaptcha is correct. Should not be available publicly.
	Recaptcha2ServerKey string `json:"recaptcha2ServerKey" yaml:"recaptcha2ServerKey"`
	// Secret which would be used for recaptcha2 on the client side. Can be publicly visible, and upon authenticating users it would be sent to front-end.
	Recaptcha2ClientKey string `json:"recaptcha2ClientKey" yaml:"recaptcha2ClientKey"`
	// Enables user to make 2FA using apps such as google authenticator or microsoft authenticator.
	EnableTotp emigo.Nullable[bool] `json:"enableTotp" yaml:"enableTotp"`
	// Forces the user to setup a 2FA in order to access their account. Users which did not setup this won't be affected.
	ForceTotp emigo.Nullable[bool] `json:"forceTotp" yaml:"forceTotp"`
	// Forces users who want to create account using phone number to also set a password on their account
	ForcePasswordOnPhone emigo.Nullable[bool] `json:"forcePasswordOnPhone" yaml:"forcePasswordOnPhone"`
	// Forces the creation of account using phone number to ask for user first name and last name
	ForcePersonNameOnPhone emigo.Nullable[bool] `json:"forcePersonNameOnPhone" yaml:"forcePersonNameOnPhone"`
	// The unique-id of the email provider service, which will be used to send the messages using it's service.
	GeneralEmailProviderId emigo.Nullable[string] `json:"generalEmailProviderId" yaml:"generalEmailProviderId"`
	// The unique-id of the general service which would be used to send text messages (sms).
	GeneralGsmProviderId emigo.Nullable[string] `json:"generalGsmProviderId" yaml:"generalGsmProviderId"`
	// The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
	InviteToWorkspaceContentId emigo.Nullable[string] `json:"inviteToWorkspaceContentId" yaml:"inviteToWorkspaceContentId"`
	// The unique-id of the template used to fill the message for email one-time-password requests.
	EmailOtpContentId emigo.Nullable[string] `json:"emailOtpContentId" yaml:"emailOtpContentId"`
	// The unique-id of the template used for OTP text messages, including the one time password code.
	SmsOtpContentId emigo.Nullable[string]  `json:"smsOtpContentId" yaml:"smsOtpContentId"`
	WorkspaceId     emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId          emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt       abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt       abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *WorkspaceConfigDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

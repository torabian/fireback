package messagingdefs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for messagingConfigDto
type MessagingConfigDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
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
	CreatedAt       abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt       abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
	WorkspaceId     string                  `json:"workspaceId" yaml:"workspaceId"`
}

func (x *MessagingConfigDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetMessagingConfigDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
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
			Name: prefix + "created-at",
			Type: "complex",
		},
		{
			Name: prefix + "updated-at",
			Type: "complex",
		},
		{
			Name: prefix + "workspace-id",
			Type: "string",
		},
	}
}
func CastMessagingConfigDtoFromCli(c emigo.CliCastable) MessagingConfigDto {
	data := MessagingConfigDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
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
	if c.IsSet("workspace-id") {
		data.WorkspaceId = c.String("workspace-id")
	}
	return data
}

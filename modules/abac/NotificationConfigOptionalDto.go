package abac

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for notificationConfigOptionalDto
type NotificationConfigOptionalDto struct {
	UniqueId                               emigo.Nullable[string]  `json:"uniqueId" yaml:"uniqueId"`
	CascadeToSubWorkspaces                 emigo.Nullable[bool]    `json:"cascadeToSubWorkspaces" yaml:"cascadeToSubWorkspaces"`
	ForcedCascadeEmailProvider             emigo.Nullable[bool]    `json:"forcedCascadeEmailProvider" yaml:"forcedCascadeEmailProvider"`
	GeneralEmailProviderId                 emigo.Nullable[string]  `json:"generalEmailProviderId" yaml:"generalEmailProviderId"`
	GeneralGsmProviderId                   emigo.Nullable[string]  `json:"generalGsmProviderId" yaml:"generalGsmProviderId"`
	InviteToWorkspaceContent               emigo.Nullable[string]  `json:"inviteToWorkspaceContent" yaml:"inviteToWorkspaceContent"`
	InviteToWorkspaceContentExcerpt        emigo.Nullable[string]  `json:"inviteToWorkspaceContentExcerpt" yaml:"inviteToWorkspaceContentExcerpt"`
	InviteToWorkspaceContentDefault        emigo.Nullable[string]  `json:"inviteToWorkspaceContentDefault" yaml:"inviteToWorkspaceContentDefault"`
	InviteToWorkspaceContentDefaultExcerpt emigo.Nullable[string]  `json:"inviteToWorkspaceContentDefaultExcerpt" yaml:"inviteToWorkspaceContentDefaultExcerpt"`
	InviteToWorkspaceTitle                 emigo.Nullable[string]  `json:"inviteToWorkspaceTitle" yaml:"inviteToWorkspaceTitle"`
	InviteToWorkspaceTitleDefault          emigo.Nullable[string]  `json:"inviteToWorkspaceTitleDefault" yaml:"inviteToWorkspaceTitleDefault"`
	InviteToWorkspaceSenderId              emigo.Nullable[string]  `json:"inviteToWorkspaceSenderId" yaml:"inviteToWorkspaceSenderId"`
	AccountCenterEmailSenderId             emigo.Nullable[string]  `json:"accountCenterEmailSenderId" yaml:"accountCenterEmailSenderId"`
	ForgetPasswordContent                  emigo.Nullable[string]  `json:"forgetPasswordContent" yaml:"forgetPasswordContent"`
	ForgetPasswordContentExcerpt           emigo.Nullable[string]  `json:"forgetPasswordContentExcerpt" yaml:"forgetPasswordContentExcerpt"`
	ForgetPasswordContentDefault           emigo.Nullable[string]  `json:"forgetPasswordContentDefault" yaml:"forgetPasswordContentDefault"`
	ForgetPasswordContentDefaultExcerpt    emigo.Nullable[string]  `json:"forgetPasswordContentDefaultExcerpt" yaml:"forgetPasswordContentDefaultExcerpt"`
	ForgetPasswordTitle                    emigo.Nullable[string]  `json:"forgetPasswordTitle" yaml:"forgetPasswordTitle"`
	ForgetPasswordTitleDefault             emigo.Nullable[string]  `json:"forgetPasswordTitleDefault" yaml:"forgetPasswordTitleDefault"`
	ForgetPasswordSenderId                 emigo.Nullable[string]  `json:"forgetPasswordSenderId" yaml:"forgetPasswordSenderId"`
	AcceptLanguage                         emigo.Nullable[string]  `json:"acceptLanguage" yaml:"acceptLanguage"`
	ConfirmEmailSenderId                   emigo.Nullable[string]  `json:"confirmEmailSenderId" yaml:"confirmEmailSenderId"`
	ConfirmEmailContent                    emigo.Nullable[string]  `json:"confirmEmailContent" yaml:"confirmEmailContent"`
	ConfirmEmailContentExcerpt             emigo.Nullable[string]  `json:"confirmEmailContentExcerpt" yaml:"confirmEmailContentExcerpt"`
	ConfirmEmailContentDefault             emigo.Nullable[string]  `json:"confirmEmailContentDefault" yaml:"confirmEmailContentDefault"`
	ConfirmEmailContentDefaultExcerpt      emigo.Nullable[string]  `json:"confirmEmailContentDefaultExcerpt" yaml:"confirmEmailContentDefaultExcerpt"`
	ConfirmEmailTitle                      emigo.Nullable[string]  `json:"confirmEmailTitle" yaml:"confirmEmailTitle"`
	ConfirmEmailTitleDefault               emigo.Nullable[string]  `json:"confirmEmailTitleDefault" yaml:"confirmEmailTitleDefault"`
	WorkspaceId                            emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId                                 emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt                              abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt                              abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *NotificationConfigOptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetNotificationConfigOptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "cascade-to-sub-workspaces",
			Type: "bool?",
		},
		{
			Name: prefix + "forced-cascade-email-provider",
			Type: "bool?",
		},
		{
			Name: prefix + "general-email-provider-id",
			Type: "string?",
		},
		{
			Name: prefix + "general-gsm-provider-id",
			Type: "string?",
		},
		{
			Name: prefix + "invite-to-workspace-content",
			Type: "string?",
		},
		{
			Name: prefix + "invite-to-workspace-content-excerpt",
			Type: "string?",
		},
		{
			Name: prefix + "invite-to-workspace-content-default",
			Type: "string?",
		},
		{
			Name: prefix + "invite-to-workspace-content-default-excerpt",
			Type: "string?",
		},
		{
			Name: prefix + "invite-to-workspace-title",
			Type: "string?",
		},
		{
			Name: prefix + "invite-to-workspace-title-default",
			Type: "string?",
		},
		{
			Name: prefix + "invite-to-workspace-sender-id",
			Type: "string?",
		},
		{
			Name: prefix + "account-center-email-sender-id",
			Type: "string?",
		},
		{
			Name: prefix + "forget-password-content",
			Type: "string?",
		},
		{
			Name: prefix + "forget-password-content-excerpt",
			Type: "string?",
		},
		{
			Name: prefix + "forget-password-content-default",
			Type: "string?",
		},
		{
			Name: prefix + "forget-password-content-default-excerpt",
			Type: "string?",
		},
		{
			Name: prefix + "forget-password-title",
			Type: "string?",
		},
		{
			Name: prefix + "forget-password-title-default",
			Type: "string?",
		},
		{
			Name: prefix + "forget-password-sender-id",
			Type: "string?",
		},
		{
			Name: prefix + "accept-language",
			Type: "string?",
		},
		{
			Name: prefix + "confirm-email-sender-id",
			Type: "string?",
		},
		{
			Name: prefix + "confirm-email-content",
			Type: "string?",
		},
		{
			Name: prefix + "confirm-email-content-excerpt",
			Type: "string?",
		},
		{
			Name: prefix + "confirm-email-content-default",
			Type: "string?",
		},
		{
			Name: prefix + "confirm-email-content-default-excerpt",
			Type: "string?",
		},
		{
			Name: prefix + "confirm-email-title",
			Type: "string?",
		},
		{
			Name: prefix + "confirm-email-title-default",
			Type: "string?",
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
func CastNotificationConfigOptionalDtoFromCli(c emigo.CliCastable) NotificationConfigOptionalDto {
	data := NotificationConfigOptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("cascade-to-sub-workspaces") {
		emigo.ParseNullable(c.String("cascade-to-sub-workspaces"), &data.CascadeToSubWorkspaces)
	}
	if c.IsSet("forced-cascade-email-provider") {
		emigo.ParseNullable(c.String("forced-cascade-email-provider"), &data.ForcedCascadeEmailProvider)
	}
	if c.IsSet("general-email-provider-id") {
		emigo.ParseNullable(c.String("general-email-provider-id"), &data.GeneralEmailProviderId)
	}
	if c.IsSet("general-gsm-provider-id") {
		emigo.ParseNullable(c.String("general-gsm-provider-id"), &data.GeneralGsmProviderId)
	}
	if c.IsSet("invite-to-workspace-content") {
		emigo.ParseNullable(c.String("invite-to-workspace-content"), &data.InviteToWorkspaceContent)
	}
	if c.IsSet("invite-to-workspace-content-excerpt") {
		emigo.ParseNullable(c.String("invite-to-workspace-content-excerpt"), &data.InviteToWorkspaceContentExcerpt)
	}
	if c.IsSet("invite-to-workspace-content-default") {
		emigo.ParseNullable(c.String("invite-to-workspace-content-default"), &data.InviteToWorkspaceContentDefault)
	}
	if c.IsSet("invite-to-workspace-content-default-excerpt") {
		emigo.ParseNullable(c.String("invite-to-workspace-content-default-excerpt"), &data.InviteToWorkspaceContentDefaultExcerpt)
	}
	if c.IsSet("invite-to-workspace-title") {
		emigo.ParseNullable(c.String("invite-to-workspace-title"), &data.InviteToWorkspaceTitle)
	}
	if c.IsSet("invite-to-workspace-title-default") {
		emigo.ParseNullable(c.String("invite-to-workspace-title-default"), &data.InviteToWorkspaceTitleDefault)
	}
	if c.IsSet("invite-to-workspace-sender-id") {
		emigo.ParseNullable(c.String("invite-to-workspace-sender-id"), &data.InviteToWorkspaceSenderId)
	}
	if c.IsSet("account-center-email-sender-id") {
		emigo.ParseNullable(c.String("account-center-email-sender-id"), &data.AccountCenterEmailSenderId)
	}
	if c.IsSet("forget-password-content") {
		emigo.ParseNullable(c.String("forget-password-content"), &data.ForgetPasswordContent)
	}
	if c.IsSet("forget-password-content-excerpt") {
		emigo.ParseNullable(c.String("forget-password-content-excerpt"), &data.ForgetPasswordContentExcerpt)
	}
	if c.IsSet("forget-password-content-default") {
		emigo.ParseNullable(c.String("forget-password-content-default"), &data.ForgetPasswordContentDefault)
	}
	if c.IsSet("forget-password-content-default-excerpt") {
		emigo.ParseNullable(c.String("forget-password-content-default-excerpt"), &data.ForgetPasswordContentDefaultExcerpt)
	}
	if c.IsSet("forget-password-title") {
		emigo.ParseNullable(c.String("forget-password-title"), &data.ForgetPasswordTitle)
	}
	if c.IsSet("forget-password-title-default") {
		emigo.ParseNullable(c.String("forget-password-title-default"), &data.ForgetPasswordTitleDefault)
	}
	if c.IsSet("forget-password-sender-id") {
		emigo.ParseNullable(c.String("forget-password-sender-id"), &data.ForgetPasswordSenderId)
	}
	if c.IsSet("accept-language") {
		emigo.ParseNullable(c.String("accept-language"), &data.AcceptLanguage)
	}
	if c.IsSet("confirm-email-sender-id") {
		emigo.ParseNullable(c.String("confirm-email-sender-id"), &data.ConfirmEmailSenderId)
	}
	if c.IsSet("confirm-email-content") {
		emigo.ParseNullable(c.String("confirm-email-content"), &data.ConfirmEmailContent)
	}
	if c.IsSet("confirm-email-content-excerpt") {
		emigo.ParseNullable(c.String("confirm-email-content-excerpt"), &data.ConfirmEmailContentExcerpt)
	}
	if c.IsSet("confirm-email-content-default") {
		emigo.ParseNullable(c.String("confirm-email-content-default"), &data.ConfirmEmailContentDefault)
	}
	if c.IsSet("confirm-email-content-default-excerpt") {
		emigo.ParseNullable(c.String("confirm-email-content-default-excerpt"), &data.ConfirmEmailContentDefaultExcerpt)
	}
	if c.IsSet("confirm-email-title") {
		emigo.ParseNullable(c.String("confirm-email-title"), &data.ConfirmEmailTitle)
	}
	if c.IsSet("confirm-email-title-default") {
		emigo.ParseNullable(c.String("confirm-email-title-default"), &data.ConfirmEmailTitleDefault)
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

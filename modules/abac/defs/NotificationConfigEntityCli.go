//go:build !wasm

package abacdefs

import (
	"encoding"
	"github.com/torabian/emi/emigo"
)

func GetNotificationConfigEntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "cascade-to-sub-workspaces",
			Type: "bool",
		},
		{
			Name: prefix + "forced-cascade-email-provider",
			Type: "bool",
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
			Type: "string",
		},
		{
			Name: prefix + "invite-to-workspace-content-excerpt",
			Type: "string",
		},
		{
			Name: prefix + "invite-to-workspace-content-default",
			Type: "string",
		},
		{
			Name: prefix + "invite-to-workspace-content-default-excerpt",
			Type: "string",
		},
		{
			Name: prefix + "invite-to-workspace-title",
			Type: "string",
		},
		{
			Name: prefix + "invite-to-workspace-title-default",
			Type: "string",
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
			Type: "string",
		},
		{
			Name: prefix + "forget-password-content-excerpt",
			Type: "string",
		},
		{
			Name: prefix + "forget-password-content-default",
			Type: "string",
		},
		{
			Name: prefix + "forget-password-content-default-excerpt",
			Type: "string",
		},
		{
			Name: prefix + "forget-password-title",
			Type: "string",
		},
		{
			Name: prefix + "forget-password-title-default",
			Type: "string",
		},
		{
			Name: prefix + "forget-password-sender-id",
			Type: "string?",
		},
		{
			Name: prefix + "accept-language",
			Type: "string",
		},
		{
			Name: prefix + "confirm-email-sender-id",
			Type: "string?",
		},
		{
			Name: prefix + "confirm-email-content",
			Type: "string",
		},
		{
			Name: prefix + "confirm-email-content-excerpt",
			Type: "string",
		},
		{
			Name: prefix + "confirm-email-content-default",
			Type: "string",
		},
		{
			Name: prefix + "confirm-email-content-default-excerpt",
			Type: "string",
		},
		{
			Name: prefix + "confirm-email-title",
			Type: "string",
		},
		{
			Name: prefix + "confirm-email-title-default",
			Type: "string",
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
func CastNotificationConfigEntityFromCli(c emigo.CliCastable) NotificationConfigEntity {
	data := NotificationConfigEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("cascade-to-sub-workspaces") {
		data.CascadeToSubWorkspaces = bool(c.Bool("cascade-to-sub-workspaces"))
	}
	if c.IsSet("forced-cascade-email-provider") {
		data.ForcedCascadeEmailProvider = bool(c.Bool("forced-cascade-email-provider"))
	}
	if c.IsSet("general-email-provider-id") {
		emigo.ParseNullable(c.String("general-email-provider-id"), &data.GeneralEmailProviderId)
	}
	if c.IsSet("general-gsm-provider-id") {
		emigo.ParseNullable(c.String("general-gsm-provider-id"), &data.GeneralGsmProviderId)
	}
	if c.IsSet("invite-to-workspace-content") {
		data.InviteToWorkspaceContent = c.String("invite-to-workspace-content")
	}
	if c.IsSet("invite-to-workspace-content-excerpt") {
		data.InviteToWorkspaceContentExcerpt = c.String("invite-to-workspace-content-excerpt")
	}
	if c.IsSet("invite-to-workspace-content-default") {
		data.InviteToWorkspaceContentDefault = c.String("invite-to-workspace-content-default")
	}
	if c.IsSet("invite-to-workspace-content-default-excerpt") {
		data.InviteToWorkspaceContentDefaultExcerpt = c.String("invite-to-workspace-content-default-excerpt")
	}
	if c.IsSet("invite-to-workspace-title") {
		data.InviteToWorkspaceTitle = c.String("invite-to-workspace-title")
	}
	if c.IsSet("invite-to-workspace-title-default") {
		data.InviteToWorkspaceTitleDefault = c.String("invite-to-workspace-title-default")
	}
	if c.IsSet("invite-to-workspace-sender-id") {
		emigo.ParseNullable(c.String("invite-to-workspace-sender-id"), &data.InviteToWorkspaceSenderId)
	}
	if c.IsSet("account-center-email-sender-id") {
		emigo.ParseNullable(c.String("account-center-email-sender-id"), &data.AccountCenterEmailSenderId)
	}
	if c.IsSet("forget-password-content") {
		data.ForgetPasswordContent = c.String("forget-password-content")
	}
	if c.IsSet("forget-password-content-excerpt") {
		data.ForgetPasswordContentExcerpt = c.String("forget-password-content-excerpt")
	}
	if c.IsSet("forget-password-content-default") {
		data.ForgetPasswordContentDefault = c.String("forget-password-content-default")
	}
	if c.IsSet("forget-password-content-default-excerpt") {
		data.ForgetPasswordContentDefaultExcerpt = c.String("forget-password-content-default-excerpt")
	}
	if c.IsSet("forget-password-title") {
		data.ForgetPasswordTitle = c.String("forget-password-title")
	}
	if c.IsSet("forget-password-title-default") {
		data.ForgetPasswordTitleDefault = c.String("forget-password-title-default")
	}
	if c.IsSet("forget-password-sender-id") {
		emigo.ParseNullable(c.String("forget-password-sender-id"), &data.ForgetPasswordSenderId)
	}
	if c.IsSet("accept-language") {
		data.AcceptLanguage = c.String("accept-language")
	}
	if c.IsSet("confirm-email-sender-id") {
		emigo.ParseNullable(c.String("confirm-email-sender-id"), &data.ConfirmEmailSenderId)
	}
	if c.IsSet("confirm-email-content") {
		data.ConfirmEmailContent = c.String("confirm-email-content")
	}
	if c.IsSet("confirm-email-content-excerpt") {
		data.ConfirmEmailContentExcerpt = c.String("confirm-email-content-excerpt")
	}
	if c.IsSet("confirm-email-content-default") {
		data.ConfirmEmailContentDefault = c.String("confirm-email-content-default")
	}
	if c.IsSet("confirm-email-content-default-excerpt") {
		data.ConfirmEmailContentDefaultExcerpt = c.String("confirm-email-content-default-excerpt")
	}
	if c.IsSet("confirm-email-title") {
		data.ConfirmEmailTitle = c.String("confirm-email-title")
	}
	if c.IsSet("confirm-email-title-default") {
		data.ConfirmEmailTitleDefault = c.String("confirm-email-title-default")
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

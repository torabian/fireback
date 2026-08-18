package abacdefs

import (
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
	InviteToWorkspaceContentDefault        emigo.Nullable[string]  `json:"inviteToWorkspaceContentDefault" sql:"-" yaml:"inviteToWorkspaceContentDefault"`
	InviteToWorkspaceContentDefaultExcerpt emigo.Nullable[string]  `json:"inviteToWorkspaceContentDefaultExcerpt" sql:"-" yaml:"inviteToWorkspaceContentDefaultExcerpt"`
	InviteToWorkspaceTitle                 emigo.Nullable[string]  `json:"inviteToWorkspaceTitle" yaml:"inviteToWorkspaceTitle"`
	InviteToWorkspaceTitleDefault          emigo.Nullable[string]  `json:"inviteToWorkspaceTitleDefault" sql:"-" yaml:"inviteToWorkspaceTitleDefault"`
	InviteToWorkspaceSenderId              emigo.Nullable[string]  `json:"inviteToWorkspaceSenderId" yaml:"inviteToWorkspaceSenderId"`
	AccountCenterEmailSenderId             emigo.Nullable[string]  `json:"accountCenterEmailSenderId" yaml:"accountCenterEmailSenderId"`
	ForgetPasswordContent                  emigo.Nullable[string]  `json:"forgetPasswordContent" yaml:"forgetPasswordContent"`
	ForgetPasswordContentExcerpt           emigo.Nullable[string]  `json:"forgetPasswordContentExcerpt" yaml:"forgetPasswordContentExcerpt"`
	ForgetPasswordContentDefault           emigo.Nullable[string]  `json:"forgetPasswordContentDefault" sql:"-" yaml:"forgetPasswordContentDefault"`
	ForgetPasswordContentDefaultExcerpt    emigo.Nullable[string]  `json:"forgetPasswordContentDefaultExcerpt" sql:"-" yaml:"forgetPasswordContentDefaultExcerpt"`
	ForgetPasswordTitle                    emigo.Nullable[string]  `json:"forgetPasswordTitle" yaml:"forgetPasswordTitle"`
	ForgetPasswordTitleDefault             emigo.Nullable[string]  `json:"forgetPasswordTitleDefault" sql:"-" yaml:"forgetPasswordTitleDefault"`
	ForgetPasswordSenderId                 emigo.Nullable[string]  `json:"forgetPasswordSenderId" yaml:"forgetPasswordSenderId"`
	AcceptLanguage                         emigo.Nullable[string]  `json:"acceptLanguage" yaml:"acceptLanguage"`
	ConfirmEmailSenderId                   emigo.Nullable[string]  `json:"confirmEmailSenderId" yaml:"confirmEmailSenderId"`
	ConfirmEmailContent                    emigo.Nullable[string]  `json:"confirmEmailContent" yaml:"confirmEmailContent"`
	ConfirmEmailContentExcerpt             emigo.Nullable[string]  `json:"confirmEmailContentExcerpt" yaml:"confirmEmailContentExcerpt"`
	ConfirmEmailContentDefault             emigo.Nullable[string]  `json:"confirmEmailContentDefault" sql:"-" yaml:"confirmEmailContentDefault"`
	ConfirmEmailContentDefaultExcerpt      emigo.Nullable[string]  `json:"confirmEmailContentDefaultExcerpt" sql:"-" yaml:"confirmEmailContentDefaultExcerpt"`
	ConfirmEmailTitle                      emigo.Nullable[string]  `json:"confirmEmailTitle" yaml:"confirmEmailTitle"`
	ConfirmEmailTitleDefault               emigo.Nullable[string]  `json:"confirmEmailTitleDefault" sql:"-" yaml:"confirmEmailTitleDefault"`
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

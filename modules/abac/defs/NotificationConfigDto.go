package abacdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for notificationConfigDto
type NotificationConfigDto struct {
	UniqueId                               emigo.Nullable[string]  `json:"uniqueId" yaml:"uniqueId"`
	CascadeToSubWorkspaces                 bool                    `json:"cascadeToSubWorkspaces" yaml:"cascadeToSubWorkspaces"`
	ForcedCascadeEmailProvider             bool                    `json:"forcedCascadeEmailProvider" yaml:"forcedCascadeEmailProvider"`
	GeneralEmailProviderId                 emigo.Nullable[string]  `json:"generalEmailProviderId" yaml:"generalEmailProviderId"`
	GeneralGsmProviderId                   emigo.Nullable[string]  `json:"generalGsmProviderId" yaml:"generalGsmProviderId"`
	InviteToWorkspaceContent               string                  `json:"inviteToWorkspaceContent" yaml:"inviteToWorkspaceContent"`
	InviteToWorkspaceContentExcerpt        string                  `json:"inviteToWorkspaceContentExcerpt" yaml:"inviteToWorkspaceContentExcerpt"`
	InviteToWorkspaceContentDefault        string                  `json:"inviteToWorkspaceContentDefault" sql:"-" yaml:"inviteToWorkspaceContentDefault"`
	InviteToWorkspaceContentDefaultExcerpt string                  `json:"inviteToWorkspaceContentDefaultExcerpt" sql:"-" yaml:"inviteToWorkspaceContentDefaultExcerpt"`
	InviteToWorkspaceTitle                 string                  `json:"inviteToWorkspaceTitle" yaml:"inviteToWorkspaceTitle"`
	InviteToWorkspaceTitleDefault          string                  `json:"inviteToWorkspaceTitleDefault" sql:"-" yaml:"inviteToWorkspaceTitleDefault"`
	InviteToWorkspaceSenderId              emigo.Nullable[string]  `json:"inviteToWorkspaceSenderId" yaml:"inviteToWorkspaceSenderId"`
	AccountCenterEmailSenderId             emigo.Nullable[string]  `json:"accountCenterEmailSenderId" yaml:"accountCenterEmailSenderId"`
	ForgetPasswordContent                  string                  `json:"forgetPasswordContent" yaml:"forgetPasswordContent"`
	ForgetPasswordContentExcerpt           string                  `json:"forgetPasswordContentExcerpt" yaml:"forgetPasswordContentExcerpt"`
	ForgetPasswordContentDefault           string                  `json:"forgetPasswordContentDefault" sql:"-" yaml:"forgetPasswordContentDefault"`
	ForgetPasswordContentDefaultExcerpt    string                  `json:"forgetPasswordContentDefaultExcerpt" sql:"-" yaml:"forgetPasswordContentDefaultExcerpt"`
	ForgetPasswordTitle                    string                  `json:"forgetPasswordTitle" yaml:"forgetPasswordTitle"`
	ForgetPasswordTitleDefault             string                  `json:"forgetPasswordTitleDefault" sql:"-" yaml:"forgetPasswordTitleDefault"`
	ForgetPasswordSenderId                 emigo.Nullable[string]  `json:"forgetPasswordSenderId" yaml:"forgetPasswordSenderId"`
	AcceptLanguage                         string                  `json:"acceptLanguage" yaml:"acceptLanguage"`
	ConfirmEmailSenderId                   emigo.Nullable[string]  `json:"confirmEmailSenderId" yaml:"confirmEmailSenderId"`
	ConfirmEmailContent                    string                  `json:"confirmEmailContent" yaml:"confirmEmailContent"`
	ConfirmEmailContentExcerpt             string                  `json:"confirmEmailContentExcerpt" yaml:"confirmEmailContentExcerpt"`
	ConfirmEmailContentDefault             string                  `json:"confirmEmailContentDefault" sql:"-" yaml:"confirmEmailContentDefault"`
	ConfirmEmailContentDefaultExcerpt      string                  `json:"confirmEmailContentDefaultExcerpt" sql:"-" yaml:"confirmEmailContentDefaultExcerpt"`
	ConfirmEmailTitle                      string                  `json:"confirmEmailTitle" yaml:"confirmEmailTitle"`
	ConfirmEmailTitleDefault               string                  `json:"confirmEmailTitleDefault" sql:"-" yaml:"confirmEmailTitleDefault"`
	WorkspaceId                            emigo.Nullable[string]  `json:"workspaceId" yaml:"workspaceId"`
	UserId                                 emigo.Nullable[string]  `json:"userId" yaml:"userId"`
	CreatedAt                              abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	UpdatedAt                              abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *NotificationConfigDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

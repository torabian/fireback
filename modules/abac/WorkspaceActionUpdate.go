package abac

import (
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

func WorkspaceActionUpdate(query fireback.QueryDSL, fields *WorkspaceEntity) (*WorkspaceEntity, *fireback.IError) {

	var item WorkspaceEntity
	err := fireback.GetDbRef().
		Where(&WorkspaceEntity{UniqueId: fields.UniqueId}).
		First(&item).
		UpdateColumns(fields).Error
	if err != nil {
		return &item, fireback.GormErrorToIError(err)
	}

	return &item, nil
}

func SendInviteEmail(query fireback.QueryDSL, invite *WorkspaceInviteEntity) *fireback.IError {

	subject := "Workspace invitation"
	fromName := "Account service"
	fromEmail := "account@service.com"
	provider := &EmailProviderEntity{
		Type: EmailProviderType.Terminal,
	}
	content := `
	Hello FULL_NAME,
	You are invited to workspace WORKSPACE_NAME.

	You can continue on: INVITE_URL
	`

	config, err := NotificationConfigActionGetOneByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})

	if err == nil && config != nil {
		subject = config.InviteToWorkspaceTitle
		fromName = config.InviteToWorkspaceSender.FromName
		fromEmail = config.InviteToWorkspaceSender.FromEmailAddress
		provider = config.GeneralEmailProvider
		content = config.InviteToWorkspaceContent
	}

	content = strings.ReplaceAll(content, "FULL_NAME", invite.FirstName+" "+invite.LastName)
	content = strings.ReplaceAll(content, "INVITE_URL", "http://localhost:3000/en/join/"+invite.UniqueId)
	content = strings.ReplaceAll(content, "WORKSPACE_NAME", query.WorkspaceId)

	// Dangerous next line
	content = strings.ReplaceAll(content, "ROLE_NAME", invite.Role.Name)

	err3 := SendMail(EmailMessageContent{
		FromName:  fromName,
		FromEmail: fromEmail,
		ToName:    invite.FirstName,
		ToEmail:   invite.Email,
		Subject:   subject,
		Content:   content,
	}, provider)

	if err3 != nil {
		return fireback.GormErrorToIError(err3)
	}

	return nil
}

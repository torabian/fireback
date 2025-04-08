package abac

import (
	"strings"

	"github.com/torabian/fireback/modules/workspaces"
)

func WorkspaceActionUpdate(query workspaces.QueryDSL, fields *WorkspaceEntity) (*WorkspaceEntity, *workspaces.IError) {

	var item WorkspaceEntity
	err := workspaces.GetDbRef().
		Where(&WorkspaceEntity{UniqueId: fields.UniqueId}).
		First(&item).
		UpdateColumns(fields).Error
	if err != nil {
		return &item, workspaces.GormErrorToIError(err)
	}

	return &item, nil
}

func SendInviteEmail(query workspaces.QueryDSL, invite *WorkspaceInviteEntity) *workspaces.IError {

	config, err := NotificationConfigActionGetOneByWorkspace(workspaces.QueryDSL{WorkspaceId: ROOT_VAR})

	if err != nil {
		return err
	}

	if config == nil {
		return workspaces.Create401Error(&AbacMessages.EmailConfigurationIsNotAvailable, []string{})
	}

	if config.InviteToWorkspaceSender == nil {
		return workspaces.Create401Error(&AbacMessages.UserWhichHasThisTokenDoesNotExist, []string{})
	}

	content := config.InviteToWorkspaceContent
	content = strings.ReplaceAll(content, "FULL_NAME", invite.FirstName+" "+invite.LastName)
	content = strings.ReplaceAll(content, "INVITE_URL", "http://localhost:3000/en/join/"+invite.UniqueId)
	content = strings.ReplaceAll(content, "WORKSPACE_NAME", query.WorkspaceId)

	// Dangerous next line
	content = strings.ReplaceAll(content, "ROLE_NAME", invite.Role.Name)

	err3 := SendMail(EmailMessageContent{
		FromName:  config.InviteToWorkspaceSender.FromName,
		FromEmail: config.InviteToWorkspaceSender.FromEmailAddress,
		ToName:    invite.FirstName,
		ToEmail:   invite.Email,
		Subject:   config.InviteToWorkspaceTitle,
		Content:   content,
	}, config.GeneralEmailProvider)

	if err3 != nil {
		return workspaces.GormErrorToIError(err3)
	}

	return nil
}

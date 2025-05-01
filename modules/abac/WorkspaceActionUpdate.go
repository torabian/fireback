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

	config, err := NotificationConfigActionGetOneByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})

	if err != nil {
		return err
	}

	if config == nil {
		return fireback.Create401Error(&AbacMessages.EmailConfigurationIsNotAvailable, []string{})
	}

	if config.InviteToWorkspaceSender == nil {
		return fireback.Create401Error(&AbacMessages.UserWhichHasThisTokenDoesNotExist, []string{})
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
		return fireback.GormErrorToIError(err3)
	}

	return nil
}

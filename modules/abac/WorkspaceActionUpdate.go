package abac

import (
	"strings"

	"github.com/torabian/fireback/modules/abac/messaging"
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

	inviteToWorkspaceSenderId, hasSender := config.InviteToWorkspaceSenderId.Get()
	if !hasSender || *inviteToWorkspaceSenderId == "" {
		return fireback.Create401Error(&AbacMessages.InviteToWorkspaceMailSenderMissing, []string{})
	}
	sender, senderErr := messaging.EmailSenderActions.GetOne(fireback.QueryDSL{UniqueId: *inviteToWorkspaceSenderId})
	if senderErr != nil {
		return senderErr
	}

	content := config.InviteToWorkspaceContent
	content = strings.ReplaceAll(content, "FULL_NAME", invite.FirstName+" "+invite.LastName)
	content = strings.ReplaceAll(content, "INVITE_URL", "http://localhost:3000/en/join/"+invite.UniqueId)
	content = strings.ReplaceAll(content, "WORKSPACE_NAME", query.WorkspaceId)

	// Dangerous next line
	roleName := ""
	if role, roleErr := RoleActions.GetOne(fireback.QueryDSL{UniqueId: invite.RoleId.OrDefault("")}); roleErr == nil && role != nil {
		roleName = role.Name
	}
	content = strings.ReplaceAll(content, "ROLE_NAME", roleName)

	var provider *messaging.EmailProviderEntity
	if generalEmailProviderId, ok := config.GeneralEmailProviderId.Get(); ok && *generalEmailProviderId != "" {
		provider, _ = messaging.EmailProviderActions.GetOne(fireback.QueryDSL{UniqueId: *generalEmailProviderId})
	}

	err3 := messaging.SendMail(messaging.EmailMessageContent{
		FromName:  sender.FromName,
		FromEmail: sender.FromEmailAddress,
		ToName:    invite.FirstName,
		ToEmail:   invite.Email,
		Subject:   config.InviteToWorkspaceTitle,
		Content:   content,
	}, provider)

	if err3 != nil {
		return fireback.GormErrorToIError(err3)
	}

	return nil
}

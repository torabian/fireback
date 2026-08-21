package abac

import (
	"strings"

	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/abac/messaging"
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func WorkspaceActionUpdate(query fireback.QueryDSL, fields *abacdefs.WorkspaceEntity) (*abacdefs.WorkspaceEntity, *fireback.IError) {

	var item abacdefs.WorkspaceEntity
	// Bug fix: .Where(...).First(&item).UpdateColumns(fields) chained the read and the
	// write onto the very same *gorm.DB statement builder - First already consumes/sets
	// the statement's table from the Where model, so the subsequent UpdateColumns
	// re-inferred the table a second time on top of it, and every update failed with
	// `table name "workspace_entities" specified more than once (SQLSTATE 42712)`. Also,
	// the not-found case was only checked *after* the write had already been attempted.
	// Split into two independent statements instead: fetch first (bailing out on error/
	// not-found before writing anything), then a fresh Model(&item).UpdateColumns(fields).
	if err := fireback.GetDbRef().
		Where(&abacdefs.WorkspaceEntity{UniqueId: fields.UniqueId}).
		First(&item).Error; err != nil {
		return &item, fireback.GormErrorToIError(err)
	}

	if err := fireback.GetDbRef().Model(&item).UpdateColumns(fields).Error; err != nil {
		return &item, fireback.GormErrorToIError(err)
	}

	return &item, nil
}

func SendInviteEmail(query fireback.QueryDSL, invite *abacdefs.WorkspaceInviteEntity) *fireback.IError {

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

	var provider *messagingdefs.EmailProviderEntity
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

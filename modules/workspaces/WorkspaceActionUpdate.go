package workspaces

import (
	"strings"
)

func WorkspaceActionUpdate(query QueryDSL, fields *WorkspaceEntity) (*WorkspaceEntity, *IError) {

	var item WorkspaceEntity
	err := GetDbRef().
		Where(&WorkspaceEntity{UniqueId: fields.UniqueId}).
		First(&item).
		UpdateColumns(fields).Error
	if err != nil {
		return &item, GormErrorToIError(err)
	}

	return &item, nil
}

func SendInviteEmail(query QueryDSL, invite *WorkspaceInviteEntity) *IError {

	config, err := NotificationConfigActionGetOneByWorkspace(QueryDSL{WorkspaceId: ROOT_VAR})

	if err != nil {
		return err
	}

	if config == nil {
		return Create401Error(&WorkspacesMessages.EmailConfigurationIsNotAvailable, []string{})
	}

	if config.InviteToWorkspaceSender == nil {
		return Create401Error(&WorkspacesMessages.UserWhichHasThisTokenDoesNotExist, []string{})
	}

	content := *config.InviteToWorkspaceContent
	content = strings.ReplaceAll(content, "FULL_NAME", *invite.FirstName+" "+*invite.LastName)
	content = strings.ReplaceAll(content, "INVITE_URL", "http://localhost:3000/en/join/"+invite.UniqueId)
	content = strings.ReplaceAll(content, "WORKSPACE_NAME", query.WorkspaceId)

	// Dangerous next line
	content = strings.ReplaceAll(content, "ROLE_NAME", *invite.Role.Name)

	err3 := SendMail(EmailMessageContent{
		FromName:  *config.InviteToWorkspaceSender.FromName,
		FromEmail: *config.InviteToWorkspaceSender.FromEmailAddress,
		ToName:    *invite.FirstName,
		ToEmail:   *invite.Value,
		Subject:   *config.InviteToWorkspaceTitle,
		Content:   content,
	}, config.GeneralEmailProvider)

	if err3 != nil {
		return GormErrorToIError(err3)
	}

	return nil
}

// func WorkspaceInviteActionUpdate(query QueryDSL, fields *WorkspaceInviteEntity) (*WorkspaceInviteEntity, *IError) {

// 	fmt.Println(fields)
// 	var item WorkspaceInviteEntity
// 	err := GetDbRef().
// 		Where(&WorkspaceInviteEntity{UniqueId: fields.UniqueId}).
// 		First(&item).
// 		UpdateColumns(fields).Error
// 	if err != nil {
// 		return &item, GormErrorToIError(err)
// 	}

/// /  important @todo
// 	SendInviteEmail(query, fields)

// 	return &item, nil
// }

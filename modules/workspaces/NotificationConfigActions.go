package workspaces

import (
	"io/ioutil"

	seeders "github.com/torabian/fireback/modules/workspaces/seeders"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NotificationConfigActionCreate(
	dto *NotificationConfigEntity, query QueryDSL,
) (*NotificationConfigEntity, *IError) {
	return NotificationConfigActionCreateFn(dto, query)
}

func NotificationConfigActionUpdate(
	query QueryDSL,
	fields *NotificationConfigEntity,
) (*NotificationConfigEntity, *IError) {
	return NotificationConfigActionUpdateFn(query, fields)
}

func NotificationTestMailAction(
	dto *TestMailDto, query QueryDSL,
) (*OkayResponseDto, *IError) {

	q := query
	q.UniqueId = *dto.SenderId
	item, err := EmailSenderActionGetOne(q)

	if err != nil {
		return nil, err
	}

	conf, err2 := NotificationWorkspaecConfigActionGet(query)

	if err2 != nil {
		return nil, GormErrorToIError(err2)
	}

	err3 := SendMail(EmailMessageContent{
		FromEmail: *item.FromEmailAddress,
		FromName:  *item.FromName,
		ToName:    *dto.ToName,
		ToEmail:   *dto.ToEmail,
		Subject:   *dto.Subject,
		Content:   *dto.Content,
	}, conf.GeneralEmailProvider)

	if err3 != nil {
		return nil, GormErrorToIError(err)
	}

	return nil, GormErrorToIError(err)
}

func NotificationWorkspaecConfigActionGet(query QueryDSL) (*NotificationConfigEntity, *IError) {

	var item *NotificationConfigEntity

	q := GetDbRef()
	err := q.Preload("GeneralEmailProvider").
		Preload("InviteToWorkspaceSender").
		Preload("ForgetPasswordSender").
		Preload("ConfirmEmailSender").
		Where(RealEscape("workspace_id = ?", query.WorkspaceId)).First(&item).Error

	everything := "*"
	if err == gorm.ErrRecordNotFound {
		item = &NotificationConfigEntity{
			UniqueId:       UUID(),
			WorkspaceId:    &query.WorkspaceId,
			AcceptLanguage: &everything,
		}

		err = q.Create(&item).Error

		if err != nil {
			return item, GormErrorToIError(err)
		}
	}

	if item.ForgetPasswordContent == nil || *item.ForgetPasswordContent == "" {
		item.ForgetPasswordContent = &ForgetPasswordDefaultTemplate
	}
	item.ForgetPasswordContentDefault = &ForgetPasswordDefaultTemplate
	if item.ForgetPasswordTitle == nil || *item.ForgetPasswordTitle == "" {
		item.ForgetPasswordTitle = &ForgetPasswordDefaultTitle
	}
	item.ForgetPasswordTitleDefault = &ForgetPasswordDefaultTitle

	if item.InviteToWorkspaceContent == nil || *item.InviteToWorkspaceContent == "" {
		item.InviteToWorkspaceContent = &InviteWorkspaceTemplate
	}
	item.InviteToWorkspaceContentDefault = &InviteWorkspaceTemplate
	if item.InviteToWorkspaceTitle == nil || *item.InviteToWorkspaceTitle == "" {
		item.InviteToWorkspaceTitle = &InviteWorkspaceTitle
	}
	item.InviteToWorkspaceTitleDefault = &InviteWorkspaceTitle

	if item.ConfirmEmailContent == nil || *item.ConfirmEmailContent == "" {
		item.ConfirmEmailContent = &ConfirmMailTemplate
	}
	item.ConfirmEmailContentDefault = &ConfirmMailTemplate
	if item.ConfirmEmailTitle == nil || *item.ConfirmEmailTitle == "" {
		item.ConfirmEmailTitle = &ConfirmMailTitle
	}
	item.ConfirmEmailTitleDefault = &ConfirmMailTitle

	if err != nil {
		return item, GormErrorToIError(err)
	}

	return item, nil

}

var ForgetPasswordDefaultTemplate string
var ForgetPasswordDefaultTitle string

var ConfirmMailTemplate string
var ConfirmMailTitle string

var InviteWorkspaceTemplate string
var InviteWorkspaceTitle string

func init() {

	ForgetPasswordDefaultTitle = "Reset password"
	ConfirmMailTitle = "Confirm your email address"
	InviteWorkspaceTitle = "You are invited!"

	f, _ := seeders.ViewsFs.Open("forget-password.html")
	body, _ := ioutil.ReadAll(f)
	f.Close()
	ForgetPasswordDefaultTemplate = string(body)

	f, _ = seeders.ViewsFs.Open("confirm-mail.html")
	body, _ = ioutil.ReadAll(f)
	f.Close()
	ConfirmMailTemplate = string(body)

	f, _ = seeders.ViewsFs.Open("invitation-to-workspace.html")
	body, _ = ioutil.ReadAll(f)
	f.Close()
	InviteWorkspaceTemplate = string(body)
}

func NotificationWorkspaceConfigActionUpdate(
	query QueryDSL,
	fields *NotificationConfigEntity,
) (*NotificationConfigEntity, *IError) {

	NotificationConfigEntityPreSanitize(fields, query)
	var item NotificationConfigEntity
	q := GetDbRef().
		Where(&NotificationConfigEntity{WorkspaceId: &query.WorkspaceId}).
		First(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, GormErrorToIError(err)
	}

	err = GetDbRef().
		Preload(clause.Associations).
		Where(&NotificationConfigEntity{WorkspaceId: &query.WorkspaceId}).
		First(&item).Error

	return &item, nil
}

package abac

import (
	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NotificationConfigActionCreate(
	dto *NotificationConfigEntity, query fireback.QueryDSL,
) (*NotificationConfigEntity, *fireback.IError) {
	return NotificationConfigActionCreateFn(dto, query)
}

func NotificationConfigActionUpdate(
	query fireback.QueryDSL,
	fields *NotificationConfigEntity,
) (*NotificationConfigEntity, *fireback.IError) {
	return NotificationConfigActionUpdateFn(query, fields)
}

func NotificationTestMailAction(
	dto *TestMailDto, query fireback.QueryDSL,
) (*OkayResponseDto, *fireback.IError) {

	q := query
	q.UniqueId = dto.SenderId
	item, err := EmailSenderActions.GetOne(q)

	if err != nil {
		return nil, err
	}

	conf, err2 := NotificationWorkspaecConfigActionGet(query)

	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	err3 := SendMail(EmailMessageContent{
		FromEmail: item.FromEmailAddress,
		FromName:  item.FromName,
		ToName:    dto.ToName,
		ToEmail:   dto.ToEmail,
		Subject:   dto.Subject,
		Content:   dto.Content,
	}, conf.GeneralEmailProvider)

	if err3 != nil {
		return nil, fireback.GormErrorToIError(err)
	}

	return nil, fireback.GormErrorToIError(err)
}

func NotificationWorkspaecConfigActionGet(query fireback.QueryDSL) (*NotificationConfigEntity, *fireback.IError) {

	var item *NotificationConfigEntity

	q := fireback.GetDbRef()
	err := q.Preload("GeneralEmailProvider").
		Preload("InviteToWorkspaceSender").
		Preload("ForgetPasswordSender").
		Preload("ConfirmEmailSender").
		Where(fireback.RealEscape("workspace_id = ?", query.WorkspaceId)).First(&item).Error

	if err == gorm.ErrRecordNotFound {
		item = &NotificationConfigEntity{
			UniqueId:       fireback.UUID(),
			WorkspaceId:    fireback.NewString(query.WorkspaceId),
			AcceptLanguage: "*",
		}

		err = q.Create(&item).Error

		if err != nil {
			return item, fireback.GormErrorToIError(err)
		}
	}

	if item.ForgetPasswordContent == "" {
		item.ForgetPasswordContent = ForgetPasswordDefaultTemplate
	}
	item.ForgetPasswordContentDefault = ForgetPasswordDefaultTemplate
	if item.ForgetPasswordTitle == "" {
		item.ForgetPasswordTitle = ForgetPasswordDefaultTitle
	}
	item.ForgetPasswordTitleDefault = ForgetPasswordDefaultTitle

	if item.InviteToWorkspaceContent == "" {
		item.InviteToWorkspaceContent = InviteWorkspaceTemplate
	}
	item.InviteToWorkspaceContentDefault = InviteWorkspaceTemplate
	if item.InviteToWorkspaceTitle == "" {
		item.InviteToWorkspaceTitle = InviteWorkspaceTitle
	}
	item.InviteToWorkspaceTitleDefault = InviteWorkspaceTitle

	if item.ConfirmEmailContent == "" {
		item.ConfirmEmailContent = ConfirmMailTemplate
	}
	item.ConfirmEmailContentDefault = ConfirmMailTemplate
	if item.ConfirmEmailTitle == "" {
		item.ConfirmEmailTitle = ConfirmMailTitle
	}
	item.ConfirmEmailTitleDefault = ConfirmMailTitle

	if err != nil {
		return item, fireback.GormErrorToIError(err)
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

	// ForgetPasswordDefaultTitle = "Reset password"
	// ConfirmMailTitle = "Confirm your email address"
	// InviteWorkspaceTitle = "You are invited!"

	// f, _ := seeders.ViewsFs.Open("forget-password.html")
	// body, _ := ioutil.ReadAll(f)
	// f.Close()
	// ForgetPasswordDefaultTemplate = string(body)

	// f, _ = seeders.ViewsFs.Open("confirm-mail.html")
	// body, _ = ioutil.ReadAll(f)
	// f.Close()
	// ConfirmMailTemplate = string(body)

	// f, _ = seeders.ViewsFs.Open("invitation-to-workspace.html")
	// body, _ = ioutil.ReadAll(f)
	// f.Close()
	// InviteWorkspaceTemplate = string(body)
}

func NotificationWorkspaceConfigActionUpdate(
	query fireback.QueryDSL,
	fields *NotificationConfigEntity,
) (*NotificationConfigEntity, *fireback.IError) {

	NotificationConfigEntityPreSanitize(fields, query)
	var item NotificationConfigEntity
	q := fireback.GetDbRef().
		Where(&NotificationConfigEntity{WorkspaceId: fireback.NewString(query.WorkspaceId)}).
		First(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, fireback.GormErrorToIError(err)
	}

	err = fireback.GetDbRef().
		Preload(clause.Associations).
		Where(&NotificationConfigEntity{WorkspaceId: fireback.NewString(query.WorkspaceId)}).
		First(&item).Error

	return &item, nil
}

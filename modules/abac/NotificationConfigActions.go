package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/messaging"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"gorm.io/gorm"
)

// notificationConfig had permRewrite root.modules -> root.manage, and
// security: { writeOnRoot: true, resolveStrategy: workspace } in the old yaml: the old
// generated code set ResolveStrategy: "workspace" on every action, and AllowOnRoot: true
// only on Create/Update/Delete - preserved here exactly.
var notificationConfigPerms = NewCrudPermissionSet("root.manage", "notification-config", "notification config")
var PERM_ROOT_NOTIFICATION_CONFIG = notificationConfigPerms.Wildcard
var PERM_ROOT_NOTIFICATION_CONFIG_QUERY = notificationConfigPerms.Query
var PERM_ROOT_NOTIFICATION_CONFIG_CREATE = notificationConfigPerms.Create
var PERM_ROOT_NOTIFICATION_CONFIG_UPDATE = notificationConfigPerms.Update
var PERM_ROOT_NOTIFICATION_CONFIG_DELETE = notificationConfigPerms.Delete
var ALL_NOTIFICATION_CONFIG_PERMISSIONS = notificationConfigPerms.All

var NotificationConfigActions = NewEntityActionsBundle[NotificationConfigEntity]()

func notificationConfigSecurity(perm application.PermissionInfo, allowOnRoot bool) *fireback.SecurityModel {
	return &fireback.SecurityModel{
		ActionRequires:  []application.PermissionInfo{perm},
		ResolveStrategy: "workspace",
		AllowOnRoot:     allowOnRoot,
	}
}

func NotificationConfigBrowseAction(c NotificationConfigBrowseActionRequest) (*NotificationConfigBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, notificationConfigSecurity(PERM_ROOT_NOTIFICATION_CONFIG_QUERY, false))
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := NotificationConfigActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &NotificationConfigBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func NotificationConfigGetAction(c NotificationConfigGetActionRequest) (*NotificationConfigGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, notificationConfigSecurity(PERM_ROOT_NOTIFICATION_CONFIG_QUERY, false))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := NotificationConfigActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &NotificationConfigGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func NotificationConfigCreateAction(c NotificationConfigCreateActionRequest) (*NotificationConfigCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, notificationConfigSecurity(PERM_ROOT_NOTIFICATION_CONFIG_CREATE, true))
	if err != nil {
		return nil, err
	}
	entity := &NotificationConfigEntity{
		CascadeToSubWorkspaces:          c.Body.CascadeToSubWorkspaces,
		ForcedCascadeEmailProvider:      c.Body.ForcedCascadeEmailProvider,
		GeneralEmailProviderId:          c.Body.GeneralEmailProviderId,
		GeneralGsmProviderId:            c.Body.GeneralGsmProviderId,
		InviteToWorkspaceContent:        c.Body.InviteToWorkspaceContent,
		InviteToWorkspaceContentExcerpt: c.Body.InviteToWorkspaceContentExcerpt,
		InviteToWorkspaceTitle:          c.Body.InviteToWorkspaceTitle,
		InviteToWorkspaceSenderId:       c.Body.InviteToWorkspaceSenderId,
		AccountCenterEmailSenderId:      c.Body.AccountCenterEmailSenderId,
		ForgetPasswordContent:           c.Body.ForgetPasswordContent,
		ForgetPasswordContentExcerpt:    c.Body.ForgetPasswordContentExcerpt,
		ForgetPasswordTitle:             c.Body.ForgetPasswordTitle,
		ForgetPasswordSenderId:          c.Body.ForgetPasswordSenderId,
		AcceptLanguage:                  c.Body.AcceptLanguage,
		ConfirmEmailSenderId:            c.Body.ConfirmEmailSenderId,
		ConfirmEmailContent:             c.Body.ConfirmEmailContent,
		ConfirmEmailContentExcerpt:      c.Body.ConfirmEmailContentExcerpt,
		ConfirmEmailTitle:               c.Body.ConfirmEmailTitle,
		WorkspaceId:                     emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := NotificationConfigActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &NotificationConfigCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

// NotificationConfigUpdateAction upserts by workspace (distinctBy: workspace in the old
// yaml): Query.WorkspaceId (resolved from ResolveStrategy: workspace above) is the
// condition for finding-or-creating the row, not the uniqueId path param - preserved
// from the old NotificationConfigUpdateExec's FirstOrCreate-by-workspace behavior.
func NotificationConfigUpdateAction(c NotificationConfigUpdateActionRequest) (*NotificationConfigUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, notificationConfigSecurity(PERM_ROOT_NOTIFICATION_CONFIG_UPDATE, true))
	if err != nil {
		return nil, err
	}

	dbref := fireback.GetDbRef()
	var item NotificationConfigEntity
	if err2 := dbref.
		Where(&NotificationConfigEntity{WorkspaceId: emigo.NullableOf(query.WorkspaceId)}).
		FirstOrCreate(&item).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	changes := map[string]interface{}{}
	if v, ok := c.Body.CascadeToSubWorkspaces.Get(); ok {
		changes["CascadeToSubWorkspaces"] = *v
	}
	if v, ok := c.Body.ForcedCascadeEmailProvider.Get(); ok {
		changes["ForcedCascadeEmailProvider"] = *v
	}
	if v, ok := c.Body.GeneralEmailProviderId.Get(); ok {
		changes["GeneralEmailProviderId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.GeneralGsmProviderId.Get(); ok {
		changes["GeneralGsmProviderId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.InviteToWorkspaceContent.Get(); ok {
		changes["InviteToWorkspaceContent"] = *v
	}
	if v, ok := c.Body.InviteToWorkspaceContentExcerpt.Get(); ok {
		changes["InviteToWorkspaceContentExcerpt"] = *v
	}
	if v, ok := c.Body.InviteToWorkspaceTitle.Get(); ok {
		changes["InviteToWorkspaceTitle"] = *v
	}
	if v, ok := c.Body.InviteToWorkspaceSenderId.Get(); ok {
		changes["InviteToWorkspaceSenderId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.AccountCenterEmailSenderId.Get(); ok {
		changes["AccountCenterEmailSenderId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.ForgetPasswordContent.Get(); ok {
		changes["ForgetPasswordContent"] = *v
	}
	if v, ok := c.Body.ForgetPasswordContentExcerpt.Get(); ok {
		changes["ForgetPasswordContentExcerpt"] = *v
	}
	if v, ok := c.Body.ForgetPasswordTitle.Get(); ok {
		changes["ForgetPasswordTitle"] = *v
	}
	if v, ok := c.Body.ForgetPasswordSenderId.Get(); ok {
		changes["ForgetPasswordSenderId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.AcceptLanguage.Get(); ok {
		changes["AcceptLanguage"] = *v
	}
	if v, ok := c.Body.ConfirmEmailSenderId.Get(); ok {
		changes["ConfirmEmailSenderId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.ConfirmEmailContent.Get(); ok {
		changes["ConfirmEmailContent"] = *v
	}
	if v, ok := c.Body.ConfirmEmailContentExcerpt.Get(); ok {
		changes["ConfirmEmailContentExcerpt"] = *v
	}
	if v, ok := c.Body.ConfirmEmailTitle.Get(); ok {
		changes["ConfirmEmailTitle"] = *v
	}

	if len(changes) > 0 {
		if err2 := dbref.Model(&item).Updates(changes).Error; err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}
	}

	var updated NotificationConfigEntity
	if err2 := dbref.Where("id = ?", item.Id).First(&updated).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &NotificationConfigUpdateActionResponse{Payload: fireback.GResponseSingleItem(&updated)}, nil
}

func NotificationConfigAwareDeletePreviewAction(c NotificationConfigAwareDeletePreviewActionRequest) (*NotificationConfigAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, notificationConfigSecurity(PERM_ROOT_NOTIFICATION_CONFIG_DELETE, true)); err != nil {
		return nil, err
	}
	uniqueIds := NotificationConfigAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := NotificationConfigEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &NotificationConfigAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func NotificationConfigAwareDeleteAction(c NotificationConfigAwareDeleteActionRequest) (*NotificationConfigAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, notificationConfigSecurity(PERM_ROOT_NOTIFICATION_CONFIG_DELETE, true)); err != nil {
		return nil, err
	}
	if err2 := NotificationConfigEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &NotificationConfigAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// --- Hand business logic recovered from the pre-migration NotificationConfigActions.go ---
// (custom /notification/testmail and /notification/workspace/config routes, wired
// directly in NotificationModule.go now instead of through the old Module3Action +
// AppendNotificationConfigRouter indirection.)

func NotificationTestMailAction(
	dto *TestMailDto, query fireback.QueryDSL,
) (*OkayResponseDto, *fireback.IError) {

	q := query
	q.UniqueId = dto.SenderId
	item, err := messaging.EmailSenderActions.GetOne(q)

	if err != nil {
		return nil, err
	}

	conf, err2 := NotificationWorkspaecConfigActionGet(query)

	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	var provider *messaging.EmailProviderEntity
	if v, ok := conf.GeneralEmailProviderId.Get(); ok {
		provider, err2 = messaging.EmailProviderActions.GetOne(fireback.QueryDSL{UniqueId: *v})
		if err2 != nil {
			return nil, err2
		}
	}

	err3 := messaging.SendMail(messaging.EmailMessageContent{
		FromEmail: item.FromEmailAddress,
		FromName:  item.FromName,
		ToName:    dto.ToName,
		ToEmail:   dto.ToEmail,
		Subject:   dto.Subject,
		Content:   dto.Content,
	}, provider)

	if err3 != nil {
		return nil, fireback.CastToIError(err3)
	}

	return &OkayResponseDto{}, nil
}

// NotificationWorkspaecConfigActionGet gets-or-creates the notificationConfig row for
// query.WorkspaceId, and fills in the *Default/*DefaultExcerpt fields (sql:"-", never
// persisted) with the built-in fallback templates - preserved from the pre-migration
// hand file. The old code also preloaded the GeneralEmailProvider/InviteToWorkspaceSender/
// ForgetPasswordSender/ConfirmEmailSender relations here; those relations don't exist
// anymore (see the string uniqueId FK convention note in NotificationConfigEntity.go),
// so the preloads are dropped - callers needing the actual provider/sender now resolve
// them explicitly by id (see NotificationTestMailAction above).
func NotificationWorkspaecConfigActionGet(query fireback.QueryDSL) (*NotificationConfigEntity, *fireback.IError) {

	var item *NotificationConfigEntity

	q := fireback.GetDbRef()
	err := q.Where(fireback.RealEscape("workspace_id = ?", query.WorkspaceId)).First(&item).Error

	if err == gorm.ErrRecordNotFound {
		item = &NotificationConfigEntity{
			UniqueId:       fireback.UUID(),
			WorkspaceId:    emigo.NullableOf(query.WorkspaceId),
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

	if err != nil && err != gorm.ErrRecordNotFound {
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

// NotificationWorkspaceConfigActionUpdate is the handler behind the custom PATCH
// /notification/workspace/config route: a full-column overwrite (UpdateColumns, not a
// partial Nullable-field diff) of the workspace's notificationConfig row, keyed by
// workspace rather than uniqueId - preserved verbatim from the pre-migration hand file.
func NotificationWorkspaceConfigActionUpdate(
	query fireback.QueryDSL,
	fields *NotificationConfigEntity,
) (*NotificationConfigEntity, *fireback.IError) {

	var item NotificationConfigEntity
	q := fireback.GetDbRef().
		Where(&NotificationConfigEntity{WorkspaceId: emigo.NullableOf(query.WorkspaceId)}).
		First(&item)

	err := q.UpdateColumns(fields).Error
	if err != nil {
		return nil, fireback.GormErrorToIError(err)
	}

	err = fireback.GetDbRef().
		Where(&NotificationConfigEntity{WorkspaceId: emigo.NullableOf(query.WorkspaceId)}).
		First(&item).Error
	if err != nil {
		return nil, fireback.GormErrorToIError(err)
	}

	return &item, nil
}

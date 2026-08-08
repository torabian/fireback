package messaging

import (
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// emailSender's old security block was { writeOnRoot: true }: the old generated code
// only set AllowOnRoot: true on Create/Update/Delete, leaving Query/Get plain -
// preserved here exactly.
var emailSenderPerms = NewCrudPermissionSet("root.manage", "email-sender", "email sender")
var PERM_ROOT_EMAIL_SENDER = emailSenderPerms.Wildcard
var PERM_ROOT_EMAIL_SENDER_QUERY = emailSenderPerms.Query
var PERM_ROOT_EMAIL_SENDER_CREATE = emailSenderPerms.Create
var PERM_ROOT_EMAIL_SENDER_UPDATE = emailSenderPerms.Update
var PERM_ROOT_EMAIL_SENDER_DELETE = emailSenderPerms.Delete
var ALL_EMAIL_SENDER_PERMISSIONS = emailSenderPerms.All

// EmailSenderActions preserves the exact bundle shape (GetOne/Query/...) the old Module3
// entity compiler generated - abac's NotificationConfigActions.go/NotificationCli.go call
// into it directly (as messaging.EmailSenderActions).
var EmailSenderActions = NewEntityActionsBundle[EmailSenderEntity]()

func EmailSenderBrowseAction(c EmailSenderBrowseActionRequest) (*EmailSenderBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_SENDER_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := EmailSenderActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailSenderBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func EmailSenderGetAction(c EmailSenderGetActionRequest) (*EmailSenderGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_SENDER_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := EmailSenderActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailSenderGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func EmailSenderCreateAction(c EmailSenderCreateActionRequest) (*EmailSenderCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_SENDER_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	entity := &EmailSenderEntity{
		FromName:         c.Body.FromName,
		FromEmailAddress: c.Body.FromEmailAddress,
		ReplyTo:          c.Body.ReplyTo,
		NickName:         c.Body.NickName,
	}
	created, err2 := EmailSenderActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailSenderCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func EmailSenderUpdateAction(c EmailSenderUpdateActionRequest) (*EmailSenderUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_SENDER_UPDATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &EmailSenderEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.FromName.Get(); ok {
		fields.FromName = *v
	}
	if v, ok := c.Body.FromEmailAddress.Get(); ok {
		fields.FromEmailAddress = *v
	}
	if v, ok := c.Body.ReplyTo.Get(); ok {
		fields.ReplyTo = *v
	}
	if v, ok := c.Body.NickName.Get(); ok {
		fields.NickName = *v
	}
	updated, err2 := EmailSenderActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &EmailSenderUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func EmailSenderAwareDeletePreviewAction(c EmailSenderAwareDeletePreviewActionRequest) (*EmailSenderAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_SENDER_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := EmailSenderAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := EmailSenderEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &EmailSenderAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func EmailSenderAwareDeleteAction(c EmailSenderAwareDeleteActionRequest) (*EmailSenderAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_SENDER_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := EmailSenderEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &EmailSenderAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// GetEmailSenderAsStringList formats email senders for interactive CLI prompts (see
// abac's EmailProviderTestCmd, which calls this as messaging.GetEmailSenderAsStringList).
func GetEmailSenderAsStringList(items []*EmailSenderEntity) ([]string, error) {
	result := []string{}
	for _, entity := range items {
		result = append(result, entity.UniqueId+" >>> "+entity.FromEmailAddress+" - "+entity.FromName)
	}
	return result, nil
}

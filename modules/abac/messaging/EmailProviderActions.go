package messaging

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// emailProvider's old security block was { writeOnRoot: true, readOnRoot: true }: the old
// generated code set AllowOnRoot: true on every action (Query/Get/Create/Update/Delete) -
// preserved here exactly.
var emailProviderPerms = NewCrudPermissionSet("root.manage", "email-provider", "email provider")
var PERM_ROOT_EMAIL_PROVIDER = emailProviderPerms.Wildcard
var PERM_ROOT_EMAIL_PROVIDER_QUERY = emailProviderPerms.Query
var PERM_ROOT_EMAIL_PROVIDER_CREATE = emailProviderPerms.Create
var PERM_ROOT_EMAIL_PROVIDER_UPDATE = emailProviderPerms.Update
var PERM_ROOT_EMAIL_PROVIDER_DELETE = emailProviderPerms.Delete
var ALL_EMAIL_PROVIDER_PERMISSIONS = emailProviderPerms.All

var EmailProviderActions = NewEntityActionsBundle[EmailProviderEntity]()

func emailProviderSecurity(perm application.PermissionInfo) *fireback.SecurityModel {
	return &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{perm}, AllowOnRoot: true}
}

func EmailProviderBrowseAction(c EmailProviderBrowseActionRequest) (*EmailProviderBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, emailProviderSecurity(PERM_ROOT_EMAIL_PROVIDER_QUERY))
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := EmailProviderActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailProviderBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func EmailProviderGetAction(c EmailProviderGetActionRequest) (*EmailProviderGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, emailProviderSecurity(PERM_ROOT_EMAIL_PROVIDER_QUERY))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := EmailProviderActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailProviderGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func EmailProviderCreateAction(c EmailProviderCreateActionRequest) (*EmailProviderCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, emailProviderSecurity(PERM_ROOT_EMAIL_PROVIDER_CREATE))
	if err != nil {
		return nil, err
	}
	entity := &EmailProviderEntity{Type: c.Body.Type, Title: c.Body.Title, Config: c.Body.Config}
	created, err2 := EmailProviderActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailProviderCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func EmailProviderUpdateAction(c EmailProviderUpdateActionRequest) (*EmailProviderUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, emailProviderSecurity(PERM_ROOT_EMAIL_PROVIDER_UPDATE))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &EmailProviderEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Type.Get(); ok {
		fields.Type = *v
	}
	if v, ok := c.Body.Title.Get(); ok {
		fields.Title = *v
	}
	fields.Config = c.Body.Config
	updated, err2 := EmailProviderActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &EmailProviderUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func EmailProviderAwareDeletePreviewAction(c EmailProviderAwareDeletePreviewActionRequest) (*EmailProviderAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, emailProviderSecurity(PERM_ROOT_EMAIL_PROVIDER_DELETE)); err != nil {
		return nil, err
	}
	uniqueIds := EmailProviderAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := EmailProviderEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &EmailProviderAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func EmailProviderAwareDeleteAction(c EmailProviderAwareDeleteActionRequest) (*EmailProviderAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, emailProviderSecurity(PERM_ROOT_EMAIL_PROVIDER_DELETE)); err != nil {
		return nil, err
	}
	if err2 := EmailProviderEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &EmailProviderAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// --- Hand business logic recovered from the pre-migration EmailProviderEntity.go ---

type IEmailSender interface {
	Send(message EmailMessageContent) error
}

func SendMailViaSendGrid(message EmailMessageContent, apiKey string) (*rest.Response, error) {

	from := mail.NewEmail(message.FromName, message.FromEmail)
	to := mail.NewEmail(message.ToName, message.ToEmail)

	message2 := mail.NewSingleEmail(from, message.Subject, to, message.Content, message.Content)

	client := sendgrid.NewSendClient(apiKey)
	res, err := client.Send(message2)

	if err != nil {
		return nil, err
	}

	return res, nil
}

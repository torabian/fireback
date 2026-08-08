package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var emailConfirmationPerms = NewCrudPermissionSet("root.modules", "email-confirmation", "email confirmation")
var PERM_ROOT_EMAIL_CONFIRMATION = emailConfirmationPerms.Wildcard
var PERM_ROOT_EMAIL_CONFIRMATION_QUERY = emailConfirmationPerms.Query
var PERM_ROOT_EMAIL_CONFIRMATION_CREATE = emailConfirmationPerms.Create
var PERM_ROOT_EMAIL_CONFIRMATION_UPDATE = emailConfirmationPerms.Update
var PERM_ROOT_EMAIL_CONFIRMATION_DELETE = emailConfirmationPerms.Delete
var ALL_EMAIL_CONFIRMATION_PERMISSIONS = emailConfirmationPerms.All

var EmailConfirmationActions = NewEntityActionsBundle[EmailConfirmationEntity]()

func EmailConfirmationBrowseAction(c EmailConfirmationBrowseActionRequest) (*EmailConfirmationBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := EmailConfirmationActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailConfirmationBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func EmailConfirmationGetAction(c EmailConfirmationGetActionRequest) (*EmailConfirmationGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := EmailConfirmationActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailConfirmationGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func EmailConfirmationCreateAction(c EmailConfirmationCreateActionRequest) (*EmailConfirmationCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &EmailConfirmationEntity{
		UserId:    c.Body.UserId,
		Status:    c.Body.Status,
		Email:     c.Body.Email,
		Key:       c.Body.Key,
		ExpiresAt: c.Body.ExpiresAt,
	}
	created, err2 := EmailConfirmationActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &EmailConfirmationCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func EmailConfirmationUpdateAction(c EmailConfirmationUpdateActionRequest) (*EmailConfirmationUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &EmailConfirmationEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.UserId.Get(); ok {
		fields.UserId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.Status.Get(); ok {
		fields.Status = *v
	}
	if v, ok := c.Body.Email.Get(); ok {
		fields.Email = *v
	}
	if v, ok := c.Body.Key.Get(); ok {
		fields.Key = *v
	}
	if v, ok := c.Body.ExpiresAt.Get(); ok {
		fields.ExpiresAt = *v
	}
	updated, err2 := EmailConfirmationActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &EmailConfirmationUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func EmailConfirmationAwareDeletePreviewAction(c EmailConfirmationAwareDeletePreviewActionRequest) (*EmailConfirmationAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := EmailConfirmationAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := EmailConfirmationEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &EmailConfirmationAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func EmailConfirmationAwareDeleteAction(c EmailConfirmationAwareDeleteActionRequest) (*EmailConfirmationAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := EmailConfirmationEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &EmailConfirmationAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

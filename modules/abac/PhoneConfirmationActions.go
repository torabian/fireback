package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var phoneConfirmationPerms = NewCrudPermissionSet("root.manage", "phone-confirmation", "phone confirmation")
var PERM_ROOT_PHONE_CONFIRMATION = phoneConfirmationPerms.Wildcard
var PERM_ROOT_PHONE_CONFIRMATION_QUERY = phoneConfirmationPerms.Query
var PERM_ROOT_PHONE_CONFIRMATION_CREATE = phoneConfirmationPerms.Create
var PERM_ROOT_PHONE_CONFIRMATION_UPDATE = phoneConfirmationPerms.Update
var PERM_ROOT_PHONE_CONFIRMATION_DELETE = phoneConfirmationPerms.Delete
var ALL_PHONE_CONFIRMATION_PERMISSIONS = phoneConfirmationPerms.All

var PhoneConfirmationActions = NewEntityActionsBundle[PhoneConfirmationEntity]()

func PhoneConfirmationBrowseAction(c PhoneConfirmationBrowseActionRequest) (*PhoneConfirmationBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PhoneConfirmationActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PhoneConfirmationBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PhoneConfirmationGetAction(c PhoneConfirmationGetActionRequest) (*PhoneConfirmationGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PhoneConfirmationActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PhoneConfirmationGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PhoneConfirmationCreateAction(c PhoneConfirmationCreateActionRequest) (*PhoneConfirmationCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &PhoneConfirmationEntity{
		UserId:      c.Body.UserId,
		Status:      c.Body.Status,
		PhoneNumber: c.Body.PhoneNumber,
		Key:         c.Body.Key,
		ExpiresAt:   c.Body.ExpiresAt,
	}
	created, err2 := PhoneConfirmationActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &PhoneConfirmationCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PhoneConfirmationUpdateAction(c PhoneConfirmationUpdateActionRequest) (*PhoneConfirmationUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &PhoneConfirmationEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.UserId.Get(); ok {
		fields.UserId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.Status.Get(); ok {
		fields.Status = *v
	}
	if v, ok := c.Body.PhoneNumber.Get(); ok {
		fields.PhoneNumber = *v
	}
	if v, ok := c.Body.Key.Get(); ok {
		fields.Key = *v
	}
	if v, ok := c.Body.ExpiresAt.Get(); ok {
		fields.ExpiresAt = *v
	}
	updated, err2 := PhoneConfirmationActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &PhoneConfirmationUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PhoneConfirmationAwareDeletePreviewAction(c PhoneConfirmationAwareDeletePreviewActionRequest) (*PhoneConfirmationAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := PhoneConfirmationAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := PhoneConfirmationEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PhoneConfirmationAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PhoneConfirmationAwareDeleteAction(c PhoneConfirmationAwareDeleteActionRequest) (*PhoneConfirmationAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := PhoneConfirmationEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PhoneConfirmationAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

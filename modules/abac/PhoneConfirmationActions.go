package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
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

var PhoneConfirmationActions = NewEntityActionsBundle[abacdefs.PhoneConfirmationEntity]()

func PhoneConfirmationBrowseAction(c abacdefs.PhoneConfirmationBrowseActionRequest) (*abacdefs.PhoneConfirmationBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PhoneConfirmationActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PhoneConfirmationBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PhoneConfirmationGetAction(c abacdefs.PhoneConfirmationGetActionRequest) (*abacdefs.PhoneConfirmationGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PhoneConfirmationActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PhoneConfirmationGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PhoneConfirmationCreateAction(c abacdefs.PhoneConfirmationCreateActionRequest) (*abacdefs.PhoneConfirmationCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &abacdefs.PhoneConfirmationEntity{
		UserId:      c.Body.UserId,
		Status:      c.Body.Status,
		PhoneNumber: c.Body.PhoneNumber,
		Key:         c.Body.Key,
		ExpiresAt:   c.Body.ExpiresAt,
		// Without this, the row is invisible to PhoneConfirmationBrowseAction's
		// workspace-scoped query (see GetSqlContext) - same workspace-stamping fix as
		// EmailConfirmationCreateAction/PassportCreateAction/AppMenuCreateAction.
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := PhoneConfirmationActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PhoneConfirmationCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PhoneConfirmationUpdateAction(c abacdefs.PhoneConfirmationUpdateActionRequest) (*abacdefs.PhoneConfirmationUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &abacdefs.PhoneConfirmationEntity{UniqueId: c.Params.UniqueId}
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
	return &abacdefs.PhoneConfirmationUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PhoneConfirmationAwareDeletePreviewAction(c abacdefs.PhoneConfirmationAwareDeletePreviewActionRequest) (*abacdefs.PhoneConfirmationAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.PhoneConfirmationAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.PhoneConfirmationEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.PhoneConfirmationAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PhoneConfirmationAwareDeleteAction(c abacdefs.PhoneConfirmationAwareDeleteActionRequest) (*abacdefs.PhoneConfirmationAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PHONE_CONFIRMATION_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.PhoneConfirmationEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.PhoneConfirmationAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

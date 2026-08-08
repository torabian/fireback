package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var publicJoinKeyPerms = NewCrudPermissionSet("root.modules", "public-join-key", "public join key")
var PERM_ROOT_PUBLIC_JOIN_KEY = publicJoinKeyPerms.Wildcard
var PERM_ROOT_PUBLIC_JOIN_KEY_QUERY = publicJoinKeyPerms.Query
var PERM_ROOT_PUBLIC_JOIN_KEY_CREATE = publicJoinKeyPerms.Create
var PERM_ROOT_PUBLIC_JOIN_KEY_UPDATE = publicJoinKeyPerms.Update
var PERM_ROOT_PUBLIC_JOIN_KEY_DELETE = publicJoinKeyPerms.Delete
var ALL_PUBLIC_JOIN_KEY_PERMISSIONS = publicJoinKeyPerms.All

var PublicJoinKeyActions = NewEntityActionsBundle[PublicJoinKeyEntity]()

func PublicJoinKeyBrowseAction(c PublicJoinKeyBrowseActionRequest) (*PublicJoinKeyBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PublicJoinKeyActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PublicJoinKeyBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PublicJoinKeyGetAction(c PublicJoinKeyGetActionRequest) (*PublicJoinKeyGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PublicJoinKeyActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PublicJoinKeyGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PublicJoinKeyCreateAction(c PublicJoinKeyCreateActionRequest) (*PublicJoinKeyCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &PublicJoinKeyEntity{RoleId: c.Body.RoleId, WorkspaceId: c.Body.WorkspaceId}
	created, err2 := PublicJoinKeyActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &PublicJoinKeyCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PublicJoinKeyUpdateAction(c PublicJoinKeyUpdateActionRequest) (*PublicJoinKeyUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &PublicJoinKeyEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.RoleId.Get(); ok {
		fields.RoleId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.WorkspaceId.Get(); ok {
		fields.WorkspaceId = emigo.NullableOf(*v)
	}
	updated, err2 := PublicJoinKeyActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &PublicJoinKeyUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PublicJoinKeyAwareDeletePreviewAction(c PublicJoinKeyAwareDeletePreviewActionRequest) (*PublicJoinKeyAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := PublicJoinKeyAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := PublicJoinKeyEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PublicJoinKeyAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PublicJoinKeyAwareDeleteAction(c PublicJoinKeyAwareDeleteActionRequest) (*PublicJoinKeyAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PUBLIC_JOIN_KEY_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := PublicJoinKeyEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PublicJoinKeyAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

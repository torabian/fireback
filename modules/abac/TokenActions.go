package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

// token's old security block was { writeOnRoot: true } (permRewrite root.modules -> root.manage):
// the old generated code only set AllowOnRoot: true on Create/Update/Delete, leaving Query/Get
// plain - preserved here exactly.
var tokenPerms = NewCrudPermissionSet("root.manage", "token", "token")
var PERM_ROOT_TOKEN = tokenPerms.Wildcard
var PERM_ROOT_TOKEN_QUERY = tokenPerms.Query
var PERM_ROOT_TOKEN_CREATE = tokenPerms.Create
var PERM_ROOT_TOKEN_UPDATE = tokenPerms.Update
var PERM_ROOT_TOKEN_DELETE = tokenPerms.Delete
var ALL_TOKEN_PERMISSIONS = tokenPerms.All

var TokenActions = NewEntityActionsBundle[TokenEntity]()

func TokenBrowseAction(c TokenBrowseActionRequest) (*TokenBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_TOKEN_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := TokenActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &TokenBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func TokenGetAction(c TokenGetActionRequest) (*TokenGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_TOKEN_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := TokenActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &TokenGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func TokenCreateAction(c TokenCreateActionRequest) (*TokenCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_TOKEN_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	entity := &TokenEntity{UserId: c.Body.UserId, Token: c.Body.Token, ValidUntil: c.Body.ValidUntil}
	created, err2 := TokenActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &TokenCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func TokenUpdateAction(c TokenUpdateActionRequest) (*TokenUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_TOKEN_UPDATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &TokenEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.UserId.Get(); ok {
		fields.UserId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.Token.Get(); ok {
		fields.Token = *v
	}
	fields.ValidUntil = c.Body.ValidUntil
	updated, err2 := TokenActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &TokenUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func TokenAwareDeletePreviewAction(c TokenAwareDeletePreviewActionRequest) (*TokenAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_TOKEN_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := TokenAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := TokenEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &TokenAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func TokenAwareDeleteAction(c TokenAwareDeleteActionRequest) (*TokenAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_TOKEN_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := TokenEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &TokenAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

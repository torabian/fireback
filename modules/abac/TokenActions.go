package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
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

var TokenActions = NewEntityActionsBundle[abacdefs.TokenEntity]()

func TokenBrowseAction(c abacdefs.TokenBrowseActionRequest) (*abacdefs.TokenBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_TOKEN_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := TokenActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.TokenBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func TokenGetAction(c abacdefs.TokenGetActionRequest) (*abacdefs.TokenGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_TOKEN_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := TokenActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.TokenGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func TokenCreateAction(c abacdefs.TokenCreateActionRequest) (*abacdefs.TokenCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_TOKEN_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	// Without this, the row is invisible to TokenBrowseAction's workspace-scoped query
	// (see GetSqlContext) - same workspace-stamping fix as
	// EmailConfirmationCreateAction/PassportCreateAction/AppMenuCreateAction.
	entity := &abacdefs.TokenEntity{UserId: c.Body.UserId, Token: c.Body.Token, ValidUntil: c.Body.ValidUntil, WorkspaceId: emigo.NullableOf(query.WorkspaceId)}
	created, err2 := TokenActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.TokenCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func TokenUpdateAction(c abacdefs.TokenUpdateActionRequest) (*abacdefs.TokenUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_TOKEN_UPDATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &abacdefs.TokenEntity{UniqueId: c.Params.UniqueId}
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
	return &abacdefs.TokenUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func TokenAwareDeletePreviewAction(c abacdefs.TokenAwareDeletePreviewActionRequest) (*abacdefs.TokenAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_TOKEN_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.TokenAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.TokenEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.TokenAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func TokenAwareDeleteAction(c abacdefs.TokenAwareDeleteActionRequest) (*abacdefs.TokenAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_TOKEN_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.TokenEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.TokenAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

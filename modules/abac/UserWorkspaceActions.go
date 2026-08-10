package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var userWorkspacePerms = NewCrudPermissionSet("root.manage", "user-workspace", "user workspace")
var PERM_ROOT_USER_WORKSPACE = userWorkspacePerms.Wildcard
var PERM_ROOT_USER_WORKSPACE_QUERY = userWorkspacePerms.Query
var PERM_ROOT_USER_WORKSPACE_CREATE = userWorkspacePerms.Create
var PERM_ROOT_USER_WORKSPACE_UPDATE = userWorkspacePerms.Update
var PERM_ROOT_USER_WORKSPACE_DELETE = userWorkspacePerms.Delete
var ALL_USER_WORKSPACE_PERMISSIONS = userWorkspacePerms.All

// UserWorkspaceActions preserves the exact bundle shape (Create/Query/...) the old
// Module3 entity compiler generated - AcceptInviteAction.go, AuthFlow.go,
// ClassicSigninAction.go, UserCli.go and WorkspaceCoreFeatures.go all call into it
// directly.
var UserWorkspaceActions = NewEntityActionsBundle[abacdefs.UserWorkspaceEntity]()

func userWorkspaceSecurity(perm application.PermissionInfo) *fireback.SecurityModel {
	return &fireback.SecurityModel{
		ActionRequires:  []application.PermissionInfo{perm},
		ResolveStrategy: fireback.ResolveStrategyUser,
	}
}

func UserWorkspaceBrowseAction(c abacdefs.UserWorkspaceBrowseActionRequest) (*abacdefs.UserWorkspaceBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, userWorkspaceSecurity(PERM_ROOT_USER_WORKSPACE_QUERY))
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := UserWorkspaceActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.UserWorkspaceBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func UserWorkspaceGetAction(c abacdefs.UserWorkspaceGetActionRequest) (*abacdefs.UserWorkspaceGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, userWorkspaceSecurity(PERM_ROOT_USER_WORKSPACE_QUERY))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := UserWorkspaceActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.UserWorkspaceGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func UserWorkspaceCreateAction(c abacdefs.UserWorkspaceCreateActionRequest) (*abacdefs.UserWorkspaceCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, userWorkspaceSecurity(PERM_ROOT_USER_WORKSPACE_CREATE))
	if err != nil {
		return nil, err
	}
	entity := &abacdefs.UserWorkspaceEntity{UserId: c.Body.UserId, WorkspaceId: c.Body.WorkspaceId}
	created, err2 := UserWorkspaceActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.UserWorkspaceCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func UserWorkspaceUpdateAction(c abacdefs.UserWorkspaceUpdateActionRequest) (*abacdefs.UserWorkspaceUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, userWorkspaceSecurity(PERM_ROOT_USER_WORKSPACE_UPDATE))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &abacdefs.UserWorkspaceEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.UserId.Get(); ok {
		fields.UserId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.WorkspaceId.Get(); ok {
		fields.WorkspaceId = emigo.NullableOf(*v)
	}
	updated, err2 := UserWorkspaceActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.UserWorkspaceUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func UserWorkspaceAwareDeletePreviewAction(c abacdefs.UserWorkspaceAwareDeletePreviewActionRequest) (*abacdefs.UserWorkspaceAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, userWorkspaceSecurity(PERM_ROOT_USER_WORKSPACE_DELETE)); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.UserWorkspaceAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.UserWorkspaceEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.UserWorkspaceAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func UserWorkspaceAwareDeleteAction(c abacdefs.UserWorkspaceAwareDeleteActionRequest) (*abacdefs.UserWorkspaceAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, userWorkspaceSecurity(PERM_ROOT_USER_WORKSPACE_DELETE)); err != nil {
		return nil, err
	}
	if err2 := abacdefs.UserWorkspaceEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.UserWorkspaceAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

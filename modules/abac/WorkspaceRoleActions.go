package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var workspaceRolePerms = NewCrudPermissionSet("root.manage", "workspace-role", "workspace role")
var PERM_ROOT_WORKSPACE_ROLE = workspaceRolePerms.Wildcard
var PERM_ROOT_WORKSPACE_ROLE_QUERY = workspaceRolePerms.Query
var PERM_ROOT_WORKSPACE_ROLE_CREATE = workspaceRolePerms.Create
var PERM_ROOT_WORKSPACE_ROLE_UPDATE = workspaceRolePerms.Update
var PERM_ROOT_WORKSPACE_ROLE_DELETE = workspaceRolePerms.Delete
var ALL_WORKSPACE_ROLE_PERMISSIONS = workspaceRolePerms.All

// WorkspaceRoleActions preserves the exact bundle shape (Create/...) the old Module3
// entity compiler generated - AcceptInviteAction.go, AuthFlow.go, UserCli.go and
// WorkspaceCoreFeatures.go all call into it directly.
var WorkspaceRoleActions = NewEntityActionsBundle[abacdefs.WorkspaceRoleEntity]()

func WorkspaceRoleBrowseAction(c abacdefs.WorkspaceRoleBrowseActionRequest) (*abacdefs.WorkspaceRoleBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WorkspaceRoleActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceRoleBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func WorkspaceRoleGetAction(c abacdefs.WorkspaceRoleGetActionRequest) (*abacdefs.WorkspaceRoleGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := WorkspaceRoleActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceRoleGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WorkspaceRoleCreateAction(c abacdefs.WorkspaceRoleCreateActionRequest) (*abacdefs.WorkspaceRoleCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &abacdefs.WorkspaceRoleEntity{UserWorkspaceId: c.Body.UserWorkspaceId, RoleId: c.Body.RoleId}
	created, err2 := WorkspaceRoleActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceRoleCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func WorkspaceRoleUpdateAction(c abacdefs.WorkspaceRoleUpdateActionRequest) (*abacdefs.WorkspaceRoleUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &abacdefs.WorkspaceRoleEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.UserWorkspaceId.Get(); ok {
		fields.UserWorkspaceId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.RoleId.Get(); ok {
		fields.RoleId = emigo.NullableOf(*v)
	}
	updated, err2 := WorkspaceRoleActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceRoleUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func WorkspaceRoleAwareDeletePreviewAction(c abacdefs.WorkspaceRoleAwareDeletePreviewActionRequest) (*abacdefs.WorkspaceRoleAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.WorkspaceRoleAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.WorkspaceRoleEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.WorkspaceRoleAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WorkspaceRoleAwareDeleteAction(c abacdefs.WorkspaceRoleAwareDeleteActionRequest) (*abacdefs.WorkspaceRoleAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_ROLE_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.WorkspaceRoleEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.WorkspaceRoleAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

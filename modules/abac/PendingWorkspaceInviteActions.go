package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var pendingWorkspaceInvitePerms = NewCrudPermissionSet("root.modules", "pending-workspace-invite", "pending workspace invite")
var PERM_ROOT_PENDING_WORKSPACE_INVITE = pendingWorkspaceInvitePerms.Wildcard
var PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY = pendingWorkspaceInvitePerms.Query
var PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE = pendingWorkspaceInvitePerms.Create
var PERM_ROOT_PENDING_WORKSPACE_INVITE_UPDATE = pendingWorkspaceInvitePerms.Update
var PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE = pendingWorkspaceInvitePerms.Delete
var ALL_PENDING_WORKSPACE_INVITE_PERMISSIONS = pendingWorkspaceInvitePerms.All

var PendingWorkspaceInviteActions = NewEntityActionsBundle[PendingWorkspaceInviteEntity]()

func PendingWorkspaceInviteBrowseAction(c PendingWorkspaceInviteBrowseActionRequest) (*PendingWorkspaceInviteBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PendingWorkspaceInviteActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PendingWorkspaceInviteBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PendingWorkspaceInviteGetAction(c PendingWorkspaceInviteGetActionRequest) (*PendingWorkspaceInviteGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PendingWorkspaceInviteActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PendingWorkspaceInviteGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PendingWorkspaceInviteCreateAction(c PendingWorkspaceInviteCreateActionRequest) (*PendingWorkspaceInviteCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &PendingWorkspaceInviteEntity{
		Value:         c.Body.Value,
		Type:          c.Body.Type,
		CoverLetter:   c.Body.CoverLetter,
		WorkspaceName: c.Body.WorkspaceName,
	}
	if v, ok := c.Body.RoleId.Get(); ok {
		entity.RoleId = emigo.NullableOf(*v)
	}
	created, err2 := PendingWorkspaceInviteActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &PendingWorkspaceInviteCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PendingWorkspaceInviteUpdateAction(c PendingWorkspaceInviteUpdateActionRequest) (*PendingWorkspaceInviteUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &PendingWorkspaceInviteEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Value.Get(); ok {
		fields.Value = *v
	}
	if v, ok := c.Body.Type.Get(); ok {
		fields.Type = *v
	}
	if v, ok := c.Body.CoverLetter.Get(); ok {
		fields.CoverLetter = *v
	}
	if v, ok := c.Body.WorkspaceName.Get(); ok {
		fields.WorkspaceName = *v
	}
	if v, ok := c.Body.RoleId.Get(); ok {
		fields.RoleId = emigo.NullableOf(*v)
	}
	updated, err2 := PendingWorkspaceInviteActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &PendingWorkspaceInviteUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PendingWorkspaceInviteAwareDeletePreviewAction(c PendingWorkspaceInviteAwareDeletePreviewActionRequest) (*PendingWorkspaceInviteAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := PendingWorkspaceInviteAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := PendingWorkspaceInviteEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PendingWorkspaceInviteAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PendingWorkspaceInviteAwareDeleteAction(c PendingWorkspaceInviteAwareDeleteActionRequest) (*PendingWorkspaceInviteAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := PendingWorkspaceInviteEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PendingWorkspaceInviteAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

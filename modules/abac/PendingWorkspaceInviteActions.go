package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
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

var PendingWorkspaceInviteActions = NewEntityActionsBundle[abacdefs.PendingWorkspaceInviteEntity]()

func PendingWorkspaceInviteBrowseAction(c abacdefs.PendingWorkspaceInviteBrowseActionRequest) (*abacdefs.PendingWorkspaceInviteBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PendingWorkspaceInviteActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PendingWorkspaceInviteBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PendingWorkspaceInviteGetAction(c abacdefs.PendingWorkspaceInviteGetActionRequest) (*abacdefs.PendingWorkspaceInviteGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PendingWorkspaceInviteActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PendingWorkspaceInviteGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PendingWorkspaceInviteCreateAction(c abacdefs.PendingWorkspaceInviteCreateActionRequest) (*abacdefs.PendingWorkspaceInviteCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &abacdefs.PendingWorkspaceInviteEntity{
		Value:         c.Body.Value,
		Type:          c.Body.Type,
		CoverLetter:   c.Body.CoverLetter,
		WorkspaceName: c.Body.WorkspaceName,
		// Without this, the row is invisible to PendingWorkspaceInviteBrowseAction's
		// workspace-scoped query (see GetSqlContext) - same workspace-stamping fix as
		// EmailConfirmationCreateAction/PassportCreateAction/AppMenuCreateAction.
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
	}
	if v, ok := c.Body.RoleId.Get(); ok {
		entity.RoleId = emigo.NullableOf(*v)
	}
	created, err2 := PendingWorkspaceInviteActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PendingWorkspaceInviteCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PendingWorkspaceInviteUpdateAction(c abacdefs.PendingWorkspaceInviteUpdateActionRequest) (*abacdefs.PendingWorkspaceInviteUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &abacdefs.PendingWorkspaceInviteEntity{UniqueId: c.Params.UniqueId}
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
	return &abacdefs.PendingWorkspaceInviteUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PendingWorkspaceInviteAwareDeletePreviewAction(c abacdefs.PendingWorkspaceInviteAwareDeletePreviewActionRequest) (*abacdefs.PendingWorkspaceInviteAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.PendingWorkspaceInviteAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.PendingWorkspaceInviteEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.PendingWorkspaceInviteAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PendingWorkspaceInviteAwareDeleteAction(c abacdefs.PendingWorkspaceInviteAwareDeleteActionRequest) (*abacdefs.PendingWorkspaceInviteAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PENDING_WORKSPACE_INVITE_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.PendingWorkspaceInviteEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.PendingWorkspaceInviteAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

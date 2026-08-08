package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var workspaceInvitePerms = NewCrudPermissionSet("root.modules", "workspace-invite", "workspace invite")
var PERM_ROOT_WORKSPACE_INVITE = workspaceInvitePerms.Wildcard
var PERM_ROOT_WORKSPACE_INVITE_QUERY = workspaceInvitePerms.Query
var PERM_ROOT_WORKSPACE_INVITE_CREATE = workspaceInvitePerms.Create
var PERM_ROOT_WORKSPACE_INVITE_UPDATE = workspaceInvitePerms.Update
var PERM_ROOT_WORKSPACE_INVITE_DELETE = workspaceInvitePerms.Delete
var ALL_WORKSPACE_INVITE_PERMISSIONS = workspaceInvitePerms.All

// WorkspaceInviteActions preserves the exact bundle shape (GetOne/Create/Update/Query/
// RemoveEnqueue/...) the old Module3 entity compiler generated - AcceptInviteAction.go
// already calls into it directly.
var WorkspaceInviteActions = NewEntityActionsBundle[WorkspaceInviteEntity]()

func WorkspaceInviteBrowseAction(c WorkspaceInviteBrowseActionRequest) (*WorkspaceInviteBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WorkspaceInviteActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceInviteBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func WorkspaceInviteGetAction(c WorkspaceInviteGetActionRequest) (*WorkspaceInviteGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := WorkspaceInviteActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceInviteGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WorkspaceInviteCreateAction(c WorkspaceInviteCreateActionRequest) (*WorkspaceInviteCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &WorkspaceInviteEntity{
		PublicKey:         c.Body.PublicKey,
		CoverLetter:       c.Body.CoverLetter,
		TargetUserLocale:  c.Body.TargetUserLocale,
		Email:             c.Body.Email,
		Phonenumber:       c.Body.Phonenumber,
		WorkspaceId:       c.Body.WorkspaceId,
		FirstName:         c.Body.FirstName,
		LastName:          c.Body.LastName,
		ForceEmailAddress: c.Body.ForceEmailAddress,
		ForcePhoneNumber:  c.Body.ForcePhoneNumber,
		RoleId:            c.Body.RoleId,
	}
	created, err2 := WorkspaceInviteActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceInviteCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func WorkspaceInviteUpdateAction(c WorkspaceInviteUpdateActionRequest) (*WorkspaceInviteUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &WorkspaceInviteEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.PublicKey.Get(); ok {
		fields.PublicKey = *v
	}
	if v, ok := c.Body.CoverLetter.Get(); ok {
		fields.CoverLetter = *v
	}
	if v, ok := c.Body.TargetUserLocale.Get(); ok {
		fields.TargetUserLocale = *v
	}
	if v, ok := c.Body.Email.Get(); ok {
		fields.Email = *v
	}
	if v, ok := c.Body.Phonenumber.Get(); ok {
		fields.Phonenumber = *v
	}
	if v, ok := c.Body.WorkspaceId.Get(); ok {
		fields.WorkspaceId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.FirstName.Get(); ok {
		fields.FirstName = *v
	}
	if v, ok := c.Body.LastName.Get(); ok {
		fields.LastName = *v
	}
	if v, ok := c.Body.ForceEmailAddress.Get(); ok {
		fields.ForceEmailAddress = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.ForcePhoneNumber.Get(); ok {
		fields.ForcePhoneNumber = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.RoleId.Get(); ok {
		fields.RoleId = emigo.NullableOf(*v)
	}
	updated, err2 := WorkspaceInviteActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceInviteUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func WorkspaceInviteAwareDeletePreviewAction(c WorkspaceInviteAwareDeletePreviewActionRequest) (*WorkspaceInviteAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := WorkspaceInviteAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := WorkspaceInviteEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceInviteAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WorkspaceInviteAwareDeleteAction(c WorkspaceInviteAwareDeleteActionRequest) (*WorkspaceInviteAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_WORKSPACE_INVITE_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := WorkspaceInviteEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceInviteAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

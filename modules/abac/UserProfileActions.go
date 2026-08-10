package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var userProfilePerms = NewCrudPermissionSet("root.modules", "user-profile", "user profile")
var PERM_ROOT_USER_PROFILE = userProfilePerms.Wildcard
var PERM_ROOT_USER_PROFILE_QUERY = userProfilePerms.Query
var PERM_ROOT_USER_PROFILE_CREATE = userProfilePerms.Create
var PERM_ROOT_USER_PROFILE_UPDATE = userProfilePerms.Update
var PERM_ROOT_USER_PROFILE_DELETE = userProfilePerms.Delete
var ALL_USER_PROFILE_PERMISSIONS = userProfilePerms.All

var UserProfileActions = NewEntityActionsBundle[UserProfileEntity]()

func UserProfileBrowseAction(c UserProfileBrowseActionRequest) (*UserProfileBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_PROFILE_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := UserProfileActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &UserProfileBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func UserProfileGetAction(c UserProfileGetActionRequest) (*UserProfileGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_PROFILE_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := UserProfileActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &UserProfileGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func UserProfileCreateAction(c UserProfileCreateActionRequest) (*UserProfileCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_PROFILE_CREATE}})
	if err != nil {
		return nil, err
	}
	// Without this, the row is invisible to UserProfileBrowseAction's workspace-scoped
	// query (see GetSqlContext) - same workspace-stamping fix as
	// EmailConfirmationCreateAction/PassportCreateAction/AppMenuCreateAction.
	entity := &UserProfileEntity{FirstName: c.Body.FirstName, LastName: c.Body.LastName, WorkspaceId: emigo.NullableOf(query.WorkspaceId)}
	created, err2 := UserProfileActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &UserProfileCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func UserProfileUpdateAction(c UserProfileUpdateActionRequest) (*UserProfileUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_PROFILE_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &UserProfileEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.FirstName.Get(); ok {
		fields.FirstName = *v
	}
	if v, ok := c.Body.LastName.Get(); ok {
		fields.LastName = *v
	}
	updated, err2 := UserProfileActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &UserProfileUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func UserProfileAwareDeletePreviewAction(c UserProfileAwareDeletePreviewActionRequest) (*UserProfileAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_PROFILE_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := UserProfileAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := UserProfileEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &UserProfileAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func UserProfileAwareDeleteAction(c UserProfileAwareDeleteActionRequest) (*UserProfileAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_PROFILE_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := UserProfileEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &UserProfileAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var preferencePerms = NewCrudPermissionSet("root.modules", "preference", "preference")
var PERM_ROOT_PREFERENCE = preferencePerms.Wildcard
var PERM_ROOT_PREFERENCE_QUERY = preferencePerms.Query
var PERM_ROOT_PREFERENCE_CREATE = preferencePerms.Create
var PERM_ROOT_PREFERENCE_UPDATE = preferencePerms.Update
var PERM_ROOT_PREFERENCE_DELETE = preferencePerms.Delete
var ALL_PREFERENCE_PERMISSIONS = preferencePerms.All

var PreferenceActions = NewEntityActionsBundle[abacdefs.PreferenceEntity]()

func PreferenceBrowseAction(c abacdefs.PreferenceBrowseActionRequest) (*abacdefs.PreferenceBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PREFERENCE_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PreferenceActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PreferenceBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PreferenceGetAction(c abacdefs.PreferenceGetActionRequest) (*abacdefs.PreferenceGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PREFERENCE_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PreferenceActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PreferenceGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PreferenceCreateAction(c abacdefs.PreferenceCreateActionRequest) (*abacdefs.PreferenceCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PREFERENCE_CREATE}})
	if err != nil {
		return nil, err
	}
	// Without this, the row is invisible to PreferenceBrowseAction's workspace-scoped
	// query (see GetSqlContext) - same workspace-stamping fix as
	// EmailConfirmationCreateAction/PassportCreateAction/AppMenuCreateAction.
	entity := &abacdefs.PreferenceEntity{Timezone: c.Body.Timezone, WorkspaceId: emigo.NullableOf(query.WorkspaceId)}
	created, err2 := PreferenceActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PreferenceCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PreferenceUpdateAction(c abacdefs.PreferenceUpdateActionRequest) (*abacdefs.PreferenceUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PREFERENCE_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &abacdefs.PreferenceEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Timezone.Get(); ok {
		fields.Timezone = *v
	}
	updated, err2 := PreferenceActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.PreferenceUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PreferenceAwareDeletePreviewAction(c abacdefs.PreferenceAwareDeletePreviewActionRequest) (*abacdefs.PreferenceAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PREFERENCE_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.PreferenceAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.PreferenceEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.PreferenceAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PreferenceAwareDeleteAction(c abacdefs.PreferenceAwareDeleteActionRequest) (*abacdefs.PreferenceAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PREFERENCE_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.PreferenceEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.PreferenceAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

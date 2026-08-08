package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

var preferencePerms = NewCrudPermissionSet("root.modules", "preference", "preference")
var PERM_ROOT_PREFERENCE = preferencePerms.Wildcard
var PERM_ROOT_PREFERENCE_QUERY = preferencePerms.Query
var PERM_ROOT_PREFERENCE_CREATE = preferencePerms.Create
var PERM_ROOT_PREFERENCE_UPDATE = preferencePerms.Update
var PERM_ROOT_PREFERENCE_DELETE = preferencePerms.Delete
var ALL_PREFERENCE_PERMISSIONS = preferencePerms.All

var PreferenceActions = NewEntityActionsBundle[PreferenceEntity]()

func PreferenceBrowseAction(c PreferenceBrowseActionRequest) (*PreferenceBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PREFERENCE_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PreferenceActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PreferenceBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PreferenceGetAction(c PreferenceGetActionRequest) (*PreferenceGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PREFERENCE_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PreferenceActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PreferenceGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PreferenceCreateAction(c PreferenceCreateActionRequest) (*PreferenceCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PREFERENCE_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &PreferenceEntity{Timezone: c.Body.Timezone}
	created, err2 := PreferenceActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &PreferenceCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PreferenceUpdateAction(c PreferenceUpdateActionRequest) (*PreferenceUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PREFERENCE_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &PreferenceEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Timezone.Get(); ok {
		fields.Timezone = *v
	}
	updated, err2 := PreferenceActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &PreferenceUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PreferenceAwareDeletePreviewAction(c PreferenceAwareDeletePreviewActionRequest) (*PreferenceAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PREFERENCE_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := PreferenceAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := PreferenceEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PreferenceAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PreferenceAwareDeleteAction(c PreferenceAwareDeleteActionRequest) (*PreferenceAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PREFERENCE_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := PreferenceEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PreferenceAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

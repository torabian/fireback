package interfacetools

import (
	"reflect"

	"github.com/torabian/emi/emigo"
	interfacetoolsdefs "github.com/torabian/fireback/modules/abac/interfacetools/defs"
	seeders "github.com/torabian/fireback/modules/abac/interfacetools/seeders/TimezoneGroup"
	"github.com/torabian/fireback/modules/fireback"
)

func TimezoneGroupSyncSeeders() {
	fireback.SeederFromFSImport(
		fireback.QueryDSL{WorkspaceId: fireback.USER_SYSTEM},
		TimezoneGroupActions.Create,
		reflect.ValueOf(&interfacetoolsdefs.TimezoneGroupEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}

// timezoneGroup has queryScope: public in the old yaml - every generated action there
// used an empty, non-nil &fireback.SecurityModel{} (no ActionRequires, no
// ResolveStrategy override), rather than nil - preserved here exactly.
var timezoneGroupPerms = NewCrudPermissionSet("root.modules", "timezone-group", "timezone group")
var PERM_ROOT_TIMEZONE_GROUP = timezoneGroupPerms.Wildcard
var PERM_ROOT_TIMEZONE_GROUP_QUERY = timezoneGroupPerms.Query
var PERM_ROOT_TIMEZONE_GROUP_CREATE = timezoneGroupPerms.Create
var PERM_ROOT_TIMEZONE_GROUP_UPDATE = timezoneGroupPerms.Update
var PERM_ROOT_TIMEZONE_GROUP_DELETE = timezoneGroupPerms.Delete
var ALL_TIMEZONE_GROUP_PERMISSIONS = timezoneGroupPerms.All

var TimezoneGroupActions = NewEntityActionsBundle[interfacetoolsdefs.TimezoneGroupEntity]()

func TimezoneGroupBrowseAction(c interfacetoolsdefs.TimezoneGroupBrowseActionRequest) (*interfacetoolsdefs.TimezoneGroupBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := TimezoneGroupActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &interfacetoolsdefs.TimezoneGroupBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func TimezoneGroupGetAction(c interfacetoolsdefs.TimezoneGroupGetActionRequest) (*interfacetoolsdefs.TimezoneGroupGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := TimezoneGroupActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &interfacetoolsdefs.TimezoneGroupGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func TimezoneGroupCreateAction(c interfacetoolsdefs.TimezoneGroupCreateActionRequest) (*interfacetoolsdefs.TimezoneGroupCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{})
	if err != nil {
		return nil, err
	}
	// Without this, the row is invisible to TimezoneGroupBrowseAction's workspace-scoped
	// query (see GetSqlContext) - same workspace-stamping fix as
	// modules/abac/messaging/EmailProviderActions.go's EmailProviderCreateAction.
	entity := &interfacetoolsdefs.TimezoneGroupEntity{
		Title:       c.Body.Title,
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
		UserId:      emigo.NullableOf(query.UserId),
	}
	created, err2 := TimezoneGroupActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &interfacetoolsdefs.TimezoneGroupCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func TimezoneGroupUpdateAction(c interfacetoolsdefs.TimezoneGroupUpdateActionRequest) (*interfacetoolsdefs.TimezoneGroupUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &interfacetoolsdefs.TimezoneGroupEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Title.Get(); ok {
		fields.Title = *v
	}
	updated, err2 := TimezoneGroupActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &interfacetoolsdefs.TimezoneGroupUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func TimezoneGroupAwareDeletePreviewAction(c interfacetoolsdefs.TimezoneGroupAwareDeletePreviewActionRequest) (*interfacetoolsdefs.TimezoneGroupAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{}); err != nil {
		return nil, err
	}
	uniqueIds := interfacetoolsdefs.TimezoneGroupAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := interfacetoolsdefs.TimezoneGroupEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &interfacetoolsdefs.TimezoneGroupAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func TimezoneGroupAwareDeleteAction(c interfacetoolsdefs.TimezoneGroupAwareDeleteActionRequest) (*interfacetoolsdefs.TimezoneGroupAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{}); err != nil {
		return nil, err
	}
	if err2 := interfacetoolsdefs.TimezoneGroupEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &interfacetoolsdefs.TimezoneGroupAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

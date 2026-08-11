package interfacetools

import (
	"reflect"

	"github.com/torabian/emi/emigo"
	interfacetoolsdefs "github.com/torabian/fireback/modules/abac/interfacetools/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// Permission keys are preserved verbatim from the old Module3-generated
// TableViewSizingEntity.dyno.go (this entity had no permRewrite override), so existing
// role/capability records in any live database keep matching these completeKeys.
var PERM_ROOT_TABLE_VIEW_SIZING = application.PermissionInfo{
	CompleteKey: "root.modules.abac.table-view-sizing.*",
	Name:        "Entire table view sizing actions (*)",
}
var PERM_ROOT_TABLE_VIEW_SIZING_DELETE = application.PermissionInfo{
	CompleteKey: "root.modules.abac.table-view-sizing.delete",
	Name:        "Delete table view sizing",
}
var PERM_ROOT_TABLE_VIEW_SIZING_CREATE = application.PermissionInfo{
	CompleteKey: "root.modules.abac.table-view-sizing.create",
	Name:        "Create table view sizing",
}
var PERM_ROOT_TABLE_VIEW_SIZING_UPDATE = application.PermissionInfo{
	CompleteKey: "root.modules.abac.table-view-sizing.update",
	Name:        "Update table view sizing",
}
var PERM_ROOT_TABLE_VIEW_SIZING_QUERY = application.PermissionInfo{
	CompleteKey: "root.modules.abac.table-view-sizing.query",
	Name:        "Query table view sizing",
}
var ALL_TABLE_VIEW_SIZING_PERMISSIONS = []application.PermissionInfo{
	PERM_ROOT_TABLE_VIEW_SIZING_DELETE,
	PERM_ROOT_TABLE_VIEW_SIZING_CREATE,
	PERM_ROOT_TABLE_VIEW_SIZING_UPDATE,
	PERM_ROOT_TABLE_VIEW_SIZING_QUERY,
	PERM_ROOT_TABLE_VIEW_SIZING,
}

func TableViewSizingBrowseAction(c interfacetoolsdefs.TableViewSizingBrowseActionRequest) (*interfacetoolsdefs.TableViewSizingBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_QUERY},
	})
	if err != nil {
		return nil, err
	}
	q := *query

	refl := reflect.ValueOf(&interfacetoolsdefs.TableViewSizingEntity{})
	items, qrm, err2 := fireback.QueryEntitiesPointer[interfacetoolsdefs.TableViewSizingEntity](q, refl)
	if err2 != nil {
		return nil, err2
	}

	return &interfacetoolsdefs.TableViewSizingBrowseActionResponse{
		Payload: fireback.GResponseQuery(items, qrm, &q),
	}, nil
}

func TableViewSizingGetAction(c interfacetoolsdefs.TableViewSizingGetActionRequest) (*interfacetoolsdefs.TableViewSizingGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_QUERY},
	})
	if err != nil {
		return nil, err
	}
	q := *query
	q.UniqueId = c.Params.UniqueId

	refl := reflect.ValueOf(&interfacetoolsdefs.TableViewSizingEntity{})
	item, err2 := fireback.GetOneEntity[interfacetoolsdefs.TableViewSizingEntity](q, refl)
	if err2 != nil {
		return nil, err2
	}

	return &interfacetoolsdefs.TableViewSizingGetActionResponse{
		Payload: fireback.GResponseSingleItem(item),
	}, nil
}

func TableViewSizingCreateAction(c interfacetoolsdefs.TableViewSizingCreateActionRequest) (*interfacetoolsdefs.TableViewSizingCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_CREATE},
	})
	if err != nil {
		return nil, err
	}
	q := *query

	// tableName is validate:"required" (see InterfaceTools.emi.yml) but nothing was
	// actually enforcing it - same fix as
	// modules/abac/messaging/EmailProviderActions.go's EmailProviderCreateAction.
	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}

	entity := interfacetoolsdefs.TableViewSizingEntity{
		TableName:   c.Body.TableName,
		Sizes:       c.Body.Sizes,
		WorkspaceId: emigo.NullableOf(q.WorkspaceId),
		UserId:      emigo.NullableOf(q.UserId),
	}

	created, err2 := fireback.CreateEntity(entity)
	if err2 != nil {
		return nil, err2
	}

	return &interfacetoolsdefs.TableViewSizingCreateActionResponse{
		Payload: fireback.GResponseSingleItem(created),
	}, nil
}

func TableViewSizingUpdateAction(c interfacetoolsdefs.TableViewSizingUpdateActionRequest) (*interfacetoolsdefs.TableViewSizingUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_UPDATE},
	})
	if err != nil {
		return nil, err
	}
	q := *query
	q.UniqueId = c.Params.UniqueId

	fields := &interfacetoolsdefs.TableViewSizingEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.TableName.Get(); ok {
		fields.TableName = *v
	}
	if v, ok := c.Body.Sizes.Get(); ok {
		fields.Sizes = *v
	}

	updated, err2 := fireback.UpdateEntity(q, fields)
	if err2 != nil && err2.HttpCode == 404 {
		// The front-end (see CommonListManager.tsx's table-column-width persistence)
		// addresses this entity by a caller-chosen uniqueId (a per-table, per-user key),
		// not a server-generated one - so the first "update" for a given key is actually
		// a create. Preserved from the old distinctBy-less upsert-by-uniqueId behavior.
		fields.WorkspaceId = emigo.NullableOf(q.WorkspaceId)
		fields.UserId = emigo.NullableOf(q.UserId)
		created, err3 := fireback.CreateEntity(*fields)
		if err3 != nil {
			return nil, err3
		}
		return &interfacetoolsdefs.TableViewSizingUpdateActionResponse{
			Payload: fireback.GResponseSingleItem(created),
		}, nil
	}
	if err2 != nil {
		return nil, err2
	}

	return &interfacetoolsdefs.TableViewSizingUpdateActionResponse{
		Payload: fireback.GResponseSingleItem(updated),
	}, nil
}

func TableViewSizingAwareDeletePreviewAction(c interfacetoolsdefs.TableViewSizingAwareDeletePreviewActionRequest) (*interfacetoolsdefs.TableViewSizingAwareDeletePreviewActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_DELETE},
	})
	if err != nil {
		return nil, err
	}

	uniqueIds := interfacetoolsdefs.TableViewSizingAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := interfacetoolsdefs.TableViewSizingEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &interfacetoolsdefs.TableViewSizingAwareDeletePreviewActionResponse{
		Payload: fireback.GResponseSingleItem(preview),
	}, nil
}

func TableViewSizingAwareDeleteAction(c interfacetoolsdefs.TableViewSizingAwareDeleteActionRequest) (*interfacetoolsdefs.TableViewSizingAwareDeleteActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ActionRequires: []application.PermissionInfo{PERM_ROOT_TABLE_VIEW_SIZING_DELETE},
	})
	if err != nil {
		return nil, err
	}

	if err2 := interfacetoolsdefs.TableViewSizingEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &interfacetoolsdefs.TableViewSizingAwareDeleteActionResponse{
		Payload: fireback.GResponseSingleItem(struct{}{}),
	}, nil
}

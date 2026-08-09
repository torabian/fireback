package abac

import (
	"log"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"gorm.io/gorm"
)

// Capability's permission CompleteKeys keep their original "root.manage.fireback.capability.*"
// strings from when CapabilityEntity lived in modules/fireback - NewCrudPermissionSet would
// derive "root.manage.abac.capability.*" instead, which would silently invalidate every
// already-granted role/capability record in a live database, so these are declared by hand,
// unchanged, rather than through that helper.
var PERM_ROOT_CAPABILITY = application.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.*",
	Name:        "Entire capability actions (*)",
}
var PERM_ROOT_CAPABILITY_DELETE = application.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.delete",
	Name:        "Delete capability",
}
var PERM_ROOT_CAPABILITY_CREATE = application.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.create",
	Name:        "Create capability",
}
var PERM_ROOT_CAPABILITY_UPDATE = application.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.update",
	Name:        "Update capability",
}
var PERM_ROOT_CAPABILITY_QUERY = application.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.query",
	Name:        "Query capability",
}
var ALL_CAPABILITY_PERMISSIONS = []application.PermissionInfo{
	PERM_ROOT_CAPABILITY_DELETE,
	PERM_ROOT_CAPABILITY_CREATE,
	PERM_ROOT_CAPABILITY_UPDATE,
	PERM_ROOT_CAPABILITY_QUERY,
	PERM_ROOT_CAPABILITY,
}

// CapabilityUpsertPermissionFn is fireback.UpsertPermission's real, CapabilityEntity-backed
// body (moved verbatim from the old modules/fireback/CoreUtils.go's UpsertPermission) -
// wired into fireback.UpsertPermission from WorkspaceModuleSetup below, the same
// injection-point pattern AuthorizeRequest/WithAuthorizationFn/WithAuthorizationPure/
// WithSocketAuthorization already use there.
func CapabilityUpsertPermissionFn(permInfo *application.PermissionInfo, hasChildren bool, db *gorm.DB) {
	var entity *CapabilityEntity = nil
	perm := permInfo.CompleteKey

	if hasChildren {
		perm = perm + ".*"
	}

	if db.Where(CapabilityEntity{UniqueId: perm}).First(&entity).Error != nil {
		err := db.Create(&CapabilityEntity{
			UniqueId:    perm,
			Description: permInfo.Description,
			Name:        permInfo.Name,
		}).Error

		if err != nil {
			log.Fatalln("Cannot start the app because a permission creation failed.", perm, err)
		}
	}
}

// GetCapabilitiesAction keeps its original name (not CapabilityBrowseAction) since that's
// what FirebackModule.go used to wire it under, and other code may already reference it.
//
// Unlike every other entity in this package, CapabilityEntity is NOT workspace-scoped (it's
// a single global catalog, no workspace_id/user_id column at all) - so, same as the original
// implementation, this calls CapabilityEntityActions.Browse (the entity's own generated,
// unscoped query) directly instead of going through EntityActionsBundle/QueryEntitiesPointer,
// which unconditionally filters by workspace_id and would fail with "column workspace_id does
// not exist" on this table.
func GetCapabilitiesAction(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
	})
	if err != nil {
		return nil, err
	}

	res, qrm, err2 := CapabilityEntityActions.Browse(fireback.GetDbRef(), CapabilityBrowseActionQuery{}, "")
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &CapabilityBrowseActionResponse{
		Payload: fireback.GResponseQuery(res, &fireback.QueryResultMeta{
			TotalItems: qrm.TotalItems,
			Cursor:     qrm.Cursor,
		}, query),
	}, nil
}

// See GetCapabilitiesAction - not workspace-scoped, so this calls the entity's own
// generated, unscoped CapabilityEntityActions.Get directly.
func CapabilityGetAction(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
	}); err != nil {
		return nil, err
	}

	res, err2 := CapabilityEntityActions.Get(fireback.GetDbRef(), c.Params.UniqueId)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &CapabilityGetActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}

// CapabilityCreateAction inserts a new row directly into the capability catalog.
// AllowOnRoot is deliberate (not just PERM_ROOT_CAPABILITY_CREATE on its own): unlike
// every other entity in this package, CapabilityEntity has no workspace_id/user_id
// column at all, and CapabilityUpsertPermissionFn - the only other writer of this table -
// only ever runs during the root migration pass. Letting a non-root workspace insert
// rows here would pollute the single, global, cross-workspace catalog every
// RolePermissionTree in every workspace renders from, so creation is restricted to the
// root workspace the same way the seeded catalog itself is populated.
func CapabilityCreateAction(c CapabilityCreateActionRequest) (*CapabilityCreateActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_CAPABILITY_CREATE},
	}); err != nil {
		return nil, err
	}

	entity := &CapabilityEntity{
		Name:        c.Body.Name,
		Description: c.Body.Description,
	}
	// UniqueId is the entity's actual meaningful "key" (e.g. "students.*" below) rather
	// than an opaque id, so - unlike RoleCreateAction/WorkspaceTypeCreateAction, which
	// always let the db default generate it - it's taken straight from the body when the
	// caller sets it, falling back to gen_random_uuid() (see CapabilityEntity's gorm tag)
	// only when they don't.
	if v, ok := c.Body.UniqueId.Get(); ok && v != nil {
		entity.UniqueId = *v
	}

	created, err2 := CapabilityEntityActions.Create(fireback.GetDbRef(), entity)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &CapabilityCreateActionResponse{
		Payload: fireback.GResponseSingleItem(created),
	}, nil
}

// See GetCapabilitiesAction - not workspace-scoped, so this calls the entity's own
// generated, unscoped CapabilityEntityActions.Update directly.
func CapabilityUpdateAction(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_CAPABILITY_UPDATE},
	}); err != nil {
		return nil, err
	}

	res, err2 := CapabilityEntityActions.Update(fireback.GetDbRef(), c.Params.UniqueId, c.Body)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &CapabilityUpdateActionResponse{
		Payload: fireback.GResponseSingleItem(res),
	}, nil
}

func CapabilityAwareDeleteAction(c CapabilityAwareDeleteActionRequest) (*CapabilityAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []application.PermissionInfo{PERM_ROOT_CAPABILITY_DELETE},
	}); err != nil {
		return nil, err
	}

	if err2 := CapabilityEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &CapabilityAwareDeleteActionResponse{
		Payload: fireback.GResponseSingleItem(struct{}{}),
	}, nil
}

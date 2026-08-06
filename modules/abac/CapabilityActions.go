package abac

import (
	"log"

	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
)

// Capability's permission CompleteKeys keep their original "root.manage.fireback.capability.*"
// strings from when CapabilityEntity lived in modules/fireback - NewCrudPermissionSet would
// derive "root.manage.abac.capability.*" instead, which would silently invalidate every
// already-granted role/capability record in a live database, so these are declared by hand,
// unchanged, rather than through that helper.
var PERM_ROOT_CAPABILITY = fireback.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.*",
	Name:        "Entire capability actions (*)",
}
var PERM_ROOT_CAPABILITY_DELETE = fireback.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.delete",
	Name:        "Delete capability",
}
var PERM_ROOT_CAPABILITY_CREATE = fireback.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.create",
	Name:        "Create capability",
}
var PERM_ROOT_CAPABILITY_UPDATE = fireback.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.update",
	Name:        "Update capability",
}
var PERM_ROOT_CAPABILITY_QUERY = fireback.PermissionInfo{
	CompleteKey: "root.manage.fireback.capability.query",
	Name:        "Query capability",
}
var ALL_CAPABILITY_PERMISSIONS = []fireback.PermissionInfo{
	PERM_ROOT_CAPABILITY_DELETE,
	PERM_ROOT_CAPABILITY_CREATE,
	PERM_ROOT_CAPABILITY_UPDATE,
	PERM_ROOT_CAPABILITY_QUERY,
	PERM_ROOT_CAPABILITY,
}

var CapabilityActions = NewEntityActionsBundle[CapabilityEntity]()

// CapabilityUpsertPermissionFn is fireback.UpsertPermission's real, CapabilityEntity-backed
// body (moved verbatim from the old modules/fireback/CoreUtils.go's UpsertPermission) -
// wired into fireback.UpsertPermission from WorkspaceModuleSetup below, the same
// injection-point pattern AuthorizeRequest/WithAuthorizationFn/WithAuthorizationPure/
// WithSocketAuthorization already use there.
func CapabilityUpsertPermissionFn(permInfo *fireback.PermissionInfo, hasChildren bool, db *gorm.DB) {
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
func GetCapabilitiesAction(c CapabilityBrowseActionRequest) (*CapabilityBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []fireback.PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
	})
	if err != nil {
		return nil, err
	}

	items, qrm, err2 := CapabilityActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}

	return &CapabilityBrowseActionResponse{
		Payload: fireback.GResponseQuery(items, qrm, query),
	}, nil
}

func CapabilityGetAction(c CapabilityGetActionRequest) (*CapabilityGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []fireback.PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
	})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId

	item, err2 := CapabilityActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}

	return &CapabilityGetActionResponse{
		Payload: fireback.GResponseSingleItem(item),
	}, nil
}

func CapabilityUpdateAction(c CapabilityUpdateActionRequest) (*CapabilityUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []fireback.PermissionInfo{PERM_ROOT_CAPABILITY_UPDATE},
	})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId

	fields := &CapabilityEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Name.Get(); ok {
		fields.Name = *v
	}
	if v, ok := c.Body.Description.Get(); ok {
		fields.Description = *v
	}

	updated, err2 := CapabilityActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}

	return &CapabilityUpdateActionResponse{
		Payload: fireback.GResponseSingleItem(updated),
	}, nil
}

func CapabilityAwareDeleteAction(c CapabilityAwareDeleteActionRequest) (*CapabilityAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []fireback.PermissionInfo{PERM_ROOT_CAPABILITY_DELETE},
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

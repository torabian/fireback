package interfacetools

import (
	"reflect"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// CrudPermissionSet and NewEntityActionsBundle/EntityActionsBundle below are a
// deliberate, small duplication of modules/abac/GenericEntityActions.go's identically
// named helpers - see modules/abac/messaging/GenericEntityActions.go's copy of this
// same comment for the full rationale (abac needs to call into this package, so this
// package importing abac back for these two generic helpers would be an import cycle).
// They're pure generic scaffolding (no abac-specific types involved), so a second (now
// third) copy is low-risk; if this pattern grows further, promoting all three to
// fireback core would be worth it instead of yet another copy.

// CrudPermissionSet mirrors the 5 permissions (query/create/update/delete/wildcard) the
// old Module3 entity compiler generated for every entity. prefix is "root.modules" or
// "root.manage"; kebabName/displayName should match exactly what the old generated code
// used, to preserve existing role/capability records in any live database.
type CrudPermissionSet struct {
	Wildcard application.PermissionInfo
	Query    application.PermissionInfo
	Create   application.PermissionInfo
	Update   application.PermissionInfo
	Delete   application.PermissionInfo
	All      []application.PermissionInfo
}

// NewCrudPermissionSet builds the 5 CompleteKeys exactly as the old Module3-generated
// <Entity>Entity.dyno.go files did: prefix + ".abac." (the module name segment - kept as
// "abac" even here, unchanged, since these entities' CompleteKeys were already granted
// under that module name before moving to this package - changing it would silently
// invalidate every existing role/capability record) + kebabName + the operation.
func NewCrudPermissionSet(prefix, kebabName, displayName string) CrudPermissionSet {
	base := prefix + ".abac." + kebabName
	w := application.PermissionInfo{CompleteKey: base + ".*", Name: "Entire " + displayName + " actions (*)"}
	q := application.PermissionInfo{CompleteKey: base + ".query", Name: "Query " + displayName}
	c := application.PermissionInfo{CompleteKey: base + ".create", Name: "Create " + displayName}
	u := application.PermissionInfo{CompleteKey: base + ".update", Name: "Update " + displayName}
	d := application.PermissionInfo{CompleteKey: base + ".delete", Name: "Delete " + displayName}
	return CrudPermissionSet{
		Wildcard: w, Query: q, Create: c, Update: u, Delete: d,
		All: []application.PermissionInfo{d, c, u, q, w},
	}
}

// EntityActionsBundle mirrors the exact shape of the <Entity>ActionsSig bundle the old
// Module3 entity compiler used to generate for every entity.
type EntityActionsBundle[T any] struct {
	Update         func(query fireback.QueryDSL, dto *T) (*T, *fireback.IError)
	Create         func(dto *T, query fireback.QueryDSL) (*T, *fireback.IError)
	Upsert         func(dto *T, query fireback.QueryDSL) (*T, *fireback.IError)
	SeederInit     func() *T
	RemoveEnqueue  func(request fireback.DeleteRequest, query fireback.QueryDSL) (*fireback.DeleteResponse, *fireback.IError)
	MultiInsert    func(dtos []*T, query fireback.QueryDSL) ([]*T, *fireback.IError)
	GetOne         func(query fireback.QueryDSL) (*T, *fireback.IError)
	GetByWorkspace func(query fireback.QueryDSL) (*T, *fireback.IError)
	Query          func(query fireback.QueryDSL) ([]*T, *fireback.QueryResultMeta, *fireback.IError)
}

// NewEntityActionsBundle builds an EntityActionsBundle[T] on top of fireback's generic,
// tenant-scoped CRUD helpers - see modules/abac/GenericEntityActions.go's own copy for
// the full rationale of each method.
func NewEntityActionsBundle[T any]() EntityActionsBundle[T] {
	refl := func() reflect.Value {
		var zero T
		return reflect.ValueOf(&zero)
	}

	return EntityActionsBundle[T]{
		Create: func(dto *T, query fireback.QueryDSL) (*T, *fireback.IError) {
			created, err := fireback.CreateEntity(*dto)
			if err != nil {
				return nil, err
			}
			return &created, nil
		},
		Update: func(query fireback.QueryDSL, dto *T) (*T, *fireback.IError) {
			return fireback.UpdateEntity(query, dto)
		},
		Upsert: func(dto *T, query fireback.QueryDSL) (*T, *fireback.IError) {
			return nil, nil
		},
		SeederInit: func() *T {
			var zero T
			return &zero
		},
		RemoveEnqueue: func(request fireback.DeleteRequest, query fireback.QueryDSL) (*fireback.DeleteResponse, *fireback.IError) {
			return fireback.RemoveEntityEnqueue[T](request, query, refl())
		},
		MultiInsert: func(dtos []*T, query fireback.QueryDSL) ([]*T, *fireback.IError) {
			if len(dtos) == 0 {
				return dtos, nil
			}
			dbref := query.Tx
			if dbref == nil {
				dbref = fireback.GetDbRef()
			}
			if err := dbref.Create(&dtos).Error; err != nil {
				return nil, fireback.GormErrorToIError(err)
			}
			return dtos, nil
		},
		GetOne: func(query fireback.QueryDSL) (*T, *fireback.IError) {
			return fireback.GetOneEntity[T](query, refl())
		},
		GetByWorkspace: func(query fireback.QueryDSL) (*T, *fireback.IError) {
			return fireback.GetOneByWorkspaceEntity[T](query, refl())
		},
		Query: func(query fireback.QueryDSL) ([]*T, *fireback.QueryResultMeta, *fireback.IError) {
			return fireback.QueryEntitiesPointer[T](query, refl())
		},
	}
}

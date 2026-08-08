package abac

import (
	"reflect"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// StreamEntityQuery mirrors the paginated streaming iterator (<Entity>EntityStream) the
// old Module3 entity compiler generated for entities that needed to walk every row
// (e.g. CheckPassportMethodsAction.go's PassportMethodEntityStream usage) - queryFn is
// typically an EntityActionsBundle[T]'s own Query method.
func StreamEntityQuery[T any](queryFn func(fireback.QueryDSL) ([]*T, *fireback.QueryResultMeta, *fireback.IError), q fireback.QueryDSL) (chan []*T, *fireback.QueryResultMeta, *fireback.IError) {
	cn := make(chan []*T)
	q.ItemsPerPage = 50
	q.StartIndex = 0
	_, qrm, err := queryFn(q)
	if err != nil {
		return nil, nil, err
	}
	go func() {
		defer close(cn)
		for i := 0; i <= int(qrm.TotalAvailableItems)-1; i++ {
			items, _, _ := queryFn(q)
			i += q.ItemsPerPage
			q.StartIndex = i
			cn <- items
		}
	}()
	return cn, qrm, nil
}

// CrudPermissionSet mirrors the 5 permissions (query/create/update/delete/wildcard) the
// old Module3 entity compiler generated for every entity. prefix is "root.modules" or
// "root.manage" (the entity's permRewrite, if any); kebabName/displayName should match
// exactly what the old generated code used, to preserve existing role/capability
// records in any live database.
type CrudPermissionSet struct {
	Wildcard application.PermissionInfo
	Query    application.PermissionInfo
	Create   application.PermissionInfo
	Update   application.PermissionInfo
	Delete   application.PermissionInfo
	All      []application.PermissionInfo
}

// NewCrudPermissionSet builds the 5 CompleteKeys exactly as the old Module3-generated
// <Entity>Entity.dyno.go files did: prefix ("root.modules" or "root.manage", depending
// on whether the entity had permRewrite) + ".abac." (the module name segment - every
// entity in this package lives in the "abac" module) + kebabName + the operation.
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
// Module3 entity compiler used to generate for every entity (Update/Create/Upsert/
// SeederInit/RemoveEnqueue/MultiInsert/GetOne/GetByWorkspace/Query). A number of
// hand-written call sites across the codebase (AcceptInviteAction.go, AuthFlow.go, ...)
// already call into these exact method names/signatures on a `var <Entity>Actions`
// value, so entities migrated to Abac.emi.yml keep offering the same bundle - built
// generically here once, instead of per-entity codegen - and none of those callers need
// to change.
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
// tenant-scoped CRUD helpers (CreateEntity/UpdateEntity/GetOneEntity/
// GetOneByWorkspaceEntity/QueryEntitiesPointer/RemoveEntityEnqueue) - the same helpers
// the old Module3-generated code called internally. Upsert is intentionally a no-op
// returning (nil, nil): the old generated code never implemented it either (every
// entity's TableViewSizingActionUpsertFn-equivalent was always a stub).
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

package abac

import (
	"reflect"

	seeders "github.com/torabian/fireback/modules/abac/seeders/PassportMethod"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

func PassportMethodSyncSeeders() {
	fireback.SeederFromFSImport(
		fireback.QueryDSL{WorkspaceId: fireback.USER_SYSTEM},
		PassportMethodActions.Create,
		reflect.ValueOf(&PassportMethodEntity{}).Elem(),
		&seeders.ViewsFs,
		[]string{},
		true,
	)
}

var PassportMethodType = &xPassportMethodType{Email: "email", Phone: "phone", Google: "google", Facebook: "facebook"}

type xPassportMethodType struct {
	Email    string
	Phone    string
	Google   string
	Facebook string
}

var PassportMethodRegion = &xPassportMethodRegion{Global: "global"}

type xPassportMethodRegion struct {
	Global string
}

var passportMethodPerms = NewCrudPermissionSet("root.manage", "passport-method", "passport method")
var PERM_ROOT_PASSPORT_METHOD = passportMethodPerms.Wildcard
var PERM_ROOT_PASSPORT_METHOD_QUERY = passportMethodPerms.Query
var PERM_ROOT_PASSPORT_METHOD_CREATE = passportMethodPerms.Create
var PERM_ROOT_PASSPORT_METHOD_UPDATE = passportMethodPerms.Update
var PERM_ROOT_PASSPORT_METHOD_DELETE = passportMethodPerms.Delete
var ALL_PASSPORT_METHOD_PERMISSIONS = passportMethodPerms.All

var PassportMethodActions = NewEntityActionsBundle[PassportMethodEntity]()

func PassportMethodEntityStream(q fireback.QueryDSL) (chan []*PassportMethodEntity, *fireback.QueryResultMeta, *fireback.IError) {
	return StreamEntityQuery(PassportMethodActions.Query, q)
}

// passportMethod's security block was { writeOnRoot: true, readOnRoot: true, resolveStrategy: workspace },
// which the old generated code translated into AllowOnRoot: true, ResolveStrategy: "workspace" uniformly
// across every action - preserved here exactly.
func passportMethodSecurity(perm application.PermissionInfo) *fireback.SecurityModel {
	return &fireback.SecurityModel{
		ActionRequires:  []application.PermissionInfo{perm},
		ResolveStrategy: fireback.ResolveStrategyWorkspace,
		AllowOnRoot:     true,
	}
}

func PassportMethodBrowseAction(c PassportMethodBrowseActionRequest) (*PassportMethodBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, passportMethodSecurity(PERM_ROOT_PASSPORT_METHOD_QUERY))
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PassportMethodActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PassportMethodBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PassportMethodGetAction(c PassportMethodGetActionRequest) (*PassportMethodGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, passportMethodSecurity(PERM_ROOT_PASSPORT_METHOD_QUERY))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PassportMethodActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PassportMethodGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PassportMethodCreateAction(c PassportMethodCreateActionRequest) (*PassportMethodCreateActionResponse, error) {

	if err := fireback.CommonStructValidatorPointer(&c.Body, false); !fireback.IsNilish(err) {
		return nil, err
	}

	query, err := fireback.ResolveActionContext(c, passportMethodSecurity(PERM_ROOT_PASSPORT_METHOD_CREATE))
	if err != nil {
		return nil, err
	}
	entity := &PassportMethodEntity{
		Type:      c.Body.Type,
		Region:    c.Body.Region,
		ClientKey: c.Body.ClientKey,
	}
	created, err2 := PassportMethodActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &PassportMethodCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PassportMethodUpdateAction(c PassportMethodUpdateActionRequest) (*PassportMethodUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, passportMethodSecurity(PERM_ROOT_PASSPORT_METHOD_UPDATE))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &PassportMethodEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Type.Get(); ok {
		fields.Type = *v
	}
	if v, ok := c.Body.Region.Get(); ok {
		fields.Region = *v
	}
	if v, ok := c.Body.ClientKey.Get(); ok {
		fields.ClientKey = *v
	}
	updated, err2 := PassportMethodActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &PassportMethodUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PassportMethodAwareDeletePreviewAction(c PassportMethodAwareDeletePreviewActionRequest) (*PassportMethodAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, passportMethodSecurity(PERM_ROOT_PASSPORT_METHOD_DELETE)); err != nil {
		return nil, err
	}
	uniqueIds := PassportMethodAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := PassportMethodEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PassportMethodAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PassportMethodAwareDeleteAction(c PassportMethodAwareDeleteActionRequest) (*PassportMethodAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, passportMethodSecurity(PERM_ROOT_PASSPORT_METHOD_DELETE)); err != nil {
		return nil, err
	}
	if err2 := PassportMethodEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PassportMethodAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

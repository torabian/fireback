package abac

import (
	"reflect"

	"github.com/torabian/emi/emigo"
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

// passportMethodDuplicateCheck rejects a (type, region) pair another row already
// occupies - the pair must be globally unique, never duplicated, regardless of
// workspace (passport methods are root-only content in the first place - see
// passportMethodSecurity). excludeUniqueId lets an update skip the row being updated
// itself, so a no-op update (or one that doesn't touch type/region) doesn't trip on
// the row's own existing values.
func passportMethodDuplicateCheck(methodType string, region string, excludeUniqueId string) *fireback.IError {
	db := fireback.GetDbRef().Model(&PassportMethodEntity{}).
		Where(&PassportMethodEntity{Type: methodType, Region: region})

	if excludeUniqueId != "" {
		db = db.Where("unique_id <> ?", excludeUniqueId)
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return fireback.GormErrorToIError(err)
	}

	if count > 0 {
		return &fireback.IError{
			HttpCode: 400,
			Message:  AbacMessages.PassportMethodAlreadyExists,
		}
	}

	return nil
}

func PassportMethodCreateAction(c PassportMethodCreateActionRequest) (*PassportMethodCreateActionResponse, error) {

	if err := fireback.CommonStructValidatorPointer(&c.Body, false); !fireback.IsNilish(err) {
		return nil, err
	}

	query, err := fireback.ResolveActionContext(c, passportMethodSecurity(PERM_ROOT_PASSPORT_METHOD_CREATE))
	if err != nil {
		return nil, err
	}

	if dupErr := passportMethodDuplicateCheck(c.Body.Type, c.Body.Region, ""); dupErr != nil {
		return nil, dupErr
	}

	entity := &PassportMethodEntity{
		Type:        c.Body.Type,
		Region:      c.Body.Region,
		ClientKey:   c.Body.ClientKey,
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
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

	newType, typeChanging := c.Body.Type.Get()
	if typeChanging {
		fields.Type = *newType
	}
	newRegion, regionChanging := c.Body.Region.Get()
	if regionChanging {
		fields.Region = *newRegion
	}
	if v, ok := c.Body.ClientKey.Get(); ok {
		fields.ClientKey = *v
	}

	// Only re-check uniqueness when type and/or region are actually part of this
	// patch - an update that leaves both untouched can't create a new duplicate.
	// Whichever of the two isn't being changed still needs to be read off the
	// current row, so the pair being checked reflects what the row will actually
	// end up as, not just the half that was patched.
	if typeChanging || regionChanging {
		checkType, checkRegion := fields.Type, fields.Region
		if !typeChanging || !regionChanging {
			current, currErr := PassportMethodActions.GetOne(*query)
			if currErr != nil {
				return nil, currErr
			}
			if !typeChanging {
				checkType = current.Type
			}
			if !regionChanging {
				checkRegion = current.Region
			}
		}
		if dupErr := passportMethodDuplicateCheck(checkType, checkRegion, c.Params.UniqueId); dupErr != nil {
			return nil, dupErr
		}
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

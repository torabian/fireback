package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

// passport had permRewrite root.modules -> root.manage, and security: { writeOnRoot: true }
// in the old yaml: the old generated code set AllowOnRoot: true on Create/Update/Delete
// only, leaving Query/Get plain - preserved here exactly.
var passportPerms = NewCrudPermissionSet("root.manage", "passport", "passport")
var PERM_ROOT_PASSPORT = passportPerms.Wildcard
var PERM_ROOT_PASSPORT_QUERY = passportPerms.Query
var PERM_ROOT_PASSPORT_CREATE = passportPerms.Create
var PERM_ROOT_PASSPORT_UPDATE = passportPerms.Update
var PERM_ROOT_PASSPORT_DELETE = passportPerms.Delete
var ALL_PASSPORT_PERMISSIONS = passportPerms.All

var PassportActions = NewEntityActionsBundle[PassportEntity]()

var PASSPORT_METHOD_EMAIL = "email"
var PASSPORT_METHOD_PHONE = "phone"

// PassportTypes recovered from the pre-migration PassportEntity.go hand file.
var PassportTypes = newPassportFactory()

func newPassportFactory() *passportType {
	return &passportType{
		EmailPassword: "EmailPassword",
		Google:        "Google",
		PhoneNumber:   "PhoneNumber",
	}
}

type passportType struct {
	EmailPassword string
	Google        string
	PhoneNumber   string
}

func GetUserByPassport2(value string) (*UserEntity, *PassportEntity, *fireback.IError) {
	passport := &PassportEntity{}
	err := fireback.GetDbRef().Where(&PassportEntity{Value: value}).First(passport).Error

	if err != nil {
		return nil, nil, fireback.GormErrorToIError(err)
	}

	user := &UserEntity{}
	fireback.GetDbRef().Where(&UserEntity{UniqueId: passport.UserId.OrDefault("")}).First(user)

	return user, passport, nil
}

func GetUserByPassport(value string) (*UserEntity, *PassportEntity, error) {
	passport := &PassportEntity{}
	fireback.GetDbRef().Where(&PassportEntity{Value: value}).First(passport)

	user := &UserEntity{}
	fireback.GetDbRef().Where(&UserEntity{UniqueId: passport.UserId.OrDefault("")}).First(user)

	return user, passport, nil
}

func PassportBrowseAction(c PassportBrowseActionRequest) (*PassportBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PassportActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PassportBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PassportGetAction(c PassportGetActionRequest) (*PassportGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PassportActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PassportGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PassportCreateAction(c PassportCreateActionRequest) (*PassportCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	// type/value are validate:"required" (see PassportDto.go) but nothing was actually
	// enforcing it - same fix as
	// modules/abac/messaging/EmailProviderActions.go's EmailProviderCreateAction /
	// modules/abac/interfacetools's TableViewSizingCreateAction.
	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}
	entity := &PassportEntity{
		ThirdPartyVerifier: c.Body.ThirdPartyVerifier,
		Type:               c.Body.Type,
		UserId:             c.Body.UserId,
		Value:              c.Body.Value,
		TotpSecret:         c.Body.TotpSecret,
		TotpConfirmed:      c.Body.TotpConfirmed,
		Password:           c.Body.Password,
		Confirmed:          c.Body.Confirmed,
		AccessToken:        c.Body.AccessToken,
		// AllowOnRoot above means query.WorkspaceId is always "root" here anyway (any
		// other workspace is rejected before reaching this point) - matches "Passport
		// and user always belong to the root workspace" elsewhere (see
		// UnsafeGenerateUser). Without this, the row is invisible to
		// PassportBrowseAction's workspace-scoped query (see GetSqlContext) - same
		// workspace-stamping fix as EmailConfirmationCreateAction/
		// AppMenuCreateAction/TimezoneGroupCreateAction.
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := PassportActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &PassportCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PassportUpdateAction(c PassportUpdateActionRequest) (*PassportUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_UPDATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &PassportEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.ThirdPartyVerifier.Get(); ok {
		fields.ThirdPartyVerifier = *v
	}
	if v, ok := c.Body.Type.Get(); ok {
		fields.Type = *v
	}
	if v, ok := c.Body.UserId.Get(); ok {
		fields.UserId.Set(v)
	}
	if v, ok := c.Body.Value.Get(); ok {
		fields.Value = *v
	}
	if v, ok := c.Body.TotpSecret.Get(); ok {
		fields.TotpSecret = *v
	}
	if v, ok := c.Body.TotpConfirmed.Get(); ok {
		fields.TotpConfirmed.Set(v)
	}
	if v, ok := c.Body.Password.Get(); ok {
		fields.Password = *v
	}
	if v, ok := c.Body.Confirmed.Get(); ok {
		fields.Confirmed.Set(v)
	}
	if v, ok := c.Body.AccessToken.Get(); ok {
		fields.AccessToken = *v
	}
	updated, err2 := PassportActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &PassportUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PassportAwareDeletePreviewAction(c PassportAwareDeletePreviewActionRequest) (*PassportAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := PassportAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := PassportEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PassportAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PassportAwareDeleteAction(c PassportAwareDeleteActionRequest) (*PassportAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_PASSPORT_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := PassportEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PassportAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

// publicAuthentication's old security block was { writeOnRoot: true }: the old generated
// code only set AllowOnRoot: true on Create/Update/Delete, leaving Query/Get plain -
// preserved here exactly.
var publicAuthenticationPerms = NewCrudPermissionSet("root.manage", "public-authentication", "public authentication")
var PERM_ROOT_PUBLIC_AUTHENTICATION = publicAuthenticationPerms.Wildcard
var PERM_ROOT_PUBLIC_AUTHENTICATION_QUERY = publicAuthenticationPerms.Query
var PERM_ROOT_PUBLIC_AUTHENTICATION_CREATE = publicAuthenticationPerms.Create
var PERM_ROOT_PUBLIC_AUTHENTICATION_UPDATE = publicAuthenticationPerms.Update
var PERM_ROOT_PUBLIC_AUTHENTICATION_DELETE = publicAuthenticationPerms.Delete
var ALL_PUBLIC_AUTHENTICATION_PERMISSIONS = publicAuthenticationPerms.All

var PublicAuthenticationActions = NewEntityActionsBundle[PublicAuthenticationEntity]()

func PublicAuthenticationBrowseAction(c PublicAuthenticationBrowseActionRequest) (*PublicAuthenticationBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PUBLIC_AUTHENTICATION_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := PublicAuthenticationActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PublicAuthenticationBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func PublicAuthenticationGetAction(c PublicAuthenticationGetActionRequest) (*PublicAuthenticationGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PUBLIC_AUTHENTICATION_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := PublicAuthenticationActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &PublicAuthenticationGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func PublicAuthenticationCreateAction(c PublicAuthenticationCreateActionRequest) (*PublicAuthenticationCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PUBLIC_AUTHENTICATION_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	entity := &PublicAuthenticationEntity{
		UserId:              c.Body.UserId,
		TotpSecret:          c.Body.TotpSecret,
		TotpLink:            c.Body.TotpLink,
		PassportId:          c.Body.PassportId,
		SessionSecret:       c.Body.SessionSecret,
		PassportValue:       c.Body.PassportValue,
		IsInCreationProcess: c.Body.IsInCreationProcess,
		Status:              c.Body.Status,
		BlockedUntil:        c.Body.BlockedUntil,
		Otp:                 c.Body.Otp,
		RecoveryAbsoluteUrl: c.Body.RecoveryAbsoluteUrl,
		WorkspaceId:         c.Body.WorkspaceId,
	}
	created, err2 := PublicAuthenticationActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &PublicAuthenticationCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func PublicAuthenticationUpdateAction(c PublicAuthenticationUpdateActionRequest) (*PublicAuthenticationUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PUBLIC_AUTHENTICATION_UPDATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &PublicAuthenticationEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.UserId.Get(); ok {
		fields.UserId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.TotpSecret.Get(); ok {
		fields.TotpSecret = *v
	}
	if v, ok := c.Body.TotpLink.Get(); ok {
		fields.TotpLink = *v
	}
	if v, ok := c.Body.PassportId.Get(); ok {
		fields.PassportId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.SessionSecret.Get(); ok {
		fields.SessionSecret = *v
	}
	if v, ok := c.Body.PassportValue.Get(); ok {
		fields.PassportValue = *v
	}
	if v, ok := c.Body.IsInCreationProcess.Get(); ok {
		fields.IsInCreationProcess = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.Status.Get(); ok {
		fields.Status = *v
	}
	if v, ok := c.Body.BlockedUntil.Get(); ok {
		fields.BlockedUntil = *v
	}
	if v, ok := c.Body.Otp.Get(); ok {
		fields.Otp = *v
	}
	if v, ok := c.Body.RecoveryAbsoluteUrl.Get(); ok {
		fields.RecoveryAbsoluteUrl = *v
	}
	if v, ok := c.Body.WorkspaceId.Get(); ok {
		fields.WorkspaceId = emigo.NullableOf(*v)
	}
	updated, err2 := PublicAuthenticationActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &PublicAuthenticationUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func PublicAuthenticationAwareDeletePreviewAction(c PublicAuthenticationAwareDeletePreviewActionRequest) (*PublicAuthenticationAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PUBLIC_AUTHENTICATION_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := PublicAuthenticationAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := PublicAuthenticationEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PublicAuthenticationAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func PublicAuthenticationAwareDeleteAction(c PublicAuthenticationAwareDeleteActionRequest) (*PublicAuthenticationAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{PERM_ROOT_PUBLIC_AUTHENTICATION_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := PublicAuthenticationEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &PublicAuthenticationAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

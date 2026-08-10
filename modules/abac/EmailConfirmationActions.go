package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
)

var emailConfirmationPerms = NewCrudPermissionSet("root.modules", "email-confirmation", "email confirmation")
var PERM_ROOT_EMAIL_CONFIRMATION = emailConfirmationPerms.Wildcard
var PERM_ROOT_EMAIL_CONFIRMATION_QUERY = emailConfirmationPerms.Query
var PERM_ROOT_EMAIL_CONFIRMATION_CREATE = emailConfirmationPerms.Create
var PERM_ROOT_EMAIL_CONFIRMATION_UPDATE = emailConfirmationPerms.Update
var PERM_ROOT_EMAIL_CONFIRMATION_DELETE = emailConfirmationPerms.Delete
var ALL_EMAIL_CONFIRMATION_PERMISSIONS = emailConfirmationPerms.All

var EmailConfirmationActions = NewEntityActionsBundle[abacdefs.EmailConfirmationEntity]()

func EmailConfirmationBrowseAction(c abacdefs.EmailConfirmationBrowseActionRequest) (*abacdefs.EmailConfirmationBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := EmailConfirmationActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.EmailConfirmationBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func EmailConfirmationGetAction(c abacdefs.EmailConfirmationGetActionRequest) (*abacdefs.EmailConfirmationGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := EmailConfirmationActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.EmailConfirmationGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func EmailConfirmationCreateAction(c abacdefs.EmailConfirmationCreateActionRequest) (*abacdefs.EmailConfirmationCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_CREATE}})
	if err != nil {
		return nil, err
	}
	entity := &abacdefs.EmailConfirmationEntity{
		UserId:    c.Body.UserId,
		Status:    c.Body.Status,
		Email:     c.Body.Email,
		Key:       c.Body.Key,
		ExpiresAt: c.Body.ExpiresAt,
		// Without this, the row is invisible to EmailConfirmationBrowseAction's
		// workspace-scoped query (see GetSqlContext) - same workspace-stamping fix as
		// modules/abac/messaging/EmailProviderActions.go's EmailProviderCreateAction /
		// modules/abac/interfacetools's AppMenuCreateAction/TimezoneGroupCreateAction.
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := EmailConfirmationActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.EmailConfirmationCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func EmailConfirmationUpdateAction(c abacdefs.EmailConfirmationUpdateActionRequest) (*abacdefs.EmailConfirmationUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &abacdefs.EmailConfirmationEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.UserId.Get(); ok {
		fields.UserId = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.Status.Get(); ok {
		fields.Status = *v
	}
	if v, ok := c.Body.Email.Get(); ok {
		fields.Email = *v
	}
	if v, ok := c.Body.Key.Get(); ok {
		fields.Key = *v
	}
	if v, ok := c.Body.ExpiresAt.Get(); ok {
		fields.ExpiresAt = *v
	}
	updated, err2 := EmailConfirmationActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.EmailConfirmationUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func EmailConfirmationAwareDeletePreviewAction(c abacdefs.EmailConfirmationAwareDeletePreviewActionRequest) (*abacdefs.EmailConfirmationAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.EmailConfirmationAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.EmailConfirmationEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.EmailConfirmationAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func EmailConfirmationAwareDeleteAction(c abacdefs.EmailConfirmationAwareDeleteActionRequest) (*abacdefs.EmailConfirmationAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_EMAIL_CONFIRMATION_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.EmailConfirmationEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.EmailConfirmationAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

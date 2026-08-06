package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

// workspaceConfig had permRewrite root.modules -> root.manage, distinctBy: workspace,
// and security: { writeOnRoot: true, readOnRoot: true, resolveStrategy: workspace } in
// the old yaml: the old generated code set ResolveStrategy: "workspace" and
// AllowOnRoot: true on every action (Query/Get/Create/Update/Delete) - preserved here
// exactly.
var workspaceConfigPerms = NewCrudPermissionSet("root.manage", "workspace-config", "workspace config")
var PERM_ROOT_WORKSPACE_CONFIG = workspaceConfigPerms.Wildcard
var PERM_ROOT_WORKSPACE_CONFIG_QUERY = workspaceConfigPerms.Query
var PERM_ROOT_WORKSPACE_CONFIG_CREATE = workspaceConfigPerms.Create
var PERM_ROOT_WORKSPACE_CONFIG_UPDATE = workspaceConfigPerms.Update
var PERM_ROOT_WORKSPACE_CONFIG_DELETE = workspaceConfigPerms.Delete
var ALL_WORKSPACE_CONFIG_PERMISSIONS = workspaceConfigPerms.All

var WorkspaceConfigActions = NewEntityActionsBundle[WorkspaceConfigEntity]()

func workspaceConfigSecurity(perm fireback.PermissionInfo) *fireback.SecurityModel {
	return &fireback.SecurityModel{
		ActionRequires:  []fireback.PermissionInfo{perm},
		ResolveStrategy: "workspace",
		AllowOnRoot:     true,
	}
}

func WorkspaceConfigBrowseAction(c WorkspaceConfigBrowseActionRequest) (*WorkspaceConfigBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_QUERY))
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WorkspaceConfigActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceConfigBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func WorkspaceConfigGetAction(c WorkspaceConfigGetActionRequest) (*WorkspaceConfigGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_QUERY))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := WorkspaceConfigActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceConfigGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WorkspaceConfigCreateAction(c WorkspaceConfigCreateActionRequest) (*WorkspaceConfigCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_CREATE))
	if err != nil {
		return nil, err
	}
	entity := &WorkspaceConfigEntity{
		EnableRecaptcha2:           c.Body.EnableRecaptcha2,
		EnableOtp:                  c.Body.EnableOtp,
		RequireOtpOnSignup:         c.Body.RequireOtpOnSignup,
		RequireOtpOnSignin:         c.Body.RequireOtpOnSignin,
		Recaptcha2ServerKey:        c.Body.Recaptcha2ServerKey,
		Recaptcha2ClientKey:        c.Body.Recaptcha2ClientKey,
		EnableTotp:                 c.Body.EnableTotp,
		ForceTotp:                  c.Body.ForceTotp,
		ForcePasswordOnPhone:       c.Body.ForcePasswordOnPhone,
		ForcePersonNameOnPhone:     c.Body.ForcePersonNameOnPhone,
		GeneralEmailProviderId:     c.Body.GeneralEmailProviderId,
		GeneralGsmProviderId:       c.Body.GeneralGsmProviderId,
		InviteToWorkspaceContentId: c.Body.InviteToWorkspaceContentId,
		EmailOtpContentId:          c.Body.EmailOtpContentId,
		SmsOtpContentId:            c.Body.SmsOtpContentId,
		WorkspaceId:                emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := WorkspaceConfigActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceConfigCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

// WorkspaceConfigUpdateAction upserts by workspace (distinctBy: workspace in the old
// yaml): Query.WorkspaceId (resolved from ResolveStrategy: workspace above) is the
// condition for finding-or-creating the row, not the uniqueId path param.
func WorkspaceConfigUpdateAction(c WorkspaceConfigUpdateActionRequest) (*WorkspaceConfigUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_UPDATE))
	if err != nil {
		return nil, err
	}

	dbref := fireback.GetDbRef()
	var item WorkspaceConfigEntity
	if err2 := dbref.
		Where(&WorkspaceConfigEntity{WorkspaceId: emigo.NullableOf(query.WorkspaceId)}).
		FirstOrCreate(&item).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	changes := map[string]interface{}{}
	if v, ok := c.Body.EnableRecaptcha2.Get(); ok {
		changes["EnableRecaptcha2"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.EnableOtp.Get(); ok {
		changes["EnableOtp"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.RequireOtpOnSignup.Get(); ok {
		changes["RequireOtpOnSignup"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.RequireOtpOnSignin.Get(); ok {
		changes["RequireOtpOnSignin"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.Recaptcha2ServerKey.Get(); ok {
		changes["Recaptcha2ServerKey"] = *v
	}
	if v, ok := c.Body.Recaptcha2ClientKey.Get(); ok {
		changes["Recaptcha2ClientKey"] = *v
	}
	if v, ok := c.Body.EnableTotp.Get(); ok {
		changes["EnableTotp"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.ForceTotp.Get(); ok {
		changes["ForceTotp"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.ForcePasswordOnPhone.Get(); ok {
		changes["ForcePasswordOnPhone"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.ForcePersonNameOnPhone.Get(); ok {
		changes["ForcePersonNameOnPhone"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.GeneralEmailProviderId.Get(); ok {
		changes["GeneralEmailProviderId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.GeneralGsmProviderId.Get(); ok {
		changes["GeneralGsmProviderId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.InviteToWorkspaceContentId.Get(); ok {
		changes["InviteToWorkspaceContentId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.EmailOtpContentId.Get(); ok {
		changes["EmailOtpContentId"] = emigo.NullableOf(*v)
	}
	if v, ok := c.Body.SmsOtpContentId.Get(); ok {
		changes["SmsOtpContentId"] = emigo.NullableOf(*v)
	}

	if len(changes) > 0 {
		if err2 := dbref.Model(&item).Updates(changes).Error; err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}
	}

	var updated WorkspaceConfigEntity
	if err2 := dbref.Where("id = ?", item.Id).First(&updated).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceConfigUpdateActionResponse{Payload: fireback.GResponseSingleItem(&updated)}, nil
}

func WorkspaceConfigAwareDeletePreviewAction(c WorkspaceConfigAwareDeletePreviewActionRequest) (*WorkspaceConfigAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_DELETE)); err != nil {
		return nil, err
	}
	uniqueIds := WorkspaceConfigAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := WorkspaceConfigEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceConfigAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WorkspaceConfigAwareDeleteAction(c WorkspaceConfigAwareDeleteActionRequest) (*WorkspaceConfigAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_DELETE)); err != nil {
		return nil, err
	}
	if err2 := WorkspaceConfigEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceConfigAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

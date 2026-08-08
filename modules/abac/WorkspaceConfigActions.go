package abac

import (
	"errors"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
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

// The two GET/PATCH /workspace-config/distinct permissions predate NewCrudPermissionSet
// (they're not part of the standard 5), so their CompleteKeys are declared by hand,
// unchanged from the old Module3-generated code.
var PERM_ROOT_WORKSPACE_CONFIG_GET_DISTINCT_WORKSPACE = fireback.PermissionInfo{
	CompleteKey: "root.manage.abac.workspace-config.get-distinct-workspace",
	Name:        "Get workspace config Distinct",
}
var PERM_ROOT_WORKSPACE_CONFIG_UPDATE_DISTINCT_WORKSPACE = fireback.PermissionInfo{
	CompleteKey: "root.manage.abac.workspace-config.update-distinct-workspace",
	Name:        "Update workspace config Distinct",
}
var ALL_WORKSPACE_CONFIG_PERMISSIONS = append(
	workspaceConfigPerms.All,
	PERM_ROOT_WORKSPACE_CONFIG_GET_DISTINCT_WORKSPACE,
	PERM_ROOT_WORKSPACE_CONFIG_UPDATE_DISTINCT_WORKSPACE,
)

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

// workspaceConfigChangesFromOptionalDto turns whichever fields are actually set on body
// into a gorm .Updates()-ready map - shared by every WorkspaceConfig update path
// (per-workspace and the root-only /distinct one below).
func workspaceConfigChangesFromOptionalDto(body WorkspaceConfigOptionalDto) map[string]interface{} {
	changes := map[string]interface{}{}
	if v, ok := body.EnableRecaptcha2.Get(); ok {
		changes["EnableRecaptcha2"] = emigo.NullableOf(*v)
	}
	if v, ok := body.EnableOtp.Get(); ok {
		changes["EnableOtp"] = emigo.NullableOf(*v)
	}
	if v, ok := body.RequireOtpOnSignup.Get(); ok {
		changes["RequireOtpOnSignup"] = emigo.NullableOf(*v)
	}
	if v, ok := body.RequireOtpOnSignin.Get(); ok {
		changes["RequireOtpOnSignin"] = emigo.NullableOf(*v)
	}
	if v, ok := body.Recaptcha2ServerKey.Get(); ok {
		changes["Recaptcha2ServerKey"] = *v
	}
	if v, ok := body.Recaptcha2ClientKey.Get(); ok {
		changes["Recaptcha2ClientKey"] = *v
	}
	if v, ok := body.EnableTotp.Get(); ok {
		changes["EnableTotp"] = emigo.NullableOf(*v)
	}
	if v, ok := body.ForceTotp.Get(); ok {
		changes["ForceTotp"] = emigo.NullableOf(*v)
	}
	if v, ok := body.ForcePasswordOnPhone.Get(); ok {
		changes["ForcePasswordOnPhone"] = emigo.NullableOf(*v)
	}
	if v, ok := body.ForcePersonNameOnPhone.Get(); ok {
		changes["ForcePersonNameOnPhone"] = emigo.NullableOf(*v)
	}
	if v, ok := body.GeneralEmailProviderId.Get(); ok {
		changes["GeneralEmailProviderId"] = emigo.NullableOf(*v)
	}
	if v, ok := body.GeneralGsmProviderId.Get(); ok {
		changes["GeneralGsmProviderId"] = emigo.NullableOf(*v)
	}
	if v, ok := body.InviteToWorkspaceContentId.Get(); ok {
		changes["InviteToWorkspaceContentId"] = emigo.NullableOf(*v)
	}
	if v, ok := body.EmailOtpContentId.Get(); ok {
		changes["EmailOtpContentId"] = emigo.NullableOf(*v)
	}
	if v, ok := body.SmsOtpContentId.Get(); ok {
		changes["SmsOtpContentId"] = emigo.NullableOf(*v)
	}
	return changes
}

// workspaceConfigUpsertByWorkspace is the shared upsert-by-workspace logic (distinctBy:
// workspace in the old yaml) behind WorkspaceConfigUpdateAction - Query.WorkspaceId
// (resolved from ResolveStrategy: workspace) is the condition for finding-or-creating the
// row, not a uniqueId path param.
func workspaceConfigUpsertByWorkspace(query fireback.QueryDSL, body WorkspaceConfigOptionalDto) (*WorkspaceConfigEntity, *fireback.IError) {
	dbref := fireback.GetDbRef()
	var item WorkspaceConfigEntity
	if err2 := dbref.
		Where(&WorkspaceConfigEntity{WorkspaceId: emigo.NullableOf(query.WorkspaceId)}).
		FirstOrCreate(&item).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	changes := workspaceConfigChangesFromOptionalDto(body)
	if len(changes) > 0 {
		if err2 := dbref.Model(&item).Updates(changes).Error; err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}
	}

	var updated WorkspaceConfigEntity
	if err2 := dbref.Where("id = ?", item.Id).First(&updated).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &updated, nil
}

func WorkspaceConfigUpdateAction(c WorkspaceConfigUpdateActionRequest) (*WorkspaceConfigUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_UPDATE))
	if err != nil {
		return nil, err
	}
	updated, err2 := workspaceConfigUpsertByWorkspace(*query, c.Body)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceConfigUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

// WorkspaceConfigDistinctGetAction restores the old Module3-generated GET
// /workspace-config/distinct route. Unlike every other WorkspaceConfig action, this one
// isn't scoped to whatever workspace the caller currently happens to be acting in - it
// always configures the root workspace specifically, so it deliberately bypasses
// WorkspaceConfigActions.GetByWorkspace (which filters by the caller's own
// query.WorkspaceId) and queries workspace_id = 'root' directly via plain gorm instead.
// No uniqueId is involved either way - the row is found purely by that filter. Returns
// an empty WorkspaceConfigEntity (not a 404) when none exists yet, matching the old
// behavior.
func WorkspaceConfigDistinctGetAction(c WorkspaceConfigDistinctGetActionRequest) (*WorkspaceConfigDistinctGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_GET_DISTINCT_WORKSPACE)); err != nil {
		return nil, err
	}

	var item WorkspaceConfigEntity
	err2 := fireback.GetDbRef().Where(&WorkspaceConfigEntity{WorkspaceId: emigo.NullableOf(ROOT_VAR)}).First(&item).Error
	if err2 != nil {
		if errors.Is(err2, gorm.ErrRecordNotFound) {
			return &WorkspaceConfigDistinctGetActionResponse{Payload: fireback.GResponseSingleItem(&WorkspaceConfigEntity{})}, nil
		}
		return nil, fireback.GormErrorToIError(err2)
	}

	return &WorkspaceConfigDistinctGetActionResponse{Payload: fireback.GResponseSingleItem(&item)}, nil
}

// WorkspaceConfigDistinctUpdateAction restores the old Module3-generated PATCH
// /workspace-config/distinct route. Same root-only reasoning as
// WorkspaceConfigDistinctGetAction above - always finds-or-creates the row scoped to the
// root workspace specifically (workspace_id = 'root'), via plain gorm, regardless of the
// caller's own current workspace, and needs no uniqueId to do it.
func WorkspaceConfigDistinctUpdateAction(c WorkspaceConfigDistinctUpdateActionRequest) (*WorkspaceConfigDistinctUpdateActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_UPDATE_DISTINCT_WORKSPACE)); err != nil {
		return nil, err
	}

	db := fireback.GetDbRef()
	var item WorkspaceConfigEntity
	if err2 := db.
		Where(&WorkspaceConfigEntity{WorkspaceId: emigo.NullableOf(ROOT_VAR)}).
		FirstOrCreate(&item).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	changes := workspaceConfigChangesFromOptionalDto(c.Body)
	if len(changes) > 0 {
		if err2 := db.Model(&item).Updates(changes).Error; err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}
	}

	var updated WorkspaceConfigEntity
	if err2 := db.Where("id = ?", item.Id).First(&updated).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &WorkspaceConfigDistinctUpdateActionResponse{Payload: fireback.GResponseSingleItem(&updated)}, nil
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

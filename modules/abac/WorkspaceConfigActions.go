package abac

import (
	"errors"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
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
var PERM_ROOT_WORKSPACE_CONFIG_GET_DISTINCT_WORKSPACE = application.PermissionInfo{
	CompleteKey: "root.manage.abac.workspace-config.get-distinct-workspace",
	Name:        "Get workspace config Distinct",
}
var PERM_ROOT_WORKSPACE_CONFIG_UPDATE_DISTINCT_WORKSPACE = application.PermissionInfo{
	CompleteKey: "root.manage.abac.workspace-config.update-distinct-workspace",
	Name:        "Update workspace config Distinct",
}
var ALL_WORKSPACE_CONFIG_PERMISSIONS = append(
	workspaceConfigPerms.All,
	PERM_ROOT_WORKSPACE_CONFIG_GET_DISTINCT_WORKSPACE,
	PERM_ROOT_WORKSPACE_CONFIG_UPDATE_DISTINCT_WORKSPACE,
)

var WorkspaceConfigActions = NewEntityActionsBundle[abacdefs.WorkspaceConfigEntity]()

func workspaceConfigSecurity(perm application.PermissionInfo) *fireback.SecurityModel {
	return &fireback.SecurityModel{
		ActionRequires:  []application.PermissionInfo{perm},
		ResolveStrategy: "workspace",
		AllowOnRoot:     true,
	}
}

func WorkspaceConfigBrowseAction(c abacdefs.WorkspaceConfigBrowseActionRequest) (*abacdefs.WorkspaceConfigBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_QUERY))
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WorkspaceConfigActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceConfigBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func WorkspaceConfigGetAction(c abacdefs.WorkspaceConfigGetActionRequest) (*abacdefs.WorkspaceConfigGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_QUERY))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := WorkspaceConfigActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceConfigGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WorkspaceConfigCreateAction(c abacdefs.WorkspaceConfigCreateActionRequest) (*abacdefs.WorkspaceConfigCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_CREATE))
	if err != nil {
		return nil, err
	}
	entity := &abacdefs.WorkspaceConfigEntity{
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
	return &abacdefs.WorkspaceConfigCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

// workspaceConfigChangesFromOptionalDto turns whichever fields are actually set on body
// into a gorm .Updates()-ready map - shared by every WorkspaceConfig update path
// (per-workspace and the root-only /distinct one below).
//
// Bug fix: this used to do `if v, ok := body.X.Get(); ok { changes[...] = emigo.NullableOf(*v) }`
// for every Nullable[bool]/Nullable[string] field. Get()'s `ok` is true both when the
// caller sent a real value AND when they explicitly sent JSON `null` for the field (that's
// the whole point of Nullable[T] - it distinguishes "omitted" from "set to null") - in the
// null case v itself is nil, so `*v` panicked with a nil pointer dereference. Since there's
// no gin.Recovery() registered on this server, that panic wasn't caught by gin at all - it
// unwound all the way to net/http's own per-connection recovery, which just closes the
// socket without writing any response ("empty reply from server" to the client) rather than
// a clean 500. Any WorkspaceConfig PATCH that included an explicit `null` for one of these
// fields (a common shape for a form that hasn't touched a given field yet) hit this.
// Fixed by switching to the same IsSet()-and-pass-the-whole-Nullable-through pattern the
// sibling generated WorkspaceConfigEntityUpdateFn already uses correctly: Nullable[T]
// implements driver.Valuer, so gorm calls its Value() (nil-safe by construction) to resolve
// the SQL value - including for the plain `string` Recaptcha2*Key columns, which accept a
// Nullable[string] here the same way. No manual dereference left anywhere in this function.
func workspaceConfigChangesFromOptionalDto(body abacdefs.WorkspaceConfigOptionalDto) map[string]interface{} {
	changes := map[string]interface{}{}
	if body.EnableRecaptcha2.IsSet() {
		changes["EnableRecaptcha2"] = body.EnableRecaptcha2
	}
	if body.EnableOtp.IsSet() {
		changes["EnableOtp"] = body.EnableOtp
	}
	if body.RequireOtpOnSignup.IsSet() {
		changes["RequireOtpOnSignup"] = body.RequireOtpOnSignup
	}
	if body.RequireOtpOnSignin.IsSet() {
		changes["RequireOtpOnSignin"] = body.RequireOtpOnSignin
	}
	if body.Recaptcha2ServerKey.IsSet() {
		changes["Recaptcha2ServerKey"] = body.Recaptcha2ServerKey
	}
	if body.Recaptcha2ClientKey.IsSet() {
		changes["Recaptcha2ClientKey"] = body.Recaptcha2ClientKey
	}
	if body.EnableTotp.IsSet() {
		changes["EnableTotp"] = body.EnableTotp
	}
	if body.ForceTotp.IsSet() {
		changes["ForceTotp"] = body.ForceTotp
	}
	if body.ForcePasswordOnPhone.IsSet() {
		changes["ForcePasswordOnPhone"] = body.ForcePasswordOnPhone
	}
	if body.ForcePersonNameOnPhone.IsSet() {
		changes["ForcePersonNameOnPhone"] = body.ForcePersonNameOnPhone
	}
	if body.GeneralEmailProviderId.IsSet() {
		changes["GeneralEmailProviderId"] = body.GeneralEmailProviderId
	}
	if body.GeneralGsmProviderId.IsSet() {
		changes["GeneralGsmProviderId"] = body.GeneralGsmProviderId
	}
	if body.InviteToWorkspaceContentId.IsSet() {
		changes["InviteToWorkspaceContentId"] = body.InviteToWorkspaceContentId
	}
	if body.EmailOtpContentId.IsSet() {
		changes["EmailOtpContentId"] = body.EmailOtpContentId
	}
	if body.SmsOtpContentId.IsSet() {
		changes["SmsOtpContentId"] = body.SmsOtpContentId
	}
	return changes
}

// workspaceConfigUpsertByWorkspace is the shared upsert-by-workspace logic (distinctBy:
// workspace in the old yaml) behind WorkspaceConfigUpdateAction - Query.WorkspaceId
// (resolved from ResolveStrategy: workspace) is the condition for finding-or-creating the
// row, not a uniqueId path param.
func workspaceConfigUpsertByWorkspace(query fireback.QueryDSL, body abacdefs.WorkspaceConfigOptionalDto) (*abacdefs.WorkspaceConfigEntity, *fireback.IError) {
	dbref := fireback.GetDbRef()
	var item abacdefs.WorkspaceConfigEntity
	if err2 := dbref.
		Where(&abacdefs.WorkspaceConfigEntity{WorkspaceId: emigo.NullableOf(query.WorkspaceId)}).
		FirstOrCreate(&item).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	changes := workspaceConfigChangesFromOptionalDto(body)
	if len(changes) > 0 {
		if err2 := dbref.Model(&item).Updates(changes).Error; err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}
	}

	var updated abacdefs.WorkspaceConfigEntity
	if err2 := dbref.Where("id = ?", item.Id).First(&updated).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &updated, nil
}

func WorkspaceConfigUpdateAction(c abacdefs.WorkspaceConfigUpdateActionRequest) (*abacdefs.WorkspaceConfigUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_UPDATE))
	if err != nil {
		return nil, err
	}
	updated, err2 := workspaceConfigUpsertByWorkspace(*query, c.Body)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.WorkspaceConfigUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

// WorkspaceConfigDistinctGetAction restores the old Module3-generated GET
// /workspace-config/distinct route. Unlike every other WorkspaceConfig action, this one
// isn't scoped to whatever workspace the caller currently happens to be acting in - it
// always configures the root workspace specifically, so it deliberately bypasses
// WorkspaceConfigActions.GetByWorkspace (which filters by the caller's own
// query.WorkspaceId) and queries workspace_id = 'root' directly via plain gorm instead.
// No uniqueId is involved either way - the row is found purely by that filter. Returns
// an empty abacdefs.WorkspaceConfigEntity (not a 404) when none exists yet, matching the old
// behavior.
func WorkspaceConfigDistinctGetAction(c abacdefs.WorkspaceConfigDistinctGetActionRequest) (*abacdefs.WorkspaceConfigDistinctGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_GET_DISTINCT_WORKSPACE)); err != nil {
		return nil, err
	}

	var item abacdefs.WorkspaceConfigEntity
	err2 := fireback.GetDbRef().Where(&abacdefs.WorkspaceConfigEntity{WorkspaceId: emigo.NullableOf(ROOT_VAR)}).First(&item).Error
	if err2 != nil {
		if errors.Is(err2, gorm.ErrRecordNotFound) {
			return &abacdefs.WorkspaceConfigDistinctGetActionResponse{Payload: fireback.GResponseSingleItem(&abacdefs.WorkspaceConfigEntity{})}, nil
		}
		return nil, fireback.GormErrorToIError(err2)
	}

	return &abacdefs.WorkspaceConfigDistinctGetActionResponse{Payload: fireback.GResponseSingleItem(&item)}, nil
}

// WorkspaceConfigDistinctUpdateAction restores the old Module3-generated PATCH
// /workspace-config/distinct route. Same root-only reasoning as
// WorkspaceConfigDistinctGetAction above - always finds-or-creates the row scoped to the
// root workspace specifically (workspace_id = 'root'), via plain gorm, regardless of the
// caller's own current workspace, and needs no uniqueId to do it.
func WorkspaceConfigDistinctUpdateAction(c abacdefs.WorkspaceConfigDistinctUpdateActionRequest) (*abacdefs.WorkspaceConfigDistinctUpdateActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_UPDATE_DISTINCT_WORKSPACE)); err != nil {
		return nil, err
	}

	db := fireback.GetDbRef()
	var item abacdefs.WorkspaceConfigEntity
	if err2 := db.
		Where(&abacdefs.WorkspaceConfigEntity{WorkspaceId: emigo.NullableOf(ROOT_VAR)}).
		FirstOrCreate(&item).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	changes := workspaceConfigChangesFromOptionalDto(c.Body)
	if len(changes) > 0 {
		if err2 := db.Model(&item).Updates(changes).Error; err2 != nil {
			return nil, fireback.GormErrorToIError(err2)
		}
	}

	var updated abacdefs.WorkspaceConfigEntity
	if err2 := db.Where("id = ?", item.Id).First(&updated).Error; err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	return &abacdefs.WorkspaceConfigDistinctUpdateActionResponse{Payload: fireback.GResponseSingleItem(&updated)}, nil
}

func WorkspaceConfigAwareDeletePreviewAction(c abacdefs.WorkspaceConfigAwareDeletePreviewActionRequest) (*abacdefs.WorkspaceConfigAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_DELETE)); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.WorkspaceConfigAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.WorkspaceConfigEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.WorkspaceConfigAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WorkspaceConfigAwareDeleteAction(c abacdefs.WorkspaceConfigAwareDeleteActionRequest) (*abacdefs.WorkspaceConfigAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceConfigSecurity(PERM_ROOT_WORKSPACE_CONFIG_DELETE)); err != nil {
		return nil, err
	}
	if err2 := abacdefs.WorkspaceConfigEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.WorkspaceConfigAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

// workspaceType had permRewrite root.modules -> root.manage, and
// security: { writeOnRoot: true, readOnRoot: true } in the old yaml: the old generated
// code set AllowOnRoot: true on every action (Query/Get/Create/Update/Delete) -
// preserved here exactly. features.mock/features.msync were both explicitly false in
// the old yaml, so no mock/msync support is needed here.
var workspaceTypePerms = NewCrudPermissionSet("root.manage", "workspace-type", "workspace type")
var PERM_ROOT_WORKSPACE_TYPE = workspaceTypePerms.Wildcard
var PERM_ROOT_WORKSPACE_TYPE_QUERY = workspaceTypePerms.Query
var PERM_ROOT_WORKSPACE_TYPE_CREATE = workspaceTypePerms.Create
var PERM_ROOT_WORKSPACE_TYPE_UPDATE = workspaceTypePerms.Update
var PERM_ROOT_WORKSPACE_TYPE_DELETE = workspaceTypePerms.Delete
var ALL_WORKSPACE_TYPE_PERMISSIONS = workspaceTypePerms.All

// WorkspaceTypeMessages mirrors the entity-scoped messages: block workspaceType had in
// the old yaml - there's no Emi equivalent for entity-scoped messages, so it's
// hand-declared here.
var WorkspaceTypeMessages = struct {
	RoleIsNecessary             fireback.ErrorItem
	RoleIsNotAccessible         fireback.ErrorItem
	OnlyRootRoleIsAccepted      fireback.ErrorItem
	RoleNeedsToHaveCapabilities fireback.ErrorItem
	CannotCreateWorkspaceType   fireback.ErrorItem
	CannotModifyWorkspaceType   fireback.ErrorItem
}{
	RoleIsNecessary: fireback.ErrorItem{
		"$": "RoleIsNecessary", "en": "Role needs to be defined and exist.",
	},
	RoleIsNotAccessible: fireback.ErrorItem{
		"$": "RoleIsNotAccessible", "en": "Role is not accessible unfortunately. Make sure you the role chose exists.",
	},
	OnlyRootRoleIsAccepted: fireback.ErrorItem{
		"$": "OnlyRootRoleIsAccepted", "en": "You can only select a role which is created or belong to 'root' workspace.",
	},
	RoleNeedsToHaveCapabilities: fireback.ErrorItem{
		"$": "RoleNeedsToHaveCapabilities", "en": "Role needs to have at least one capability before could be assigned.",
	},
	CannotCreateWorkspaceType: fireback.ErrorItem{
		"$": "CannotCreateWorkspaceType", "en": "You cannot create workspace type due to some validation errors.",
	},
	CannotModifyWorkspaceType: fireback.ErrorItem{
		"$": "CannotModifyWorkspaceType", "en": "You cannot modify workspace type due to some validation errors.",
	},
}

var WorkspaceTypeActions = NewEntityActionsBundle[WorkspaceTypeEntity]()

func workspaceTypeSecurity(perm fireback.PermissionInfo) *fireback.SecurityModel {
	return &fireback.SecurityModel{ActionRequires: []fireback.PermissionInfo{perm}, AllowOnRoot: true}
}

func WorkspaceTypeBrowseAction(c WorkspaceTypeBrowseActionRequest) (*WorkspaceTypeBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_QUERY))
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := WorkspaceTypeActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceTypeBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func WorkspaceTypeGetAction(c WorkspaceTypeGetActionRequest) (*WorkspaceTypeGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_QUERY))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := WorkspaceTypeActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceTypeGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func WorkspaceTypeCreateAction(c WorkspaceTypeCreateActionRequest) (*WorkspaceTypeCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_CREATE))
	if err != nil {
		return nil, err
	}
	entity := &WorkspaceTypeEntity{
		Title:       c.Body.Title,
		Description: c.Body.Description,
		Slug:        c.Body.Slug,
		RoleId:      c.Body.RoleId,
	}
	created, err2 := WorkspaceTypeActionCreate(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceTypeCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func WorkspaceTypeUpdateAction(c WorkspaceTypeUpdateActionRequest) (*WorkspaceTypeUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_UPDATE))
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &WorkspaceTypeEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Title.Get(); ok {
		fields.Title = *v
	}
	if v, ok := c.Body.Description.Get(); ok {
		fields.Description = *v
	}
	if v, ok := c.Body.Slug.Get(); ok {
		fields.Slug = *v
	}
	if v, ok := c.Body.RoleId.Get(); ok {
		fields.RoleId = *v
	}
	updated, err2 := WorkspaceTypeActionUpdate(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &WorkspaceTypeUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func WorkspaceTypeAwareDeletePreviewAction(c WorkspaceTypeAwareDeletePreviewActionRequest) (*WorkspaceTypeAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_DELETE)); err != nil {
		return nil, err
	}
	uniqueIds := WorkspaceTypeAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := WorkspaceTypeEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceTypeAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func WorkspaceTypeAwareDeleteAction(c WorkspaceTypeAwareDeleteActionRequest) (*WorkspaceTypeAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, workspaceTypeSecurity(PERM_ROOT_WORKSPACE_TYPE_DELETE)); err != nil {
		return nil, err
	}
	if err2 := WorkspaceTypeEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &WorkspaceTypeAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// --- Hand business logic recovered from the pre-migration WorkspaceTypeEntity.go ---

func WorkspaceTypeActionCreate(
	dto *WorkspaceTypeEntity, query fireback.QueryDSL,
) (*WorkspaceTypeEntity, *fireback.IError) {

	if errors := ValidateTheWorkspaceTypeEntity(dto); len(errors) > 0 {
		return nil, &fireback.IError{
			Message:  WorkspaceTypeMessages.CannotCreateWorkspaceType,
			HttpCode: 400,
			Errors:   errors,
		}
	}

	created, err := WorkspaceTypeActions.Create(dto, query)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func WorkspaceTypeActionUpdate(
	query fireback.QueryDSL,
	fields *WorkspaceTypeEntity,
) (*WorkspaceTypeEntity, *fireback.IError) {

	if errors := ValidateTheWorkspaceTypeEntity(fields); len(errors) > 0 {
		return nil, &fireback.IError{
			Message:  WorkspaceTypeMessages.CannotModifyWorkspaceType,
			HttpCode: 400,
			Errors:   errors,
		}
	}

	return WorkspaceTypeActions.Update(query, fields)
}

func ValidateRoleAndItsExistence(roleId emigo.Nullable[string]) (*RoleEntity, []*fireback.IErrorItem) {
	items := []*fireback.IErrorItem{}

	if !roleId.IsNull() {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.RoleIsNecessary,
		})

		return nil, items
	}

	if role, err := RoleActions.GetOne(fireback.QueryDSL{UniqueId: roleId.OrDefault("")}); err != nil {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.RoleIsNotAccessible,
		})
		return nil, items
	} else {
		if role == nil {
			items = append(items, &fireback.IErrorItem{
				Location: "roleId",
				Message:  &WorkspaceTypeMessages.RoleIsNotAccessible,
			})

			return nil, items
		} else {
			if len(RoleCapabilitiesListIdGet(role)) == 0 {
				items = append(items, &fireback.IErrorItem{
					Location: "roleId",
					Message:  &WorkspaceTypeMessages.RoleNeedsToHaveCapabilities,
				})
				return nil, items
			}

			return role, nil
		}
	}
}

// Before write or update we need some extra validation for this.
// It's important to check if the role actually exists, and has some previliges
// before making it available
func ValidateTheWorkspaceTypeEntity(fields *WorkspaceTypeEntity) []*fireback.IErrorItem {
	items := []*fireback.IErrorItem{}
	role, roleErrors := ValidateRoleAndItsExistence(emigo.NullableOf(fields.RoleId))
	if len(roleErrors) != 0 {
		return roleErrors
	}

	if !role.WorkspaceId.IsSet() || role.WorkspaceId.OrDefault("") != ROOT_VAR {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.OnlyRootRoleIsAccepted,
		})

		return items
	}

	return items
}

func WorkspaceTypeActionPublicQuery(query fireback.QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionRes, *fireback.QueryResultMeta, *fireback.IError) {
	// Make this API public, so the signup screen can get it.
	// At this moment, we just move things back as are, but maybe later we need
	// to add some limits on what kind of information is going out.
	query.WorkspaceId = "root"
	query.UserId = "root"

	items, qr, err := WorkspaceTypeActions.Query(query)

	var all []*QueryWorkspaceTypesPubliclyActionRes

	for _, item := range items {
		if item.UniqueId == "root" {
			continue
		}

		all = append(all, &QueryWorkspaceTypesPubliclyActionRes{
			Title:       item.Title,
			Description: item.Description,
			UniqueId:    item.UniqueId,
			Slug:        item.Slug,
		})
	}

	return all, qr, err
}

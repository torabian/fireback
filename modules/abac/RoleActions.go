package abac

import (
	"strings"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// role had no security: block in the old yaml, so its old generated code used the
// default SecurityModel on every action (no AllowOnRoot) - preserved here exactly.
var rolePerms = NewCrudPermissionSet("root.modules", "role", "role")
var PERM_ROOT_ROLE = rolePerms.Wildcard
var PERM_ROOT_ROLE_QUERY = rolePerms.Query
var PERM_ROOT_ROLE_CREATE = rolePerms.Create
var PERM_ROOT_ROLE_UPDATE = rolePerms.Update
var PERM_ROOT_ROLE_DELETE = rolePerms.Delete
var ALL_ROLE_PERMISSIONS = rolePerms.All

// RoleMessages mirrors the entity-scoped messages: block role had in the old yaml -
// there's no Emi equivalent for entity-scoped messages, so it's hand-declared here.
var RoleMessages = struct {
	RoleNeedsOneCapability fireback.ErrorItem
	RoleNameReserved       fireback.ErrorItem
}{
	RoleNeedsOneCapability: fireback.ErrorItem{
		"$":  "RoleNeedsOneCapability",
		"en": "Role atleast needs one capability to be selected.",
	},
	RoleNameReserved: fireback.ErrorItem{
		"$":  "RoleNameReserved",
		"en": "\"root\" is a reserved role name and uniqueId - it cannot be used for a new role.",
	},
}

// RoleCapabilitiesListIdOf builds the complexes.JSON value for CapabilitiesListId from a
// plain []string of capability completeKeys - the JSON complex has no literal syntax of
// its own, so this is the constructor call sites should use.
func RoleCapabilitiesListIdOf(ids []string) complexes.JSON {
	return *complexes.JSONFrom(ids)
}

// RoleCapabilitiesListIdGet reads CapabilitiesListId back out as a plain []string.
func RoleCapabilitiesListIdGet(role *RoleEntity) []string {
	if role == nil {
		return nil
	}
	ids, _ := complexes.JSONTo[[]string](role.CapabilitiesListId)
	return ids
}

var RoleActions = NewEntityActionsBundle[RoleEntity]()

func init() {
	// filterRolePermissions/the RoleNeedsOneCapability guard deliberately do NOT wrap
	// RoleActions.Create/.Update here (unlike most other entities' overrides) - they're
	// applied only inside RoleCreateAction/RoleUpdateAction below, the actual API/CLI
	// entry points. That matches the pre-migration behavior exactly: the old
	// filterPermissions only ever mutated the transient, non-persisted
	// CapabilitiesListId convenience field (gorm:"-" sql:"-" in the old schema) used by
	// user-facing role create/update requests - it never touched the real persisted
	// grant, the Capabilities many-to-many relation, which internal/bootstrap callers
	// (WorkspaceCoreFeatures.go's signup flow, UserCli.go's SyncWorkspaceDefaultRoles,
	// CreateRootRoleInWorkspace) set directly and were therefore never filtered.
	// CapabilitiesListId is now the one and only, actually-persisted field, so wrapping
	// RoleActions.Create/.Update directly (as originally done here) would apply the
	// filter to those internal/bootstrap callers too - which have no acting-user
	// permission context to filter against, so every capability gets stripped and
	// RoleNeedsOneCapability fires on every signup. Filtering only at the action layer,
	// where ResolveActionContext has actually populated the caller's real permissions,
	// restores the old exemption for internal callers.

	// The root role (uniqueId "root") is excluded from removal whenever InternalQuery is
	// already set - preserved verbatim from the pre-migration RoleEntity.go hand file
	// (note: it does NOT add the exclusion when InternalQuery starts out empty).
	baseRemoveEnqueue := RoleActions.RemoveEnqueue
	RoleActions.RemoveEnqueue = func(request fireback.DeleteRequest, query fireback.QueryDSL) (*fireback.DeleteResponse, *fireback.IError) {
		if query.InternalQuery != "" {
			query.InternalQuery += " and unique_id != 'root'"
		}
		return baseRemoveEnqueue(request, query)
	}

	// The root role is only ever visible when querying from the root workspace itself -
	// hidden entirely from every other workspace's role list/lookup, not merely marked
	// read-only, since it's the super-admin role and has no business appearing in a
	// non-root workspace's UI at all. When it IS visible (root workspace), it's reported
	// as non-deletable/non-updatable in the response.
	baseQuery := RoleActions.Query
	RoleActions.Query = func(query fireback.QueryDSL) ([]*RoleEntity, *fireback.QueryResultMeta, *fireback.IError) {
		roles, qrm, err := baseQuery(query)
		if len(roles) > 0 {
			filtered := roles[:0]
			for _, role := range roles {
				if role.UniqueId == ROOT_VAR {
					if query.WorkspaceId != ROOT_VAR {
						if qrm != nil {
							qrm.TotalItems--
							qrm.TotalAvailableItems--
						}
						continue
					}
					f := false
					role.IsDeletable = emigo.NullableOf(f)
					role.IsUpdatable = emigo.NullableOf(f)
				}
				filtered = append(filtered, role)
			}
			roles = filtered
		}
		return roles, qrm, err
	}

	// Same visibility rule as RoleActions.Query above, for the single-item lookup path
	// (RoleGetAction) - without this, the root role would still be directly fetchable by
	// uniqueId from a non-root workspace even though it never shows up in that
	// workspace's role list. Reported as a plain 404 (not a permission error), since the
	// point is that it's not visible at all from here, not that it exists but is
	// forbidden.
	baseGetOne := RoleActions.GetOne
	RoleActions.GetOne = func(query fireback.QueryDSL) (*RoleEntity, *fireback.IError) {
		role, err := baseGetOne(query)
		if err != nil {
			return nil, err
		}
		if role != nil && role.UniqueId == ROOT_VAR && query.WorkspaceId != ROOT_VAR {
			return nil, &fireback.IError{
				Message:  fireback.FirebackMessages.ResourceNotFound,
				HttpCode: 404,
			}
		}
		return role, nil
	}
}

func filterRolePermissions(dto *RoleEntity, query fireback.QueryDSL) {
	workspaceAccesses, rolesPermission := GetWorkspaceAndUserAccesses(query)

	ids := RoleCapabilitiesListIdGet(dto)

	// Let's filter out the permissions that user actually doesn't have
	itemsFiltered := []string{}
	for _, capability := range ids {
		meetsUser := MeetsCheck([]application.PermissionInfo{{CompleteKey: capability}}, rolesPermission)
		meetsWorkspace := MeetsCheck([]application.PermissionInfo{{CompleteKey: capability}}, workspaceAccesses)

		if !meetsUser || !meetsWorkspace {
			continue
		}

		itemsFiltered = append(itemsFiltered, capability)
	}

	dto.CapabilitiesListId = RoleCapabilitiesListIdOf(itemsFiltered)
}

func RoleBrowseAction(c RoleBrowseActionRequest) (*RoleBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_ROLE_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := RoleActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &RoleBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func RoleGetAction(c RoleGetActionRequest) (*RoleGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_ROLE_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := RoleActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &RoleGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func RoleCreateAction(c RoleCreateActionRequest) (*RoleCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_ROLE_CREATE}})
	if err != nil {
		return nil, err
	}

	// "root" is the one seeded, protected super-admin role (see RoleActions.Query's
	// visibility filter and RemoveEnqueue's delete guard below) - a caller can never
	// create another role that claims that name or uniqueId, regardless of workspace,
	// case-insensitively on the name (an "Root"/"ROOT" role would be just as confusing).
	if strings.EqualFold(c.Body.Name, ROOT_VAR) {
		return nil, fireback.Create401Error(&RoleMessages.RoleNameReserved, []string{})
	}
	if v, ok := c.Body.UniqueId.Get(); ok && strings.EqualFold(*v, ROOT_VAR) {
		return nil, fireback.Create401Error(&RoleMessages.RoleNameReserved, []string{})
	}

	entity := &RoleEntity{
		Name:               c.Body.Name,
		CapabilitiesListId: c.Body.CapabilitiesListId,
		WorkspaceId:        emigo.NullableOf(query.WorkspaceId),
	}
	filterRolePermissions(entity, *query)
	if len(RoleCapabilitiesListIdGet(entity)) == 0 {
		return nil, fireback.Create401Error(&RoleMessages.RoleNeedsOneCapability, []string{})
	}
	created, err2 := RoleActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &RoleCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func RoleUpdateAction(c RoleUpdateActionRequest) (*RoleUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_ROLE_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &RoleEntity{UniqueId: c.Params.UniqueId, CapabilitiesListId: c.Body.CapabilitiesListId}
	if v, ok := c.Body.Name.Get(); ok {
		fields.Name = *v
	}
	filterRolePermissions(fields, *query)
	if len(RoleCapabilitiesListIdGet(fields)) == 0 {
		return nil, fireback.Create401Error(&RoleMessages.RoleNeedsOneCapability, []string{})
	}
	updated, err2 := RoleActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &RoleUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func RoleAwareDeletePreviewAction(c RoleAwareDeletePreviewActionRequest) (*RoleAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_ROLE_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := RoleAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := RoleEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &RoleAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func RoleAwareDeleteAction(c RoleAwareDeleteActionRequest) (*RoleAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_ROLE_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := RoleEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &RoleAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

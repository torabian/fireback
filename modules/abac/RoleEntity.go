package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	RoleActions.Create = func(dto *RoleEntity, query fireback.QueryDSL) (*RoleEntity, *fireback.IError) {
		filterPermissions(dto, query)

		if len(dto.CapabilitiesListId) == 0 && len(dto.Capabilities) == 0 {
			return nil, fireback.Create401Error(&RoleMessages.RoleNeedsOneCapability, []string{})
		}

		return RoleActionCreateFn(dto, query)
	}

	RoleActions.Update = func(query fireback.QueryDSL, dto *RoleEntity) (*RoleEntity, *fireback.IError) {
		filterPermissions(dto, query)

		if len(dto.CapabilitiesListId) == 0 && len(dto.Capabilities) == 0 {
			return nil, fireback.Create401Error(&RoleMessages.RoleNeedsOneCapability, []string{})
		}

		return RoleActionUpdateFn(query, dto)
	}

	RoleActions.RemoveEnqueue = func(request fireback.DeleteRequest, query fireback.QueryDSL) (*fireback.DeleteResponse, *fireback.IError) {
		if query.InternalQuery != "" {
			query.InternalQuery += " and unique_id != 'root'"
		}

		return RoleActionRemoveEnqueueFn(request, query)
	}

	RoleActions.Query = func(query fireback.QueryDSL) ([]*RoleEntity, *fireback.QueryResultMeta, *fireback.IError) {
		roles, qrm, err := RoleActionQueryFn(query)

		if len(roles) > 0 {
			for _, role := range roles {
				if role.UniqueId == ROOT_VAR {
					f := false
					role.IsDeletable = &f
					role.IsUpdatable = &f
				}
			}
		}

		return roles, qrm, err
	}

}

func filterPermissions(dto *RoleEntity, query fireback.QueryDSL) {
	workspaceAccesses, rolesPermission := fireback.GetWorkspaceAndUserAccesses(query)

	// Let's filter out the permissions that user actually doesn't have
	itemsFiltered := []string{}

	for _, capability := range dto.CapabilitiesListId {
		meetsUser := fireback.MeetsCheck([]fireback.PermissionInfo{{CompleteKey: capability}}, rolesPermission)
		meetsWorkspace := fireback.MeetsCheck([]fireback.PermissionInfo{{CompleteKey: capability}}, workspaceAccesses)

		if !meetsUser || !meetsWorkspace {
			continue
		}

		itemsFiltered = append(itemsFiltered, capability)
	}

	dto.CapabilitiesListId = itemsFiltered
}

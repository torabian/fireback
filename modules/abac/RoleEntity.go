package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	RoleActions.Create = func(dto *RoleEntity, query workspaces.QueryDSL) (*RoleEntity, *workspaces.IError) {
		filterPermissions(dto, query)

		if len(dto.CapabilitiesListId) == 0 && len(dto.Capabilities) == 0 {
			return nil, workspaces.Create401Error(&RoleMessages.RoleNeedsOneCapability, []string{})
		}

		return RoleActionCreateFn(dto, query)
	}

	RoleActions.Update = func(query workspaces.QueryDSL, dto *RoleEntity) (*RoleEntity, *workspaces.IError) {
		filterPermissions(dto, query)

		if len(dto.CapabilitiesListId) == 0 && len(dto.Capabilities) == 0 {
			return nil, workspaces.Create401Error(&RoleMessages.RoleNeedsOneCapability, []string{})
		}

		return RoleActionUpdateFn(query, dto)
	}

}

func filterPermissions(dto *RoleEntity, query workspaces.QueryDSL) {
	workspaceAccesses, rolesPermission := workspaces.GetWorkspaceAndUserAccesses(query)

	// Let's filter out the permissions that user actually doesn't have
	itemsFiltered := []string{}

	for _, capability := range dto.CapabilitiesListId {
		meetsUser := workspaces.MeetsCheck([]workspaces.PermissionInfo{{CompleteKey: capability}}, rolesPermission)
		meetsWorkspace := workspaces.MeetsCheck([]workspaces.PermissionInfo{{CompleteKey: capability}}, workspaceAccesses)

		if !meetsUser || !meetsWorkspace {
			continue
		}

		itemsFiltered = append(itemsFiltered, capability)
	}

	dto.CapabilitiesListId = itemsFiltered
}

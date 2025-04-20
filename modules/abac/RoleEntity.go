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

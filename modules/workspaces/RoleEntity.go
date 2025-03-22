package workspaces

func init() {
	RoleActions.Create = func(dto *RoleEntity, query QueryDSL) (*RoleEntity, *IError) {
		filterPermissions(dto, query)

		if len(dto.CapabilitiesListId) == 0 && len(dto.Capabilities) == 0 {
			return nil, Create401Error(&RoleMessages.RoleNeedsOneCapability, []string{})
		}

		return RoleActionCreateFn(dto, query)
	}

	RoleActions.Update = func(query QueryDSL, dto *RoleEntity) (*RoleEntity, *IError) {
		filterPermissions(dto, query)

		if len(dto.CapabilitiesListId) == 0 && len(dto.Capabilities) == 0 {
			return nil, Create401Error(&RoleMessages.RoleNeedsOneCapability, []string{})
		}

		return RoleActionUpdateFn(query, dto)
	}

}

func filterPermissions(dto *RoleEntity, query QueryDSL) {
	workspaceAccesses, rolesPermission := GetWorkspaceAndUserAccesses(query)

	// Let's filter out the permissions that user actually doesn't have
	itemsFiltered := []string{}

	for _, capability := range dto.CapabilitiesListId {
		meetsUser := meetsCheck([]PermissionInfo{{CompleteKey: capability}}, rolesPermission)
		meetsWorkspace := meetsCheck([]PermissionInfo{{CompleteKey: capability}}, workspaceAccesses)

		if !meetsUser || !meetsWorkspace {
			continue
		}

		itemsFiltered = append(itemsFiltered, capability)
	}

	dto.CapabilitiesListId = itemsFiltered
}

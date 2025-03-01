package workspaces

func UserWorkspaceActionCreate(
	dto *UserWorkspaceEntity, query QueryDSL,
) (*UserWorkspaceEntity, *IError) {
	return UserWorkspaceActionCreateFn(dto, query)
}
func UserWorkspaceActionUpdate(
	query QueryDSL,
	fields *UserWorkspaceEntity,
) (*UserWorkspaceEntity, *IError) {
	return UserWorkspaceActionUpdateFn(query, fields)
}

func CastUserRoleWorkspacDtoMap(data map[string]*UserRoleWorkspaceDto) []UserRoleWorkspaceDto {
	items := []UserRoleWorkspaceDto{}
	for _, v := range data {
		items = append(items, *v)
	}
	return items
}

func UserWorkspacePostFormatter(dto *UserWorkspaceEntity, query QueryDSL) {

	roles := map[string]*UserRoleWorkspaceDto{}

	for _, urwItem := range query.UserRoleWorkspacePermissions {
		if urwItem.Type == "" || urwItem.CapabilityId == "" || urwItem.RoleId == "" || urwItem.WorkspaceId == "" {
			continue
		}
		if urwItem.Type == "account_restrict" && urwItem.WorkspaceId == dto.WorkspaceId.String {
			if roles[urwItem.RoleId] == nil {
				roles[urwItem.RoleId] = &UserRoleWorkspaceDto{
					RoleId:       urwItem.RoleId,
					Capabilities: []string{},
				}
			}
			roles[urwItem.RoleId].Capabilities = append(roles[urwItem.RoleId].Capabilities, urwItem.CapabilityId)
		}
	}

	for _, urwItem := range query.UserRoleWorkspacePermissions {
		if urwItem.Type == "" || urwItem.CapabilityId == "" || urwItem.RoleId == "" || urwItem.WorkspaceId == "" {
			continue
		}
		if urwItem.Type == "workspace_restrict" && urwItem.WorkspaceId == dto.WorkspaceId.String {
			dto.WorkspacePermissions = append(dto.WorkspacePermissions, urwItem.CapabilityId)
		}

		if urwItem.Type == "account_restrict" {
			dto.UserPermissions = append(dto.UserPermissions, urwItem.CapabilityId)
			dto.RolePermission = CastUserRoleWorkspacDtoMap(roles)
		}
	}

}

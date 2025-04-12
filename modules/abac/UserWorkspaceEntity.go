package abac

func CastUserRoleWorkspacDtoMap(data map[string]*UserRoleWorkspaceDto) []UserRoleWorkspaceDto {
	items := []UserRoleWorkspaceDto{}
	for _, v := range data {
		items = append(items, *v)
	}
	return items
}

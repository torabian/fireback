package workspaces

func WorkspaceTypeActionCreate(
	dto *WorkspaceTypeEntity, query QueryDSL,
) (*WorkspaceTypeEntity, *IError) {

	if errors := ValidateTheWorkspaceTypeEntity(dto); len(errors) > 0 {
		return nil, &IError{
			Message:  WorkspaceTypeMessages.CannotCreateWorkspaceType,
			HttpCode: 400,
			Errors:   errors,
		}
	}

	return WorkspaceTypeActionCreateFn(dto, query)
}

func WorkspaceTypeActionUpdate(
	query QueryDSL,
	fields *WorkspaceTypeEntity,
) (*WorkspaceTypeEntity, *IError) {

	if errors := ValidateTheWorkspaceTypeEntity(fields); len(errors) > 0 {
		return nil, &IError{
			Message:  WorkspaceTypeMessages.CannotModifyWorkspaceType,
			HttpCode: 400,
			Errors:   errors,
		}
	}

	return WorkspaceTypeActionUpdateFn(query, fields)
}

func ValidateRoleAndItsExsitence(roleId String) (*RoleEntity, []*IErrorItem) {
	items := []*IErrorItem{}

	if !roleId.Valid {
		items = append(items, &IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.RoleIsNecessary,
		})

		return nil, items
	}

	if role, err := RoleActions.GetOne(QueryDSL{UniqueId: roleId.String}); err != nil {
		items = append(items, &IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.RoleIsNotAccessible,
		})
		return nil, items
	} else {
		if role == nil {
			items = append(items, &IErrorItem{
				Location: "roleId",
				Message:  &WorkspaceTypeMessages.RoleIsNotAccessible,
			})

			return nil, items
		} else {
			if len(role.Capabilities) == 0 {
				items = append(items, &IErrorItem{
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
func ValidateTheWorkspaceTypeEntity(fields *WorkspaceTypeEntity) []*IErrorItem {
	items := []*IErrorItem{}
	role, roleErrors := ValidateRoleAndItsExsitence(fields.RoleId)
	if len(roleErrors) != 0 {
		return roleErrors
	}

	if !role.WorkspaceId.Valid || role.WorkspaceId.String != ROOT_VAR {
		items = append(items, &IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.OnlyRootRoleIsAccepted,
		})

		return items
	}

	return items
}

func WorkspaceTypeActionPublicQuery(query QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionResDto, *QueryResultMeta, error) {
	// Make this API public, so the signup screen can get it.
	// At this moment, we just move things back as are, but maybe later we need
	// to add some limits on what kind of information is going out.
	query.WorkspaceId = "root"
	query.UserId = "root"

	items, qr, err := WorkspaceTypeActions.Query(query)
	var all []*QueryWorkspaceTypesPubliclyActionResDto

	for _, item := range items {
		if item.UniqueId == "root" {
			continue
		}

		all = append(all, &QueryWorkspaceTypesPubliclyActionResDto{
			Title:       item.Title,
			Description: item.Description,
			UniqueId:    item.UniqueId,
			Slug:        item.Slug,
		})
	}

	return all, qr, err
}

func init() {
	QueryWorkspaceTypesPubliclyActionImp = func(q QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionResDto, *QueryResultMeta, *IError) {
		res, qrm, err := WorkspaceTypeActionPublicQuery(q)
		return res, qrm, CastToIError(err)
	}
}

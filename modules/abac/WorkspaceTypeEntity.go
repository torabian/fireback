package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

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

	return WorkspaceTypeActionCreateFn(dto, query)
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

	return WorkspaceTypeActionUpdateFn(query, fields)
}

func ValidateRoleAndItsExsitence(roleId fireback.String) (*RoleEntity, []*fireback.IErrorItem) {
	items := []*fireback.IErrorItem{}

	if !roleId.Valid {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.RoleIsNecessary,
		})

		return nil, items
	}

	if role, err := RoleActions.GetOne(fireback.QueryDSL{UniqueId: roleId.String}); err != nil {
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
			if len(role.Capabilities) == 0 {
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
	role, roleErrors := ValidateRoleAndItsExsitence(fields.RoleId)
	if len(roleErrors) != 0 {
		return roleErrors
	}

	if !role.WorkspaceId.Valid || role.WorkspaceId.String != ROOT_VAR {
		items = append(items, &fireback.IErrorItem{
			Location: "roleId",
			Message:  &WorkspaceTypeMessages.OnlyRootRoleIsAccepted,
		})

		return items
	}

	return items
}

func WorkspaceTypeActionPublicQuery(query fireback.QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionResDto, *fireback.QueryResultMeta, *fireback.IError) {
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
	QueryWorkspaceTypesPubliclyActionImp = func(q fireback.QueryDSL) ([]*QueryWorkspaceTypesPubliclyActionResDto, *fireback.QueryResultMeta, *fireback.IError) {
		res, qrm, err := WorkspaceTypeActionPublicQuery(q)
		return res, qrm, fireback.CastToIError(err)
	}
}

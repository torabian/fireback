package abac

import (
	"errors"
	"strings"

	queries "github.com/torabian/fireback/modules/abac/queries"
	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
)

type GenerateUserDto struct {
	// The user we want to create, also can include the person object
	// for personal information it can create that.
	user *UserEntity

	createUser bool

	// Workspace that this user will be assigend to
	workspace *WorkspaceEntity

	createWorkspace bool

	// The roles that this user will have
	role *RoleEntity

	createRole bool

	// Restrict means, if any operation failed, we rolle back
	// In some scenarios, some entities might existing
	restricted bool

	createPassport bool
	passport       *PassportEntity
}

/**
*	A general function to create a user, and generate a session with that
*	It's necessary to know you need to do some initial tests before using this function
*	and it would never be exported to public access directly
 */

func CreateWorkspaceAndAssignUser(dto *GenerateUserDto, q fireback.QueryDSL, session *UserSessionDto) *fireback.IError {
	workspaceId := dto.workspace.UniqueId
	q.WorkspaceId = workspaceId

	q.UserId = dto.user.UniqueId
	dto.workspace.WorkspaceId = fireback.NewString(workspaceId)

	var actualWorkspace *WorkspaceEntity = nil
	if existingWs, err9 := WorkspaceActions.GetOne(fireback.QueryDSL{UniqueId: dto.workspace.UniqueId, Tx: q.Tx}); err9 == nil && existingWs != nil {
		if existingWs.UniqueId == dto.workspace.UniqueId {
			actualWorkspace = existingWs
		}
	}

	if actualWorkspace == nil {

		if ws, err := WorkspaceActions.Create(dto.workspace, q); err != nil {
			if dto.restricted {
				return err
			}
		} else {
			actualWorkspace = ws
		}
	}

	workspaceId = actualWorkspace.UniqueId
	q.WorkspaceId = actualWorkspace.UniqueId

	var userWorkspace *UserWorkspaceEntity
	// This is a bit special table, I did not want introduce a new concept
	// In fireback, so it would be like this to modify things directly.

	// let's find that link, if not exists create it.
	if errFinding := q.Tx.Model(&UserWorkspaceEntity{}).Where(&UserWorkspaceEntity{
		WorkspaceId: fireback.NewString(workspaceId),
		UserId:      fireback.NewString(q.UserId),
	}).Find(&userWorkspace); errFinding != nil {
		if createdWorkspace, err := UserWorkspaceActions.Create(&UserWorkspaceEntity{
			WorkspaceId: fireback.NewString(workspaceId),
			UserId:      fireback.NewString(q.UserId),
		}, q); err == nil {
			userWorkspace = createdWorkspace
		}
	}
	if userWorkspace != nil {
		session.UserWorkspaces = []*UserWorkspaceEntity{userWorkspace}
	}

	return nil
}

func runTransaction[T any](
	entity *T, query fireback.QueryDSL,
	fn func(tx *gorm.DB) error,
) (*T, *fireback.IError) {

	vf := fireback.GetRef(query).Transaction(fn)

	if vf != nil {
		return nil, fireback.CastToIError(vf)
	}
	return entity, nil
}

// This is core function of creating a new user in the system.
// All passport methods, need to pass through this logic in order to
// create account publicly.
func UnsafeGenerateUser(dto *GenerateUserDto, q fireback.QueryDSL) (*UserSessionDto, *fireback.IError) {
	session := &UserSessionDto{}

	return runTransaction(session, q, func(tx *gorm.DB) error {
		q.Tx = tx

		if dto.createPassport && dto.passport != nil {
			dto.passport.UserId = fireback.NewString(dto.user.UniqueId)

			// Passport and user always belong to the root workspace
			dto.passport.WorkspaceId = fireback.NewString(ROOT_VAR)
			q.WorkspaceId = ROOT_VAR
			q.UserId = dto.user.UniqueId
			if passportdb, err := PassportActions.Create(dto.passport, q); err != nil {
				if dto.restricted {
					return err
				}
			} else {
				session.Passport = passportdb
				// dto.user.PassportsListId = []string{passportdb.UniqueId}
			}
		}

		if dto.createUser && dto.user != nil {
			q.UserId = dto.user.UniqueId
			if _, err := UserActionCreate(dto.user, q); err != nil {
				if dto.restricted {
					return err
				}
			}

			session.User = dto.user
		}

		if dto.createWorkspace && dto.workspace != nil {
			if err5 := CreateWorkspaceAndAssignUser(dto, q, session); err5 != nil {
				return err5
			}
		}

		if dto.createRole && dto.role != nil {

			// Make sure the q.WorkspaceId is not root anymore
			q2 := q
			q2.WorkspaceId = dto.workspace.UniqueId
			if _, err := RoleActions.Create(dto.role, q2); err != nil {
				if dto.restricted {
					return err
				}
			}

			// Note: here we skipped to add the workspace role into the session
			// this is used somewhere else
			wre := &WorkspaceRoleEntity{
				UserWorkspaceId: fireback.NewString(session.UserWorkspaces[0].UniqueId),
				RoleId:          fireback.NewString(dto.role.UniqueId),
				WorkspaceId:     fireback.NewString(dto.workspace.UniqueId),
			}

			wsid := q.WorkspaceId
			q.WorkspaceId = dto.workspace.UniqueId
			if _, err := WorkspaceRoleActions.Create(wre, q); err != nil {
				if dto.restricted {
					return err
				}
			}
			q.WorkspaceId = wsid
		}

		// For creating a user, we need at least the user to be available
		if session.User == nil {
			return errors.New("USER_IS_MISSING")
		}

		// Token for the session is essential, a session without a token
		// has absolutely no use.
		if token, err := session.User.AuthorizeWithToken(q); err != nil {
			return err
		} else {
			session.Token = token
		}

		return nil
	})

}

/**
*	Return the definition of operation, make sure it does not
*	Do any effect on the database
**/
func GetOsHostUserRoleWorkspaceDef() (*UserEntity, *RoleEntity, *WorkspaceEntity) {
	osUser := fireback.GetOsUserWithPhone()
	name := osUser.Name + "'s workspace"
	user := &UserEntity{
		UniqueId:    "OS_USER_" + osUser.Uid,
		WorkspaceId: fireback.NewString(ROOT_VAR),
		FirstName:   osUser.Username,
		LastName:    osUser.Username,
	}

	wid := "OS_WORKSPACE_" + osUser.Uid
	workspace := &WorkspaceEntity{
		Name:        name,
		UniqueId:    wid,
		WorkspaceId: fireback.NewString(wid),
		LinkerId:    fireback.NewString(ROOT_VAR),
		ParentId:    fireback.NewString(ROOT_VAR),
		TypeId:      fireback.NewString(ROOT_VAR),
	}

	osRole := "OS User"
	role := &RoleEntity{
		UniqueId:    "ROLE_WORKSPACE_" + osUser.Uid,
		Name:        osRole,
		WorkspaceId: fireback.NewString(workspace.UniqueId),
		Capabilities: []*fireback.CapabilityEntity{
			{
				UniqueId:    ROOT_ALL_MODULES,
				Visibility:  fireback.NewString("A"),
				WorkspaceId: fireback.NewString("system"),
			},
		},
	}

	return user, role, workspace
}

func DetectSignupMechanismOverValue(value string) (string, *fireback.IError) {
	if strings.Contains(value, "@") {
		return PASSPORT_METHOD_EMAIL, nil
	}
	if strings.Contains(value, "+") {
		return PASSPORT_METHOD_PHONE, nil
	}

	return "", fireback.Create401Error(&AbacMessages.PassportNotAvailable, []string{})

}

func GetEmailPassportSignupMechanism(dto *ClassicSignupActionReqDto) (*UserEntity, *RoleEntity, *WorkspaceEntity, *PassportEntity) {

	userId := fireback.UUID()
	workspaceId := fireback.UUID()
	roleId := fireback.UUID()
	passportId := fireback.UUID()

	user := &UserEntity{
		UniqueId:  userId,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	wname := "workspace"
	workspace := &WorkspaceEntity{
		UniqueId: workspaceId,
		Name:     wname,
		LinkerId: fireback.NewString(ROOT_VAR),
		ParentId: fireback.NewString(ROOT_VAR),
		TypeId:   dto.WorkspaceTypeId,
	}

	osRole := "Admin"
	role := &RoleEntity{
		UniqueId:    roleId,
		Name:        osRole,
		WorkspaceId: fireback.NewString(workspace.UniqueId),
		Capabilities: []*fireback.CapabilityEntity{
			{
				UniqueId:    ROOT_ALL_MODULES,
				Visibility:  fireback.NewString("A"),
				WorkspaceId: fireback.NewString("system"),
			},
		},
	}

	method, _ := DetectSignupMechanismOverValue(dto.Value)
	passwordHashed := ""
	if strings.TrimSpace(dto.Password) != "" {
		genPass, _ := fireback.HashPassword(dto.Password)
		passwordHashed = genPass
	}

	passport := &PassportEntity{
		Type:     method,
		Password: passwordHashed,
		Value:    dto.Value,
		UniqueId: passportId,
	}

	return user, role, workspace, passport
}

func GetUserAccessLevels(query fireback.QueryDSL) (*UserAccessLevelDto, *fireback.IError) {

	access := &UserAccessLevelDto{}
	query.ItemsPerPage = 1000

	items, _, err := fireback.UnsafeQuerySqlFromFs[UserRoleWorkspacePermissionDto](
		&queries.QueriesFs, "UserRolePermission", query,
	)

	if err != nil {
		return nil, fireback.CastToIError(err)
	}

	ws := fireback.UserAccessPerWorkspaceDto{}

	for _, item := range items {
		if ws[item.WorkspaceId] == nil {
			ws[item.WorkspaceId] = &struct {
				WorkspacesAccesses []string
				UserRoles          map[string]*struct {
					Name     string
					Accesses []string
				}
			}{}
		}

		if item.Type == "account_restrict" {
			if ws[item.WorkspaceId].UserRoles[item.RoleId] == nil {
				ws[item.WorkspaceId].UserRoles = map[string]*struct {
					Name     string
					Accesses []string
				}{}
				ws[item.WorkspaceId].UserRoles[item.RoleId] = &struct {
					Name     string
					Accesses []string
				}{}
			}
			ws[item.WorkspaceId].UserRoles[item.RoleId].Accesses = append(ws[item.WorkspaceId].UserRoles[item.RoleId].Accesses, item.CapabilityId)
			ws[item.WorkspaceId].UserRoles[item.RoleId].Name = item.RoleName
		}

		if item.Type == "workspace_restrict" {
			ws[item.WorkspaceId].WorkspacesAccesses = append(ws[item.WorkspaceId].WorkspacesAccesses, item.CapabilityId)
		}
	}

	access.UserAccessPerWorkspace = &ws

	return access, nil
}

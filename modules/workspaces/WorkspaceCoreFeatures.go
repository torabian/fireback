package workspaces

import (
	"encoding/json"
	"errors"
	"strings"

	queries "github.com/torabian/fireback/modules/workspaces/queries"
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

func CreateWorkspaceAndAssignUser(dto *GenerateUserDto, q QueryDSL, session *UserSessionDto) *IError {
	workspaceId := dto.workspace.UniqueId
	q.WorkspaceId = workspaceId

	q.UserId = dto.user.UniqueId
	dto.workspace.WorkspaceId = NewString(workspaceId)
	var actualWorkspace *WorkspaceEntity = nil
	if ws, err := WorkspaceActions.Create(dto.workspace, q); err != nil {
		if dto.restricted {
			return err
		}
	} else {
		actualWorkspace = ws

		workspaceId = actualWorkspace.UniqueId
		q.WorkspaceId = actualWorkspace.UniqueId
	}

	// This is a bit special table, I did not want introduce a new concept
	// In fireback, so it would be like this to modify things directly.
	if userWorkspace, err := UserWorkspaceActions.Create(&UserWorkspaceEntity{
		WorkspaceId: NewString(workspaceId),
		UserId:      NewString(q.UserId),
	}, q); err != nil {
		if dto.restricted {
			return err
		} else {
			session.UserWorkspaces = []*UserWorkspaceEntity{userWorkspace}
		}
	} else {
		session.UserWorkspaces = []*UserWorkspaceEntity{userWorkspace}
	}

	return nil
}

func runTransaction[T any](
	entity *T, query QueryDSL,
	fn func(tx *gorm.DB) error,
) (*T, *IError) {

	vf := GetRef(query).Transaction(fn)

	if vf != nil {
		return nil, CastToIError(vf)
	}
	return entity, nil
}

// This is core function of creating a new user in the system.
// All passport methods, need to pass through this logic in order to
// create account publicly.
func UnsafeGenerateUser(dto *GenerateUserDto, q QueryDSL) (*UserSessionDto, *IError) {
	session := &UserSessionDto{}

	return runTransaction(session, q, func(tx *gorm.DB) error {
		q.Tx = tx

		if dto.createPassport && dto.passport != nil {
			dto.passport.UserId = NewString(dto.user.UniqueId)

			// Passport and user always belong to the root workspace
			dto.passport.WorkspaceId = NewString(ROOT_VAR)
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
				UserWorkspaceId: NewString(session.UserWorkspaces[0].UniqueId),
				RoleId:          NewString(dto.role.UniqueId),
				WorkspaceId:     NewString(dto.workspace.UniqueId),
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
	osUser := GetOsUserWithPhone()
	name := osUser.Name + "'s workspace"
	user := &UserEntity{
		UniqueId:    "OS_USER_" + osUser.Uid,
		WorkspaceId: NewString(ROOT_VAR),
		FirstName:   osUser.Username,
		LastName:    osUser.Username,
	}

	wid := "OS_WORKSPACE_" + osUser.Uid
	workspace := &WorkspaceEntity{
		Name:        name,
		UniqueId:    wid,
		WorkspaceId: NewString(wid),
		LinkerId:    NewString(ROOT_VAR),
		ParentId:    NewString(ROOT_VAR),
		TypeId:      NewString(ROOT_VAR),
	}

	osRole := "OS User"
	role := &RoleEntity{
		UniqueId:    "ROLE_WORKSPACE_" + osUser.Uid,
		Name:        osRole,
		WorkspaceId: NewString(workspace.UniqueId),
		Capabilities: []*CapabilityEntity{
			{
				UniqueId:    ROOT_ALL_MODULES,
				Visibility:  NewString("A"),
				WorkspaceId: NewString("system"),
			},
		},
	}

	return user, role, workspace
}

func DetectSignupMechanismOverValue(value string) (string, *IError) {
	if strings.Contains(value, "@") {
		return PASSPORT_METHOD_EMAIL, nil
	}
	if strings.Contains(value, "+") {
		return PASSPORT_METHOD_PHONE, nil
	}

	return "", Create401Error(&WorkspacesMessages.PassportNotAvailable, []string{})

}

func GetEmailPassportSignupMechanism(dto *ClassicSignupActionReqDto) (*UserEntity, *RoleEntity, *WorkspaceEntity, *PassportEntity) {

	userId := UUID()
	workspaceId := UUID()
	roleId := UUID()
	passportId := UUID()

	user := &UserEntity{
		UniqueId:  userId,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	wname := "workspace"
	workspace := &WorkspaceEntity{
		UniqueId: workspaceId,
		Name:     wname,
		LinkerId: NewString(ROOT_VAR),
		ParentId: NewString(ROOT_VAR),
		TypeId:   dto.WorkspaceTypeId,
	}

	osRole := "Admin"
	role := &RoleEntity{
		UniqueId:    roleId,
		Name:        osRole,
		WorkspaceId: NewString(workspace.UniqueId),
		Capabilities: []*CapabilityEntity{
			{
				UniqueId:    ROOT_ALL_MODULES,
				Visibility:  NewString("A"),
				WorkspaceId: NewString("system"),
			},
		},
	}

	method, _ := DetectSignupMechanismOverValue(dto.Value)
	passwordHashed := ""
	if strings.TrimSpace(dto.Password) != "" {
		genPass, _ := HashPassword(dto.Password)
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

type UserAccessPerWorkspaceDto map[string]*struct {

	// The access which are available to this workspace, not to the specific user.
	// Even a user has access to many things, these accesses need to reduce those
	WorkspacesAccesses []string

	// The permissions which user has access to
	UserRoles map[string]*struct {
		Name     string
		Accesses []string
	}
}

func (x *UserAccessPerWorkspaceDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

func GetUserAccessLevels(query QueryDSL) (*UserAccessLevelDto, *IError) {

	access := &UserAccessLevelDto{}
	query.ItemsPerPage = 1000

	items, _, err := UnsafeQuerySqlFromFs[UserRoleWorkspacePermissionDto](
		&queries.QueriesFs, "UserRolePermission", query,
	)

	if err != nil {
		return nil, CastToIError(err)
	}

	ws := UserAccessPerWorkspaceDto{}

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

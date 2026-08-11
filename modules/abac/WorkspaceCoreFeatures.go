package abac

import (
	"errors"
	"strings"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	queries "github.com/torabian/fireback/modules/abac/queries"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/security"
	"gorm.io/gorm"
)

type GenerateUserDto struct {
	// The user we want to create, also can include the person object
	// for personal information it can create that.
	user *abacdefs.UserEntity

	createUser bool

	// Workspace that this user will be assigend to
	workspace *abacdefs.WorkspaceEntity

	createWorkspace bool

	// The roles that this user will have
	role *abacdefs.RoleEntity

	createRole bool

	// Restrict means, if any operation failed, we rolle back
	// In some scenarios, some entities might existing
	restricted bool

	createPassport bool
	passport       *abacdefs.PassportEntity
}

/**
*	A general function to create a user, and generate a session with that
*	It's necessary to know you need to do some initial tests before using this function
*	and it would never be exported to public access directly
 */

func CreateWorkspaceAndAssignUser(dto *GenerateUserDto, q fireback.QueryDSL, session *abacdefs.UserSessionDto) *fireback.IError {
	workspaceId := dto.workspace.UniqueId
	q.WorkspaceId = workspaceId

	q.UserId = dto.user.UniqueId
	dto.workspace.WorkspaceId = emigo.NullableOf(workspaceId)

	var actualWorkspace *abacdefs.WorkspaceEntity = nil
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
			// restricted is false: creation failed but the caller wants to proceed
			// anyway (see the doc comment above dto.restricted) - there's nothing to
			// link a user to below, so stop here rather than falling through to a nil
			// actualWorkspace dereference. Not reachable from any caller today (both
			// CreateWorkspaceAction and UnsafeGenerateUser always pass restricted:
			// true), but left defensive since dto.restricted exists specifically to
			// allow this path.
			return nil
		} else {
			actualWorkspace = ws
		}
	}

	// dto.workspace only ever carries the caller-supplied fields (e.g. just Name) -
	// actualWorkspace is the real, fully-populated row (generated UniqueId included).
	// Without writing it back here, CreateWorkspaceAction's response always reported
	// workspaceId:"" even on a successful create, since it reads back dto.workspace
	// (the same pointer it originally passed in) after this function returns.
	dto.workspace = actualWorkspace
	workspaceId = actualWorkspace.UniqueId
	q.WorkspaceId = actualWorkspace.UniqueId

	var userWorkspace *abacdefs.UserWorkspaceEntity
	// This is a bit special table, I did not want introduce a new concept
	// In fireback, so it would be like this to modify things directly.

	// let's find that link, if not exists create it. q.Tx is nil whenever this is
	// reached outside a runTransaction-wrapped caller (e.g. CreateWorkspaceAction calls
	// this directly, with no transaction of its own) - fireback.GetRef falls back to
	// the plain db ref in that case, same as everywhere else in this codebase; calling
	// q.Tx.Model(...) directly on a nil *gorm.DB panicked with a nil pointer
	// dereference on every single call.
	if errFinding := fireback.GetRef(q).Model(&abacdefs.UserWorkspaceEntity{}).Where(&abacdefs.UserWorkspaceEntity{
		WorkspaceId: emigo.NullableOf(workspaceId),
		UserId:      emigo.NullableOf(q.UserId),
	}).Find(&userWorkspace); errFinding != nil {
		if createdWorkspace, err := UserWorkspaceActions.Create(&abacdefs.UserWorkspaceEntity{
			WorkspaceId: emigo.NullableOf(workspaceId),
			UserId:      emigo.NullableOf(q.UserId),
		}, q); err == nil {
			userWorkspace = createdWorkspace
		}
	}
	if userWorkspace != nil {
		session.UserWorkspaces = emigo.CollectionReplace([]abacdefs.UserWorkspaceEntity{*userWorkspace})
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
func UnsafeGenerateUser(dto *GenerateUserDto, q fireback.QueryDSL) (*abacdefs.UserSessionDto, *fireback.IError) {
	session := &abacdefs.UserSessionDto{}

	return runTransaction(session, q, func(tx *gorm.DB) error {
		q.Tx = tx

		if dto.createPassport && dto.passport != nil {
			dto.passport.UserId = emigo.NullableOf(dto.user.UniqueId)

			// Passport and user always belong to the root workspace
			dto.passport.WorkspaceId = emigo.NullableOf(ROOT_VAR)
			q.WorkspaceId = ROOT_VAR
			q.UserId = dto.user.UniqueId
			if passportdb, err := PassportActions.Create(dto.passport, q); err != nil {
				if dto.restricted {
					return err
				}
			} else {
				session.Passport.Set(*passportdb)
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

			session.User.Set(*dto.user)
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
			wre := &abacdefs.WorkspaceRoleEntity{
				UserWorkspaceId: emigo.NullableOf(session.UserWorkspaces.Items[0].UniqueId),
				RoleId:          emigo.NullableOf(dto.role.UniqueId),
				WorkspaceId:     emigo.NullableOf(dto.workspace.UniqueId),
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
		user, _ := session.User.Get()
		if user == nil {
			return errors.New("USER_IS_MISSING")
		}

		// Token for the session is essential, a session without a token
		// has absolutely no use.
		if token, err := AuthorizeWithToken(&user.Item, q); err != nil {
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
func GetOsHostUserRoleWorkspaceDef() (*abacdefs.UserEntity, *abacdefs.RoleEntity, *abacdefs.WorkspaceEntity) {
	osUser := fireback.GetOsUserWithPhone()
	name := osUser.Name + "'s workspace"
	user := &abacdefs.UserEntity{
		UniqueId:  "OS_USER_" + osUser.Uid,
		FirstName: osUser.Username,
		LastName:  osUser.Username,
	}

	wid := "OS_WORKSPACE_" + osUser.Uid
	workspace := &abacdefs.WorkspaceEntity{
		Name:        name,
		UniqueId:    wid,
		WorkspaceId: emigo.NullableOf(wid),
		ParentId:    emigo.NullableOf(ROOT_VAR),
		TypeId:      ROOT_VAR,
	}

	osRole := "OS User"
	role := &abacdefs.RoleEntity{
		UniqueId:           "ROLE_WORKSPACE_" + osUser.Uid,
		Name:               osRole,
		WorkspaceId:        emigo.NullableOf(workspace.UniqueId),
		CapabilitiesListId: RoleCapabilitiesListIdOf([]string{ROOT_ALL_MODULES}),
	}

	return user, role, workspace
}

func DetectSignupMechanismOverValue(value string) (string, *fireback.IError) {
	if strings.HasPrefix(value, "anonymous_") {
		return ANONYMOUS_AUTHENTICATION, nil
	}
	if strings.Contains(value, "@") {
		return PASSPORT_METHOD_EMAIL, nil
	}
	if strings.Contains(value, "+") {
		return PASSPORT_METHOD_PHONE, nil
	}

	return "", fireback.Create401Error(&AbacMessages.PassportNotAvailable, []string{})

}

func GetEmailPassportSignupMechanism(dto *abacdefs.ClassicSignupActionReq) (*abacdefs.UserEntity, *abacdefs.RoleEntity, *abacdefs.WorkspaceEntity, *abacdefs.PassportEntity) {

	userId := fireback.UUID()
	workspaceId := fireback.UUID()
	roleId := fireback.UUID()
	passportId := fireback.UUID()

	user := &abacdefs.UserEntity{
		UniqueId:  userId,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	wname := "workspace"
	workspace := &abacdefs.WorkspaceEntity{
		UniqueId: workspaceId,
		Name:     wname,
		ParentId: emigo.NullableOf(ROOT_VAR),
		TypeId:   dto.WorkspaceTypeId.OrDefault(""),
	}

	osRole := "Admin"
	role := &abacdefs.RoleEntity{
		UniqueId:           roleId,
		Name:               osRole,
		WorkspaceId:        emigo.NullableOf(workspace.UniqueId),
		CapabilitiesListId: RoleCapabilitiesListIdOf(ResolveWorkspaceTypeRoleCapabilities(workspace.TypeId)),
	}

	method, _ := DetectSignupMechanismOverValue(dto.Value)
	passwordHashed := ""
	if strings.TrimSpace(dto.Password) != "" {
		genPass, _ := security.HashPassword(dto.Password)
		passwordHashed = genPass
	}

	passport := &abacdefs.PassportEntity{
		Type:     method,
		Password: passwordHashed,
		Value:    dto.Value,
		UniqueId: passportId,
	}

	return user, role, workspace, passport
}

// ResolveWorkspaceTypeRoleCapabilities looks up the role a workspace type is
// configured with (WorkspaceTypeEntity.RoleId - validated on create/update by
// ValidateTheWorkspaceTypeEntity in WorkspaceTypeActions.go to exist, belong to
// root, have at least one capability, and never be root/root.* itself) and
// returns that role's capabilities - what a fresh signup into that workspace
// type should actually be granted.
//
// Bug fix: GetEmailPassportSignupMechanism previously hardcoded every fresh
// signup's own role to ROOT_ALL_MODULES ("root.modules.*") regardless of which
// workspace type was selected, so different workspace types never actually
// granted different permissions to whoever signed up as them - only a
// workspace-level capability list (fed from the type's role via a separate
// path, GetWorkspaceAndUserAccesses) ever reflected the chosen type, while the
// signed-up user's own role capabilities did not. Since MeetsAccessLevel
// requires *both* the user's own role and their workspace to satisfy a given
// permission (see Permissions.go), that mismatch could make a type's own
// intended capabilities entirely unreachable whenever they fell outside
// "root.modules.*" (e.g. anything under the "root.manage.*" prefix).
//
// Falls back to ROOT_ALL_MODULES (the previous, always-used default) if the
// workspace type or its role can't be resolved for any reason - erring on the
// side of the prior behavior rather than leaving a fresh signup with zero
// capabilities. Uses bare GetOne lookups (no WorkspaceId/UserAccessPerWorkspace
// on the query) the same way ValidateRoleAndItsExistence already does for the
// same role, since this runs before any session/access-scope exists for the
// user being created.
func ResolveWorkspaceTypeRoleCapabilities(workspaceTypeId string) []string {
	fallback := []string{ROOT_ALL_MODULES}

	if workspaceTypeId == "" {
		return fallback
	}

	workspaceType, err := WorkspaceTypeActions.GetOne(fireback.QueryDSL{UniqueId: workspaceTypeId})
	if err != nil || workspaceType == nil {
		return fallback
	}

	role, err := RoleActions.GetOne(fireback.QueryDSL{UniqueId: workspaceType.RoleId})
	if err != nil || role == nil {
		return fallback
	}

	capabilities := RoleCapabilitiesListIdGet(role)
	if len(capabilities) == 0 {
		return fallback
	}

	return capabilities
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
				Name               string
				WorkspacesAccesses []string
				UserRoles          map[string]*struct {
					Name     string
					Accesses []string
				}
			}{}
		}

		ws[item.WorkspaceId].Name = item.WorkspaceName

		if item.Type == "account_restrict" {
			// Bug fix: this used to re-init the whole UserRoles map (wiping every
			// other role already accumulated for this same workspace) whenever it hit
			// a *second* role_id it hadn't seen yet - the map-is-nil check and the
			// key-is-nil check were conflated into one condition. In practice this
			// meant a user holding more than one role in the same workspace (e.g. via
			// a Workspaceabacdefs.RoleEntity granting an extra role alongside their normal one)
			// silently lost every role but the last one processed - for root, that
			// could drop the seeded "root.*" wildcard role itself, causing
			// MeetsAccessLevel to reject actions root should always be allowed.
			if ws[item.WorkspaceId].UserRoles == nil {
				ws[item.WorkspaceId].UserRoles = map[string]*struct {
					Name     string
					Accesses []string
				}{}
			}
			if ws[item.WorkspaceId].UserRoles[item.RoleId] == nil {
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

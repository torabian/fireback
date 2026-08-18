package abac

import (
	"errors"
	"fmt"
	"os"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
)

func GetRoleByUniqueId(Id string) *abacdefs.RoleEntity {
	workspace := &abacdefs.RoleEntity{}
	fireback.GetDbRef().Where(abacdefs.RoleEntity{UniqueId: Id}).First(&workspace)

	return workspace
}

func GetWorkspaceByUniqueId(Id string) *abacdefs.WorkspaceEntity {
	workspace := &abacdefs.WorkspaceEntity{}
	fireback.GetDbRef().Where(abacdefs.WorkspaceEntity{UniqueId: Id}).First(&workspace)

	return workspace
}

func SyncWorkspaceDefaultWorkspaceTypes(db *gorm.DB, roles []*abacdefs.WorkspaceTypeEntity) error {
	var root = "root"
	return db.Transaction(func(tx *gorm.DB) error {

		for _, role := range roles {

			item := &abacdefs.WorkspaceTypeEntity{}
			err := tx.
				Model(&abacdefs.WorkspaceTypeEntity{}).
				Where(&abacdefs.WorkspaceTypeEntity{
					WorkspaceId: emigo.NullableOf(ROOT_VAR),
					UniqueId:    role.UniqueId,
				}).First(item).Error

			if err == gorm.ErrRecordNotFound {
				_, err := WorkspaceTypeActionCreate(role, fireback.QueryDSL{Tx: tx, WorkspaceId: root})

				if err != nil {
					return err
				}

			}
		}

		// fmt.Println("✓ Default roles are synchronized")

		return nil
	})

}

func SyncWorkspaceDefaultRoles(db *gorm.DB, roles []*abacdefs.RoleEntity) error {

	return db.Transaction(func(tx *gorm.DB) error {

		for _, role := range roles {
			item := &abacdefs.RoleEntity{}
			err := tx.
				Model(&abacdefs.RoleEntity{}).
				Where(&abacdefs.RoleEntity{WorkspaceId: role.WorkspaceId, UniqueId: role.UniqueId}).First(item).Error

			if err == gorm.ErrRecordNotFound {
				_, err := RoleActions.Create(role, fireback.QueryDSL{Tx: tx, WorkspaceId: role.WorkspaceId.OrDefault("")})

				if err != nil {
					return err
				}

			}
		}

		// fmt.Println("✓ Default roles are synchronized")

		return nil
	})

}

/**
*	Returns os user in the system, if it's added to fireback database.
*	You need to create user, workspace and it's roles before thi function to give you the user.
**/
func GetOsUserInFireback() (*abacdefs.UserEntity, error) {
	currentUser := fireback.GetOsUserWithPhone()

	var user *abacdefs.UserEntity = nil

	err2 := fireback.GetDbRef().Where(fireback.RealEscape("unique_id = ?", "OS_"+currentUser.Uid)).First(&user).Error
	if err2 != nil {
		return nil, err2
	}

	return user, nil
}

func SigninWithOsUser2(q fireback.QueryDSL) (*abacdefs.UserSessionDto, *fireback.IError) {
	user, role, workspace := GetOsHostUserRoleWorkspaceDef()

	return UnsafeGenerateUser(&GenerateUserDto{
		user:            user,
		workspace:       workspace,
		role:            role,
		createUser:      true,
		createWorkspace: true,
		createRole:      true,

		// We want always to be able to login regardless
		restricted: false,
	}, q)
}

func WorkpaceTypeToString(items []*abacdefs.WorkspaceTypeEntity) []string {
	result := []string{}
	for _, item := range items {
		result = append(result, item.UniqueId+" >>> "+item.Title+"("+item.Slug+")")
	}

	return result
}

type AdminCreationInfo struct {
	Token       string
	WorkspaceAs string
}

func CreateAdminTransaction(dto *abacdefs.ClassicSignupActionReq, setForRoot bool, query fireback.QueryDSL) (AdminCreationInfo, error) {

	result := AdminCreationInfo{}

	err := fireback.GetDbRef().Transaction(func(tx *gorm.DB) error {

		query.Tx = tx

		user, role, workspace, passport := GetEmailPassportSignupMechanism(dto)
		session, sessionError := UnsafeGenerateUser(&GenerateUserDto{

			createUser:      true,
			createWorkspace: true,
			createRole:      true,
			createPassport:  true,

			user:      user,
			role:      role,
			workspace: workspace,
			passport:  passport,

			// We want always to be able to login regardless
			restricted: true,
		}, query)

		if sessionError != nil {
			return sessionError
		}

		if session == nil {
			return errors.New("Session has not been created.")
		}

		if len(session.UserWorkspaces.Items) == 0 {
			return errors.New("User has no workspaces after generation")
		}

		workspaceAs := session.UserWorkspaces.Items[0].WorkspaceId.OrDefault("")

		if setForRoot {
			user, _ := session.User.Get()

			// GetEmailPassportSignupMechanism/UnsafeGenerateUser above always create a
			// throwaway "workspace" + "Admin" role first (createWorkspace/createRole are
			// unconditionally true, since the non-root signup path needs them) and
			// enrolled this user into it via CreateWorkspaceAndAssignUser - that's the
			// UserWorkspace row session.UserWorkspaces.Items[0] refers to. Bug fix: this
			// throwaway membership was never cleaned up once root's own "root" workspace
			// membership was granted below, so every root bootstrap (every `auth
			// --in-root=true` call, including withFirebackServer()'s in every Cypress
			// spec) left root belonging to *two* workspaces - the intended "root" one
			// plus this orphaned "Admin"-role one - which made WithSelfServiceRoutes.tsx
			// stop auto-selecting a workspace and show the "Select workspace" picker
			// instead (see workspace-role-capability-scoping.cy.ts, which assumes root
			// only ever has the one workspace).
			throwawayUserWorkspaceId := session.UserWorkspaces.Items[0].UniqueId

			query.WorkspaceId = ROOT_VAR
			workspaceAs = ROOT_VAR
			// user.Item.UserId is the "created/owned by" metadata field (blank for a
			// fresh signup, since nobody else created this user) - the user's actual
			// identity is UniqueId. Using UserId here left every root UserWorkspace/
			// WorkspaceRole row created below with an empty user reference.
			query.UserId = user.Item.UniqueId
			createdUserWorkspace, err2 := UserWorkspaceActions.Create(&abacdefs.UserWorkspaceEntity{
				UniqueId:    fireback.UUID(),
				UserId:      emigo.NullableOf(user.Item.UniqueId),
				WorkspaceId: emigo.NullableOf(ROOT_VAR),
			}, query)

			if err2 != nil {
				return err2
			}

			_, err3 := WorkspaceRoleActions.Create(&abacdefs.WorkspaceRoleEntity{
				UserWorkspaceId: emigo.NullableOf(createdUserWorkspace.UniqueId),
				RoleId:          emigo.NullableOf(ROOT_VAR),
				WorkspaceId:     emigo.NullableOf(ROOT_VAR),
			}, query)

			if err3 != nil {
				return err3
			}

			if _, err4 := WorkspaceRoleActions.RemoveEnqueue(fireback.DeleteRequest{
				Query:          "user_workspace_id = " + throwawayUserWorkspaceId,
				ForceImmediate: true,
			}, query); err4 != nil {
				return err4
			}

			if _, err5 := UserWorkspaceActions.RemoveEnqueue(fireback.DeleteRequest{
				Query:          "unique_id = " + throwawayUserWorkspaceId,
				ForceImmediate: true,
			}, query); err5 != nil {
				return err5
			}
		}

		exePath, err4 := os.Executable()
		if err4 == nil {
			fmt.Println("Workspace changed to :::", workspaceAs, " run `"+exePath+" ws view` to see the access scope")
		}

		result.WorkspaceAs = workspaceAs
		result.Token = session.Token

		return nil
	})

	return result, err
}

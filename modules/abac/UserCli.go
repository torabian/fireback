package abac

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

var ROOT_ALL_ACCESS = "root.*"
var ROOT_ALL_MODULES = "root.modules.*"

var OS_SIGNIN_CAPABILITIES []*abacdefs.CapabilityEntity = []*abacdefs.CapabilityEntity{
	{UniqueId: ROOT_ALL_ACCESS, Name: "Root"},
}

var TokenParseInformation cli.Command = cli.Command{

	Name:    "parse",
	Aliases: []string{"r", "del", "delete"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "token",
			Usage: "The token information",
			Value: "",
		},
	},
	Usage: "Extracts a token information, either JWT or internal tokens and prints on screen",
	Action: func(ctx context.Context, c *cli.Command) error {
		token := c.String("token")
		user, err := GetUserFromToken(token)
		if err != nil {
			log.Fatal("User which has this token does not exists")
		}

		query := fireback.CommonCliQueryDSLBuilder(c)
		query.UserId = user.UniqueId
		access, _ := GetUserAccessLevels(query)

		fmt.Println("Workspaces associated:")
		fmt.Println(access.Json())

		return nil
	},
}

type CliUserCreationDto struct {
	FirstName string
	LastName  string
	IsRoot    bool
}

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

/**
*  Call this when you are going to initialize a server, it will create root workspaces
*  It will create root workspace, assign the role to it.
 */
func RepairTheWorkspaces() error {
	{

		if role := GetRoleByUniqueId("root"); role == nil || role.UniqueId == "" {
			if _, err2 := CreateRootRoleInWorkspace("root"); err2 != nil {
				if !strings.Contains(err2.Error(), "Duplicate") {

					fmt.Println(err2)
				}
			}
		}
	}
	{
		item := &abacdefs.WorkspaceTypeEntity{}
		err := fireback.GetDbRef().Model(&abacdefs.WorkspaceTypeEntity{}).Where(&abacdefs.WorkspaceTypeEntity{UniqueId: "root"}).First(item).Error
		system := "system"
		if err == gorm.ErrRecordNotFound {
			err = fireback.GetDbRef().Create(&abacdefs.WorkspaceTypeEntity{WorkspaceId: emigo.NullableOf(system), UniqueId: "root", RoleId: ROOT_VAR}).Error
			if err != nil {
				return err
			}
		}
	}

	{

		item := &abacdefs.WorkspaceEntity{}
		err := fireback.GetDbRef().Model(&abacdefs.WorkspaceEntity{}).Where(&abacdefs.WorkspaceEntity{UniqueId: "root"}).First(item).Error

		description := "The root system which holds entire software data tree"
		if err == gorm.ErrRecordNotFound {
			err = fireback.GetDbRef().Create(&abacdefs.WorkspaceEntity{
				UniqueId: "root", Name: ROOT_VAR, Description: description,
				WorkspaceId: emigo.NullableOf(ROOT_VAR),
				TypeId:      ROOT_VAR,
			}).Error

			if err != nil {
				return err
			}

			_, err2 := CreateRootRoleInWorkspace("root")

			if err2 != nil && !strings.Contains(err2.Error(), "Duplicate") {
				return err2
			}

		}

		ws := GetWorkspaceByUniqueId("root")
		if ws == nil || ws.UniqueId != "root" {
			return errors.New(("ROOT_WORKSPACE_DOES_NOT_EXISTS"))
		}
	}

	{
		item := &abacdefs.WorkspaceEntity{}
		err := fireback.GetDbRef().Model(&abacdefs.WorkspaceEntity{}).Where(&abacdefs.WorkspaceEntity{UniqueId: "system"}).First(item).Error
		system := "system"
		if err == gorm.ErrRecordNotFound {
			description := "The workspace content which applies to everyworkspace"
			err = fireback.GetDbRef().Create(&abacdefs.WorkspaceEntity{WorkspaceId: emigo.NullableOf(system), UniqueId: "system", Name: system, Description: description}).Error

			if err != nil {
				return err
			}
		}

		ws := GetWorkspaceByUniqueId("system")
		if ws == nil || ws.UniqueId != "system" {
			return errors.New(("SYSTEM_WORKSPACE_DOES_NOT_EXISTS"))
		}
	}

	return nil
}

func CreateRootRoleInWorkspace(workspaceId string) (*abacdefs.RoleEntity, error) {
	sampleName := "Root Access"
	entity := &abacdefs.RoleEntity{
		UniqueId:           "root",
		WorkspaceId:        emigo.NullableOf(ROOT_VAR),
		Name:               sampleName,
		CapabilitiesListId: RoleCapabilitiesListIdOf([]string{ROOT_ALL_ACCESS}),
	}

	err := fireback.GetDbRef().
		Where(&abacdefs.RoleEntity{UniqueId: "root"}).
		FirstOrCreate(&entity).Error

	return entity, err
}

func CreateUserFromSchema(t *CliUserCreationDto) (*abacdefs.UserEntity, error) {

	userUniqueId := fireback.UUID()
	user := &abacdefs.UserEntity{
		UniqueId: userUniqueId,
	}

	err := fireback.GetDbRef().Create(&user).Error

	return user, err
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

func CreateUserInteractiveQuestions(query fireback.QueryDSL) (*abacdefs.ClassicSignupActionReq, bool, error) {
	dto := &abacdefs.ClassicSignupActionReq{}
	setForRoot := true
	defaultValue := "a@a.com"

	if result := fireback.AskForSelect("Method", []string{PASSPORT_METHOD_EMAIL, PASSPORT_METHOD_PHONE}); result != "" {
		dto.Type = result
		if result == PASSPORT_METHOD_PHONE {
			defaultValue = "+1000"
		}
	}

	if result := fireback.AskForInput(fireback.ToUpper(dto.Type), defaultValue); result != "" {
		dto.Value = result
	}

	if result := fireback.AskForInput("Password", "123321"); result != "" {
		dto.Password = result
	}

	if result := fireback.AskForInput("First name", "Ali"); result != "" {
		dto.FirstName = result
	}

	if result := fireback.AskForInput("Last name", "Torabi"); result != "" {
		dto.LastName = result
	}

	items, _, _ := WorkspaceTypeActions.Query(query)
	if result := fireback.AskForSelect("Workspace Type", WorkpaceTypeToString(items)); result != "" {
		dto.WorkspaceTypeId = emigo.NullableOf(result)
	}

	if result := fireback.AskForSelect("Add to root group? (workspace, role)", []string{"yes", "no"}); result != "" {
		if result == "no" {
			setForRoot = false
		} else if result == "yes" {
			setForRoot = true
		}
	}

	return dto, setForRoot, nil
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

func InteractiveUserAdmin(query fireback.QueryDSL) (AdminCreationInfo, error) {
	dto, setForRoot, _ := CreateUserInteractiveQuestions(query)
	return CreateAdminTransaction(dto, setForRoot, query)
}

// UserMockActionCliHandler is a hand-written cli command (not emi-generated, unlike the
// rest of UserCliFn's commands) that creates one or more mock users with realistic,
// randomly generated information - name, avatar, birth date, ip, primary address, phone
// number, job title, company and bio. It's built directly on top of
// UserActions.SeederInit (the same generator the seeders framework calls) and
// UserActionCreate (the same helper UserCreateAction's HTTP/CLI endpoint uses), so mock
// users are indistinguishable from ones created through the regular "user create"
// command, just filled in for you. Useful for local development and demos where an
// empty user list isn't very convincing.
func UserMockActionCliHandler() *cli.Command {
	return &cli.Command{
		Name:  "mock",
		Usage: "Creates one or more mock users with realistic randomly generated information (name, address, phone, job title, company, bio, ...) - useful for local development and demos.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "count",
				Usage: "How many mock users to create",
				Value: 1,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			count := int(c.Int("count"))
			if count <= 0 {
				count = 1
			}

			created := make([]*abacdefs.UserEntity, 0, count)
			for i := 0; i < count; i++ {
				entity := UserActions.SeederInit()
				entity.UniqueId = fireback.UUID()
				user, err := UserActionCreate(entity, fireback.QueryDSL{})
				if err != nil {
					return err
				}
				created = append(created, user)
			}

			out, encErr := json.MarshalIndent(map[string]any{
				"data": map[string]any{"items": created},
			}, "", "  ")
			if encErr != nil {
				return encErr
			}
			fmt.Println(string(out))
			return nil
		},
	}
}

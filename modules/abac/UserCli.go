package abac

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

var ROOT_ALL_ACCESS = "root.*"
var ROOT_ALL_MODULES = "root.modules.*"

var OS_SIGNIN_CAPABILITIES []*fireback.CapabilityEntity = []*fireback.CapabilityEntity{
	{UniqueId: ROOT_ALL_ACCESS, Visibility: fireback.NewString("A"), Name: "Root"},
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
	Action: func(c *cli.Context) error {
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

func GetRoleByUniqueId(Id string) *RoleEntity {
	workspace := &RoleEntity{}
	fireback.GetDbRef().Where(RoleEntity{UniqueId: Id}).First(&workspace)

	return workspace
}

func GetWorkspaceByUniqueId(Id string) *WorkspaceEntity {
	workspace := &WorkspaceEntity{}
	fireback.GetDbRef().Where(WorkspaceEntity{UniqueId: Id}).First(&workspace)

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
		item := &WorkspaceTypeEntity{}
		err := fireback.GetDbRef().Model(&WorkspaceTypeEntity{}).Where(&WorkspaceTypeEntity{UniqueId: "root"}).First(item).Error
		system := "system"
		if err == gorm.ErrRecordNotFound {
			err = fireback.GetDbRef().Create(&WorkspaceTypeEntity{WorkspaceId: fireback.NewString(system), UniqueId: "root", RoleId: fireback.NewString(ROOT_VAR)}).Error
			if err != nil {
				return err
			}
		}
	}

	{

		item := &WorkspaceEntity{}
		err := fireback.GetDbRef().Model(&WorkspaceEntity{}).Where(&WorkspaceEntity{UniqueId: "root"}).First(item).Error

		description := "The root system which holds entire software data tree"
		if err == gorm.ErrRecordNotFound {
			err = fireback.GetDbRef().Create(&WorkspaceEntity{
				UniqueId: "root", Name: ROOT_VAR, Description: description,
				WorkspaceId: fireback.NewString(ROOT_VAR),
				TypeId:      fireback.NewString(ROOT_VAR),
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
		item := &WorkspaceEntity{}
		err := fireback.GetDbRef().Model(&WorkspaceEntity{}).Where(&WorkspaceEntity{UniqueId: "system"}).First(item).Error
		system := "system"
		if err == gorm.ErrRecordNotFound {
			description := "The workspace content which applies to everyworkspace"
			err = fireback.GetDbRef().Create(&WorkspaceEntity{WorkspaceId: fireback.NewString(system), UniqueId: "system", Name: system, Description: description}).Error

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

func CreateRootRoleInWorkspace(workspaceId string) (*RoleEntity, error) {
	sampleName := "Root Access"
	entity := &RoleEntity{
		UniqueId:    "root",
		WorkspaceId: fireback.NewString(ROOT_VAR),
		Name:        sampleName,
		Capabilities: []*fireback.CapabilityEntity{
			{
				WorkspaceId: fireback.NewString("system"),
				Visibility:  fireback.NewString("A"),
				UniqueId:    ROOT_ALL_ACCESS,
			},
		},
	}

	err := fireback.GetDbRef().
		Where(&RoleEntity{UniqueId: "root"}).
		FirstOrCreate(&entity).Error

	return entity, err
}

func CreateUserFromSchema(t *CliUserCreationDto) (*UserEntity, error) {

	userUniqueId := fireback.UUID()
	user := &UserEntity{
		UniqueId: userUniqueId,
	}

	err := fireback.GetDbRef().Create(&user).Error

	return user, err
}

func SyncWorkspaceDefaultWorkspaceTypes(db *gorm.DB, roles []*WorkspaceTypeEntity) error {
	var root = "root"
	return db.Transaction(func(tx *gorm.DB) error {

		for _, role := range roles {

			item := &WorkspaceTypeEntity{}
			err := tx.
				Model(&WorkspaceTypeEntity{}).
				Where(&WorkspaceTypeEntity{WorkspaceId: fireback.NewString(ROOT_VAR), UniqueId: role.UniqueId}).First(item).Error

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

func SyncWorkspaceDefaultRoles(db *gorm.DB, roles []*RoleEntity) error {

	return db.Transaction(func(tx *gorm.DB) error {

		for _, role := range roles {
			item := &RoleEntity{}
			err := tx.
				Model(&RoleEntity{}).
				Where(&RoleEntity{WorkspaceId: role.WorkspaceId, UniqueId: role.UniqueId}).First(item).Error

			if err == gorm.ErrRecordNotFound {
				_, err := RoleActions.Create(role, fireback.QueryDSL{Tx: tx, WorkspaceId: role.WorkspaceId.String})

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
func GetOsUserInFireback() (*UserEntity, error) {
	currentUser := fireback.GetOsUserWithPhone()

	var user *UserEntity = nil

	err2 := fireback.GetDbRef().Where(fireback.RealEscape("unique_id = ?", "OS_"+currentUser.Uid)).First(&user).Error
	if err2 != nil {
		return nil, err2
	}

	return user, nil
}

func SigninWithOsUser2(q fireback.QueryDSL) (*UserSessionDto, *fireback.IError) {
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

func WorkpaceTypeToString(items []*WorkspaceTypeEntity) []string {
	result := []string{}
	for _, item := range items {
		result = append(result, item.UniqueId+" >>> "+item.Title+"("+item.Slug+")")
	}

	return result
}

func CreateUserInteractiveQuestions(query fireback.QueryDSL) (*ClassicSignupActionReqDto, bool, error) {
	dto := &ClassicSignupActionReqDto{}
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
		dto.WorkspaceTypeId = fireback.NewString(result)
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

func CreateAdminTransaction(dto *ClassicSignupActionReqDto, setForRoot bool, query fireback.QueryDSL) error {
	appConfig := fireback.GetConfig()

	return fireback.GetDbRef().Transaction(func(tx *gorm.DB) error {

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

		if len(session.UserWorkspaces) == 0 {
			return errors.New("User has no workspaces after generation")
		}

		workspaceAs := session.UserWorkspaces[0].WorkspaceId.String

		if setForRoot {

			query.WorkspaceId = ROOT_VAR
			workspaceAs = ROOT_VAR
			query.UserId = session.User.UserId.String
			_, err2 := UserWorkspaceActions.Create(&UserWorkspaceEntity{
				UniqueId:    fireback.UUID(),
				UserId:      session.User.UserId,
				WorkspaceId: fireback.NewString(ROOT_VAR),
			}, query)

			if err2 != nil {
				return err2
			}

			_, err3 := WorkspaceRoleActions.Create(&WorkspaceRoleEntity{
				RoleId:      fireback.NewString(ROOT_VAR),
				WorkspaceId: fireback.NewString(ROOT_VAR),
			}, query)

			if err3 != nil {
				return err3
			}
		}

		exePath, err4 := os.Executable()
		if err4 == nil {
			fmt.Println("Workspace changed to :::", workspaceAs, " run `"+exePath+" ws view` to see the access scope")
		}

		appConfig.CliWorkspace = workspaceAs
		appConfig.CliToken = session.Token
		appConfig.Save(".env")

		return nil
	})
}

func InteractiveUserAdmin(query fireback.QueryDSL) error {
	dto, setForRoot, _ := CreateUserInteractiveQuestions(query)
	return CreateAdminTransaction(dto, setForRoot, query)
}

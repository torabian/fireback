package workspaces

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/urfave/cli"
	"gorm.io/gorm"
)

var OS_SIGNIN_CAPABILITIES []*CapabilityEntity = []*CapabilityEntity{
	{UniqueId: "root/*"},
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

		query := CommonCliQueryDSLBuilder(c)
		query.UserId = user.UniqueId
		access, _ := GetUserAccessLevels(query)

		fmt.Println("Workspaces associated:")
		fmt.Println(access.Workspaces)
		fmt.Println(access.Capabilities)
		fmt.Println(access.SQL)

		return nil
	},
}

type CliUserCreationDto struct {
	FirstName string
	LastName  string
	IsRoot    bool
}

func getRoleEntityAsListItem(items []*RoleEntity) ([]string, error) {

	result := []string{}
	for _, role := range items {
		result = append(result, role.UniqueId+" >>> "+*role.Name)
	}
	return result, nil
}

func getWorkspaceEntitiesAsListItem(items []*WorkspaceEntity) ([]string, error) {

	result := []string{}
	for _, entity := range items {
		result = append(result, entity.UniqueId+" >>> "+*entity.Name)
	}
	return result, nil
}

/**
*  Call this when you are going to initialize a server, it will create root workspaces
*  It will create root workspace, assign the role to it.
 */
func RepairTheWorkspaces() error {
	{
		root := "root"

		item := &WorkspaceEntity{}
		err := GetDbRef().Model(&WorkspaceEntity{}).Where(&WorkspaceEntity{UniqueId: "root"}).First(item).Error

		description := "The root system which holds entire software data tree"
		if err == gorm.ErrRecordNotFound {
			err = GetDbRef().Create(&WorkspaceEntity{
				UniqueId: "root", Name: &root, Description: &description,
				WorkspaceId: &root,
			}).Error

			if err != nil {
				return err
			}

			_, err2 := CreateRootRoleInWorkspace("root")

			if err2 != nil {
				return err2
			}

		}

		ws := GetWorkspaceByUniqueId("root")
		if ws == nil || ws.UniqueId != "root" {
			return errors.New(("ROOT_WORKSPACE_DOES_NOT_EXISTS"))
		}
	}
	{

		if role := GetRoleByUniqueId("root"); role == nil || role.UniqueId == "" {
			if _, err2 := CreateRootRoleInWorkspace("root"); err2 != nil {
				fmt.Println(err2)
			}
		}
	}
	{
		item := &WorkspaceEntity{}
		err := GetDbRef().Model(&WorkspaceEntity{}).Where(&WorkspaceEntity{UniqueId: "system"}).First(item).Error
		system := "system"
		if err == gorm.ErrRecordNotFound {
			description := "The workspace content which applies to everyworkspace"
			err = GetDbRef().Create(&WorkspaceEntity{WorkspaceId: &system, UniqueId: "system", Name: &system, Description: &description}).Error

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
	sampleName := "Root Administrator"
	entity := &RoleEntity{
		UniqueId:    "root",
		WorkspaceId: &workspaceId,
		Name:        &sampleName,
		Capabilities: []*CapabilityEntity{
			{
				UniqueId: "root/*",
			},
		},
	}

	err := GetDbRef().Create(entity).Error

	return entity, err
}

func CreateUserFromSchema(t *CliUserCreationDto) (*UserEntity, error) {

	userUniqueId := UUID()
	user := &UserEntity{
		UniqueId: userUniqueId,
	}

	err := GetDbRef().Create(&user).Error

	return user, err
}

func SyncWorkspaceDefaultWorkspaceTypes(db *gorm.DB, roles []*WorkspaceTypeEntity) error {
	var root = "root"
	return db.Transaction(func(tx *gorm.DB) error {

		for _, role := range roles {

			item := &WorkspaceTypeEntity{}
			err := tx.
				Model(&WorkspaceTypeEntity{}).
				Where(&WorkspaceTypeEntity{WorkspaceId: &root, UniqueId: role.UniqueId}).First(item).Error

			if err == gorm.ErrRecordNotFound {
				fmt.Println("Create1", err)
				_, err := WorkspaceTypeActionCreate(role, QueryDSL{Tx: tx, WorkspaceId: root})
				fmt.Println("Create2", err)

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
				_, err := RoleActionCreate(role, QueryDSL{Tx: tx, WorkspaceId: *role.WorkspaceId})

				if err != nil {
					return err
				}

			}
		}

		// fmt.Println("✓ Default roles are synchronized")

		return nil
	})

}

func SigninUser(uniqueId string) string {

	until := time.Now().Add(time.Hour * time.Duration(12)).String()
	tokenString := GenerateSecureToken(32)
	GetDbRef().Create(&TokenEntity{
		UniqueId:   tokenString,
		UserId:     &uniqueId,
		ValidUntil: &until,
	})

	return tokenString
}

/**
*	Returns os user in the system, if it's added to fireback database.
*	You need to create user, workspace and it's roles before thi function to give you the user.
**/
func GetOsUserInFireback() (*UserEntity, error) {
	currentUser := GetOsUserWithPhone()

	var user *UserEntity = nil

	err2 := GetDbRef().Where("unique_id = ?", "OS_"+currentUser.Uid).First(&user).Error
	if err2 != nil {
		return nil, err2
	}

	return user, nil
}

// /**
// *	This is used for those apps which need to be used locally.
// *	OS User will be who ever we have installed.
// **/

// func SigninWithOsUser() (*UserEntity, string, error) {
// 	osUser := GetOsUserWithPhone()

// 	// If the user already exists, then give it back
// 	user, err2 := GetOsUserInFireback()
// 	if err2 == nil && user != nil {
// 		tokenString, _ := GenerateToken("OS_" + osUser.Uid)
// 		return user, tokenString, nil
// 	}

// 	// var token string

// 	err := GetDbRef().Transaction(func(tx *gorm.DB) error {

// 		user = &UserEntity{
// 			UniqueId: "OS_" + osUser.Uid,
// 			Person: &PersonEntity{
// 				FirstName: &osUser.Username,
// 			},
// 		}
// 		err := tx.Create(&user).Error

// 		if err != nil {
// 			return err
// 		}

// 		root := "root"
// 		name := osUser.Name + "'s workspace"
// 		workspace := &WorkspaceEntity{
// 			Name:     &name,
// 			UniqueId: "OS_WS_" + osUser.Uid,
// 			LinkerId: &root,
// 		}

// 		workspaceErr := tx.Create(&workspace).Error

// 		if workspaceErr != nil {
// 			return workspaceErr
// 		}

// 		osRole := "OS User"
// 		roleD := &RoleEntity{
// 			UniqueId:    UUID(),
// 			Name:        &osRole,
// 			WorkspaceId: &workspace.UniqueId,
// 			Capabilities: []*CapabilityEntity{
// 				{UniqueId: "root/*"},
// 			},
// 		}

// 		role, err2 := RoleActionCreate(roleD, QueryDSL{Tx: tx})

// 		if err2 != nil {
// 			return err
// 		}

// 		linker := UserRoleWorkspaceEntity{
// 			UserId:      &user.UniqueId,
// 			Role:        role,
// 			UniqueId:    UUID(),
// 			WorkspaceId: &workspace.UniqueId,
// 		}

// 		return tx.Create(&linker).Error

// 	})

// 	if err != nil {
// 		return nil, "", err
// 	}

// 	tokenString, err := GenerateToken(user.UniqueId)

// 	if err != nil {
// 		return nil, "", err
// 	}

// 	return user, tokenString, nil
// }

func (x *UserEntity) AuthorizeWithToken(q QueryDSL) (string, error) {
	tokenString := GenerateSecureToken(32)

	until := time.Now().Add(time.Hour * time.Duration(12)).String()

	err3 := GetRef(q).Create(&TokenEntity{
		UniqueId:   tokenString,
		UserId:     &x.UniqueId,
		ValidUntil: &until,
	}).Error

	if err3 != nil {
		return "", err3
	}

	return tokenString, nil
}

func SigninWithOsUser2(q QueryDSL) (*UserSessionDto, *IError) {
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

func GenerateToken(userId string) (string, error) {
	tokenString := GenerateSecureToken(32)

	until := time.Now().Add(time.Hour * time.Duration(12)).String()

	err3 := GetDbRef().Create(&TokenEntity{
		UniqueId:   tokenString,
		UserId:     &userId,
		ValidUntil: &until,
	}).Error

	if err3 != nil {
		return "", err3
	}

	return tokenString, nil
}

func InteractiveUserAdmin(query QueryDSL) error {
	dto := &ClassicSignupActionReqDto{}
	setForRoot := true
	if result := AskForInput("First name", "Ali"); result != "" {
		dto.FirstName = &result
	}

	if result := AskForInput("Last name", "Torabi"); result != "" {
		dto.LastName = &result
	}

	if result := AskForSelect("Method", []string{"email", "phonenumber"}); result != "" {
		dto.Type = &result
	}

	if result := AskForInput(ToUpper(*dto.Type), "admin"); result != "" {
		dto.Value = &result
	}

	if result := AskForInput("Password", "admin"); result != "" {
		dto.Password = &result
	}

	if result := AskForSelect("Add to root group? (workspace, role)", []string{"yes", "no"}); result != "" {
		if result == "no" {
			setForRoot = false
		} else if result == "yes" {
			setForRoot = true
		}
	}

	return dbref.Transaction(func(tx *gorm.DB) error {

		query.Tx = tx

		session, err := ClassicSignupAction(dto, query)
		if err != nil {
			return err
		}

		workspaceAs := *session.UserWorkspaces[0].WorkspaceId

		if setForRoot {

			query.WorkspaceId = ROOT_VAR
			workspaceAs = ROOT_VAR
			query.UserId = *session.User.UserId
			_, err2 := UserWorkspaceActionCreate(&UserWorkspaceEntity{
				UniqueId:    UUID(),
				UserId:      session.User.UserId,
				WorkspaceId: &ROOT_VAR,
			}, query)

			if err2 != nil {
				return err2
			}

			_, err3 := WorkspaceRoleActionCreate(&WorkspaceRoleEntity{
				RoleId:      &ROOT_VAR,
				WorkspaceId: &ROOT_VAR,
			}, query)

			if err3 != nil {
				return err3
			}
		}

		cfg := GetAppConfig()
		fmt.Println("Workspace changed to :::", workspaceAs, " run `f ws view` to see the access scope")

		cfg.WorkspaceAs = workspaceAs
		cfg.Token = *session.Token
		cfg.Save()

		return nil
	})
}

func InteractiveCreateUserInCli() *UserEntity {
	dto := &CliUserCreationDto{}
	result := AskForInput("First name", "")
	if result != "" {
		dto.FirstName = result
	}
	result = AskForInput("Last name", "")
	if result != "" {
		dto.LastName = result
	}

	items, meta, err := GetSystemWorkspacesAction(QueryDSL{ItemsPerPage: 20})

	if err != nil {
		fmt.Println(err.Error())
	}

	workspaces, err := getWorkspaceEntitiesAsListItem(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	// This is always there in database, so do not add it.
	// workspaces = append([]string{"root >>> The system root"}, ..)

	selectedWorkspace := ""
	if meta.TotalItems <= 20 {
		selectedWorkspace = AskForSelect("Select the workspace", workspaces)
		index := strings.Index(selectedWorkspace, ">>>")
		selectedWorkspace = strings.Trim(selectedWorkspace[0:index], " ")
	} else {
		selectedWorkspace = AskForInput("Too many workspaces, enter the unique id", "")
	}

	// 2. Ask user role
	roles, err := GetRolesInsideWorkspaceById(selectedWorkspace)

	if len(roles) == 0 {
		result = AskForSelect("There are no roles with root privilegs in the root workspace. Create now?", []string{"yes", "no"})

		if result == "yes" {

			role, err := CreateRootRoleInWorkspace(selectedWorkspace)

			if err != nil {
				log.Fatal(err)
			}
			roles = append(roles, role)
		}
	}

	selectRole, _ := getRoleEntityAsListItem(roles)
	selectedRoleId := AskForSelect("Which role you want to assign to this new user?", selectRole)
	index := strings.Index(selectedRoleId, ">>>")
	selectedRoleId = strings.Trim(selectedRoleId[0:index], " ")

	if user, err := CreateUserFromSchema(dto); err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		// ConnectWorkspaceUserToRole(selectedWorkspace, user, selectedRoleId)
		token := SigninUser(user.UniqueId)

		fmt.Println("Token:", token)
		fmt.Println("UserId:", user.UniqueId)

		return user
	}
}

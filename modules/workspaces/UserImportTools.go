package workspaces

/*
	This file is a focus on creating users, their passports, workspaces related to them
	The content could simulate a real software with 1000 users, and complex relation between
	them. It's mostly for testing and demo, and QA of the product itself, and other projects
	you might build using fireback.

	Managing user, role workspace is one of the top difficult tasks in a backend system
	and firebak tries to simplify that process for you.
*/

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
	seeders "github.com/torabian/fireback/modules/workspaces/mocks/User"
)

func ImportFromFs(req *ImportUserActionReqDto, q QueryDSL) (*OkayResponseDto, *IError) {

	var content ContentImport[UserImportDto]
	if err := ReadYamlFileEmbed[ContentImport[UserImportDto]](&seeders.ViewsFs, "fake-random-users.yml", &content); err != nil {
		return nil, IErrorFromString("Import file did not found")
	}
	bar := progressbar.Default(int64(len(content.Items)))
	for _, item := range content.Items {
		user, role, workspace, passport := CreateUserCatalog(&item)
		if _, err := UnsafeGenerateUser(&GenerateUserDto{
			user:            user,
			workspace:       workspace,
			role:            role,
			passport:        passport,
			createUser:      true,
			createWorkspace: true,
			createRole:      true,
			createPassport:  true,

			// We want always to be able to login regardless
			restricted: true,
		}, q); err != nil {
			fmt.Println("Error:", err)
		} else {

		}
		bar.Add(1)
		// time.Sleep(time)
	}

	return &OkayResponseDto{}, nil
}

func CreateUserCatalog(dto *UserImportDto) (*UserEntity, *RoleEntity, *WorkspaceEntity, *PassportEntity) {

	user := &UserEntity{
		UniqueId: "ux_" + *dto.Passports[0].Value,
		Person: &PersonEntity{
			UniqueId:  "ux_" + *dto.Passports[0].Value,
			FirstName: dto.Person.FirstName,
			LastName:  dto.Person.LastName,
		},
	}

	passwordHashed, _ := HashPassword(*dto.Passports[0].Password)
	method, _ := DetectSignupMechanismOverValue(*dto.Passports[0].Value)

	passport := &PassportEntity{
		UniqueId: "ps_" + *dto.Passports[0].Value,
		Value:    dto.Passports[0].Value,
		Password: &passwordHashed,
		Type:     &method,
	}

	// For now, it's random. But make sure later we have the track of workspaces
	wid := UUID()
	workspace := &WorkspaceEntity{
		Name:        dto.Person.FirstName,
		UniqueId:    wid,
		WorkspaceId: &wid,
		LinkerId:    &ROOT_VAR,
		ParentId:    &ROOT_VAR,
		TypeId:      &ROOT_VAR,
	}

	role := &RoleEntity{
		UniqueId:    "ROLE_WORKSPACE_" + UUID(),
		Name:        dto.Person.FirstName,
		WorkspaceId: &workspace.UniqueId,
		Capabilities: []*CapabilityEntity{
			{UniqueId: "root/*"},
		},
	}

	return user, role, workspace, passport
}

func init() {
	ImportUserActionImp = ImportFromFs
}

package abac

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
	"github.com/torabian/fireback/modules/fireback"
	seeders "github.com/torabian/fireback/modules/fireback/mocks/User"
)

func ImportFromFs(req *ImportUserActionReqDto, q fireback.QueryDSL) (*OkayResponseDto, *fireback.IError) {

	var content fireback.ContentImport[UserImportDto]
	if err := fireback.ReadYamlFileEmbed[fireback.ContentImport[UserImportDto]](&seeders.ViewsFs, "fake-random-users.yml", &content); err != nil {
		return nil, fireback.Create401Error(&AbacMessages.FileNotFound, []string{})
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
		UniqueId: "ux_" + dto.Passports[0].Value,
	}

	passwordHashed, _ := fireback.HashPassword(dto.Passports[0].Password)
	method, _ := DetectSignupMechanismOverValue(dto.Passports[0].Value)

	passport := &PassportEntity{
		UniqueId: "ps_" + dto.Passports[0].Value,
		Value:    dto.Passports[0].Value,
		Password: passwordHashed,
		Type:     method,
	}

	// For now, it's random. But make sure later we have the track of workspaces
	wid := fireback.UUID()
	workspace := &WorkspaceEntity{

		UniqueId:    wid,
		WorkspaceId: fireback.NewString(wid),
		LinkerId:    fireback.NewString(ROOT_VAR),
		ParentId:    fireback.NewString(ROOT_VAR),
		TypeId:      fireback.NewString(ROOT_VAR),
	}

	role := &RoleEntity{
		UniqueId: "ROLE_WORKSPACE_" + fireback.UUID(),

		WorkspaceId: fireback.NewString(workspace.UniqueId),
		Capabilities: []*fireback.CapabilityEntity{
			{UniqueId: ROOT_ALL_ACCESS, Visibility: fireback.NewString("A")},
		},
	}

	return user, role, workspace, passport
}

func init() {
	ImportUserActionImp = ImportFromFs
}

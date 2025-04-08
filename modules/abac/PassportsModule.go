package abac

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func PassportsModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{

		// it must write on the workspaces instead
		Name: "abac",
	}

	module.ProvideMockWriterHandler(func(languages []string) {
		for _, lang := range languages {
			var result *UserSessionDto
			if result != nil {

				workspaces.WriteMockDataToFile(lang, "", "UserSessionDto", gin.H{
					"data": gin.H{
						"user":     result.User,
						"passport": result.Passport,
						"token":    result.Token,
						// "userRoleWorkspaces": result.UserRoleWorkspaces,
					},
				})
			}
		}
	})

	module.ProvidePermissionHandler(
		ALL_PASSPORT_PERMISSIONS,
		ALL_PASSPORT_METHOD_PERMISSIONS,
		ALL_PUBLIC_JOIN_KEY_PERMISSIONS,
		ALL_ROLE_PERMISSIONS,
		ALL_USER_PERMISSIONS,
	)

	module.Actions = [][]workspaces.Module3Action{
		GetPassportMethodModule3Actions(),
		GetPassportModule3Actions(),
		GetPublicJoinKeyModule3Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {

		return dbref.AutoMigrate(
			&EmailConfirmationEntity{},
			&PhoneConfirmationEntity{},
			&PublicAuthenticationEntity{},
			&PassportEntity{},
			&PassportMethodEntity{},
			&PublicJoinKeyEntity{},
		)
	})

	module.ProvideCliHandlers([]cli.Command{
		PassportCli,
		PublicJoinKeyCliFn(),
		PublicAuthenticationCliFn(),
	})

	return module
}

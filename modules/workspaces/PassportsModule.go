package workspaces

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func PassportsModuleSetup() *ModuleProvider {
	module := &ModuleProvider{

		// it must write on the workspaces instead
		Name: "workspaces",
	}

	module.ProvideMockWriterHandler(func(languages []string) {
		for _, lang := range languages {
			var result *UserSessionDto
			if result != nil {

				WriteMockDataToFile(lang, "", "UserSessionDto", gin.H{
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
	)

	module.Actions = [][]Module3Action{
		GetPassportMethodModule3Actions(),
		GetPassportModule3Actions(),
		GetPublicJoinKeyModule3Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {

		return dbref.AutoMigrate(
			&EmailConfirmationEntity{},
			&PhoneConfirmationEntity{},
			&ForgetPasswordEntity{},
			&PassportEntity{},
			&PassportMethodEntity{},
			&PassportMethodEntityPolyglot{},
			&PublicJoinKeyEntity{},
		)
	})

	module.ProvideCliHandlers([]cli.Command{
		PassportCli,
	})

	return module
}

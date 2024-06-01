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
	firstName := "John"
	lastName := "Doe"
	email := "john"
	password := "123321"
	account1 := &ClassicAuthDto{FirstName: &firstName, LastName: &lastName, Value: &email, Password: &password}
	firstName = "علی"
	lastName = "ترابی"
	email = "ali"
	password = "123321"
	account2 := &ClassicAuthDto{FirstName: &firstName, LastName: &lastName, Value: &email, Password: &password}

	module.ProvideMockImportHandler(func() {
		f := QueryDSL{WorkspaceId: "root"}
		PassportActionEmailSignup(account1, f)
		PassportActionEmailSignup(account2, f)
	})
	module.ProvideMockWriterHandler(func(languages []string) {

		// 1. Write some users based on different langauges
		// f := QueryDSL{WorkspaceId: "root"}
		for _, lang := range languages {
			var result *UserSessionDto
			// var urw []*UserRoleWorkspaceEntity

			// if lang == "fa" {
			// 	result, _ = PassportActionEmailSignin(&EmailAccountSigninDto{Email: account1.Email, Password: account1.Password}, f)
			// 	urw2, _, _ := GetUserWorkspacesAndRolesAction(QueryDSL{UserId: result.User.UniqueId})
			// 	urw = urw2
			// } else {
			// 	result, _ = PassportActionEmailSignin(&EmailAccountSigninDto{Email: account2.Email, Password: account2.Password}, f)
			// 	urw2, _, _ := GetUserWorkspacesAndRolesAction(QueryDSL{UserId: result.User.UniqueId})
			// 	urw = urw2
			// }

			// WriteMockDataToFile(lang, "", "UserRoleWorkspaces", QueryEntitySuccessResult(f, urw, &QueryResultMeta{TotalItems: 100}))
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

	module.ProvideTranslationList(PassportTranslations)
	module.Actions = [][]Module2Action{
		GetPassportMethodModule2Actions(),
		GetPassportModule2Actions(),
		GetPublicJoinKeyModule2Actions(),
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
		PublicJoinKeyCliFn(),
	})

	return module
}

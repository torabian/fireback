package abac

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/torabian/fireback/modules/abac/migrations"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module3Definitions embed.FS

func AppMenuWriteQueryCteMock(ctx fireback.MockQueryContext) {
	for _, lang := range ctx.Languages {
		itemsPerPage := 9999
		if ctx.ItemsPerPage > 0 {
			itemsPerPage = ctx.ItemsPerPage
		}
		f := fireback.QueryDSL{ItemsPerPage: itemsPerPage, Language: lang, WithPreloads: ctx.WithPreloads, Deep: true}
		items, count, _ := AppMenuActions.CteQuery(f)
		result := fireback.QueryEntitySuccessResult(f, items, count)
		fireback.WriteMockDataToFile(lang, "", "AppMenu", result)
	}
}

func workspaceModuleCore(module *fireback.ModuleProvider) {

	module.ProvidePermissionHandler(
		ALL_WORKSPACE_CONFIG_PERMISSIONS,
		ALL_WORKSPACE_TYPE_PERMISSIONS,
		ALL_EMAIL_SENDER_PERMISSIONS,
		ALL_EMAIL_PROVIDER_PERMISSIONS,
		ALL_NOTIFICATION_CONFIG_PERMISSIONS,
		ALL_GSM_PROVIDER_PERMISSIONS,
		ALL_WORKSPACE_INVITE_PERMISSIONS,
		ALL_BACKUP_TABLE_META_PERMISSIONS,
		ALL_TABLE_VIEW_SIZING_PERMISSIONS,
		ALL_APP_MENU_PERMISSIONS,
		ALL_REGIONAL_CONTENT_PERMISSIONS,
		ALL_USER_WORKSPACE_PERMISSIONS,
		ALL_USER_PERMISSIONS,
		ALL_ROLE_PERMISSIONS,
		ALL_WORKSPACE_ROLE_PERMISSIONS,
		ALL_WORKSPACE_PERMISSIONS,
		ALL_PERM_ABAC_MODULE,
		ALL_TIMEZONE_GROUP_PERMISSIONS,
	)

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {
		items := []interface{}{
			&UserEntity{},
			&TokenEntity{},
			&PreferenceEntity{},
			&RoleEntity{},
			&WorkspaceEntity{},
			&WorkspaceInviteEntity{},
			&WorkspaceConfigEntity{},
			&WorkspaceTypeEntity{},
			&WorkspaceTypeEntityPolyglot{},
			&GsmProviderEntity{},
			&BackupTableMetaEntity{},
			&WorkspaceRoleEntity{},
			&UserWorkspaceEntity{},
			&RegionalContentEntity{},
			&TableViewSizingEntity{},
			&AppMenuEntity{},
			&AppMenuEntityPolyglot{},
			&TimezoneGroupEntity{},
			&TimezoneGroupEntityPolyglot{},
		}

		items2 := []interface{}{}
		items2 = append(items2, items...)

		for _, item := range items2 {

			if err := dbref.AutoMigrate(item); err != nil {
				fmt.Println("Migrating entity issue:", fireback.GetInterfaceName(item))
				return err
			}
		}

		// This is an important function, to create the root workspace.
		// root workspaces is the only, main workspace, which has every other workspace under it.
		return RepairTheWorkspaces()
	})

}

type MicroserviceSetupConfig struct {
	AuthorizationResolver WithAuthorizationPureImpl
}

// Inject this into any project as a complete solution
func AbacCompleteModules() []*fireback.ModuleProvider {
	return []*fireback.ModuleProvider{
		WorkspaceModuleSetup(),
		DriveModuleSetup(),
		NotificationModuleSetup(),
		PassportsModuleSetup(),
	}
}

func WorkspaceModuleSetup() *fireback.ModuleProvider {

	// Default Fireback authorization. You can Override this on microservices
	fireback.WithAuthorizationPure = WithAuthorizationPureDefault
	fireback.WithAuthorizationFn = WithAuthorizationFn
	fireback.WithSocketAuthorization = WithSocketAuthorization

	module := &fireback.ModuleProvider{
		Name:               "abac",
		Definitions:        &Module3Definitions,
		OnEnvInit:          OnInitEnvHook,
		GoMigrateDirectory: &migrations.MigrationsFs,
	}

	workspaceModuleCore(module)

	module.ProvideMockWriterHandler(func(languages []string) {
		// WorkspaceTypeWriteQueryMock(MockQueryContext{Languages: languages})
		// GsmProviderWriteQueryMock(MockQueryContext{Languages: languages})
		// AppMenuWriteQueryCteMock(MockQueryContext{Languages: languages})
	})

	module.ProvideTests(fireback.UserTests,
		[]fireback.Test{
			fireback.TestNewModuleProjectGen,
		},
		AppMenuTests,
		fireback.IntelisenseTest,
	)

	module.ProvideSeederImportHandler(func() {
		// We do not use syncing here.
		// Because fireback is being imported by other modules,
		// they might want their own unique menu items
		// sync items in the fireback/main or desktop one manually for this project.
		// for other projects extending fireback you can use here.
		TimezoneGroupSyncSeeders()
	})

	module.ProvideMockImportHandler(func() {
		// GsmProviderImportMocks()
	})

	module.Actions = [][]fireback.Module3Action{
		GetUserModule3Actions(),
		GetWorkspaceModule3Actions(),
		GetRoleModule3Actions(),
		GetWorkspaceTypeModule3Actions(),
		GetGsmProviderModule3Actions(),
		GetWorkspaceInviteModule3Actions(),
		GetBackupTableMetaModule3Actions(),
		GetTableViewSizingModule3Actions(),
		GetAppMenuModule3Actions(),
		GetEmailConfirmationModule3Actions(),
		AbacCustomActions(),
		GetUserWorkspaceModule3Actions(),
		GetWorkspaceRoleModule3Actions(),
		GetTimezoneGroupModule3Actions(),
		GetWorkspaceConfigModule3Actions(),
		GetRegionalContentModule3Actions(),
		// {
		// 	AS_FIREBACK_ACTION,
		// },
	}

	module.ProvideCliHandlers([]cli.Command{
		RoleCliFn(),
		UserCliFn(),
		WorkspaceCliFn(),
		MiscCli,
		TimezoneGroupCliFn(),
	})

	module.ProvideCliHandlers([]cli.Command{AuthFlow, AbacActions})

	return module
}

var AuthFlow cli.Command = cli.Command{
	Name:  "authorize",
	Usage: "Complete authorization flow via cli, similar to what would a public web app do.",
	Action: func(c *cli.Context) error {

		fmt.Println("In authorization flow, first of all we need to check what are the available actions publicly, to the user.")
		fmt.Println("On such authorization, there is no security model, since all actions are accessible publicly.")

		methods := []string{}

		// Since it's public, no need for any query dsl creation
		query := fireback.QueryDSL{ItemsPerPage: 9999}

		res, err := CheckPassportMethodsImpl(CheckPassportMethodsActionRequest{CliCtx: c}, query)
		if err != nil {
			return err
		}

		var passport CheckPassportMethodsActionRes
		if res, ok := res.Payload.(fireback.GoogleResponse[*CheckPassportMethodsActionRes]); ok {
			passport = *res.Data.Item
		} else {
			log.Fatalln("Checking passport methods publicly has failed, might be not available, or result is not back.")
			return nil
		}
		green := color.New(color.FgGreen)
		red := color.New(color.FgRed)

		fmt.Print("Authentication with email: ")
		if passport.Email {
			green.Print(passport.Email, "\n")
			methods = append(methods, "email")
		} else {
			red.Print(passport.Email, "\n")
		}

		fmt.Print("Authentication with phone: ")
		if passport.Phone {
			green.Print(passport.Phone, "\n")
			methods = append(methods, "phone")
		} else {
			red.Print(passport.Phone, "\n")
		}

		fmt.Print("Authentication with facebook: ")
		if passport.Facebook {
			green.Print(passport.Facebook, "\n")
			fmt.Println("Facebook App Id:", passport.FacebookAppId)
			methods = append(methods, "facebook")
		} else {
			red.Print(passport.Facebook, "\n")
		}

		fmt.Print("Authentication with google: ")
		if passport.Google {
			green.Print(passport.Google, "\n")
			fmt.Println("Google App Id:", passport.GoogleOAuthClientKey)
			methods = append(methods, "google")
		} else {
			red.Print(passport.Google, "\n")
		}

		fmt.Print("Recaptcha enabled: ")
		if passport.EnabledRecaptcha2 {
			green.Print(passport.EnabledRecaptcha2, "\n")
			fmt.Println("Recaptcha2 Client key:", passport.Recaptcha2ClientKey)
		} else {
			red.Print(passport.EnabledRecaptcha2, "\n")
		}

		selectedMethod := fireback.AskForSelect("Continue with method", methods)

		fmt.Println("Continuing with method: ", selectedMethod)

		if selectedMethod == "email" || selectedMethod == "phone" {
			label := "Enter the account"
			if selectedMethod == "email" {
				label = "Enter email address"
			} else if selectedMethod == "phone" {
				label = "Enter phone number"
			}

			value := fireback.AskForInput(label, "a@a.com")
			m, e := CheckClassicPassportAction(&CheckClassicPassportActionReqDto{
				Value: value,
			}, query)

			if e != nil {
				return e
			}

			fmt.Println("Flags we got: ", strings.Join(m.Flags, ","))
			fmt.Println("Next steps: ", strings.Join(m.Next, ","))
			if m.OtpInfo != nil {
				fmt.Println("Also otp information are present.")
				fmt.Println("Blocked until:", m.OtpInfo.BlockedUntil)
				fmt.Println("Second to unblock:", m.OtpInfo.SecondsToUnblock)
				fmt.Println("SuspendUntil:", m.OtpInfo.SuspendUntil)
				fmt.Println("Valid until:", m.OtpInfo.ValidUntil)
			}

			if len(m.Next) == 0 {
				fmt.Println("There are no next steps specified based on given account. This can be issue, why there are no next steps available at all.")
				os.Exit(2)
			}

			var nextStep = ""
			if len(m.Next) > 1 {
				nextStep = fireback.AskForSelect("How to continue?", m.Next)
			} else {
				fmt.Println("Currently only a single next step is available: ", m.Next[0])
				nextStep = m.Next[0]
			}

			if nextStep == "create-with-password" {
				workspaceType := cliSelectWorkspace()

				dto := ClassicSignupActionReqDto{
					Value: value,
					Type:  workspaceType.UniqueId,
				}
				if result := fireback.AskForInput("First name", "Ali"); result != "" {
					dto.FirstName = result
				}

				if result := fireback.AskForInput("Last name", "Torabi"); result != "" {
					dto.LastName = result
				}

				if result := fireback.AskForInput("Password", "123321"); result != "" {
					dto.Password = result
				}

				res2, err := ClassicSignupAction(&dto, query)
				if err != nil {
					return err
				}

				// Here needs to implemented.
				fmt.Println(res2.Json())

			}

			if nextStep == "signin-with-password" {
				var password = ""
				if result := fireback.AskForInput("Password", "123321"); result != "" {
					password = result
				}

				if signin, err := ClassicSigninAction(&ClassicSigninActionReqDto{
					Value:    value,
					Password: password,
				}, query); err != nil {
					return err
				} else {
					fmt.Println("Signin next steps: ", signin.Next)

					// In case the session is available, it's successful and checking further steps
					// is not required.
					if signin.Session != nil {
						var selectedWorkspace = ""
						if signin.Session.User.IsSet() && !signin.Session.User.IsNull() {
							fmt.Println("Signin successful as: ", signin.Session.User.Ptr().FirstName, signin.Session.User.Ptr().LastName)
						} else {
							fmt.Println("Successful signin, but no user is associated with this session")
						}

						// Check the workspaces. If there are more than 1, we ask user to choose.
						if len(signin.Session.UserWorkspaces) > 1 {
							var workspaces = []string{}
							for _, item := range signin.Session.UserWorkspaces {
								workspaces = append(workspaces, fmt.Sprintf("%v", item.WorkspaceId.String))
							}

							selectedWorkspace = fireback.AskForSelect("You have more than one workspace assigned to your account. Choose which one to continue", workspaces)
						}

						fmt.Println("Completed with:")
						fmt.Println("Token:", signin.Session.Token)
						if selectedWorkspace != "" {
							fmt.Println("Workspace Id:", selectedWorkspace)
						}

						config := fireback.GetConfig()
						config.CliToken = signin.Session.Token

						// Even if workspace is empty, we need to clear it.
						config.CliWorkspace = selectedWorkspace

						config.Save(".env")
					}

				}
			}

		}

		return nil
	},
}

func cliSelectWorkspace() *QueryWorkspaceTypesPubliclyActionResDto {
	var selectedWorkspace *QueryWorkspaceTypesPubliclyActionResDto = nil
	workspaceTypes, _, err := WorkspaceTypeActionPublicQuery(fireback.QueryDSL{ItemsPerPage: 9999})
	if err != nil {
		log.Fatalln("Error on reading workspace types from database: %w", err)
	}

	workspacesChoises := []string{}

	if len(workspaceTypes) == 0 {
		fmt.Println("You cannot create an account via this tool, because you do not have any workspace types setup. You need to use an root account, to create at least one workspace type, and then public flow of authentication would work.")
		os.Exit(1)
	} else if len(workspaceTypes) == 1 {
		selectedWorkspace = workspaceTypes[0]
	} else {
		for _, item := range workspaceTypes {
			workspacesChoises = append(workspacesChoises, fmt.Sprint("%v >>> %v (%v)", item.UniqueId, item.Title, item.Slug))
		}
		selectedWTId := fireback.AskForSelect("Which workspace type (account type) you are going to create?", workspacesChoises)
		for _, item := range workspaceTypes {
			if item.UniqueId == selectedWTId {
				selectedWorkspace = item
			}
		}
	}

	fmt.Println("Selected workspace type: ", selectedWorkspace.Title, " with id: ", selectedWorkspace.UniqueId)

	return selectedWorkspace
}

var AbacActions cli.Command = cli.Command{
	Name:        "abac",
	Usage:       "All actions which are available for abac module",
	Subcommands: GetAbacActionsCli(),
}

package abac

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
)

var ANONYMOUS_AUTHENTICATION = "anonymous"

func discoverPassportMethodsAndPrint(c *cli.Context) []string {
	fmt.Println("In authorization flow, first of all we need to check what are the available actions publicly, to the user.")
	fmt.Println("On such authorization, there is no security model, since all actions are accessible publicly.")

	methods := []string{
		fmt.Sprintf("%v >>> %v", ANONYMOUS_AUTHENTICATION, "Anonymous authentication"),
	}

	// Since it's public, no need for any query dsl creation
	query := fireback.QueryDSL{ItemsPerPage: 9999}

	res, err := CheckPassportMethodsImpl(CheckPassportMethodsActionRequest{CliCtx: c}, query)
	if err != nil {
		log.Fatalln("Error on checking passport methods: %w", err)
	}

	var passport CheckPassportMethodsActionRes
	if res, ok := res.Payload.(fireback.GoogleResponse[*CheckPassportMethodsActionRes]); ok {
		passport = *res.Data.Item
	} else {
		log.Fatalln("Checking passport methods publicly has failed, might be not available, or result is not back.")
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

	return methods
}

// This is a great a example, which implements all of the features of authentication in Fireback abac.
// important is, you can implement different flow, using functions which are available in module,
// this is one of the very common structures that you can see in most web apps.
// First of all, you check what are the available options, such as email, phone, google.
// Then user enters their account (email), and from there we check if we can login or signup.
// if otp is enforced, even if user has an account in the system will be vague, and we protect
// the users information in a system.
// This tool, is a cli only tool, because it would allow root account creation as well, in case
// system has no user accounts created.
var AuthFlow cli.Command = cli.Command{
	Name:  "authorize",
	Usage: "All in one authorization tool into abac module, creates, authenticates end-to-end and can set cli workspace token.",
	Action: func(c *cli.Context) error {
		query := fireback.QueryDSL{ItemsPerPage: 9999}

		methods := discoverPassportMethodsAndPrint(c)

		selectedMethod := fireback.AskForSelect("Continue with method", methods)

		fmt.Println("Continuing with method: ", selectedMethod)

		if selectedMethod == "email" || selectedMethod == "phone" || selectedMethod == ANONYMOUS_AUTHENTICATION {
			prefix := "a@a.com"
			label := "Enter the account"
			switch selectedMethod {
			case "email":
				label = "Enter email address"
			case "phone":
				label = "Enter phone number"
			case ANONYMOUS_AUTHENTICATION:
				label = "Enter the anonymous prefix, which can be a random string generated as cookie from browser, with anonymous_* prefix"
				prefix = fmt.Sprintf("anonymous_%v", fireback.UUID())
			}

			value := fireback.AskForInput(label, prefix)
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

				// Now when we are creating a user in cli mode, which shouldn't be able
				// in web version, we can allow a direct root user to be created.
				// In such scenario, a workspace type is not needed.

				workspaceType := UNSAFE_allow_selection_of_workspace_type()

				dto := ClassicSignupActionReqDto{
					Value: value,
					Type:  workspaceType.UniqueId,
				}

				// We need such information, non-anonymous account creation.
				// For anonymous, it's enough to have a unique value, which can be random
				// value generated upon user initiate of website, and store it in cookie,
				// so we keep track of his behavior without him being logged in.
				if selectedMethod != ANONYMOUS_AUTHENTICATION {

					if result := fireback.AskForInput("First name", "Ali"); result != "" {
						dto.FirstName = result
					}

					if result := fireback.AskForInput("Last name", "Torabi"); result != "" {
						dto.LastName = result
					}

					if result := fireback.AskForInput("Password", "123321"); result != "" {
						dto.Password = result
					}
				}

				res2, err := ClassicSignupAction(&dto, query)
				if err != nil {
					return err
				}

				if workspaceType.UniqueId == ROOT_VAR {
					user, _ := res2.Session.User.Get()

					query.WorkspaceId = ROOT_VAR
					query.UserId = user.UserId.String
					_, err2 := UserWorkspaceActions.Create(&UserWorkspaceEntity{
						UniqueId:    fireback.UUID(),
						UserId:      user.UserId,
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

				// Now we need to select a workspace.
				var selectedWorkspace = ""
				var workspaces = []string{}
				for _, item := range res2.Session.UserWorkspaces {
					workspaces = append(workspaces, fmt.Sprintf("%v", item.WorkspaceId.String))
				}

				if workspaceType.UniqueId == ROOT_VAR {
					workspaces = append(workspaces, ROOT_VAR)
				}

				if res2.Session != nil {
					fmt.Println("Token:", res2.Session.Token)
					if fireback.AskBoolean("Session is created. Do you want to authorize the cli as well?") {

						if len(workspaces) > 1 {
							selectedWorkspace = fireback.AskForSelect("You have more than one workspace assigned to your account. Choose one to continue", workspaces)
						} else if len(workspaces) == 1 {
							selectedWorkspace = workspaces[0]
						}

						config := fireback.GetConfig()
						config.CliToken = res2.Session.Token
						config.CliWorkspace = selectedWorkspace
						config.Save(".env")
					}
				}

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
						config.CliWorkspace = selectedWorkspace

						config.Save(".env")
					}

				}
			}

		}

		return nil
	},
}

func UNSAFE_allow_selection_of_workspace_type() *QueryWorkspaceTypesPubliclyActionResDto {
	var selectedWorkspace *QueryWorkspaceTypesPubliclyActionResDto = nil
	workspaceTypes, _, err := WorkspaceTypeActionPublicQuery(fireback.QueryDSL{ItemsPerPage: 9999})
	if err != nil {
		log.Fatalln("Error on reading workspace types from database: %w", err)
	}

	workspacesChoises := []string{ROOT_VAR}

	if len(workspaceTypes) == 0 {
		fmt.Println("There are no workspace types available, it means only you can create a root account via this tool.")
		selectedWorkspace = &QueryWorkspaceTypesPubliclyActionResDto{
			UniqueId: ROOT_VAR,
			Title:    "ROOT",
		}
	} else {
		for _, item := range workspaceTypes {
			workspacesChoises = append(workspacesChoises, fmt.Sprintf("%v >>> %v (%v)", item.UniqueId, item.Title, item.Slug))
		}
		selectedWTId := fireback.AskForSelect("Which workspace type (account type) you are going to create?", workspacesChoises)
		for _, item := range workspaceTypes {
			if item.UniqueId == selectedWTId {
				selectedWorkspace = item
			}
		}
	}

	return selectedWorkspace
}

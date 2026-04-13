package abac

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pquerna/otp/totp"
	emigo "github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"
)

var ANONYMOUS_AUTHENTICATION = "anonymous"

func discoverPassportMethodsAndPrint(c *cli.Context) []string {
	fmt.Println("In authorization flow, first of all we need to check what are the available actions publicly, to the user.")
	fmt.Println("On such authorization, there is no security model, since all actions are accessible publicly.")

	methods := []string{
		fmt.Sprintf("%v >>> %v", ANONYMOUS_AUTHENTICATION, "Anonymous, can be also created as root, with a unique identifier"),
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
		methods = append(methods, "email >>> Create or sign-in with an email address, otp, totp might be required.")
	} else {
		red.Print(passport.Email, "\n")
	}

	fmt.Print("Authentication with phone: ")
	if passport.Phone {
		green.Print(passport.Phone, "\n")
		methods = append(methods, "phone >>> Phone number sign-in or create account, otp, totp might be required.")
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

// @WARNING: CLI ONLY. NEVER EXPOSE IN HTTP or Public operation, allows root access.
// This is a great a example, which implements all of the features of authentication in Fireback abac.
// important is, you can implement different flow, using functions which are available in module,
// this is one of the very common structures that you can see in most web apps.
// First of all, you check what are the available options, such as email, phone, google.
// Then user enters their account (email), and from there we check if we can login or signup.
// if otp is enforced, even if user has an account in the system will be vague, and we protect
// the users information in a system.
// This tool, is a cli only tool, because it would allow root account creation as well, in case
// system has no user accounts created.
func IntegrateAuthFlow(c *cli.Context) error {
	query := fireback.QueryDSL{ItemsPerPage: 9999}

	var sessionSecret = ""
	methods := discoverPassportMethodsAndPrint(c)

	selectedMethod := fireback.AskForSelect("Available authentication methods via cli (might be different than web)", methods)

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
			label = "Passport (account) anonymous_* prefix"
			prefix = fmt.Sprintf("anonymous_%v", fireback.UUID())
		}

		value := fireback.AskForInput(label, prefix)
		query.C = c
		mresponse, e := CheckClassicPassportAction(CheckClassicPassportActionRequest{
			Body: CheckClassicPassportActionReq{
				Value: value,
			},
			CliCtx: c,
		}, query)

		if e != nil {
			return e
		}

		var m *CheckClassicPassportActionRes

		if mf, ok := mresponse.Payload.(fireback.GoogleResponse[*CheckClassicPassportActionRes]); ok {
			m = mf.Data.Item
		}

		fmt.Println("Flags we got: ", strings.Join(m.Flags, ","))
		fmt.Println("Next steps: ", strings.Join(m.Next, ","))
		if m.OtpInfo.IsSet() {
			otp := m.OtpInfo.Ptr()
			fmt.Println("Also otp information are present.")
			fmt.Println("Blocked until:", otp.BlockedUntil)
			fmt.Println("Second to unblock:", otp.SecondsToUnblock)
			fmt.Println("SuspendUntil:", otp.SuspendUntil)
			fmt.Println("Valid until:", otp.ValidUntil)
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

		if nextStep == "otp" {
			otpCode := fireback.AskForInput("Enter the otp code. You might see it a bit above this command.", "")

			res, err := ClassicPassportOtpAction(&ClassicPassportOtpActionReqDto{
				Value: value, Otp: otpCode,
			}, query)

			if err != nil {
				fmt.Println("Not nil")
				return err
			}

			if res.ContinueWithCreation {
				fmt.Println("We continue to create account.")
				nextStep = "create-with-password"
				sessionSecret = res.SessionSecret
			} else if res.Session != nil {

				/// Now we need to set the session here, but code is duplicated multiple times
			}
		}

		if nextStep == "create-with-password" {

			// Now when we are creating a user in cli mode, which shouldn't be able
			// in web version, we can allow a direct root user to be created.
			// In such scenario, a workspace type is not needed.

			workspaceType := UNSAFE_allow_selection_of_workspace_type()

			dto := ClassicSignupActionReq{
				Value:           value,
				Type:            selectedMethod,
				WorkspaceTypeId: emigo.NullableOf(workspaceType.UniqueId),
				SessionSecret:   sessionSecret,
			}

			// We need such information, non-anonymous account creation.
			// For anonymous, it's enough to have a unique value, which can be random
			// value generated upon user initiate of website, and store it in cookie,
			// so we keep track of his behavior without him being logged in.
			if selectedMethod == ANONYMOUS_AUTHENTICATION {
				dto.FirstName = "Anonymous"
				dto.LastName = "Anonymous"
				dto.Password = "Anonymous"
			} else {

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

			query.WorkspaceId = ROOT_VAR
			result, err := ClassicSignupAction(ClassicSignupActionRequest{
				Body:   dto,
				CliCtx: c,
			}, query)

			if err != nil {
				return err
			}

			resEnvelope, ok := result.Payload.(fireback.GoogleResponse[ClassicSignupActionRes])
			if !ok {
				fmt.Println("Internal error on getting account creation results")
				os.Exit(4)
			}
			res2 := resEnvelope.Data.Item

			if res2.ContinueToTotp {
				fmt.Println("You need to setup time based token (totp) for this account.")
				fmt.Println("URL: ", res2.TotpUrl)

				u, err := url.Parse(res2.TotpUrl)
				if err != nil {
					panic(err)
				}

				secret := u.Query().Get("secret")
				fmt.Println("Secret:", secret)

				fmt.Println("You have to store the secret somewhere - usually done in the mobile app.")
				fmt.Println("Also, in order to avoid mobile app, and see the code, you can run this command:")
				fmt.Println("----")
				fmt.Println("" + fireback.GetExePath() + " misc totp --secret " + secret + " \n ")
				fmt.Println("----")

				now := time.Now()
				code, err := totp.GenerateCode(secret, now)
				if err != nil {
					return err
				}

				m, err := ConfirmClassicPassportTotpAction(&ConfirmClassicPassportTotpActionReqDto{
					Value:    value,
					Password: dto.Password,
					TotpCode: code,
				}, query)

				if m.Session != nil {
					authenticateCliWithSession(m.Session, workspaceType.UniqueId)
				}

			}

			if workspaceType.UniqueId == ROOT_VAR && res2.Session.Token != "" {
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

			if res2.Session.Token != "" {
				authenticateCliWithSession(&res2.Session, workspaceType.UniqueId)
			}

		}

		if nextStep == "signin-with-password" {
			var password = ""
			if result := fireback.AskForInput("Password", "123321"); result != "" {
				password = result
			}

			if result, err := ClassicSigninAction(ClassicSigninActionRequest{
				Body: ClassicSigninActionReq{
					Value:    value,
					Password: password,
				},
			}, query); err != nil {
				return err
			} else {

				resEnvelope, ok := result.Payload.(fireback.GoogleResponse[ClassicSigninActionRes])
				if !ok {
					fmt.Println("Critical internal error on casting signin result")
					os.Exit(1)
				}

				signin := resEnvelope.Data.Item
				fmt.Println("Signin next steps: ", signin.Next)

				// In case the session is available, it's successful and checking further steps
				// is not required.
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

	return nil
}

var AuthFlow cli.Command = cli.Command{
	Name:      "authorize",
	ShortName: "auth",
	Usage:     "All in one authorization tool into abac module, creates, authenticates end-to-end and can set cli workspace token.",
	Action: func(c *cli.Context) error {

		// In case that there are flags, means the interactive operation is not needed
		// this is useful to create an account in one go.
		if c.NumFlags() > 0 {
			appConfig := fireback.GetConfig()
			dto := CastClassicSignupActionReqFromCli(c)
			query := fireback.CommonCliQueryDSLBuilder(c)

			fmt.Println("Type", dto.Type)
			if result, err := CreateAdminTransaction(&dto, c.Bool("in-root"), query); err != nil {
				log.Fatalln(err)
			} else {
				appConfig.CliWorkspace = result.WorkspaceAs
				appConfig.CliToken = result.Token
				appConfig.Save(".env")
			}
			return nil
		}

		return IntegrateAuthFlow(c)
	},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "in-root",
			Usage: "Append this user to root group",
		},
		&cli.StringFlag{
			Name:  "value",
			Usage: "value",
		},
		&cli.StringFlag{
			Name:  "workspace-type-id",
			Usage: "The workspace type id, you can use 'root'",
		},
		&cli.StringFlag{
			Name:  "session-secrent",
			Usage: "The secret generated through the otp authentication process.",
		},
		&cli.StringFlag{
			Name:  "type",
			Usage: "One of: 'phonenumber', 'email'",
		},
		&cli.StringFlag{
			Name:  "password",
			Usage: "password",
		},
		&cli.StringFlag{
			Name:  "first-name",
			Usage: "firstName",
		},
		&cli.StringFlag{
			Name:  "last-name",
			Usage: "lastName",
		},
	},
}

func authenticateCliWithSession(session *UserSessionDto, workspaceTypeId string) {
	// Now we need to select a workspace.
	var selectedWorkspace = ""
	var workspaces = []string{}
	for _, item := range session.UserWorkspaces {
		workspaces = append(workspaces, fmt.Sprintf("%v", item.WorkspaceId.String))
	}

	if workspaceTypeId == ROOT_VAR {
		workspaces = append(workspaces, ROOT_VAR)
	}

	fmt.Println("Token:", session.Token)
	if fireback.AskBoolean("Session is created. Do you want to authorize the cli as well?") {

		if len(workspaces) > 1 {
			selectedWorkspace = fireback.AskForSelect("You have more than one workspace assigned to your account. Choose one to continue", workspaces)
		} else if len(workspaces) == 1 {
			selectedWorkspace = workspaces[0]
		}

		config := fireback.GetConfig()
		config.CliToken = session.Token
		config.CliWorkspace = selectedWorkspace
		config.Save(".env")
	}
}

func UNSAFE_allow_selection_of_workspace_type() *QueryWorkspaceTypesPubliclyActionRes {
	var selectedWorkspace *QueryWorkspaceTypesPubliclyActionRes = &QueryWorkspaceTypesPubliclyActionRes{
		UniqueId: ROOT_VAR,
		Title:    "ROOT",
	}

	workspaceTypes, _, err := WorkspaceTypeActionPublicQuery(fireback.QueryDSL{ItemsPerPage: 9999})
	if err != nil {
		log.Fatalln("Error on reading workspace types from database: %w", err)
	}

	workspacesChoises := []string{ROOT_VAR}

	for _, item := range workspaceTypes {
		workspacesChoises = append(workspacesChoises, fmt.Sprintf("%v >>> %v (%v)", item.UniqueId, item.Title, item.Slug))
	}
	selectedWTId := fireback.AskForSelect("Which workspace type (account type) you are going to create?", workspacesChoises)
	for _, item := range workspaceTypes {
		if item.UniqueId == selectedWTId {
			selectedWorkspace = item
		}
	}

	return selectedWorkspace
}

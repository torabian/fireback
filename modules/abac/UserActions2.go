//go:build !wasm

package abac

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
)

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
	Action: func(ctx context.Context, c *cli.Command) error {
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

// UserCliFn mirrors the old Module3-generated grouped "user" cli command (minus the
// import/export/dev commands, which had no hand-written equivalent to recover) - tokens,
// accept-invite and user-invitations are nested here too, since they're user-scoped
// actions without their own cli home (preserved from the pre-migration abacdefs.UserEntity.go
// hand file's "Tokens are related to users, so let's move them there." comment).
func UserCliFn() *cli.Command {
	return &cli.Command{
		Name:        "user",
		Description: `Manage the users who are in the current app (root only)`,
		Usage:       `Manage the users who are in the current app (root only)`,
		Commands: []*cli.Command{
			abacdefs.UserBrowseActionCliHandler(UserBrowseAction),
			abacdefs.UserGetActionCliHandler(UserGetAction),
			abacdefs.UserCreateActionCliHandler(UserCreateAction),
			abacdefs.UserUpdateActionCliHandler(UserUpdateAction),
			abacdefs.UserAwareDeletePreviewActionCliHandler(UserAwareDeletePreviewAction),
			abacdefs.UserAwareDeleteActionCliHandler(UserAwareDeleteAction),
			abacdefs.TokenBrowseActionCliHandler(TokenBrowseAction),
			abacdefs.TokenGetActionCliHandler(TokenGetAction),
			abacdefs.TokenCreateActionCliHandler(TokenCreateAction),
			abacdefs.TokenUpdateActionCliHandler(TokenUpdateAction),
			abacdefs.TokenAwareDeletePreviewActionCliHandler(TokenAwareDeletePreviewAction),
			abacdefs.TokenAwareDeleteActionCliHandler(TokenAwareDeleteAction),
			abacdefs.AcceptInviteActionCliHandler(AcceptInviteAction),
			abacdefs.UserInvitationsActionCliHandler(UserInvitationsAction),
			UserMockActionCliHandler(),
		},
	}
}

func CreateUserInteractiveQuestions(query fireback.QueryDSL) (*abacdefs.ClassicSignupActionReq, bool, error) {
	dto := &abacdefs.ClassicSignupActionReq{}
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
		dto.WorkspaceTypeId = emigo.NullableOf(result)
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

func InteractiveUserAdmin(query fireback.QueryDSL) (AdminCreationInfo, error) {
	dto, setForRoot, _ := CreateUserInteractiveQuestions(query)
	return CreateAdminTransaction(dto, setForRoot, query)
}

// UserMockActionCliHandler is a hand-written cli command (not emi-generated, unlike the
// rest of UserCliFn's commands) that creates one or more mock users with realistic,
// randomly generated information - name, avatar, birth date, ip, primary address, phone
// number, job title, company and bio. It's built directly on top of
// UserActions.SeederInit (the same generator the seeders framework calls) and
// UserActionCreate (the same helper UserCreateAction's HTTP/CLI endpoint uses), so mock
// users are indistinguishable from ones created through the regular "user create"
// command, just filled in for you. Useful for local development and demos where an
// empty user list isn't very convincing.
func UserMockActionCliHandler() *cli.Command {
	return &cli.Command{
		Name:  "mock",
		Usage: "Creates one or more mock users with realistic randomly generated information (name, address, phone, job title, company, bio, ...) - useful for local development and demos.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "count",
				Usage: "How many mock users to create",
				Value: 1,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			count := int(c.Int("count"))
			if count <= 0 {
				count = 1
			}

			created := make([]*abacdefs.UserEntity, 0, count)
			for i := 0; i < count; i++ {
				entity := UserActions.SeederInit()
				entity.UniqueId = fireback.UUID()
				user, err := UserActionCreate(entity, fireback.QueryDSL{})
				if err != nil {
					return err
				}
				created = append(created, user)
			}

			out, encErr := json.MarshalIndent(map[string]any{
				"data": map[string]any{"items": created},
			}, "", "  ")
			if encErr != nil {
				return encErr
			}
			fmt.Println(string(out))
			return nil
		},
	}
}

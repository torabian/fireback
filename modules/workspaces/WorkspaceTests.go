package workspaces

import (
	"fmt"
	"log"
	"testing"

	"github.com/fatih/color"
)

type WorkspaceCreationCtx struct {
	F QueryDSL
}

type Test struct {
	Name     string
	Function func(t *TestContext) error
}

type TestContext struct {
	testing.T
	F QueryDSL
	// Include any necessary fields or methods for test context
}

func (tc *TestContext) ErrorLn(message string, args ...interface{}) {
	// Implement error reporting mechanism
	c := color.New(color.FgRed)
	c.Print("Test Has failed unfortunately:")
	log.Fatalln(message, args)

}

func (tc *TestContext) Log(message string, args ...interface{}) {
	// Implement error reporting mechanism
	c := color.New(color.Faint)
	c.Println(message, args)

}

func (tc *TestContext) Errorf(message string, args ...interface{}) {
	// Implement error reporting mechanism
	c := color.New(color.FgRed)
	c.Print("Test Has failed unfortunately:")
	log.Fatalln(message, args)

}
func (tc *TestContext) Error(message string, args ...interface{}) {
	// Implement error reporting mechanism
	c := color.New(color.FgRed)
	c.Println(message, args)

}

func RunTests(F QueryDSL) {
	testContext := &TestContext{F: F} // Initialize test context

	tests := []Test{
		{
			Name: "Workspace without name must fail",
			Function: func(t *TestContext) error {

				if result, err := WorkspaceActionCreate(&WorkspaceEntity{}, t.F); err == nil {
					t.ErrorLn("Workspace name cannot be empty", result.UniqueId, "has been created unexpectely")
					return err
				}

				return nil
			},
		},
		{
			Name: "Should be able to create a workspace, and another workspace under that",
			Function: func(t *TestContext) error {

				parentWsName := "Main workspace"
				childWsName := "Child workspace"
				parent, err := WorkspaceActionCreate(&WorkspaceEntity{Name: &parentWsName}, t.F)
				if err != nil {
					t.ErrorLn("First workspace did not create to begin with:", err)
					return err
				}

				child, err2 := WorkspaceActionCreate(&WorkspaceEntity{
					Name:     &childWsName,
					ParentId: &parent.UniqueId,
				}, t.F)
				if err2 != nil {
					t.ErrorLn("Second workspace did not created")
					return err
				}

				if *child.ParentId != parent.UniqueId {
					t.ErrorLn(
						"Expected the parent id of child to be unique id, but it's not",
						"parent id:",
						parent.UniqueId,
						"child id",
						child.UniqueId,
						"child parent id",
						*child.ParentId,
					)
				}
				return nil
			},
		},
		{
			Name: "It should be able to find the created workspaces, and they must be representing each other",
			Function: func(t *TestContext) error {

				if _, _, err := WorkspaceActionQuery(t.F); err != nil {
					t.ErrorLn("Workspaces could not be queried from database")
					return err
				}

				return nil
			},
		},
		{
			Name: "Should be able to authorize the user using the host os credentials",
			Function: func(t *TestContext) error {

				if session, err := SigninWithOsUser2(t.F); err != nil {
					t.ErrorLn("Error on signin", err)
					return err
				} else {

					user, _, workspace := GetOsHostUserRoleWorkspaceDef()

					// Test the user to have correct values
					t.F.UniqueId = user.UniqueId
					if userdb, err := UserActionGetOne(t.F); err != nil {
						t.ErrorLn("Error on finding the created user in database", err)
						return err
					} else {
						if userdb.UniqueId != user.UniqueId {
							t.ErrorLn("Generated user is not correct", err)
							return err
						}
					}

					// Test the created workspace
					t.F.UniqueId = workspace.UniqueId
					t.F.WorkspaceId = workspace.UniqueId

					if workspacedb, err := WorkspaceActionGetOne(t.F); err != nil {
						t.ErrorLn("Error on finding created workspace in database", err)
						return err
					} else {
						if workspacedb.UniqueId != *workspacedb.WorkspaceId {
							t.ErrorLn("Workspace id must be equal to it's unique id in database, but its not", err, workspacedb)
							return err
						}

						if workspace.UniqueId != *workspacedb.WorkspaceId {
							t.ErrorLn(
								"Created workspace id is not equal to the definition workspace",
								err,
								"in db:",
								workspacedb,
								"in workspace:",
								workspace,
							)
							return err
						}
					}

					if session.Token == nil || len(*session.Token) < 10 {
						t.ErrorLn(
							"Token from the session has length less than 10, it's suspisous",
						)
					}

					if len(session.UserWorkspaces) < 1 {
						t.ErrorLn(
							"User workspaces need to be exactly one, and need to be created",
						)
					}

					if *session.UserWorkspaces[0].WorkspaceId != workspace.UniqueId {
						t.ErrorLn("Workspace id is not same as definition")
					}

					if *session.UserWorkspaces[0].UserId != user.UniqueId {
						t.ErrorLn("user if of created workspace is not same as expected user id")
					}

					if *session.User.WorkspaceId != ROOT_VAR {
						t.ErrorLn("User in the session does not belong to the 'root' workspace. Every user, belongs to root workspace regardless, and can be assigned other workspaces using userWorkspace table")
					}

				}

				return nil
			},
		},
	}

	for _, test := range tests {
		err := test.Function(testContext)
		if err == nil {
			c := color.New(color.FgGreen)
			fmt.Print("\u2713 Test \"")
			c.Print(test.Name)
			fmt.Print("\" Has passed successfully")
		}
		fmt.Println("")

	}
}

func TestRunner(ctx *TestContext, tests []Test) {
	for _, test := range tests {
		err := test.Function(ctx)
		if err == nil {
			c := color.New(color.FgGreen)
			fmt.Print("\u2713 Test \"")
			c.Print(test.Name)
			fmt.Print("\" Has passed successfully")
		}
		fmt.Println("")

	}
}

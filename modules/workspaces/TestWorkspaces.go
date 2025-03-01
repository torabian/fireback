package workspaces

var WorkspaceCreationTests = []Test{
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
			parent, err := WorkspaceActionCreate(&WorkspaceEntity{
				Name:   parentWsName,
				TypeId: NewString(ROOT_VAR),
			}, t.F)
			if err != nil {
				t.ErrorLn("First workspace did not create to begin with:", err)
				return err
			}

			child, err2 := WorkspaceActionCreate(&WorkspaceEntity{
				Name:     childWsName,
				ParentId: NewString(parent.UniqueId),
				TypeId:   NewString(ROOT_VAR),
			}, t.F)

			if err2 != nil {
				t.ErrorLn("Second workspace did not created")
				return err
			}

			if child.ParentId.String != parent.UniqueId {
				t.ErrorLn(
					"Expected the parent id of child to be unique id, but it's not",
					"parent id:",
					parent.UniqueId,
					"child id",
					child.UniqueId,
					"child parent id",
					child.ParentId.String,
				)
			}
			return nil
		},
	},
	{
		Name: "It should be able to find the created workspaces, and they must be representing each other",
		Function: func(t *TestContext) error {

			if _, _, err := WorkspaceActionQuery(t.F); err != nil {
				t.ErrorLn("Workspaces could not be queried from database", err.Error())
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
				t.F.WorkspaceId = ROOT_VAR
				t.F.UserHas = []string{ROOT_ALL_ACCESS}
				if workspacedb, err := WorkspaceActionGetOne(t.F); err != nil {
					t.ErrorLn("Error on finding created workspace in database", err)
					return err
				} else {
					if workspacedb.UniqueId != workspacedb.WorkspaceId.String {
						t.ErrorLn("Workspace id must be equal to it's unique id in database, but its not", err, workspacedb)
						return err
					}

					if workspace.UniqueId != workspacedb.WorkspaceId.String {
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

				if session.Token == "" || len(session.Token) < 10 {
					t.ErrorLn(
						"Token from the session has length less than 10, it's suspisous",
					)
				}

				if len(session.UserWorkspaces) < 1 {
					t.ErrorLn(
						"User workspaces need to be exactly one, and need to be created",
					)
				}

				if session.UserWorkspaces[0].WorkspaceId.String != workspace.UniqueId {
					t.ErrorLn("Workspace id is not same as definition")
				}

				if session.UserWorkspaces[0].UserId.String != user.UniqueId {
					t.ErrorLn("user if of created workspace is not same as expected user id")
				}

				if session.User.WorkspaceId.String != ROOT_VAR {
					t.ErrorLn("User in the session does not belong to the 'root' workspace. Every user, belongs to root workspace regardless, and can be assigned other workspaces using userWorkspace table")
				}

			}

			return nil
		},
	},
}

package workspaces

func (x *UserEntity) FullName() string {
	if x.Person == nil {
		return ""
	}

	full := ""

	if x.Person.FirstName != "" {
		full += x.Person.FirstName
	}

	if x.Person.LastName != "" {
		full += " " + x.Person.LastName
	}

	return full

}

func init() {

	UserActions.SeederInit = func() *UserEntity {

		return &UserEntity{
			Person: &PersonEntity{
				UniqueId:  UUID(),
				FirstName: "Ali",
				LastName:  "Torabi",
			},
		}
	}
	// Tokens are related to users, so let's move them there.
	UserCliCommands = append(
		UserCliCommands,
		TokenCliFn(),
		CreateRootUser,
		AcceptInviteActionCmd,
		UserInvitationsActionCmd,
	)
}

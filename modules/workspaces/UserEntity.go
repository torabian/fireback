package workspaces

func (x *UserEntity) FullName() string {
	if x.Person == nil {
		return ""
	}

	full := ""

	if x.Person.FirstName != nil {
		full += *x.Person.FirstName
	}

	if x.Person.LastName != nil {
		full += " " + *x.Person.LastName
	}

	return full

}

func init() {

	UserActions.SeederInit = func() *UserEntity {
		ali := "Ali"
		torabi := "Torabi"

		return &UserEntity{
			Person: &PersonEntity{
				UniqueId:  UUID(),
				FirstName: &ali,
				LastName:  &torabi,
			},
		}
	}
	// Tokens are related to users, so let's move them there.
	UserCliCommands = append(
		UserCliCommands,
		TokenCliFn(),
		CreateRootUser,
	)
}

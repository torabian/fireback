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

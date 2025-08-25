package mcore

func (x *Module3) PublicName() string {
	return ToUpper(x.Name)
}

func (x *Module3Field) PublicName() string {
	return ToUpper(x.Name)
}

func (x *Module3FieldMatch) PublicName() string {
	if x.Dto == nil {
		return ""
	}

	return ToUpper(*x.Dto) + "Dto"
}

func (x *Module3Field) PrivateName() string {
	return x.Name
}

func (x *Module3Action) Upper() string {
	return ToUpper(x.Name)
}

package workspaces

func prependScript(name string) string {
	return `

	func Cast` + ToUpper(name) + `FieldsFromJson(schema *workspaces.JSON) ([]*` + ToUpper(name) + `Fields, *workspaces.IError) {
		form := JSONSchema{}
	
		if err := form.fromJson(schema); err != nil {
			return nil, workspaces.GormErrorToIError(err)
		}
	
		fields := []*` + ToUpper(name) + `Fields{}
		for key, field := range flattenFields("", form.Properties) {
			key := key
			field := field
			fields = append(fields, &` + ToUpper(name) + `Fields{
				Type: &field.Type,
				Name: &key,
			})
		}
	
		return fields, nil
	}
	`
}

func prependCreateScript(name string) string {
	return `
	if fields, err := Cast` + ToUpper(name) + `FieldsFromJson(dto.JsonSchema); err != nil {
		return nil, err
	} else {
		dto.Fields = fields
	}
	`
}

func prependUpdateScript(name string) string {
	return `
	if fields2, err := Cast` + ToUpper(name) + `FieldsFromJson(fields.JsonSchema); err != nil {
		return nil, err
	} else {
		fields.Fields = fields2
	}
	`
}

func EavMacro(macro Module2Macro, x *Module2) {
	key := macro.Name

	form := Module2Entity{
		Name:                key,
		PrependScript:       prependScript(key),
		PrependCreateScript: prependCreateScript(key),
		PrependUpdateScript: prependUpdateScript(key),
		Fields: []*Module2Field{
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "description",
				Type: "string",
			},
			{
				Name: "uiSchema",
				Type: "json",
			},
			{
				Name: "jsonSchema",
				Type: "json",
			},
			{
				Name: "fields",
				Type: "array",
				Fields: []*Module2Field{
					{
						Name:   key,
						Type:   "one",
						Target: ToUpper(key) + "Entity",
					},
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "type",
						Type: "string",
					},
				},
			},
		},
	}

	submissionFields := []*Module2Field{
		{
			Name:     key,
			Type:     "one",
			Target:   ToUpper(key) + "Entity",
			Validate: "required",
		},
		{
			Name: "Values",
			Type: "array",
			Fields: []*Module2Field{
				{
					Name:   key + "Field",
					Type:   "one",
					Target: ToUpper(key) + "Fields",
				},
				{
					Name: "value",
					Type: "text",
				},
			},
		},
	}

	submissionFields = append(submissionFields, macro.Fields...)

	formSubmission := Module2Entity{
		Name:   key + "Submission",
		Fields: submissionFields,
	}

	x.Entities = append(x.Entities, form, formSubmission)
}

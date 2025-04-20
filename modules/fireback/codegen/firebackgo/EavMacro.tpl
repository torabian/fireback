func Cast{{ .Key }}FieldsFromJson(schema *{{ .wsprefix }}JSON) ([]*{{ .Key }}Fields, *{{ .wsprefix }}IError) {
    form := {{ .wsprefix }}JSONSchema{}

    if err := form.FromJson(schema); err != nil {
        return nil, {{ .wsprefix }}GormErrorToIError(err)
    }

    fields := []*{{ .Key }}Fields{}
    for key, field := range {{ .wsprefix }}FlattenFields("", form.Properties) {
        key := key
        field := field
        fields = append(fields, &{{ .Key }}Fields{
            Type: &field.Type,
            Name: &key,
        })
    }

    return fields, nil
}


func ComputeValueFromInterface(row *{{ .Key }}SubmissionValues, value interface{}) {

	switch value := value.(type) {
	case int64:

		row.ValueInt64 = &value
	case float64:
		row.ValueFloat64 = &value
	case string:
		row.ValueString = &value
	case bool:
		row.ValueBoolean = &value
	}

}

func FindFieldId(fields []*{{ .Key }}Fields, fieldName string) string {
	for _, field := range fields {
		if *field.Name == fieldName {
			return field.UniqueId
		}
	}
	return ""
}

func SubmergeDataObjectWithValuesArray(
	data *{{ .wsprefix }}JSON,
	fields []*{{ .Key }}Fields,
) []*{{ .Key }}SubmissionValues {

	items := []*{{ .Key }}SubmissionValues{}

    if (data == nil ) {
        return items
    }

	var data3 map[string]interface{}
	// var json = jsoniter.ConfigCompatibleWithStandardLibrary
	// json.UnmarshalFromString(data.String(), &data3)
	json.Unmarshal([]byte(data.String()), &data3)

	for k, v := range {{ .wsprefix }}FlattenData(data3, "") {

		fieldUniqueId := FindFieldId(fields, k)
		if fieldUniqueId == "" {
			continue
		}

		row := &{{ .Key }}SubmissionValues{
			{{ .Key }}FieldId: &fieldUniqueId,
		}
		ComputeValueFromInterface(row, v)

		items = append(items, row)
	}

	return items
}

func {{ .Key}}SubmissionCastFieldsToEavAndValidate(dto *{{ .Key }}SubmissionEntity, query {{ .wsprefix }}QueryDSL) *{{ .wsprefix }}IError {
    if dto == nil || dto.ProductId == nil {
        return nil
    }
	id := query.UniqueId
	query.UniqueId = *dto.{{ .Key }}Id
	form, err := {{ .Key }}ActionGetOne(query)
	if err != nil {
		return err
	}

	query.UniqueId = id

	dto.Values = SubmergeDataObjectWithValuesArray(dto.Data, form.Fields)

	if err0 := {{ .wsprefix }}ValidateEavSchema(form.JsonSchema, dto.Data); err0 != nil {
		return err0
	}

	return nil
}
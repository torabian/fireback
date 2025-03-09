package workspaces

func init() {
	// Override the implementation with our actual code.

	UserPassportsActionImp = UserPassportsAction
}
func UserPassportsAction(q QueryDSL) ([]*UserPassportsActionResDto, *QueryResultMeta, *IError) {

	passports := []PassportEntity{}
	err := GetRef(q).Where(PassportEntity{UserId: NewString(q.UserId)}).Find(&passports).Error
	if err != nil {
		return nil, nil, CastToIError(err)
	}

	result := []*UserPassportsActionResDto{}
	for _, item := range passports {
		result = append(result, &UserPassportsActionResDto{
			Value:         item.Value,
			Type:          item.Type,
			UniqueId:      item.UniqueId,
			TotpConfirmed: item.TotpConfirmed,
		})
	}

	return result, nil, nil
}

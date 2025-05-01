package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.

	UserPassportsActionImp = UserPassportsAction
}
func UserPassportsAction(q fireback.QueryDSL) ([]*UserPassportsActionResDto, *fireback.QueryResultMeta, *fireback.IError) {

	passports := []PassportEntity{}
	err := fireback.GetRef(q).Debug().Where(PassportEntity{UserId: fireback.NewString(q.UserId)}).Find(&passports).Error
	if err != nil {
		return nil, nil, fireback.CastToIError(err)
	}

	result := []*UserPassportsActionResDto{}
	for _, item := range passports {
		result = append(result, &UserPassportsActionResDto{
			Value:         item.Value,
			Type:          item.Type,
			UniqueId:      item.UniqueId,
			TotpConfirmed: item.TotpConfirmed.Bool,
		})
	}

	return result, nil, nil
}

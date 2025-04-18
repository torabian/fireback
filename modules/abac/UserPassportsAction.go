package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.

	UserPassportsActionImp = UserPassportsAction
}
func UserPassportsAction(q workspaces.QueryDSL) ([]*UserPassportsActionResDto, *workspaces.QueryResultMeta, *workspaces.IError) {

	passports := []PassportEntity{}
	err := workspaces.GetRef(q).Debug().Where(PassportEntity{UserId: workspaces.NewString(q.UserId)}).Find(&passports).Error
	if err != nil {
		return nil, nil, workspaces.CastToIError(err)
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

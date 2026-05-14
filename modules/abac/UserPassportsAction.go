package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	UserPassportsImpl = UserPassportsAction
}

func UserPassportsAction(c UserPassportsActionRequest, q fireback.QueryDSL) (*UserPassportsActionResponse, error) {

	passports := []PassportEntity{}
	err := fireback.GetRef(q).Where(PassportEntity{UserId: emigo.NullableOf(q.UserId)}).Find(&passports).Error
	if err != nil {
		return nil, fireback.CastToIError(err)
	}

	result := []UserPassportsActionRes{}
	for _, item := range passports {
		result = append(result, UserPassportsActionRes{
			Value:         item.Value,
			Type:          item.Type,
			UniqueId:      item.UniqueId,
			TotpConfirmed: item.TotpConfirmed.OrDefault(false),
		})
	}

	return &UserPassportsActionResponse{
		Payload: fireback.GResponseQuery(result, nil, &q),
	}, nil
}

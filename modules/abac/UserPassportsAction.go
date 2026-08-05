package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func UserPassportsAction(c UserPassportsActionRequest) (*UserPassportsActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

	passports := []PassportEntity{}
	err2 := fireback.GetRef(q).Where(PassportEntity{UserId: emigo.NullableOf(q.UserId)}).Find(&passports).Error
	if err2 != nil {
		return nil, fireback.CastToIError(err2)
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

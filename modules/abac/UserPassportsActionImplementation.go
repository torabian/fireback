package abac

import (
	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func UserPassportsAction(c abacdefs.UserPassportsActionRequest) (*abacdefs.UserPassportsActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
	})
	if err != nil {
		return nil, err
	}
	q := *query

	passports := []abacdefs.PassportEntity{}
	err2 := fireback.GetRef(q).Where(abacdefs.PassportEntity{UserId: emigo.NullableOf(q.UserId)}).Find(&passports).Error
	if err2 != nil {
		return nil, fireback.CastToIError(err2)
	}

	result := []abacdefs.UserPassportsActionRes{}
	for _, item := range passports {
		result = append(result, abacdefs.UserPassportsActionRes{
			Value:         item.Value,
			Type:          item.Type,
			UniqueId:      item.UniqueId,
			TotpConfirmed: item.TotpConfirmed.OrDefault(false),
		})
	}

	return &abacdefs.UserPassportsActionResponse{
		Payload: fireback.GResponseQuery(result, nil, &q),
	}, nil
}

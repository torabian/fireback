package abac

import (
	"encoding/json"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func SendEmailWithProviderAction(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func CastEmailProviderEntityFromCli(c emigo.CliCastable) EmailProviderEntity {
	var result EmailProviderEntity
	json.Unmarshal([]byte(c.String("email-provider")), &result)

	return result
}

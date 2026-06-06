package abac

import (
	"encoding/json"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	SendEmailWithProviderImpl = func(c SendEmailWithProviderActionRequest, query fireback.QueryDSL) (*SendEmailWithProviderActionResponse, error) {
		return nil, nil
	}
}

func CastEmailProviderEntityFromCli(c emigo.CliCastable) EmailProviderEntity {
	var result EmailProviderEntity
	json.Unmarshal([]byte(c.String("email-provider")), &result)

	return result
}

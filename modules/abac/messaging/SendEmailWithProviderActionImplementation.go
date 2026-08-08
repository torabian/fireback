package messaging

import (
	"github.com/torabian/fireback/modules/fireback"
)

func SendEmailWithProviderAction(c SendEmailWithProviderActionRequest) (*SendEmailWithProviderActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

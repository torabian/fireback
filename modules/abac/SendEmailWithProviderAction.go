package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	SendEmailWithProviderImpl = func(c SendEmailWithProviderActionRequest, query fireback.QueryDSL) (*SendEmailWithProviderActionResponse, error) {
		return nil, nil
	}
}

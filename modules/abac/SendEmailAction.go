package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	SendEmailImpl = func(c SendEmailActionRequest, query fireback.QueryDSL) (*SendEmailActionResponse, error) {
		return nil, nil
	}
}

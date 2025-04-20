package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	SendEmailWithProviderActionImp = SendEmailWithProviderAction
}
func SendEmailWithProviderAction(
	req *SendEmailWithProviderActionReqDto,
	q fireback.QueryDSL) (*SendEmailWithProviderActionResDto,
	*fireback.IError,
) {
	// Implement the logic here.
	return nil, nil
}

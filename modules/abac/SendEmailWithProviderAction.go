package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	SendEmailWithProviderActionImp = SendEmailWithProviderAction
}
func SendEmailWithProviderAction(
	req *SendEmailWithProviderActionReqDto,
	q workspaces.QueryDSL) (*SendEmailWithProviderActionResDto,
	*workspaces.IError,
) {
	// Implement the logic here.
	return nil, nil
}

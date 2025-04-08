package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	SendEmailActionImp = SendEmailAction
}
func SendEmailAction(
	req *SendEmailActionReqDto,
	q workspaces.QueryDSL) (*SendEmailActionResDto,
	*workspaces.IError,
) {
	// Implement the logic here.
	return nil, nil
}

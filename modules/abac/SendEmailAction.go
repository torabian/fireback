package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	SendEmailActionImp = SendEmailAction
}
func SendEmailAction(
	req *SendEmailActionReqDto,
	q fireback.QueryDSL) (*SendEmailActionResDto,
	*fireback.IError,
) {
	// Implement the logic here.
	return nil, nil
}

package workspaces

func init() {
	// Override the implementation with our actual code.
	SendEmailActionImp = SendEmailAction
}
func SendEmailAction(
	req *SendEmailActionReqDto,
	q QueryDSL) (*SendEmailActionResDto,
	*IError,
) {
	// Implement the logic here.
	return nil, nil
}

package workspaces

func init() {
	// Override the implementation with our actual code.
	SendEmailWithProviderActionImp = SendEmailWithProviderAction
}
func SendEmailWithProviderAction(
	req *SendEmailWithProviderActionReqDto,
	q QueryDSL) (*SendEmailWithProviderActionResDto,
	*IError,
) {
	// Implement the logic here.
	return nil, nil
}

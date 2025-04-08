package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	GsmSendSmsWithProviderActionImp = GsmSendSmsWithProviderAction
}

func GsmSendSmsWithProviderAction(req *GsmSendSmsWithProviderActionReqDto, q workspaces.QueryDSL) (*GsmSendSmsWithProviderActionResDto, *workspaces.IError) {

	if err := GsmSendSmsWithProviderActionReqValidator(req); err != nil {
		return nil, err
	}

	return GsmSendSMS(req.GsmProviderId.String, req.Body, []string{req.ToNumber})
}

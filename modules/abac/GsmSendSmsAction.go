package abac

import "github.com/torabian/fireback/modules/workspaces"

func init() {
	// Override the implementation with our actual code.
	GsmSendSmsActionImp = GsmSendSmsAction
}

func GsmSendSmsAction(req *GsmSendSmsActionReqDto, q workspaces.QueryDSL) (*GsmSendSmsActionResDto, *workspaces.IError) {

	if err := GsmSendSmsActionReqValidator(req); err != nil {
		return nil, err
	}
	if res, err := GsmSendSMSUsingNotificationConfig(req.Body, []string{req.ToNumber}); err != nil {
		return nil, err
	} else {
		return &GsmSendSmsActionResDto{
			QueueId: res.QueueId,
		}, nil
	}
}

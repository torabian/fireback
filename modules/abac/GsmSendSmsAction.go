package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	GsmSendSmsActionImp = GsmSendSmsAction
}

func GsmSendSmsAction(req *GsmSendSmsActionReqDto, q fireback.QueryDSL) (*GsmSendSmsActionResDto, *fireback.IError) {

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

package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	GsmSendSmsImpl = func(c GsmSendSmsActionRequest, query fireback.QueryDSL) (*GsmSendSmsActionResponse, error) {
		req := c.Body

		if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
			return nil, err
		}
		if res, err := GsmSendSMSUsingNotificationConfig(req.Body, []string{req.ToNumber}); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsActionResponse{
				Payload: GsmSendSmsActionRes{
					QueueId: res.QueueId,
				},
			}, nil
		}
	}
}

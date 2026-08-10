package abac

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func GsmSendSmsAction(c abacdefs.GsmSendSmsActionRequest) (*abacdefs.GsmSendSmsActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	req := c.Body

	if err2 := fireback.CommonStructValidatorPointer(&req, false); err2 != nil {
		return nil, err2
	}
	if res, err2 := GsmSendSMSUsingNotificationConfig(req.Body, []string{req.ToNumber}); err2 != nil {
		return nil, err2
	} else {
		return &abacdefs.GsmSendSmsActionResponse{
			Payload: abacdefs.GsmSendSmsActionRes{
				QueueId: res.QueueId,
			},
		}, nil
	}
}

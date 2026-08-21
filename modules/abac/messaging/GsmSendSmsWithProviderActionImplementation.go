package messaging

import (
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func GsmSendSmsWithProviderAction(c messagingdefs.GsmSendSmsWithProviderActionRequest) (*messagingdefs.GsmSendSmsWithProviderActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	req := c.Body
	if err2 := fireback.CommonStructValidatorPointer(&req, false); err2 != nil {
		return nil, err2
	}

	if res, err2 := GsmSendSMS(req.GsmProviderId, req.Body, []string{req.ToNumber}); err2 != nil {
		return nil, err2
	} else {
		return &messagingdefs.GsmSendSmsWithProviderActionResponse{
			Payload: res,
		}, nil
	}
}

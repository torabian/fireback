package messaging

import "github.com/torabian/fireback/modules/fireback"

func GsmSendSmsWithProviderAction(c GsmSendSmsWithProviderActionRequest) (*GsmSendSmsWithProviderActionResponse, error) {
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
		return &GsmSendSmsWithProviderActionResponse{
			Payload: res,
		}, nil
	}
}

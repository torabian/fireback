package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	GsmSendSmsWithProviderImpl = GsmSendSmsWithProviderActionfunc
}

func GsmSendSmsWithProviderActionfunc(c GsmSendSmsWithProviderActionRequest, query fireback.QueryDSL) (*GsmSendSmsWithProviderActionResponse, error) {

	req := c.Body
	if err := fireback.CommonStructValidatorPointer(&req, false); err != nil {
		return nil, err
	}

	if res, err := GsmSendSMS(req.GsmProviderId, req.Body, []string{req.ToNumber}); err != nil {
		return nil, err
	} else {
		return &GsmSendSmsWithProviderActionResponse{
			Payload: res,
		}, nil
	}

}

package messaging

import (
	"github.com/torabian/fireback/modules/fireback"
)

func SendEmailAction(c SendEmailActionRequest) (*SendEmailActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	req := c.Body
	if err2 := fireback.CommonStructValidatorPointer(&req, false); err2 != nil {
		return nil, err2
	}

	provider, err2 := EmailProviderActions.GetOne(fireback.QueryDSL{
		UniqueId: req.ProviderId,
	})

	if fireback.IsErr(err2) {
		return nil, err2
	}

	if err3 := SendMail(EmailMessageContent{
		FromName:  "Test",
		FromEmail: "test@test.com",
		ToName:    req.ToAddress,
		ToEmail:   req.ToAddress,
		Subject:   "Testing email",
		Content:   req.Body,
	}, provider); err3 != nil {
		return nil, err3
	}

	return &SendEmailActionResponse{
		Payload: &SendEmailActionRes{QueueId: fireback.UUID()},
	}, nil
}

package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

func SendEmailAction(c SendEmailActionRequest) (*SendEmailActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	provider, err2 := EmailProviderActions.GetOne(fireback.QueryDSL{
		UniqueId: c.Body.ProviderId,
	})

	if fireback.IsErr(err2) {
		return nil, err2
	}

	if err3 := SendMail(EmailMessageContent{
		FromName:  "Test",
		FromEmail: "test@test.com",
		ToName:    "Test reciever",
		ToEmail:   "test@test.com",
		Subject:   "Testing email",
		Content:   "Hello :)",
	}, provider); err3 != nil {
		return nil, err3
	}

	return nil, nil
}

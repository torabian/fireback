package abac

import (
	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	SendEmailImpl = func(c SendEmailActionRequest, query fireback.QueryDSL) (*SendEmailActionResponse, error) {
		provider, err := EmailProviderActions.GetOne(fireback.QueryDSL{
			UniqueId: c.Body.ProviderId,
		})

		if fireback.IsErr(err) {
			return nil, err
		}

		if err := SendMail(EmailMessageContent{
			FromName:  "Test",
			FromEmail: "test@test.com",
			ToName:    "Test reciever",
			ToEmail:   "test@test.com",
			Subject:   "Testing email",
			Content:   "Hello :)",
		}, provider); err != nil {
			return nil, err
		}

		return nil, nil
	}
}

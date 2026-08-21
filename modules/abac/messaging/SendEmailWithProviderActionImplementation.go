package messaging

import (
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func SendEmailWithProviderAction(c messagingdefs.SendEmailWithProviderActionRequest) (*messagingdefs.SendEmailWithProviderActionResponse, error) {
	_, err := fireback.ResolveActionContext(c, nil)
	if err != nil {
		return nil, err
	}

	req := c.Body

	// Only validate ToAddress/Body here, not the whole request struct: req.EmailProvider
	// is an emigo.One[EmailProviderEntity], used purely as an id selector (see the
	// req.EmailProvider.Item.UniqueId lookup below) - but go-playground's validator
	// recurses into nested structs by default, so validating &req directly would also
	// enforce EmailProviderEntity's OWN validate:"required" tags (e.g. "type") against
	// whatever zero-valued entity sits in .Item, rejecting a perfectly valid
	// {"emailProvider": {"uniqueId": "..."}} selector with a bogus "emailProvider.item.type
	// is required" error.
	if err2 := fireback.CommonStructValidatorPointer(&struct {
		ToAddress string `validate:"required"`
		Body      string `validate:"required"`
	}{ToAddress: req.ToAddress, Body: req.Body}, false); err2 != nil {
		return nil, err2
	}

	provider, err2 := EmailProviderActions.GetOne(fireback.QueryDSL{
		UniqueId: req.EmailProvider.Item.UniqueId,
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

	return &messagingdefs.SendEmailWithProviderActionResponse{
		Payload: &messagingdefs.SendEmailWithProviderActionRes{QueueId: fireback.UUID()},
	}, nil
}

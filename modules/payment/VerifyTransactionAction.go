package payment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	VerifyTransactionActionImp = VerifyTransactionAction
}

func VerifyTransactionAction(req *VerifyTransactionActionReqDto, q fireback.QueryDSL) (string, *fireback.IError) {
	// We are not having that much transaction yet,
	// so we are querying configuration for each transaction to make sure we have latest config.
	// take out this part and sync it on change maybe.
	config, err := PaymentParameterActions.GetByWorkspace(fireback.QueryDSL{WorkspaceId: fireback.ROOT_VAR})
	if err != nil {
		return "", err
	}

	merchantId := config.PosId
	merchantIdInt64, err2 := strconv.ParseInt(merchantId, 10, 64)
	if err2 != nil {
		return "", fireback.CastToIError(err)
	}

	body := &VerifyTransactionRemoteBody{
		SessionId:  req.SessionId,
		MerchantId: merchantIdInt64,
		PosId:      merchantIdInt64,
		Amount:     req.Amount,
		OrderId:    req.OrderId,
		Currency:   config.ComputedCurrency(),
	}

	body.SignWithCrc(config.Crc)

	options := &fireback.RemoteRequestOptions{
		Url: config.VerifyApiUrl,
	}

	res, _, err := PaymentRemotes.VerifyTransaction(body, http.Header{
		"Authorization": []string{config.FullBasicAuthHeader()},
	}, options)

	fmt.Println(res)
	if err != nil || res.Error != "" {
		return "", err
	}

	return "Verified", nil
}

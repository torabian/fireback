package payment

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

func init() {
	// Override the implementation with our actual code.
	RegisterTransactionActionImp = RegisterTransactionAction
}

func RegisterTransactionAction(req *RegisterTransactionActionReqDto, q fireback.QueryDSL) (string, *fireback.IError) {
	// We are not having that much transaction yet,
	// so we are querying configuration for each transaction to make sure we have latest config.
	// take out this part and sync it on change maybe.
	config, err := PaymentParameterActions.GetByWorkspace(fireback.QueryDSL{WorkspaceId: fireback.ROOT_VAR})
	if err != nil {
		return "", err
	}

	// Let's validate the configuration as well, and kill process earlier.
	if err2 := ValidatePaymentParameters(config); err2 != nil {
		return "", err2
	}

	merchantId := config.PosId
	merchantIdInt64, err2 := strconv.ParseInt(merchantId, 10, 64)
	if err2 != nil {
		return "", fireback.CastToIError(err)
	}

	sessionId := fireback.UUID()

	body := &RegisterTransactionRemoteBody{
		MerchantId:  merchantIdInt64,
		PosId:       merchantIdInt64,
		Email:       req.Email,
		Amount:      req.Amount,
		Description: req.Description,
		Country:     config.Country,
		UrlReturn:   config.UrlReturn,
		UrlStatus:   config.UrlStatus,
		SessionId:   sessionId,
		Currency:    config.ComputedCurrency(),
	}

	sign := GenerateSignForRegister(body, config.Crc)
	body.Sign = sign
	options := &fireback.RemoteRequestOptions{
		Url: config.RegisterApiUrl,
	}

	if config.RegisterApiUrl != "" {
		options.Url = config.RegisterApiUrl
	}

	res, _, err := PaymentRemotes.RegisterTransaction(body, http.Header{
		"Authorization": []string{config.FullBasicAuthHeader()},
	}, options)

	if err != nil || res.Error != "" {
		return "", err
	} else {
		if res != nil && res.Data != nil && res.Data.Token != "" {
			return strings.ReplaceAll(config.PaymentPageUrl, "{TOKEN}", res.Data.Token), nil
		}
		return "", nil
	}
}

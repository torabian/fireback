package payment

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/torabian/fireback/modules/fireback"
)

// The code here is for registering a payment as first step in przelewy24.pl

func init() {
	RegisterTransactionActionImp = RegisterTransactionAction
	VerifyTransactionActionImp = VerifyTransactionAction
}

// We need to check if the payment provider is configurated correctly before proceed to the payment
func ValidatePaymentParameters(config *PaymentParameterEntity) *fireback.IError {
	if config == nil {
		return fireback.Create401Error(&PaymentMessages.PaymentProviderFullyMissing, nil)
	}

	items := []*fireback.IErrorItem{}
	if config.Crc == "" {
		items = append(items, &fireback.IErrorItem{
			Message:  &PaymentMessages.PaymentProviderMissingCrc,
			Location: "crc",
		})
	}

	if config.PosId == "" {
		items = append(items, &fireback.IErrorItem{
			Message:  &PaymentMessages.PaymentProviderMissingPosId,
			Location: "posId",
		})
	}

	if config.SecretId == "" {
		items = append(items, &fireback.IErrorItem{
			Message:  &PaymentMessages.PaymentProviderMissingSecretId,
			Location: "secretId",
		})
	}

	if config.UrlReturn == "" {
		items = append(items, &fireback.IErrorItem{
			Message:  &PaymentMessages.PaymentProviderMissingUrlReturn,
			Location: "urlReturn",
		})
	}

	if config.UrlStatus == "" {
		items = append(items, &fireback.IErrorItem{
			Message:  &PaymentMessages.PaymentProviderMissingUrlStatus,
			Location: "urlStatus",
		})
	}

	if config.MerchantId == "" {
		items = append(items, &fireback.IErrorItem{
			Message:  &PaymentMessages.PaymentProviderMissingMerchantId,
			Location: "merchantId",
		})
	}

	if len(items) == 0 {
		return nil
	}

	return fireback.Create401ErrorWithItems(&PaymentMessages.PaymentProviderMissingConfiguration, items)
}

func (x *VerifyTransactionRemoteBody) SignWithCrc(crc string) {
	data := VerifyCrcData{
		SessionID: x.SessionId,
		OrderId:   x.OrderId,
		Amount:    x.Amount,
		Currency:  x.Currency,
		CRC:       crc,
	}

	s := SignSha512(data)
	x.Sign = s
}

func (x *PaymentParameterEntity) ComputedCurrency() string {

	// Currency has been hard coded in favor of a client.
	currency := "PLN"
	if x.Currency != "" {
		currency = x.Currency
	}

	return currency
}

func (x *PaymentParameterEntity) FullBasicAuthHeader() string {
	data := x.PosId + ":" + x.SecretId
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(data))
}

type VerifyCrcData struct {
	SessionID string `json:"sessionId"`
	OrderId   int64  `json:"merchantId"`
	Amount    int64  `json:"amount"` // Adjust type if needed
	Currency  string `json:"currency"`
	CRC       string `json:"crc"`
}

type RegisterCrcData struct {
	SessionID  string  `json:"sessionId"`
	MerchantID float64 `json:"merchantId"`
	Amount     float64 `json:"amount"` // Adjust type if needed
	Currency   string  `json:"currency"`
	CRC        string  `json:"crc"`
}

func SignSha512(data any) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Default().Println("Json failed on merchant data unfortunately for security reasons cannot log that")
	}

	hash := sha512.New384()

	hash.Write(jsonData)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GenerateSignForRegister(body *RegisterTransactionRemoteBody, crc string) string {
	data := RegisterCrcData{
		SessionID:  body.SessionId,
		MerchantID: float64(body.MerchantId),
		Amount:     float64(body.Amount),
		Currency:   body.Currency,
		CRC:        crc,
	}

	return SignSha512(data)
}

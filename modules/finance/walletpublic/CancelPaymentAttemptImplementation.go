package walletpublic

import (
	"errors"
	"net/http"

	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
	"github.com/torabian/fireback/modules/fireback"
	"gorm.io/gorm"
)

// CancelPaymentAttemptAction cancels one of the caller's own payment attempts, only
// while it's still "pending" - once a gateway has moved it past that (requires_action or
// a terminal state) cancelling here could race the gateway's own webhook.
func CancelPaymentAttemptAction(c walletpublicdefs.CancelPaymentAttemptActionRequest) (*walletpublicdefs.CancelPaymentAttemptActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.CancelPaymentAttemptActionResponse{StatusCode: status, Payload: payload}, nil
	}

	tx := fireback.GetDbRef()
	var attempt walletdefs.WalletPaymentAttemptEntity
	if err := tx.Preload("Wallet").Preload("Gateway").First(&attempt, "unique_id = ?", c.Body.AttemptId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status, payload := emiIErrorPayload(attemptNotFoundError(), query)
			return &walletpublicdefs.CancelPaymentAttemptActionResponse{StatusCode: status, Payload: payload}, nil
		}
		return nil, err
	}
	if attempt.Wallet == nil || !ownsWallet(attempt.Wallet, query.UserId) {
		status, payload := emiIErrorPayload(attemptNotFoundError(), query)
		return &walletpublicdefs.CancelPaymentAttemptActionResponse{StatusCode: status, Payload: payload}, nil
	}
	if attempt.Status != "pending" {
		status, payload := emiIErrorPayload(attemptNotCancellableError(attempt.Status), query)
		return &walletpublicdefs.CancelPaymentAttemptActionResponse{StatusCode: status, Payload: payload}, nil
	}

	if err := tx.Model(&attempt).Update("status", "cancelled").Error; err != nil {
		return nil, err
	}
	gatewayCode := ""
	if attempt.Gateway != nil {
		gatewayCode = attempt.Gateway.Code
	}
	return &walletpublicdefs.CancelPaymentAttemptActionResponse{
		Payload: fireback.GResponseSingleItem(walletPaymentAttemptViewDtoFromEntity(&attempt, gatewayCode)),
	}, nil
}

func attemptNotFoundError() *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusNotFound,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.attemptNotFound",
			"en": "Payment attempt not found.",
			"fa": "تلاش پرداخت یافت نشد.",
			"ru": "Попытка оплаты не найдена.",
			"pl": "Nie znaleziono próby płatności.",
		},
	}
}

func attemptNotCancellableError(status string) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusConflict,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.attemptNotCancellable",
			"en": "Only a pending payment attempt can be cancelled.",
			"fa": "فقط تلاش پرداخت در حالت در انتظار قابل لغو است.",
			"ru": "Отменить можно только ожидающую попытку оплаты.",
			"pl": "Anulować można tylko oczekującą próbę płatności.",
		},
		MessageParams: map[string]any{"status": status},
	}
}

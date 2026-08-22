package walletpublic

import (
	"context"
	"net/http"
	"time"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/complexes"
	wallet "github.com/torabian/fireback/modules/finance/wallet"
	walletdefs "github.com/torabian/fireback/modules/finance/wallet/defs"
	walletpublicdefs "github.com/torabian/fireback/modules/finance/walletpublic/defs"
)

// TopupAction starts a topup: creates a pending walletPaymentAttempt and asks the
// gateway's GatewayAdapter (see modules/wallet/GatewayAdapter.go) to initiate payment.
// idempotencyKey makes retrying a timed-out call safe - a retry with the same key returns
// the original attempt instead of creating a second one / calling the gateway twice.
func TopupAction(c walletpublicdefs.TopupActionRequest) (*walletpublicdefs.TopupActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ResolveStrategy: fireback.ResolveStrategyUser})
	if err != nil {
		return nil, err
	}
	if ierr := fireback.CommonStructValidatorPointer(&c.Body, false); ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.TopupActionResponse{StatusCode: status, Payload: payload}, nil
	}

	tx := fireback.GetDbRef()
	w, ierr := resolveOwnedWallet(tx, c.Body.WalletId, query.UserId)
	if ierr != nil {
		status, payload := emiIErrorPayload(ierr, query)
		return &walletpublicdefs.TopupActionResponse{StatusCode: status, Payload: payload}, nil
	}
	if w.Status != "active" {
		status, payload := emiIErrorPayload(walletNotActiveForTopupError(w.Status), query)
		return &walletpublicdefs.TopupActionResponse{StatusCode: status, Payload: payload}, nil
	}

	var gateway walletdefs.WalletGatewayEntity
	if err := tx.First(&gateway, "code = ? AND is_active = ?", c.Body.GatewayCode, true).Error; err != nil {
		status, payload := emiIErrorPayload(unknownGatewayError(c.Body.GatewayCode), query)
		return &walletpublicdefs.TopupActionResponse{StatusCode: status, Payload: payload}, nil
	}
	adapter, ok := wallet.GetGatewayAdapter(c.Body.GatewayCode)
	if !ok {
		status, payload := emiIErrorPayload(unknownGatewayError(c.Body.GatewayCode), query)
		return &walletpublicdefs.TopupActionResponse{StatusCode: status, Payload: payload}, nil
	}

	// Idempotent retry: a prior call with this key already created an attempt (whether
	// or not InitiatePayment itself finished) - return it instead of starting a second
	// one at the gateway.
	var existing walletdefs.WalletPaymentAttemptEntity
	if err := tx.First(&existing, "idempotency_key = ?", c.Body.IdempotencyKey).Error; err == nil {
		return &walletpublicdefs.TopupActionResponse{Payload: topupResFromAttempt(&existing, gateway.Code)}, nil
	}

	attempt := &walletdefs.WalletPaymentAttemptEntity{
		WalletId:       w.Id,
		GatewayId:      gateway.Id,
		Purpose:        "topup",
		Amount:         c.Body.Amount,
		Currency:       w.Currency,
		Status:         "pending",
		IdempotencyKey: c.Body.IdempotencyKey,
		ReturnUrl:      c.Body.ReturnUrl,
		CreatedAt:      complexes.XDate(time.Now().Format("2006-01-02")),
	}
	if err := tx.Create(attempt).Error; err != nil {
		return nil, err
	}

	result, err := adapter.InitiatePayment(context.Background(), attempt, &gateway)
	if err != nil {
		tx.Model(attempt).Updates(map[string]any{"status": "failed", "failure_reason": err.Error()})
		status, payload := emiIErrorPayload(gatewayInitiationFailedError(err), query)
		return &walletpublicdefs.TopupActionResponse{StatusCode: status, Payload: payload}, nil
	}

	updates := map[string]any{}
	if result.GatewayReference != "" {
		updates["gateway_reference"] = result.GatewayReference
	}
	if result.RawRequest != nil {
		updates["raw_request"] = *complexes.JSONFrom(result.RawRequest)
	}
	if result.RawResponse != nil {
		updates["raw_response"] = *complexes.JSONFrom(result.RawResponse)
	}
	if len(updates) > 0 {
		tx.Model(attempt).Updates(updates)
	}
	if result.GatewayReference != "" {
		attempt.GatewayReference.Set(&result.GatewayReference)
	}

	return &walletpublicdefs.TopupActionResponse{
		StatusCode: http.StatusCreated,
		Payload: &walletpublicdefs.TopupActionRes{
			Attempt: walletpublicdefs.TopupActionResAttempt{
				UniqueId:    attempt.UniqueId,
				Status:      attempt.Status,
				GatewayCode: gateway.Code,
			},
			RedirectUrl:  optionalString(result.RedirectUrl),
			ClientSecret: optionalString(result.ClientSecret),
		},
	}, nil
}

func topupResFromAttempt(a *walletdefs.WalletPaymentAttemptEntity, gatewayCode string) *walletpublicdefs.TopupActionRes {
	return &walletpublicdefs.TopupActionRes{
		Attempt: walletpublicdefs.TopupActionResAttempt{
			UniqueId:    a.UniqueId,
			Status:      a.Status,
			GatewayCode: gatewayCode,
		},
	}
}

func walletNotActiveForTopupError(status string) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusConflict,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.notActive",
			"en": "This wallet is not active and cannot be topped up.",
			"fa": "این کیف پول فعال نیست و امکان شارژ آن وجود ندارد.",
			"ru": "Этот кошелёк неактивен, пополнение недоступно.",
			"pl": "Ten portfel jest nieaktywny i nie można go zasilić.",
		},
		MessageParams: map[string]any{"status": status},
	}
}

func unknownGatewayError(code string) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusBadRequest,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.unknownGateway",
			"en": "This payment gateway is not available.",
			"fa": "این درگاه پرداخت در دسترس نیست.",
			"ru": "Этот платёжный шлюз недоступен.",
			"pl": "Ta bramka płatności jest niedostępna.",
		},
		MessageParams: map[string]any{"gatewayCode": code},
	}
}

func gatewayInitiationFailedError(err error) *fireback.IError {
	return &fireback.IError{
		HttpCode: http.StatusBadGateway,
		Message: fireback.ErrorItem{
			"$":  "wallet.errors.gatewayInitiationFailed",
			"en": "The payment gateway could not start this payment.",
			"fa": "درگاه پرداخت نتوانست این پرداخت را آغاز کند.",
			"ru": "Платёжный шлюз не смог инициировать этот платёж.",
			"pl": "Bramka płatności nie mogła rozpocząć tej płatności.",
		},
		MessageParams: map[string]any{"detail": err.Error()},
	}
}

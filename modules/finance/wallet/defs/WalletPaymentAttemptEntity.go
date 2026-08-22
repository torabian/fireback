package walletdefs

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"github.com/torabian/fireback/modules/fireback/complexes"
	"gorm.io/gorm"
)

// The base class definition for walletPaymentAttemptEntity
type WalletPaymentAttemptEntity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:varchar(100);default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	// The wallet this attempt would credit/debit if it succeeds.
	Wallet *WalletEntity `gorm:"foreignKey:WalletId;references:Id" json:"wallet" validate:"required" yaml:"wallet"`
	// The gateway this attempt is routed through.
	Gateway *WalletGatewayEntity `gorm:"foreignKey:GatewayId;references:Id" json:"gateway" validate:"required" yaml:"gateway"`
	// What this attempt is for.
	Purpose string `json:"purpose" validate:"required,oneof=topup purchase withdrawal" yaml:"purpose"`
	// Requested amount as a minor-units string, in currency.
	Amount string `json:"amount" validate:"required" yaml:"amount"`
	// Currency code for amount - must match the wallet's currency.
	Currency string `json:"currency" validate:"required" yaml:"currency"`
	// Current lifecycle state of this attempt.
	Status string `json:"status" validate:"required,oneof=pending requires_action succeeded failed cancelled expired" yaml:"status"`
	// The gateway's own id for this attempt (e.g. a PaymentIntent id or a transaction hash), once known. Indexed for webhook lookups.
	GatewayReference emigo.Nullable[string] `json:"gatewayReference" yaml:"gatewayReference"`
	// Caller-supplied key making topup-initiation safe to retry without creating duplicate attempts at the gateway.
	IdempotencyKey string `gorm:"unique;not null" json:"idempotencyKey" validate:"required" yaml:"idempotencyKey"`
	// When this attempt was created.
	CreatedAt complexes.XDate `json:"createdAt" yaml:"createdAt"`
	// Human-readable reason, populated when status is "failed".
	FailureReason emigo.Nullable[string] `json:"failureReason" yaml:"failureReason"`
	// The raw request sent to the gateway when initiating this attempt.
	RawRequest complexes.JSON `json:"rawRequest" yaml:"rawRequest"`
	// The raw response/init payload received back from the gateway.
	RawResponse complexes.JSON `json:"rawResponse" yaml:"rawResponse"`
	// The ledger entry that was created once this attempt succeeded. Empty until then.
	WalletTransaction WalletTransactionEntity `gorm:"foreignKey:WalletTransactionId;references:Id" json:"walletTransaction" yaml:"walletTransaction"`
	// When this attempt expires if not completed (gateway-dependent). Empty if the gateway doesn't impose one.
	ExpiresAt complexes.XDate `json:"expiresAt" yaml:"expiresAt"`
	// When this attempt reached a terminal status. Empty until then.
	CompletedAt complexes.XDate `json:"completedAt" yaml:"completedAt"`
	// Where to send the caller's browser back to once a redirect-based gateway (e.g. Przelewy24, ZarinPal) completes the payment. Not needed by gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow).
	ReturnUrl           emigo.Nullable[string] `json:"returnUrl" yaml:"returnUrl"`
	WalletId            int64                  `gorm:"index" json:"-" yaml:"-"`
	GatewayId           int64                  `gorm:"index" json:"-" yaml:"-"`
	WalletTransactionId int64                  `gorm:"index" json:"-" yaml:"-"`
}

func (x *WalletPaymentAttemptEntity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetWalletPaymentAttemptEntityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name:        prefix + "wallet",
			Type:        "class",
			Description: "The wallet this attempt would credit/debit if it succeeds.",
		},
		{
			Name:        prefix + "gateway",
			Type:        "class",
			Description: "The gateway this attempt is routed through.",
		},
		{
			Name:        prefix + "purpose",
			Type:        "string",
			Description: "What this attempt is for.",
		},
		{
			Name:        prefix + "amount",
			Type:        "string",
			Description: "Requested amount as a minor-units string, in currency.",
		},
		{
			Name:        prefix + "currency",
			Type:        "string",
			Description: "Currency code for amount - must match the wallet's currency.",
		},
		{
			Name:        prefix + "status",
			Type:        "string",
			Description: "Current lifecycle state of this attempt.",
		},
		{
			Name:        prefix + "gateway-reference",
			Type:        "string?",
			Description: "The gateway's own id for this attempt (e.g. a PaymentIntent id or a transaction hash), once known. Indexed for webhook lookups.",
		},
		{
			Name:        prefix + "idempotency-key",
			Type:        "string",
			Description: "Caller-supplied key making topup-initiation safe to retry without creating duplicate attempts at the gateway.",
		},
		{
			Name:        prefix + "created-at",
			Type:        "complex",
			Description: "When this attempt was created.",
		},
		{
			Name:        prefix + "failure-reason",
			Type:        "string?",
			Description: "Human-readable reason, populated when status is \"failed\".",
		},
		{
			Name:        prefix + "raw-request",
			Type:        "complex",
			Description: "The raw request sent to the gateway when initiating this attempt.",
		},
		{
			Name:        prefix + "raw-response",
			Type:        "complex",
			Description: "The raw response/init payload received back from the gateway.",
		},
		{
			Name:        prefix + "wallet-transaction",
			Type:        "class?",
			Description: "The ledger entry that was created once this attempt succeeded. Empty until then.",
		},
		{
			Name:        prefix + "expires-at",
			Type:        "complex",
			Description: "When this attempt expires if not completed (gateway-dependent). Empty if the gateway doesn't impose one.",
		},
		{
			Name:        prefix + "completed-at",
			Type:        "complex",
			Description: "When this attempt reached a terminal status. Empty until then.",
		},
		{
			Name:        prefix + "return-url",
			Type:        "string?",
			Description: "Where to send the caller's browser back to once a redirect-based gateway (e.g. Przelewy24, ZarinPal) completes the payment. Not needed by gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow).",
		},
		{
			Name: prefix + "wallet-id",
			Type: "int64",
		},
		{
			Name: prefix + "gateway-id",
			Type: "int64",
		},
		{
			Name: prefix + "wallet-transaction-id",
			Type: "int64",
		},
	}
}
func CastWalletPaymentAttemptEntityFromCli(c emigo.CliCastable) WalletPaymentAttemptEntity {
	data := WalletPaymentAttemptEntity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("purpose") {
		data.Purpose = c.String("purpose")
	}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("gateway-reference") {
		emigo.ParseNullable(c.String("gateway-reference"), &data.GatewayReference)
	}
	if c.IsSet("idempotency-key") {
		data.IdempotencyKey = c.String("idempotency-key")
	}
	if c.IsSet("created-at") {
		if u, ok := any(&data.CreatedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("created-at")))
		}
	}
	if c.IsSet("failure-reason") {
		emigo.ParseNullable(c.String("failure-reason"), &data.FailureReason)
	}
	if c.IsSet("raw-request") {
		if u, ok := any(&data.RawRequest).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("raw-request")))
		}
	}
	if c.IsSet("raw-response") {
		if u, ok := any(&data.RawResponse).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("raw-response")))
		}
	}
	if c.IsSet("expires-at") {
		if u, ok := any(&data.ExpiresAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("expires-at")))
		}
	}
	if c.IsSet("completed-at") {
		if u, ok := any(&data.CompletedAt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("completed-at")))
		}
	}
	if c.IsSet("return-url") {
		emigo.ParseNullable(c.String("return-url"), &data.ReturnUrl)
	}
	if c.IsSet("wallet-id") {
		data.WalletId = int64(c.Int64("wallet-id"))
	}
	if c.IsSet("gateway-id") {
		data.GatewayId = int64(c.Int64("gateway-id"))
	}
	if c.IsSet("wallet-transaction-id") {
		data.WalletTransactionId = int64(c.Int64("wallet-transaction-id"))
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// WalletPaymentAttemptEntityCreateFn creates a new WalletPaymentAttemptEntity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func WalletPaymentAttemptEntityCreateFn(tx *gorm.DB, dto *WalletPaymentAttemptEntity) (*WalletPaymentAttemptEntity, error) {
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(dto).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dto, nil
}

// WalletPaymentAttemptEntityUpdateFn applies a partial update to the WalletPaymentAttemptEntity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// WalletPaymentAttemptEntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func WalletPaymentAttemptEntityUpdateFn(tx *gorm.DB, uniqueId string, input WalletPaymentAttemptOptionalDto) (*WalletPaymentAttemptEntity, error) {
	var entity WalletPaymentAttemptEntity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Wallet.IsSet() {
			if input.Wallet.Operation != "select" {
				return fmt.Errorf("wallet: updating a one/one? relation only supports the \"select\" operation (link to an existing row by its uniqueId), got %q", input.Wallet.Operation)
			}
			var selectorId string
			if s, ok := input.Wallet.Selector.(string); ok {
				selectorId = s
			}
			resolvedId, err := emigorm.ReconcileOne[WalletEntity](tx, input.Wallet.Operation, selectorId, nil)
			if err != nil {
				return err
			}
			changes["WalletId"] = resolvedId
		}
		if input.Gateway.IsSet() {
			if input.Gateway.Operation != "select" {
				return fmt.Errorf("gateway: updating a one/one? relation only supports the \"select\" operation (link to an existing row by its uniqueId), got %q", input.Gateway.Operation)
			}
			var selectorId string
			if s, ok := input.Gateway.Selector.(string); ok {
				selectorId = s
			}
			resolvedId, err := emigorm.ReconcileOne[WalletGatewayEntity](tx, input.Gateway.Operation, selectorId, nil)
			if err != nil {
				return err
			}
			changes["GatewayId"] = resolvedId
		}
		if input.WalletTransaction.IsSet() {
			if input.WalletTransaction.Operation != "select" {
				return fmt.Errorf("walletTransaction: updating a one/one? relation only supports the \"select\" operation (link to an existing row by its uniqueId), got %q", input.WalletTransaction.Operation)
			}
			var selectorId string
			if s, ok := input.WalletTransaction.Selector.(string); ok {
				selectorId = s
			}
			resolvedId, err := emigorm.ReconcileOne[WalletTransactionEntity](tx, input.WalletTransaction.Operation, selectorId, nil)
			if err != nil {
				return err
			}
			changes["WalletTransactionId"] = resolvedId
		}
		if input.Purpose.IsSet() {
			changes["Purpose"] = input.Purpose
		}
		if input.Amount.IsSet() {
			changes["Amount"] = input.Amount
		}
		if input.Currency.IsSet() {
			changes["Currency"] = input.Currency
		}
		if input.Status.IsSet() {
			changes["Status"] = input.Status
		}
		if input.GatewayReference.IsSet() {
			changes["GatewayReference"] = input.GatewayReference
		}
		if input.IdempotencyKey.IsSet() {
			changes["IdempotencyKey"] = input.IdempotencyKey
		}
		changes["CreatedAt"] = input.CreatedAt
		if input.FailureReason.IsSet() {
			changes["FailureReason"] = input.FailureReason
		}
		changes["RawRequest"] = input.RawRequest
		changes["RawResponse"] = input.RawResponse
		changes["ExpiresAt"] = input.ExpiresAt
		changes["CompletedAt"] = input.CompletedAt
		if input.ReturnUrl.IsSet() {
			changes["ReturnUrl"] = input.ReturnUrl
		}
		if len(changes) > 0 {
			if err := tx.Model(&entity).Updates(changes).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var updated WalletPaymentAttemptEntity
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// WalletPaymentAttemptEntityGetFn looks up a single WalletPaymentAttemptEntity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func WalletPaymentAttemptEntityGetFn(tx *gorm.DB, uniqueId string) (*WalletPaymentAttemptEntity, error) {
	var entity WalletPaymentAttemptEntity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// WalletPaymentAttemptEntityBrowseFn returns WalletPaymentAttemptEntity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func WalletPaymentAttemptEntityBrowseFn(tx *gorm.DB, qs WalletPaymentAttemptBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletPaymentAttemptEntity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&WalletPaymentAttemptEntity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*WalletPaymentAttemptEntity
	paged := emigorm.ApplyQueryPage(emigorm.ApplyQueryCursor(emigorm.ApplyQuerySort(filtered, qs.Sort), qs.Cursor), qs.StartIndex, qs.ItemsPerPage)
	if err := paged.Find(&items).Error; err != nil {
		return nil, nil, err
	}
	meta := &emigo.QueryResultMeta{
		TotalItems: total,
		Cursor:     emigorm.BuildQueryCursor(items),
	}
	return items, meta, nil
}

// WalletPaymentAttemptEntityAwareDeleteAffected reports one relation of WalletPaymentAttemptEntity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on WalletPaymentAttemptEntity itself, so deleting WalletPaymentAttemptEntity doesn't cascade into them.
type WalletPaymentAttemptEntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// WalletPaymentAttemptEntityAwareDeletePreview is the result of WalletPaymentAttemptEntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts WalletPaymentAttemptEntityAwareDeleteFn would delete/clear
// alongside the WalletPaymentAttemptEntity row(s) themselves.
type WalletPaymentAttemptEntityAwareDeletePreview struct {
	Message  string                                          `json:"message"`
	Affected []WalletPaymentAttemptEntityAwareDeleteAffected `json:"affected"`
}

// WalletPaymentAttemptEntityAwareDeletePreviewFn looks up the WalletPaymentAttemptEntity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// WalletPaymentAttemptEntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling WalletPaymentAttemptEntityAwareDeleteFn.
func WalletPaymentAttemptEntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*WalletPaymentAttemptEntityAwareDeletePreview, error) {
	var rows []*WalletPaymentAttemptEntity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &WalletPaymentAttemptEntityAwareDeletePreview{Message: "No matching WalletPaymentAttemptEntity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []WalletPaymentAttemptEntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d WalletPaymentAttemptEntity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &WalletPaymentAttemptEntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// WalletPaymentAttemptEntityAwareDeleteFn deletes the WalletPaymentAttemptEntity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation WalletPaymentAttemptEntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func WalletPaymentAttemptEntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*WalletPaymentAttemptEntity
		if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
			return err
		}
		if len(rows) == 0 {
			return nil
		}
		ids := make([]int64, len(rows))
		for i := range rows {
			ids[i] = rows[i].Id
		}
		return tx.Where("id IN ?", ids).Delete(&WalletPaymentAttemptEntity{}).Error
	})
}

// WalletPaymentAttemptEntityActionsSig bundles the actions available for WalletPaymentAttemptEntity. Extend this (and
// WalletPaymentAttemptEntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type WalletPaymentAttemptEntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *WalletPaymentAttemptEntity) (*WalletPaymentAttemptEntity, error)
	Update             func(tx *gorm.DB, uniqueId string, input WalletPaymentAttemptOptionalDto) (*WalletPaymentAttemptEntity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*WalletPaymentAttemptEntity, error)
	Browse             func(tx *gorm.DB, qs WalletPaymentAttemptBrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*WalletPaymentAttemptEntity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*WalletPaymentAttemptEntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var WalletPaymentAttemptEntityActions WalletPaymentAttemptEntityActionsSig = WalletPaymentAttemptEntityActionsSig{
	Create:             WalletPaymentAttemptEntityCreateFn,
	Update:             WalletPaymentAttemptEntityUpdateFn,
	Get:                WalletPaymentAttemptEntityGetFn,
	Browse:             WalletPaymentAttemptEntityBrowseFn,
	AwareDeletePreview: WalletPaymentAttemptEntityAwareDeletePreviewFn,
	AwareDelete:        WalletPaymentAttemptEntityAwareDeleteFn,
}
